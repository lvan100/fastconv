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
	"log"
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
	l.Append(1)
	encodeValue(l, &l.buf[0], srcValue)
	decodeValue(l, &l.buf[0], destValue)
	return nil
}

type Kind int

const (
	Nil = Kind(iota)
	Bool
	Int
	Uint
	Float
	String
	Bytes
	Slice
	Map
)

type Value struct {
	Type   Kind
	Name   string
	Value  [8]byte
	Length int // 孩子数量
	First  int // 首个孩子
}

func (p *Value) Bool() bool {
	return *(*bool)(unsafe.Pointer(&p.Value))
}

func (p *Value) SetBool(b bool) {
	*(*bool)(unsafe.Pointer(&p.Value)) = b
}

func (p *Value) Int() int64 {
	return *(*int64)(unsafe.Pointer(&p.Value))
}

func (p *Value) SetInt(i int64) {
	*(*int64)(unsafe.Pointer(&p.Value)) = i
}

func (p *Value) Uint() uint64 {
	return *(*uint64)(unsafe.Pointer(&p.Value))
}

func (p *Value) SetUint(u uint64) {
	*(*uint64)(unsafe.Pointer(&p.Value)) = u
}

func (p *Value) Float() float64 {
	return *(*float64)(unsafe.Pointer(&p.Value))
}

func (p *Value) SetFloat(f float64) {
	*(*float64)(unsafe.Pointer(&p.Value)) = f
}

func (p *Value) String() string {
	return *(*string)(unsafe.Pointer(&p.Value))
}

func (p *Value) SetString(s string) {
	*(*string)(unsafe.Pointer(&p.Value)) = s
}

func (p *Value) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&p.Value))
}

func (p *Value) SetBytes(s []byte) {
	*(*[]byte)(unsafe.Pointer(&p.Value)) = s
}

type Buffer struct {
	buf []Value
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}

func (b *Buffer) Append(n int) {
	b.grow(n)
	for i := 0; i < n; i++ {
		b.buf = append(b.buf, Value{})
	}
}

func (b *Buffer) grow(n int) {
	c := cap(b.buf)
	l := len(b.buf)
	m := l + n
	if m > c {
		buf := make([]Value, l, ((m*3/2)/64)*64)
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

func encodeValue(l *Buffer, p *Value, v reflect.Value) {
	if v.IsValid() {
		typeEncoder(v.Type())(l, p, v)
	}
}

// TypeFor returns the [Type] that represents the type argument T.
func TypeFor[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

var typeSliceByte = TypeFor[[]interface{}]()
var typeSliceInterface = TypeFor[[]interface{}]()
var typeMapStringInterface = TypeFor[map[string]interface{}]()

var encoderCache sync.Map // map[string]string

type encoderFunc func(l *Buffer, p *Value, v reflect.Value)

var fastEncoders []encoderFunc

func init() {
	fastEncoders = []encoderFunc{
		nil,              //Invalid
		valueEncoder,     //Bool
		valueEncoder,     //Int
		valueEncoder,     //Int8
		valueEncoder,     //Int16
		valueEncoder,     //Int32
		valueEncoder,     //Int64
		valueEncoder,     //Uint
		valueEncoder,     //Uint8
		valueEncoder,     //Uint16
		valueEncoder,     //Uint32
		valueEncoder,     //Uint64
		nil,              //Uintptr
		valueEncoder,     //Float32
		valueEncoder,     //Float64
		nil,              //Complex64
		nil,              //Complex128
		nil,              //Array
		nil,              //Chan
		nil,              //Func
		interfaceEncoder, //Interface
		nil,              //Map
		nil,              //Pointer
		nil,              //Slice
		valueEncoder,     //String
		nil,              //Struct
		nil,              //UnsafePointer
	}
}

func typeEncoder(t reflect.Type) encoderFunc {

	switch k := t.Kind(); k {
	case reflect.Slice:
		if t == typeSliceInterface {
			return sliceInterfaceEncoder
		} else if t == typeSliceByte {
			return sliceByteEncoder
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
	fi, loaded := encoderCache.LoadOrStore(t, encoderFunc(func(l *Buffer, p *Value, v reflect.Value) {
		wg.Wait()
		f(l, p, v)
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
		return func(l *Buffer, p *Value, v reflect.Value) {}
	}
}

func sliceByteEncoder(l *Buffer, p *Value, v reflect.Value) {
	p.Type = Bytes
	p.SetBytes(v.Interface().([]byte))
}

func sliceInterfaceEncoder(l *Buffer, p *Value, v reflect.Value) {
	s := v.Interface().([]interface{})
	n := len(s)
	p.Type = Slice
	if n == 0 {
		return
	}
	p.Length = n
	end := len(l.buf)
	p.First = end
	l.Append(n)
	for i, sValue := range s {
		if sValue == nil {
			l.buf[end+i] = Value{Type: Nil}
		} else {
			encodeValue(l, &l.buf[end+i], reflect.ValueOf(sValue))
		}
	}
}

func mapStringInterfaceEncoder(l *Buffer, p *Value, v reflect.Value) {
	m := v.Interface().(map[string]interface{})
	n := len(m)
	p.Type = Map
	if n == 0 {
		return
	}
	p.Length = n
	end := len(l.buf)
	p.First = end
	l.Append(n)
	i := 0
	for mKey, mValue := range m {
		if mValue == nil {
			l.buf[end+i] = Value{Type: Nil}
		} else {
			encodeValue(l, &l.buf[end+i], reflect.ValueOf(mValue))
		}
		l.buf[end+i].Name = mKey
		i++
	}
}

func validMapKey(key reflect.Value) (string, bool) {
	if key.Kind() != reflect.String {
		return "", false
	}
	return key.String(), true
}

func interfaceEncoder(l *Buffer, p *Value, v reflect.Value) {
	if v.IsNil() {
		p.Type = Nil
		return
	}
	encodeValue(l, p, v.Elem())
}

func valueEncoder(l *Buffer, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		p.Type = Bool
		p.SetBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p.Type = Int
		p.SetInt(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		p.Type = Uint
		p.SetUint(v.Uint())
	case reflect.Float32, reflect.Float64:
		p.Type = Float
		p.SetFloat(v.Float())
	case reflect.String:
		p.Type = String
		p.SetString(v.String())
	}
}

type ptrEncoder struct {
	elemEnc encoderFunc
}

func newPtrEncoder(t reflect.Type) encoderFunc {
	e := ptrEncoder{typeEncoder(t.Elem())}
	return e.encode
}

func (e ptrEncoder) encode(l *Buffer, p *Value, v reflect.Value) {
	if v.IsNil() {
		p.Type = Nil
		return
	}
	e.elemEnc(l, p, v.Elem())
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func newArrayEncoder(t reflect.Type) encoderFunc {
	e := arrayEncoder{typeEncoder(t.Elem())}
	return e.encode
}

func (e arrayEncoder) encode(l *Buffer, p *Value, v reflect.Value) {
	n := v.Len()
	p.Type = Slice
	if n == 0 {
		return
	}
	p.Length = n
	end := len(l.buf)
	p.First = end
	l.Append(n)
	for i := 0; i < n; i++ {
		e.elemEnc(l, &l.buf[end+i], v.Index(i))
	}
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func newMapEncoder(t reflect.Type) encoderFunc {
	e := mapEncoder{typeEncoder(t.Elem())}
	return e.encode
}

func (e mapEncoder) encode(l *Buffer, p *Value, v reflect.Value) {
	n := v.Len()
	p.Type = Map
	if n == 0 {
		return
	}
	p.Length = n
	end := len(l.buf)
	p.First = end
	l.Append(n)
	i := 0
	iter := v.MapRange()
	for iter.Next() {
		strKey, valid := validMapKey(iter.Key())
		if !valid {
			continue
		}
		e.elemEnc(l, &l.buf[end+i], iter.Value())
		l.buf[end+i].Name = strKey
		i++
	}
}

type structEncoder struct {
	fields structFields
}

func newStructEncoder(t reflect.Type) encoderFunc {
	e := structEncoder{fields: cachedTypeFields(t)}
	return e.encode
}

func (e structEncoder) encode(l *Buffer, p *Value, v reflect.Value) {
	n := len(e.fields.list)
	p.Type = Map
	if n == 0 {
		return
	}
	p.Length = n
	end := len(l.buf)
	p.First = end
	l.Append(n)
	for j := range e.fields.list {
		f := &e.fields.list[j]
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
			l.buf[end+j] = Value{Type: Nil, Name: f.name}
			continue
		}
		f.encoder(l, &l.buf[end+j], fv)
		l.buf[end+j].Name = f.name
	}
}

func decodeValue(l *Buffer, p *Value, v reflect.Value) {
	if p.Type == Nil {
		return
	}
	v = makeValue(v)
	switch p.Type {
	case Bool:
		decodeBool(p.Bool(), v)
	case Int:
		decodeInt(p.Int(), v)
	case Uint:
		decodeUint(p.Uint(), v)
	case Float:
		decodeFloat(p.Float(), v)
	case String:
		decodeString(p.String(), v)
	case Bytes:
		decodeBytes(p.Bytes(), v)
	case Slice:
		decodeSlice(l, p, v)
	case Map:
		decodeMap(l, p, v)
	default:
		log.Println("should never reach here")
	}
}

func valueInterface(l *Buffer, p Value) interface{} {
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
		arr := l.buf[p.First : p.First+p.Length]
		return arrayInterface(l, arr)
	case Map:
		arr := l.buf[p.First : p.First+p.Length]
		return objectInterface(l, arr)
	default:
		log.Println("should never reach here")
		return nil
	}
}

func arrayInterface(l *Buffer, v []Value) []interface{} {
	r := make([]interface{}, len(v))
	for i, p := range v {
		r[i] = valueInterface(l, p)
	}
	return r
}

func objectInterface(l *Buffer, v []Value) map[string]interface{} {
	r := make(map[string]interface{}, len(v))
	for _, p := range v {
		r[p.Name] = valueInterface(l, p)
	}
	return r
}

// decodeBool decodes bool value to v.
func decodeBool(b bool, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(b))
	case reflect.Bool:
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := int64(0)
		if b {
			i = 1
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := uint64(0)
		if b {
			u = 1
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f := float64(0)
		if b {
			f = 1
		}
		v.SetFloat(f)
	case reflect.String:
		s := strconv.FormatBool(b)
		v.SetString(s)
	default:
		panic(nil)
	}
}

// decodeInt decodes int value to v.
func decodeInt(i int64, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(i))
	case reflect.Bool:
		v.SetBool(i == 1)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(i))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(i))
	case reflect.String:
		s := strconv.FormatInt(i, 64)
		v.SetString(s)
	default:
		panic(nil)
	}
}

// decodeUint decodes uint value to v.
func decodeUint(u uint64, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(u))
	case reflect.Bool:
		v.SetBool(u == 1)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(u))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(u))
	case reflect.String:
		s := strconv.FormatUint(u, 64)
		v.SetString(s)
	default:
		panic(nil)
	}
}

// decodeFloat decodes float value to v.
func decodeFloat(f float64, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(f))
	case reflect.Bool:
		v.SetBool(f == 1)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(f))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(f))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(f)
	case reflect.String:
		s := strconv.FormatFloat(f, 'f', -1, 64)
		v.SetString(s)
	default:
		panic(nil)
	}
}

// decodeString decodes string value to v.
func decodeString(s string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(s))
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			panic(err)
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			panic(err)
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(s)
	default:
		panic(nil)
	}
}

// decodeBytes decodes []byte value to v.
func decodeBytes(b []byte, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(b))
	case reflect.Bool:
		b, err := strconv.ParseBool(string(b))
		if err != nil {
			panic(err)
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			panic(err)
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			panic(err)
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(string(b), 64)
		if err != nil {
			panic(err)
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(string(b))
	default:
		panic(nil)
	}
}

// decodeSlice decodes slice value to v.
func decodeSlice(l *Buffer, p *Value, v reflect.Value) {
	var arr []Value
	if p.Length > 0 {
		arr = l.buf[p.First : p.First+p.Length]
	}
	switch v.Kind() {
	case reflect.Interface:
		arr := arrayInterface(l, arr)
		v.Set(reflect.ValueOf(arr))
	default:
		n := len(arr)
		if v.Kind() == reflect.Slice {
			v.Set(reflect.MakeSlice(v.Type(), n, n))
		}
		i := 0
		for ; i < n; i++ {
			if i < v.Len() {
				decodeValue(l, &l.buf[p.First+i], v.Index(i))
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

// decodeMap decodes map value to v.
func decodeMap(l *Buffer, p *Value, v reflect.Value) {
	t := v.Type()
	switch v.Kind() {
	case reflect.Interface:
		var arr []Value
		if p.Length > 0 {
			arr = l.buf[p.First : p.First+p.Length]
		}
		oi := objectInterface(l, arr)
		v.Set(reflect.ValueOf(oi))
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			return
		}
		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}
		decodeMapToMap(l, p, v, t)
	case reflect.Struct:
		decodeMapToStruct(l, p, v, t)
	}
}

func decodeMapToMap(l *Buffer, p *Value, v reflect.Value, t reflect.Type) {
	elemType := t.Elem()
	for i := 0; i < p.Length; i++ {
		elemValue := reflect.New(elemType).Elem()
		decodeValue(l, &l.buf[p.First+i], elemValue)
		keyValue := reflect.ValueOf(p.Name)
		v.SetMapIndex(keyValue, elemValue)
	}
}

func decodeMapToStruct(l *Buffer, p *Value, v reflect.Value, t reflect.Type) {
	fields := cachedTypeFields(t)
	for i := 0; i < p.Length; i++ {
		e := &l.buf[p.First+i]
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
		decodeValue(l, e, subValue)
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

var fieldCache sync.Map // map[reflect.Type]structFields

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) structFields {
	if f, ok := fieldCache.Load(t); ok {
		return f.(structFields)
	}
	f, _ := fieldCache.LoadOrStore(t, typeFields(t))
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
		f.encoder = typeEncoder(typeByIndex(t, f.index))
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
