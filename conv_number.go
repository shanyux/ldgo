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

const (
	_NUMBER_TYPE_NIL   = 0
	_NUMBER_TYPE_OCT   = 'o'
	_NUMBER_TYPE_DEC   = 'i'
	_NUMBER_TYPE_HEX   = 'x'
	_NUMBER_TYPE_FLOAT = 'f'
	_NUMBER_TYPE_TRUE  = 'T'
	_NUMBER_TYPE_FALSE = 'F'
)

type convNumberReader struct {
	buff     []byte
	pos      int
	begin    int
	end      int
	typNum   byte
	negative bool
}

func (that *convNumberReader) IsValid() bool {
	return that.pos < len(that.buff)
}

func (that *convNumberReader) Peek() byte {
	if that.pos < len(that.buff) {
		return that.buff[that.pos]
	}
	return 0
}

func (that *convNumberReader) Next() {
	if that.pos < len(that.buff) {
		that.pos++
	}
}

func (that *convNumberReader) StripSpace() {
	for {
		if !IsSpace(that.Peek()) {
			return
		}
		that.Next()
	}
}

func (that *convNumberReader) ReadString(s string) bool {
	for i := 1; i < len(s); i++ {
		that.Next()
		c := that.Peek()
		if ToLower(c) != s[i] {
			return false
		}
	}
	return true
}

func (that *convNumberReader) MarkBegin() { that.begin = that.pos }
func (that *convNumberReader) MarkEnd()   { that.end = that.pos }

func (that *convNumberReader) ReadOctal() bool {
	if c := that.Peek(); c < '0' || c > '7' {
		return false
	}
	for {
		that.Next()
		if c := that.Peek(); c < '0' || c > '7' {
			break
		}
	}
	return true
}

func (that *convNumberReader) ReadDecimal() bool {
	if !IsDigit(that.Peek()) {
		return false
	}
	for {
		that.Next()
		if !IsDigit(that.Peek()) {
			break
		}
	}
	return true
}

func (that *convNumberReader) ReadHexadecimal() bool {
	if !IsXDigit(that.Peek()) {
		return false
	}
	for {
		that.Next()
		if !IsXDigit(that.Peek()) {
			break
		}
	}
	return true
}

func (that *convNumberReader) ReadExponentiation() bool {
	switch that.Peek() {
	case '-':
		fallthrough
	case '+':
		that.Next()
	}
	return that.ReadDecimal()
}

func (that *convNumberReader) TestNumberType() {
	switch that.Peek() {
	case 't':
		fallthrough
	case 'T':
		if that.ReadString("true") {
			that.typNum = _NUMBER_TYPE_TRUE
		}
		return
	case 'f':
		fallthrough
	case 'F':
		if that.ReadString("false") {
			that.typNum = _NUMBER_TYPE_FALSE
		}
		return
	case 'n':
		fallthrough
	case 'N':
		if that.ReadString("null") {
			that.typNum = _NUMBER_TYPE_FALSE
		}
		return
	}

	// detect positive or negative
	for {
		switch that.Peek() {
		case '-':
			that.negative = !that.negative
			fallthrough
		case '+':
			that.Next()
			continue
		}
		break
	}

	switch that.Peek() {
	case '0':
		that.Next()
		switch c := that.Peek(); c {
		case '.':
			// 0.123
			that.MarkBegin()
			that.Next()
			that.ReadDecimal()
			if c := that.Peek(); c == 'e' || c == 'E' {
				that.Next()
				if !that.ReadExponentiation() {
					that.typNum = _NUMBER_TYPE_NIL
					return
				}
				that.ReadDecimal()
			}
			that.typNum = _NUMBER_TYPE_FLOAT
			that.MarkEnd()
			return

		case 'x':
			fallthrough
		case 'X':
			// 0x123
			that.Next()
			that.MarkBegin()
			if !that.ReadHexadecimal() {
				return
			}
			that.MarkEnd()
			that.typNum = _NUMBER_TYPE_HEX

		default:
			// 0123
			that.MarkBegin()
			that.ReadOctal()
			that.MarkEnd()
			that.typNum = _NUMBER_TYPE_OCT
		}

	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		that.MarkBegin()
		that.Next()
		that.typNum = _NUMBER_TYPE_DEC
		that.ReadDecimal()
		if that.Peek() == '.' {
			that.Next()
			that.typNum = _NUMBER_TYPE_FLOAT
			that.ReadDecimal()
		}
		if c := that.Peek(); c == 'e' || c == 'E' {
			that.Next()
			if !that.ReadExponentiation() {
				that.typNum = _NUMBER_TYPE_NIL
				return
			}
			that.typNum = _NUMBER_TYPE_FLOAT
			that.ReadDecimal()
		}
		that.MarkEnd()

	case '.':
		that.MarkBegin()
		that.Next()
		if !IsDigit(that.Peek()) {
			return
		}
		that.ReadDecimal()
		if c := that.Peek(); c == 'e' || c == 'E' {
			that.Next()
			if !that.ReadExponentiation() {
				that.typNum = _NUMBER_TYPE_NIL
				return
			}
		}
		that.typNum = _NUMBER_TYPE_FLOAT
		that.MarkEnd()
	}
}

func (that *convNumberReader) GetString() string {
	if that.typNum == _NUMBER_TYPE_NIL {
		return ""
	}
	return BytesToStrUnsafe(that.buff[that.begin:that.end])
}

func testStringNumberType(b []byte) (byte, bool, string) {
	if len(b) == 0 {
		return _NUMBER_TYPE_NIL, false, ""
	}

	that := &convNumberReader{
		buff: b,
	}

	that.StripSpace()
	that.TestNumberType()
	s := that.GetString()
	// log.Printf("str:%s, type:%c, sub:%s", b, that.typNum, s)
	return that.typNum, that.negative, s
}

func convStrToOct(s string) (int64, error) {
	var n uint64
	for _, c := range s {
		n *= 8
		n += uint64(c) - '0'
	}
	return int64(n), nil
}

func convStrToDec(s string) (int64, error) {
	var n uint64
	for _, c := range s {
		n *= 10
		n += uint64(c) - '0'
	}
	return int64(n), nil
}

func convStrToHex(s string) (int64, error) {
	var n uint64
	for _, c := range s {
		n *= 16
		if c >= 'A' {
			n += uint64((c&0xdf)-'A') + 10
		} else {
			n += uint64(c) - '0'
		}
	}
	return int64(n), nil
}

func convStrToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func convBool(b []byte) (bool, error) {
	typNum, _, str := testStringNumberType(b)
	switch typNum {
	case _NUMBER_TYPE_TRUE:
		return true, nil
	case _NUMBER_TYPE_FALSE:
		return false, nil
	case _NUMBER_TYPE_OCT:
		d, err := convStrToOct(str)
		return d != 0, err
	case _NUMBER_TYPE_DEC:
		d, err := convStrToDec(str)
		return d != 0, err
	case _NUMBER_TYPE_HEX:
		d, err := convStrToHex(str)
		return d != 0, err
	case _NUMBER_TYPE_FLOAT:
		f, err := convStrToFloat(str)
		return int64(f) != 0, err
	}
	return false, _ERR_INVALID_SYNTAX
}

func convInt(b []byte) (int64, error) {
	typNum, negative, str := testStringNumberType(b)
	switch typNum {
	case _NUMBER_TYPE_TRUE:
		return 1, nil
	case _NUMBER_TYPE_FALSE:
		return 0, nil
	case _NUMBER_TYPE_OCT:
		d, err := convStrToOct(str)
		if negative {
			d = -d
		}
		return d, err
	case _NUMBER_TYPE_DEC:
		d, err := convStrToDec(str)
		if negative {
			d = -d
		}
		return d, err
	case _NUMBER_TYPE_HEX:
		d, err := convStrToHex(str)
		if negative {
			d = -d
		}
		return d, err
	case _NUMBER_TYPE_FLOAT:
		f, err := convStrToFloat(str)
		if negative {
			f = -f
		}
		return int64(f), err
	}
	return 0, _ERR_INVALID_SYNTAX
}

func convUint(b []byte) (uint64, error) {
	typNum, negative, str := testStringNumberType(b)
	switch typNum {
	case _NUMBER_TYPE_TRUE:
		return 1, nil
	case _NUMBER_TYPE_FALSE:
		return 0, nil
	case _NUMBER_TYPE_OCT:
		d, err := convStrToOct(str)
		if negative {
			d = -d
		}
		return uint64(d), err
	case _NUMBER_TYPE_DEC:
		d, err := convStrToDec(str)
		if negative {
			d = -d
		}
		return uint64(d), err
	case _NUMBER_TYPE_HEX:
		d, err := convStrToHex(str)
		if negative {
			d = -d
		}
		return uint64(d), err
	case _NUMBER_TYPE_FLOAT:
		f, err := convStrToFloat(str)
		if negative {
			f = -f
		}
		return uint64(f), err
	}
	return 0, _ERR_INVALID_SYNTAX
}

func convFloat(b []byte) (float64, error) {
	typNum, negative, str := testStringNumberType(b)
	switch typNum {
	case _NUMBER_TYPE_TRUE:
		return 1, nil
	case _NUMBER_TYPE_FALSE:
		return 0, nil
	case _NUMBER_TYPE_OCT:
		d, err := convStrToOct(str)
		if negative {
			return float64(-d), err
		}
		return float64(uint64(d)), err
	case _NUMBER_TYPE_DEC:
		d, err := convStrToDec(str)
		if negative {
			return float64(-d), err
		}
		return float64(uint64(d)), err
	case _NUMBER_TYPE_HEX:
		d, err := convStrToHex(str)
		if negative {
			return float64(-d), err
		}
		return float64(uint64(d)), err
	case _NUMBER_TYPE_FLOAT:
		f, err := convStrToFloat(str)
		if negative {
			f = -f
		}
		return f, err
	}
	return 0, _ERR_INVALID_SYNTAX
}
