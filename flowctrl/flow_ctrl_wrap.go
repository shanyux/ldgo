/*
 * Copyright (C) distroy
 */

package flowctrl

func Wrap(fcs []FlowCtrl) FlowCtrl {
	if len(fcs) == 1 {
		return fcs[0]
	}
	return flowCtrlWrap{fcs: fcs}
}

type flowCtrlWrap struct {
	fcs []FlowCtrl
}

func (that flowCtrlWrap) Close() {
	for _, fc := range that.fcs {
		fc.Close()
	}
}

func (that flowCtrlWrap) Apply(n int64) int64 {
	r := n
	for _, fc := range that.fcs {
		r = fc.Apply(r)
	}
	return r
}

func (that flowCtrlWrap) TryApply(n int64) int64 {
	r := n
	for _, fc := range that.fcs {
		r = fc.TryApply(r)
	}
	return r
}
