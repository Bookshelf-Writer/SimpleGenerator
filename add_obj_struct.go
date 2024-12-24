package SimpleGenerator

import (
	"fmt"
	"reflect"
	"sort"
)

// //////////////////////////////////////////////////////////////

func generateDefinitions(t reflect.Type, visited map[reflect.Type]bool) {
	if t == nil {
		return
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Struct:
		if visited[t] {
			return
		}
		visited[t] = true

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			generateDefinitions(field.Type, visited)
		}
	case reflect.Slice, reflect.Array:
		generateDefinitions(t.Elem(), visited)
	case reflect.Map:
		generateDefinitions(t.Key(), visited)
		generateDefinitions(t.Elem(), visited)
	}
}

// //

func (gen *GeneratorObj) AddObjStruct(obj interface{}) *GeneratorObj {
	t := reflect.TypeOf(obj)
	visited := make(map[reflect.Type]bool)

	generateDefinitions(t, visited)

	sortArr := make([]string, 0)
	sortMap := make(map[string]reflect.Type)

	for tp := range visited {
		sortArr = append(sortArr, tp.String())
		sortMap[tp.String()] = tp
	}
	sort.Strings(sortArr)

	for _, key := range sortArr {
		gen.structDefinition(sortMap[key])
	}

	return gen
}

func (gen *GeneratorObj) structDefinition(t reflect.Type) {
	if t == nil {
		gen.LN()
		return
	}
	if t.Kind() != reflect.Struct {
		gen.LN()
		return
	}

	gen.Sprintf("type %s struct {\n", t.Name())

	bufString := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		fieldTypeStr := typeToString(fieldType)
		bufString = append(bufString, fmt.Sprintf("\t%s %s", field.Name, fieldTypeStr))
	}

	sort.Strings(bufString)
	for _, txt := range bufString {
		gen.PrintLN(txt)
	}

	gen.PrintLN("}")
}
