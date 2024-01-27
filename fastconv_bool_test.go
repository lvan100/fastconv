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

func TestEncode_Bool(t *testing.T) {
	encodeSuccess(t, false, GetBuffer().Bool("", false))
	encodeSuccess(t, Ptr(true), GetBuffer().Bool("", true))
}

func TestDecode_Bool(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Bool("", false), false)
	decodeSuccess[interface{}](t, GetBuffer().Bool("", true), true)

	decodeSuccess(t, GetBuffer().Bool("", false), false)
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(true))

	decodeSuccess(t, GetBuffer().Bool("", false), int(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int8(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int16(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int32(0))
	decodeSuccess(t, GetBuffer().Bool("", false), int64(0))

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

	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint8(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint16(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint32(1)))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(uint64(1)))

	decodeSuccess(t, GetBuffer().Bool("", false), float32(0))
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr(float64(1)))

	decodeSuccess(t, GetBuffer().Bool("", false), "false")
	decodeSuccess(t, GetBuffer().Bool("", true), Ptr("true"))

	decodeFail[uintptr](t, GetBuffer().Bool("", false), errors.New("can't convert bool to uintptr"))
	decodeFail[complex64](t, GetBuffer().Bool("", false), errors.New("can't convert bool to complex64"))
	decodeFail[complex128](t, GetBuffer().Bool("", false), errors.New("can't convert bool to complex128"))
	decodeFail[[]int](t, GetBuffer().Bool("", false), errors.New("can't convert bool to []int"))
	decodeFail[[3]int](t, GetBuffer().Bool("", false), errors.New("can't convert bool to [3]int"))
	decodeFail[func()](t, GetBuffer().Bool("", false), errors.New("can't convert bool to func()"))
	decodeFail[chan int](t, GetBuffer().Bool("", false), errors.New("can't convert bool to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Bool("", false), errors.New("can't convert bool to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Bool("", false), errors.New("can't convert bool to struct {}"))

	decodeFail[*uintptr](t, GetBuffer().Bool("", true), errors.New("can't convert bool to uintptr"))
	decodeFail[*complex64](t, GetBuffer().Bool("", true), errors.New("can't convert bool to complex64"))
	decodeFail[*complex128](t, GetBuffer().Bool("", true), errors.New("can't convert bool to complex128"))
	decodeFail[*[]int](t, GetBuffer().Bool("", true), errors.New("can't convert bool to []int"))
	decodeFail[*[3]int](t, GetBuffer().Bool("", true), errors.New("can't convert bool to [3]int"))
	decodeFail[*func()](t, GetBuffer().Bool("", true), errors.New("can't convert bool to func()"))
	decodeFail[*chan int](t, GetBuffer().Bool("", true), errors.New("can't convert bool to chan int"))
	decodeFail[*map[string]bool](t, GetBuffer().Bool("", true), errors.New("can't convert bool to map[string]bool"))
	decodeFail[*struct{}](t, GetBuffer().Bool("", true), errors.New("can't convert bool to struct {}"))
}

func TestEncode_Bools(t *testing.T) {
	encodeSuccess(t, []bool{false, true, false}, GetBuffer().Bools("", []bool{false, true, false}))
}

func TestDecode_Bools(t *testing.T) {

	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []bool{false, true, false})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []int{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []int8{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []int16{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []int32{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []int64{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []uint{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []uint8{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []uint16{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []uint32{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []uint64{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []float32{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []float64{0, 1, 0})
	decodeSuccess(t, GetBuffer().Bools("", []bool{false, true, false}), []string{"false", "true", "false"})
	decodeSuccess[interface{}](t, GetBuffer().Bools("", []bool{false, true, false}), []bool{false, true, false})

	decodeFail[uintptr](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to uintptr"))
	decodeFail[complex64](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to complex64"))
	decodeFail[complex128](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to complex128"))
	decodeFail[func()](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to func()"))
	decodeFail[chan int](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to struct {}"))

	decodeFail[[]uintptr](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []uintptr"))
	decodeFail[[]complex64](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []complex64"))
	decodeFail[[]complex128](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []complex128"))
	decodeFail[[]func()](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []func()"))
	decodeFail[[]chan int](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []chan int"))
	decodeFail[[]map[string]bool](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []map[string]bool"))
	decodeFail[[]struct{}](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to []struct {}"))

	decodeFail[[2]bool](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]bool overflow"))
	decodeFail[[2]int](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]int overflow"))
	decodeFail[[2]int8](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]int8 overflow"))
	decodeFail[[2]int16](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]int16 overflow"))
	decodeFail[[2]int32](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]int32 overflow"))
	decodeFail[[2]int64](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]int64 overflow"))
	decodeFail[[2]uint](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]uint overflow"))
	decodeFail[[2]uint8](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]uint8 overflow"))
	decodeFail[[2]uint16](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]uint16 overflow"))
	decodeFail[[2]uint32](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]uint32 overflow"))
	decodeFail[[2]uint64](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]uint64 overflow"))
	decodeFail[[2]float32](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]float32 overflow"))
	decodeFail[[2]float64](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]float64 overflow"))
	decodeFail[[2]string](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("array [2]string overflow"))

	decodeSuccess[[3]bool](t, GetBuffer().Bools("", []bool{true, false, true}), [3]bool{true, false, true})
	decodeSuccess[[3]int](t, GetBuffer().Bools("", []bool{true, false, true}), [3]int{1, 0, 1})
	decodeSuccess[[3]int8](t, GetBuffer().Bools("", []bool{true, false, true}), [3]int8{1, 0, 1})
	decodeSuccess[[3]int16](t, GetBuffer().Bools("", []bool{true, false, true}), [3]int16{1, 0, 1})
	decodeSuccess[[3]int32](t, GetBuffer().Bools("", []bool{true, false, true}), [3]int32{1, 0, 1})
	decodeSuccess[[3]int64](t, GetBuffer().Bools("", []bool{true, false, true}), [3]int64{1, 0, 1})
	decodeSuccess[[3]uint](t, GetBuffer().Bools("", []bool{true, false, true}), [3]uint{1, 0, 1})
	decodeSuccess[[3]uint8](t, GetBuffer().Bools("", []bool{true, false, true}), [3]uint8{1, 0, 1})
	decodeSuccess[[3]uint16](t, GetBuffer().Bools("", []bool{true, false, true}), [3]uint16{1, 0, 1})
	decodeSuccess[[3]uint32](t, GetBuffer().Bools("", []bool{true, false, true}), [3]uint32{1, 0, 1})
	decodeSuccess[[3]uint64](t, GetBuffer().Bools("", []bool{true, false, true}), [3]uint64{1, 0, 1})
	decodeSuccess[[3]float32](t, GetBuffer().Bools("", []bool{true, false, true}), [3]float32{1, 0, 1})
	decodeSuccess[[3]float64](t, GetBuffer().Bools("", []bool{true, false, true}), [3]float64{1, 0, 1})
	decodeSuccess[[3]string](t, GetBuffer().Bools("", []bool{true, false, true}), [3]string{"true", "false", "true"})

	decodeFail[[3]uintptr](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]uintptr"))
	decodeFail[[3]complex64](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]complex64"))
	decodeFail[[3]complex128](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]complex128"))
	decodeFail[[3]func()](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]func()"))
	decodeFail[[3]chan int](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]chan int"))
	decodeFail[[3]map[string]bool](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]map[string]bool"))
	decodeFail[[3]struct{}](t, GetBuffer().Bools("", []bool{false, true, false}), errors.New("can't convert []bool to [3]struct {}"))
}
