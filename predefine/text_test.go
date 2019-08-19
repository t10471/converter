package predefine

import (
	"reflect"
	"testing"
)

func TestToUtf8(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Success Convert",
			args{string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66})},
			"照会",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUtf8(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUtf8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToJis(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Success Convert",
			args{"照会"},
			string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJis(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToWiden(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success Convert",
			args{"ﾗジｵドﾗﾏ"},
			"ラジオドラマ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToWiden(tt.args.s); got != tt.want {
				t.Errorf("ToWiden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNarrow(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success Convert",
			args{"ラジオドラマ"},
			"ﾗジｵドﾗﾏ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNarrow(tt.args.s); got != tt.want {
				t.Errorf("ToNarrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNFD(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success Convert",
			args{"ラジオドラマ"},
			"ラジオドラマ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNFD(tt.args.s); got != tt.want {
				t.Errorf("ToNFD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNFC(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success Convert",
			args{"ラジオドラマ"},
			"ラジオドラマ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNFC(tt.args.s); got != tt.want {
				t.Errorf("ToNFC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHex(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name  string
		args  args
		wantH []string
	}{
		{
			"Success Convert",
			args{string([]byte{1, 255, 15})},
			[]string{"0x01", "0xff", "0x0f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotH := ToHex(tt.args.a); !reflect.DeepEqual(gotH, tt.wantH) {
				t.Errorf("ToHex() = %v, want %v", gotH, tt.wantH)
			}
		})
	}
}

func Test_wrap(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success Wrap exist wrap",
			args{string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66})},
			string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66}),
		},
		{
			"Success Wrap exist wrap left only",
			args{string([]byte{27, 36, 66, 62, 72, 50, 113})},
			string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66}),
		},
		{
			"Success Wrap exist wrap right only",
			args{string([]byte{62, 72, 50, 113, 27, 40, 66})},
			string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66}),
		},
		{
			"Success Wrap not exist wrap",
			args{string([]byte{62, 72, 50, 113})},
			string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66}),
		},
		{
			"Success Wrap exist wrap",
			args{string([]byte{62, 72})},
			string([]byte{62, 72}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrap(tt.args.a); !reflect.DeepEqual([]byte(got), []byte(tt.want)) {
				t.Errorf("wrap() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func TestRemoveEscapeSequence(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success RemoveEscapeSequence exist wrap",
			args{string([]byte{27, 36, 66, 62, 72, 50, 113, 27, 40, 66})},
			string([]byte{62, 72, 50, 113}),
		},
		{
			"Success RemoveEscapeSequence exist wrap left only",
			args{string([]byte{27, 36, 66, 62, 72, 50, 113})},
			string([]byte{62, 72, 50, 113}),
		},
		{
			"Success RemoveEscapeSequence exist wrap right only",
			args{string([]byte{62, 72, 50, 113, 27, 40, 66})},
			string([]byte{62, 72, 50, 113}),
		},
		{
			"Success RemoveEscapeSequence not exist wrap",
			args{string([]byte{62, 72, 50, 113})},
			string([]byte{62, 72, 50, 113}),
		},
		{
			"Success RemoveEscapeSequence not exist wrap",
			args{string([]byte{62, 72})},
			string([]byte{62, 72}),
		},
		{
			"Success RemoveEscapeSequence not exist wrap",
			args{string([]byte{27, 36, 66, 62, 72})},
			string([]byte{62, 72}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveEscapeSequence(tt.args.a); !reflect.DeepEqual([]byte(got), []byte(tt.want)) {
				t.Errorf("RemoveEscapeSequence() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}
