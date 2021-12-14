/*
 * Copyright (C) distroy
 */

package ldsort

type Ints []int

func (s Ints) Len() int      { return len(s) }
func (s Ints) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Ints) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortInts(a []int)          { internalSort(Ints(a)) }
func UniqInts(a []int) []int    { return a[:uniq(Ints(a))] }
func IsSortedInts(a []int) bool { return internalIsSorted(Ints(a)) }
func SearchInts(a []int, x int) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexInts(a []int, x int) int {
	if idx := SearchInts(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Int64s []int64

func (s Int64s) Len() int      { return len(s) }
func (s Int64s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Int64s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortInt64s(a []int64)          { internalSort(Int64s(a)) }
func UniqInt64s(a []int64) []int64  { return a[:uniq(Int64s(a))] }
func IsSortedInt64s(a []int64) bool { return internalIsSorted(Int64s(a)) }
func SearchInt64s(a []int64, x int64) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexInt64s(a []int64, x int64) int {
	if idx := SearchInt64s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Int32s []int32

func (s Int32s) Len() int      { return len(s) }
func (s Int32s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Int32s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortInt32s(a []int32)          { internalSort(Int32s(a)) }
func UniqInt32s(a []int32) []int32  { return a[:uniq(Int32s(a))] }
func IsSortedInt32s(a []int32) bool { return internalIsSorted(Int32s(a)) }
func SearchInt32s(a []int32, x int32) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexInt32s(a []int32, x int32) int {
	if idx := SearchInt32s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Int16s []int16

func (s Int16s) Len() int      { return len(s) }
func (s Int16s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Int16s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortInt16s(a []int16)          { internalSort(Int16s(a)) }
func UniqInt16s(a []int16) []int16  { return a[:uniq(Int16s(a))] }
func IsSortedInt16s(a []int16) bool { return internalIsSorted(Int16s(a)) }
func SearchInt16s(a []int16, x int16) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexInt16s(a []int16, x int16) int {
	if idx := SearchInt16s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Int8s []int8

func (s Int8s) Len() int      { return len(s) }
func (s Int8s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Int8s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortInt8s(a []int8)          { internalSort(Int8s(a)) }
func UniqInt8s(a []int8) []int8   { return a[:uniq(Int8s(a))] }
func IsSortedInt8s(a []int8) bool { return internalIsSorted(Int8s(a)) }
func SearchInt8s(a []int8, x int8) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexInt8s(a []int8, x int8) int {
	if idx := SearchInt8s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Uints []uint

func (s Uints) Len() int      { return len(s) }
func (s Uints) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Uints) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortUints(a []uint)          { internalSort(Uints(a)) }
func UniqUints(a []uint) []uint   { return a[:uniq(Uints(a))] }
func IsSortedUints(a []uint) bool { return internalIsSorted(Uints(a)) }
func SearchUints(a []uint, x uint) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexUints(a []uint, x uint) int {
	if idx := SearchUints(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Uint64s []uint64

func (s Uint64s) Len() int      { return len(s) }
func (s Uint64s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Uint64s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortUint64s(a []uint64)          { internalSort(Uint64s(a)) }
func UniqUint64s(a []uint64) []uint64 { return a[:uniq(Uint64s(a))] }
func IsSortedUint64s(a []uint64) bool { return internalIsSorted(Uint64s(a)) }
func SearchUint64s(a []uint64, x uint64) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexUint64s(a []uint64, x uint64) int {
	if idx := SearchUint64s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Uint32s []uint32

func (s Uint32s) Len() int      { return len(s) }
func (s Uint32s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Uint32s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortUint32s(a []uint32)          { internalSort(Uint32s(a)) }
func UniqUint32s(a []uint32) []uint32 { return a[:uniq(Uint32s(a))] }
func IsSortedUint32s(a []uint32) bool { return internalIsSorted(Uint32s(a)) }
func SearchUint32s(a []uint32, x uint32) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexUint32s(a []uint32, x uint32) int {
	if idx := SearchUint32s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Uint16s []uint16

func (s Uint16s) Len() int      { return len(s) }
func (s Uint16s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Uint16s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortUint16s(a []uint16)          { internalSort(Uint16s(a)) }
func UniqUint16s(a []uint16) []uint16 { return a[:uniq(Uint16s(a))] }
func IsSortedUint16s(a []uint16) bool { return internalIsSorted(Uint16s(a)) }
func SearchUint16s(a []uint16, x uint16) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexUint16s(a []uint16, x uint16) int {
	if idx := SearchUint16s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Uint8s []uint8

func (s Uint8s) Len() int      { return len(s) }
func (s Uint8s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Uint8s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortUint8s(a []uint8)          { internalSort(Uint8s(a)) }
func UniqUint8s(a []uint8) []uint8  { return a[:uniq(Uint8s(a))] }
func IsSortedUint8s(a []uint8) bool { return internalIsSorted(Uint8s(a)) }
func SearchUint8s(a []uint8, x uint8) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexUint8s(a []uint8, x uint8) int {
	if idx := SearchUint8s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Float64s []float64

func (s Float64s) Len() int      { return len(s) }
func (s Float64s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Float64s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortFloat64s(a []float64)           { internalSort(Float64s(a)) }
func UniqFloat64s(a []float64) []float64 { return a[:uniq(Float64s(a))] }
func IsSortedFloat64s(a []float64) bool  { return internalIsSorted(Float64s(a)) }
func SearchFloat64s(a []float64, x float64) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexFloat64s(a []float64, x float64) int {
	if idx := SearchFloat64s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}

type Float32s []float32

func (s Float32s) Len() int      { return len(s) }
func (s Float32s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Float32s) Compare(i, j int) int {
	if s[i] == s[j] {
		return 0
	} else if s[i] < s[j] {
		return -1
	}
	return +1
}

func SortFloat32s(a []float32)           { internalSort(Float32s(a)) }
func UniqFloat32s(a []float32) []float32 { return a[:uniq(Float32s(a))] }
func IsSortedFloat32s(a []float32) bool  { return internalIsSorted(Float32s(a)) }
func SearchFloat32s(a []float32, x float32) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexFloat32s(a []float32, x float32) int {
	if idx := SearchFloat32s(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}
