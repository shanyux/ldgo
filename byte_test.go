/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"testing"
)

func TestByteToUpper(t *testing.T) {
	b0 := byte('a')
	b1 := byte('A')
	for i := byte(0); i < 'z'-'a'; i++ {
		c0 := b0 + i
		c1 := b1 + i
		r0 := ToUpper(c0)
		if r0 != c1 {
			t.Errorf("ToUpper(%c) = %c", c0, r0)
		} else {
			t.Logf("ToUpper(%c) = %c", c0, r0)
		}
	}
}

func TestByteToLower(t *testing.T) {
	b0 := byte('A')
	b1 := byte('a')
	for i := byte(0); i < 'z'-'a'; i++ {
		c0 := b0 + i
		c1 := b1 + i
		r0 := ToLower(c0)
		if r0 != c1 {
			t.Errorf("ToLower(%c) = %c", c0, r0)
		} else {
			t.Logf("ToLower(%c) = %c", c0, r0)
		}
	}
}
