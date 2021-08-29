/*
 * Copyright (C) distroy
 */

package lderr

import "net/http"

// error definitions: [-1, -999]
var (
	ErrUnkown       = NewError(http.StatusOK, -1, "unknown error")
	ErrServicePanic = NewError(http.StatusServiceUnavailable, -2, "service panic")

	ErrCtxCanceled         = NewError(http.StatusOK, -11, "context canceled")
	ErrCtxDeadlineExceeded = NewError(http.StatusOK, -12, "context deadline exceeded")

	ErrNonAuthoritativeInfo = NewError(http.StatusNonAuthoritativeInfo, -101, "http non authoritative info")
	ErrUnauthorized         = NewError(http.StatusUnauthorized, -102, "http unauthorized")
	ErrInternalServerError  = NewError(http.StatusInternalServerError, -103, "http internal server error")
	ErrTooManyRequests      = NewError(http.StatusTooManyRequests, -104, "too many requests")

	ErrInvalidRequestType = NewError(http.StatusOK, -201, "input request type must be pointer to struct")
	ErrParseRequest       = NewError(http.StatusOK, -202, "parse request error")
	ErrRequestValidation  = NewError(http.StatusOK, -203, "request validation error")
	ErrDuplicateRequest   = NewError(http.StatusOK, -204, "duplicate request")

	ErrHttpCall          = NewError(http.StatusOK, -301, "http call request error")
	ErrHttpReadBody      = NewError(http.StatusOK, -302, "http read body error")
	ErrHttpNewRequest    = NewError(http.StatusOK, -303, "http new request error")
	ErrHttpInvalidStatus = NewError(http.StatusOK, -304, "http invalid status")

	ErrRpcCall      = NewError(http.StatusOK, -311, "rpc call error")
	ErrRpcTimeout   = NewError(http.StatusOK, -312, "rpc timeout")
	ErrRpcBokenPipe = NewError(http.StatusOK, -313, "rpc boken pip")

	ErrDataNotFound  = NewError(http.StatusOK, -801, "data not found")
	ErrDuplicateData = NewError(http.StatusOK, -802, "duplicate data")

	ErrDbQuery    = NewError(http.StatusOK, -811, "database query error")
	ErrDbExcute   = NewError(http.StatusOK, -812, "database execute error")
	ErrDbTxBegin  = NewError(http.StatusOK, -813, "database tx begin error")
	ErrDbTxCommit = NewError(http.StatusOK, -814, "database tx commit error")

	ErrCacheRead  = NewError(http.StatusOK, -821, "cache read error")
	ErrCacheWrite = NewError(http.StatusOK, -821, "cache write error")
)
