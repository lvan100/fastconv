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

//func TestConvert_Strings(t *testing.T) {
//
//	success[[]string, []bool](t, []string{"true", "false", "true"}, []bool{true, false, true})
//	success[[]string, []int](t, []string{"1", "2", "3"}, []int{1, 2, 3})
//	success[[]string, []int8](t, []string{"1", "2", "3"}, []int8{1, 2, 3})
//	success[[]string, []int16](t, []string{"1", "2", "3"}, []int16{1, 2, 3})
//	success[[]string, []int32](t, []string{"1", "2", "3"}, []int32{1, 2, 3})
//	success[[]string, []int64](t, []string{"1", "2", "3"}, []int64{1, 2, 3})
//	success[[]string, []uint](t, []string{"1", "2", "3"}, []uint{1, 2, 3})
//	success[[]string, []uint8](t, []string{"1", "2", "3"}, []uint8{1, 2, 3})
//	success[[]string, []uint16](t, []string{"1", "2", "3"}, []uint16{1, 2, 3})
//	success[[]string, []uint32](t, []string{"1", "2", "3"}, []uint32{1, 2, 3})
//	success[[]string, []uint64](t, []string{"1", "2", "3"}, []uint64{1, 2, 3})
//	success[[]string, []float32](t, []string{"1", "2", "3"}, []float32{1, 2, 3})
//	success[[]string, []float64](t, []string{"1", "2", "3"}, []float64{1, 2, 3})
//	success[[]string, []string](t, []string{"1", "2", "3"}, []string{"1", "2", "3"})
//	success[[]string, interface{}](t, []string{"1", "2", "3"}, []string{"1", "2", "3"})
//
//	fail[[]string, []bool](t, []string{"x3"}, &strconv.NumError{Func: "ParseBool", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []int](t, []string{"x3"}, &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []int8](t, []string{"x3"}, &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []int16](t, []string{"x3"}, &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []int32](t, []string{"x3"}, &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []int64](t, []string{"x3"}, &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []uint](t, []string{"x3"}, &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []uint8](t, []string{"x3"}, &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []uint16](t, []string{"x3"}, &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []uint32](t, []string{"x3"}, &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []uint64](t, []string{"x3"}, &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []float32](t, []string{"x3"}, &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
//	fail[[]string, []float64](t, []string{"x3"}, &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
//
//	fail[[]string, uintptr](t, []string{"3"}, errors.New("can't convert []string to uintptr"))
//	fail[[]string, complex64](t, []string{"3"}, errors.New("can't convert []string to complex64"))
//	fail[[]string, complex128](t, []string{"3"}, errors.New("can't convert []string to complex128"))
//	fail[[]string, func()](t, []string{"3"}, errors.New("can't convert []string to func()"))
//	fail[[]string, chan int](t, []string{"3"}, errors.New("can't convert []string to chan int"))
//	fail[[]string, map[string]bool](t, []string{"3"}, errors.New("can't convert []string to map[string]bool"))
//	fail[[]string, struct{}](t, []string{"3"}, errors.New("can't convert []string to struct {}"))
//
//	fail[[]string, []uintptr](t, []string{"3"}, errors.New("can't convert []string to []uintptr"))
//	fail[[]string, []complex64](t, []string{"3"}, errors.New("can't convert []string to []complex64"))
//	fail[[]string, []complex128](t, []string{"3"}, errors.New("can't convert []string to []complex128"))
//	fail[[]string, []func()](t, []string{"3"}, errors.New("can't convert []string to []func()"))
//	fail[[]string, []chan int](t, []string{"3"}, errors.New("can't convert []string to []chan int"))
//	fail[[]string, []map[string]bool](t, []string{"3"}, errors.New("can't convert []string to []map[string]bool"))
//	fail[[]string, []struct{}](t, []string{"3"}, errors.New("can't convert []string to []struct {}"))
//
//	//success[string, *bool](t, "true", Ptr(true))
//	//success[string, *bool](t, "false", Ptr(false))
//	//success[string, *int](t, "3", Ptr(int(3)))
//	//success[string, *int8](t, "3", Ptr(int8(3)))
//	//success[string, *int16](t, "3", Ptr(int16(3)))
//	//success[string, *int32](t, "3", Ptr(int32(3)))
//	//success[string, *int64](t, "3", Ptr(int64(3)))
//	//success[string, *uint](t, "3", Ptr(uint(3)))
//	//success[string, *uint8](t, "3", Ptr(uint8(3)))
//	//success[string, *uint16](t, "3", Ptr(uint16(3)))
//	//success[string, *uint32](t, "3", Ptr(uint32(3)))
//	//success[string, *uint64](t, "3", Ptr(uint64(3)))
//	//success[string, *float32](t, "3", Ptr(float32(3)))
//	//success[string, *float64](t, "3", Ptr(float64(3)))
//	//success[string, *string](t, "3", Ptr("3"))
//	//
//	//fail[string, *bool](t, "x3", &strconv.NumError{Func: "ParseBool", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *int](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *int8](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *int16](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *int32](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *int64](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *uint](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *uint8](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *uint16](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *uint32](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *uint64](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *float32](t, "x3", &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
//	//fail[string, *float64](t, "x3", &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
//	//
//	//fail[string, *uintptr](t, "3", errors.New("can't convert string to uintptr"))
//	//fail[string, *complex64](t, "3", errors.New("can't convert string to complex64"))
//	//fail[string, *complex128](t, "3", errors.New("can't convert string to complex128"))
//	//fail[string, *[]int](t, "3", errors.New("can't convert string to []int"))
//	//fail[string, *[3]int](t, "3", errors.New("can't convert string to [3]int"))
//	//fail[string, *func()](t, "3", errors.New("can't convert string to func()"))
//	//fail[string, *chan int](t, "3", errors.New("can't convert string to chan int"))
//	//fail[string, *map[string]bool](t, "3", errors.New("can't convert string to map[string]bool"))
//	//fail[string, *struct{}](t, "3", errors.New("can't convert string to struct {}"))
//}
