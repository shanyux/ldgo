/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"math/big"
	"strconv"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05-0700"
)

func ToByte(v interface{}) (byte, error) { return ToUint8(v) }

func ToBool(v interface{}) (bool, error) {
	switch vv := v.(type) {
	case bool:
		return vv, nil

	case int:
		return vv != 0, nil
	case int8:
		return vv != 0, nil
	case int16:
		return vv != 0, nil
	case int32:
		return vv != 0, nil
	case int64:
		return vv != 0, nil

	case uint:
		return vv != 0, nil
	case uint8:
		return vv != 0, nil
	case uint16:
		return vv != 0, nil
	case uint32:
		return vv != 0, nil
	case uint64:
		return vv != 0, nil

	case float32:
		return vv != 0, nil
	case float64:
		return vv != 0, nil

	case big.Float:
		return vv.Cmp(newBigFloatZero()) != 0, nil
	case *big.Float:
		return vv.Cmp(newBigFloatZero()) != 0, nil

	case decimalNumber:
		return !vv.IsZero(), nil
	case *decimalNumber:
		return !vv.IsZero(), nil

	case []byte:
		return convBool(vv)
	case string:
		return convBool(StrToBytesUnsafe(vv))
	}
	return false, _ERR_UNKOWN_TYPE
}

func ToInt(v interface{}) (int, error) {
	n, err := ToInt64(v)
	return int(n), err
}
func ToInt8(v interface{}) (int8, error) {
	n, err := ToInt64(v)
	return int8(n), err
}
func ToInt16(v interface{}) (int16, error) {
	n, err := ToInt64(v)
	return int16(n), err
}
func ToInt32(v interface{}) (int32, error) {
	n, err := ToInt64(v)
	return int32(n), err
}

func ToInt64(v interface{}) (int64, error) {
	switch vv := v.(type) {
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil

	case int:
		return int64(vv), nil
	case int8:
		return int64(vv), nil
	case int16:
		return int64(vv), nil
	case int32:
		return int64(vv), nil
	case int64:
		return int64(vv), nil

	case uint:
		return int64(vv), nil
	case uint8:
		return int64(vv), nil
	case uint16:
		return int64(vv), nil
	case uint32:
		return int64(vv), nil
	case uint64:
		return int64(vv), nil

	case float32:
		return int64(vv), nil
	case float64:
		return int64(vv), nil

	case big.Float:
		r, _ := vv.Int64()
		return r, nil
	case *big.Float:
		r, _ := vv.Int64()
		return r, nil

	case decimalNumber:
		return vv.BigInt().Int64(), nil
	case *decimalNumber:
		return vv.BigInt().Int64(), nil

	case []byte:
		return convInt(vv)
	case string:
		return convInt(StrToBytesUnsafe(vv))
	}
	return 0, _ERR_UNKOWN_TYPE
}

func ToUint(v interface{}) (uint, error) {
	n, err := ToUint64(v)
	return uint(n), err
}
func ToUint8(v interface{}) (uint8, error) {
	n, err := ToUint64(v)
	return uint8(n), err
}
func ToUint16(v interface{}) (uint16, error) {
	n, err := ToUint64(v)
	return uint16(n), err
}
func ToUint32(v interface{}) (uint32, error) {
	n, err := ToUint64(v)
	return uint32(n), err
}

func ToUint64(v interface{}) (uint64, error) {
	switch vv := v.(type) {
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil

	case int:
		return uint64(vv), nil
	case int8:
		return uint64(vv), nil
	case int16:
		return uint64(vv), nil
	case int32:
		return uint64(vv), nil
	case int64:
		return uint64(vv), nil

	case uint:
		return uint64(vv), nil
	case uint8:
		return uint64(vv), nil
	case uint16:
		return uint64(vv), nil
	case uint32:
		return uint64(vv), nil
	case uint64:
		return uint64(vv), nil

	case float32:
		return uint64(vv), nil
	case float64:
		return uint64(vv), nil

	case big.Float:
		r, _ := vv.Int64()
		return uint64(r), nil
	case *big.Float:
		r, _ := vv.Uint64()
		return uint64(r), nil

	case decimalNumber:
		return vv.BigInt().Uint64(), nil
	case *decimalNumber:
		return vv.BigInt().Uint64(), nil

	case []byte:
		return convUint(vv)
	case string:
		return convUint(StrToBytesUnsafe(vv))
	}
	return 0, _ERR_UNKOWN_TYPE
}

func ToFloat32(v interface{}) (float32, error) {
	switch vv := v.(type) {
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil

	case int:
		return float32(vv), nil
	case int8:
		return float32(vv), nil
	case int16:
		return float32(vv), nil
	case int32:
		return float32(vv), nil
	case int64:
		return float32(vv), nil

	case uint:
		return float32(vv), nil
	case uint8:
		return float32(vv), nil
	case uint16:
		return float32(vv), nil
	case uint32:
		return float32(vv), nil
	case uint64:
		return float32(vv), nil

	case float32:
		return float32(vv), nil
	case float64:
		return float32(vv), nil

	case big.Float:
		r, _ := vv.Float32()
		return r, nil
	case *big.Float:
		r, _ := vv.Float32()
		return r, nil

	case decimalNumber:
		r, _ := vv.Rat().Float32()
		return r, nil
	case *decimalNumber:
		r, _ := vv.Rat().Float32()
		return r, nil

	case []byte:
		f, err := convFloat(vv)
		r, _ := f.Rat().Float32()
		return r, err
	case string:
		f, err := convFloat(StrToBytesUnsafe(vv))
		r, _ := f.Rat().Float32()
		return r, err
	}
	return 0, _ERR_UNKOWN_TYPE
}

func ToFloat64(v interface{}) (float64, error) {
	switch vv := v.(type) {
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil

	case int:
		return float64(vv), nil
	case int8:
		return float64(vv), nil
	case int16:
		return float64(vv), nil
	case int32:
		return float64(vv), nil
	case int64:
		return float64(vv), nil

	case uint:
		return float64(vv), nil
	case uint8:
		return float64(vv), nil
	case uint16:
		return float64(vv), nil
	case uint32:
		return float64(vv), nil
	case uint64:
		return float64(vv), nil

	case float32:
		return float64(vv), nil
	case float64:
		return float64(vv), nil

	case big.Float:
		r, _ := vv.Float64()
		return r, nil
	case *big.Float:
		r, _ := vv.Float64()
		return r, nil

	case decimalNumber:
		r, _ := vv.Rat().Float64()
		return r, nil
	case *decimalNumber:
		r, _ := vv.Rat().Float64()
		return r, nil

	case []byte:
		f, err := convFloat(vv)
		r, _ := f.Rat().Float64()
		return r, err
	case string:
		f, err := convFloat(StrToBytesUnsafe(vv))
		r, _ := f.Rat().Float64()
		return r, err
	}
	return 0, _ERR_UNKOWN_TYPE
}

func ToString(v interface{}) (string, error) {
	switch vv := v.(type) {
	case bool:
		return strconv.FormatBool(vv), nil

	case int:
		return strconv.FormatInt(int64(vv), 10), nil
	case int8:
		return strconv.FormatInt(int64(vv), 10), nil
	case int16:
		return strconv.FormatInt(int64(vv), 10), nil
	case int32:
		return strconv.FormatInt(int64(vv), 10), nil
	case int64:
		return strconv.FormatInt(int64(vv), 10), nil

	case uint:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(vv), 10), nil

	case float32:
		return strconv.FormatFloat(float64(vv), 'f', -1, 64), nil
	case float64:
		return strconv.FormatFloat(float64(vv), 'f', -1, 64), nil

	case big.Float:
		return vv.String(), nil
	case *big.Float:
		return vv.String(), nil

	case decimalNumber:
		return vv.String(), nil
	case *decimalNumber:
		return vv.String(), nil

	case []byte:
		return BytesToStrUnsafe(vv), nil
	case string:
		return vv, nil

	case time.Time:
		return vv.Format(timeFormat), nil
	case *time.Time:
		return vv.Format(timeFormat), nil

	case time.Duration:
		return vv.String(), nil
	case *time.Duration:
		return vv.String(), nil
	}
	return "", _ERR_UNKOWN_TYPE
}
func toBigFloat(v interface{}) (*big.Float, error) {
	f := &big.Float{}
	switch vv := v.(type) {
	case bool:
		if vv {
			return f.SetInt64(1), nil
		} else {
			return f.SetInt64(0), nil
		}

	case int:
		return f.SetInt64(int64(vv)), nil
	case int8:
		return f.SetInt64(int64(vv)), nil
	case int16:
		return f.SetInt64(int64(vv)), nil
	case int32:
		return f.SetInt64(int64(vv)), nil
	case int64:
		return f.SetInt64(vv), nil

	case uint:
		return f.SetUint64(uint64(vv)), nil
	case uint8:
		return f.SetUint64(uint64(vv)), nil
	case uint16:
		return f.SetUint64(uint64(vv)), nil
	case uint32:
		return f.SetUint64(uint64(vv)), nil
	case uint64:
		return f.SetUint64(vv), nil

	case float32:
		return f.SetFloat64(float64(vv)), nil
	case float64:
		return f.SetFloat64(vv), nil

	case big.Float:
		return &vv, nil
	case *big.Float:
		return vv, nil

	case decimalNumber:
		return vv.BigFloat(), nil
	case *decimalNumber:
		return vv.BigFloat(), nil

	case []byte:
		d, err := convFloat(vv)
		return d.BigFloat(), err
	case string:
		d, err := convFloat(StrToBytesUnsafe(vv))
		return d.BigFloat(), err
	}

	return newBigFloatZero(), _ERR_UNKOWN_TYPE
}

func toDecimal(v interface{}) (decimalNumber, error) {
	switch vv := v.(type) {
	case bool:
		if vv {
			return newDecimalFromInt(1), nil
		} else {
			return newDecimalZero(), nil
		}

	case int:
		return newDecimalFromInt(int64(vv)), nil
	case int8:
		return newDecimalFromInt(int64(vv)), nil
	case int16:
		return newDecimalFromInt(int64(vv)), nil
	case int32:
		return newDecimalFromInt(int64(vv)), nil
	case int64:
		return newDecimalFromInt(vv), nil

	case uint:
		return newDecimalFromUint(uint64(vv)), nil
	case uint8:
		return newDecimalFromUint(uint64(vv)), nil
	case uint16:
		return newDecimalFromUint(uint64(vv)), nil
	case uint32:
		return newDecimalFromUint(uint64(vv)), nil
	case uint64:
		return newDecimalFromUint(vv), nil

	case float32:
		return newDecimalFromFloat(float64(vv)), nil
	case float64:
		return newDecimalFromFloat(vv), nil

	case big.Float:
		return newDecimalFromBigFloat(&vv), nil
	case *big.Float:
		return newDecimalFromBigFloat(vv), nil

	case decimalNumber:
		return vv, nil
	case *decimalNumber:
		return *vv, nil

	case []byte:
		return convFloat(vv)
	case string:
		return convFloat(StrToBytesUnsafe(vv))
	}
	return newDecimalZero(), _ERR_UNKOWN_TYPE
}
