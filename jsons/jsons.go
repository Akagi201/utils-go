// Package jsons contains some json related helper functions.
package jsons

import (
	"bytes"
	"encoding/json"
)

// JSONPrettyPrint pretty print raw json string to indent string
func JSONPrettyPrint(in, prefix, indent string) string {
	var out bytes.Buffer
	if err := json.Indent(&out, []byte(in), prefix, indent); err != nil {
		return in
	}
	return out.String()
}

// CompactJSON compact json input with insignificant space characters elided
func CompactJSON(in string) string {
	var out bytes.Buffer
	if err := json.Compact(&out, []byte(in)); err != nil {
		return in
	}
	return out.String()
}

// PrettyPrintMap pretty print a map to indent json string
func PrettyPrintMap(m map[string]any) string {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}
