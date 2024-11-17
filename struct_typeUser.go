package SimpleGenerator

import "fmt"

////////////////////////////////////////////////////////////////

type GeneratorUserTypeObj struct {
	obj *GeneratorObj

	name  string
	types string

	pathImport string
}

func (obj *GeneratorObj) newType() *GeneratorUserTypeObj {
	t := new(GeneratorUserTypeObj)
	t.obj = obj
	return t
}

func (t *GeneratorUserTypeObj) id() string {
	return t.name + t.pathImport
}

func (t *GeneratorUserTypeObj) isImport() bool {
	return t.types == "importType"
}

//

func (t *GeneratorUserTypeObj) Name() string {
	if t.isImport() {
		if t.pathImport == "" {
			return t.name
		}

		return fmt.Sprintf("%s.%s", t.obj.imports[t.pathImport], t.name)
	}

	if t.types == "GlobalTypes" {
		return t.name
	}

	return fmt.Sprintf("%sType", t.name)
}
