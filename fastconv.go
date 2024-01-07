/*
 * Copyright 2023 the original author or authors.
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
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unsafe"
)

func Convert(src, dest any) error {

	srcValue := reflect.ValueOf(src)
	if !srcValue.IsValid() || srcValue.IsNil() {
		return nil
	}

	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.IsNil() {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(destValue)}
	}

	l := newBuffer()
	defer bufferPool.Put(l)

	e := &encodeState{Buffer: l}
	e.Append(1)
	c := &e.buf[0]
	c.Parent = -1
	encodeValue(e, 0, c, srcValue)

	//fmt.Println("----------------")
	//n := len(l.buf)
	//for i := 0; i < n; i++ {
	//	p := l.buf[i]
	//	if p.Type == Slice || p.Type == Map {
	//		if p.Length > 0 {
	//			fmt.Println("current:", i, "parent:", p.Parent, "children", p.First, "~", p.First+p.Length-1)
	//		} else {
	//			fmt.Println("current:", i, "parent:", p.Parent)
	//		}
	//	} else {
	//		fmt.Println("current:", i, "parent:", p.Parent)
	//	}
	//}

	d := &decodeState{Buffer: l}
	decodeValue(d, &d.buf[0], destValue)
	return d.savedError
}

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

func newBuffer() *Buffer {
	if v := bufferPool.Get(); v != nil {
		e := v.(*Buffer)
		e.Reset()
		return e
	}
	return &Buffer{
		buf: make([]Value, 0, 512),
	}
}

type encodeState struct {
	*Buffer
	savedError error
}

func (e *encodeState) addErrorContext(p *Value, err error) {

}

// encodeValue encodes v to p [*Value] which stored in l [*Buffer].
func encodeValue(e *encodeState, current int, p *Value, v reflect.Value) {
	if v.IsValid() {
		cachedTypeEncoder(v.Type())(e, current, p, v)
	}
}

type encoderFunc func(e *encodeState, current int, p *Value, v reflect.Value)

var (
	fastEncoders []encoderFunc
	encoderCache sync.Map // map[reflect.Type]encoderFunc
	stFieldCache sync.Map // map[reflect.Type]structFields
)

// Ptr returns the pointer to the given value.
func Ptr[T any](t T) *T {
	return &t
}

// TypeFor returns the [reflect.Type] that represents the type T.
func TypeFor[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

var (
	typeSliceInterface     = TypeFor[[]interface{}]()
	typeMapStringInterface = TypeFor[map[string]interface{}]()
)

func init() {
	fastEncoders = []encoderFunc{
		nil,              // Invalid
		boolEncoder,      // Bool
		intEncoder,       // Int
		intEncoder,       // Int8
		intEncoder,       // Int16
		intEncoder,       // Int32
		intEncoder,       // Int64
		uintEncoder,      // Uint
		uintEncoder,      // Uint8
		uintEncoder,      // Uint16
		uintEncoder,      // Uint32
		uintEncoder,      // Uint64
		nil,              // Uintptr
		floatEncoder,     // Float32
		floatEncoder,     // Float64
		nil,              // Complex64
		nil,              // Complex128
		nil,              // Array
		nil,              // Chan
		nil,              // Func
		interfaceEncoder, // Interface
		nil,              // Map
		nil,              // Pointer
		nil,              // Slice
		stringEncoder,    // String
		nil,              // Struct
		nil,              // UnsafePointer
	}
}

// cachedTypeEncoder gets the encoderFunc stored in the cache of type t.
func cachedTypeEncoder(t reflect.Type) encoderFunc {

	// fast judgement to avoid concurrent competition
	switch k := t.Kind(); k {
	case reflect.Slice:
		if t == typeSliceInterface {
			return sliceInterfaceEncoder
		}
	case reflect.Map:
		if t == typeMapStringInterface {
			return mapStringInterfaceEncoder
		}
	default:
		if f := fastEncoders[k]; f != nil {
			return f
		}
	}

	if fi, ok := encoderCache.Load(t); ok {
		return fi.(encoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoderCache.LoadOrStore(t, encoderFunc(func(e *encodeState, current int, p *Value, v reflect.Value) {
		wg.Wait()
		f(e, current, p, v)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = newTypeEncoder(t)
	wg.Done()
	encoderCache.Store(t, f)
	return f
}

// newTypeEncoder constructs an encoderFunc for a type.
func newTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Pointer:
		return newPtrEncoder(t)
	case reflect.Array, reflect.Slice:
		return newArrayEncoder(t)
	case reflect.Map:
		return newMapEncoder(t)
	case reflect.Struct:
		return newStructEncoder(t)
	default:
		return func(e *encodeState, current int, p *Value, v reflect.Value) {}
	}
}

// boolEncoder is the encoderFunc of bool.
func boolEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	p.Type = Bool
	p.SetBool(v.Bool())
}

// intEncoder is the encoderFunc of int, int8, int16, int32 and int64.
func intEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	p.Type = Int
	p.SetInt(v.Int())
}

// uintEncoder is the encoderFunc of uint, uint8, uint16, uint32 and uint64.
func uintEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	p.Type = Uint
	p.SetUint(v.Uint())
}

// floatEncoder is the encoderFunc of float32 and float64.
func floatEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	p.Type = Float
	p.SetFloat(v.Float())
}

// stringEncoder is the encoderFunc of string.
func stringEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	p.Type = String
	old := p.Parent
	p.SetString(v.String())
	if old != p.Parent { // should never happen
		panic(fmt.Errorf("!!! parent was unexpectedly modified"))
	}
}

// interfaceEncoder is the encoderFunc of interface{}.
func interfaceEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	if v.IsNil() {
		p.Type = Nil
		return
	}
	encodeValue(e, current, p, v.Elem())
}

// sliceByteEncoder is the encoderFunc of []byte.
func sliceByteEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	p.Type = Bytes
	old := p.Parent
	p.SetBytes(v.Interface().([]byte))
	if p.Parent != old { // should never happen
		panic(fmt.Errorf("!!! parent was unexpectedly modified"))
	}
}

// sliceInterfaceEncoder is the encoderFunc of []interface{}.
func sliceInterfaceEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	s := v.Interface().([]interface{})
	n := len(s)
	p.Type = Slice
	if n == 0 {
		return
	}
	p.Length = n
	end := len(e.buf)
	p.First = end
	e.Append(n)
	for i, sValue := range s {
		if sValue == nil {
			e.buf[end+i] = Value{Type: Nil, Parent: current}
		} else {
			c := &e.buf[end+i]
			c.Parent = current
			encodeValue(e, end+i, c, reflect.ValueOf(sValue))
		}
	}
}

// mapStringInterfaceEncoder is the encoderFunc of map[string]interface{}.
func mapStringInterfaceEncoder(e *encodeState, current int, p *Value, v reflect.Value) {
	m := v.Interface().(map[string]interface{})
	n := len(m)
	p.Type = Map
	if n == 0 {
		return
	}
	p.Length = n
	end := len(e.buf)
	p.First = end
	e.Append(n)
	i := 0
	for mKey, mValue := range m { // no need to sort keys
		if mValue == nil {
			e.buf[end+i] = Value{Type: Nil, Name: mKey, Parent: current}
		} else {
			c := &e.buf[end+i]
			c.Name = mKey
			c.Parent = current
			encodeValue(e, end+i, c, reflect.ValueOf(mValue))
		}
		i++
	}
}

// ptrEncoder is the encoderFunc of pointer type (*T).
type ptrEncoder struct {
	elemEnc encoderFunc
}

func newPtrEncoder(t reflect.Type) encoderFunc {
	e := ptrEncoder{cachedTypeEncoder(t.Elem())}
	return e.encode
}

func (pe ptrEncoder) encode(e *encodeState, current int, p *Value, v reflect.Value) {
	if v.IsNil() {
		p.Type = Nil
		return
	}
	pe.elemEnc(e, current, p, v.Elem())
}

// arrayEncoder is the encoderFunc of array and slice.
type arrayEncoder struct {
	elemEnc encoderFunc
}

func newArrayEncoder(t reflect.Type) encoderFunc {
	e := arrayEncoder{cachedTypeEncoder(t.Elem())}
	return e.encode
}

func (ae arrayEncoder) encode(e *encodeState, current int, p *Value, v reflect.Value) {
	n := v.Len()
	p.Type = Slice
	if n == 0 {
		return
	}
	p.Length = n
	end := len(e.buf)
	p.First = end
	e.Append(n)
	for i := 0; i < n; i++ {
		c := &e.buf[end+i]
		c.Parent = current
		ae.elemEnc(e, end+i, c, v.Index(i))
	}
}

// mapEncoder is the encoderFunc of map.
type mapEncoder struct {
	elemEnc encoderFunc
}

func newMapEncoder(t reflect.Type) encoderFunc {
	e := mapEncoder{cachedTypeEncoder(t.Elem())}
	return e.encode
}

func validMapKey(k reflect.Value) (string, bool) {
	switch k.Kind() {
	case reflect.String:
		return k.String(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(k.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(k.Uint(), 10), true
	default:
		return fmt.Sprint(k.Interface()), false
	}
}

func (me mapEncoder) encode(e *encodeState, current int, p *Value, v reflect.Value) {
	n := v.Len()
	p.Type = Map
	if n == 0 {
		return
	}
	p.Length = n
	end := len(e.buf)
	p.First = end
	e.Append(n)
	i := 0
	iter := v.MapRange()
	for iter.Next() { // no need to sort keys
		var valid bool
		c := &e.buf[end+i]
		c.Parent = current
		c.Name, valid = validMapKey(iter.Key())
		if !valid {
			e.addErrorContext(c, errors.New("invalid map key"))
			i++
			continue
		}
		me.elemEnc(e, end+i, c, iter.Value())
		i++
	}
}

// structEncoder is the encoderFunc of struct.
type structEncoder struct {
	fields structFields
}

func newStructEncoder(t reflect.Type) encoderFunc {
	e := structEncoder{fields: cachedTypeFields(t)}
	return e.encode
}

func (se structEncoder) encode(e *encodeState, current int, p *Value, v reflect.Value) {
	n := len(se.fields.list)
	p.Type = Map
	if n == 0 {
		return
	}
	p.Length = n
	end := len(e.buf)
	p.First = end
	e.Append(n)
	for j := range se.fields.list {
		f := &se.fields.list[j]
		fv := v
		breakNil := false
		for _, i := range f.index {
			if fv.Kind() == reflect.Pointer {
				if fv.IsNil() {
					breakNil = true
					break
				}
				fv = fv.Elem()
			}
			fv = fv.Field(i)
		}
		if breakNil {
			e.buf[end+j] = Value{Type: Nil, Name: f.name, Parent: current}
			continue
		}
		c := &e.buf[end+j]
		c.Name = f.name
		c.Parent = current
		f.encoder(e, end+j, c, fv)
	}
}

type decoderFunc func(d *decodeState, p *Value, v reflect.Value)

var decoders []decoderFunc

func init() {
	decoders = []decoderFunc{
		nil,          //Invalid
		nil,          // Nil
		decodeBool,   // Bool
		decodeInt,    // Int
		decodeUint,   // Uint
		decodeFloat,  // Float
		decodeString, // String
		decodeBytes,  // Bytes
		decodeSlice,  // Slice
		decodeMap,    // Map
	}
}

type decodeState struct {
	*Buffer
	savedError error
}

func (d *decodeState) addErrorContext(p *Value, err error) {

}

// decodeValue decodes the value stored in p [*Value] into v [reflect.Value].
func decodeValue(d *decodeState, p *Value, v reflect.Value) {
	if f := decoders[p.Type]; f != nil {
		f(d, p, makeValue(v))
	}
}

// valueInterface returns p's underlying value as [interface{}].
func valueInterface(d *decodeState, p *Value) interface{} {
	switch p.Type {
	case Nil:
		return nil
	case Bool:
		return p.Bool()
	case Int:
		return p.Int()
	case Uint:
		return p.Uint()
	case Float:
		return p.Float()
	case String:
		return p.String()
	case Bytes:
		return p.Bytes()
	case Slice:
		arr := d.buf[p.First : p.First+p.Length]
		return arrayInterface(d, arr)
	case Map:
		arr := d.buf[p.First : p.First+p.Length]
		return objectInterface(d, arr)
	}
	return nil
}

// arrayInterface returns p's underlying value as [[]interface{}].
func arrayInterface(d *decodeState, v []Value) []interface{} {
	r := make([]interface{}, len(v))
	for i, p := range v {
		r[i] = valueInterface(d, &p)
	}
	return r
}

// objectInterface returns p's underlying value as [map[string]interface{}].
func objectInterface(d *decodeState, v []Value) map[string]interface{} {
	r := make(map[string]interface{}, len(v))
	for _, p := range v {
		r[p.Name] = valueInterface(d, &p)
	}
	return r
}

// decodeBool decodes a bool value into v.
func decodeBool(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Bool()))
	case reflect.Bool:
		v.SetBool(p.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := int64(0)
		if p.Bool() {
			i = 1
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := uint64(0)
		if p.Bool() {
			u = 1
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f := float64(0)
		if p.Bool() {
			f = 1
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(strconv.FormatBool(p.Bool()))
	}
}

// decodeInt decodes a int64 value into v.
func decodeInt(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Int()))
	case reflect.Bool:
		v.SetBool(p.Int() == 1)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(p.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(p.Int()))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(p.Int()))
	case reflect.String:
		v.SetString(strconv.FormatInt(p.Int(), 64))
	}
}

// decodeUint decodes a uint64 value into v.
func decodeUint(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Uint()))
	case reflect.Bool:
		v.SetBool(p.Uint() == 1)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(p.Uint()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(p.Uint())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(p.Uint()))
	case reflect.String:
		v.SetString(strconv.FormatUint(p.Uint(), 64))
	}
}

// decodeFloat decodes a float64 value into v.
func decodeFloat(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Float()))
	case reflect.Bool:
		v.SetBool(p.Float() == 1)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(p.Float()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(p.Float()))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(p.Float())
	case reflect.String:
		v.SetString(strconv.FormatFloat(p.Float(), 'f', -1, 64))
	}
}

// decodeString decodes a string value into v.
func decodeString(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.String()))
	case reflect.Bool:
		b, err := strconv.ParseBool(p.String())
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(p.String(), 10, 64)
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(p.String(), 10, 64)
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(p.String(), 64)
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(p.String())
	}
}

// decodeBytes decodes []byte value
func decodeBytes(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Bytes()))
	case reflect.Bool:
		b, err := strconv.ParseBool(string(p.Bytes()))
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(string(p.Bytes()), 10, 64)
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(string(p.Bytes()), 10, 64)
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(string(p.Bytes()), 64)
		if err != nil {
			d.addErrorContext(p, err)
			break
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(string(p.Bytes()))
	}
}

// decodeSlice decodes slice value
func decodeSlice(d *decodeState, p *Value, v reflect.Value) {
	var arr []Value
	if p.Length > 0 {
		arr = d.buf[p.First : p.First+p.Length]
	}
	switch v.Kind() {
	case reflect.Interface:
		arr := arrayInterface(d, arr)
		v.Set(reflect.ValueOf(arr))
	default:
		n := len(arr)
		if v.Kind() == reflect.Slice {
			v.Set(reflect.MakeSlice(v.Type(), n, n))
		}
		i := 0
		for ; i < n; i++ {
			if i < v.Len() {
				decodeValue(d, &d.buf[p.First+i], v.Index(i))
			}
		}
		if i < v.Len() {
			if v.Kind() == reflect.Array {
				z := reflect.Zero(v.Type().Elem())
				for ; i < v.Len(); i++ {
					v.Index(i).Set(z)
				}
			} else {
				v.SetLen(i)
			}
		}
	}
}

// decodeMap decodes map value
func decodeMap(d *decodeState, p *Value, v reflect.Value) {
	t := v.Type()
	switch v.Kind() {
	case reflect.Interface:
		var arr []Value
		if p.Length > 0 {
			arr = d.buf[p.First : p.First+p.Length]
		}
		oi := objectInterface(d, arr)
		v.Set(reflect.ValueOf(oi))
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			return
		}
		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}
		et := t.Elem()
		for i := 0; i < p.Length; i++ {
			ev := reflect.New(et).Elem()
			decodeValue(d, &d.buf[p.First+i], ev)
			v.SetMapIndex(reflect.ValueOf(p.Name), ev)
		}
	case reflect.Struct:
		decodeStruct(d, p, v, t)
	}
}

func decodeStruct(d *decodeState, p *Value, v reflect.Value, t reflect.Type) {
	fields := cachedTypeFields(t)
	for i := 0; i < p.Length; i++ {
		e := &d.buf[p.First+i]
		f, ok := fields.byExactName[e.Name]
		if !ok {
			continue
		}
		subValue := v
		for _, j := range f.index {
			if subValue.Kind() == reflect.Ptr {
				if subValue.IsNil() {
					if !subValue.CanSet() {
						subValue = reflect.Value{}
						break
					}
					subValue.Set(reflect.New(subValue.Type().Elem()))
				}
				subValue = subValue.Elem()
			}
			subValue = subValue.Field(j)
		}
		decodeValue(d, e, subValue)
	}
}

func makeValue(v reflect.Value) reflect.Value {
	for {
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && e.Elem().Kind() == reflect.Ptr {
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}

		v = v.Elem()
	}
	return v
}

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) string {
	tag, _, _ = strings.Cut(tag, ",")
	return tag
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
			// Backslash and quote chars are reserved, but
			// otherwise any punctuation chars are allowed
			// in a tag name.
		case !unicode.IsLetter(c) && !unicode.IsDigit(c):
			return false
		}
	}
	return true
}

// A field represents a single field found in a struct.
type structField struct {
	name    string
	tag     bool
	index   []int
	typ     reflect.Type
	encoder encoderFunc
}

type structFields struct {
	list        []structField
	byExactName map[string]*structField
}

// byIndex sorts field by index sequence.
type byIndex []structField

func (x byIndex) Len() int { return len(x) }

func (x byIndex) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x byIndex) Less(i, j int) bool {
	for k, xik := range x[i].index {
		if k >= len(x[j].index) {
			return false
		}
		if xik != x[j].index[k] {
			return xik < x[j].index[k]
		}
	}
	return len(x[i].index) < len(x[j].index)
}

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) structFields {
	if f, ok := stFieldCache.Load(t); ok {
		return f.(structFields)
	}
	f, _ := stFieldCache.LoadOrStore(t, typeFields(t))
	return f.(structFields)
}

// typeFields returns a list of fields that JSON should recognize for the given type.
// The algorithm is breadth-first search over the set of structs to include - the top struct
// and then any reachable anonymous structs.
func typeFields(t reflect.Type) structFields {
	// Anonymous fields to explore at the current level and the next.
	var current []structField
	next := []structField{{typ: t}}

	// Count of queued names for current level and the next.
	var count, nextCount map[reflect.Type]int

	// Types already visited at an earlier level.
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []structField

	for len(next) > 0 {
		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {
			if visited[f.typ] {
				continue
			}
			visited[f.typ] = true

			// Scan f.typ for fields to include.
			for i := 0; i < f.typ.NumField(); i++ {
				sf := f.typ.Field(i)
				if sf.Anonymous {
					st := sf.Type
					if st.Kind() == reflect.Pointer {
						st = st.Elem()
					}
					if !sf.IsExported() && t.Kind() != reflect.Struct {
						// Ignore embedded fields of unexported non-struct types.
						continue
					}
					// Do not ignore embedded fields of unexported struct types
					// since they may have exported fields.
				} else if !sf.IsExported() {
					// Ignore unexported non-embedded fields.
					continue
				}
				tag := sf.Tag.Get("json")
				if tag == "-" {
					continue
				}
				name := parseTag(tag)
				if !isValidTag(name) {
					name = ""
				}
				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Name() == "" && ft.Kind() == reflect.Pointer {
					// Follow pointer.
					ft = ft.Elem()
				}

				// Record found field and index sequence.
				if name != "" || !sf.Anonymous || ft.Kind() != reflect.Struct {
					tagged := name != ""
					if name == "" {
						name = sf.Name
					}
					field := structField{
						name:  name,
						tag:   tagged,
						index: index,
						typ:   ft,
					}
					fields = append(fields, field)
					if count[f.typ] > 1 {
						// If there were multiple instances, add a second,
						// so that the annihilation code will see a duplicate.
						// It only cares about the distinction between 1 or 2,
						// so don't bother generating any more copies.
						fields = append(fields, fields[len(fields)-1])
					}
					continue
				}

				// Record new anonymous struct to explore in next round.
				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, structField{name: ft.Name(), index: index, typ: ft})
				}
			}
		}
	}

	sort.Slice(fields, func(i, j int) bool {
		x := fields
		// sort field by name, breaking ties with depth, then
		// breaking ties with "name came from json tag", then
		// breaking ties with index sequence.
		if x[i].name != x[j].name {
			return x[i].name < x[j].name
		}
		if len(x[i].index) != len(x[j].index) {
			return len(x[i].index) < len(x[j].index)
		}
		if x[i].tag != x[j].tag {
			return x[i].tag
		}
		return byIndex(x).Less(i, j)
	})

	// Delete all fields that are hidden by the Go rules for embedded fields,
	// except that fields with JSON tags are promoted.

	// The fields are sorted in primary order of name, secondary order
	// of field index length. Loop over names; for each name, delete
	// hidden fields by choosing the one dominant field that survives.
	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
		// One iteration per name.
		// Find the sequence of fields with the name of this first field.
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
			fj := fields[i+advance]
			if fj.name != name {
				break
			}
		}
		if advance == 1 { // Only one field with this name
			out = append(out, fi)
			continue
		}
		dominant, ok := dominantField(fields[i : i+advance])
		if ok {
			out = append(out, dominant)
		}
	}

	fields = out
	sort.Sort(byIndex(fields))

	for i := range fields {
		f := &fields[i]
		f.encoder = cachedTypeEncoder(typeByIndex(t, f.index))
	}

	exactNameIndex := make(map[string]*structField, len(fields))
	for i, field := range fields {
		exactNameIndex[field.name] = &fields[i]
	}
	return structFields{fields, exactNameIndex}
}

func typeByIndex(t reflect.Type, index []int) reflect.Type {
	for _, i := range index {
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
		t = t.Field(i).Type
	}
	return t
}

// dominantField looks through the fields, all of which are known to
// have the same name, to find the single field that dominates the
// others using Go's embedding rules, modified by the presence of
// JSON tags. If there are multiple top-level fields, the boolean
// will be false: This condition is an error in Go and we skip all
// the fields.
func dominantField(fields []structField) (structField, bool) {
	// The fields are sorted in increasing index-length order, then by presence of tag.
	// That means that the first field is the dominant one. We need only check
	// for error cases: two fields at top level, either both tagged or neither tagged.
	if len(fields) > 1 && len(fields[0].index) == len(fields[1].index) && fields[0].tag == fields[1].tag {
		return structField{}, false
	}
	return fields[0], true
}
