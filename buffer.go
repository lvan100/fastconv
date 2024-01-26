/*
 * Copyright 2024 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fastconv

import (
	"errors"
	"slices"
	"sort"
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

var (
	ZeroValue = Value{}
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

// SetNil sets v's underlying value.
func (p *Value) SetNil() {
	p.Type = Nil
}

// Bool returns v's underlying value.
func (p *Value) Bool() bool {
	return *(*bool)(unsafe.Pointer(&p.Data))
}

// SetBool sets v's underlying value.
func (p *Value) SetBool(b bool) {
	p.Type = Bool
	*(*bool)(unsafe.Pointer(&p.Data)) = b
}

// Int returns v's underlying value.
func (p *Value) Int() int64 {
	return *(*int64)(unsafe.Pointer(&p.Data))
}

// SetInt sets v's underlying value.
func (p *Value) SetInt(i int64) {
	p.Type = Int
	*(*int64)(unsafe.Pointer(&p.Data)) = i
}

// Uint returns v's underlying value.
func (p *Value) Uint() uint64 {
	return *(*uint64)(unsafe.Pointer(&p.Data))
}

// SetUint sets v's underlying value.
func (p *Value) SetUint(u uint64) {
	p.Type = Uint
	*(*uint64)(unsafe.Pointer(&p.Data)) = u
}

// Float returns v's underlying value.
func (p *Value) Float() float64 {
	return *(*float64)(unsafe.Pointer(&p.Data))
}

// SetFloat sets v's underlying value.
func (p *Value) SetFloat(f float64) {
	p.Type = Float
	*(*float64)(unsafe.Pointer(&p.Data)) = f
}

// String returns v's underlying value.
func (p *Value) String() string {
	return *(*string)(unsafe.Pointer(&p.Data))
}

// SetString sets v's underlying value.
func (p *Value) SetString(s string) {
	p.Type = String
	*(*string)(unsafe.Pointer(&p.Data)) = s
	if p.Type != String { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Bools returns v's underlying value.
func (p *Value) Bools() []bool {
	return *(*[]bool)(unsafe.Pointer(&p.Data))
}

// SetBools sets v's underlying value.
func (p *Value) SetBools(s []bool) {
	p.Type = Bools
	*(*[]bool)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Bools { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Ints returns v's underlying value.
func (p *Value) Ints() []int {
	return *(*[]int)(unsafe.Pointer(&p.Data))
}

// SetInts sets v's underlying value.
func (p *Value) SetInts(s []int) {
	p.Type = Ints
	*(*[]int)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Ints { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Int8s returns v's underlying value.
func (p *Value) Int8s() []int8 {
	return *(*[]int8)(unsafe.Pointer(&p.Data))
}

// SetInt8s sets v's underlying value.
func (p *Value) SetInt8s(s []int8) {
	p.Type = Int8s
	*(*[]int8)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Int8s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Int16s returns v's underlying value.
func (p *Value) Int16s() []int16 {
	return *(*[]int16)(unsafe.Pointer(&p.Data))
}

// SetInt16s sets v's underlying value.
func (p *Value) SetInt16s(s []int16) {
	p.Type = Int16s
	*(*[]int16)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Int16s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Int32s returns v's underlying value.
func (p *Value) Int32s() []int32 {
	return *(*[]int32)(unsafe.Pointer(&p.Data))
}

// SetInt32s sets v's underlying value.
func (p *Value) SetInt32s(s []int32) {
	p.Type = Int32s
	*(*[]int32)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Int32s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Int64s returns v's underlying value.
func (p *Value) Int64s() []int64 {
	return *(*[]int64)(unsafe.Pointer(&p.Data))
}

// SetInt64s sets v's underlying value.
func (p *Value) SetInt64s(s []int64) {
	p.Type = Int64s
	*(*[]int64)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Int64s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Uints returns v's underlying value.
func (p *Value) Uints() []uint {
	return *(*[]uint)(unsafe.Pointer(&p.Data))
}

// SetUints sets v's underlying value.
func (p *Value) SetUints(s []uint) {
	p.Type = Uints
	*(*[]uint)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Uints { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Uint8s returns v's underlying value.
func (p *Value) Uint8s() []uint8 {
	return *(*[]uint8)(unsafe.Pointer(&p.Data))
}

// SetUint8s sets v's underlying value.
func (p *Value) SetUint8s(s []uint8) {
	p.Type = Uint8s
	*(*[]uint8)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Uint8s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Uint16s returns v's underlying value.
func (p *Value) Uint16s() []uint16 {
	return *(*[]uint16)(unsafe.Pointer(&p.Data))
}

// SetUint16s sets v's underlying value.
func (p *Value) SetUint16s(s []uint16) {
	p.Type = Uint16s
	*(*[]uint16)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Uint16s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Uint32s returns v's underlying value.
func (p *Value) Uint32s() []uint32 {
	return *(*[]uint32)(unsafe.Pointer(&p.Data))
}

// SetUint32s sets v's underlying value.
func (p *Value) SetUint32s(s []uint32) {
	p.Type = Uint32s
	*(*[]uint32)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Uint32s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Uint64s returns v's underlying value.
func (p *Value) Uint64s() []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&p.Data))
}

// SetUint64s sets v's underlying value.
func (p *Value) SetUint64s(s []uint64) {
	p.Type = Uint64s
	*(*[]uint64)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Uint64s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Float32s returns v's underlying value.
func (p *Value) Float32s() []float32 {
	return *(*[]float32)(unsafe.Pointer(&p.Data))
}

// SetFloat32s sets v's underlying value.
func (p *Value) SetFloat32s(s []float32) {
	p.Type = Float32s
	*(*[]float32)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Float32s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Float64s returns v's underlying value.
func (p *Value) Float64s() []float64 {
	return *(*[]float64)(unsafe.Pointer(&p.Data))
}

// SetFloat64s sets v's underlying value.
func (p *Value) SetFloat64s(s []float64) {
	p.Type = Float64s
	*(*[]float64)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Float64s { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// Strings returns v's underlying value.
func (p *Value) Strings() []string {
	return *(*[]string)(unsafe.Pointer(&p.Data))
}

// SetStrings sets v's underlying value.
func (p *Value) SetStrings(s []string) {
	p.Type = Strings
	*(*[]string)(unsafe.Pointer(&p.Data)) = s
	if p.Type != Strings { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// A Buffer is a variable-sized buffer of Value.
type Buffer struct {
	buf []Value
}

// Reset resets the buffer to be empty.
func (b *Buffer) Reset() {
	for i := 0; i < len(b.buf); i++ {
		b.buf[i] = ZeroValue
	}
	b.buf = b.buf[:0]
}

// Grow grows the buffer to guarantee space for n more [Value]s.
func (b *Buffer) Grow(n int) {
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
	b.buf = b.buf[:l+n]
}

func (b *Buffer) Nil(name string) *Buffer {
	v := Value{Name: name}
	v.SetNil()
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Bool(name string, x bool) *Buffer {
	v := Value{Name: name}
	v.SetBool(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Int(name string, x int64) *Buffer {
	v := Value{Name: name}
	v.SetInt(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Uint(name string, x uint64) *Buffer {
	v := Value{Name: name}
	v.SetUint(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Float(name string, x float64) *Buffer {
	v := Value{Name: name}
	v.SetFloat(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) String(name string, x string) *Buffer {
	v := Value{Name: name}
	v.SetString(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Bools(name string, x []bool) *Buffer {
	v := Value{Name: name}
	v.SetBools(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Ints(name string, x []int) *Buffer {
	v := Value{Name: name}
	v.SetInts(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Int8s(name string, x []int8) *Buffer {
	v := Value{Name: name}
	v.SetInt8s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Int16s(name string, x []int16) *Buffer {
	v := Value{Name: name}
	v.SetInt16s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Int32s(name string, x []int32) *Buffer {
	v := Value{Name: name}
	v.SetInt32s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Int64s(name string, x []int64) *Buffer {
	v := Value{Name: name}
	v.SetInt64s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Uints(name string, x []uint) *Buffer {
	v := Value{Name: name}
	v.SetUints(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Uint8s(name string, x []uint8) *Buffer {
	v := Value{Name: name}
	v.SetUint8s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Uint16s(name string, x []uint16) *Buffer {
	v := Value{Name: name}
	v.SetUint16s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Uint32s(name string, x []uint32) *Buffer {
	v := Value{Name: name}
	v.SetUint32s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Uint64s(name string, x []uint64) *Buffer {
	v := Value{Name: name}
	v.SetUint64s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Float32s(name string, x []float32) *Buffer {
	v := Value{Name: name}
	v.SetFloat32s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Float64s(name string, x []float64) *Buffer {
	v := Value{Name: name}
	v.SetFloat64s(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Strings(name string, x []string) *Buffer {
	v := Value{Name: name}
	v.SetStrings(x)
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Slice(name string, length, first int) *Buffer {
	v := Value{
		Length: length,
		First:  first,
		Type:   Slice,
		Name:   name,
	}
	b.buf = append(b.buf, v)
	return b
}

func (b *Buffer) Map(name string, length, first int) *Buffer {
	v := Value{
		Length: length,
		First:  first,
		Type:   Map,
		Name:   name,
	}
	b.buf = append(b.buf, v)
	return b
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
