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

type sortIntSlice []int

func (s sortIntSlice) Len() int           { return len(s) }
func (s sortIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortIntSlice) Less(i, j int) bool { return s[i] < s[j] }

func SortInt(l []int)          { sort.Sort(sortIntSlice(l)) }
func IsSortedInt(l []int) bool { return sort.IsSorted(sortIntSlice(l)) }
func SearchInt(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortInt64Slice []int64

func (s sortInt64Slice) Len() int           { return len(s) }
func (s sortInt64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortInt64Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortInt64(l []int64)          { sort.Sort(sortInt64Slice(l)) }
func IsSortedInt64(l []int64) bool { return sort.IsSorted(sortInt64Slice(l)) }
func SearchInt64(a []int64, x int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortInt32Slice []int32

func (s sortInt32Slice) Len() int           { return len(s) }
func (s sortInt32Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortInt32Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortInt32(l []int32)          { sort.Sort(sortInt32Slice(l)) }
func IsSortedInt32(l []int32) bool { return sort.IsSorted(sortInt32Slice(l)) }
func SearchInt32(a []int32, x int32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortInt16Slice []int16

func (s sortInt16Slice) Len() int           { return len(s) }
func (s sortInt16Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortInt16Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortInt16(l []int16)          { sort.Sort(sortInt16Slice(l)) }
func IsSortedInt16(l []int16) bool { return sort.IsSorted(sortInt16Slice(l)) }
func SearchInt16(a []int16, x int16) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortInt8Slice []int8

func (s sortInt8Slice) Len() int           { return len(s) }
func (s sortInt8Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortInt8Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortInt8(l []int8)          { sort.Sort(sortInt8Slice(l)) }
func IsSortedInt8(l []int8) bool { return sort.IsSorted(sortInt8Slice(l)) }
func SearchInt8(a []int8, x int8) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortUintSlice []uint

func (s sortUintSlice) Len() int           { return len(s) }
func (s sortUintSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortUintSlice) Less(i, j int) bool { return s[i] < s[j] }

func SortUint(l []uint)          { sort.Sort(sortUintSlice(l)) }
func IsSortedUint(l []uint) bool { return sort.IsSorted(sortUintSlice(l)) }
func SearchUint(a []uint, x uint) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortUint64Slice []uint64

func (s sortUint64Slice) Len() int           { return len(s) }
func (s sortUint64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortUint64Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortUint64(l []uint64)          { sort.Sort(sortUint64Slice(l)) }
func IsSortedUint64(l []uint64) bool { return sort.IsSorted(sortUint64Slice(l)) }
func SearchUint64(a []uint64, x uint64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortUint32Slice []uint32

func (s sortUint32Slice) Len() int           { return len(s) }
func (s sortUint32Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortUint32Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortUint32(l []uint32)          { sort.Sort(sortUint32Slice(l)) }
func IsSortedUint32(l []uint32) bool { return sort.IsSorted(sortUint32Slice(l)) }
func SearchUint32(a []uint32, x uint32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortUint16Slice []uint16

func (s sortUint16Slice) Len() int           { return len(s) }
func (s sortUint16Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortUint16Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortUint16(l []uint16)          { sort.Sort(sortUint16Slice(l)) }
func IsSortedUint16(l []uint16) bool { return sort.IsSorted(sortUint16Slice(l)) }
func SearchUint16(a []uint16, x uint16) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortUint8Slice []uint8

func (s sortUint8Slice) Len() int           { return len(s) }
func (s sortUint8Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortUint8Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortUint8(l []uint8)          { sort.Sort(sortUint8Slice(l)) }
func IsSortedUint8(l []uint8) bool { return sort.IsSorted(sortUint8Slice(l)) }
func SearchUint8(a []uint8, x uint8) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortFloat64Slice []float64

func (s sortFloat64Slice) Len() int           { return len(s) }
func (s sortFloat64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortFloat64Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortFloat64(l []float64)          { sort.Sort(sortFloat64Slice(l)) }
func IsSortedFloat64(l []float64) bool { return sort.IsSorted(sortFloat64Slice(l)) }
func SearchFloat64(a []float64, x float64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

type sortFloat32Slice []float32

func (s sortFloat32Slice) Len() int           { return len(s) }
func (s sortFloat32Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortFloat32Slice) Less(i, j int) bool { return s[i] < s[j] }

func SortFloat32(l []float32)          { sort.Sort(sortFloat32Slice(l)) }
func IsSortedFloat32(l []float32) bool { return sort.IsSorted(sortFloat32Slice(l)) }
func SearchFloat32(a []float32, x float32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}
