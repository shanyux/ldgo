/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/distroy/ldgo/ldcmp"
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

	fprintArg(buf, reflect.ValueOf(args[0]))
	for _, arg := range args[1:] {
		buf.AppendByte(' ')
		fprintArg(buf, reflect.ValueOf(arg))
	}

	buf.TrimNewline()
	text := buf.String()
	buf.Free()

	return text
}

func fprintArg(b *buffer.Buffer, v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		if v.Pointer() == 0 {
			fprintPointer(b, v)
			return
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		fprintSlice(b, v)

	case reflect.Map:
		fprintMap(b, v)

	case reflect.Struct:
		fprintStruct(b, v)

	case reflect.String:
		b.AppendString(v.String())

	case reflect.Bool:
		b.AppendBool(v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b.AppendInt(v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		b.AppendUint(v.Uint())

	case reflect.Float64:
		b.AppendFloat(v.Float(), 64)

	case reflect.Float32:
		b.AppendFloat(v.Float(), 32)

	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		fprintPointer(b, v)

	default:
		fmt.Fprint(b, v.Interface())
	}
}

func fprintSlice(b *buffer.Buffer, v reflect.Value) {
	b.AppendString("[")
	for i := 0; i < v.Len(); i++ {
		if i != 0 {
			b.AppendString(", ")
		}
		fprintArg(b, reflect.ValueOf(v.Index(i).Interface()))
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
		fprintArg(b, reflect.ValueOf(field.Interface()))
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
		fprintArg(b, reflect.ValueOf(kv[0].Interface()))
		b.AppendByte(':')
		fprintArg(b, reflect.ValueOf(kv[1].Interface()))
	}
	b.AppendByte(']')
}

type sortedMap [][2]reflect.Value

func (o sortedMap) Len() int           { return len(o) }
func (o sortedMap) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o sortedMap) Less(i, j int) bool { return ldcmp.CompareReflect(o[i][0], o[j][0]) < 0 }
