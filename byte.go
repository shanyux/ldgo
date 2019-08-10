/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = context.Canceled
var _ = errors.New
var _ = fmt.Sprintf
var _ = log.New
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

func ToLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func ToUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}
