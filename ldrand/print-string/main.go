/*
 * Copyright (C) distroy
 */

package main

import (
	"encoding/hex"
	"fmt"

	"github.com/distroy/ldgo/ldflag"
	"github.com/distroy/ldgo/ldmath"
	"github.com/distroy/ldgo/ldrand"
)

type Flags struct {
	Count int `flag:"default:10"`
	Size  int `flag:"default:16"`
}

func main() {
	flags := &Flags{}
	ldflag.MustParse(flags)

	flags.Count = ldmath.MaxInt(flags.Count, 1)
	flags.Size = ldmath.MaxInt(flags.Size, 1)

	buf := make([]byte, flags.Size)

	for i := 0; i < flags.Count; i++ {
		ldrand.Read(buf)
		fmt.Println(hex.EncodeToString(buf))
	}
}