/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"
)

func WrapGinWeb(router interface{}) Router {
	return wrapRouter(wrapGinWebRouter(router))
}

type ginWebWapper struct {
	ginWapperBase

	router reflect.Value
	method reflect.Method
}

func wrapGinWebRouter(r interface{}) *ginWebWapper {
	routerV := reflect.ValueOf(r)
	routerT := routerV.Type()

	method, ok := routerT.MethodByName("Handle")
	if !ok {
		panicf("gin web router must have the method `Handle`")
	}

	w := &ginWebWapper{
		router: routerV,
		method: method,
	}

	methodT := method.Type
	methodName := method.Name

	numIn := methodT.NumIn()
	if numIn != 5 {
		panicf("input parameter count of router handle must be 5. %s", methodName)
	}

	if !methodT.IsVariadic() {
		panicf("final input parameter of router handle must be an ...parameter. %s", methodName)
	}

	if methodT.In(1).Kind() != reflect.String && methodT.In(2).Kind() != reflect.String {
		panicf("1st and 2nd input parameter of router handle must be `string`. %s", methodName)
	}

	if typ := methodT.In(3); !typeOfGinHandlerFunc.ConvertibleTo(typ) {
		panicf("3rd input parameter of router handle must be `interface{}`. %s", methodName)
	}

	if typ := methodT.In(4); !typeOfGinHandlersChain.ConvertibleTo(typ) {
		panicf("3rd input parameter of router handle must be `[]func(*gin.Context)`. %s", methodName)
	}

	return w
}

func (w *ginWebWapper) clone() *ginWebWapper {
	c := *w
	return &c
}

func (w *ginWebWapper) WithAppPath(path string) routerAdapter {
	w = w.clone()
	w.setAppPath(path)
	return w
}

func (w *ginWebWapper) Group(relativePath string, midwares ...Midware) routerAdapter {
	w = w.clone()
	w.basePath = w.calculateAbsolutePath(relativePath)
	w.midwares = w.combineMidwares(midwares)
	return w
}

func (w *ginWebWapper) Use(midwares ...Midware) routerAdapter {
	w = w.clone()
	w.midwares = w.combineMidwares(midwares)
	return w
}

func (w *ginWebWapper) Handle(method, path string, h Handler, ms ...Midware) routerAdapter {
	fullPath := w.calculateFullPath(path)
	absPath := w.calculateAbsolutePath(path)

	handler := wrapHandler(h)
	handler.setMethod(method)
	handler.setPath(fullPath)

	midwares := w.combineMidwares(ms).WithMethod(method, fullPath).Get()

	ins := make([]reflect.Value, 0, len(midwares)+4)
	ins = append(ins, w.router)
	ins = append(ins, reflect.ValueOf(method))
	ins = append(ins, reflect.ValueOf(absPath))
	ins = append(ins, reflect.ValueOf(handler.Do))
	for _, m := range midwares {
		ins = append(ins, reflect.ValueOf(m))
	}
	w.method.Func.Call(ins[:])
	return w
}
