/*
 * Copyright 2023 the original author or authors.
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

// DeepCopy returns a "deep" copy of v.
func DeepCopy[T any](v T) (T, error) {
	l := GetBuffer()
	defer PutBuffer(l)
	var r T
	if err := Encode(l, v); err != nil {
		return r, err
	}
	if err := Decode(l, &r); err != nil {
		return r, err
	}
	return r, nil
}

// Convert converts src to dest by "deep" copy.
func Convert(src, dest any) error {
	l := GetBuffer()
	defer PutBuffer(l)
	if err := Encode(l, src); err != nil {
		return err
	}
	if err := Decode(l, dest); err != nil {
		return err
	}
	return nil
}
