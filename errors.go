package SimpleGenerator

////////////////////////////////////////////////////////////////

func (obj *GeneratorObj) catchError(err error) {
	obj.errors = append(obj.errors, err)
}
