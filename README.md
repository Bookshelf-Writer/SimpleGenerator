![Fork GitHub Release](https://img.shields.io/github/v/release/Bookshelf-Writer/SimpleGenerator)
![Tests](https://github.com/Bookshelf-Writer/SimpleGenerator/actions/workflows/go-test.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/Bookshelf-Writer/SimpleGenerator)](https://goreportcard.com/report/github.com/Bookshelf-Writer/SimpleGenerator)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/Bookshelf-Writer/SimpleGenerator?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Bookshelf-Writer/SimpleGenerator?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/Bookshelf-Writer/SimpleGenerator)

# SimpleGenerator
Simple Go-file generator. No syntax checking. Autostyling.

## INSTALL

```bash
go get -u github.com/Bookshelf-Writer/SimpleGenerator
```

---

## INITIALIZATION

Methods:
- [generator.go](generator.go)

### NewFile(packageName)

```
var obj *GeneratorObj = NewFile("package_name")
```

Initialization by package name. The default directory is the launch point.

### NewFilePath(packagePath)

```
var obj *GeneratorObj = NewFilePath("package/path/name")
```

Initialization along the path to the directory. The name of the last directory is the name of the package.

### NewFilePathName(packagePath, packageName)

```
var obj *GeneratorObj = NewFilePathName("package_name", "package/path/name")
```

Initialization along the directory path indicating a custom package name.

---

## IMPORT

Methods:
- [generator.go](generator.go)
- [types.go](types.go)

### NewImport(path, alias)

```
var path string = obj.NewImport("github.com/dave/jennifer/jen", "")
```

Adding a new import with the ability to specify an alias.
The `_` and `.` aliases are not currently supported.

### NewTypeImport(path, name)

```
var JenFileType *GeneratorUserTypeObj = obj.NewTypeImport("github.com/dave/jennifer/jen", "File")
```

Initializing a specific type from an import.
If a package with a similar path has not been initialized before, initialization will occur automatically.

### NewType(name)

```
var GlobalTypeNameType *GeneratorUserTypeObj = obj.NewType("GlobalTypeName")
```

Initialize a third party type visible within the package.

### AddType(name, interface)

```
var NewLocalType *GeneratorUserTypeObj = obj.AddType("NewLocal", byte(0))
```

Initialize a local type by specifying the base type.
Adds this type to the generation if this type has not been initialized previously.

### AddObjStruct(interface)

```
obj.AddObjStruct(MyObj)
```

Recursively forms a structure based on the passed object.
Does not support any modifiers, sorting the output alphabetically. 

### AddObjValue(name, interface)

```
obj.AddObjValue("ValueNameObj", ValueNameObj)
```

It recursively goes through the structure and forms a data object. 
All basic types and nesting of non-standard ones are supported.

---

## EXAMPLE

##### Generator:
```go

var obj = SimpleGenerator.NewFilePathName("example", "main")

func main() {
	obj.NewImport("github.com/dave/jennifer/jen", "exAlias")

	jenFile := obj.NewTypeImport("github.com/dave/jennifer/jen", "File")
	timeDuration := obj.NewTypeImport("time", "Duration")

	//

	obj.Comment("Initializing simple constants")
	obj.AddConst(map[string]SimpleGenerator.GeneratorValueObj{
		"testConst1": {Val: 1234, Types: timeDuration, Comment: jenFile.Name()},
		"testConst2": {Val: "text", Types: obj.TypeString(), Format: ""},
	})

	obj.Comment("Initializing structure")
	testStruct := obj.AddStruct("testStruct", map[string]SimpleGenerator.GeneratorTypeObj{
		"ttt1": {Types: jenFile, IsLink: true},
		"ttt2": {Types: timeDuration, IsArray: 4},
	})

	obj.Comment("Creating a Function")
	obj.AddFunc(
		"TestFunc",
		map[string]SimpleGenerator.GeneratorTypeObj{
			"name": {Types: obj.TypeString()},
		},
		map[string]SimpleGenerator.GeneratorTypeObj{
			"err": {Types: obj.TypeError()},
		},
		testStruct,
		func(gen *SimpleGenerator.GeneratorObj) {
			gen.PrintLN("if name == testConst2 {")
			gen.Offset(1).WriteString("err = parent.ttt1.Save(name)").LN()
			gen.PrintLN("}")
		},
		func(gen *SimpleGenerator.GeneratorObj) {
			gen.LN().AddConst(map[string]SimpleGenerator.GeneratorValueObj{
				"funcConst": {Val: "text", Types: gen.TypeString(), Format: ""},
			})
		},
	)

	obj.SeparatorX8().LN()

	obj.Comment("Initializing a map")
	obj.AddMap("exList", obj.TypeString(), obj.TypeBool(), map[SimpleGenerator.GeneratorValueObj]SimpleGenerator.GeneratorValueObj{
		{Val: "1", Format: ""}: {Val: true},
		{Val: "3", Format: ""}: {Val: true},
		{Val: "2", Format: ""}: {Val: false, Comment: "hello"},
	})

	obj.SeparatorX8().LN()

	obj.Comment("Creating an ENUM-like structure")
	obj.ConstructEnum("User", "Users", byte(0), map[string]SimpleGenerator.GeneratorValueObj{
		"Def":    {Val: 0},
		"Su":     {Val: 3, Comment: "superuser"},
		"Normal": {Val: 1},
		"Admin":  {Val: 2},
	})

	obj.SeparatorX8().LN()

	obj.AddValue(map[string]SimpleGenerator.GeneratorValueObj{
		"Ok":    {Val: "1", Types: obj.TypeString(), Format: ""},
		"NotOk": {Val: 1, Types: obj.TypeString(), Format: ""},
	})

	//

	for p, err := range obj.Errors() {
		fmt.Println(p, "\t", err)
	}
	fmt.Println(obj.Save("GENERATE.go"))
}
```

##### Generated code:
```go
/**  This file is automatically generated  **/
/**  File generator: main.go:85  **/

package main

import (
	exAlias "github.com/dave/jennifer/jen"
	"time"
)

////////////////////////////////////////////////////////////////

// Initializing simple constants
const (
	testConst1 time.Duration = 1234 // exAlias.File
	testConst2 string        = "text"
)

// Initializing structure
type testStructObj struct {
	ttt1 *exAlias.File
	ttt2 [4]time.Duration
}

// Creating a Function
func (parent *testStructObj) TestFunc(name string) (err error) {
	if name == testConst2 {
		err = parent.ttt1.Save(name)
	}

	const funcConst string = "text"

	return
}

////////////////////////////////

// Initializing a map
var exListMap = map[string]bool{
	"1": true,
	"2": false, // hello
	"3": true,
}

////////////////////////////////

// Creating an ENUM-like structure
type UsersType uint8

const (
	UserDef    UsersType = 0
	UserNormal UsersType = 1
	UserAdmin  UsersType = 2
	UserSu     UsersType = 3 // superuser
)

////////////////////////////////

var (
	NotOk string = "%!s(int=1)"
	Ok    string = "1"
)
```

---

---

### Mirrors

- https://git.bookshelf-writer.fun/Bookshelf-Writer/SimpleGenerator