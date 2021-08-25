/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/distroy/ldgo/lderr"
	"github.com/gin-gonic/gin"
)

type (
	inConvType  = func(Context) (reflect.Value, Error)
	outConvType = func(Context, []reflect.Value)
)

// These routines end in 'f' and take a format string.

func panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}

// These routines do not take a format string

type wrapper struct {
	Name    string
	Type    reflect.Type
	InConvs []inConvType
	OutConv outConvType
}

func (w *wrapper) returnError(c Context, err Error) {
	response := &CommResponse{
		Sequence: c.Sequence(),
		Cost:     time.Since(c.GetBeginTime()).String(),
		ErrCode:  err.Code(),
		ErrMsg:   err.Error(),
		Data:     struct{}{},
	}

	c.Set(GIN_KEY_RESPONSE, response)
	c.JSON(err.Status(), response)
}

func (w *wrapper) returnResponse(c Context, rsp interface{}) {
	if rsp == nil {
		rsp = struct{}{}
	}

	response := &CommResponse{
		Sequence: c.Sequence(),
		Cost:     time.Since(c.GetBeginTime()).String(),
		Data:     rsp,
	}

	c.Set(GIN_KEY_RESPONSE, response)
	c.JSON(http.StatusOK, response)
}

func (w *wrapper) getOutConv0() outConvType {
	return func(c Context, v []reflect.Value) {}
}

func (w *wrapper) getOutConv1(outType reflect.Type) outConvType {
	errType := outType
	if !w.isType(errType, _TYPE_OF_ERROR) && !w.isType(errType, _TYPE_OF_COMM_ERROR) {
		panicf("%s output parameter type should be `ldgin.Error` or `error`", w.Name)
	}

	return func(c Context, outs []reflect.Value) {
		out0 := outs[0].Interface()
		if err := out0; err != nil {
			w.returnError(c, lderr.Wrap(err.(error)))
			return
		}
	}
}

func (w *wrapper) getAllInConvs(t reflect.Type) []inConvType {
	switch t.NumIn() {
	case 1:
		inType := t.In(0)
		if !w.isType(_TYPE_OF_CONTEXT, inType) && !w.isType(_TYPE_OF_GIN_CONTEXT, inType) {
			panicf("%s input parameter type should be `ldgin.Context` or `*gin.Context`", w.Name)
		}

	case 2:
		inType := t.In(0)
		if !w.isType(_TYPE_OF_CONTEXT, inType) && !w.isType(_TYPE_OF_GIN_CONTEXT, inType) {
			panicf("%s first input parameter type should be `ldgin.Context` or `*gin.Context`", w.Name)
		}

	default:
		panicf("%s input parameter count should be 1 or 2", w.Name)
	}

	ins := make([]inConvType, 0, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		t := t.In(i)
		conv := w.getInConv(t)
		ins = append(ins, conv)
	}

	return ins
}

func (w *wrapper) getInConv(t reflect.Type) inConvType {
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

func (w *wrapper) isType(child, parent reflect.Type) bool {
	if child == parent {
		return true
	}
	if parent.Kind() == reflect.Interface && child.Implements(parent) {
		return true
	}
	return false
}

func (w *wrapper) getParserFunc(t reflect.Type) func(Context, reflect.Value) Error {
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

func (w *wrapper) getValidatorFunc(t reflect.Type) func(Context, reflect.Value) Error {
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

func (w *wrapper) call(g *gin.Context, h reflect.Value) {
	c := GetContext(g)
	defer func() {
		if e := recover(); e != nil {
			err, _ := e.(error)
			if err == nil {
				err = fmt.Errorf("%v", e)
			}
			w.returnError(c, lderr.Wrap(err))
		}
	}()

	ins := make([]reflect.Value, 0, len(w.InConvs))
	for _, conv := range w.InConvs {
		v, err := conv(c)
		if err != nil {
			w.returnError(c, err)
			return
		}
		ins = append(ins, v)
	}

	outs := h.Call(ins)
	w.OutConv(c, outs)
}
