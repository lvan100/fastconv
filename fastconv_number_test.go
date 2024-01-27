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
