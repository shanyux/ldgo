/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestVariableHook(t *testing.T) {
	var _TestVariableHook *int
	convey.Convey(t.Name(), t, func() {
		v := 100

		patches := NewPatches()
		patches.Apply(VariableHook{
			Target: &_TestVariableHook,
			Double: &v,
		})

		convey.So(_TestVariableHook, convey.ShouldNotBeNil)
		convey.So(*_TestVariableHook, convey.ShouldEqual, v)

		patches.Reset()
		convey.So(_TestVariableHook, convey.ShouldBeNil)
	})
}
