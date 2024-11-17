package SimpleGenerator

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
)

////////////////////////////////////////////////////////////////

func (gen *GeneratorObj) NewTypeImport(path, name string) *GeneratorUserTypeObj {
	t := gen.newType()

	t.types = "importType"
	t.name = name
	t.pathImport = path

	_, ok := gen.imports[path]
	if !ok {
		gen.imports[path] = filepath.Base(path)
	}

	gen.types[t.id()] = t
	return t
}

func (gen *GeneratorObj) NewType(name string) *GeneratorUserTypeObj {
	t := gen.newType()
	t.name = name
	t.types = "GlobalTypes"

	_, ok := gen.types[t.id()]
	if !ok {
		gen.types[t.id()] = t
	}

	return t
}

func (gen *GeneratorObj) AddType(name string, i interface{}) *GeneratorUserTypeObj {
	if value, ok := i.(*GeneratorUserTypeObj); ok {
		gen.catchError(errors.New("invalid type: *GeneratorUserTypeObj " + name))
		return value
	}

	t := gen.newType()
	t.name = name
	t.types = reflect.TypeOf(i).String()

	_, ok := gen.types[t.id()]
	if !ok {
		gen.PrintLN(fmt.Sprintf("type %s %s", t.Name(), t.types)).LN()
		gen.types[t.id()] = t
	} else {
		gen.catchError(errors.New("add type: type already exists " + name))
	}

	return t
}
