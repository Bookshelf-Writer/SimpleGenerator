package SimpleGenerator

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

////////////////////////////////////////////////////////////////

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

	var buf []string
	for nameVal, _ := range vals {
		buf = append(buf, nameVal)
	}
	sort.Strings(buf)

	for _, nameVal := range buf {
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

func (gen *GeneratorObj) AddValue(vals map[string]GeneratorValueObj) *GeneratorObj {
	return gen.addVals("var", vals)
}

func (gen *GeneratorObj) AddConst(vals map[string]GeneratorValueObj) *GeneratorObj {
	return gen.addVals("const", vals)
}

////

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

	var buf []string
	for nameVal, _ := range vals {
		buf = append(buf, nameVal)
	}
	sort.Strings(buf)

	for _, nameVal := range buf {
		gen.Offset(1)
		value := vals[nameVal]

		gen.WriteString(nameVal)

		if name == "struct" && value.Types != nil {
			gen.WriteString("\t" + value.String())
		}

		//

		if value.Comment != "" {
			gen.Print("\t// " + value.Comment)
		}

		gen.LN()
	}

	gen.PrintLN("}").LN()

	return t
}

//

func (gen *GeneratorObj) AddStruct(name string, vals map[string]GeneratorTypeObj) *GeneratorUserTypeObj {
	return gen.addStructs("struct", name, vals)
}

func (gen *GeneratorObj) AddInterface(name string, vals map[string]GeneratorTypeObj) *GeneratorUserTypeObj {
	return gen.addStructs("interface", name, vals)
}

////

func (gen *GeneratorObj) AddFunc(
	name string,
	input, output map[string]GeneratorTypeObj,
	parent *GeneratorUserTypeObj,
	body ...func(gen *GeneratorObj),
) {

	var bufStringArr []string
	params := func(m map[string]GeneratorTypeObj) {
		bufStringArr = []string{}
		for nameVal, val := range m {
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

////

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