/*
 * Copyright (C) distroy
 */

package lditer

type Iterable[Iter comparable] interface {
	comparable

	Prev() Iter
	Next() Iter
}

type Iterator[Data any] interface {
	Data() Data
	Prev() Iterator[Data]
	Next() Iterator[Data]
}

type Rangable interface {
	comparable

	HasNext() bool
	Next()
}

type Range[Data any] interface {
	Data() Data
	HasNext() bool
	Next()
}

type iterable[Data any, Iter comparable] interface {
	comparable

	Data() Data
	Prev() Iter
	Next() Iter
}

func MakeIter[Data any, Iter iterable[Data, Iter]](iter Iter) Iterator[Data] {
	return iterator[Data, Iter]{iter: iter}
}

func MakeRange[Data any, Iter iterable[Data, Iter]](begin, end Iter) Range[Data] {
	var zero Iter
	if begin == zero {
		return &_range[Data, Iter]{}
	}
	return &_range[Data, Iter]{
		begin: begin,
		end:   end,
	}
}

type iterator[Data any, Iter iterable[Data, Iter]] struct {
	iter Iter
}

func (i iterator[Data, Iter]) Data() Data           { return i.iter.Data() }
func (i iterator[Data, Iter]) Prev() Iterator[Data] { return MakeIter[Data](i.iter.Prev()) }
func (i iterator[Data, Iter]) Next() Iterator[Data] { return MakeIter[Data](i.iter.Next()) }

type _range[Data any, Iter iterable[Data, Iter]] struct {
	begin Iter
	end   Iter
}

func (r *_range[Data, Iter]) Data() Data    { return r.begin.Data() }
func (r *_range[Data, Iter]) HasNext() bool { return r.begin != r.end }
func (r *_range[Data, Iter]) Next()         { r.begin = r.begin.Next() }
