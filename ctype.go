/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

func IsDigit(c byte) bool { return ldcore.IsDigit(c) }
func IsLower(c byte) bool { return ldcore.IsLower(c) }
func IsUpper(c byte) bool { return ldcore.IsUpper(c) }
func IsPrint(c byte) bool { return ldcore.IsPrint(c) }

func IsCtrl(c byte) bool { return ldcore.IsCtrl(c) }

func IsBlank(c byte) bool { return ldcore.IsBlank(c) }

func IsSpace(c byte) bool { return ldcore.IsSpace(c) }

func IsXDigit(c byte) bool { return ldcore.IsXDigit(c) }

func IsPunct(c byte) bool { return ldcore.IsPunct(c) }

// IsAlpha = IsUpper || IsLower
func IsAlpha(c byte) bool { return ldcore.IsAlpha(c) }

// IsAlNum = IsAlpha || IsDigit
func IsAlNum(c byte) bool { return ldcore.IsAlNum(c) }
