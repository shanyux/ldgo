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

type ErrorWithDetails interface {
	Error

	Details() []string
}

func New(status, code int, message string) Error {
	var err Error = &commError{
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

	return &commError{
		error:  err,
		status: d.Status(),
		code:   d.Code(),
	}
}

func GetByCode(code int) Error {
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

func (e *commError) Status() int { return e.status }
func (e *commError) Code() int   { return e.code }

type strError struct {
	text string
}

func (e strError) Error() string { return e.text }

func WithDetail(err Error, details ...string) ErrorWithDetails {
	return WithDetails(err, details)
}

func WithDetails(err Error, details []string) ErrorWithDetails {
	var d []string
	switch v := err.(type) {
	case *detailsError:
		if len(details) == 0 {
			return v
		}

		err = v.err
		t := v.Details()
		d = make([]string, 0, len(details)+len(t))
		d = append(d, t...)
		d = append(d, details...)

	case ErrorWithDetails:
		if len(details) == 0 {
			return v
		}

		t := v.Details()
		d = make([]string, 0, len(details)+len(t))
		d = append(d, t...)
		d = append(d, details...)

	default:
		d = details
	}

	return &detailsError{
		err:     err,
		details: d,
	}
}

type detailsError struct {
	err Error

	details []string
}

func (e *detailsError) Error() string     { return e.err.Error() }
func (e *detailsError) Status() int       { return e.err.Status() }
func (e *detailsError) Code() int         { return e.err.Code() }
func (e *detailsError) Details() []string { return e.details }
