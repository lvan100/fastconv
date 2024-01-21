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

func TestConvert_Bool(t *testing.T) {

	t.Run("false", func(t *testing.T) {
		success[bool, bool](t, false, false)
		success[bool, int](t, false, 0)
		success[bool, int8](t, false, 0)
		success[bool, int16](t, false, 0)
		success[bool, int32](t, false, 0)
		success[bool, int64](t, false, 0)
		success[bool, uint](t, false, 0)
		success[bool, uint8](t, false, 0)
		success[bool, uint16](t, false, 0)
		success[bool, uint32](t, false, 0)
		success[bool, uint64](t, false, 0)
		success[bool, float32](t, false, 0)
		success[bool, float64](t, false, 0)
		success[bool, string](t, false, "false")
		success[bool, interface{}](t, false, false)

		fail[bool, uintptr](t, false, errors.New("can't convert bool to uintptr"))
		fail[bool, complex64](t, false, errors.New("can't convert bool to complex64"))
		fail[bool, complex128](t, false, errors.New("can't convert bool to complex128"))
		fail[bool, []int](t, false, errors.New("can't convert bool to []int"))
		fail[bool, [3]int](t, false, errors.New("can't convert bool to [3]int"))
		fail[bool, func()](t, false, errors.New("can't convert bool to func()"))
		fail[bool, chan int](t, false, errors.New("can't convert bool to chan int"))
		fail[bool, map[string]bool](t, false, errors.New("can't convert bool to map[string]bool"))
		fail[bool, struct{}](t, false, errors.New("can't convert bool to struct {}"))

		success[bool, *bool](t, false, Ptr(false))
		success[bool, *int](t, false, Ptr(int(0)))
		success[bool, *int8](t, false, Ptr(int8(0)))
		success[bool, *int16](t, false, Ptr(int16(0)))
		success[bool, *int32](t, false, Ptr(int32(0)))
		success[bool, *int64](t, false, Ptr(int64(0)))
		success[bool, *uint](t, false, Ptr(uint(0)))
		success[bool, *uint8](t, false, Ptr(uint8(0)))
		success[bool, *uint16](t, false, Ptr(uint16(0)))
		success[bool, *uint32](t, false, Ptr(uint32(0)))
		success[bool, *uint64](t, false, Ptr(uint64(0)))
		success[bool, *float32](t, false, Ptr(float32(0)))
		success[bool, *float64](t, false, Ptr(float64(0)))
		success[bool, *string](t, false, Ptr("false"))

		fail[bool, *uintptr](t, false, errors.New("can't convert bool to uintptr"))
		fail[bool, *complex64](t, false, errors.New("can't convert bool to complex64"))
		fail[bool, *complex128](t, false, errors.New("can't convert bool to complex128"))
		fail[bool, *[]int](t, false, errors.New("can't convert bool to []int"))
		fail[bool, *[3]int](t, false, errors.New("can't convert bool to [3]int"))
		fail[bool, *func()](t, false, errors.New("can't convert bool to func()"))
		fail[bool, *chan int](t, false, errors.New("can't convert bool to chan int"))
		fail[bool, *map[string]bool](t, false, errors.New("can't convert bool to map[string]bool"))
		fail[bool, *struct{}](t, false, errors.New("can't convert bool to struct {}"))
	})

	t.Run("true", func(t *testing.T) {
		success[bool, bool](t, true, true)
		success[bool, int](t, true, 1)
		success[bool, int8](t, true, 1)
		success[bool, int16](t, true, 1)
		success[bool, int32](t, true, 1)
		success[bool, int64](t, true, 1)
		success[bool, uint](t, true, 1)
		success[bool, uint8](t, true, 1)
		success[bool, uint16](t, true, 1)
		success[bool, uint32](t, true, 1)
		success[bool, uint64](t, true, 1)
		success[bool, float32](t, true, 1)
		success[bool, float64](t, true, 1)
		success[bool, string](t, true, "true")
		success[bool, interface{}](t, true, true)

		fail[bool, uintptr](t, true, errors.New("can't convert bool to uintptr"))
		fail[bool, complex64](t, true, errors.New("can't convert bool to complex64"))
		fail[bool, complex128](t, true, errors.New("can't convert bool to complex128"))
		fail[bool, []int](t, true, errors.New("can't convert bool to []int"))
		fail[bool, [3]int](t, true, errors.New("can't convert bool to [3]int"))
		fail[bool, func()](t, true, errors.New("can't convert bool to func()"))
		fail[bool, chan int](t, true, errors.New("can't convert bool to chan int"))
		fail[bool, map[string]bool](t, true, errors.New("can't convert bool to map[string]bool"))
		fail[bool, struct{}](t, true, errors.New("can't convert bool to struct {}"))

		success[bool, *bool](t, true, Ptr(true))
		success[bool, *int](t, true, Ptr(int(1)))
		success[bool, *int8](t, true, Ptr(int8(1)))
		success[bool, *int16](t, true, Ptr(int16(1)))
		success[bool, *int32](t, true, Ptr(int32(1)))
		success[bool, *int64](t, true, Ptr(int64(1)))
		success[bool, *uint](t, true, Ptr(uint(1)))
		success[bool, *uint8](t, true, Ptr(uint8(1)))
		success[bool, *uint16](t, true, Ptr(uint16(1)))
		success[bool, *uint32](t, true, Ptr(uint32(1)))
		success[bool, *uint64](t, true, Ptr(uint64(1)))
		success[bool, *float32](t, true, Ptr(float32(1)))
		success[bool, *float64](t, true, Ptr(float64(1)))
		success[bool, *string](t, true, Ptr(string("true")))

		fail[bool, *uintptr](t, true, errors.New("can't convert bool to uintptr"))
		fail[bool, *complex64](t, true, errors.New("can't convert bool to complex64"))
		fail[bool, *complex128](t, true, errors.New("can't convert bool to complex128"))
		fail[bool, *[]int](t, true, errors.New("can't convert bool to []int"))
		fail[bool, *[3]int](t, true, errors.New("can't convert bool to [3]int"))
		fail[bool, *func()](t, true, errors.New("can't convert bool to func()"))
		fail[bool, *chan int](t, true, errors.New("can't convert bool to chan int"))
		fail[bool, *map[string]bool](t, true, errors.New("can't convert bool to map[string]bool"))
		fail[bool, *struct{}](t, true, errors.New("can't convert bool to struct {}"))
	})
}
