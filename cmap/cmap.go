// Package cmap Concurrent safe Map based on map and sync.RWMutex
package cmap

import (
	"sync"
)

// Cmap concurrent safe map
type Cmap struct {
	m map[interface{}]interface{}
	sync.RWMutex
}

// New new cmap
func New() *Cmap {
	return &Cmap{
		m: make(map[interface{}]interface{}),
	}
}

// Get get from cmap return k's value
func (cm *Cmap) Get(k interface{}) (interface{}, bool) {
	cm.RLock()
	defer cm.RUnlock()
	if val, ok := cm.m[k]; ok {
		return val, true
	}
	return nil, false
}

// Set maps the given key and value.
// Return false if the key is already in the map.
func (cm *Cmap) Set(k interface{}, v interface{}) bool {
	cm.Lock()
	defer cm.Unlock()
	_, ok := cm.m[k]
	cm.m[k] = v
	return !ok
}

// Delete delete k in cmap
func (cm *Cmap) Delete(k interface{}) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.m, k)
}

// Has returns true if k exists in the map
func (cm *Cmap) Has(k interface{}) bool {
	cm.RLock()
	defer cm.RUnlock()
	if _, ok := cm.m[k]; !ok {
		return false
	}
	return true
}

// Len returns the number of items in a set.
func (cm *Cmap) Len() int {
	cm.RLock()
	defer cm.RUnlock()
	return len(cm.m)
}

// IsEmpty checks for emptiness
func (cm *Cmap) IsEmpty() bool {
	return cm.Len() == 0
}

// Clear removes all items from the set
func (cm *Cmap) Clear() {
	cm.Lock()
	defer cm.Unlock()
	cm.m = make(map[interface{}]interface{})
}

// Keys return all the keys in cmap
func (cm *Cmap) Keys() []interface{} {
	cm.RLock()
	defer cm.RUnlock()
	s := make([]interface{}, cm.Len())
	for k := range cm.m {
		s = append(s, k)
	}
	return s
}

// Values return all the values in cmap
func (cm *Cmap) Values() []interface{} {
	cm.RLock()
	defer cm.RUnlock()

	s := make([]interface{}, cm.Len())
	for _, v := range cm.m {
		s = append(s, v)
	}
	return s
}
