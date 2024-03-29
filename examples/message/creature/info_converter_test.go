// Code generated by "encoder -type=Info"; DO NOT EDIT.
package creature

import (
	"bytes"
	"reflect"
	"testing"
)

func TestInfoConverter_MarshalBuffer(t *testing.T) {
	type fields struct {
		Type   [1]byte
		Detail [1]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    *bytes.Buffer
		wantErr bool
	}{
		{
			"OK",
			fields{[1]byte{0x01}, [1]byte{0x00}},
			func() *bytes.Buffer {
				buffer := new(bytes.Buffer)
				buffer.Write([]byte{0x01, 0x00})
				return buffer
			}(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &InfoConverter{
				Type:   tt.fields.Type,
				Detail: tt.fields.Detail,
			}
			got, err := cv.MarshalBuffer()
			if (err != nil) != tt.wantErr {
				t.Errorf("InfoConverter.MarshalBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfoConverter.MarshalBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoConverter_UnmarshalBuffer(t *testing.T) {
	type fields struct {
		Type   [1]byte
		Detail [1]byte
	}
	type args struct {
		buffer *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"OK",
			fields{[1]byte{0x01}, [1]byte{0x00}},
			func() args {
				buffer := new(bytes.Buffer)
				buffer.Write([]byte{0x01, 0x00})
				return args{buffer}
			}(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &InfoConverter{
				Type:   tt.fields.Type,
				Detail: tt.fields.Detail,
			}
			if err := cv.UnmarshalBuffer(tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("InfoConverter.UnmarshalBuffer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInfoConverter_ConvertFrom(t *testing.T) {
	type fields struct {
		Type   [1]byte
		Detail [1]byte
	}
	type args struct {
		original Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"OK",
			fields{[1]byte{0x01}, [1]byte{0x00}},
			args{Info{TypePlantae, DetailPlantaeRubiales}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &InfoConverter{
				Type:   tt.fields.Type,
				Detail: tt.fields.Detail,
			}
			if err := cv.ConvertFrom(tt.args.original); (err != nil) != tt.wantErr {
				t.Errorf("InfoConverter.ConvertFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInfoConverter_ToOriginal(t *testing.T) {
	type fields struct {
		Type   [1]byte
		Detail [1]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    Info
		wantErr bool
	}{
		{
			"OK",
			fields{[1]byte{0x01}, [1]byte{0x00}},
			Info{TypePlantae, DetailPlantaeRubiales},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &InfoConverter{
				Type:   tt.fields.Type,
				Detail: tt.fields.Detail,
			}
			got, err := cv.ToOriginal()
			if (err != nil) != tt.wantErr {
				t.Errorf("InfoConverter.ToOriginal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfoConverter.ToOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}
