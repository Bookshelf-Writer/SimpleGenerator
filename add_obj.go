package SimpleGenerator

import "reflect"

// //////////////////////////////////////////////////////////////

func typeToString(t reflect.Type) string {
	if t == nil {
		return "interface{}"
	}

	switch t.Kind() {
	case reflect.Ptr:
		elem := t.Elem()
		if elem.Kind() == reflect.Struct {
			return "*" + elem.Name()
		}
		return "*" + typeToString(elem)
	case reflect.Slice, reflect.Array:
		return "[]" + typeToString(t.Elem())
	case reflect.Map:
		return "map[" + typeToString(t.Key()) + "]" + typeToString(t.Elem())
	case reflect.Struct:
		return t.Name()
	default:
		return t.String()
	}
}
