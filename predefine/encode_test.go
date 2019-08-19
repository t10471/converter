package predefine

import (
	"reflect"
	"testing"
)

func Test_int2Uint(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		args   args
		wantUi uint64
	}{
		{
			"Positive number",
			args{257},
			257,
		},
		{
			"Negative number",
			args{-257},
			18446744073709551359,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUi := int2Uint(tt.args.i); gotUi != tt.wantUi {
				t.Errorf("int2Uint() = %v, want %v", gotUi, tt.wantUi)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Success reverse",
			args{[]byte{0, 1, 2}},
			[]byte{2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverse(tt.args.a); !reflect.DeepEqual(tt.args.a, tt.want) {
				t.Errorf("reverse() = %v, want %v", tt.args.a, tt.want)
			}
		})
	}
}

func Test_bigEndian_Uint2bytes(t *testing.T) {
	type args struct {
		i    uint64
		size int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Success convert Uint2bytes",
			args{258, 2},
			[]byte{0x1, 0x2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint2bytes(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.Uint2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_Int2bytes(t *testing.T) {
	type args struct {
		i    int
		size int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Success convert Int2bytes Positive Number",
			args{259, 2},
			[]byte{0x1, 0x3},
		},
		{
			"Success convert Int2bytes Negative Number",
			args{-259, 2},
			[]byte{0xfe, 0xfd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2bytes(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.Int2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_Bytes2uint(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"Success convert Bytes2uint",
			args{[]byte{0x1, 0x2}},
			258,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2uint(tt.args.bytes); got != tt.want {
				t.Errorf("bigEndian.Bytes2uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_Bytes2int(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Success convert Bytes2int Positive Number",
			args{[]byte{0x1, 0x3}},
			259,
		},
		{
			"Success convert Bytes2int Negative Number",
			args{[]byte{0xfe, 0xfd}},
			-259,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2int(tt.args.bytes); got != tt.want {
				t.Errorf("bigEndian.Bytes2int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_Uint2bytes(t *testing.T) {
	type args struct {
		i    uint64
		size int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Success convert Uint2bytes",
			args{258, 2},
			[]byte{0x2, 0x1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint2bytes(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.Uint2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_Int2bytes(t *testing.T) {
	type args struct {
		i    int
		size int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Success convert Int2bytes Positive Number",
			args{259, 2},
			[]byte{0x3, 0x1},
		},
		{
			"Success convert Int2bytes Negative Number",
			args{-259, 2},
			[]byte{0xfd, 0xfe},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2bytes(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.Int2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_Bytes2uint(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"Success convert Bytes2uint",
			args{[]byte{0x2, 0x1}},
			258,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2uint(tt.args.bytes); got != tt.want {
				t.Errorf("littleEndian.Bytes2uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_Bytes2int(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Success convert Bytes2int Positive Number",
			args{[]byte{0x3, 0x1}},
			259,
		},
		{
			"Success convert Bytes2int Negative Number",
			args{[]byte{0xfd, 0xfe}},
			-259,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2int(tt.args.bytes); got != tt.want {
				t.Errorf("littleEndian.Bytes2int() = %v, want %v", got, tt.want)
			}
		})
	}
}
