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
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now
var _ = testing.Main

func testConvInt(t testing.TB, s string, r int64) {
	n, _ := convInt(StrToBytes(s))
	t.Logf("conv str to int. s:%s, n:%d", s, n)
}

func TestConvInt(t *testing.T) {
	testConvInt(t, "123", 123)
	testConvInt(t, "123.1", 123)
	testConvInt(t, ".1", 0)
	testConvInt(t, "0123", 0123)
	testConvInt(t, "+-0123", -0123)
	testConvInt(t, "123e+1", 1230)
	testConvInt(t, "123e++1", 1230)
	testConvInt(t, "123e-1", 12)
	testConvInt(t, "0x123", 0x123)
	testConvInt(t, "--0X123", 0x123)
	testConvInt(t, "-+-0x123", 0x123)
}
