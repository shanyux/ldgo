/*
 * Copyright (C) distroy
 */

package ldredis

import "errors"

var (
	ErrMutexLocked    = errors.New("redis mutex had been locked")
	ErrMutexLocking   = errors.New("redis mutex is locking")
	ErrMutexNotExists = errors.New("redis mutex is not exists")
	ErrMutexNotMatch  = errors.New("redis mutex is not match")
)
