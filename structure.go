package main

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/fatih/structtag"
	"github.com/favclip/genbase"
)

type FieldSource struct {
	Name            string
	Names           []*ast.Ident
	TypeStr         string
	Type            ast.Expr
	Length          uint8
	Label           string
	IsEmbedded      bool
	IsStructure     bool
	UseGenerateType bool
	BaseTypeName    string
	Refs            []string
	IsSlice         bool
}

type importDef struct {
	Ident string
	Path  string
}

type UseCase int

//go:generate stringer -type=UseCase

const (
	Both UseCase = iota
	Encoder
	Decoder
)

type structureSource struct {
	TypeName         string
	UseCase          string
	Package          *Package
	ImportDefs       []*importDef
	Name             string
	file             *File
	ImportSpecs      []*ast.ImportSpec
	FieldSources     []*FieldSource
	GenerateTypes    generateTypes
	GenerateTypeList []generateType
}

type generateTypes mapset.Set

type generateType struct {
	BaseTypeName string
	Length       uint8
}

func getFields(ts *ast.TypeSpec) ([]*ast.Field, error) {
	structType, ok := interface{}(ts.Type).(*ast.StructType)
	if !ok {
		return nil, ErrNotStructType
	}
	var fields []*ast.Field
	for _, field := range structType.Fields.List {
		fields = append(fields, field)
	}
	return fields, nil
}

func newStructureSource(ts *ast.TypeSpec, typeName string, file *File) (*structureSource, error) {
	ss := new(structureSource)
	ss.Name = typeName
	ss.file = file
	ss.GenerateTypes = mapset.NewSet()
	fis, err := getFields(ts)
	if err != nil {
		return nil, err
	}
	for _, fi := range fis {
		f, err := ss.makeFieldSource(fi)
		if err != nil {
			return nil, err
		}
		ss.FieldSources = append(ss.FieldSources, f)
	}
	var xs []generateType
	for x := range ss.GenerateTypes.Iter() {
		xs = append(xs, x.(generateType))
	}
	ss.GenerateTypeList = xs
	return ss, nil
}

func (ss *structureSource) makeFieldSource(fi *ast.Field) (*FieldSource, error) {
	f := new(FieldSource)
	if ts, err := genbase.ExprToTypeName(fi.Type); err == nil {
		f.TypeStr = ts
	}
	if _, isSlice := fi.Type.(*ast.ArrayType); isSlice {
		f.IsSlice = true
		f.TypeStr = strings.TrimLeft(f.TypeStr, "[]")
	}

	if len(fi.Names) >= 1 {
		f.Name = fi.Names[0].String()
	}
	f.Names = fi.Names
	if f.Name == "" {
		f.IsEmbedded = true
		f.Name = f.TypeStr
	}
	f.Type = fi.Type
	if fi.Tag == nil {
		f.IsStructure = true
		return f, nil
	}
	f, err := ss.parseTag(fi, f)
	if err != nil {
		return nil, err
	}
	baseTypeName, err := ExprToBaseTypeName(fi.Type)
	if err != nil {
		return nil, err
	}
	if ss.isGenerateTarget(baseTypeName) {
		ss.GenerateTypes.Add(generateType{baseTypeName, f.Length})
		f.UseGenerateType = true
		f.BaseTypeName = baseTypeName
	}
	return f, nil
}

func (ss *structureSource) parseTag(fi *ast.Field, f *FieldSource) (*FieldSource, error) {

	s := strings.Trim(fi.Tag.Value, "`")
	tags, err := structtag.Parse(s)
	if err != nil {
		return nil, err
	}
	t, err := tags.Get("label")
	if err != nil {
		return nil, err
	}
	f.Label = t.Value()
	if f.IsSlice {
		t, err = tags.Get("length")
		if err != nil {
			return nil, err
		}
	} else {
		t, err = tags.Get("byte")
		if err != nil {
			return nil, err
		}
	}
	i, err := strconv.Atoi(t.Value())
	if err != nil {
		return nil, err
	}
	f.Length = uint8(i)
	t, err = tags.Get("ref")
	if err == nil {
		f.Refs = strings.Split(t.Value(), ",")
	}
	return f, nil
}

func ExprToBaseTypeName(expr ast.Expr) (string, error) {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name, nil
	}
	if star, ok := expr.(*ast.StarExpr); ok {
		x, err := ExprToBaseTypeName(star.X)
		if err != nil {
			return "", nil
		}
		return x, nil
	}
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		sel, err := ExprToBaseTypeName(selector.Sel)
		if err != nil {
			return "", nil
		}
		return sel, nil
	}
	if array, ok := expr.(*ast.ArrayType); ok {
		x, err := ExprToBaseTypeName(array.Elt)
		if err != nil {
			return "", nil
		}
		return x, nil
	}
	return "", fmt.Errorf("can't detect type name")
}

func (ss *structureSource) isGenerateTarget(n string) bool {
	return contains([]string{"Blank", "Ebcdic", "Hex", "PackedDecimal"}, n)
}
