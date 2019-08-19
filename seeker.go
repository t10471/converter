package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/favclip/genbase"
)

type typeSeeker struct {
	name string
	file *File
	ss   *structureSource
	err  error
}

type seekInfo struct {
	name string
	file *File
}

type ImportSeeker struct {
	ss         *structureSource
	importPkgs []*seekInfo
	err        error
}

func (t *typeSeeker) seek(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE {
		return true
	}
	found := false
	for _, spec := range decl.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		if ts.Name.String() != t.name {
			continue
		}
		t.ss, t.err = newStructureSource(ts, t.name, t.file)
		if t.err != nil {
			return true
		}
		found = true
	}
	return !found
}

func (i *ImportSeeker) addImportStr(path, ident string) {
	if strings.HasPrefix(path, `"`) && strings.HasSuffix(path, `"`) {
		path = path[1 : len(path)-1]
	}
	var ss []string
	for _, x := range i.ss.ImportDefs {
		ss = append(ss, x.Path)
	}
	if contains(ss, path) {
		return
	}
	i.ss.ImportDefs = append(i.ss.ImportDefs, &importDef{Ident: ident, Path: path})
}

func contains(ss []string, s string) bool {
	for _, a := range ss {
		if a == s {
			return true
		}
	}
	return false
}

func (i *ImportSeeker) addImports(ss *structureSource) {
	for _, f := range ss.FieldSources {
		needImport, packageIdent := genbase.IsReferenceToOtherPackage(f.Type)
		if !needImport {
			continue
		}
		importSpec := i.FindImportSpecByIdent(ss.file, packageIdent)
		if importSpec == nil {
			continue
		}
		i.addImport(importSpec)
	}
}

func (i *ImportSeeker) FindImportSpecByIdent(file *File, packageIdent string) *ast.ImportSpec {
	for _, imp := range file.file.Imports {
		if imp.Name != nil && imp.Name.Name == packageIdent {
			return imp
		} else if strings.HasSuffix(imp.Path.Value, fmt.Sprintf(`/%s"`, packageIdent)) {
			return imp
		} else if imp.Path.Value == fmt.Sprintf(`"%s"`, packageIdent) {
			return imp
		}
	}
	return nil
}

func (i *ImportSeeker) addImport(importSpec *ast.ImportSpec) {
	path := importSpec.Path.Value
	ident := ""
	if importSpec.Name != nil {
		ident = importSpec.Name.Name
	}
	i.addImportStr(path, ident)
}
