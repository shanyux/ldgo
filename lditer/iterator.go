/*
 * Copyright (C) distroy
 */

package lditer

type Iterable[Iter comparable] interface {
	comparable

	Prev() Iter
	Next() Iter
}

type ConstIterator[Data any] interface {
	Get() Data
	Prev() ConstIterator[Data]
	Next() ConstIterator[Data]
}

type Iterator[Data any] interface {
	Set(d Data)
	Get() Data
	Prev() Iterator[Data]
	Next() Iterator[Data]
}

type Rangable interface {
	comparable

	HasNext() bool
	Next()
}

type ConstRange[Data any] interface {
	Get() Data
	HasNext() bool
	Next()
}

type Range[Data any] interface {
	ConstRange[Data]

	Set(d Data)
}

type constIterable[Data any, Iter comparable] interface {
	Iterable[Iter]

	Get() Data
}

type iterable[Data any, Iter comparable] interface {
	constIterable[Data, Iter]

	Set(d Data)
}

func Iter[Data any, Iter iterable[Data, Iter]](iter Iter) Iterator[Data] {
	return iterator[Data, Iter]{iter: iter}
}

func makeIter[Data any, Iter iterable[Data, Iter]](iter Iter) Iterator[Data] {
	return iterator[Data, Iter]{iter: iter}
}

func ConstIter[Data any, Iter constIterable[Data, Iter]](iter Iter) ConstIterator[Data] {
	return constIterator[Data, Iter]{iter: iter}
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

func MakeConstRange[Data any, Iter constIterable[Data, Iter]](begin, end Iter) ConstRange[Data] {
	var zero Iter
	if begin == zero {
		return &_constRange[Data, Iter]{}
	}
	return &_constRange[Data, Iter]{
		begin: begin,
		end:   end,
	}
}

// iterator begin

type iterator[Data any, Iter iterable[Data, Iter]] struct {
	iter Iter
}

func (i iterator[Data, Iter]) Set(d Data)           { i.iter.Set(d) }
func (i iterator[Data, Iter]) Get() Data            { return i.iter.Get() }
func (i iterator[Data, Iter]) Prev() Iterator[Data] { return makeIter[Data](i.iter.Prev()) }
func (i iterator[Data, Iter]) Next() Iterator[Data] { return makeIter[Data](i.iter.Next()) }

// iterator end

// const iterator begin

type constIterator[Data any, Iter constIterable[Data, Iter]] struct {
	iter Iter
}

func (i constIterator[Data, Iter]) Get() Data { return i.iter.Get() }
func (i constIterator[Data, Iter]) Prev() ConstIterator[Data] {
	return ConstIter[Data](i.iter.Prev())
}
func (i constIterator[Data, Iter]) Next() ConstIterator[Data] {
	return ConstIter[Data](i.iter.Next())
}

// const iterator end

// range begin

type _range[Data any, Iter iterable[Data, Iter]] struct {
	begin Iter
	end   Iter
}

func (r *_range[Data, Iter]) Set(d Data)    { r.begin.Set(d) }
func (r *_range[Data, Iter]) Get() Data     { return r.begin.Get() }
func (r *_range[Data, Iter]) HasNext() bool { return r.begin != r.end }
func (r *_range[Data, Iter]) Next()         { r.begin = r.begin.Next() }

// range end

// const range begin

type _constRange[Data any, Iter constIterable[Data, Iter]] struct {
	begin Iter
	end   Iter
}

func (r *_constRange[Data, Iter]) Get() Data     { return r.begin.Get() }
func (r *_constRange[Data, Iter]) HasNext() bool { return r.begin != r.end }
func (r *_constRange[Data, Iter]) Next()         { r.begin = r.begin.Next() }

// const range end
