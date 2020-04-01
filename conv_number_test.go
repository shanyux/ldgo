/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = context.Canceled
var _ = errors.New
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now
var _ = testing.Main

func testConvInt(t testing.TB, s string, r int64) {
	n, _ := convInt(StrToBytesUnsafe(s))
	if n == r {
		t.Logf("conv str to int succ. s:%s, n:%d", s, n)
	} else {
		t.Errorf("conv str to int fail. s:%s, n:%d, expected:%d", s, n, r)
	}
}

func testConvUint(t testing.TB, s string, r uint64) {
	n, _ := convUint(StrToBytesUnsafe(s))
	if n == r {
		t.Logf("conv str to uint succ. s:%s, n:%d", s, n)
	} else {
		t.Errorf("conv str to uint fail. s:%s, n:%d, expected:%d", s, n, r)
	}
}

func testConvFloat(t testing.TB, s string, r float64) {
	n, _ := convFloat(StrToBytesUnsafe(s))
	if n == r {
		t.Logf("conv str to float succ. s:%s, n:%s", s, strconv.FormatFloat(n, 'f', -1, 64))
	} else {
		t.Errorf("conv str to float fail. s:%s, n:%s, expected:%s", s,
			strconv.FormatFloat(n, 'f', -1, 64), strconv.FormatFloat(r, 'f', -1, 64))
	}
}

func TestConvInt(t *testing.T) {
	testConvInt(t, "123", 123)
	testConvInt(t, "123.1", 123)
	testConvInt(t, ".1", 0)
	testConvInt(t, "0123", 0123)
	testConvInt(t, "+-0123", -0123)
	testConvInt(t, "123e+1", 1230)
	testConvInt(t, "123e1", 1230)
	testConvInt(t, "123e-1", 12)
	testConvInt(t, "0x123", 0x123)
	testConvInt(t, "--0X123", 0x123)
	testConvInt(t, "-+-0x123", 0x123)
	testConvInt(t, "-+-0xeFaB", 0xeFaB)
}

func TestConvUint(t *testing.T) {
	testConvUint(t, "123", 123)
	testConvUint(t, "123.1", 123)
	testConvUint(t, ".1", 0)
	testConvUint(t, "0123", 0123)
	// testConvUint(t, "+-0123", uint64(int64(-0123)))
	testConvUint(t, "123e+1", 1230)
	testConvUint(t, "123e1", 1230)
	testConvUint(t, "123e-1", 12)
	testConvUint(t, "0x123", 0x123)
	testConvUint(t, "--0X123", 0x123)
	testConvUint(t, "-+-0x123", 0x123)
	testConvUint(t, "-+-0xeFaB", 0xeFaB)
}

func TestConvFloat(t *testing.T) {
	testConvFloat(t, "123", 123)
	testConvFloat(t, "123.1", 123.1)
	testConvFloat(t, ".1", 0.1)
	testConvFloat(t, "0123", 0123)
	// testConvFloat(t, "+-0123", uint64(int64(-0123)))
	testConvFloat(t, "123e+1", 1230)
	testConvFloat(t, "123e1", 1230)
	testConvFloat(t, "123e-1", 12.3)
	testConvFloat(t, "0x123", 0x123)
	testConvFloat(t, "--0X123", 0x123)
	testConvFloat(t, "-+-0x123", 0x123)
	testConvFloat(t, "-+-0xeFaB", 0xeFaB)
	testConvFloat(t, "0.30129", 0.30129)
	testConvFloat(t, "0.30129e3", 301.29)
}

func TestXXX(t *testing.T) {
}
