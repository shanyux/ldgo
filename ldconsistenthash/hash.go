/*
 * Copyright (C) distroy
 */

package ldconsistenthash

import "crypto/md5"

func md5Hash(data []byte) uint32 {
	h := md5.Sum(data)
	n := ((uint32(h[3]) << 24) | (uint32(h[2]) << 16) | (uint32(h[1]) << 8) | uint32(h[0]))
	return n
}
