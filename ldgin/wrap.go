/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func Wrap(f Handler) func(c *gin.Context) {
	w := wrapHandler(f)
	return w.Do
}

type (
	inConvType  = func(Context) (reflect.Value, Error)
	outConvType = func(Context, []reflect.Value)
)

type handler struct {
	Type    reflect.Type
	Value   reflect.Value
	InConvs []inConvType
	outConv outConvType
}

func wrapHandler(f Handler) *handler {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("handler should be a function")
	}

	w := &handler{
		Type:  t,
		Value: reflect.ValueOf(f),
	}

	switch t.NumIn() {
	case 1:
		inType := t.In(0)
		if !w.isType(_TYPE_OF_CONTEXT, inType) && !w.isType(_TYPE_OF_GIN_CONTEXT, inType) {
			panic("handler input parameter type should be `ldgin.Context` or `*gin.Context`")
		}

	case 2:
		inType := t.In(0)
		if !w.isType(_TYPE_OF_CONTEXT, inType) && !w.isType(_TYPE_OF_GIN_CONTEXT, inType) {
			panic("handler first input parameter type should be `ldgin.Context` or `*gin.Context`")
		}

	default:
		panic("handler input parameter count should be 1 or 2")
	}

	ins := make([]inConvType, 0, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		t := t.In(i)
		conv := w.getInConv(t)
		ins = append(ins, conv)
	}
	w.InConvs = ins

	switch t.NumOut() {
	case 0:
		w.outConv = w.getOutConv0()
		break

	case 1:
		outType := t.Out(0)
		w.outConv = w.getOutConv1(outType)

	case 2:
		outTypes := []reflect.Type{t.Out(0), t.Out(1)}
		w.outConv = w.getOutConv2(outTypes)

	default:
		panic("handler output parameter count should be 0, 1 or 2")
	}

	return w
}

func (w *handler) returnError(c Context, e interface{}) {
	err, ok := e.(Error)
	if !ok {
		err = commError{
			message: e.(error).Error(),
			status:  http.StatusOK,
			code:    -1,
		}
	}
	response := &CommResponse{
		Sequence: c.Sequence(),
		Cost:     time.Since(c.BeginTime()).String(),
		ErrCode:  err.Code(),
		ErrMsg:   err.Error(),
		Data:     struct{}{},
	}

	c.Set(GIN_KEY_RESPONSE, response)
	c.JSON(err.Status(), response)
}

func (w *handler) returnResponse(c Context, rsp interface{}) {
	if rsp == nil {
		rsp = struct{}{}
	}

	response := &CommResponse{
		Sequence: c.Sequence(),
		Cost:     time.Since(c.BeginTime()).String(),
		Data:     rsp,
	}

	c.Set(GIN_KEY_RESPONSE, response)
	c.JSON(http.StatusOK, response)
}

func (w *handler) getOutConv0() outConvType {
	return func(c Context, v []reflect.Value) {}
}

func (w *handler) getOutConv1(outType reflect.Type) outConvType {
	errType := outType
	if !w.isType(errType, _TYPE_OF_ERROR) && !w.isType(errType, _TYPE_OF_COMM_ERROR) {
		panic("handler output parameter type should be `ldgin.Error` or `error`")
	}

	return func(c Context, outs []reflect.Value) {
		out0 := outs[0].Interface()
		if err := out0; err != nil {
			w.returnError(c, err)
			return
		}

		w.returnResponse(c, nil)
	}
}

func (w *handler) getOutConv2(outTypes []reflect.Type) outConvType {
	rspType := outTypes[0]
	errType := outTypes[1]

	if !w.isType(errType, _TYPE_OF_ERROR) && !w.isType(errType, _TYPE_OF_COMM_ERROR) {
		panic("handler second output parameter type should be `ldgin.Error` or `error`")
	}

	if w.isType(rspType, _TYPE_OF_RENDERER) {
		return func(c Context, outs []reflect.Value) {
			out0 := outs[0].Interface()
			out1 := outs[1].Interface()

			if err := out1; err != nil {
				w.returnError(c, err)
				return
			}

			c.Set(GIN_KEY_RENDERER, out0)

			render := out0.(Renderer)
			render.Render(c)
		}
	}

	if w.isType(rspType, _TYPE_OF_GIN_RENDERER) {
		return func(c Context, outs []reflect.Value) {
			out0 := outs[0].Interface()
			out1 := outs[1].Interface()

			if err := out1; err != nil {
				w.returnError(c, err)
				return
			}

			c.Set(GIN_KEY_RENDERER, out0)

			render := out0.(GinRenderer)
			render.Render(c.Gin())
		}
	}

	return func(c Context, outs []reflect.Value) {
		out0 := outs[0].Interface()
		out1 := outs[1].Interface()

		if err := out1; err != nil {
			w.returnError(c, err)
			return
		}

		w.returnResponse(c, out0)
	}
}

func (w *handler) getInConv(t reflect.Type) inConvType {
	switch {
	case w.isType(_TYPE_OF_CONTEXT, t):
		return func(c Context) (reflect.Value, Error) {
			return reflect.ValueOf(c), nil
		}

	case w.isType(_TYPE_OF_GIN_CONTEXT, t):
		return func(c Context) (reflect.Value, Error) {
			return reflect.ValueOf(c.Gin()), nil
		}
	}

	convs := make([]func(Context, reflect.Value) Error, 0, 2)
	convs = append(convs, w.getParserFunc(t))
	if f := w.getValidatorFunc(t); f != nil {
		convs = append(convs, f)
	}

	return func(c Context) (reflect.Value, Error) {
		v := reflect.New(t.Elem())

		for _, f := range convs {
			if err := f(c, v); err != nil {
				return v, err
			}
		}

		c.Set(GIN_KEY_REQUEST, v.Interface())
		return v, nil
	}
}
func (w *handler) isType(child, parent reflect.Type) bool {
	if child == parent {
		return true
	}
	if parent.Kind() == reflect.Interface && child.Implements(parent) {
		return true
	}
	return false
}

func (w *handler) getParserFunc(t reflect.Type) func(Context, reflect.Value) Error {
	if w.isType(t, _TYPE_OF_PARSER) {
		return func(c Context, v reflect.Value) Error {
			paser, _ := v.Interface().(Parser)
			return paser.Parse(c)
		}
	}

	if w.isType(t, _TYPE_OF_GIN_PARSER) {
		return func(c Context, v reflect.Value) Error {
			paser, _ := v.Interface().(GinParser)
			return paser.Parse(c.Gin())
		}
	}

	return func(c Context, v reflect.Value) Error {
		return shouldBind(c, v.Interface())
	}
}

func (w *handler) getValidatorFunc(t reflect.Type) func(Context, reflect.Value) Error {
	if w.isType(t, _TYPE_OF_VALIDATER) {
		return func(c Context, v reflect.Value) Error {
			validator, _ := v.Interface().(Validator)
			return validator.Validate(c)
		}
	}

	if w.isType(t, _TYPE_OF_GIN_VALIDATER) {
		return func(c Context, v reflect.Value) Error {
			validator, _ := v.Interface().(GinValidator)
			return validator.Validate(c.Gin())
		}
	}

	return nil
}

func (w *handler) Do(g *gin.Context) {
	c := GetContext(g)
	ins := make([]reflect.Value, 0, len(w.InConvs))
	for _, conv := range w.InConvs {
		v, err := conv(c)
		if err != nil {
			w.returnError(c, err)
			return
		}
		ins = append(ins, v)
	}

	outs := w.Value.Call(ins)
	if len(outs) == 0 {
		return
	}

	w.outConv(c, outs)
}
