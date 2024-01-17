package fastconv

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Encode todo 补充注释
func Encode(l *Buffer, v interface{}) (err error) {

	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.IsNil() {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if ce, ok := r.(*convError); ok {
				err = ce.error
			} else {
				panic(r)
			}
		}
	}()

	e := &encodeState{Buffer: l}
	e.Append(1)
	c := &e.buf[0]
	encodeValue(e, c, rv)
	return
}

type encodeState struct {
	*Buffer
}

// encodeValue encodes v to p [*Value] which stored in l [*Buffer].
func encodeValue(e *encodeState, p *Value, v reflect.Value) {
	if v.IsValid() {
		cachedTypeEncoder(v.Type())(e, p, v)
	}
}

type encoderFunc func(e *encodeState, p *Value, v reflect.Value)

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
	fi, loaded := encoderCache.LoadOrStore(t, encoderFunc(func(e *encodeState, p *Value, v reflect.Value) {
		wg.Wait()
		f(e, p, v)
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
		panic(&convError{fmt.Errorf("unsupported type %s", t)})
	}
}

// boolEncoder is the encoderFunc of bool.
func boolEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Bool
	p.SetBool(v.Bool())
}

// intEncoder is the encoderFunc of int, int8, int16, int32 and int64.
func intEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Int
	p.SetInt(v.Int())
}

// uintEncoder is the encoderFunc of uint, uint8, uint16, uint32 and uint64.
func uintEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Uint
	p.SetUint(v.Uint())
}

// floatEncoder is the encoderFunc of float32 and float64.
func floatEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Float
	p.SetFloat(v.Float())
}

// stringEncoder is the encoderFunc of string.
func stringEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = String
	old := p.Type
	p.SetString(v.String())
	if old != p.Type { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// interfaceEncoder is the encoderFunc of interface{}.
func interfaceEncoder(e *encodeState, p *Value, v reflect.Value) {
	if v.IsNil() {
		p.Type = Nil
		return
	}
	encodeValue(e, p, v.Elem())
}

// sliceInterfaceEncoder is the encoderFunc of []interface{}.
func sliceInterfaceEncoder(e *encodeState, p *Value, v reflect.Value) {
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
			e.buf[end+i] = Value{Type: Nil}
		} else {
			c := &e.buf[end+i]
			encodeValue(e, c, reflect.ValueOf(sValue))
		}
	}
}

// mapStringInterfaceEncoder is the encoderFunc of map[string]interface{}.
func mapStringInterfaceEncoder(e *encodeState, p *Value, v reflect.Value) {
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
			e.buf[end+i] = Value{Type: Nil, Name: mKey}
		} else {
			c := &e.buf[end+i]
			c.Name = mKey
			encodeValue(e, c, reflect.ValueOf(mValue))
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

func (pe ptrEncoder) encode(e *encodeState, p *Value, v reflect.Value) {
	if v.IsNil() {
		p.Type = Nil
		return
	}
	pe.elemEnc(e, p, v.Elem())
}

// arrayEncoder is the encoderFunc of array and slice.
type arrayEncoder struct {
	elemEnc encoderFunc
}

var fastSliceEncoders []encoderFunc

func init() {
	fastSliceEncoders = []encoderFunc{
		nil,             // Invalid
		boolsEncoder,    // Bool
		intsEncoder,     // Int
		int8sEncoder,    // Int8
		int16sEncoder,   // Int16
		int32sEncoder,   // Int32
		int64sEncoder,   // Int64
		uintsEncoder,    // Uint
		uint8sEncoder,   // Uint8
		uint16sEncoder,  // Uint16
		uint32sEncoder,  // Uint32
		uint64sEncoder,  // Uint64
		nil,             // Uintptr
		float32sEncoder, // Float32
		float64sEncoder, // Float64
		nil,             // Complex64
		nil,             // Complex128
		nil,             // Array
		nil,             // Chan
		nil,             // Func
		nil,             // Interface
		nil,             // Map
		nil,             // Pointer
		nil,             // Slice
		stringsEncoder,  // String
		nil,             // Struct
		nil,             // UnsafePointer
	}
}

// boolsEncoder is the encoderFunc of type []bool.
func boolsEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Bools
	old := p.Type
	p.SetBools(v.Interface().([]bool))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// boolsEncoder is the encoderFunc of type []int.
func intsEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Ints
	old := p.Type
	p.SetInts(v.Interface().([]int))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// int8sEncoder is the encoderFunc of type []int8.
func int8sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Int8s
	old := p.Type
	p.SetInt8s(v.Interface().([]int8))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// int16sEncoder is the encoderFunc of type []int16.
func int16sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Int16s
	old := p.Type
	p.SetInt16s(v.Interface().([]int16))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// int32sEncoder is the encoderFunc of type []int32.
func int32sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Int32s
	old := p.Type
	p.SetInt32s(v.Interface().([]int32))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// int64sEncoder is the encoderFunc of type []int64.
func int64sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Int64s
	old := p.Type
	p.SetInt64s(v.Interface().([]int64))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// uintsEncoder is the encoderFunc of type []uint.
func uintsEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Uints
	old := p.Type
	p.SetUints(v.Interface().([]uint))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// uint8sEncoder is the encoderFunc of type []uint8.
func uint8sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Uint8s
	old := p.Type
	p.SetUint8s(v.Interface().([]uint8))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// uint16sEncoder is the encoderFunc of type []uint16.
func uint16sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Uint16s
	old := p.Type
	p.SetUint16s(v.Interface().([]uint16))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// uint32sEncoder is the encoderFunc of type []uint32.
func uint32sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Uint32s
	old := p.Type
	p.SetUint32s(v.Interface().([]uint32))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// uint64sEncoder is the encoderFunc of type []uin64.
func uint64sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Uint64s
	old := p.Type
	p.SetUint64s(v.Interface().([]uint64))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// float32sEncoder is the encoderFunc of type []float32.
func float32sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Float32s
	old := p.Type
	p.SetFloat32s(v.Interface().([]float32))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// float64sEncoder is the encoderFunc of type []float64.
func float64sEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Float64s
	old := p.Type
	p.SetFloat64s(v.Interface().([]float64))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

// stringsEncoder is the encoderFunc of type []string.
func stringsEncoder(e *encodeState, p *Value, v reflect.Value) {
	p.Type = Strings
	old := p.Type
	p.SetStrings(v.Interface().([]string))
	if p.Type != old { // should never happen
		panic(errors.New("!!! Type was unexpectedly modified"))
	}
}

func newArrayEncoder(t reflect.Type) encoderFunc {
	et := t.Elem()
	if f := fastSliceEncoders[et.Kind()]; f != nil {
		return f
	}
	e := arrayEncoder{cachedTypeEncoder(et)}
	return e.encode
}

func (ae arrayEncoder) encode(e *encodeState, p *Value, v reflect.Value) {
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
		ae.elemEnc(e, c, v.Index(i))
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

func (me mapEncoder) encode(e *encodeState, p *Value, v reflect.Value) {
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
		c.Name, valid = validMapKey(iter.Key())
		if !valid {
			panic(&convError{fmt.Errorf("invalid map key type %s", iter.Key().Type())})
		}
		me.elemEnc(e, c, iter.Value())
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

func (se structEncoder) encode(e *encodeState, p *Value, v reflect.Value) {
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
			e.buf[end+j] = Value{Type: Nil, Name: f.name}
			continue
		}
		c := &e.buf[end+j]
		c.Name = f.name
		f.encoder(e, c, fv)
	}
}

// A structField represents a single field found in a struct.
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
					if !sf.IsExported() && st.Kind() != reflect.Struct {
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
				name, _, _ := strings.Cut(tag, ",")
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
