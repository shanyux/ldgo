/*
 * Copyright (C) distroy
 */

package lderr

import (
	"context"
	"net/http"
)

func GetCode(err error, def ...int) int {
	if err == nil {
		return 0
	}

	if e := getMatchError(err); e != nil {
		return e.Code()
	}

	switch v := err.(type) {
	case interface{ Code() int }:
		return v.Code()
	}

	if len(def) > 0 {
		return def[0]
	}

	return errCodeUnkown
}

func GetStatus(err error, def ...int) int {
	if err == nil {
		return http.StatusOK
	}

	if e := getMatchError(err); e != nil {
		return e.Status()
	}

	switch v := err.(type) {
	case interface{ Status() int }:
		return v.Status()
	}

	if len(def) > 0 {
		return def[0]
	}
	return errStatusUnkonw
}

func GetMessage(err error) string {
	if err == nil {
		return errMessageSucess
	}
	return err.Error()
}

func GetDetails(err error) []string {
	if err == nil {
		return nil
	}
	switch v := err.(type) {
	case interface{ Details() []string }:
		return v.Details()

	case interface{ Detail() string }:
		return []string{v.Detail()}
	}
	return nil
}

func getMatchError(err error) Error {
	switch err {
	case nil:
		return ErrSuccess

	case context.Canceled:
		return ErrCtxCanceled

	case context.DeadlineExceeded:
		return ErrCtxDeadlineExceeded
	}
	return nil
}
