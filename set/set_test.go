package set

import "testing"

func TestBasic(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Del(1)
	s.Del(2)
	if s.Has(2) || s.Len() != 0 {
		t.FailNow()
	}
}

func TestSet_Copy(t *testing.T) {
	s := New[string]()
	s.Add("hello")
	s2 := s.Copy()
	s2.Del("hello")
	if s.Len() != 1 || !s.Has("hello") ||
		s2.Len() != 0 || s2.Has("hello") {
		t.FailNow()
	}
}

func TestEqualEmpty(t *testing.T) {
	s := New[uintptr]()
	if !s.Eq(Set[uintptr](nil)) {
		t.FailNow()
	}
}

func TestNotEqual(t *testing.T) {
	s := New[uintptr]()
	s2 := New[uintptr]()
	s.Add(1)
	if s.Eq(s2) {
		t.FailNow()
	}
	s2.Add(0)
	if s.Eq(s2) {
		t.FailNow()
	}
}

func TestSlice(t *testing.T) {
	s := New[string]()
	s.Add("hello")
	s.Add("wow")
	slice := s.ToSlice()
	s2 := FromSlice(slice)
	if !s.Eq(s2) {
		t.FailNow()
	}
}
