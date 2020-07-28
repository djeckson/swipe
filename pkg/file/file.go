package file

import (
	"bytes"
	"fmt"
	"go/format"
	"os/exec"
	"path/filepath"

	"github.com/swipe-io/swipe/pkg/importer"
)

type File struct {
	bytes.Buffer
	PkgName   string
	PkgPath   string
	Version   string
	OutputDir string
	Filename  string
	Errs      []error
	Importer  *importer.Importer
}

func (f *File) frameGO() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("//+build !swipe\n\n")
	buf.WriteString("// Code generated by Swipe " + f.Version + ". DO NOT EDIT.\n\n")
	buf.WriteString("//go:generate swipe\n")
	buf.WriteString("package ")
	buf.WriteString(f.PkgName)
	buf.WriteString("\n\n")

	if f.Importer.HasImports() {
		buf.WriteString("import (\n")
		for _, imp := range f.Importer.SortedImports() {
			_, _ = fmt.Fprint(&buf, imp)
		}
		buf.WriteString(")\n\n")
	}
	buf.Write(f.Bytes())

	goSrc := buf.Bytes()
	fmtSrc, err := format.Source(goSrc)
	if err == nil {
		goSrc = fmtSrc
	}
	return goSrc, err
}

func (f *File) frameJS() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteString("// Code generated by Swipe " + f.Version + ". DO NOT EDIT.\n\n")
	buf.Write(f.Bytes())

	cmd := exec.Command("prettier", "--stdin-filepath", "prettier.js")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		defer stdin.Close()
		_, _ = stdin.Write(buf.Bytes())
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (f *File) Frame() ([]byte, error) {
	ext := filepath.Ext(f.Filename)
	switch ext {
	default:
		return f.Bytes(), nil
	case ".go":
		return f.frameGO()
	case ".js":
		return f.frameJS()

	}
}
