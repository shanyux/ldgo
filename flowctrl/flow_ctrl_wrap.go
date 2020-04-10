/*
 * Copyright (C) distroy
 */

package flowctrl

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

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
