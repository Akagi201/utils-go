package flags_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/utils-go/flags"
)

func TestMap(t *testing.T) {
	want := map[string]string{
		"hello":   "world",
		"goodbye": "moon",
	}

	var m flags.Map
	for wantK, wantV := range want {
		err := m.Set(fmt.Sprintf("%s:%s", wantK, wantV))
		if err != nil {
			t.Fatal(err)
		}
		gotV, ok := m[wantK]
		if !ok {
			t.Errorf("missing map entry with key %s\n", wantK)
		}
		if gotV != wantV {
			t.Errorf("expected map value %s, got %s\n", wantV, gotV)
		}
	}

	invalids := []string{"key=", "=value", "key=value=?"}
	for _, invalid := range invalids {
		if err := m.Set(invalid); err == nil {
			t.Error("expected error, got nil")
		}
	}
}
