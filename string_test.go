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
	"testing"
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
var _ = testing.Main

func TestStrMapReplace(t *testing.T) {
	s0 := "a}{{a}{abc}{b}{}.{{a}}"
	m := map[string]string{
		"a": "A",
		"b": "B",
	}
	s1 := StrMapReplace(s0, m, "{", "}")
	t.Logf("str: %s", s0)
	t.Logf("res: %s", s1)
}
