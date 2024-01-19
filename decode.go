package fastconv

import (
	"fmt"
	"reflect"
	"strconv"
)

// Decode todo 补充注释
func Decode(l *Buffer, v interface{}) (err error) {

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidDecodeError{reflect.TypeOf(v)}
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

	d := &decodeState{Buffer: l}
	decodeValue(d, &d.buf[0], rv)
	return
}

// An InvalidDecodeError describes an invalid argument passed to Decode.
// (The argument to Decode must be a non-nil pointer.)
type InvalidDecodeError struct {
	Type reflect.Type
}

func (e *InvalidDecodeError) Error() string {
	if e.Type == nil {
		return "fastconv: Decode(nil)"
	}
	if e.Type.Kind() != reflect.Pointer {
		return "fastconv: Decode(non-pointer " + e.Type.String() + ")"
	}
	return "fastconv: Decode(nil " + e.Type.String() + ")"
}

type decoderFunc func(d *decodeState, p *Value, v reflect.Value)

var decoders []decoderFunc

func init() {
	decoders = []decoderFunc{
		nil,            // Invalid
		nil,            // Nil
		decodeBool,     // Bool
		decodeInt,      // Int
		decodeUint,     // Uint
		decodeFloat,    // Float
		decodeString,   // String
		decodeBools,    // Bools
		decodeInts,     // Ints
		decodeInt8s,    // Int8s
		decodeInt16s,   // Int16s
		decodeInt32s,   // Int32s
		decodeInt64s,   // Int64s
		decodeUints,    // Uints
		decodeUint8s,   // Uint8s
		decodeUint16s,  // Uint16s
		decodeUint32s,  // Uint32s
		decodeUint64s,  // Uint64s
		decodeFloat32s, // Float32s
		decodeFloat64s, // Float64s
		decodeStrings,  // Strings
		decodeSlice,    // Slice
		decodeMap,      // Map
	}
}

type decodeState struct {
	*Buffer
}

// decodeValue decodes the value stored in p [*Value] into v [reflect.Value].
func decodeValue(d *decodeState, p *Value, v reflect.Value) {
	if f := decoders[p.Type]; f != nil {
		f(d, p, makeValue(v))
	}
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// dupSlice duplicates a slice of type bool, Number, string.
func dupSlice[T bool | Number | string](v []T) []T {
	r := make([]T, len(v), len(v))
	copy(r, v)
	return r
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
	case Bools:
		return dupSlice(p.Bools())
	case Ints:
		return dupSlice(p.Ints())
	case Int8s:
		return dupSlice(p.Int8s())
	case Int16s:
		return dupSlice(p.Int16s())
	case Int32s:
		return dupSlice(p.Int32s())
	case Int64s:
		return dupSlice(p.Int64s())
	case Uints:
		return dupSlice(p.Uints())
	case Uint8s:
		return dupSlice(p.Uint8s())
	case Uint16s:
		return dupSlice(p.Uint16s())
	case Uint32s:
		return dupSlice(p.Uint32s())
	case Uint64s:
		return dupSlice(p.Uint64s())
	case Float32s:
		return dupSlice(p.Float32s())
	case Float64s:
		return dupSlice(p.Float64s())
	case Strings:
		return dupSlice(p.Strings())
	case Slice:
		arr := d.buf[p.First : p.First+p.Length]
		return arrayInterface(d, arr)
	case Map:
		arr := d.buf[p.First : p.First+p.Length]
		return objectInterface(d, arr)
	default: // should never happen
		panic(fmt.Errorf("unexpected kind %v", p.Type))
	}
}

// arrayInterface returns p's underlying value as [[]interface{}].
func arrayInterface(d *decodeState, v []Value) []interface{} {
	r := make([]interface{}, len(v), len(v))
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
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
	}
}

// decodeInt decodes a int64 value into v.
func decodeInt(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Int()))
	case reflect.Bool:
		v.SetBool(p.Int() != 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(p.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(p.Int()))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(p.Int()))
	case reflect.String:
		v.SetString(strconv.FormatInt(p.Int(), 10))
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
	}
}

// decodeUint decodes a uint64 value into v.
func decodeUint(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Uint()))
	case reflect.Bool:
		v.SetBool(p.Uint() != 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(p.Uint()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(p.Uint())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(p.Uint()))
	case reflect.String:
		v.SetString(strconv.FormatUint(p.Uint(), 10))
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
	}
}

// decodeFloat decodes a float64 value into v.
func decodeFloat(d *decodeState, p *Value, v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(p.Float()))
	case reflect.Bool:
		v.SetBool(p.Float() != 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(p.Float()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(p.Float()))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(p.Float())
	case reflect.String:
		v.SetString(strconv.FormatFloat(p.Float(), 'f', -1, 64))
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
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
			panic(&convError{err})
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(p.String(), 10, 64)
		if err != nil {
			panic(&convError{err})
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(p.String(), 10, 64)
		if err != nil {
			panic(&convError{err})
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(p.String(), 64)
		if err != nil {
			panic(&convError{err})
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(p.String())
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
	}
}

func checkArrayLength(srcLen int, v reflect.Value) int {
	var n int
	switch v.Kind() {
	case reflect.Slice:
		n = srcLen
	case reflect.Array:
		n = v.Len()
		if srcLen > n {
			panic(&convError{fmt.Errorf("array overflow %s", v.Type())})
		}
	default:
		panic(&convError{fmt.Errorf("can't convert to type %s", v.Type())})
	}
	return n
}

// decodeBoolsToNumbers n >= len(s)
func decodeBoolsToNumbers[R Number](s []bool, n int) []R {
	r := make([]R, n, n)
	for i := 0; i < len(s); i++ {
		var v R
		if s[i] {
			v = 1
		}
		r[i] = v
	}
	return r
}

// decodeBools decodes []bool value
func decodeBools(d *decodeState, p *Value, v reflect.Value) {
	n := checkArrayLength(p.Length, v)
	et := v.Type().Elem()
	ek := et.Kind()
	if reflect.Bool == ek {
		r := make([]bool, n, n)
		copy(r, p.Bools())
		v.Set(reflect.ValueOf(r))
		return
	}
	switch ek {
	case reflect.Int:
		r := decodeBoolsToNumbers[int](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int8:
		r := decodeBoolsToNumbers[int8](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int16:
		r := decodeBoolsToNumbers[int16](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int32:
		r := decodeBoolsToNumbers[int32](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int64:
		r := decodeBoolsToNumbers[int64](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint:
		r := decodeBoolsToNumbers[uint](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint8:
		r := decodeBoolsToNumbers[uint8](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint16:
		r := decodeBoolsToNumbers[uint16](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint32:
		r := decodeBoolsToNumbers[uint32](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint64:
		r := decodeBoolsToNumbers[uint64](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Float32:
		r := decodeBoolsToNumbers[float32](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Float64:
		r := decodeBoolsToNumbers[float64](p.Bools(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.String:
		r := make([]string, n, n)
		for i, b := range p.Bools() {
			r[i] = strconv.FormatBool(b)
		}
		v.Set(reflect.ValueOf(r))
	default:
		panic(&convError{fmt.Errorf("can't convert to type %s", v.Type())})
	}
}

// dupNumberSlice n >= len(s)
func dupNumberSlice[T, R Number](s []T, n int) []R {
	r := make([]R, n, n)
	for i := 0; i < len(s); i++ {
		r[i] = R(s[i])
	}
	return r
}

func decodeNumbers[T Number](k reflect.Kind, s []T, v reflect.Value) {
	n := checkArrayLength(len(s), v)
	et := v.Type().Elem()
	ek := et.Kind()
	if k == ek {
		r := make([]T, n, n)
		copy(r, s)
		v.Set(reflect.ValueOf(r))
		return
	}
	switch ek {
	case reflect.Bool:
		r := make([]bool, n, n)
		for i := 0; i < len(s); i++ {
			r[i] = s[i] != 0
		}
		v.Set(reflect.ValueOf(r))
	case reflect.Int:
		r := dupNumberSlice[T, int](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int8:
		r := dupNumberSlice[T, int8](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int16:
		r := dupNumberSlice[T, int16](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int32:
		r := dupNumberSlice[T, int32](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int64:
		r := dupNumberSlice[T, int64](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint:
		r := dupNumberSlice[T, uint](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint8:
		r := dupNumberSlice[T, uint8](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint16:
		r := dupNumberSlice[T, uint16](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint32:
		r := dupNumberSlice[T, uint32](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint64:
		r := dupNumberSlice[T, uint64](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Float32:
		r := dupNumberSlice[T, float32](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.Float64:
		r := dupNumberSlice[T, float64](s, n)
		v.Set(reflect.ValueOf(r))
	case reflect.String:
		r := make([]string, n, n)
		switch k {
		case reflect.Float32, reflect.Float64:
			for i := 0; i < len(s); i++ {
				r[i] = strconv.FormatFloat(float64(s[i]), 'f', -1, 64)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			for i := 0; i < len(s); i++ {
				r[i] = strconv.FormatInt(int64(s[i]), 10)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			for i := 0; i < len(s); i++ {
				r[i] = strconv.FormatUint(uint64(s[i]), 10)
			}
		}
		v.Set(reflect.ValueOf(r))
	default:
		panic(&convError{fmt.Errorf("can't convert to type %s", v.Type())})
	}
}

// decodeInts decodes []int value
func decodeInts(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Int, p.Ints(), v)
}

// decodeInt8s decodes []int8 value
func decodeInt8s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Int8, p.Int8s(), v)
}

// decodeInt16s decodes []int16 value
func decodeInt16s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Int16, p.Int16s(), v)
}

// decodeInt32s decodes []int32 value
func decodeInt32s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Int32, p.Int32s(), v)
}

// decodeInt64s decodes []int64 value
func decodeInt64s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Int64, p.Int64s(), v)
}

// decodeUints decodes []uint value
func decodeUints(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Uint, p.Uints(), v)
}

// decodeUint8s decodes []uint8 value
func decodeUint8s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Uint8, p.Uint8s(), v)
}

// decodeUint16s decodes []uint16 value
func decodeUint16s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Uint16, p.Uint16s(), v)
}

// decodeUint32s decodes []uint32 value
func decodeUint32s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Uint32, p.Uint32s(), v)
}

// decodeUint64s decodes []uint64 value
func decodeUint64s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Uint64, p.Uint64s(), v)
}

// decodeFloat32s decodes []float32 value
func decodeFloat32s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Float32, p.Float32s(), v)
}

// decodeFloat64s decodes []float64 value
func decodeFloat64s(d *decodeState, p *Value, v reflect.Value) {
	decodeNumbers(reflect.Float64, p.Float64s(), v)
}

// decodeStringsToInts n >= len(s)
func decodeStringsToInts[R Number](s []string, n int) []R {
	r := make([]R, n, n)
	for i := 0; i < len(s); i++ {
		v, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			panic(&convError{err})
		}
		r[i] = R(v)
	}
	return r
}

// decodeStringsToUints n >= len(s)
func decodeStringsToUints[R Number](s []string, n int) []R {
	r := make([]R, n, n)
	for i := 0; i < len(s); i++ {
		v, err := strconv.ParseUint(s[i], 10, 64)
		if err != nil {
			panic(&convError{err})
		}
		r[i] = R(v)
	}
	return r
}

// decodeStringsToFloats n >= len(s)
func decodeStringsToFloats[R Number](s []string, n int) []R {
	r := make([]R, n, n)
	for i := 0; i < len(s); i++ {
		v, err := strconv.ParseFloat(s[i], 64)
		if err != nil {
			panic(&convError{err})
		}
		r[i] = R(v)
	}
	return r
}

// decodeStrings decodes []string value
func decodeStrings(d *decodeState, p *Value, v reflect.Value) {
	n := checkArrayLength(p.Length, v)
	et := v.Type().Elem()
	ek := et.Kind()
	if reflect.String == ek {
		r := make([]string, n, n)
		copy(r, p.Strings())
		v.Set(reflect.ValueOf(r))
		return
	}
	switch ek {
	case reflect.Bool:
		r := make([]bool, n, n)
		for i, s := range p.Strings() {
			b, err := strconv.ParseBool(s)
			if err != nil {
				panic(&convError{err})
			}
			r[i] = b
		}
		v.Set(reflect.ValueOf(r))
	case reflect.Int:
		r := decodeStringsToInts[int](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int8:
		r := decodeStringsToInts[int8](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int16:
		r := decodeStringsToInts[int16](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int32:
		r := decodeStringsToInts[int32](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Int64:
		r := decodeStringsToInts[int64](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint:
		r := decodeStringsToUints[uint](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint8:
		r := decodeStringsToUints[uint8](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint16:
		r := decodeStringsToUints[uint16](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint32:
		r := decodeStringsToUints[uint32](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Uint64:
		r := decodeStringsToUints[uint64](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Float32:
		r := decodeStringsToFloats[float32](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	case reflect.Float64:
		r := decodeStringsToFloats[float64](p.Strings(), n)
		v.Set(reflect.ValueOf(r))
	default:
		panic(&convError{fmt.Errorf("can't convert to type %s", v.Type())})
	}
}

// decodeSlice decodes slice value
func decodeSlice(d *decodeState, p *Value, v reflect.Value) {
	var arr []Value
	if p.Length > 0 {
		arr = d.buf[p.First : p.First+p.Length]
	}
	n := len(arr)
	switch v.Kind() {
	case reflect.Interface:
		r := arrayInterface(d, arr)
		v.Set(reflect.ValueOf(r))
	case reflect.Array:
		m := v.Len()
		if n > m {
			panic(&convError{fmt.Errorf("array overflow %s", v.Type())})
		}
		i := 0
		for ; i < n; i++ {
			decodeValue(d, &d.buf[p.First+i], v.Index(i))
		}
		if i < m {
			z := reflect.Zero(v.Type().Elem())
			for ; i < m; i++ {
				v.Index(i).Set(z)
			}
		}
	case reflect.Slice:
		v.Set(reflect.MakeSlice(v.Type(), n, n))
		for i := 0; i < n; i++ {
			decodeValue(d, &d.buf[p.First+i], v.Index(i))
		}
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
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
		r := objectInterface(d, arr)
		v.Set(reflect.ValueOf(r))
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			panic(&convError{fmt.Errorf("unsupported map key type %s", t.Key())})
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
	default:
		panic(&convError{fmt.Errorf("unsupported type %s", v.Type())})
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
						err := fmt.Errorf("fastconv: cannot set embedded pointer to unexported struct: %v", subValue.Type().Elem())
						panic(&convError{err})
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

// makeValue walks down v allocating pointers as needed, until it gets to a non-pointer.
func makeValue(v reflect.Value) reflect.Value {
	for {
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Pointer && !e.IsNil() {
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Pointer {
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
