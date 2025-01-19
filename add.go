package SimpleGenerator

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// //////////////////////////////////////////////////////////////

func (gen *GeneratorObj) addVals(name string, vals map[string]GeneratorValueObj) *GeneratorObj {
	l := len(vals)

	if name != "var" && name != "const" {
		gen.catchError(errors.New("invalid value name " + name))
		return gen
	}

	switch l {
	case 0:
		gen.catchError(errors.New("value name is empty"))
		return gen

	case 1:
		gen.Print(name + " ")
	}

	//

	if l > 1 {
		gen.PrintLN(name + " (")
	}

	for _, nameVal := range sortMapKey(vals) {
		if l > 1 {
			gen.Offset(1)
		}
		value := vals[nameVal]

		gen.WriteString(nameVal + " " + value.Types.Name() + " = ")

		//

		gen.WriteString(value.String())

		if value.Comment != "" {
			gen.Print("\t// " + value.Comment)
		}

		gen.LN()
	}

	if l > 1 {
		gen.PrintLN(")").LN()
	}

	return gen
}

//

// AddValue Adding an Array of Variables
func (gen *GeneratorObj) AddValue(vals map[string]GeneratorValueObj) *GeneratorObj {
	return gen.addVals("var", vals)
}

// AddConst Adding an Array of Constants
func (gen *GeneratorObj) AddConst(vals map[string]GeneratorValueObj) *GeneratorObj {
	return gen.addVals("const", vals)
}

// //

func (gen *GeneratorObj) addStructs(name, fName string, vals map[string]GeneratorTypeObj) *GeneratorUserTypeObj {

	if name != "struct" && name != "interface" {
		gen.catchError(errors.New("invalid value name " + name))
		return nil
	}

	//

	var t *GeneratorUserTypeObj
	switch name {
	case "struct":
		t = gen.NewType(fName + "Obj")
	case "interface":
		t = gen.NewType(fName + "Interface")
	}

	gen.PrintLN("type " + t.Name() + " " + name + " {")

	for _, nameVal := range sortMapKey(vals) {
		gen.Offset(1)
		value := vals[nameVal]

		gen.WriteString(nameVal)

		if name == "struct" && value.Types != nil {
			gen.WriteString("\t" + value.String())
		}

		//

		if len(value.Tags) > 0 {
			gen.Print("\t`")
			for _, tagName := range sortMapKey(value.Tags) {
				gen.Print(tagName + ":\"" + value.Tags[tagName] + "\",")
			}
			gen.Print("`")
		}

		if value.Comment != "" {
			gen.Print("\t// " + value.Comment)
		}

		gen.LN()
	}

	gen.PrintLN("}").LN()

	return t
}

//

// AddStruct Adding a structure description based on an array of variables
func (gen *GeneratorObj) AddStruct(name string, vals map[string]GeneratorTypeObj) *GeneratorUserTypeObj {
	return gen.addStructs("struct", name, vals)
}

// AddInterface Adding an interface description using an array of variables
func (gen *GeneratorObj) AddInterface(name string, vals map[string]GeneratorTypeObj) *GeneratorUserTypeObj {
	return gen.addStructs("interface", name, vals)
}

// //

// AddFunc Adding a Method
func (gen *GeneratorObj) AddFunc(
	name string,
	input, output map[string]GeneratorTypeObj,
	parent *GeneratorUserTypeObj,
	body ...func(gen *GeneratorObj),
) {

	var bufStringArr []string
	params := func(m map[string]GeneratorTypeObj) {
		bufStringArr = []string{}
		for _, nameVal := range sortMapKey(m) {
			val := m[nameVal]
			if val.Types == nil {
				continue
			}
			bufStringArr = append(bufStringArr, fmt.Sprintf("%s %s", nameVal, val.String()))
		}
		gen.Sprintf(" (%s)", strings.Join(bufStringArr, ", "))
	}

	//

	gen.Print("func ")
	if parent != nil && !parent.isImport() {
		gen.Sprintf("(parent *%s) ", parent.Name())
	}
	gen.WriteString(name)

	params(input)
	if len(output) > 0 {
		params(output)
	}

	gen.WriteString(" {\n")
	gen.OffsetAdd()
	for _, f := range body {
		f(gen)
	}

	if len(output) > 0 {
		gen.LN().PrintLN("return")
	}

	gen.OffsetRemove()
	gen.PrintLN("}").LN()
}

// //

// AddMap Adding a map by array of values
func (gen *GeneratorObj) AddMap(
	name string,
	key, element *GeneratorUserTypeObj,
	body map[GeneratorValueObj]GeneratorValueObj,
) *GeneratorUserTypeObj {

	t := gen.NewType(name + "Map")
	gen.Print("var ").Sprintf("%s = map[%s]%s{", t.Name(), key.Name(), element.Name()).LN()

	bufID := []string{}
	bufArr := make(map[string]string, 0)
	for nameVal, val := range body {
		if nameVal.Val == nil || val.Val == nil {
			continue
		}

		line := fmt.Sprintf("%s:\t %s,", nameVal.String(), val.String())
		if val.Comment != "" {
			line += "\t// " + val.Comment
		}

		bufID = append(bufID, nameVal.String())
		bufArr[nameVal.String()] = line
	}

	sort.Strings(bufID)
	for _, ID := range bufID {
		gen.Offset(1).PrintLN(bufArr[ID])
	}

	gen.PrintLN("}").LN()
	return t
}
