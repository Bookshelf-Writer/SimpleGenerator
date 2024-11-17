package SimpleGenerator

////////////////////////////////////////////////////////////////

func (gen *GeneratorObj) TypeError() *GeneratorUserTypeObj {
	return gen.NewType("error")
}

func (gen *GeneratorObj) TypeString() *GeneratorUserTypeObj {
	return gen.NewType("string")
}

func (gen *GeneratorObj) TypeBool() *GeneratorUserTypeObj {
	return gen.NewType("bool")
}

func (gen *GeneratorObj) TypeUInt8() *GeneratorUserTypeObj {
	return gen.NewType("uint8")
}

func (gen *GeneratorObj) TypeUInt16() *GeneratorUserTypeObj {
	return gen.NewType("uint16")
}

func (gen *GeneratorObj) TypeUInt32() *GeneratorUserTypeObj {
	return gen.NewType("uint32")
}

func (gen *GeneratorObj) TypeUInt64() *GeneratorUserTypeObj {
	return gen.NewType("uint64")
}

func (gen *GeneratorObj) TypeInt8() *GeneratorUserTypeObj {
	return gen.NewType("int8")
}

func (gen *GeneratorObj) TypeInt16() *GeneratorUserTypeObj {
	return gen.NewType("int16")
}

func (gen *GeneratorObj) TypeInt32() *GeneratorUserTypeObj {
	return gen.NewType("int32")
}

func (gen *GeneratorObj) TypeInt64() *GeneratorUserTypeObj {
	return gen.NewType("int64")
}

func (gen *GeneratorObj) TypeFloat32() *GeneratorUserTypeObj {
	return gen.NewType("float32")
}

func (gen *GeneratorObj) TypeFloat64() *GeneratorUserTypeObj {
	return gen.NewType("float64")
}

func (gen *GeneratorObj) TypeComplex64() *GeneratorUserTypeObj {
	return gen.NewType("complex64")
}

func (gen *GeneratorObj) TypeComplex128() *GeneratorUserTypeObj {
	return gen.NewType("complex128")
}

func (gen *GeneratorObj) TypeInt() *GeneratorUserTypeObj {
	return gen.NewType("int")
}

func (gen *GeneratorObj) TypeUInt() *GeneratorUserTypeObj {
	return gen.NewType("uint")
}

func (gen *GeneratorObj) TypeUIntPtr() *GeneratorUserTypeObj {
	return gen.NewType("uintptr")
}

func (gen *GeneratorObj) TypeByte() *GeneratorUserTypeObj {
	return gen.NewType("byte")
}

func (gen *GeneratorObj) TypeRune() *GeneratorUserTypeObj {
	return gen.NewType("rune")
}

func (gen *GeneratorObj) TypeAny() *GeneratorUserTypeObj {
	return gen.NewType("any")
}

func (gen *GeneratorObj) TypeComparable() *GeneratorUserTypeObj {
	return gen.NewType("comparable")
}

////

func (gen *GeneratorObj) TypeTimeTime() *GeneratorUserTypeObj {
	return gen.NewTypeImport("time", "Time")
}

func (gen *GeneratorObj) TypeTimeTimer() *GeneratorUserTypeObj {
	return gen.NewTypeImport("time", "Timer")
}

func (gen *GeneratorObj) TypeTimeTicker() *GeneratorUserTypeObj {
	return gen.NewTypeImport("time", "Ticker")
}

func (gen *GeneratorObj) TypeTimeLocation() *GeneratorUserTypeObj {
	return gen.NewTypeImport("time", "Location")
}

func (gen *GeneratorObj) TypeTimeDuration() *GeneratorUserTypeObj {
	return gen.NewTypeImport("time", "Duration")
}
