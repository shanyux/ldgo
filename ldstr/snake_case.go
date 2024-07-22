/*
 * Copyright (C) distroy
 */

package ldstr

import (
	"bytes"
	"unicode"

	"github.com/distroy/ldgo/v2/ldconv"
	"github.com/distroy/ldgo/v2/ldsort"
)

func ToSnakeCase(s string, sep ...rune) string {
	if s == "" {
		return s
	}
	if len(sep) > 0 && sep[0] != 0 {
		return toSnakeCase(s, sep[0])
	}
	return toSnakeCase(s, '_')
}

func toSnakeCase(s string, sep rune) string {
	reader := bytes.NewBuffer(ldconv.StrToBytesUnsafe(s))

	writer := bytes.NewBuffer(nil)
	writer.Grow(len(s) * 2)

	lastSep := true
	writeSep := func() {
		if lastSep {
			return
		}
		writer.WriteRune(sep)
		lastSep = true
	}
	writeRune := func(c rune) {
		if c == sep {
			writeSep()
			return
		}
		writer.WriteRune(c)
		lastSep = false
	}

	loopReadRunes(reader, func(i, size int, curr rune) bool {
		if unicode.IsSpace(curr) {
			writeRune(sep)
			return true

		} else if !unicode.IsUpper(curr) {
			writeRune(curr)
			return true

		} else if i > 0 {
			writeRune(sep)
		}

		last := unicode.ToLower(curr)
		loopReadRunes(reader, func(i, _ int, curr rune) bool {
			if unicode.IsSpace(curr) {
				writeRune(last)
				last = sep
				return false

			} else if unicode.IsUpper(curr) {
				writeRune(last)
				last = unicode.ToLower(curr)
				return true

			} else if i > 0 {
				writeRune(sep)
			}

			writeRune(last)
			last = unicode.ToLower(curr)
			return false
		})
		writeRune(last)
		return true
	})

	raw := writer.Bytes()
	if l := len(raw); rune(raw[l-1]) == sep {
		raw = raw[:l-1]
	}
	return ldconv.BytesToStrUnsafe(raw)
}

func ToCamelCase(s string, seps ...[]rune) string {
	if s == "" {
		return s
	}
	if len(seps) > 0 && len(seps[0]) > 0 {
		ldsort.SortInt32s(seps[0])
		return toCamelCase(s, seps[0])
	}
	return toCamelCase(s, []rune{'-', '_'})
}

func toCamelCase(s string, seps []rune) string {
	reader := bytes.NewBuffer(ldconv.StrToBytesUnsafe(s))
	writer := bytes.NewBuffer(nil)
	writer.Grow(len(s))

	lastSep := true
	loopReadRunes(reader, func(i, size int, curr rune) bool {
		if unicode.IsSpace(curr) || ldsort.IndexInt32s(seps, curr) >= 0 {
			lastSep = true
			return true

		} else if unicode.IsUpper(curr) {
			lastSep = false
			writer.WriteRune(curr)
			return true

		} else if lastSep {
			lastSep = false
			writer.WriteRune(unicode.ToUpper(curr))
			return true
		}

		lastSep = false
		writer.WriteRune(curr)
		return true
	})

	raw := writer.Bytes()
	return ldconv.BytesToStrUnsafe(raw)
}

func loopReadRunes(r *bytes.Buffer, fBody func(i int, size int, c rune) bool) {
	for i := 0; ; i++ {
		c, size, err := r.ReadRune()
		if err != nil {
			break
		}
		if ok := fBody(i, size, c); !ok {
			break
		}
	}
}
