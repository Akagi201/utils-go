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

type Person struct {
	ID    int
	Name  string
	Money float64
	Male  bool
}

type Group struct {
	People []*Person
}

func TestMapSliceToString(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}
	s, err := slices.ToStrings(g.People, "Name")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], "George")
	assert.Equal(s[1], "Jeff")
	assert.Equal(s[2], "Ted")
	assert.Equal(s[3], "Luda")

	s, err = slices.ToStrings(g.People, "ID")
	assert.Equal(slices.ErrNotString, err)
}

func TestMapSliceToInt(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}

	s, err := slices.ToInts(g.People, "ID")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], 0)
	assert.Equal(s[1], 1)
	assert.Equal(s[2], 2)
	assert.Equal(s[3], 3)

	s, err = slices.ToInts(g.People, "Money")
	assert.Equal(slices.ErrNotInt, err)
}

func TestMapSliceToFloat(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}

	s, err := slices.ToFloats(g.People, "Money")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], 42.42)
	assert.Equal(s[1], 0.0)
	assert.Equal(s[2], 50.0)
	assert.Equal(s[3], 100.0)

	s, err = slices.ToFloats(g.People, "ID")
	assert.Equal(slices.ErrNotFloat, err)
}

func TestMapSliceToBool(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}

	s, err := slices.ToBools(g.People, "Male")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], true)
	assert.Equal(s[1], true)
	assert.Equal(s[2], true)
	assert.Equal(s[3], false)

	s, err = slices.ToBools(g.People, "Name")
	assert.Equal(slices.ErrNotBool, err)
}

type Tag struct {
	Name string
}

type Post struct {
	Title string
	Tags  []*Tag
}

func Example() {
	post := &Post{
		Title: "GOLANG",
		Tags: []*Tag{
			{"Go"}, {"Golang"}, {"Gopher"},
		},
	}
	s, _ := slices.ToStrings(post.Tags, "Name")
	fmt.Println(s)

	// Output:
	// [Go Golang Gopher]
}
