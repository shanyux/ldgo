/*
 * Copyright (C) distroy
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/distroy/ldgo/ldrand"
)

func printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintf(os.Stdout, "\n")
}

func main() {
	r := rand.New(ldrand.NewFastSource(time.Now().Unix()))
	printf("uint64:")
	for i := 0; i < 50; i++ {
		printf("%#0.16x", r.Uint64())
	}
	printf("")

	printf("uint32:")
	for i := 0; i < 50; i++ {
		printf("%#0.8x", r.Uint32())
	}
	printf("")

	printf("mod 16:")
	for i := 0; i < 50; i++ {
		printf("%d", r.Intn(16))
	}
	printf("")
}
