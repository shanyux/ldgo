/*
 * Copyright (C) distroy
 */

package ldgin

import "net/http"

var (
	ERR_INVALID_REQUEST_TYPE = NewError(http.StatusOK, -1, "input request type must be pointer to struct")
	ERR_PARSE_REQUEST_FAIL   = NewError(http.StatusOK, -1, "parse request fail")
)

type commError struct {
	status  int
	code    int
	message string
}

func NewError(status, code int, message string) Error {
	return commError{
		status:  status,
		code:    code,
		message: message,
	}
}

func (e commError) Status() int   { return e.status }
func (e commError) Code() int     { return e.code }
func (e commError) Error() string { return e.message }
