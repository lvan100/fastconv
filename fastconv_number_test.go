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
	"testing"
)

func TestEncode_Int(t *testing.T) {

	encodeSuccess(t, int(3), GetBuffer().Int("", 3))
	encodeSuccess(t, int8(3), GetBuffer().Int("", 3))
	encodeSuccess(t, int16(3), GetBuffer().Int("", 3))
	encodeSuccess(t, int32(3), GetBuffer().Int("", 3))
	encodeSuccess(t, int64(3), GetBuffer().Int("", 3))

	encodeSuccess(t, Ptr(int(3)), GetBuffer().Int("", 3))
	encodeSuccess(t, Ptr(int8(3)), GetBuffer().Int("", 3))
	encodeSuccess(t, Ptr(int16(3)), GetBuffer().Int("", 3))
	encodeSuccess(t, Ptr(int32(3)), GetBuffer().Int("", 3))
	encodeSuccess(t, Ptr(int64(3)), GetBuffer().Int("", 3))
}

func TestDecode_Int(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Int("", 3), int64(3))

	decodeSuccess(t, GetBuffer().Int("", 0), false)
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(true))

	decodeSuccess(t, GetBuffer().Int("", 3), int(3))
	decodeSuccess(t, GetBuffer().Int("", 3), int8(3))
	decodeSuccess(t, GetBuffer().Int("", 3), int16(3))
	decodeSuccess(t, GetBuffer().Int("", 3), int32(3))
	decodeSuccess(t, GetBuffer().Int("", 3), int64(3))

	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(int(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(int8(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(int16(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(int32(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(int64(3)))

	decodeSuccess(t, GetBuffer().Int("", 3), uint(3))
	decodeSuccess(t, GetBuffer().Int("", 3), uint8(3))
	decodeSuccess(t, GetBuffer().Int("", 3), uint16(3))
	decodeSuccess(t, GetBuffer().Int("", 3), uint32(3))
	decodeSuccess(t, GetBuffer().Int("", 3), uint64(3))

	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(uint(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(uint8(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(uint16(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(uint32(3)))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(uint64(3)))

	decodeSuccess(t, GetBuffer().Int("", 3), float32(3))
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr(float64(3)))

	decodeSuccess(t, GetBuffer().Int("", 3), "3")
	decodeSuccess(t, GetBuffer().Int("", 3), Ptr("3"))

	decodeFail[uintptr](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to uintptr"))
	decodeFail[complex64](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to complex64"))
	decodeFail[complex128](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to complex128"))
	decodeFail[[]int](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to []int"))
	decodeFail[[3]int](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to [3]int"))
	decodeFail[func()](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to func()"))
	decodeFail[chan int](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to struct {}"))

	decodeFail[*uintptr](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to uintptr"))
	decodeFail[*complex64](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to complex64"))
	decodeFail[*complex128](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to complex128"))
	decodeFail[*[]int](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to []int"))
	decodeFail[*[3]int](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to [3]int"))
	decodeFail[*func()](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to func()"))
	decodeFail[*chan int](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to chan int"))
	decodeFail[*map[string]bool](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to map[string]bool"))
	decodeFail[*struct{}](t, GetBuffer().Int("", 3), errors.New("can't convert int64 to struct {}"))
}

func TestEncode_Ints(t *testing.T) {
	encodeSuccess(t, []int{0, 1, 2}, GetBuffer().Ints("", []int{0, 1, 2}))
	encodeSuccess(t, []int8{0, 1, 2}, GetBuffer().Int8s("", []int8{0, 1, 2}))
	encodeSuccess(t, []int16{0, 1, 2}, GetBuffer().Int16s("", []int16{0, 1, 2}))
	encodeSuccess(t, []int32{0, 1, 2}, GetBuffer().Int32s("", []int32{0, 1, 2}))
	encodeSuccess(t, []int64{0, 1, 2}, GetBuffer().Int64s("", []int64{0, 1, 2}))
}

func TestDecode_Ints(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Ints("", []int{0, 1}), []int{0, 1})
	decodeSuccess[interface{}](t, GetBuffer().Int8s("", []int8{0, 1}), []int8{0, 1})
	decodeSuccess[interface{}](t, GetBuffer().Int16s("", []int16{0, 1}), []int16{0, 1})
	decodeSuccess[interface{}](t, GetBuffer().Int32s("", []int32{0, 1}), []int32{0, 1})
	decodeSuccess[interface{}](t, GetBuffer().Int64s("", []int64{0, 1}), []int64{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1, 2}), []bool{false, true, true})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1, 2}), []bool{false, true, true})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1, 2}), []bool{false, true, true})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1, 2}), []bool{false, true, true})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1, 2}), []bool{false, true, true})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []int{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []int{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []int{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []int{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []int{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []int8{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []int8{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []int8{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []int8{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []int8{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []int16{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []int16{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []int16{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []int16{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []int16{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []int32{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []int32{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []int32{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []int32{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []int32{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []int64{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []int64{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []int64{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []int64{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []int64{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []uint{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []uint{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []uint{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []uint{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []uint{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []uint8{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []uint8{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []uint8{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []uint8{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []uint8{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []uint16{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []uint16{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []uint16{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []uint16{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []uint16{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []uint32{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []uint32{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []uint32{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []uint32{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []uint32{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []uint64{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []uint64{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []uint64{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []uint64{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []uint64{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []float32{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []float32{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []float32{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []float32{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []float32{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []float64{0, 1})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []float64{0, 1})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []float64{0, 1})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []float64{0, 1})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []float64{0, 1})

	decodeSuccess(t, GetBuffer().Ints("", []int{0, 1}), []string{"0", "1"})
	decodeSuccess(t, GetBuffer().Int8s("", []int8{0, 1}), []string{"0", "1"})
	decodeSuccess(t, GetBuffer().Int16s("", []int16{0, 1}), []string{"0", "1"})
	decodeSuccess(t, GetBuffer().Int32s("", []int32{0, 1}), []string{"0", "1"})
	decodeSuccess(t, GetBuffer().Int64s("", []int64{0, 1}), []string{"0", "1"})

	decodeFail[uintptr](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to uintptr"))
	decodeFail[complex64](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to complex64"))
	decodeFail[complex128](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to complex128"))
	decodeFail[func()](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to func()"))
	decodeFail[chan int](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to struct {}"))

	decodeFail[[]uintptr](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []uintptr"))
	decodeFail[[]complex64](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []complex64"))
	decodeFail[[]complex128](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []complex128"))
	decodeFail[[]func()](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []func()"))
	decodeFail[[]chan int](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []chan int"))
	decodeFail[[]map[string]bool](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []map[string]bool"))
	decodeFail[[]struct{}](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to []struct {}"))

	decodeFail[[2]bool](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]bool overflow"))
	decodeFail[[2]int](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]int overflow"))
	decodeFail[[2]int8](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]int8 overflow"))
	decodeFail[[2]int16](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]int16 overflow"))
	decodeFail[[2]int32](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]int32 overflow"))
	decodeFail[[2]int64](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]int64 overflow"))
	decodeFail[[2]uint](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]uint overflow"))
	decodeFail[[2]uint8](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]uint8 overflow"))
	decodeFail[[2]uint16](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]uint16 overflow"))
	decodeFail[[2]uint32](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]uint32 overflow"))
	decodeFail[[2]uint64](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]uint64 overflow"))
	decodeFail[[2]float32](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]float32 overflow"))
	decodeFail[[2]float64](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]float64 overflow"))
	decodeFail[[2]string](t, GetBuffer().Ints("", []int{0, 1, 2}), errors.New("array [2]string overflow"))

	decodeSuccess[[3]bool](t, GetBuffer().Ints("", []int{0, 1, 0}), [3]bool{false, true, false})
	decodeSuccess[[3]int](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]int{0, 1, 2})
	decodeSuccess[[3]int8](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]int8{0, 1, 2})
	decodeSuccess[[3]int16](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]int16{0, 1, 2})
	decodeSuccess[[3]int32](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]int32{0, 1, 2})
	decodeSuccess[[3]int64](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]int64{0, 1, 2})
	decodeSuccess[[3]uint](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]uint{0, 1, 2})
	decodeSuccess[[3]uint8](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]uint8{0, 1, 2})
	decodeSuccess[[3]uint16](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]uint16{0, 1, 2})
	decodeSuccess[[3]uint32](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]uint32{0, 1, 2})
	decodeSuccess[[3]uint64](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]uint64{0, 1, 2})
	decodeSuccess[[3]float32](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]float32{0, 1, 2})
	decodeSuccess[[3]float64](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]float64{0, 1, 2})
	decodeSuccess[[3]string](t, GetBuffer().Ints("", []int{0, 1, 2}), [3]string{"0", "1", "2"})

	decodeFail[[3]uintptr](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]uintptr"))
	decodeFail[[3]complex64](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]complex64"))
	decodeFail[[3]complex128](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]complex128"))
	decodeFail[[3]func()](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]func()"))
	decodeFail[[3]chan int](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]chan int"))
	decodeFail[[3]map[string]bool](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]map[string]bool"))
	decodeFail[[3]struct{}](t, GetBuffer().Ints("", []int{0, 1}), errors.New("can't convert []int to [3]struct {}"))
}

func TestEncode_Uint(t *testing.T) {

	encodeSuccess(t, uint(3), GetBuffer().Uint("", 3))
	encodeSuccess(t, uint8(3), GetBuffer().Uint("", 3))
	encodeSuccess(t, uint16(3), GetBuffer().Uint("", 3))
	encodeSuccess(t, uint32(3), GetBuffer().Uint("", 3))
	encodeSuccess(t, uint64(3), GetBuffer().Uint("", 3))

	encodeSuccess(t, Ptr(uint(3)), GetBuffer().Uint("", 3))
	encodeSuccess(t, Ptr(uint8(3)), GetBuffer().Uint("", 3))
	encodeSuccess(t, Ptr(uint16(3)), GetBuffer().Uint("", 3))
	encodeSuccess(t, Ptr(uint32(3)), GetBuffer().Uint("", 3))
	encodeSuccess(t, Ptr(uint64(3)), GetBuffer().Uint("", 3))
}

func TestDecode_Uint(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Uint("", 3), uint64(3))

	decodeSuccess(t, GetBuffer().Uint("", 0), false)
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(true))

	decodeSuccess(t, GetBuffer().Uint("", 3), int(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), int8(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), int16(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), int32(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), int64(3))

	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(int(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(int8(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(int16(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(int32(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(int64(3)))

	decodeSuccess(t, GetBuffer().Uint("", 3), uint(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), uint8(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), uint16(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), uint32(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), uint64(3))

	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(uint(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(uint8(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(uint16(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(uint32(3)))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(uint64(3)))

	decodeSuccess(t, GetBuffer().Uint("", 3), float32(3))
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr(float64(3)))

	decodeSuccess(t, GetBuffer().Uint("", 3), "3")
	decodeSuccess(t, GetBuffer().Uint("", 3), Ptr("3"))

	decodeFail[uintptr](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to uintptr"))
	decodeFail[complex64](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to complex64"))
	decodeFail[complex128](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to complex128"))
	decodeFail[[]int](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to []int"))
	decodeFail[[3]int](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to [3]int"))
	decodeFail[func()](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to func()"))
	decodeFail[chan int](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to struct {}"))

	decodeFail[*uintptr](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to uintptr"))
	decodeFail[*complex64](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to complex64"))
	decodeFail[*complex128](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to complex128"))
	decodeFail[*[]int](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to []int"))
	decodeFail[*[3]int](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to [3]int"))
	decodeFail[*func()](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to func()"))
	decodeFail[*chan int](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to chan int"))
	decodeFail[*map[string]bool](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to map[string]bool"))
	decodeFail[*struct{}](t, GetBuffer().Uint("", 3), errors.New("can't convert uint64 to struct {}"))
}

func TestEncode_Uints(t *testing.T) {
	encodeSuccess(t, []uint{0, 1, 2}, GetBuffer().Uints("", []uint{0, 1, 2}))
	encodeSuccess(t, []uint8{0, 1, 2}, GetBuffer().Uint8s("", []uint8{0, 1, 2}))
	encodeSuccess(t, []uint16{0, 1, 2}, GetBuffer().Uint16s("", []uint16{0, 1, 2}))
	encodeSuccess(t, []uint32{0, 1, 2}, GetBuffer().Uint32s("", []uint32{0, 1, 2}))
	encodeSuccess(t, []uint64{0, 1, 2}, GetBuffer().Uint64s("", []uint64{0, 1, 2}))
}

func TestEncode_Float(t *testing.T) {
	encodeSuccess(t, float32(3), GetBuffer().Float("", 3))
	encodeSuccess(t, Ptr(float64(3)), GetBuffer().Float("", 3))
}

func TestDecode_Float(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Float("", 3), float64(3))

	decodeSuccess(t, GetBuffer().Float("", 0), false)
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(true))

	decodeSuccess(t, GetBuffer().Float("", 3), int(3))
	decodeSuccess(t, GetBuffer().Float("", 3), int8(3))
	decodeSuccess(t, GetBuffer().Float("", 3), int16(3))
	decodeSuccess(t, GetBuffer().Float("", 3), int32(3))
	decodeSuccess(t, GetBuffer().Float("", 3), int64(3))

	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(int(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(int8(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(int16(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(int32(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(int64(3)))

	decodeSuccess(t, GetBuffer().Float("", 3), uint(3))
	decodeSuccess(t, GetBuffer().Float("", 3), uint8(3))
	decodeSuccess(t, GetBuffer().Float("", 3), uint16(3))
	decodeSuccess(t, GetBuffer().Float("", 3), uint32(3))
	decodeSuccess(t, GetBuffer().Float("", 3), uint64(3))

	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(uint(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(uint8(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(uint16(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(uint32(3)))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(uint64(3)))

	decodeSuccess(t, GetBuffer().Float("", 3), float32(3))
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr(float64(3)))

	decodeSuccess(t, GetBuffer().Float("", 3), "3")
	decodeSuccess(t, GetBuffer().Float("", 3), Ptr("3"))

	decodeFail[uintptr](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to uintptr"))
	decodeFail[complex64](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to complex64"))
	decodeFail[complex128](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to complex128"))
	decodeFail[[]int](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to []int"))
	decodeFail[[3]int](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to [3]int"))
	decodeFail[func()](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to func()"))
	decodeFail[chan int](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to struct {}"))

	decodeFail[*uintptr](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to uintptr"))
	decodeFail[*complex64](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to complex64"))
	decodeFail[*complex128](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to complex128"))
	decodeFail[*[]int](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to []int"))
	decodeFail[*[3]int](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to [3]int"))
	decodeFail[*func()](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to func()"))
	decodeFail[*chan int](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to chan int"))
	decodeFail[*map[string]bool](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to map[string]bool"))
	decodeFail[*struct{}](t, GetBuffer().Float("", 3), errors.New("can't convert float64 to struct {}"))
}

func TestEncode_Floats(t *testing.T) {
	encodeSuccess(t, []float32{0, 1, 2}, GetBuffer().Float32s("", []float32{0, 1, 2}))
	encodeSuccess(t, []float64{0, 1, 2}, GetBuffer().Float64s("", []float64{0, 1, 2}))
}
