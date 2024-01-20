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

func TestConvert_String(t *testing.T) {

	success(t, "true", true)
	success(t, "false", false)
	success(t, "3", int(3))
	success(t, "3", int8(3))
	success(t, "3", int16(3))
	success(t, "3", int32(3))
	success(t, "3", int64(3))
	success(t, "3", uint(3))
	success(t, "3", uint8(3))
	success(t, "3", uint16(3))
	success(t, "3", uint32(3))
	success(t, "3", uint64(3))
	success(t, "3", float32(3))
	success(t, "3", float64(3))
	success(t, "3", string("3"))

	success(t, "true", Ptr(true))
	success(t, "false", Ptr(false))
	success(t, "3", Ptr(int(3)))
	success(t, "3", Ptr(int8(3)))
	success(t, "3", Ptr(int16(3)))
	success(t, "3", Ptr(int32(3)))
	success(t, "3", Ptr(int64(3)))
	success(t, "3", Ptr(uint(3)))
	success(t, "3", Ptr(uint8(3)))
	success(t, "3", Ptr(uint16(3)))
	success(t, "3", Ptr(uint32(3)))
	success(t, "3", Ptr(uint64(3)))
	success(t, "3", Ptr(float32(3)))
	success(t, "3", Ptr(float64(3)))
	success(t, "3", Ptr(string("3")))
}
