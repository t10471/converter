package predefine

import (
	"fmt"
	"math"
)

const (
	positiveSign = 0xc
	negativeSign = 0xd
)

func (bigEndian) Int2PackedDecimal(i int, size int) ([]byte, error) {
	digits := countDigits(i)
	bl := int(math.Ceil(float64(digits)/2 + 0.5))
	if bl > size {
		return nil, fmt.Errorf("invalid size %d, int %d is required %d bytes", size, i, bl)
	}
	sign := positiveSign
	if i < 0 {
		sign = negativeSign
		i *= -1
	}
	bytes := make([]byte, size)
	scale := 1
	current := 1
	index := 0
	for digits >= current {
		var lower, upper byte
		if current == 1 {
			lower = byte(sign)
			upper = byte(i % 10)
			scale *= 10
			current++
		} else {
			lower = byte(i % (scale * 10) / scale)
			scale *= 10
			current++
			if digits >= current {
				upper = byte(i % (scale * 10) / scale)
			} else {
				upper = byte(0)
			}
			scale *= 10
			current++
		}
		bytes[size-index-1] = byte((upper << 4) | lower)
		index++
	}
	return bytes, nil
}
func (bigEndian) PackedDecimal2Int(bytes []byte) int {
	i := 0
	scale := 1
	negative := false
	byteLen := len(bytes)
	index := 0

	for byteLen > index {
		digit := bytes[byteLen-index-1]
		upper := int((digit & 0xf0) >> 4)
		lower := int(digit & 0x0f)
		if index == 0 {
			if lower == negativeSign {
				negative = true
			}
			i += upper
			scale *= 10
		} else {
			i += lower * scale
			scale *= 10
			i += upper * scale
			scale *= 10
		}
		index++
	}
	if negative {
		i *= -1
	}
	return i
}

func (littleEndian) Int2PackedDecimal(i int, size int) ([]byte, error) {
	v, err := BigEndian.Int2PackedDecimal(i, size)
	if err != nil {
		return v, err
	}
	reverse(v)
	return v, nil
}

func (littleEndian) PackedDecimal2Int(bytes []byte) int {
	reverse(bytes)
	return BigEndian.PackedDecimal2Int(bytes)
}

func countDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}
