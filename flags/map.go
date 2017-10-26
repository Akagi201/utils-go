package flags

import (
	"fmt"
	"strings"
)

// Map represents map[string]string flag variable.
type Map map[string]string

// String returns the string representation of the map.
func (m *Map) String() string {
	return fmt.Sprintf("%v", *m)
}

// Set adds element to the map. It must be in key:value format, otherwise
// error is returned.
func (m *Map) Set(value string) error {
	if *m == nil {
		*m = make(map[string]string)
	}

	kv := strings.Split(value, ":")
	if len(kv) != 2 || len(kv[0]) == 0 || len(kv[1]) == 0 {
		return fmt.Errorf("unsupported map flag format: %q", value)
	}
	key, value := kv[0], kv[1]
	(*m)[key] = value
	return nil
}
