/*
 * Copyright (C) distroy
 */

package lderr

import (
	"context"
	"net/http"
)

func GetCode(err error, def ...int) int {
	switch err {
	case nil:
		return 0
	case context.Canceled:
		return ErrCtxCanceled.Code()
	case context.DeadlineExceeded:
		return ErrCtxDeadlineExceeded.Code()
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
	switch err {
	case nil:
		return http.StatusOK
	case context.Canceled:
		return ErrCtxCanceled.Status()
	case context.DeadlineExceeded:
		return ErrCtxDeadlineExceeded.Status()
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

func GetMessage(err error, def ...string) string {
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
