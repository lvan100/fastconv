package fastconv

import (
	"bytes"
	"sort"
	"strings"
	"sync"
	"unsafe"
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
	Bytes
	Slice
	Map
)

// Value is used to store a value.
// When the Type is Bool, Int, Uint, Float, only the Data field is used.
// When the Type is String, the Data and Length fields are used.
// When the Type is Bytes, the Data, Length, and First fields are used.
type Value struct {
	Type   Kind
	Name   string
	Data   uintptr
	Length int // number of children
	First  int // position of first
	Parent int // position of parent
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

// Bytes returns v's underlying value.
func (p *Value) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&p.Data))
}

// SetBytes sets v's underlying value.
func (p *Value) SetBytes(s []byte) {
	*(*[]byte)(unsafe.Pointer(&p.Data)) = s
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

var bufferPool sync.Pool

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

func PutBuffer(l *Buffer) {
	bufferPool.Put(l)
}

func EqualBuffer(x, y *Buffer) bool {
	if len(x.buf) != len(y.buf) {
		return false
	}
	prepareBuffer(x)
	prepareBuffer(y)
	return equalBuffer(x, y)
}

func prepareBuffer(x *Buffer) {
	for _, a := range x.buf {
		if a.Type == Map && a.Length > 1 {
			sort.Slice(x.buf[a.First:a.First+a.Length], func(i, j int) bool {
				return x.buf[a.First+i].Name > x.buf[a.First+j].Name
			})
		}
	}
}

func equalBuffer(x, y *Buffer) bool {
	for i := 0; i < len(x.buf); i++ {
		a := x.buf[i]
		b := y.buf[i]
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
