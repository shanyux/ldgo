/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"testing"
)

func TestMax(t *testing.T) {
	t.Logf("MaxInt(3, 4) = %d", MaxInt(3, 4))
	t.Logf("MaxInt8(3, 4) = %d", MaxInt8(3, 4))
}

func TestMin(t *testing.T) {
	t.Logf("MinInt(3, 4) = %d", MinInt(3, 4))
	t.Logf("MinInt8(3, 4) = %d", MinInt8(3, 4))
}
