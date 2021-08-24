/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// WrapHandler returns then gin Handler
func WrapMidware(f Midware) func(c *gin.Context) {
	w := wrapMidware(f)
	return w.Do
}

// WM is the alias for WrapHandler
func WM(f Midware) func(c *gin.Context) { return WrapMidware(f) }

type midware struct {
	wrapper

	Value reflect.Value
}

func wrapMidware(f Midware) *midware {
	t := reflect.TypeOf(f)
	w := &midware{
		wrapper: wrapper{
			Name: "midware",
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

	default:
		panicf("%s output parameter count should be 0 or 1", w.Name)
	}

	return w
}

func (w *midware) Do(g *gin.Context) {
	w.call(g, w.Value)
}
