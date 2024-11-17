package SimpleGenerator

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////

func (gen *GeneratorObj) Len() int {
	return gen.buf.Len()
}

func (gen *GeneratorObj) Write(data []byte) *GeneratorObj {
	gen.buf.Write(data)
	return gen
}

func (gen *GeneratorObj) WriteString(data string) *GeneratorObj {
	gen.buf.Write([]byte(data))
	return gen
}

func (gen *GeneratorObj) Del(len int) *GeneratorObj {
	gen.buf.Truncate(gen.Len() - len)
	return gen
}

//

func (gen *GeneratorObj) Repeat(chat string, size int) *GeneratorObj {
	gen.WriteString(strings.Repeat(chat, size))
	return gen
}

func (gen *GeneratorObj) Offset(size int) *GeneratorObj {
	gen.Repeat("\t", gen.nowOffset+size)
	return gen
}

func (gen *GeneratorObj) OffsetAdd() *GeneratorObj {
	gen.nowOffset++
	return gen
}

func (gen *GeneratorObj) OffsetRemove() *GeneratorObj {
	gen.nowOffset--
	if gen.nowOffset < 0 {
		gen.nowOffset = 0
	}
	return gen
}

//

func (gen *GeneratorObj) LN() *GeneratorObj {
	gen.WriteString("\n")
	return gen
}

func (gen *GeneratorObj) Print(text string) *GeneratorObj {
	gen.Offset(0).WriteString(text)
	return gen
}

func (gen *GeneratorObj) PrintLN(text string) *GeneratorObj {
	gen.Print(text).LN()
	return gen
}

////

func (gen *GeneratorObj) patchPrint(format string, ln bool, a []any) {
	switch len(a) {
	case 0:
		return

	case 1:
		gen.WriteString(fmt.Sprintf(format, a[0]))
		if ln {
			gen.WriteString(",").LN()
		}
		return
	}

	//

	if ln {
		gen.LN()
	}

	for _, v := range a {
		if ln {
			gen.Offset(1)
		}
		gen.WriteString(fmt.Sprintf(format, v)).WriteString(",")
		if ln {
			gen.LN()
		}
	}

	return
}

//

func (gen *GeneratorObj) Sprintf(format string, a ...any) *GeneratorObj {
	return gen.WriteString(fmt.Sprintf(format, a...))
}

func (gen *GeneratorObj) String(a ...any) *GeneratorObj {
	gen.patchPrint("\"%s\"", false, a)
	return gen
}

func (gen *GeneratorObj) StringLN(a ...any) *GeneratorObj {
	gen.patchPrint("\"%s\"", true, a)
	return gen
}

func (gen *GeneratorObj) Hex(a ...any) *GeneratorObj {
	gen.patchPrint("0x%02x", false, a)
	return gen
}

func (gen *GeneratorObj) HexLN(a ...any) *GeneratorObj {
	gen.patchPrint("0x%02x", true, a)
	return gen
}

func (gen *GeneratorObj) Number(a ...any) *GeneratorObj {
	gen.patchPrint("%d", false, a)
	return gen
}

func (gen *GeneratorObj) NumberLN(a ...any) *GeneratorObj {
	gen.patchPrint("%d", true, a)
	return gen
}

////

func (gen *GeneratorObj) Comment(text string) *GeneratorObj {
	bufArr := strings.Split(text, "\n")
	for _, v := range bufArr {
		gen.PrintLN("//" + v)
	}
	return gen
}

func (gen *GeneratorObj) CommentFormat(text string) *GeneratorObj {
	gen.PrintLN("/** " + text + " **/")
	return gen
}
