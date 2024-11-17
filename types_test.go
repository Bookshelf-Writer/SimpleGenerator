package SimpleGenerator

import (
	"reflect"
	"strings"
	"testing"
)

////////////////////////////////////////////////////////////////

func TestNewTypeImport(t *testing.T) {
	gen := NewFile("testPackage")
	path := "path/to/package"
	name := "TypeName"

	typ := gen.NewTypeImport(path, name)

	if typ == nil {
		t.Fatalf("NewTypeImport returned nil")
	}

	if typ.name != name {
		t.Errorf("expected name %s; got %s", name, typ.name)
	}

	if typ.types != "importType" {
		t.Errorf("expected type 'importType'; got %s", typ.types)
	}

	if alias, exists := gen.imports[path]; !exists || alias != "package" {
		t.Errorf("expected import to be registered with alias 'package'; got %s", alias)
	}

	for _, e := range gen.Errors() {
		t.Error(e)
	}
}

func TestNewType(t *testing.T) {
	gen := NewFile("testPackage")
	name := "GlobalType"

	typ := gen.NewType(name)
	if typ == nil {
		t.Fatalf("NewType returned nil")
	}

	if typ.name != name {
		t.Errorf("expected name %s; got %s", name, typ.name)
	}

	if typ.types != "GlobalTypes" {
		t.Errorf("expected type 'GlobalTypes'; got %s", typ.types)
	}

	for _, e := range gen.Errors() {
		t.Error(e)
	}
}

func TestAddType(t *testing.T) {
	gen := NewFile("testPackage")
	name := "SomeType"
	i := 42

	typ := gen.AddType(name, i)
	if typ == nil {
		t.Fatalf("AddType returned nil")
	}

	if typ.name != name {
		t.Errorf("expected name %s; got %s", name, typ.name)
	}

	if expectedType := reflect.TypeOf(i).String(); typ.types != expectedType {
		t.Errorf("expected type %s; got %s", expectedType, typ.types)
	}

	unsupportedType := &GeneratorUserTypeObj{}
	gen.AddType("Unsupported", unsupportedType)

	for _, e := range gen.Errors() {
		if len(strings.Split(e.Error(), "GeneratorUserTypeObj")) > 1 {
			continue
		}
		t.Error(e)
	}
}
