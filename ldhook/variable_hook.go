/*
 * Copyright (C) distroy
 */

package ldhook

import "github.com/agiledragon/gomonkey"

type VariableHook struct {
	Target interface{}
	Double interface{}
}

func (h VariableHook) hook(patches *gomonkey.Patches) {
	patches.ApplyGlobalVar(h.Target, h.Double)
}
