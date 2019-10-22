/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"bytes"
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
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now
var _ = testing.Main

func TestMax(t *testing.T) {
	t.Logf("MaxInt(3, 4) = %d", MaxInt(3, 4))
	t.Logf("MaxInt8(3, 4) = %d", MaxInt8(3, 4))
}

func TestMin(t *testing.T) {
	t.Logf("MinInt(3, 4) = %d", MinInt(3, 4))
	t.Logf("MinInt8(3, 4) = %d", MinInt8(3, 4))
}
