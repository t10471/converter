package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

var (
	ErrNotStructType = errors.New("type is not ast.StructType")
	loadFlag         = packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo
)

type File struct {
	pkg  *Package
	file *ast.File
}

type Package struct {
	Name  string
	typs  map[ast.Expr]types.TypeAndValue
	defs  map[*ast.Ident]types.Object
	files []*File
}

func newPackage(patterns []string) (*Package, error) {
	cfg := &packages.Config{
		Mode:  loadFlag,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("error: %d packages found", len(pkgs))
	}
	pkg := pkgs[0]
	p := &Package{
		Name:  pkg.Name,
		typs:  pkg.TypesInfo.Types,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}
	for i, file := range pkg.Syntax {
		p.files[i] = &File{
			file: file,
			pkg:  p,
		}
	}
	return p, nil

}

func (p *Package) makeStructureSource(typeName string) (*structureSource, error) {
	ss, err := p.seek(typeName)
	if err != nil {
		return nil, err
	}
	ss.TypeName = typeName
	ss.Package = p
	is := &ImportSeeker{ss, nil, nil}
	is.addImportStr("bytes", "")
	is.addImportStr("encoding/binary", "")
	is.addImportStr("github.com/t10471/converter/predefine", "")
	is.addImportStr("github.com/Intermernet/ebcdic", "")
	is.addImports(ss)
	return ss, nil
}

func (p *Package) seek(typeName string) (*structureSource, error) {
	ts := &typeSeeker{typeName, nil, nil, nil}
	for _, file := range p.files {
		if file.file == nil {
			continue
		}
		ts.file = file
		ast.Inspect(file.file, ts.seek)
		if ts.err != nil || ts.ss != nil {
			break
		}
	}
	return ts.ss, ts.err
}
