/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldtopk"

func TopkIntsAdd(b []int, k int, x int) ([]int, bool) {
	return ldtopk.TopkIntsAdd(b, k, x)
}
func TopkInt64sAdd(b []int64, k int, x int64) ([]int64, bool) {
	return ldtopk.TopkInt64sAdd(b, k, x)
}
func TopkInt32sAdd(b []int32, k int, x int32) ([]int32, bool) {
	return ldtopk.TopkInt32sAdd(b, k, x)
}
func TopkInt16sAdd(b []int16, k int, x int16) ([]int16, bool) {
	return ldtopk.TopkInt16sAdd(b, k, x)
}
func TopkInt8sAdd(b []int8, k int, x int8) ([]int8, bool) {
	return ldtopk.TopkInt8sAdd(b, k, x)
}

func TopkUintsAdd(b []uint, k int, x uint) ([]uint, bool) {
	return ldtopk.TopkUintsAdd(b, k, x)
}
func TopkUint64sAdd(b []uint64, k int, x uint64) ([]uint64, bool) {
	return ldtopk.TopkUint64sAdd(b, k, x)
}
func TopkUint32sAdd(b []uint32, k int, x uint32) ([]uint32, bool) {
	return ldtopk.TopkUint32sAdd(b, k, x)
}
func TopkUint16sAdd(b []uint16, k int, x uint16) ([]uint16, bool) {
	return ldtopk.TopkUint16sAdd(b, k, x)
}
func TopkUint8sAdd(b []uint8, k int, x uint8) ([]uint8, bool) {
	return ldtopk.TopkUint8sAdd(b, k, x)
}

func TopkFloat32sAdd(b []float32, k int, x float32) ([]float32, bool) {
	return ldtopk.TopkFloat32sAdd(b, k, x)
}
func TopkFloat64sAdd(b []float64, k int, x float64) ([]float64, bool) {
	return ldtopk.TopkFloat64sAdd(b, k, x)
}
