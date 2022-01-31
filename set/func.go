package set

// New creates an empty Set of type T.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// FromSlice returns a Set of type T with all elements from
// the given slice.
func FromSlice[T comparable](a []T) Set[T] {
	s := make(Set[T], len(a))
	for i := range a {
		s[a[i]] = struct{}{}
	}
	return s
}

// FromMapKey returns a Set of type T with all keys
// from the given map.
func FromMapKey[K comparable, V any](m map[K]V) Set[K] {
	s := make(Set[K], len(m))
	for k := range m {
		s[k] = struct{}{}
	}
	return s
}

// Union returns a Set of type T with elements
// that exists in either of the given sets.
func Union[T comparable](s1, s2 Set[T]) Set[T] {
	s := make(Set[T], len(s1)+len(s2))
	for k := range s1 {
		s[k] = struct{}{}
	}
	for k := range s2 {
		s[k] = struct{}{}
	}
	return s
}

// Intersect returns a Set of type T with elements
// that exists in both of the given sets.
func Intersect[T comparable](s1, s2 Set[T]) Set[T] {
	s := make(Set[T], min(len(s1), len(s2)))
	for k := range s1 {
		if _, ok := s2[k]; ok {
			s[k] = struct{}{}
		}
	}
	return s
}

// Minus returns a Set of type T with elements
// that exists in the first but not the second given set.
func Minus[T comparable](s1, s2 Set[T]) Set[T] {
	s := make(Set[T], len(s1))
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			s[k] = struct{}{}
		}
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
