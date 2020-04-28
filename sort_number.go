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
	"sort"
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

type SortSliceInt []int

func (s SortSliceInt) Len() int           { return len(s) }
func (s SortSliceInt) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceInt) Less(i, j int) bool { return s[i] < s[j] }

func SortInt(a []int)          { sort.Sort(SortSliceInt(a)) }
func IsSortedInt(a []int) bool { return sort.IsSorted(SortSliceInt(a)) }
func SearchInt(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceInt64 []int64

func (s SortSliceInt64) Len() int           { return len(s) }
func (s SortSliceInt64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceInt64) Less(i, j int) bool { return s[i] < s[j] }

func SortInt64(a []int64)          { sort.Sort(SortSliceInt64(a)) }
func IsSortedInt64(a []int64) bool { return sort.IsSorted(SortSliceInt64(a)) }
func SearchInt64(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceInt32 []int32

func (s SortSliceInt32) Len() int           { return len(s) }
func (s SortSliceInt32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceInt32) Less(i, j int) bool { return s[i] < s[j] }

func SortInt32(a []int32)          { sort.Sort(SortSliceInt32(a)) }
func IsSortedInt32(a []int32) bool { return sort.IsSorted(SortSliceInt32(a)) }
func SearchInt32(a []int32, x int32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceInt16 []int16

func (s SortSliceInt16) Len() int           { return len(s) }
func (s SortSliceInt16) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceInt16) Less(i, j int) bool { return s[i] < s[j] }

func SortInt16(a []int16)          { sort.Sort(SortSliceInt16(a)) }
func IsSortedInt16(a []int16) bool { return sort.IsSorted(SortSliceInt16(a)) }
func SearchInt16(a []int16, x int16) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceInt8 []int8

func (s SortSliceInt8) Len() int           { return len(s) }
func (s SortSliceInt8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceInt8) Less(i, j int) bool { return s[i] < s[j] }

func SortInt8(a []int8)          { sort.Sort(SortSliceInt8(a)) }
func IsSortedInt8(a []int8) bool { return sort.IsSorted(SortSliceInt8(a)) }
func SearchInt8(a []int8, x int8) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceUint []uint

func (s SortSliceUint) Len() int           { return len(s) }
func (s SortSliceUint) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceUint) Less(i, j int) bool { return s[i] < s[j] }

func SortUint(a []uint)          { sort.Sort(SortSliceUint(a)) }
func IsSortedUint(a []uint) bool { return sort.IsSorted(SortSliceUint(a)) }
func SearchUint(a []uint, x uint) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceUint64 []uint64

func (s SortSliceUint64) Len() int           { return len(s) }
func (s SortSliceUint64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceUint64) Less(i, j int) bool { return s[i] < s[j] }

func SortUint64(a []uint64)          { sort.Sort(SortSliceUint64(a)) }
func IsSortedUint64(a []uint64) bool { return sort.IsSorted(SortSliceUint64(a)) }
func SearchUint64(a []uint64, x uint64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceUint32 []uint32

func (s SortSliceUint32) Len() int           { return len(s) }
func (s SortSliceUint32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceUint32) Less(i, j int) bool { return s[i] < s[j] }

func SortUint32(a []uint32)          { sort.Sort(SortSliceUint32(a)) }
func IsSortedUint32(a []uint32) bool { return sort.IsSorted(SortSliceUint32(a)) }
func SearchUint32(a []uint32, x uint32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceUint16 []uint16

func (s SortSliceUint16) Len() int           { return len(s) }
func (s SortSliceUint16) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceUint16) Less(i, j int) bool { return s[i] < s[j] }

func SortUint16(a []uint16)          { sort.Sort(SortSliceUint16(a)) }
func IsSortedUint16(a []uint16) bool { return sort.IsSorted(SortSliceUint16(a)) }
func SearchUint16(a []uint16, x uint16) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceUint8 []uint8

func (s SortSliceUint8) Len() int           { return len(s) }
func (s SortSliceUint8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceUint8) Less(i, j int) bool { return s[i] < s[j] }

func SortUint8(a []uint8)          { sort.Sort(SortSliceUint8(a)) }
func IsSortedUint8(a []uint8) bool { return sort.IsSorted(SortSliceUint8(a)) }
func SearchUint8(a []uint8, x uint8) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceFloat64 []float64

func (s SortSliceFloat64) Len() int           { return len(s) }
func (s SortSliceFloat64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceFloat64) Less(i, j int) bool { return s[i] < s[j] }

func SortFloat64(a []float64)          { sort.Sort(SortSliceFloat64(a)) }
func IsSortedFloat64(a []float64) bool { return sort.IsSorted(SortSliceFloat64(a)) }
func SearchFloat64(a []float64, x float64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type SortSliceFloat32 []float32

func (s SortSliceFloat32) Len() int           { return len(s) }
func (s SortSliceFloat32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortSliceFloat32) Less(i, j int) bool { return s[i] < s[j] }

func SortFloat32(a []float32)          { sort.Sort(SortSliceFloat32(a)) }
func IsSortedFloat32(a []float32) bool { return sort.IsSorted(SortSliceFloat32(a)) }
func SearchFloat32(a []float32, x float32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}
