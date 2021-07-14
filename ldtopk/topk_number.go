/*
 * Copyright (C) distroy
 */

package ldtopk

type topkSliceInt []int

func (s topkSliceInt) Len() int                   { return len(s) }
func (s topkSliceInt) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceInt) Less(i, j interface{}) bool { return i.(int) < j.(int) }
func (s topkSliceInt) Get(i int) interface{}      { return s[i] }
func (s topkSliceInt) Set(i int, x interface{})   { s[i] = x.(int) }
func (s *topkSliceInt) Push(x interface{})        { *s = append(*s, x.(int)) }

func TopkIntsAdd(b []int, k int, x int) ([]int, bool) {
	ok := TopkAdd((*topkSliceInt)(&b), k, x)
	return b, ok
}

type topkSliceInt64 []int64

func (s topkSliceInt64) Len() int                   { return len(s) }
func (s topkSliceInt64) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceInt64) Less(i, j interface{}) bool { return i.(int64) < j.(int64) }
func (s topkSliceInt64) Get(i int) interface{}      { return s[i] }
func (s topkSliceInt64) Set(i int, x interface{})   { s[i] = x.(int64) }
func (s *topkSliceInt64) Push(x interface{})        { *s = append(*s, x.(int64)) }

func TopkInt64sAdd(b []int64, k int, x int64) ([]int64, bool) {
	ok := TopkAdd((*topkSliceInt64)(&b), k, x)
	return b, ok
}

type topkSliceInt32 []int32

func (s topkSliceInt32) Len() int                   { return len(s) }
func (s topkSliceInt32) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceInt32) Less(i, j interface{}) bool { return i.(int32) < j.(int32) }
func (s topkSliceInt32) Get(i int) interface{}      { return s[i] }
func (s topkSliceInt32) Set(i int, x interface{})   { s[i] = x.(int32) }
func (s *topkSliceInt32) Push(x interface{})        { *s = append(*s, x.(int32)) }

func TopkInt32sAdd(b []int32, k int, x int32) ([]int32, bool) {
	ok := TopkAdd((*topkSliceInt32)(&b), k, x)
	return b, ok
}

type topkSliceInt16 []int16

func (s topkSliceInt16) Len() int                   { return len(s) }
func (s topkSliceInt16) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceInt16) Less(i, j interface{}) bool { return i.(int16) < j.(int16) }
func (s topkSliceInt16) Get(i int) interface{}      { return s[i] }
func (s topkSliceInt16) Set(i int, x interface{})   { s[i] = x.(int16) }
func (s *topkSliceInt16) Push(x interface{})        { *s = append(*s, x.(int16)) }

func TopkInt16sAdd(b []int16, k int, x int16) ([]int16, bool) {
	ok := TopkAdd((*topkSliceInt16)(&b), k, x)
	return b, ok
}

type topkSliceInt8 []int8

func (s topkSliceInt8) Len() int                   { return len(s) }
func (s topkSliceInt8) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceInt8) Less(i, j interface{}) bool { return i.(int8) < j.(int8) }
func (s topkSliceInt8) Get(i int) interface{}      { return s[i] }
func (s topkSliceInt8) Set(i int, x interface{})   { s[i] = x.(int8) }
func (s *topkSliceInt8) Push(x interface{})        { *s = append(*s, x.(int8)) }

func TopkInt8sAdd(b []int8, k int, x int8) ([]int8, bool) {
	ok := TopkAdd((*topkSliceInt8)(&b), k, x)
	return b, ok
}

type topkSliceUint []uint

func (s topkSliceUint) Len() int                   { return len(s) }
func (s topkSliceUint) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceUint) Less(i, j interface{}) bool { return i.(uint) < j.(uint) }
func (s topkSliceUint) Get(i int) interface{}      { return s[i] }
func (s topkSliceUint) Set(i int, x interface{})   { s[i] = x.(uint) }
func (s *topkSliceUint) Push(x interface{})        { *s = append(*s, x.(uint)) }

func TopkUintsAdd(b []uint, k int, x uint) ([]uint, bool) {
	ok := TopkAdd((*topkSliceUint)(&b), k, x)
	return b, ok
}

type topkSliceUint64 []uint64

func (s topkSliceUint64) Len() int                   { return len(s) }
func (s topkSliceUint64) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceUint64) Less(i, j interface{}) bool { return i.(uint64) < j.(uint64) }
func (s topkSliceUint64) Get(i int) interface{}      { return s[i] }
func (s topkSliceUint64) Set(i int, x interface{})   { s[i] = x.(uint64) }
func (s *topkSliceUint64) Push(x interface{})        { *s = append(*s, x.(uint64)) }

func TopkUint64sAdd(b []uint64, k int, x uint64) ([]uint64, bool) {
	ok := TopkAdd((*topkSliceUint64)(&b), k, x)
	return b, ok
}

type topkSliceUint32 []uint32

func (s topkSliceUint32) Len() int                   { return len(s) }
func (s topkSliceUint32) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceUint32) Less(i, j interface{}) bool { return i.(uint32) < j.(uint32) }
func (s topkSliceUint32) Get(i int) interface{}      { return s[i] }
func (s topkSliceUint32) Set(i int, x interface{})   { s[i] = x.(uint32) }
func (s *topkSliceUint32) Push(x interface{})        { *s = append(*s, x.(uint32)) }

func TopkUint32sAdd(b []uint32, k int, x uint32) ([]uint32, bool) {
	ok := TopkAdd((*topkSliceUint32)(&b), k, x)
	return b, ok
}

type topkSliceUint16 []uint16

func (s topkSliceUint16) Len() int                   { return len(s) }
func (s topkSliceUint16) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceUint16) Less(i, j interface{}) bool { return i.(uint16) < j.(uint16) }
func (s topkSliceUint16) Get(i int) interface{}      { return s[i] }
func (s topkSliceUint16) Set(i int, x interface{})   { s[i] = x.(uint16) }
func (s *topkSliceUint16) Push(x interface{})        { *s = append(*s, x.(uint16)) }

func TopkUint16sAdd(b []uint16, k int, x uint16) ([]uint16, bool) {
	ok := TopkAdd((*topkSliceUint16)(&b), k, x)
	return b, ok
}

type topkSliceUint8 []uint8

func (s topkSliceUint8) Len() int                   { return len(s) }
func (s topkSliceUint8) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceUint8) Less(i, j interface{}) bool { return i.(uint8) < j.(uint8) }
func (s topkSliceUint8) Get(i int) interface{}      { return s[i] }
func (s topkSliceUint8) Set(i int, x interface{})   { s[i] = x.(uint8) }
func (s *topkSliceUint8) Push(x interface{})        { *s = append(*s, x.(uint8)) }

func TopkUint8sAdd(b []uint8, k int, x uint8) ([]uint8, bool) {
	ok := TopkAdd((*topkSliceUint8)(&b), k, x)
	return b, ok
}

type topkSliceFloat32 []float32

func (s topkSliceFloat32) Len() int                   { return len(s) }
func (s topkSliceFloat32) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceFloat32) Less(i, j interface{}) bool { return i.(float32) < j.(float32) }
func (s topkSliceFloat32) Get(i int) interface{}      { return s[i] }
func (s topkSliceFloat32) Set(i int, x interface{})   { s[i] = x.(float32) }
func (s *topkSliceFloat32) Push(x interface{})        { *s = append(*s, x.(float32)) }

func TopkFloat32sAdd(b []float32, k int, x float32) ([]float32, bool) {
	ok := TopkAdd((*topkSliceFloat32)(&b), k, x)
	return b, ok
}

type topkSliceFloat64 []float64

func (s topkSliceFloat64) Len() int                   { return len(s) }
func (s topkSliceFloat64) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceFloat64) Less(i, j interface{}) bool { return i.(float64) < j.(float64) }
func (s topkSliceFloat64) Get(i int) interface{}      { return s[i] }
func (s topkSliceFloat64) Set(i int, x interface{})   { s[i] = x.(float64) }
func (s *topkSliceFloat64) Push(x interface{})        { *s = append(*s, x.(float64)) }

func TopkFloat64sAdd(b []float64, k int, x float64) ([]float64, bool) {
	ok := TopkAdd((*topkSliceFloat64)(&b), k, x)
	return b, ok
}
