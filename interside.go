// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interside

import "strconv"

type Container interface {
	Append(value ...interface{}) Container
	Length() int
	IntSlice() []int
	UintSlice() []uint
	Int32Slice() []int32
	Int64Slice() []int64
	StringSlice() []string
	InterfaceSlice() []interface{}
}

func NewContainer() *container {
	return &container{}
}

type container struct {
	store []interface{}
}

// InterfaceSlice returns interface slice.
func (c *container) InterfaceSlice() []interface{} {
	return c.store
}

func (c *container) Append(value ...interface{}) Container {
	c.store = append(c.store, value...)

	return c
}

func (c *container) Length() int {
	return len(c.store)
}

func (c *container) IntSlice() []int {
	var data = make([]int, 0)
	for _, datum := range c.store {
		var converted int

		switch v := datum.(type) {
		case string:
			a, err := strconv.Atoi(v)
			if err != nil {
				converted = 0
				continue
			}

			converted = a
		case float32:
			converted = int(v)
		case bool:
			var tmp int

			if v == true {
				tmp = 1
			} else {
				tmp = 0
			}

			converted = tmp
		case float64:
			converted = int(v)
		case uint:
			converted = int(v)
		case uint8:
			converted = int(v)
		case uint16:
			converted = int(v)
		case uint32:
			converted = int(v)
		case uint64:
			converted = int(v)
		case int:
			converted = int(v)
		case int8:
			converted = int(v)
		case int16:
			converted = int(v)
		case int32:
			converted = int(v)
		case int64:
			converted = int(v)
		}

		data = append(data, converted)
	}

	return data
}

func (c *container) UintSlice() []uint {
	var data = make([]uint, 0)
	for _, datum := range c.store {
		var converted uint

		switch v := datum.(type) {
		case string:
			a, err := strconv.Atoi(v)
			if err != nil {
				converted = 0
				continue
			}

			converted = uint(a)
		case float32:
			converted = uint(v)
		case bool:
			var tmp uint

			if v == true {
				tmp = 1
			} else {
				tmp = 0
			}

			converted = tmp
		case float64:
			converted = uint(v)
		case uint:
			converted = uint(v)
		case uint8:
			converted = uint(v)
		case uint16:
			converted = uint(v)
		case uint32:
			converted = uint(v)
		case uint64:
			converted = uint(v)
		case int:
			converted = uint(v)
		case int8:
			converted = uint(v)
		case int16:
			converted = uint(v)
		case int32:
			converted = uint(v)
		case int64:
			converted = uint(v)
		}

		data = append(data, converted)
	}

	return data
}

func (c *container) Int32Slice() []int32 {
	var data = make([]int32, 0)
	for _, datum := range c.store {
		var converted int32

		switch v := datum.(type) {
		case string:
			a, err := strconv.Atoi(v)
			if err != nil {
				converted = 0
				continue
			}

			converted = int32(a)
		case float32:
			converted = int32(v)
		case bool:
			var tmp int32

			if v == true {
				tmp = 1
			} else {
				tmp = 0
			}

			converted = tmp
		case float64:
			converted = int32(v)
		case uint:
			converted = int32(v)
		case uint8:
			converted = int32(v)
		case uint16:
			converted = int32(v)
		case uint32:
			converted = int32(v)
		case uint64:
			converted = int32(v)
		case int:
			converted = int32(v)
		case int8:
			converted = int32(v)
		case int16:
			converted = int32(v)
		case int32:
			converted = int32(v)
		case int64:
			converted = int32(v)
		}

		data = append(data, converted)
	}

	return data
}

func (c *container) Int64Slice() []int64 {
	var data = make([]int64, 0)
	for _, datum := range c.store {
		var converted int64

		switch v := datum.(type) {
		case string:
			a, err := strconv.Atoi(v)
			if err != nil {
				converted = 0
				continue
			}

			converted = int64(a)
		case float32:
			converted = int64(v)
		case bool:
			var tmp int64

			if v == true {
				tmp = 1
			} else {
				tmp = 0
			}

			converted = tmp
		case float64:
			converted = int64(v)
		case uint:
			converted = int64(v)
		case uint8:
			converted = int64(v)
		case uint16:
			converted = int64(v)
		case uint32:
			converted = int64(v)
		case uint64:
			converted = int64(v)
		case int:
			converted = int64(v)
		case int8:
			converted = int64(v)
		case int16:
			converted = int64(v)
		case int32:
			converted = int64(v)
		case int64:
			converted = int64(v)
		}

		data = append(data, converted)
	}

	return data
}

func (c *container) StringSlice() []string {
	var data = make([]string, 0)
	for _, datum := range c.store {
		var converted string

		switch v := datum.(type) {
		case string:
			converted = v
		case float32:
			converted = strconv.FormatFloat(float64(v), 'f', 2, 32)
		case float64:
			converted = strconv.FormatFloat(v, 'f', 2, 64)
		case bool:
			var tmp string

			if v == true {
				tmp = "1"
			} else {
				tmp = "0"
			}

			converted = tmp
		case uint:
			converted = strconv.Itoa(int(v))
		case uint8:
			converted = strconv.Itoa(int(v))
		case uint16:
			converted = strconv.Itoa(int(v))
		case uint32:
			converted = strconv.Itoa(int(v))
		case uint64:
			converted = strconv.Itoa(int(v))
		case int:
			converted = strconv.Itoa(int(v))
		case int8:
			converted = strconv.Itoa(int(v))
		case int16:
			converted = strconv.Itoa(int(v))
		case int32:
			converted = strconv.Itoa(int(v))
		case int64:
			converted = strconv.Itoa(int(v))
		}

		data = append(data, converted)
	}

	return data
}
