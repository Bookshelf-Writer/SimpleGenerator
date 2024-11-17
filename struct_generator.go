package SimpleGenerator

import "bytes"

////////////////////////////////////////////////////////////////

type GeneratorObj struct {
	name string
	path string

	errors []error

	//

	types   map[string]*GeneratorUserTypeObj
	imports map[string]string

	buf       bytes.Buffer
	nowOffset int
}

func newGenerator(packageName, packagePath string) *GeneratorObj {
	obj := GeneratorObj{}

	obj.name = packageName
	obj.path = packagePath

	obj.types = make(map[string]*GeneratorUserTypeObj)
	obj.imports = make(map[string]string)

	return &obj
}

//

func (gen *GeneratorObj) Errors() []error {
	return gen.errors
}
