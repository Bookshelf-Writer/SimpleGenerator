package SimpleGenerator

import "fmt"

////////////////////////////////////////////////////////////////

type GeneratorValueObj struct {
	Val     any
	Format  any
	Comment string

	Types *GeneratorUserTypeObj
}

func (v *GeneratorValueObj) String() string {
	switch v.Format.(type) {
	case string:
		return fmt.Sprintf("\"%s\"", v.Val)

	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		return fmt.Sprintf("%d", v.Val)

	case []byte:
		return fmt.Sprintf("\"0x%02x", v.Val)
	}

	return fmt.Sprintf("%v", v.Val)
}
