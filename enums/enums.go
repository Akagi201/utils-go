// Package enums Go enum and its string representation
package enums

// https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ

// Enum holds the enumerables
type Enum struct {
	enums []string
}

// type Enum int

// String turns Enum to its string form
func (e Enum) String(v int) string {
	return e.enums[v]
}

// Iota converts string to enumerable, similar to Go's iota
func (e *Enum) Iota(s string) int {
	e.enums = append(e.enums, s)
	return len(e.enums) - 1
}

// Get lookup the given string for interal enumerable
func (e Enum) Get(s string) (int, bool) {
	for ii, vv := range e.enums {
		if vv == s {
			return ii, true
		}
	}
	return -1, false
}
