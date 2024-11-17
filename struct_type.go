package SimpleGenerator

import "fmt"

////////////////////////////////////////////////////////////////

type GeneratorTypeObj struct {
	Comment string

	Types   *GeneratorUserTypeObj
	IsLink  bool
	IsArray int
}

func (s *GeneratorTypeObj) String() string {
	ret := ""

	if s.IsArray > 0 {
		if s.IsArray > 1 {
			ret += fmt.Sprintf("[%d]", s.IsArray)
		} else {
			ret += "[]"
		}
	}
	if s.IsLink {
		ret += "*"
	}

	return ret + s.Types.Name()
}
