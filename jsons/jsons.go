// Package jsons contains some json related helper functions.
package jsons

import (
	"bytes"
	"encoding/json"
)

// JSONPrettyPrint pretty print raw json string to indent string
func JSONPrettyPrint(in, prefix, indent string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), prefix, indent)
	if err != nil {
		return in
	}
	return out.String()
}

// PrettyPrintMap pretty print a map to indent json string
func PrettyPrintMap(m map[string]interface{}) string {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}
