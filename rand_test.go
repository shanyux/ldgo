/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"log"
	"testing"
)

func TestRandString(t *testing.T) {
	log.Printf("rand string: %s", RandString(16))
	log.Printf("rand string: %s", RandString(16))
}
