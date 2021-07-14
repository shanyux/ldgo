/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldbyte"
)

func IsDigit(c byte) bool { return ldbyte.IsDigit(c) }
func IsLower(c byte) bool { return ldbyte.IsLower(c) }
func IsUpper(c byte) bool { return ldbyte.IsUpper(c) }
func IsPrint(c byte) bool { return ldbyte.IsPrint(c) }

func IsCtrl(c byte) bool { return ldbyte.IsCtrl(c) }

func IsBlank(c byte) bool { return ldbyte.IsBlank(c) }

func IsSpace(c byte) bool { return ldbyte.IsSpace(c) }

func IsXDigit(c byte) bool { return ldbyte.IsXDigit(c) }

func IsPunct(c byte) bool { return ldbyte.IsPunct(c) }

// IsAlpha = IsUpper || IsLower
func IsAlpha(c byte) bool { return ldbyte.IsAlpha(c) }

// IsAlNum = IsAlpha || IsDigit
func IsAlNum(c byte) bool { return ldbyte.IsAlNum(c) }
