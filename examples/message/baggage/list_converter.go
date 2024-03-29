// Code generated by "encoder -type=List"; DO NOT EDIT.
package baggage

// GENERATED BY YO. DO NOT EDIT.

import (
	"bytes"
	"encoding/binary"

	"github.com/Intermernet/ebcdic"
	"github.com/t10471/converter/examples/message/creature"
	"github.com/t10471/converter/predefine"
)

type ListConverter struct {
	Name          [7]byte                        `label:"名前"`
	PlantaeCount  [2]byte                        `label:"植物数"`
	PlantaeList   [10]creature.PlantaeConverter  `label:"植物リスト"`
	AnimaliaCount [2]byte                        `label:"動物数"`
	AnimaliaList  [10]creature.AnimaliaConverter `label:"動物リスト"`
	Cost          [7]byte                        `label:"金額"`
}

func (cv *ListConverter) MarshalBuffer() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, cv); err != nil {
		return nil, err
	}
	return buffer, nil
}
func (cv *ListConverter) UnmarshalBuffer(buffer *bytes.Buffer) error {
	if err := binary.Read(buffer, binary.BigEndian, cv); err != nil {
		return err
	}
	return nil
}
func (cv *ListConverter) ConvertFrom(original List) error {
	var err error
	cv.Name = encodeListEbcdic7(original.Name)
	cv.PlantaeCount = encodeListHex2(original.PlantaeCount)
	for i, x := range original.PlantaeList {
		if i >= int(original.PlantaeCount) {
			continue
		}
		err = cv.PlantaeList[i].ConvertFrom(x)
		if err != nil {
			return err
		}
	}
	cv.AnimaliaCount = encodeListHex2(original.AnimaliaCount)
	for i, x := range original.AnimaliaList {
		if i >= int(original.AnimaliaCount) {
			continue
		}
		err = cv.AnimaliaList[i].ConvertFrom(x)
		if err != nil {
			return err
		}
	}
	cv.Cost, err = encodeListPackedDecimal7(original.Cost)
	if err != nil {
		return err
	}
	return nil
}
func (cv *ListConverter) ToOriginal() (List, error) {
	var err error
	original := List{}
	original.Name = decodeListEbcdic7(cv.Name)
	original.PlantaeCount = decodeListHex2(cv.PlantaeCount)
	original.PlantaeList = make([]creature.Plantae, int(original.PlantaeCount))
	for i, x := range cv.PlantaeList {
		if i >= int(original.PlantaeCount) {
			continue
		}
		original.PlantaeList[i], err = x.ToOriginal()
		if err != nil {
			return List{}, err
		}
	}
	original.AnimaliaCount = decodeListHex2(cv.AnimaliaCount)
	original.AnimaliaList = make([]creature.Animalia, int(original.AnimaliaCount))
	for i, x := range cv.AnimaliaList {
		if i >= int(original.AnimaliaCount) {
			continue
		}
		original.AnimaliaList[i], err = x.ToOriginal()
		if err != nil {
			return List{}, err
		}
	}
	original.Cost = decodeListPackedDecimal7(cv.Cost)
	return original, nil
}
func encodeListEbcdic7(e predefine.Ebcdic) [7]byte {
	v := ebcdic.Encode([]byte(e))
	r := [7]byte{}
	copy(r[:], v)
	return r
}
func decodeListEbcdic7(b [7]byte) predefine.Ebcdic {
	return predefine.Ebcdic(ebcdic.Decode(b[:]))
}
func encodeListHex2(h predefine.Hex) [2]byte {
	v := predefine.Uint2bytes(uint64(h), 2)
	r := [2]byte{}
	copy(r[:], v)
	return r
}
func decodeListHex2(b [2]byte) predefine.Hex {
	return predefine.Hex(predefine.Bytes2uint(b[:]))
}
func encodeListPackedDecimal7(p predefine.PackedDecimal) ([7]byte, error) {
	r := [7]byte{}
	v, err := predefine.Int2PackedDecimal(int(p), 7)
	if err != nil {
		copy(r[:], bytes.Repeat([]byte{byte(0)}, int(7)))
		return r, err
	}
	copy(r[:], v)
	return r, nil
}
func decodeListPackedDecimal7(b [7]byte) predefine.PackedDecimal {
	return predefine.PackedDecimal(predefine.PackedDecimal2Int(b[:]))
}
