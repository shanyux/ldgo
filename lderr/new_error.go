/*
 * Copyright (C) distroy
 */

package lderr

import (
	"sync"
)

var errMap = &sync.Map{}

type Error interface {
	error
	Status() int
	Code() int
}

func NewError(status, code int, message string) Error {
	return New(status, code, message)
}

func New(status, code int, message string) Error {
	var err Error = commError{
		error:  strError{text: message},
		status: status,
		code:   code,
	}

	errMap.LoadOrStore(err.Code(), err)
	// errMap.Store(err.Code(), err)
	return err
}

func Wrap(err error, def ...Error) Error {
	if v, ok := err.(Error); ok {
		return v
	}

	d := ErrUnkown
	if len(def) != 0 {
		d = def[0]
	}

	return commError{
		error:  err,
		status: d.Status(),
		code:   d.Code(),
	}
}

func GetCode(code int) Error {
	v, _ := errMap.Load(code)
	if v == nil {
		return nil
	}

	err, ok := v.(Error)
	if !ok {
		return nil
	}
	return err
}

type commError struct {
	error

	status int
	code   int
}

func (e commError) Status() int { return e.status }
func (e commError) Code() int   { return e.code }

type strError struct {
	text string
}

func (e strError) Error() string { return e.text }
