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
	"strconv"
	"testing"
)

func TestEncode_String(t *testing.T) {
	encodeSuccess(t, "3", GetBuffer().String("", "3"))
	encodeSuccess(t, Ptr("x3"), GetBuffer().String("", "x3"))
}

func TestDecode_String(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().String("", "3"), "3")
	decodeSuccess[interface{}](t, GetBuffer().String("", "x3"), "x3")

	decodeSuccess(t, GetBuffer().String("", "false"), false)
	decodeSuccess(t, GetBuffer().String("", "true"), Ptr(true))

	decodeFail[bool](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseBool", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*bool](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseBool", Num: "x3", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().String("", "3"), int(3))
	decodeSuccess(t, GetBuffer().String("", "3"), int8(3))
	decodeSuccess(t, GetBuffer().String("", "3"), int16(3))
	decodeSuccess(t, GetBuffer().String("", "3"), int32(3))
	decodeSuccess(t, GetBuffer().String("", "3"), int64(3))

	decodeFail[int](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[int8](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[int16](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[int32](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[int64](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(int(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(int8(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(int16(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(int32(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(int64(3)))

	decodeFail[*int](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*int8](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*int16](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*int32](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*int64](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().String("", "3"), uint(3))
	decodeSuccess(t, GetBuffer().String("", "3"), uint8(3))
	decodeSuccess(t, GetBuffer().String("", "3"), uint16(3))
	decodeSuccess(t, GetBuffer().String("", "3"), uint32(3))
	decodeSuccess(t, GetBuffer().String("", "3"), uint64(3))

	decodeFail[uint](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[uint8](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[uint16](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[uint32](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[uint64](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(uint(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(uint8(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(uint16(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(uint32(3)))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(uint64(3)))

	decodeFail[*uint](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*uint8](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*uint16](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*uint32](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*uint64](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().String("", "3"), float32(3))
	decodeSuccess(t, GetBuffer().String("", "3"), Ptr(float64(3)))

	decodeFail[float32](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
	decodeFail[*float64](t, GetBuffer().String("", "x3"), &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().String("", "3"), "3")
	decodeSuccess(t, GetBuffer().String("", "x3"), Ptr("x3"))

	decodeFail[uintptr](t, GetBuffer().String("", "3"), errors.New("can't convert string to uintptr"))
	decodeFail[complex64](t, GetBuffer().String("", "3"), errors.New("can't convert string to complex64"))
	decodeFail[complex128](t, GetBuffer().String("", "3"), errors.New("can't convert string to complex128"))
	decodeFail[[]int](t, GetBuffer().String("", "3"), errors.New("can't convert string to []int"))
	decodeFail[[3]int](t, GetBuffer().String("", "3"), errors.New("can't convert string to [3]int"))
	decodeFail[func()](t, GetBuffer().String("", "3"), errors.New("can't convert string to func()"))
	decodeFail[chan int](t, GetBuffer().String("", "3"), errors.New("can't convert string to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().String("", "3"), errors.New("can't convert string to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().String("", "3"), errors.New("can't convert string to struct {}"))

	decodeFail[*uintptr](t, GetBuffer().String("", "3"), errors.New("can't convert string to uintptr"))
	decodeFail[*complex64](t, GetBuffer().String("", "3"), errors.New("can't convert string to complex64"))
	decodeFail[*complex128](t, GetBuffer().String("", "3"), errors.New("can't convert string to complex128"))
	decodeFail[*[]int](t, GetBuffer().String("", "3"), errors.New("can't convert string to []int"))
	decodeFail[*[3]int](t, GetBuffer().String("", "3"), errors.New("can't convert string to [3]int"))
	decodeFail[*func()](t, GetBuffer().String("", "3"), errors.New("can't convert string to func()"))
	decodeFail[*chan int](t, GetBuffer().String("", "3"), errors.New("can't convert string to chan int"))
	decodeFail[*map[string]bool](t, GetBuffer().String("", "3"), errors.New("can't convert string to map[string]bool"))
	decodeFail[*struct{}](t, GetBuffer().String("", "3"), errors.New("can't convert string to struct {}"))
}

func TestEncode_Strings(t *testing.T) {
	encodeSuccess(t, []string{"0", "1"}, GetBuffer().Strings("", []string{"0", "1"}))
}

func TestDecode_Strings(t *testing.T) {

	decodeSuccess[interface{}](t, GetBuffer().Strings("", []string{"0", "1"}), []string{"0", "1"})
	decodeSuccess[interface{}](t, GetBuffer().Strings("", []string{"x0", "x1"}), []string{"x0", "x1"})

	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []bool{false, true})
	decodeFail[[]bool](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseBool", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]bool](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseBool", Num: "x0", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []int{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []int8{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []int16{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []int32{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []int64{0, 1})

	decodeFail[[]int](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]int8](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]int16](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]int32](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]int64](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]int](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]int8](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]int16](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]int32](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]int64](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseInt", Num: "x0", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []uint{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []uint8{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []uint16{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []uint32{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []uint64{0, 1})

	decodeFail[[]uint](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]uint8](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]uint16](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]uint32](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]uint64](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]uint](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]uint8](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]uint16](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]uint32](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]uint64](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseUint", Num: "x0", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []float32{0, 1})
	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []float64{0, 1})

	decodeFail[[]float32](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseFloat", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[]float64](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseFloat", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]float32](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseFloat", Num: "x0", Err: strconv.ErrSyntax})
	decodeFail[[1]float64](t, GetBuffer().Strings("", []string{"x0"}), &strconv.NumError{Func: "ParseFloat", Num: "x0", Err: strconv.ErrSyntax})

	decodeSuccess(t, GetBuffer().Strings("", []string{"0", "1"}), []string{"0", "1"})

	decodeFail[uintptr](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to uintptr"))
	decodeFail[complex64](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to complex64"))
	decodeFail[complex128](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to complex128"))
	decodeFail[func()](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to func()"))
	decodeFail[chan int](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to chan int"))
	decodeFail[map[string]bool](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to map[string]bool"))
	decodeFail[struct{}](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to struct {}"))

	decodeFail[[]uintptr](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []uintptr"))
	decodeFail[[]complex64](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []complex64"))
	decodeFail[[]complex128](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []complex128"))
	decodeFail[[]func()](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []func()"))
	decodeFail[[]chan int](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []chan int"))
	decodeFail[[]map[string]bool](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []map[string]bool"))
	decodeFail[[]struct{}](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to []struct {}"))

	decodeFail[[2]bool](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]bool overflow"))
	decodeFail[[2]int](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]int overflow"))
	decodeFail[[2]int8](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]int8 overflow"))
	decodeFail[[2]int16](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]int16 overflow"))
	decodeFail[[2]int32](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]int32 overflow"))
	decodeFail[[2]int64](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]int64 overflow"))
	decodeFail[[2]uint](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]uint overflow"))
	decodeFail[[2]uint8](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]uint8 overflow"))
	decodeFail[[2]uint16](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]uint16 overflow"))
	decodeFail[[2]uint32](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]uint32 overflow"))
	decodeFail[[2]uint64](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]uint64 overflow"))
	decodeFail[[2]float32](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]float32 overflow"))
	decodeFail[[2]float64](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]float64 overflow"))
	decodeFail[[2]string](t, GetBuffer().Strings("", []string{"0", "1", "2"}), errors.New("array [2]string overflow"))

	decodeSuccess[[3]bool](t, GetBuffer().Strings("", []string{"0", "1", "0"}), [3]bool{false, true, false})
	decodeSuccess[[3]int](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]int{0, 1, 2})
	decodeSuccess[[3]int8](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]int8{0, 1, 2})
	decodeSuccess[[3]int16](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]int16{0, 1, 2})
	decodeSuccess[[3]int32](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]int32{0, 1, 2})
	decodeSuccess[[3]int64](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]int64{0, 1, 2})
	decodeSuccess[[3]uint](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]uint{0, 1, 2})
	decodeSuccess[[3]uint8](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]uint8{0, 1, 2})
	decodeSuccess[[3]uint16](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]uint16{0, 1, 2})
	decodeSuccess[[3]uint32](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]uint32{0, 1, 2})
	decodeSuccess[[3]uint64](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]uint64{0, 1, 2})
	decodeSuccess[[3]float32](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]float32{0, 1, 2})
	decodeSuccess[[3]float64](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]float64{0, 1, 2})
	decodeSuccess[[3]string](t, GetBuffer().Strings("", []string{"0", "1", "2"}), [3]string{"0", "1", "2"})

	decodeFail[[3]uintptr](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]uintptr"))
	decodeFail[[3]complex64](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]complex64"))
	decodeFail[[3]complex128](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]complex128"))
	decodeFail[[3]func()](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]func()"))
	decodeFail[[3]chan int](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]chan int"))
	decodeFail[[3]map[string]bool](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]map[string]bool"))
	decodeFail[[3]struct{}](t, GetBuffer().Strings("", []string{"0", "1"}), errors.New("can't convert []string to [3]struct {}"))
}
