/*
 * Copyright (C) distroy
 */

package ldatomic

var _ StoreLoadSwapper[any] = (*Interface)(nil)

// Interface provides an atomic load and store of an any typed value.
// The zero value for a Value returns nil from Load.
type Interface = Any[interface{}]

func NewInterface(d interface{}) *Interface {
	p := &Interface{}
	p.Store(d)
	return p
}

type Error = Any[error]

func NewError(d error) *Error {
	p := &Error{}
	p.Store(d)
	return p
}

type Loader[T any] interface {
	Load() T
}

type Storer[T any] interface {
	Store(d T)
}

type StoreLoader[T any] interface {
	Loader[T]
	Storer[T]
}

type StoreLoadSwapper[T any] interface {
	StoreLoader[T]

	Swap(new T) (old T)
}

type StoreLoadCompareAndSwapper[T any] interface {
	StoreLoader[T]

	Swap(new T) (old T)
	CompareAndSwap(old, new T) (swapped bool)
}
