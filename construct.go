package SimpleGenerator

import (
	"errors"
	"reflect"
	"sort"
	"strings"
)

////////////////////////////////////////////////////////////////

func (gen *GeneratorObj) SeparatorX1() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 2))
	return gen
}

func (gen *GeneratorObj) SeparatorX2() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 4))
	return gen
}

func (gen *GeneratorObj) SeparatorX3() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 6))
	return gen
}

func (gen *GeneratorObj) SeparatorX4() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 8))
	return gen
}

func (gen *GeneratorObj) SeparatorX5() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 10))
	return gen
}

func (gen *GeneratorObj) SeparatorX6() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 12))
	return gen
}

func (gen *GeneratorObj) SeparatorX7() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 14))
	return gen
}

func (gen *GeneratorObj) SeparatorX8() *GeneratorObj {
	gen.PrintLN(strings.Repeat("//", 16))
	return gen
}

////

func (gen *GeneratorObj) ConstructEnum(
	name string,
	typeName string,
	typeInterface interface{},
	values map[string]GeneratorValueObj,
) *GeneratorObj {

	if name == "" || typeName == "" || typeInterface == nil {
		gen.catchError(errors.New("construct const: invalid type"))
		return gen
	}

	//

	types := gen.newType()
	types.name = name
	types.types = reflect.TypeOf(typeInterface).String()
	_, ok := gen.types[types.id()]
	if ok {
		gen.catchError(errors.New("construct const: type already exists"))
		return gen
	}

	types = gen.AddType(typeName, typeInterface)

	//

	gen.PrintLN("const (")

	bufID := []string{}
	bufMap := make(map[string]string)
	for n, v := range values {
		bufID = append(bufID, v.String()+":"+n)
		bufMap[v.String()+":"+n] = n
	}
	sort.Strings(bufID)

	//

	for _, id := range bufID {
		nameVal := bufMap[id]
		val := values[nameVal]

		gen.Offset(1).
			Sprintf("%s%s %s = %s", name, nameVal, types.Name(), val.String())

		if val.Comment != "" {
			gen.WriteString("\t// " + val.Comment)
		}

		gen.LN()
	}

	gen.PrintLN(")").LN()
	return gen
}
