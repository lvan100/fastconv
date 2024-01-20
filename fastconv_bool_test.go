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

import "testing"

func TestConvert_Bool(t *testing.T) {

	t.Run("false", func(t *testing.T) {
		success(t, false, false)
		success(t, false, int(0))
		success(t, false, int8(0))
		success(t, false, int16(0))
		success(t, false, int32(0))
		success(t, false, int64(0))
		success(t, false, uint(0))
		success(t, false, uint8(0))
		success(t, false, uint16(0))
		success(t, false, uint32(0))
		success(t, false, uint64(0))
		success(t, false, float32(0))
		success(t, false, float64(0))
		success(t, false, string("false"))

		success(t, false, Ptr(false))
		success(t, false, Ptr(int(0)))
		success(t, false, Ptr(int8(0)))
		success(t, false, Ptr(int16(0)))
		success(t, false, Ptr(int32(0)))
		success(t, false, Ptr(int64(0)))
		success(t, false, Ptr(uint(0)))
		success(t, false, Ptr(uint8(0)))
		success(t, false, Ptr(uint16(0)))
		success(t, false, Ptr(uint32(0)))
		success(t, false, Ptr(uint64(0)))
		success(t, false, Ptr(float32(0)))
		success(t, false, Ptr(float64(0)))
		success(t, false, Ptr(string("false")))
	})

	t.Run("true", func(t *testing.T) {
		success(t, true, true)
		success(t, true, int(1))
		success(t, true, int8(1))
		success(t, true, int16(1))
		success(t, true, int32(1))
		success(t, true, int64(1))
		success(t, true, uint(1))
		success(t, true, uint8(1))
		success(t, true, uint16(1))
		success(t, true, uint32(1))
		success(t, true, uint64(1))
		success(t, true, float32(1))
		success(t, true, float64(1))
		success(t, true, string("true"))

		success(t, true, Ptr(true))
		success(t, true, Ptr(int(1)))
		success(t, true, Ptr(int8(1)))
		success(t, true, Ptr(int16(1)))
		success(t, true, Ptr(int32(1)))
		success(t, true, Ptr(int64(1)))
		success(t, true, Ptr(uint(1)))
		success(t, true, Ptr(uint8(1)))
		success(t, true, Ptr(uint16(1)))
		success(t, true, Ptr(uint32(1)))
		success(t, true, Ptr(uint64(1)))
		success(t, true, Ptr(float32(1)))
		success(t, true, Ptr(float64(1)))
		success(t, true, Ptr(string("true")))
	})
}
