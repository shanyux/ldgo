/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"

	"github.com/distroy/ldgo/lderr"
	"github.com/gin-gonic/gin"
)

// WrapHandler returns then gin Handler
func WrapHandler(f Handler) func(c *gin.Context) {
	w := wrapHandler(f)
	return w.Do
}

// Wrap is the alias for WrapHandler
func Wrap(f Handler) func(c *gin.Context) { return WrapHandler(f) }

// WH is the alias for WrapHandler
func WH(f Handler) func(c *gin.Context) { return WrapHandler(f) }

type handler struct {
	wrapper

	Value reflect.Value
}

func wrapHandler(f Handler) *handler {
	t := reflect.TypeOf(f)
	w := &handler{
		wrapper: wrapper{
			Name: "handler",
			Type: t,
		},
		Value: reflect.ValueOf(f),
	}

	if t.Kind() != reflect.Func {
		panicf("%s should be a function", w.Name)
	}

	w.InConvs = w.getAllInConvs(t)

	switch t.NumOut() {
	case 0:
		w.OutConv = w.getOutConv0()
		break

	case 1:
		outType := t.Out(0)
		w.OutConv = w.getOutConv1(outType)

	case 2:
		outTypes := []reflect.Type{t.Out(0), t.Out(1)}
		w.OutConv = w.getOutConv2(outTypes)

	default:
		panicf("%s output parameter count should be 0, 1 or 2", w.Name)
	}

	return w
}

func (w *handler) getOutConv1(outType reflect.Type) outConvType {
	errType := outType
	if !w.isType(errType, _TYPE_OF_ERROR) && !w.isType(errType, _TYPE_OF_COMM_ERROR) {
		panicf("%s output parameter type should be `ldgin.Error` or `error`", w.Name)
	}

	return func(c Context, outs []reflect.Value) {
		out0 := outs[0].Interface()
		if err := out0; err != nil {
			w.returnError(c, lderr.Wrap(err.(Error)))
			return
		}

		w.returnResponse(c, nil)
	}
}

func (w *handler) getOutConv2(outTypes []reflect.Type) outConvType {
	rspType := outTypes[0]
	errType := outTypes[1]

	if !w.isType(errType, _TYPE_OF_ERROR) && !w.isType(errType, _TYPE_OF_COMM_ERROR) {
		panicf("%s second output parameter type should be `ldgin.Error` or `error`", w.Name)
	}

	if w.isType(rspType, _TYPE_OF_RENDERER) {
		return func(c Context, outs []reflect.Value) {
			out0 := outs[0].Interface()
			out1 := outs[1].Interface()

			if err := out1; err != nil {
				w.returnError(c, lderr.Wrap(err.(Error)))
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
				w.returnError(c, lderr.Wrap(err.(Error)))
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
			w.returnError(c, lderr.Wrap(err.(Error)))
			return
		}

		w.returnResponse(c, out0)
	}
}

func (w *handler) Do(g *gin.Context) {
	w.call(g, w.Value)
}
