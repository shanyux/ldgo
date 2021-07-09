/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

type SortSliceInt = ldcore.SortSliceInt

func SortInts(a []int)              { ldcore.SortInts(a) }
func IsSortedInts(a []int) bool     { return ldcore.IsSortedInts(a) }
func SearchInts(a []int, x int) int { return ldcore.SearchInts(a, x) }

type SortSliceInt64 = ldcore.SortSliceInt64

func SortInt64s(a []int64)                { ldcore.SortInt64s(a) }
func IsSortedInt64s(a []int64) bool       { return ldcore.IsSortedInt64s((a)) }
func SearchInt64s(a []int64, x int64) int { return ldcore.SearchInt64s(a, x) }

type SortSliceInt32 = ldcore.SortSliceInt32

func SortInt32s(a []int32)                { ldcore.SortInt32s(a) }
func IsSortedInt32s(a []int32) bool       { return ldcore.IsSortedInt32s(a) }
func SearchInt32s(a []int32, x int32) int { return ldcore.SearchInt32s(a, x) }

type SortSliceInt16 = ldcore.SortSliceInt16

func SortInt16s(a []int16)                { ldcore.SortInt16s(a) }
func IsSortedInt16s(a []int16) bool       { return ldcore.IsSortedInt16s(a) }
func SearchInt16s(a []int16, x int16) int { return ldcore.SearchInt16s(a, x) }

type SortSliceInt8 = ldcore.SortSliceInt8

func SortInt8s(a []int8)               { ldcore.SortInt8s(a) }
func IsSortedInt8s(a []int8) bool      { return ldcore.IsSortedInt8s(a) }
func SearchInt8s(a []int8, x int8) int { return ldcore.SearchInt8s(a, x) }

type SortSliceUint = ldcore.SortSliceUint

func SortUints(a []uint)               { ldcore.SortUints(a) }
func IsSortedUints(a []uint) bool      { return ldcore.IsSortedUints(a) }
func SearchUints(a []uint, x uint) int { return ldcore.SearchUints(a, x) }

type SortSliceUint64 = ldcore.SortSliceUint64

func SortUint64s(a []uint64)                 { ldcore.SortUint64s(a) }
func IsSortedUint64s(a []uint64) bool        { return ldcore.IsSortedUint64s(a) }
func SearchUint64s(a []uint64, x uint64) int { return ldcore.SearchUint64s(a, x) }

type SortSliceUint32 = ldcore.SortSliceUint32

func SortUint32s(a []uint32)                 { ldcore.SortUint32s(a) }
func IsSortedUint32s(a []uint32) bool        { return ldcore.IsSortedUint32s(a) }
func SearchUint32s(a []uint32, x uint32) int { return ldcore.SearchUint32s(a, x) }

type SortSliceUint16 = ldcore.SortSliceUint16

func SortUint16s(a []uint16)                 { ldcore.SortUint16s(a) }
func IsSortedUint16s(a []uint16) bool        { return ldcore.IsSortedUint16s(a) }
func SearchUint16s(a []uint16, x uint16) int { return ldcore.SearchUint16s(a, x) }

type SortSliceUint8 = ldcore.SortSliceUint8

func SortUint8s(a []uint8)                { ldcore.SortUint8s(a) }
func IsSortedUint8s(a []uint8) bool       { return ldcore.IsSortedUint8s(a) }
func SearchUint8s(a []uint8, x uint8) int { return ldcore.SearchUint8s(a, x) }

type SortSliceFloat64 = ldcore.SortSliceFloat64

func SortFloat64s(a []float64)                  { ldcore.SortFloat64s(a) }
func IsSortedFloat64s(a []float64) bool         { return ldcore.IsSortedFloat64s(a) }
func SearchFloat64s(a []float64, x float64) int { return ldcore.SearchFloat64s(a, x) }

type SortSliceFloat32 = ldcore.SortSliceFloat32

func SortFloat32s(a []float32)                  { ldcore.SortFloat32s(a) }
func IsSortedFloat32s(a []float32) bool         { return ldcore.IsSortedFloat32s(a) }
func SearchFloat32s(a []float32, x float32) int { return ldcore.SearchFloat32s(a, x) }
