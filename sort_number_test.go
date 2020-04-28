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

func TestSortInt64(t *testing.T) {
	l := []int64{223, 562, 424, 642, 123, 496, 623, 845, 375}
	if IsSortedInt64(l) {
		t.Fatal("is sorted: ", l)
	}

	SortInt64(l)
	if !IsSortedInt64(l) {
		t.Fatal("is not sorted: ", l)
	}
	t.Log("sorted: ", l)

	n := SearchInt64(l, 223)
	t.Log("223 pos: ", n)

	n = SearchInt64(l, 300)
	t.Log("300 pos: ", n)
}
