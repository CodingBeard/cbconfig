package flatten

/**
The MIT License (MIT)

Copyright (c) 2016 Jeremy Wohl

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
 */

import (
	"errors"
	"reflect"
	"strconv"
)

type SeparatorStyle struct {
	Before string // Prepend to key
	Middle string // Add between keys
	After  string // Append to key
}

// Default styles
var (
	// Separate nested key components with dots, e.g. "a.b.1.c.d"
	DotStyle = SeparatorStyle{Middle: "."}

	// Separate with path-like slashes, e.g. a/b/1/c/d
	PathStyle = SeparatorStyle{Middle: "/"}

	// Separate ala Rails, e.g. "a[b][c][1][d]"
	RailsStyle = SeparatorStyle{Before: "[", After: "]"}

	// Separate with underscores, e.g. "a_b_1_c_d"
	UnderscoreStyle = SeparatorStyle{Middle: "_"}
)

// Nested input must be a map or slice
var NotValidInputError = errors.New("Not a valid input: map or slice")

// Flatten generates a flat map from a nested one.  The original may include values of type map, slice and scalar,
// but not struct.  Keys in the flat map will be a compound of descending map keys and slice iterations.
// The presentation of keys is set by style.  A prefix is joined to each key.
func Flatten(nested map[string]interface{}, prefix string, style SeparatorStyle) (map[string]interface{}, error) {
	flatmap := make(map[string]interface{})

	err := flatten(true, flatmap, nested, prefix, style)
	if err != nil {
		return nil, err
	}

	return flatmap, nil
}

func flatten(top bool, flatMap map[string]interface{}, nested interface{}, prefix string, style SeparatorStyle) error {
	assign := func(newKey string, v interface{}) error {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map, reflect.Slice:
			if err := flatten(false, flatMap, v, newKey, style); err != nil {
				return err
			}
		default:
			flatMap[newKey] = v
		}

		return nil
	}

	switch reflect.TypeOf(nested).Kind() {
	case reflect.Map:
		stringMap, ok := nested.(map[string]interface{})
		if ok {
			for k, v := range stringMap {
				newKey := enkey(top, prefix, k, style)
				e := assign(newKey, v)
				if e != nil {
					return e
				}
			}
		}
		interfaceMap, ok := nested.(map[interface{}]interface{})
		if ok {
			for k, v := range interfaceMap {
				newKey := enkey(top, prefix, k.(string), style)
				e := assign(newKey, v)
				if e != nil {
					return e
				}
			}
		}
	case reflect.Slice:
		flatMap[prefix] = nested
		for i, v := range nested.([]interface{}) {
			newKey := enkey(top, prefix, strconv.Itoa(i), style)
			e := assign(newKey, v)
			if e != nil {
				return e
			}
		}
	default:
		return NotValidInputError
	}

	return nil
}

func enkey(top bool, prefix, subkey string, style SeparatorStyle) string {
	key := prefix

	if top {
		key += subkey
	} else {
		key += style.Before + style.Middle + subkey + style.After
	}

	return key
}