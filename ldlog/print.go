/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"fmt"
	"reflect"
	"sort"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug  = zapcore.DebugLevel
	LevelInfo   = zapcore.InfoLevel
	LevelWarn   = zapcore.WarnLevel
	LevelError  = zapcore.ErrorLevel
	LevelDpanic = zapcore.DPanicLevel
	LevelPanic  = zapcore.PanicLevel
	LevelFatal  = zapcore.FatalLevel
)

type printWrapper struct {
	args []interface{}
}

func (w printWrapper) String() string {
	return sprintln(w.args)
}

func pw(args []interface{}) fmt.Stringer { return printWrapper{args: args} }

func sprintln(args []interface{}) string {
	if len(args) == 0 {
		return ""
	}

	buf := bufferpool.Get()

	fprintArg(buf, args[0])
	for _, arg := range args[1:] {
		buf.AppendByte(' ')
		fprintArg(buf, arg)
	}

	buf.TrimNewline()
	text := buf.String()
	buf.Free()

	return text
}

func fprintArg(b *buffer.Buffer, val interface{}) {
	switch v := val.(type) {
	case fmt.Stringer:
		b.AppendString(v.String())
		return

	case error:
		b.AppendString(v.Error())
		return
	}

	ref := reflect.ValueOf(val)
	if ref.Kind() == reflect.Ptr {
		if ref.Pointer() == 0 {
			fprintPointer(b, ref)
			return
		}
		ref = ref.Elem()
	}

	switch ref.Kind() {
	case reflect.Array, reflect.Slice:
		fprintSlice(b, ref)

	case reflect.Map:
		fprintMap(b, ref)

	case reflect.Struct:
		fprintStruct(b, ref)

	case reflect.String:
		b.AppendString(ref.String())

	case reflect.Bool:
		b.AppendBool(ref.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b.AppendInt(ref.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		b.AppendUint(ref.Uint())

	case reflect.Float64:
		b.AppendFloat(ref.Float(), 64)

	case reflect.Float32:
		b.AppendFloat(ref.Float(), 32)

	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		fprintPointer(b, ref)

	default:
		fmt.Fprint(b, ref.Interface())
	}
}

func fprintSlice(b *buffer.Buffer, v reflect.Value) {
	b.AppendString("[")
	for i := 0; i < v.Len(); i++ {
		if i != 0 {
			b.AppendString(", ")
		}
		fprintArg(b, v.Index(i).Interface())
	}
	b.AppendString("]")
}

func fprintPointer(b *buffer.Buffer, v reflect.Value) {
	p := v.Pointer()

	b.AppendByte('(')
	b.AppendString(v.Type().String())
	b.AppendString(")(")
	if p == 0 {
		b.AppendString("nil")
	} else {
		fmt.Fprintf(b, "0x%x", p)
	}
	b.AppendByte(')')
}

func fprintStruct(b *buffer.Buffer, v reflect.Value) {
	b.AppendByte('{')
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.AppendByte(',')
		}
		if name := v.Type().Field(i).Name; name != "" {
			b.AppendString(name)
			b.AppendByte(':')
		}
		field := v.Field(i)
		fprintArg(b, field.Interface())
	}
	b.AppendByte('}')
}

func fprintMap(b *buffer.Buffer, val reflect.Value) {
	m := make([][2]reflect.Value, 0, val.Len())
	for it := val.MapRange(); it.Next(); {
		m = append(m, [2]reflect.Value{it.Key(), it.Value()})
	}

	sort.Sort(sortedMap(m))

	b.AppendString("map[")
	for i, kv := range m {
		if i > 0 {
			b.AppendByte(',')
		}
		fprintArg(b, kv[0].Interface())
		b.AppendByte(':')
		fprintArg(b, kv[1].Interface())
	}
	b.AppendByte(']')
}

type sortedMap [][2]reflect.Value

func (o sortedMap) Len() int           { return len(o) }
func (o sortedMap) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o sortedMap) Less(i, j int) bool { return o.compare(o[i][0], o[j][0]) < 0 }

func (o sortedMap) compare(aVal, bVal reflect.Value) int {
	aType, bType := aVal.Type(), bVal.Type()
	if aType != bType {
		return -1 // No good answer possible, but don't return 0: they're not equal.
	}
	switch aVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a, b := aVal.Int(), bVal.Int()
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		a, b := aVal.Uint(), bVal.Uint()
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	case reflect.String:
		a, b := aVal.String(), bVal.String()
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	case reflect.Float32, reflect.Float64:
		return o.floatCompare(aVal.Float(), bVal.Float())
	case reflect.Complex64, reflect.Complex128:
		a, b := aVal.Complex(), bVal.Complex()
		if c := o.floatCompare(real(a), real(b)); c != 0 {
			return c
		}
		return o.floatCompare(imag(a), imag(b))
	case reflect.Bool:
		a, b := aVal.Bool(), bVal.Bool()
		switch {
		case a == b:
			return 0
		case a:
			return 1
		default:
			return -1
		}
	case reflect.Ptr, reflect.UnsafePointer:
		a, b := aVal.Pointer(), bVal.Pointer()
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	case reflect.Chan:
		if c, ok := o.nilCompare(aVal, bVal); ok {
			return c
		}
		ap, bp := aVal.Pointer(), bVal.Pointer()
		switch {
		case ap < bp:
			return -1
		case ap > bp:
			return 1
		default:
			return 0
		}
	case reflect.Struct:
		for i := 0; i < aVal.NumField(); i++ {
			if c := o.compare(aVal.Field(i), bVal.Field(i)); c != 0 {
				return c
			}
		}
		return 0
	case reflect.Array:
		for i := 0; i < aVal.Len(); i++ {
			if c := o.compare(aVal.Index(i), bVal.Index(i)); c != 0 {
				return c
			}
		}
		return 0
	case reflect.Interface:
		if c, ok := o.nilCompare(aVal, bVal); ok {
			return c
		}
		c := o.compare(reflect.ValueOf(aVal.Elem().Type()), reflect.ValueOf(bVal.Elem().Type()))
		if c != 0 {
			return c
		}
		return o.compare(aVal.Elem(), bVal.Elem())
	default:
		// Certain types cannot appear as keys (maps, funcs, slices), but be explicit.
		panic("bad type in compare: " + aType.String())
	}
}
func (o sortedMap) nilCompare(aVal, bVal reflect.Value) (int, bool) {
	if aVal.IsNil() {
		if bVal.IsNil() {
			return 0, true
		}
		return -1, true
	}
	if bVal.IsNil() {
		return 1, true
	}
	return 0, false
}

func (o sortedMap) floatCompare(a, b float64) int {
	switch {
	case o.isNaN(a):
		return -1 // No good answer if b is a NaN so don't bother checking.
	case o.isNaN(b):
		return 1
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

func (o sortedMap) isNaN(a float64) bool {
	return a != a
}
