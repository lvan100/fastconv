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

func TestConvert_String(t *testing.T) {

	success[string, bool](t, "true", true)
	success[string, bool](t, "false", false)
	success[string, int](t, "3", int(3))
	success[string, int8](t, "3", int8(3))
	success[string, int16](t, "3", int16(3))
	success[string, int32](t, "3", int32(3))
	success[string, int64](t, "3", int64(3))
	success[string, uint](t, "3", uint(3))
	success[string, uint8](t, "3", uint8(3))
	success[string, uint16](t, "3", uint16(3))
	success[string, uint32](t, "3", uint32(3))
	success[string, uint64](t, "3", uint64(3))
	success[string, float32](t, "3", float32(3))
	success[string, float64](t, "3", float64(3))
	success[string, string](t, "3", string("3"))
	success[string, interface{}](t, "3", string("3"))

	fail[string, bool](t, "x3", &strconv.NumError{Func: "ParseBool", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, int](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, int8](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, int16](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, int32](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, int64](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, uint](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, uint8](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, uint16](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, uint32](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, uint64](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, float32](t, "x3", &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, float64](t, "x3", &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})

	fail[string, uintptr](t, "3", errors.New("can't convert string to uintptr"))
	fail[string, complex64](t, "3", errors.New("can't convert string to complex64"))
	fail[string, complex128](t, "3", errors.New("can't convert string to complex128"))
	fail[string, []int](t, "3", errors.New("can't convert string to []int"))
	fail[string, [3]int](t, "3", errors.New("can't convert string to [3]int"))
	fail[string, func()](t, "3", errors.New("can't convert string to func()"))
	fail[string, chan int](t, "3", errors.New("can't convert string to chan int"))
	fail[string, map[string]bool](t, "3", errors.New("can't convert string to map[string]bool"))
	fail[string, struct{}](t, "3", errors.New("can't convert string to struct {}"))

	success[string, *bool](t, "true", Ptr(true))
	success[string, *bool](t, "false", Ptr(false))
	success[string, *int](t, "3", Ptr(int(3)))
	success[string, *int8](t, "3", Ptr(int8(3)))
	success[string, *int16](t, "3", Ptr(int16(3)))
	success[string, *int32](t, "3", Ptr(int32(3)))
	success[string, *int64](t, "3", Ptr(int64(3)))
	success[string, *uint](t, "3", Ptr(uint(3)))
	success[string, *uint8](t, "3", Ptr(uint8(3)))
	success[string, *uint16](t, "3", Ptr(uint16(3)))
	success[string, *uint32](t, "3", Ptr(uint32(3)))
	success[string, *uint64](t, "3", Ptr(uint64(3)))
	success[string, *float32](t, "3", Ptr(float32(3)))
	success[string, *float64](t, "3", Ptr(float64(3)))
	success[string, *string](t, "3", Ptr(string("3")))

	fail[string, *bool](t, "x3", &strconv.NumError{Func: "ParseBool", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *int](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *int8](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *int16](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *int32](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *int64](t, "x3", &strconv.NumError{Func: "ParseInt", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *uint](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *uint8](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *uint16](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *uint32](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *uint64](t, "x3", &strconv.NumError{Func: "ParseUint", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *float32](t, "x3", &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})
	fail[string, *float64](t, "x3", &strconv.NumError{Func: "ParseFloat", Num: "x3", Err: strconv.ErrSyntax})

	fail[string, *uintptr](t, "3", errors.New("can't convert string to uintptr"))
	fail[string, *complex64](t, "3", errors.New("can't convert string to complex64"))
	fail[string, *complex128](t, "3", errors.New("can't convert string to complex128"))
	fail[string, *[]int](t, "3", errors.New("can't convert string to []int"))
	fail[string, *[3]int](t, "3", errors.New("can't convert string to [3]int"))
	fail[string, *func()](t, "3", errors.New("can't convert string to func()"))
	fail[string, *chan int](t, "3", errors.New("can't convert string to chan int"))
	fail[string, *map[string]bool](t, "3", errors.New("can't convert string to map[string]bool"))
	fail[string, *struct{}](t, "3", errors.New("can't convert string to struct {}"))
}
