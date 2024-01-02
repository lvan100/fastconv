package fastconv

import (
	"bytes"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"

	"github.com/lvan100/fastconv/internal/assert"
)

func equalValues(got, expect []Value) bool {
	if len(got) != len(expect) {
		return false
	}
	for i := 0; i < len(got); i++ {
		a := got[i]
		b := expect[i]
		if a.Type != b.Type || a.Name != b.Name || a.Parent != b.Parent {
			return false
		}
		switch a.Type {
		case Nil, Bool, Int, Uint, Float, Slice, Map:
			if a.Data != b.Data || a.Length != b.Length || a.First != b.First {
				return false
			}
		case String:
			s1 := *(*string)(unsafe.Pointer(&a.Data))
			s2 := *(*string)(unsafe.Pointer(&b.Data))
			if strings.Compare(s1, s2) != 0 || a.First != b.First {
				return false
			}
		case Bytes:
			s1 := *(*[]byte)(unsafe.Pointer(&a.Data))
			s2 := *(*[]byte)(unsafe.Pointer(&b.Data))
			if bytes.Compare(s1, s2) != 0 {
				return false
			}
		}
	}
	return true
}

func Test_encodeValue(t *testing.T) {
	type args struct {
		value   interface{}
		buffer  func() []Value
		encoder map[reflect.Type]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bool",
			args: args{
				value: true,
				buffer: func() []Value {
					v := Value{Type: Bool}
					v.SetBool(true)
					return []Value{v}
				},
				encoder: map[reflect.Type]string{},
			},
		},
		{
			name: "*bool",
			args: args{
				value: Ptr(false),
				buffer: func() []Value {
					v := Value{Type: Bool}
					v.SetBool(false)
					return []Value{v}
				},
				encoder: map[reflect.Type]string{
					TypeFor[*bool](): "ptrEncoder.encode-fm",
				},
			},
		},
		{
			name: "int",
			args: args{
				value: int(3),
				buffer: func() []Value {
					v := Value{Type: Int}
					v.SetInt(int64(3))
					return []Value{v}
				},
				encoder: map[reflect.Type]string{},
			},
		},
		{
			name: "*int64",
			args: args{
				value: Ptr(int64(3)),
				buffer: func() []Value {
					v := Value{Type: Int}
					v.SetInt(int64(3))
					return []Value{v}
				},
				encoder: map[reflect.Type]string{
					TypeFor[*int64](): "ptrEncoder.encode-fm",
				},
			},
		},
		{
			name: "uint",
			args: args{
				value: uint(3),
				buffer: func() []Value {
					v := Value{Type: Uint}
					v.SetUint(uint64(3))
					return []Value{v}
				},
				encoder: map[reflect.Type]string{},
			},
		},
		{
			name: "*uint64",
			args: args{
				value: Ptr(uint64(3)),
				buffer: func() []Value {
					v := Value{Type: Uint}
					v.SetUint(uint64(3))
					return []Value{v}
				},
				encoder: map[reflect.Type]string{
					TypeFor[*uint64](): "ptrEncoder.encode-fm",
				},
			},
		},
		{
			name: "float32",
			args: args{
				value: float32(3),
				buffer: func() []Value {
					v := Value{Type: Float}
					v.SetFloat(float64(3))
					return []Value{v}
				},
				encoder: map[reflect.Type]string{},
			},
		},
		{
			name: "*float64",
			args: args{
				value: Ptr(float64(3)),
				buffer: func() []Value {
					v := Value{Type: Float}
					v.SetFloat(float64(3))
					return []Value{v}
				},
				encoder: map[reflect.Type]string{
					TypeFor[*float64](): "ptrEncoder.encode-fm",
				},
			},
		},
		{
			name: "string",
			args: args{
				value: "3",
				buffer: func() []Value {
					v := Value{Type: String}
					v.SetString("3")
					return []Value{v}
				},
				encoder: map[reflect.Type]string{},
			},
		},
		{
			name: "*string",
			args: args{
				value: Ptr("3"),
				buffer: func() []Value {
					v := Value{Type: String}
					v.SetString("3")
					return []Value{v}
				},
				encoder: map[reflect.Type]string{
					TypeFor[*string](): "ptrEncoder.encode-fm",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := reflect.ValueOf(tt.args.value)
			l := &Buffer{}
			l.Append(1)
			encodeValue(l, 0, &l.buf[0], v)
			assert.True(t, equalValues(l.buf, tt.args.buffer()))

			encoder := make(map[reflect.Type]string)
			encoderCache.Range(func(key, value any) bool {
				fn := runtime.FuncForPC(reflect.ValueOf(value.(encoderFunc)).Pointer())
				encoder[key.(reflect.Type)] = strings.TrimPrefix(fn.Name(), "github.com/lvan100/fastconv.")
				encoderCache.Delete(key)
				return true
			})
			assert.Equal(t, encoder, tt.args.encoder)
		})
	}
}
