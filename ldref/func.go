/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"runtime"
	"strings"
)

type FuncName struct {
	Full     string
	Short    string
	Path     string
	Package  string
	Method   string
	Receiver string
}

func GetFuncName(f interface{}) FuncName {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return ParseFuncName(fullName)
}

func ParseFuncName(fullName string) FuncName {
	// shortName := path.Base(fullName)
	shortName := fullName
	pathName := ""
	if idx := strings.LastIndex(fullName, "/"); idx >= 0 {
		pathName = fullName[:idx]
		shortName = fullName[idx+1:]
	}

	firstDotIndex := strings.Index(shortName, ".")
	lastDotIndex := strings.LastIndex(shortName, ".")

	packageName := shortName
	if firstDotIndex >= 0 {
		packageName = shortName[:firstDotIndex]
	}

	methodName := shortName
	if lastDotIndex >= 0 {
		methodName = shortName[lastDotIndex+1:]
	}

	receiverName := ""
	if firstDotIndex >= 0 && lastDotIndex >= 0 && firstDotIndex < lastDotIndex {
		if firstDotIndex+1 < lastDotIndex-1 && shortName[firstDotIndex+1] == '(' && shortName[lastDotIndex-1] == ')' {
			receiverName = shortName[firstDotIndex+2 : lastDotIndex-1]
		} else {
			receiverName = shortName[firstDotIndex+1 : lastDotIndex]
		}
	}

	return FuncName{
		Full:     fullName,
		Short:    shortName,
		Path:     pathName,
		Package:  packageName,
		Method:   methodName,
		Receiver: receiverName,
	}
}
