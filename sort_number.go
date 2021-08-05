/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldsort"
)

func SortInts(a []int)              { ldsort.SortInts(a) }
func IsSortedInts(a []int) bool     { return ldsort.IsSortedInts(a) }
func SearchInts(a []int, x int) int { return ldsort.SearchInts(a, x) }

func SortInt64s(a []int64)                { ldsort.SortInt64s(a) }
func IsSortedInt64s(a []int64) bool       { return ldsort.IsSortedInt64s((a)) }
func SearchInt64s(a []int64, x int64) int { return ldsort.SearchInt64s(a, x) }

func SortInt32s(a []int32)                { ldsort.SortInt32s(a) }
func IsSortedInt32s(a []int32) bool       { return ldsort.IsSortedInt32s(a) }
func SearchInt32s(a []int32, x int32) int { return ldsort.SearchInt32s(a, x) }

func SortInt16s(a []int16)                { ldsort.SortInt16s(a) }
func IsSortedInt16s(a []int16) bool       { return ldsort.IsSortedInt16s(a) }
func SearchInt16s(a []int16, x int16) int { return ldsort.SearchInt16s(a, x) }

func SortInt8s(a []int8)               { ldsort.SortInt8s(a) }
func IsSortedInt8s(a []int8) bool      { return ldsort.IsSortedInt8s(a) }
func SearchInt8s(a []int8, x int8) int { return ldsort.SearchInt8s(a, x) }

func SortUints(a []uint)               { ldsort.SortUints(a) }
func IsSortedUints(a []uint) bool      { return ldsort.IsSortedUints(a) }
func SearchUints(a []uint, x uint) int { return ldsort.SearchUints(a, x) }

func SortUint64s(a []uint64)                 { ldsort.SortUint64s(a) }
func IsSortedUint64s(a []uint64) bool        { return ldsort.IsSortedUint64s(a) }
func SearchUint64s(a []uint64, x uint64) int { return ldsort.SearchUint64s(a, x) }

func SortUint32s(a []uint32)                 { ldsort.SortUint32s(a) }
func IsSortedUint32s(a []uint32) bool        { return ldsort.IsSortedUint32s(a) }
func SearchUint32s(a []uint32, x uint32) int { return ldsort.SearchUint32s(a, x) }

func SortUint16s(a []uint16)                 { ldsort.SortUint16s(a) }
func IsSortedUint16s(a []uint16) bool        { return ldsort.IsSortedUint16s(a) }
func SearchUint16s(a []uint16, x uint16) int { return ldsort.SearchUint16s(a, x) }

func SortUint8s(a []uint8)                { ldsort.SortUint8s(a) }
func IsSortedUint8s(a []uint8) bool       { return ldsort.IsSortedUint8s(a) }
func SearchUint8s(a []uint8, x uint8) int { return ldsort.SearchUint8s(a, x) }

func SortFloat64s(a []float64)                  { ldsort.SortFloat64s(a) }
func IsSortedFloat64s(a []float64) bool         { return ldsort.IsSortedFloat64s(a) }
func SearchFloat64s(a []float64, x float64) int { return ldsort.SearchFloat64s(a, x) }

func SortFloat32s(a []float32)                  { ldsort.SortFloat32s(a) }
func IsSortedFloat32s(a []float32) bool         { return ldsort.IsSortedFloat32s(a) }
func SearchFloat32s(a []float32, x float32) int { return ldsort.SearchFloat32s(a, x) }
