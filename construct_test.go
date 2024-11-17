package SimpleGenerator

import (
	"strings"
	"testing"
)

////////////////////////////////////////////////////////////////

func TestSeparatorMethods(t *testing.T) {
	gen := NewFile("testPackage")
	tests := []struct {
		method   func() *GeneratorObj
		expected string
	}{
		{gen.SeparatorX1, strings.Repeat("//", 2) + "\n"},
		{gen.SeparatorX2, strings.Repeat("//", 4) + "\n"},
		{gen.SeparatorX3, strings.Repeat("//", 6) + "\n"},
		{gen.SeparatorX4, strings.Repeat("//", 8) + "\n"},
		{gen.SeparatorX5, strings.Repeat("//", 10) + "\n"},
		{gen.SeparatorX6, strings.Repeat("//", 12) + "\n"},
		{gen.SeparatorX7, strings.Repeat("//", 14) + "\n"},
		{gen.SeparatorX8, strings.Repeat("//", 16) + "\n"},
	}

	for _, tc := range tests {
		tc.method()
	}
	for _, e := range gen.Errors() {
		t.Error(e)
	}
}

func TestConstructEnum(t *testing.T) {
	gen := NewFile("testPackage")
	values := map[string]GeneratorValueObj{
		"Pi": {Val: 3.14},
		"E":  {Val: 2.71},
	}

	gen.ConstructEnum("testPackage", "float64", float64(0), values)
	for _, e := range gen.Errors() {
		t.Error(e)
	}
}
