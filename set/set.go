package set

type Set[T comparable] map[T]struct{}

// Add adds an element to the set.
func (s Set[T]) Add(k T) {
	s[k] = struct{}{}
}

// Del deletes an element from the set
func (s Set[T]) Del(k T) {
	delete(s, k)
}

// Has checks if an element is in the set.
func (s Set[T]) Has(k T) bool {
	_, ok := s[k]
	return ok
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Iterate iterates the set by calling the given function
// with each element in the set.
func (s Set[T]) Iterate(f func(k T)) {
	for k := range s {
		f(k)
	}
}

// Filter filters the set by the given predicament
// and adds them to a new set for return.
func (s Set[T]) Filter(p func(k T) bool) Set[T] {
	r := make(Set[T], len(s))
	for k := range s {
		if p(k) {
			r[k] = struct{}{}
		}
	}
	return r
}

// FilterInPlace filters the set by the given predicament
// and acts in-place.
func (s Set[T]) FilterInPlace(p func(k T) bool) {
	for k := range s {
		if !p(k) {
			delete(s, k)
		}
	}
}

// Copy generates a copy of the set.
func (s Set[T]) Copy() Set[T] {
	n := make(Set[T], s.Len())
	for k := range s {
		n[k] = struct{}{}
	}
	return n
}

// Eq checks if two sets are equal.
func (s Set[T]) Eq(s2 Set[T]) bool {
	if len(s) != len(s2) {
		return false
	}
	for k := range s {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// ToSlice extracts all elements into a slice.
func (s Set[T]) ToSlice() []T {
	r := make([]T, 0, len(s))
	for k := range s {
		r = append(r, k)
	}
	return r
}

// UnionWith adds all elements in the given set
// to the current set.
func (s Set[T]) UnionWith(s2 Set[T]) {
	for k := range s2 {
		s[k] = struct{}{}
	}
}

// IntersectWith removes all elements from the current
// set if it's not in the given set.
func (s Set[T]) IntersectWith(s2 Set[T]) {
	for k := range s {
		if _, ok := s2[k]; !ok {
			delete(s, k)
		}
	}
}

// MinusWith removes all elements from the current
// set if it's also in the given set.
func (s Set[T]) MinusWith(s2 Set[T]) {
	for k := range s2 {
		delete(s, k)
	}
}
