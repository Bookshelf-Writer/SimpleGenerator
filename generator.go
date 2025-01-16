package SimpleGenerator

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// //////////////////////////////////////////////////////////////

func NewFile(packageName string) *GeneratorObj {
	path, _ := os.Getwd()
	obj := newGenerator(packageName, path)
	return obj
}

func NewFilePath(packagePath string) *GeneratorObj {
	lastFolder := filepath.Base(packagePath)
	pathToLastFolder := filepath.Dir(packagePath)
	obj := newGenerator(lastFolder, pathToLastFolder)
	return obj
}

func NewFilePathName(packagePath, packageName string) *GeneratorObj {
	obj := newGenerator(packageName, packagePath)
	return obj
}

// //////

func (gen *GeneratorObj) Render(w io.Writer) (err error) {
	write := func(b []byte) {
		_, err = w.Write(b)
	}

	//

	write([]byte("// Code generated by '$ github.com/Bookshelf-Writer/SimpleGenerator'. DO NOT EDIT.\n"))
	write([]byte(fmt.Sprintf("\npackage %s\n\n", gen.name)))

	if len(gen.imports) > 0 {
		write([]byte("import (\n"))
		for path, alias := range gen.imports {

			if filepath.Base(path) == alias {
				write([]byte(fmt.Sprintf("\t\"%s\"\n", path)))
				continue
			}

			write([]byte(fmt.Sprintf("\t%s \"%s\"\n", alias, path)))
		}
		write([]byte(")\n\n"))
	}

	write([]byte(strings.Repeat("//", 32) + "\n\n"))

	if err != nil {
		return err
	}

	//

	write(gen.buf.Bytes())

	return
}

func (gen *GeneratorObj) Save(filename string) error {
	var buf bytes.Buffer

	if err := gen.Render(&buf); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if err = os.WriteFile(filepath.Join(gen.path, filename), formatted, 0644); err != nil {
		return err
	}
	return nil
}

// //////

func (gen *GeneratorObj) NewImport(path, alias string) string {
	_, ok := gen.imports[path]
	if !ok {
		if alias == "" {
			alias = filepath.Base(path)
		}
		gen.imports[path] = alias
	} else {
		gen.catchError(errors.New("Import already exists"))
	}

	return path
}
