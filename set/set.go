// Package set Set Data Structure based on map and sync.RWMutex
package set

import "sync"

// Set safe map
type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}

// New new set
func New() *Set {
	return &Set{
		m: make(map[interface{}]struct{}),
	}
}

// Add add
func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
}

// Remove deletes the specified item from the map
func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Has looks for the existence of an item
func (s *Set) Has(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len returns the number of items in a set.
func (s *Set) Len() int {
	return len(s.List())
}

// Clear removes all items from the set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[interface{}]struct{})
}

// IsEmpty checks for emptiness
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// Each call fn on each element in the Set
func (s *Set) Each(fn func(interface{})) {
	for e := range s.m {
		fn(e)
	}
}

// List Set returns a slice of all items
func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var list []interface{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
