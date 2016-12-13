package slices_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/utilgo/slices"
	"github.com/stretchr/testify/assert"
)

func TestStringsFound(t *testing.T) {
	assert := assert.New(t)

	v := "foo"
	a := []string{"bar", "foo", "bazz"}

	assert.Equal(1, slices.IndexOf(v, a))
}

func TestStringsNotFound(t *testing.T) {
	assert := assert.New(t)

	v := "quux"
	a := []string{"bar", "foo", "bazz"}

	assert.Equal(-1, slices.IndexOf(v, a))
}

func TestIntsFound(t *testing.T) {
	assert := assert.New(t)

	v := 1
	a := []int{0, 1, 2}

	assert.Equal(1, slices.IndexOf(v, a))
}

func TestIntsNotFound(t *testing.T) {
	assert := assert.New(t)

	v := 100
	a := []int{0, 1, 2}

	assert.Equal(-1, slices.IndexOf(v, a))
}

func TestIntsArrayFound(t *testing.T) {
	assert := assert.New(t)

	var v uint8 = 3
	var a = [5]byte{1, 2, 3, 4, 5}

	assert.Equal(2, slices.IndexOf(v, a))
}

func ExampleIndexOf() {

	s := []string{"a1", "a2", "aa"}

	{
		v := "aa"
		i := slices.IndexOf(v, s)

		fmt.Println(i)
	}

	{
		v := "bb"
		i := slices.IndexOf(v, s)

		fmt.Println(i)
	}

	// Output:
	// 2
	// -1
}
