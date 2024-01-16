package fastconv

import (
	"sort"
	"sync"
	"unsafe"

	"github.com/lvan100/fastconv/internal/slices"
)

// A Kind represents the type of value stored in Value.
type Kind int

const (
	Invalid = Kind(iota)
	Nil
	Bool
	Int
	Uint
	Float
	String
	Bools
	Ints
	Int8s
	Int16s
	Int32s
	Int64s
	Uints
	Uint8s
	Uint16s
	Uint32s
	Uint64s
	Float32s
	Float64s
	Strings
	Slice
	Map
)

// Value is used to store a value.
// When the Type is Bool, Int, Uint, Float, only the Data field is used.
// When the Type is String, the Data and Length fields are used.
// When the Type is [Type]s, the Data, Length, and First fields are used.
type Value struct {
	Data   uintptr
	Length int  // number of children
	First  int  // position of first
	Type   Kind // to prevent overflow
	Name   string
}

// Bool returns v's underlying value.
func (p *Value) Bool() bool {
	return *(*bool)(unsafe.Pointer(&p.Data))
}

// SetBool sets v's underlying value.
func (p *Value) SetBool(b bool) {
	*(*bool)(unsafe.Pointer(&p.Data)) = b
}

// Int returns v's underlying value.
func (p *Value) Int() int64 {
	return *(*int64)(unsafe.Pointer(&p.Data))
}

// SetInt sets v's underlying value.
func (p *Value) SetInt(i int64) {
	*(*int64)(unsafe.Pointer(&p.Data)) = i
}

// Uint returns v's underlying value.
func (p *Value) Uint() uint64 {
	return *(*uint64)(unsafe.Pointer(&p.Data))
}

// SetUint sets v's underlying value.
func (p *Value) SetUint(u uint64) {
	*(*uint64)(unsafe.Pointer(&p.Data)) = u
}

// Float returns v's underlying value.
func (p *Value) Float() float64 {
	return *(*float64)(unsafe.Pointer(&p.Data))
}

// SetFloat sets v's underlying value.
func (p *Value) SetFloat(f float64) {
	*(*float64)(unsafe.Pointer(&p.Data)) = f
}

// String returns v's underlying value.
func (p *Value) String() string {
	return *(*string)(unsafe.Pointer(&p.Data))
}

// SetString sets v's underlying value.
func (p *Value) SetString(s string) {
	*(*string)(unsafe.Pointer(&p.Data)) = s
}

// Bools returns v's underlying value.
func (p *Value) Bools() []bool {
	return *(*[]bool)(unsafe.Pointer(&p.Data))
}

// SetBools sets v's underlying value.
func (p *Value) SetBools(s []bool) {
	*(*[]bool)(unsafe.Pointer(&p.Data)) = s
}

// Ints returns v's underlying value.
func (p *Value) Ints() []int {
	return *(*[]int)(unsafe.Pointer(&p.Data))
}

// SetInts sets v's underlying value.
func (p *Value) SetInts(s []int) {
	*(*[]int)(unsafe.Pointer(&p.Data)) = s
}

// Int8s returns v's underlying value.
func (p *Value) Int8s() []int8 {
	return *(*[]int8)(unsafe.Pointer(&p.Data))
}

// SetInt8s sets v's underlying value.
func (p *Value) SetInt8s(s []int8) {
	*(*[]int8)(unsafe.Pointer(&p.Data)) = s
}

// Int16s returns v's underlying value.
func (p *Value) Int16s() []int16 {
	return *(*[]int16)(unsafe.Pointer(&p.Data))
}

// SetInt16s sets v's underlying value.
func (p *Value) SetInt16s(s []int16) {
	*(*[]int16)(unsafe.Pointer(&p.Data)) = s
}

// Int32s returns v's underlying value.
func (p *Value) Int32s() []int32 {
	return *(*[]int32)(unsafe.Pointer(&p.Data))
}

// SetInt32s sets v's underlying value.
func (p *Value) SetInt32s(s []int32) {
	*(*[]int32)(unsafe.Pointer(&p.Data)) = s
}

// Int64s returns v's underlying value.
func (p *Value) Int64s() []int64 {
	return *(*[]int64)(unsafe.Pointer(&p.Data))
}

// SetInt64s sets v's underlying value.
func (p *Value) SetInt64s(s []int64) {
	*(*[]int64)(unsafe.Pointer(&p.Data)) = s
}

// Uints returns v's underlying value.
func (p *Value) Uints() []uint {
	return *(*[]uint)(unsafe.Pointer(&p.Data))
}

// SetUints sets v's underlying value.
func (p *Value) SetUints(s []uint) {
	*(*[]uint)(unsafe.Pointer(&p.Data)) = s
}

// Uint8s returns v's underlying value.
func (p *Value) Uint8s() []uint8 {
	return *(*[]uint8)(unsafe.Pointer(&p.Data))
}

// SetUint8s sets v's underlying value.
func (p *Value) SetUint8s(s []uint8) {
	*(*[]uint8)(unsafe.Pointer(&p.Data)) = s
}

// Uint16s returns v's underlying value.
func (p *Value) Uint16s() []uint16 {
	return *(*[]uint16)(unsafe.Pointer(&p.Data))
}

// SetUint16s sets v's underlying value.
func (p *Value) SetUint16s(s []uint16) {
	*(*[]uint16)(unsafe.Pointer(&p.Data)) = s
}

// Uint32s returns v's underlying value.
func (p *Value) Uint32s() []uint32 {
	return *(*[]uint32)(unsafe.Pointer(&p.Data))
}

// SetUint32s sets v's underlying value.
func (p *Value) SetUint32s(s []uint32) {
	*(*[]uint32)(unsafe.Pointer(&p.Data)) = s
}

// Uint64s returns v's underlying value.
func (p *Value) Uint64s() []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&p.Data))
}

// SetUint64s sets v's underlying value.
func (p *Value) SetUint64s(s []uint64) {
	*(*[]uint64)(unsafe.Pointer(&p.Data)) = s
}

// Float32s returns v's underlying value.
func (p *Value) Float32s() []float32 {
	return *(*[]float32)(unsafe.Pointer(&p.Data))
}

// SetFloat32s sets v's underlying value.
func (p *Value) SetFloat32s(s []float32) {
	*(*[]float32)(unsafe.Pointer(&p.Data)) = s
}

// Float64s returns v's underlying value.
func (p *Value) Float64s() []float64 {
	return *(*[]float64)(unsafe.Pointer(&p.Data))
}

// SetFloat64s sets v's underlying value.
func (p *Value) SetFloat64s(s []float64) {
	*(*[]float64)(unsafe.Pointer(&p.Data)) = s
}

// Strings returns v's underlying value.
func (p *Value) Strings() []string {
	return *(*[]string)(unsafe.Pointer(&p.Data))
}

// SetStrings sets v's underlying value.
func (p *Value) SetStrings(s []string) {
	*(*[]string)(unsafe.Pointer(&p.Data)) = s
}

// A Buffer is a variable-sized buffer of Value.
type Buffer struct {
	buf []Value
}

// Reset resets the buffer to be empty.
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}

// Append appends n [Value]s to the buffer, growing if needed.
func (b *Buffer) Append(n int) {
	b.grow(n)
	for i := 0; i < n; i++ {
		b.buf = append(b.buf, Value{})
	}
}

// grow grows the buffer to guarantee space for n more [Value]s.
func (b *Buffer) grow(n int) {
	c := cap(b.buf)
	l := len(b.buf)
	if l+n > c {
		if c < 1024 {
			c *= 2
		} else {
			c += c / 4
		}
		buf := make([]Value, l, c)
		copy(buf, b.buf[:l])
		b.buf = buf
	}
}

// bufferPool pools the [Buffer]s.
var bufferPool sync.Pool

// GetBuffer gets a Buffer from the pool.
func GetBuffer() *Buffer {
	if v := bufferPool.Get(); v != nil {
		e := v.(*Buffer)
		e.Reset()
		return e
	}
	return &Buffer{
		buf: make([]Value, 0, 512),
	}
}

// PutBuffer returns a Buffer to the pool.
func PutBuffer(l *Buffer) {
	bufferPool.Put(l)
}

// EqualBuffer reports whether x and y are the same [Value]s.
func EqualBuffer(x, y *Buffer) bool {
	if len(x.buf) != len(y.buf) {
		return false
	}
	sortMap(x)
	sortMap(y)
	return equalBuffer(x, y)
}

func sortMap(x *Buffer) {
	for _, a := range x.buf {
		if a.Type == Map && a.Length > 1 {
			start := a.First
			end := a.First + a.Length
			sort.Slice(x.buf[start:end], func(i, j int) bool {
				return x.buf[start+i].Name > x.buf[start+j].Name
			})
		}
	}
}

func equalBuffer(x, y *Buffer) bool {
	for i := 0; i < len(x.buf); i++ {
		a, b := x.buf[i], y.buf[i]
		if a.Type != b.Type || a.Name != b.Name {
			return false
		}
		switch a.Type {
		case Nil, Bool, Int, Uint, Float, Slice, Map:
			if a.Data != b.Data || a.Length != b.Length || a.First != b.First {
				return false
			}
		case String:
			if a.String() != b.String() || a.First != b.First {
				return false
			}
		case Bools:
			if !slices.Equal(a.Bools(), b.Bools()) {
				return false
			}
		case Ints:
			if !slices.Equal(a.Ints(), b.Ints()) {
				return false
			}
		case Int8s:
			if !slices.Equal(a.Int8s(), b.Int8s()) {
				return false
			}
		case Int16s:
			if !slices.Equal(a.Int16s(), b.Int16s()) {
				return false
			}
		case Int32s:
			if !slices.Equal(a.Int32s(), b.Int32s()) {
				return false
			}
		case Int64s:
			if !slices.Equal(a.Int64s(), b.Int64s()) {
				return false
			}
		case Uints:
			if !slices.Equal(a.Uints(), b.Uints()) {
				return false
			}
		case Uint8s:
			if !slices.Equal(a.Uint8s(), b.Uint8s()) {
				return false
			}
		case Uint16s:
			if !slices.Equal(a.Uint16s(), b.Uint16s()) {
				return false
			}
		case Uint32s:
			if !slices.Equal(a.Uint32s(), b.Uint32s()) {
				return false
			}
		case Uint64s:
			if !slices.Equal(a.Uint64s(), b.Uint64s()) {
				return false
			}
		case Float32s:
			if !slices.Equal(a.Float32s(), b.Float32s()) {
				return false
			}
		case Float64s:
			if !slices.Equal(a.Float64s(), b.Float64s()) {
				return false
			}
		case Strings:
			if !slices.Equal(a.Strings(), b.Strings()) {
				return false
			}
		}
	}
	return true
}
