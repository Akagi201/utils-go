package set_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/utilgo/set"
	"github.com/stretchr/testify/assert"
)

func TestEmptySet(t *testing.T) {
	s := set.New()
	assert.Zero(t, s.Len(), "New set Len should be 0")
	assert.Zero(t, len(s.List()), "New set List length should be 0")

	s.Add(1)
	s.Add(2)
	s.Clear()
	assert.True(t, s.IsEmpty(), "Clear set should be empty")
}

func TestEach(t *testing.T) {
	s := set.New()
	s.Add(1)
	s.Add(2)
	fn := func(e interface{}) {
		assert.True(t, s.Has(e))
	}
	s.Each(fn)
}

func TestConcurrentAddAndRemove(t *testing.T) {
	s := set.New()
	jobs := 2
	done := make(chan bool, jobs)
	go func() {
		s.Add(1)
		done <- true
	}()
	go func() {
		s.Add(2)
		done <- true
	}()
	for i := 0; i < jobs; i++ {
		<-done
	}
	assert.Equal(t, 2, s.Len(), "The set should have 2 items")
	assert.True(t, s.Has(1), "The set should have 1")
	assert.True(t, s.Has(2), "The set should have 2")

	go func() {
		s.Remove(1)
		done <- true
	}()
	go func() {
		s.Remove(2)
		done <- true
	}()
	for i := 0; i < jobs; i++ {
		<-done
	}
	assert.Zero(t, s.Len(), "The set should have 0 items")
	assert.False(t, s.Has(1), "The set should have 1")
	assert.False(t, s.Has(2), "The set should have 2")
}

func ExampleSet() {
	s := set.New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(2)
	fmt.Println(s.Len())
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(2))
	fmt.Println(s.Has(3))
	// Output: 2
	// true
	// false
	// true
}
