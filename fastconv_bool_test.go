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
	"testing"
)

func TestEncode_Bool(t *testing.T) {

	encodeSuccess(t, true, GetBuffer().Bool("", true))
	encodeSuccess(t, false, GetBuffer().Bool("", false))

	encodeSuccess(t, Ptr(true), GetBuffer().Bool("", true))
	encodeSuccess(t, Ptr(false), GetBuffer().Bool("", false))
}

func TestDecode_Bool(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Bool("", true), true)
	decodeSuccess[interface{}](t, GetBuffer().Bool("", false), false)

	decodeSuccess(t, GetBuffer().Bool("", true), true)
	decodeSuccess(t, GetBuffer().Bool("", false), false)

	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(true))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(false))

	decodeSuccess(t, GetBuffer().Bool("", false), int(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int8(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int16(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int32(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int64(0))

	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(int(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(int8(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(int16(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(int32(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(int64(0)))

	decodeSuccess(t, GetBuffer().Bool("", true), int(1))
	decodeSuccess(t, GetBuffer().Bool("", true), int8(1))
	decodeSuccess(t, GetBuffer().Bool("", true), int16(1))
	decodeSuccess(t, GetBuffer().Bool("", true), int32(1))
	decodeSuccess(t, GetBuffer().Bool("", true), int64(1))

	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(int(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(int8(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(int16(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(int32(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(int64(1)))

	decodeSuccess(t, GetBuffer().Bool("", false), uint(0))
	decodeSuccess(t, GetBuffer().Bool("", false), uint8(0))
	decodeSuccess(t, GetBuffer().Bool("", false), uint16(0))
	decodeSuccess(t, GetBuffer().Bool("", false), uint32(0))
	decodeSuccess(t, GetBuffer().Bool("", false), uint64(0))

	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(uint(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(uint8(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(uint16(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(uint32(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(uint64(0)))

	decodeSuccess(t, GetBuffer().Bool("", true), uint(1))
	decodeSuccess(t, GetBuffer().Bool("", true), uint8(1))
	decodeSuccess(t, GetBuffer().Bool("", true), uint16(1))
	decodeSuccess(t, GetBuffer().Bool("", true), uint32(1))
	decodeSuccess(t, GetBuffer().Bool("", true), uint64(1))

	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint8(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint16(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint32(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint64(1)))

	decodeSuccess(t, GetBuffer().Bool("", false), float32(0))
	decodeSuccess(t, GetBuffer().Bool("", false), float64(0))

	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(float32(0)))
	decodeSuccess(t, GetBuffer().Bool("", false), Ptr(float64(0)))

	decodeSuccess(t, GetBuffer().Bool("", true), float32(1))
	decodeSuccess(t, GetBuffer().Bool("", true), float64(1))

	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(float32(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(float64(1)))

	decodeSuccess(t, GetBuffer().Bool("", false), "false")
	decodeSuccess(t, GetBuffer().Bool("", true), "true")

	decodeSuccess(t, GetBuffer().Bool("", false), Ptr("false"))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr("true"))
}

//
//func TestConvert_Bool(t *testing.T) {
//
//
//	success[bool, interface{}](t, false, false)
//	success[bool, interface{}](t, true, true)
//
//	fail[bool, uintptr](t, false, errors.New("can't convert bool to uintptr"))
//	fail[bool, uintptr](t, true, errors.New("can't convert bool to uintptr"))
//	fail[bool, complex64](t, false, errors.New("can't convert bool to complex64"))
//	fail[bool, complex64](t, true, errors.New("can't convert bool to complex64"))
//	fail[bool, complex128](t, false, errors.New("can't convert bool to complex128"))
//	fail[bool, complex128](t, true, errors.New("can't convert bool to complex128"))
//	fail[bool, []int](t, false, errors.New("can't convert bool to []int"))
//	fail[bool, []int](t, true, errors.New("can't convert bool to []int"))
//	fail[bool, [3]int](t, false, errors.New("can't convert bool to [3]int"))
//	fail[bool, [3]int](t, true, errors.New("can't convert bool to [3]int"))
//	fail[bool, func()](t, false, errors.New("can't convert bool to func()"))
//	fail[bool, func()](t, true, errors.New("can't convert bool to func()"))
//	fail[bool, chan int](t, false, errors.New("can't convert bool to chan int"))
//	fail[bool, chan int](t, true, errors.New("can't convert bool to chan int"))
//	fail[bool, map[string]bool](t, false, errors.New("can't convert bool to map[string]bool"))
//	fail[bool, map[string]bool](t, true, errors.New("can't convert bool to map[string]bool"))
//	fail[bool, struct{}](t, false, errors.New("can't convert bool to struct {}"))
//	fail[bool, struct{}](t, true, errors.New("can't convert bool to struct {}"))
//
//
//
//	fail[bool, *uintptr](t, false, errors.New("can't convert bool to uintptr"))
//	fail[bool, *uintptr](t, true, errors.New("can't convert bool to uintptr"))
//	fail[bool, *complex64](t, false, errors.New("can't convert bool to complex64"))
//	fail[bool, *complex64](t, true, errors.New("can't convert bool to complex64"))
//	fail[bool, *complex128](t, false, errors.New("can't convert bool to complex128"))
//	fail[bool, *complex128](t, true, errors.New("can't convert bool to complex128"))
//	fail[bool, *[]int](t, false, errors.New("can't convert bool to []int"))
//	fail[bool, *[]int](t, true, errors.New("can't convert bool to []int"))
//	fail[bool, *[3]int](t, false, errors.New("can't convert bool to [3]int"))
//	fail[bool, *[3]int](t, true, errors.New("can't convert bool to [3]int"))
//	fail[bool, *func()](t, false, errors.New("can't convert bool to func()"))
//	fail[bool, *func()](t, true, errors.New("can't convert bool to func()"))
//	fail[bool, *chan int](t, false, errors.New("can't convert bool to chan int"))
//	fail[bool, *chan int](t, true, errors.New("can't convert bool to chan int"))
//	fail[bool, *map[string]bool](t, false, errors.New("can't convert bool to map[string]bool"))
//	fail[bool, *map[string]bool](t, true, errors.New("can't convert bool to map[string]bool"))
//	fail[bool, *struct{}](t, false, errors.New("can't convert bool to struct {}"))
//	fail[bool, *struct{}](t, true, errors.New("can't convert bool to struct {}"))
//}

//func TestConvert_Bools(t *testing.T) {
//
//	success[[]bool, []bool](t, []bool{true, false, true}, []bool{true, false, true})
//	success[[]bool, []int](t, []bool{true, false, true}, []int{1, 0, 1})
//	success[[]bool, []int8](t, []bool{true, false, true}, []int8{1, 0, 1})
//	success[[]bool, []int16](t, []bool{true, false, true}, []int16{1, 0, 1})
//	success[[]bool, []int32](t, []bool{true, false, true}, []int32{1, 0, 1})
//	success[[]bool, []int64](t, []bool{true, false, true}, []int64{1, 0, 1})
//	success[[]bool, []uint](t, []bool{true, false, true}, []uint{1, 0, 1})
//	success[[]bool, []uint8](t, []bool{true, false, true}, []uint8{1, 0, 1})
//	success[[]bool, []uint16](t, []bool{true, false, true}, []uint16{1, 0, 1})
//	success[[]bool, []uint32](t, []bool{true, false, true}, []uint32{1, 0, 1})
//	success[[]bool, []uint64](t, []bool{true, false, true}, []uint64{1, 0, 1})
//	success[[]bool, []float32](t, []bool{true, false, true}, []float32{1, 0, 1})
//	success[[]bool, []float64](t, []bool{true, false, true}, []float64{1, 0, 1})
//	success[[]bool, []string](t, []bool{true, false, true}, []string{"true", "false", "true"})
//	success[[]bool, interface{}](t, []bool{true, false, true}, []bool{true, false, true})
//
//	fail[[]bool, uintptr](t, []bool{true, false, true}, errors.New("can't convert []bool to uintptr"))
//	fail[[]bool, complex64](t, []bool{true, false, true}, errors.New("can't convert []bool to complex64"))
//	fail[[]bool, complex128](t, []bool{true, false, true}, errors.New("can't convert []bool to complex128"))
//	fail[[]bool, func()](t, []bool{true, false, true}, errors.New("can't convert []bool to func()"))
//	fail[[]bool, chan int](t, []bool{true, false, true}, errors.New("can't convert []bool to chan int"))
//	fail[[]bool, map[string]bool](t, []bool{true, false, true}, errors.New("can't convert []bool to map[string]bool"))
//	fail[[]bool, struct{}](t, []bool{true, false, true}, errors.New("can't convert []bool to struct {}"))
//
//	fail[[]bool, [2]bool](t, []bool{true, false, true}, errors.New("array [2]bool overflow"))
//	fail[[]bool, [2]int](t, []bool{true, false, true}, errors.New("array [2]int overflow"))
//	fail[[]bool, [2]int8](t, []bool{true, false, true}, errors.New("array [2]int8 overflow"))
//	fail[[]bool, [2]int16](t, []bool{true, false, true}, errors.New("array [2]int16 overflow"))
//	fail[[]bool, [2]int32](t, []bool{true, false, true}, errors.New("array [2]int32 overflow"))
//	fail[[]bool, [2]int64](t, []bool{true, false, true}, errors.New("array [2]int64 overflow"))
//	fail[[]bool, [2]uint](t, []bool{true, false, true}, errors.New("array [2]uint overflow"))
//	fail[[]bool, [2]uint8](t, []bool{true, false, true}, errors.New("array [2]uint8 overflow"))
//	fail[[]bool, [2]uint16](t, []bool{true, false, true}, errors.New("array [2]uint16 overflow"))
//	fail[[]bool, [2]uint32](t, []bool{true, false, true}, errors.New("array [2]uint32 overflow"))
//	fail[[]bool, [2]uint64](t, []bool{true, false, true}, errors.New("array [2]uint64 overflow"))
//	fail[[]bool, [2]float32](t, []bool{true, false, true}, errors.New("array [2]float32 overflow"))
//	fail[[]bool, [2]float64](t, []bool{true, false, true}, errors.New("array [2]float64 overflow"))
//	fail[[]bool, [2]string](t, []bool{true, false, true}, errors.New("array [2]string overflow"))
//
//	success[[]bool, [3]bool](t, []bool{true, false, true}, [3]bool{true, false, true})
//	success[[]bool, [3]int](t, []bool{true, false, true}, [3]int{1, 0, 1})
//	success[[]bool, [3]int8](t, []bool{true, false, true}, [3]int8{1, 0, 1})
//	success[[]bool, [3]int16](t, []bool{true, false, true}, [3]int16{1, 0, 1})
//	success[[]bool, [3]int32](t, []bool{true, false, true}, [3]int32{1, 0, 1})
//	success[[]bool, [3]int64](t, []bool{true, false, true}, [3]int64{1, 0, 1})
//	success[[]bool, [3]uint](t, []bool{true, false, true}, [3]uint{1, 0, 1})
//	success[[]bool, [3]uint8](t, []bool{true, false, true}, [3]uint8{1, 0, 1})
//	success[[]bool, [3]uint16](t, []bool{true, false, true}, [3]uint16{1, 0, 1})
//	success[[]bool, [3]uint32](t, []bool{true, false, true}, [3]uint32{1, 0, 1})
//	success[[]bool, [3]uint64](t, []bool{true, false, true}, [3]uint64{1, 0, 1})
//	success[[]bool, [3]float32](t, []bool{true, false, true}, [3]float32{1, 0, 1})
//	success[[]bool, [3]float64](t, []bool{true, false, true}, [3]float64{1, 0, 1})
//	success[[]bool, [3]string](t, []bool{true, false, true}, [3]string{"true", "false", "true"})
//
//	fail[[]bool, []uintptr](t, []bool{true, false, true}, errors.New("can't convert []bool to []uintptr"))
//	fail[[]bool, []complex64](t, []bool{true, false, true}, errors.New("can't convert []bool to []complex64"))
//	fail[[]bool, []complex128](t, []bool{true, false, true}, errors.New("can't convert []bool to []complex128"))
//	fail[[]bool, []func()](t, []bool{true, false, true}, errors.New("can't convert []bool to []func()"))
//	fail[[]bool, []chan int](t, []bool{true, false, true}, errors.New("can't convert []bool to []chan int"))
//	fail[[]bool, []map[string]bool](t, []bool{true, false, true}, errors.New("can't convert []bool to []map[string]bool"))
//	fail[[]bool, []struct{}](t, []bool{true, false, true}, errors.New("can't convert []bool to []struct {}"))
//
//	fail[[]bool, [3]uintptr](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]uintptr"))
//	fail[[]bool, [3]complex64](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]complex64"))
//	fail[[]bool, [3]complex128](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]complex128"))
//	fail[[]bool, [3]func()](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]func()"))
//	fail[[]bool, [3]chan int](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]chan int"))
//	fail[[]bool, [3]map[string]bool](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]map[string]bool"))
//	fail[[]bool, [3]struct{}](t, []bool{true, false, true}, errors.New("can't convert []bool to [3]struct {}"))
//}
