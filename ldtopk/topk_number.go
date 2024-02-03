/*
 * Copyright (C) distroy
 */

package ldtopk

func TopkIntsAdd(b []int, k int, x int) ([]int, bool)             { return TopkAdd[int](b, k, x) }
func TopkInt64sAdd(b []int64, k int, x int64) ([]int64, bool)     { return TopkAdd[int64](b, k, x) }
func TopkInt32sAdd(b []int32, k int, x int32) ([]int32, bool)     { return TopkAdd[int32](b, k, x) }
func TopkInt16sAdd(b []int16, k int, x int16) ([]int16, bool)     { return TopkAdd[int16](b, k, x) }
func TopkInt8sAdd(b []int8, k int, x int8) ([]int8, bool)         { return TopkAdd[int8](b, k, x) }
func TopkUintsAdd(b []uint, k int, x uint) ([]uint, bool)         { return TopkAdd[uint](b, k, x) }
func TopkUint64sAdd(b []uint64, k int, x uint64) ([]uint64, bool) { return TopkAdd[uint64](b, k, x) }
func TopkUint32sAdd(b []uint32, k int, x uint32) ([]uint32, bool) { return TopkAdd[uint32](b, k, x) }
func TopkUint16sAdd(b []uint16, k int, x uint16) ([]uint16, bool) { return TopkAdd[uint16](b, k, x) }
func TopkUint8sAdd(b []uint8, k int, x uint8) ([]uint8, bool)     { return TopkAdd[uint8](b, k, x) }

func TopkUintptrsAdd(b []uintptr, k int, x uintptr) ([]uintptr, bool) {
	return TopkAdd[uintptr](b, k, x)
}

func TopkFloat32sAdd(b []float32, k int, x float32) ([]float32, bool) {
	return TopkAdd[float32](b, k, x)
}
func TopkFloat64sAdd(b []float64, k int, x float64) ([]float64, bool) {
	return TopkAdd[float64](b, k, x)
}
