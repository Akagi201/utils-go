// Package slices contains some slice related helper functions.
package slices

import (
	"reflect"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
)

var (
	// ErrNotSlice happens when the value passed is not a slice.
	ErrNotSlice = errors.New("not slice")

	// ErrNotString happens when the value of the field is not a string.
	ErrNotString = errors.New("not string")

	// ErrNotInt happens when the value of the field is not an int.
	ErrNotInt = errors.New("not int")

	// ErrNotFloat happens when the value of the field is not a float64.
	ErrNotFloat = errors.New("not float64")

	// ErrNotBool happens when the value of the field is not a bool.
	ErrNotBool = errors.New("not bool")
)

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// The second argument must be a slice or array.
func IndexOf(val any, slice any) int {
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

// ToStrings maps a field to a slice of string.
func ToStrings(slice any, fieldName string) (s []string, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case string:
						s = append(s, e.(string))
					default:
						return nil, ErrNotString
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// ToInts maps a field to a slice of int.
func ToInts(slice any, fieldName string) (s []int, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case int:
						s = append(s, e.(int))
					default:
						return nil, ErrNotInt
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// ToFloats maps a field to a slice of int.
func ToFloats(slice any, fieldName string) (s []float64, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case float64:
						s = append(s, e.(float64))
					default:
						return nil, ErrNotFloat
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// ToBools maps a field to a slice of bool.
func ToBools(slice any, fieldName string) (s []bool, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case bool:
						s = append(s, e.(bool))
					default:
						return nil, ErrNotBool
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// ToStringsUnsafe maps a field to a slice of string but not returns an error.
// If an error still happens, s will be nil.
func ToStringsUnsafe(slice any, fieldName string) []string {
	s, _ := ToStrings(slice, fieldName)
	return s
}

// ToIntsUnsafe maps a field to a slice of int but not returns an error.
// If an error still happens, s will be nil.
func ToIntsUnsafe(slice any, fieldName string) []int {
	s, _ := ToInts(slice, fieldName)
	return s
}

// ToFloatsUnsafe maps a field to a slice of float but not returns an error.
// If an error still happens, s will be nil.
func ToFloatsUnsafe(slice any, fieldName string) []float64 {
	s, _ := ToFloats(slice, fieldName)
	return s
}

// ToBoolsUnsafe maps a field to a slice of bool but not returns an error.
// If an error still happens, s will be nil.
func ToBoolsUnsafe(slice any, fieldName string) []bool {
	s, _ := ToBools(slice, fieldName)
	return s
}

// MaxInt max int of the slice
func MaxInt(slice []int) int {
	var max int
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

// MaxFloat max float of the slice
func MaxFloat(slice []float64) float64 {
	var max float64
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

// MinInt min int of the slice
func MinInt(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	min := 9223372036854775807
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min
}

// MinFloat min float of the slice
func MinFloat(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}
	var min float64 = 9223372036854775807
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min
}

func SumInt(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumFloat(slice []float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	return sum
}
