package fastconv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func Decode(l *Buffer, v interface{}) error {

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(v)}
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	d := &decodeState{Buffer: l}
	decodeValue(d, &d.buf[0], rv)
	return nil
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

// dupSlice duplicates a slice.
func dupSlice[T bool | Number | string](v []T) []T {
	r := make([]T, len(v))
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
	default:
		panic(fmt.Errorf("unexpected kind %v", p.Type))
	}
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
			panic(err)
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(p.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(p.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		v.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(p.String(), 64)
		if err != nil {
			panic(err)
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(p.String())
	}
}

// decodeBools decodes []bool value
func decodeBools(d *decodeState, p *Value, v reflect.Value) {

}

// decodeInts decodes []int value
func decodeInts(d *decodeState, p *Value, v reflect.Value) {

}

// decodeInt8s decodes []int8 value
func decodeInt8s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeInt16s decodes []int16 value
func decodeInt16s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeInt32s decodes []int32 value
func decodeInt32s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeInt64s decodes []int64 value
func decodeInt64s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeUints decodes []uint value
func decodeUints(d *decodeState, p *Value, v reflect.Value) {

}

// decodeUint8s decodes []uint8 value
func decodeUint8s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeUint16s decodes []uint16 value
func decodeUint16s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeUint32s decodes []uint32 value
func decodeUint32s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeUint64s decodes []uint64 value
func decodeUint64s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeFloat32s decodes []float32 value
func decodeFloat32s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeFloat64s decodes []float64 value
func decodeFloat64s(d *decodeState, p *Value, v reflect.Value) {

}

// decodeStrings decodes []string value
func decodeStrings(d *decodeState, p *Value, v reflect.Value) {

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
