/*
 * Copyright (C) distroy
 */

package flowctrl

func newFlowCtrlCall() *flowCtrlCall {
	return &flowCtrlCall{
		ch: make(chan struct{}, 1),
	}
}

type flowCtrlCall struct {
	ch chan struct{}
	n  int64
}
