/*
 * Copyright (C) distroy
 */

package ldatomic

import "sync/atomic"

// A Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value = atomic.Value

func NewValue(d interface{}) *Value {
	p := &Value{}
	if d != nil {
		p.Store(d)
	}
	return p
}
