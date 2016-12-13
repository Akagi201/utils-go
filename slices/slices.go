// Package slices contains some slice related helper functions.
package slices

import (
	"reflect"
)

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// The second argument must be a slice or array.
func IndexOf(val interface{}, slice interface{}) int {
	v := reflect.ValueOf(val)
	arr := reflect.ValueOf(slice)

	t := reflect.TypeOf(slice).Kind()

	if t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}

	return -1
}
