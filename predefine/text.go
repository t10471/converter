package predefine

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

func ToUtf8(s string) (string, error) {
	u8, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(wrap(s)), japanese.ISO2022JP.NewDecoder()))
	// only memory alloc error
	if err != nil {
		return "", err
	}
	return string(u8), err
}

func ToJis(s string) (string, error) {
	jis, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(s), japanese.ISO2022JP.NewEncoder()))
	// only memory alloc error
	if err != nil {
		return "", err
	}
	return string(jis), err
}

func ToWiden(s string) string {
	return width.Widen.String(s)
}

func ToNarrow(s string) string {
	return width.Narrow.String(s)
}

func ToNFD(s string) string {
	return norm.NFD.String(s)
}

func ToNFC(s string) string {
	return norm.NFC.String(s)
}

func ToHex(a string) (h []string) {
	for _, x := range []byte(a) {
		h = append(h, fmt.Sprintf("0x%02x", x))
	}
	return
}

func wrap(a string) string {
	var w []byte
	t := []byte(a)
	l := len(t)
	if l < 3 {
		return a
	}
	// start kanji ESC $ B
	start := []byte{27, 36, 66}
	// start ASCII ESC ( B
	end := []byte{27, 40, 66}
	if !reflect.DeepEqual(t[:3], start) {
		w = append(start, t...)
	} else {
		w = t
	}
	if !reflect.DeepEqual(t[len(t)-3:], end) {
		w = append(w, end...)
	}
	return string(w)
}

func RemoveEscapeSequence(a string) string {
	var w []byte
	t := []byte(a)
	l := len(t)
	if l < 3 {
		return a
	}
	// start kanji ESC $ B
	start := []byte{27, 36, 66}
	// start ASCII ESC ( B
	end := []byte{27, 40, 66}
	if reflect.DeepEqual(t[:3], start) {
		w = t[3:]
	} else {
		w = t
	}
	l = len(w)
	if l < 3 {
		return string(w)
	}
	if reflect.DeepEqual(w[l-3:], end) {
		w = w[:l-3]
	}
	return string(w)
}
