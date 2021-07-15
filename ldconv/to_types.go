/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"math/big"
	"strconv"
)

func ToByte(v interface{}) (byte, error) { return ToUint8(v) }

func ToBool(_v interface{}) (bool, error) {
	switch v := _v.(type) {
	case bool:
		return v, nil

	case int:
		return v != 0, nil
	case int8:
		return v != 0, nil
	case int16:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case int64:
		return v != 0, nil

	case uint:
		return v != 0, nil
	case uint8:
		return v != 0, nil
	case uint16:
		return v != 0, nil
	case uint32:
		return v != 0, nil
	case uint64:
		return v != 0, nil

	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil

	case big.Float:
		return v.Cmp(newBigFloatZero()) != 0, nil
	case *big.Float:
		return v.Cmp(newBigFloatZero()) != 0, nil

	case decimalNumber:
		return !v.IsZero(), nil
	case *decimalNumber:
		return !v.IsZero(), nil

	case []byte:
		return convBool(v)
	case string:
		return convBool(StrToBytesUnsafe(v))
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

func ToInt64(_v interface{}) (int64, error) {
	switch v := _v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil

	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil

	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil

	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil

	case big.Float:
		r, _ := v.Int64()
		return r, nil
	case *big.Float:
		r, _ := v.Int64()
		return r, nil

	case decimalNumber:
		return v.BigInt().Int64(), nil
	case *decimalNumber:
		return v.BigInt().Int64(), nil

	case []byte:
		return convInt(v)
	case string:
		return convInt(StrToBytesUnsafe(v))
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

func ToUint64(_v interface{}) (uint64, error) {
	switch v := _v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil

	case int:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil

	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return uint64(v), nil

	case float32:
		return uint64(v), nil
	case float64:
		return uint64(v), nil

	case big.Float:
		r, _ := v.Int64()
		return uint64(r), nil
	case *big.Float:
		r, _ := v.Uint64()
		return uint64(r), nil

	case decimalNumber:
		return v.BigInt().Uint64(), nil
	case *decimalNumber:
		return v.BigInt().Uint64(), nil

	case []byte:
		return convUint(v)
	case string:
		return convUint(StrToBytesUnsafe(v))
	}
	return 0, _ERR_UNKOWN_TYPE
}

func ToFloat32(v interface{}) (float32, error) {
	n, err := ToFloat64(v)
	return float32(n), err
}

func ToFloat64(_v interface{}) (float64, error) {
	switch v := _v.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil

	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil

	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil

	case float32:
		return float64(v), nil
	case float64:
		return float64(v), nil

	case big.Float:
		r, _ := v.Float64()
		return r, nil
	case *big.Float:
		r, _ := v.Float64()
		return r, nil

	case decimalNumber:
		r, _ := v.BigFloat().Float64()
		return r, nil
	case *decimalNumber:
		r, _ := v.BigFloat().Float64()
		return r, nil

	case []byte:
		f, err := convFloat(v)
		r, _ := f.Float64()
		return r, err
	case string:
		f, err := convFloat(StrToBytesUnsafe(v))
		r, _ := f.Float64()
		return r, err
	}
	return 0, _ERR_UNKOWN_TYPE
}

func ToString(_v interface{}) (string, error) {
	switch v := _v.(type) {
	case bool:
		return strconv.FormatBool(v), nil

	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(int64(v), 10), nil

	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(v), 10), nil

	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64), nil

	case big.Float:
		return v.String(), nil
	case *big.Float:
		return v.String(), nil

	case decimalNumber:
		return v.String(), nil
	case *decimalNumber:
		return v.String(), nil

	case []byte:
		return BytesToStrUnsafe(v), nil
	case string:
		return v, nil
	}
	return "", _ERR_UNKOWN_TYPE
}
func toBigFloat(_v interface{}) (*big.Float, error) {
	f := &big.Float{}
	switch v := _v.(type) {
	case bool:
		if v {
			return f.SetInt64(1), nil
		} else {
			return f.SetInt64(0), nil
		}

	case int:
		return f.SetInt64(int64(v)), nil
	case int8:
		return f.SetInt64(int64(v)), nil
	case int16:
		return f.SetInt64(int64(v)), nil
	case int32:
		return f.SetInt64(int64(v)), nil
	case int64:
		return f.SetInt64(v), nil

	case uint:
		return f.SetUint64(uint64(v)), nil
	case uint8:
		return f.SetUint64(uint64(v)), nil
	case uint16:
		return f.SetUint64(uint64(v)), nil
	case uint32:
		return f.SetUint64(uint64(v)), nil
	case uint64:
		return f.SetUint64(v), nil

	case float32:
		return f.SetFloat64(float64(v)), nil
	case float64:
		return f.SetFloat64(v), nil

	case big.Float:
		return &v, nil
	case *big.Float:
		return v, nil

	case decimalNumber:
		return v.BigFloat(), nil
	case *decimalNumber:
		return v.BigFloat(), nil

	case []byte:
		d, err := convFloat(v)
		return d.BigFloat(), err
	case string:
		d, err := convFloat(StrToBytesUnsafe(v))
		return d.BigFloat(), err
	}

	return newBigFloatZero(), _ERR_UNKOWN_TYPE
}

func toDecimal(_v interface{}) (decimalNumber, error) {
	switch v := _v.(type) {
	case bool:
		if v {
			return newDecimalFromInt(1), nil
		} else {
			return newDecimalZero(), nil
		}

	case int:
		return newDecimalFromInt(int64(v)), nil
	case int8:
		return newDecimalFromInt(int64(v)), nil
	case int16:
		return newDecimalFromInt(int64(v)), nil
	case int32:
		return newDecimalFromInt(int64(v)), nil
	case int64:
		return newDecimalFromInt(v), nil

	case uint:
		return newDecimalFromUint(uint64(v)), nil
	case uint8:
		return newDecimalFromUint(uint64(v)), nil
	case uint16:
		return newDecimalFromUint(uint64(v)), nil
	case uint32:
		return newDecimalFromUint(uint64(v)), nil
	case uint64:
		return newDecimalFromUint(v), nil

	case float32:
		return newDecimalFromFloat(float64(v)), nil
	case float64:
		return newDecimalFromFloat(v), nil

	case big.Float:
		return newDecimalFromBigFloat(&v), nil
	case *big.Float:
		return newDecimalFromBigFloat(v), nil

	case decimalNumber:
		return v, nil
	case *decimalNumber:
		return *v, nil

	case []byte:
		return convFloat(v)
	case string:
		return convFloat(StrToBytesUnsafe(v))
	}
	return newDecimalZero(), _ERR_UNKOWN_TYPE
}
