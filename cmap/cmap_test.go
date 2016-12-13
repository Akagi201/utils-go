package cmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Akagi201/utilgo/cmap"
)

func TestEmptyCmap(t *testing.T) {
	cm := cmap.New()

	assert.Zero(t, cm.Len(), "New cmap Len should be 0")
	assert.Zero(t, len(cm.Keys()), "New cmap Keys Len should be 0")
	assert.Zero(t, len(cm.Values()), "New cmap Keys Len should be 0")
	assert.True(t, cm.IsEmpty(), "New cmap should be empty")
}

func TestCmapOps(t *testing.T) {
	cm := cmap.New()

	assert.True(t, cm.Set("hello", "world"), "Set new key should return true")
	assert.True(t, cm.Has("hello"), "Set should Has hello")

	val, ok := cm.Get("hello")
	assert.True(t, ok, "Get should be true")
	assert.Equal(t, "world", val.(string), "value should be world")

	cm.Delete("hello")
	assert.False(t, cm.Has("hello"), "hello should be deleted")
}

func BenchmarkSet(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Set(i, i)
	}
}

func BenchmarkGet(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Get(i)
	}
}

func BenchmarkDelete(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Delete(i)
	}
}

func BenchmarkHas(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Has(i)
	}
}

func BenchmarkLen(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Len()
	}
}

func BenchmarkClear(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Clear()
	}
}

func BenchmarkKeys(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Keys()
	}
}

func BenchmarkValues(b *testing.B) {
	cm := cmap.New()
	for i := 0; i < b.N; i++ {
		cm.Values()
	}
}
