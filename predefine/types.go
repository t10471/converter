package predefine

import (
	"bytes"
)

var (
	Uint2bytes = BigEndian.Uint2bytes
	Int2bytes  = BigEndian.Int2bytes
	Bytes2int  = BigEndian.Bytes2int
	Bytes2uint = BigEndian.Bytes2uint

	Int2PackedDecimal = BigEndian.Int2PackedDecimal
	PackedDecimal2Int = BigEndian.PackedDecimal2Int
)

type StructEncoder interface {
	Encode(*bytes.Buffer) error
}

type StructDecoder interface {
	Decode(*bytes.Buffer) error
}

type StructConverter interface {
	StructEncoder
	StructDecoder
}

type Blank int

type Hex uint

type PackedDecimal uint

type Ebcdic string
