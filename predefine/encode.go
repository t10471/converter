package predefine

import (
	"encoding/binary"
)

var LittleEndian littleEndian

type littleEndian struct{}

var BigEndian bigEndian

type bigEndian struct{}

func int2Uint(i int) (ui uint64) {
	if 0 < i {
		ui = uint64(i)
	} else {
		ui = ^uint64(-i) + 1
	}
	return ui
}

func reverse(a []byte) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// Uint2bytes converts uint64 to []byte
func (bigEndian) Uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, i)
	return bytes[8-size : 8]
}

// Int2bytes converts int to []byte
func (bigEndian) Int2bytes(i int, size int) []byte {
	return BigEndian.Uint2bytes(int2Uint(i), size)
}

// Bytes2uint converts []byte to uint64
func (bigEndian) Bytes2uint(bytes []byte) uint64 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint64(append(padding, bytes...))
	return i
}

// Bytes2int converts []byte to int64
func (bigEndian) Bytes2int(bytes []byte) int64 {
	if 0x7f < bytes[0] {
		mask := uint64(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := BigEndian.Bytes2uint(bytes)
		i = (^i + 1) & mask
		return int64(-i)

	} else {
		i := BigEndian.Bytes2uint(bytes)
		return int64(i)
	}
}

// Uint2bytes converts uint64 to []byte
func (littleEndian) Uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, i)
	return bytes[0:size]
}

// Int2bytes converts int to []byte
func (littleEndian) Int2bytes(i int, size int) []byte {
	return LittleEndian.Uint2bytes(int2Uint(i), size)
}

// Bytes2uint converts []byte to uint64
func (littleEndian) Bytes2uint(bytes []byte) uint64 {
	reverse(bytes)
	return BigEndian.Bytes2uint(bytes)
}

// Bytes2int converts []byte to int64
func (littleEndian) Bytes2int(bytes []byte) int64 {
	reverse(bytes)
	return BigEndian.Bytes2int(bytes)
}

// EndianWrapper CIC is BigEndian (maybe...).
// just to be sure, for test endian wrapper
func EndianWrapper(b []byte) []byte {
	// if little endian
	//reverse(b)
	return b
}
