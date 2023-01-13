/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func testNewAny(v interface{}) interface{} {
	if v == nil {
		return v
	}

	vv := reflect.ValueOf(v)
	res := reflect.New(vv.Type())
	res.Elem().Set(vv)
	return res.Interface()
}

func TestToBool(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		want    bool
		wantErr bool
	}{
		{name: `bool`, args: false, want: false, wantErr: false},

		{name: `int`, args: int(1), want: true, wantErr: false},
		{name: `int8`, args: int8(0), want: false, wantErr: false},
		{name: `int16`, args: int16(1), want: true, wantErr: false},
		{name: `int32`, args: int32(0), want: false, wantErr: false},
		{name: `int64`, args: int64(1), want: true, wantErr: false},

		{name: `uint`, args: uint(0), want: false, wantErr: false},
		{name: `uint8`, args: uint8(1), want: true, wantErr: false},
		{name: `uint16`, args: uint16(0), want: false, wantErr: false},
		{name: `uint32`, args: uint32(1), want: true, wantErr: false},
		{name: `uint64`, args: uint64(0), want: false, wantErr: false},

		{name: `float32`, args: float32(1), want: true, wantErr: false},
		{name: `float64`, args: float64(0), want: false, wantErr: false},

		{name: `big.Float`, args: *big.NewFloat(1), want: true, wantErr: false},
		{name: `*big.Float`, args: big.NewFloat(0), want: false, wantErr: false},

		{name: `decimal.Decimal`, args: mustNewDecimalFromStr("1"), want: true, wantErr: false},
		{name: `*decimal.Decimal`, args: testNewAny(mustNewDecimalFromStr("0")), want: false, wantErr: false},

		{name: `[]byte - null`, args: []byte("null"), want: false, wantErr: false},
		{name: `[]byte - true`, args: []byte("true"), want: true, wantErr: false},

		{name: `string - false`, args: "false", want: false, wantErr: false},
		{name: `string - ""`, args: "", want: false, wantErr: true},
	}

	convey.Convey(t.Name(), t, func() {
		for _, tt := range tests {
			convey.Convey(tt.name, func() {
				got, err := ToBool(tt.args)
				if tt.wantErr {
					convey.So(err, convey.ShouldNotBeNil)
				} else {
					convey.So(err, convey.ShouldBeNil)
				}
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		}
	})
}

func TestToInt64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    any
		want    int64
		wantErr bool
	}{
		{name: `bool`, args: false, want: 0, wantErr: false},

		{name: `int`, args: int(1), want: 1, wantErr: false},
		{name: `int8`, args: int8(-123), want: -123, wantErr: false},
		{name: `int16`, args: int16(-238), want: -238, wantErr: false},
		{name: `int32`, args: int32(2345), want: 2345, wantErr: false},
		{name: `int64`, args: int64(1234), want: 1234, wantErr: false},

		{name: `uint`, args: uint(1), want: 1, wantErr: false},
		{name: `uint8`, args: uint8(123), want: 123, wantErr: false},
		{name: `uint16`, args: uint16(238), want: 238, wantErr: false},
		{name: `uint32`, args: uint32(2345), want: 2345, wantErr: false},
		{name: `uint64`, args: uint64(1234), want: 1234, wantErr: false},

		{name: `float32`, args: float32(12.34), want: 12, wantErr: false},
		{name: `float64`, args: float64(-1234.5678), want: -1234, wantErr: false},

		{name: `big.Float`, args: *big.NewFloat(12345), want: 12345, wantErr: false},
		{name: `*big.Float`, args: big.NewFloat(123.45), want: 123, wantErr: false},

		{name: `decimal.Decimal`, args: mustNewDecimalFromStr("1"), want: 1, wantErr: false},
		{name: `*decimal.Decimal`, args: testNewAny(mustNewDecimalFromStr("0")), want: 0, wantErr: false},

		{name: `json.Number`, args: json.Number("12.34"), want: 12, wantErr: false},
		{name: `*json.Number`, args: testNewAny(json.Number("-1234.5678")), want: -1234, wantErr: false},

		{name: `[]byte - null`, args: []byte("null"), want: 0, wantErr: false},
		{name: `[]byte - true`, args: []byte("true"), want: 1, wantErr: false},

		{name: `string - ++1234.56`, args: "++1234.56", want: 1234, wantErr: false},
		{name: `string - ""`, args: "", want: 0, wantErr: true},
	}
	convey.Convey(t.Name(), t, func() {
		for _, tt := range tests {
			convey.Convey(tt.name, func() {
				got, err := ToInt64(tt.args)
				if tt.wantErr {
					convey.So(err, convey.ShouldNotBeNil)
				} else {
					convey.So(err, convey.ShouldBeNil)
				}
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		}
	})
}
