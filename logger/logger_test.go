/*
 * Copyright (C) distroy
 */

package logger

import "testing"

func TestLogger(t *testing.T) {
	l := NewLogger()
	l.Error("error message")
	l.Warn("warn message")
	l.Info("info message")
}
