package SimpleGenerator

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// //////////////////////////////////////////////////////////////

func goLiteral(val reflect.Value, indent int) string {
	if !val.IsValid() {
		return "nil"
	}

	indentation := strings.Repeat("\t", indent)

	switch val.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)

	case reflect.Uint8:
		return fmt.Sprintf("%#02x", val.Uint())

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10)

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64)

	case reflect.String:
		return strconv.Quote(val.String())

	case reflect.Array, reflect.Slice:
		var sb strings.Builder

		if val.Kind() == reflect.Array {
			sb.WriteString(fmt.Sprintf("[%d]%s{\n", val.Len(), typeToString(val.Type().Elem())))
		} else {
			sb.WriteString(fmt.Sprintf("[]%s{\n", typeToString(val.Type().Elem())))
		}

		for i := 0; i < val.Len(); i++ {
			sb.WriteString(indentation + "\t" + goLiteral(val.Index(i), indent+1) + ",\n")
		}
		sb.WriteString(indentation + "}")
		return sb.String()

	case reflect.Map:
		var sb strings.Builder

		keyType := typeToString(val.Type().Key())
		elemType := typeToString(val.Type().Elem())

		sb.WriteString(fmt.Sprintf("map[%s]%s{\n", keyType, elemType))

		keys := val.MapKeys()
		bufValues := make([]string, len(keys))

		for _, k := range keys {
			mapKeyStr := goLiteral(k, indent+1)
			mapValStr := goLiteral(val.MapIndex(k), indent+1)
			bufValues = append(bufValues, fmt.Sprintf("%s\t%s: %s,\n", indentation, mapKeyStr, mapValStr))
		}

		sort.Strings(bufValues)
		for _, k := range bufValues {
			sb.WriteString(k)
		}

		sb.WriteString(indentation + "}")
		return sb.String()

	case reflect.Struct:
		var sb strings.Builder

		typ := val.Type()
		typeName := typ.Name()
		if typeName == "" {
			typeName = "struct"
		}
		sb.WriteString(typeName + "{\n")
		numFields := val.NumField()

		for i := 0; i < numFields; i++ {
			fieldVal := val.Field(i)
			fieldType := typ.Field(i)
			sb.WriteString(fmt.Sprintf("%s\t%s: %s,\n", indentation, fieldType.Name, goLiteral(fieldVal, indent+1)))
		}
		sb.WriteString(indentation + "}")
		return sb.String()

	case reflect.Interface, reflect.Pointer:
		if val.IsNil() {
			return "nil"
		}
		prefix := ""
		if val.Kind() == reflect.Pointer {
			prefix = "&"
		}
		return prefix + goLiteral(val.Elem(), indent)

	default:
		return fmt.Sprintf("%#v", val.Interface())
	}
}

// //

func (gen *GeneratorObj) AddObjValue(valName string, v interface{}) *GeneratorObj {
	gen.Sprintf("var %s = ", valName)
	gen.Print(goLiteral(reflect.ValueOf(v), 0))
	return gen
}
