// Code generated by "encoder -type={{ .TypeName }}"; DO NOT EDIT.
package {{ .Package.Name }}

// GENERATED BY YO. DO NOT EDIT.

import (
{{- range .ImportDefs }}
{{- if eq .Ident "" }}
	"{{ .Path }}"
{{- else }}
	{{ .Ident }} "{{ .Path }}"
{{- end }}
{{- end }}
)

{{- $gvName := .Name }}
{{- $gvUseCase := .UseCase }}

type {{ .Name }}Converter struct {
{{- range .FieldSources }}
{{- if eq .IsSlice true }}
	{{ .Name }} [{{ .Length }}]{{ .TypeStr }}Converter `label:"{{ .Label }}"`
{{- else if eq .IsStructure true }}
	{{- if eq .IsEmbedded true }}
	{{ .Name }}Converter
	{{- else }}
	{{ .Name }} {{ .TypeStr }}Converter
	{{- end }}
{{- else }}
	{{ .Name }} [{{ .Length }}]byte `label:"{{ .Label }}"`
{{- end }}
{{- end }}
}

{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Encoder") }}
func (cv *{{ .Name }}Converter) MarshalBuffer() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, cv); err != nil {
		return nil, err
	}
	return buffer, nil
}
{{- end }}

{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Decoder") }}
func (cv *{{ .Name }}Converter) UnmarshalBuffer(buffer *bytes.Buffer) error {
	if err := binary.Read(buffer, binary.BigEndian, cv); err != nil {
		return err
	}
	return nil
}
{{- end }}

{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Encoder") }}
func (cv *{{ .Name }}Converter) ConvertFrom(original {{ .Name }}) error {
	var err error
{{- range .FieldSources }}

	{{- if eq .IsSlice true }}
	for i, x := range original.{{ .Name }} {
		{{- if eq .SliceInfo.HasSliceCount true }}
		if i >= int(original.{{ .SliceInfo.SliceCountName }}) {
			continue
		}
		{{- end}}
		err = cv.{{ .Name }}[i].ConvertFrom(x)
		if err != nil {
			return err
		}
	}
	{{- else if eq .IsStructure true }}
	{{- $n := "" }}
		{{- if eq .IsEmbedded true }}
	{{- $n = printf "%sConverter" .Name }}
		{{- else }}
	{{- $n = .Name }}
		{{- end }}
	err = cv.{{ $n }}.ConvertFrom(original.{{ .Name }})
	if err != nil {
		return err
	}
	{{- else }}
		{{- if eq .UseGenerateType true }}
			{{- if eq .BaseTypeName "PackedDecimal" }}
	cv.{{ .Name }}, err = encode{{ $gvName }}{{ .BaseTypeName }}{{ .Length }}(original.{{ .Name }})
	if err != nil {
		return err
	}
			{{- else }}
	cv.{{ .Name }} = encode{{ $gvName }}{{ .BaseTypeName }}{{ .Length }}(original.{{ .Name }})
			{{- end }}
		{{- else }}
			{{- $l := concatWithPrefix "original." .Refs }}
	cv.{{ .Name }}, err = original.{{ .Name }}.Encode({{ $l }})
	if err != nil {
		return err
	}
		{{- end }}
	{{- end }}
{{- end }}
	return nil
}
{{- end }}

{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Decoder") }}
func (cv *{{ .Name }}Converter) ToOriginal() ({{ .Name }}, error) {
	var err error
	{{- $orignalName := printf "%s" .Name }}
	original := {{ .Name }}{}
{{- range .FieldSources }}

	{{- if eq .IsSlice true }}

		{{- if eq .SliceInfo.HasSliceCount true }}
	original.{{ .Name }} = make([]{{ .TypeStr }}, int(original.{{ .SliceInfo.SliceCountName }}))
	for i, x := range cv.{{ .Name }} {
		if i >= int(original.{{ .SliceInfo.SliceCountName }}) {
			continue
		}
		{{- else}}
	original.{{ .Name }} = make([]{{ .TypeStr }}, {{ .Length }})
	for i, x := range cv.{{ .Name }} {
		{{- end}}
		original.{{ .Name }}[i], err = x.ToOriginal()
		if err != nil {
			return {{ $orignalName }}{}, err
		}
	}
	{{- else if eq .IsStructure true }}
		{{ $n := "" }}
		{{- if eq .IsEmbedded true }}
		{{ $n = printf "%sConverter" .Name }}
		{{- else }}
		{{ $n = .Name }}
		{{- end }}
	original.{{ .Name }}, err = cv.{{ $n }}.ToOriginal()
	if err != nil {
		return {{ $orignalName }}{}, err
	}
	{{- else }}
		{{- if eq .UseGenerateType true }}
	original.{{ .Name }} = decode{{ $gvName }}{{ .BaseTypeName }}{{ .Length }}(cv.{{ .Name }})
		{{- else }}
			{{- if .Refs }}
			{{- $l := concatWithPrefix "original." .Refs }}
	err = original.{{ .Name }}.Decode(cv.{{ .Name }}, {{ $l }})
			{{- else }}
	err = original.{{ .Name }}.Decode(cv.{{ .Name }})
			{{- end }}
	if err != nil {
		return {{ $orignalName }}{}, err
	}
	{{- end }}
		{{- end }}
	{{- end }}
	return original, nil
}
{{- end }}

{{- range .GenerateTypeList }}

{{- if eq .BaseTypeName "Blank" }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Encoder") }}
func encode{{ $gvName }}Blank{{ .Length }}(b predefine.Blank) [{{ .Length }}]byte {
	r := [{{ .Length }}]byte{}
	copy(r[:], bytes.Repeat([]byte{byte(0)}, int(b)))
	return r
}
{{- end }}

{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Decoder") }}
func decode{{ $gvName }}Blank{{ .Length }}(b [{{ .Length }}]byte) predefine.Blank {
	return predefine.Blank(len(b))
}
{{- end }}
{{- else if eq .BaseTypeName "Hex" }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Encoder") }}
func encode{{ $gvName }}Hex{{ .Length }}(h predefine.Hex) [{{ .Length }}]byte {
	v := predefine.Uint2bytes(uint64(h), {{ .Length }})
	r := [{{ .Length }}]byte{}
	copy(r[:], v)
	return r
}
{{- end }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Decoder") }}
func decode{{ $gvName }}Hex{{ .Length }}(b [{{ .Length }}]byte) predefine.Hex {
	return predefine.Hex(predefine.Bytes2uint(b[:]))
}
{{- end }}
{{- else if eq .BaseTypeName "PackedDecimal" }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Encoder") }}
func encode{{ $gvName }}PackedDecimal{{ .Length }}(p predefine.PackedDecimal) ([{{ .Length }}]byte, error) {
	r := [{{ .Length }}]byte{}
	v, err := predefine.Int2PackedDecimal(int(p), {{ .Length }})
	if err != nil {
		copy(r[:], bytes.Repeat([]byte{byte(0)}, int({{ .Length }})))
		return r, err
	}
	copy(r[:], v)
	return r, nil
}
{{- end }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Decoder") }}
func decode{{ $gvName }}PackedDecimal{{ .Length }}(b [{{ .Length }}]byte) predefine.PackedDecimal {
	return predefine.PackedDecimal(predefine.PackedDecimal2Int(b[:]))
}
{{- end }}

{{- else if eq .BaseTypeName "Ebcdic" }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Encoder") }}
func encode{{ $gvName }}Ebcdic{{ .Length }}(e predefine.Ebcdic) [{{ .Length }}]byte {
	v := ebcdic.Encode([]byte(e))
	r := [{{ .Length }}]byte{}
	copy(r[:], v)
	return r
}
{{- end }}
{{- if or (eq $gvUseCase "Both") (eq $gvUseCase "Decoder") }}
func decode{{ $gvName }}Ebcdic{{ .Length }}(b [{{ .Length }}]byte) predefine.Ebcdic {
	return predefine.Ebcdic(ebcdic.Decode(b[:]))
}
{{- end }}
{{- end }}
{{- end }}

