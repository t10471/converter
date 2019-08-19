package predefine

import (
	"reflect"
	"testing"
)

func Test_bigEndian_Int2PackedDecimal(t *testing.T) {
	type args struct {
		i    int
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"Success zero",
			args{0, 4},
			[]byte{0x00, 0x00, 0x00, 0x0c},
			false,
		},
		{
			"Success positive fit",
			args{999, 2},
			[]byte{0x99, 0x9c},
			false,
		},
		{
			"Success positive more",
			args{999, 3},
			[]byte{0x00, 0x99, 0x9c},
			false,
		},
		{
			"Success cpu inquiry max positive",
			args{999999999, 5},
			[]byte{0x99, 0x99, 0x99, 0x99, 0x9c},
			false,
		},
		{
			"Success negative fit",
			args{-123, 2},
			[]byte{0x12, 0x3d},
			false,
		},
		{
			"Success negative more",
			args{-1234, 3},
			[]byte{0x01, 0x23, 0x4d},
			false,
		},
		{
			"Success cpu inquiry max negative",
			args{-999999999, 5},
			[]byte{0x99, 0x99, 0x99, 0x99, 0x9d},
			false,
		},
		{
			"Failed",
			args{-1234, 2},
			[]byte{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int2PackedDecimal(tt.args.i, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("bigEndian.Int2PackedDecimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.Int2PackedDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_PackedDecimal2Int(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want int
	}{
		{
			"Success positive fit",
			[]byte{0x99, 0x9c},
			999,
		},
		{
			"Success positive more",
			[]byte{0x01, 0x23, 0x4c},
			1234,
		},
		{
			"Success negative fit",
			[]byte{0x99, 0x9d},
			-999,
		},
		{
			"Success negative more",
			[]byte{0x01, 0x23, 0x4d},
			-1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PackedDecimal2Int(tt.args); got != tt.want {
				t.Errorf("bigEndian.PackedDecimal2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_Int2PackedDecimal(t *testing.T) {
	type args struct {
		i    int
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"Success positive fit",
			args{999, 2},
			[]byte{0x9c, 0x99},
			false,
		},
		{
			"Success positive more",
			args{999, 3},
			[]byte{0x9c, 0x99, 0x00},
			false,
		},
		{
			"Success negative fit",
			args{-123, 2},
			[]byte{0x3d, 0x12},
			false,
		},
		{
			"Success negative more",
			args{-1234, 3},
			[]byte{0x4d, 0x23, 0x01},
			false,
		},
		{
			"Failed",
			args{-1234, 2},
			[]byte{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int2PackedDecimal(tt.args.i, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("littleEndian.Int2PackedDecimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.Int2PackedDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_PackedDecimal2Int(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want int
	}{
		{
			"Success positive fit",
			[]byte{0x9c, 0x99},
			999,
		},
		{
			"Success positive more",
			[]byte{0x4c, 0x23, 0x01},
			1234,
		},
		{
			"Success negative fit",
			[]byte{0x9d, 0x99},
			-999,
		},
		{
			"Success negative more",
			[]byte{0x4d, 0x23, 0x01},
			-1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PackedDecimal2Int(tt.args); got != tt.want {
				t.Errorf("littleEndian.PackedDecimal2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countDigits(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			"Success count 0",
			args{0},
			1,
		},
		{
			"Success count 1",
			args{1},
			1,
		},
		{
			"Success count 10",
			args{10},
			2,
		},
		{
			"Success count -10",
			args{-10},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countDigits(tt.args.i); gotCount != tt.wantCount {
				t.Errorf("countDigits() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
