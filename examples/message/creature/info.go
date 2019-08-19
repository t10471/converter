package creature

import (
	"fmt"

	"github.com/t10471/converter/predefine"
)

type Type int
type type_ struct {
	Value uint8
	Label string
}

//go:generate stringer -type=Type

const (
	TypeUnused Type = iota
	TypePlantae
	TypeAnimalia
)

var typeMap = map[Type]type_{
	TypePlantae:  {0x01, "植物"},
	TypeAnimalia: {0x02, "動物"},
}

type Detail int
type detail struct {
	Value uint8
	Label string
}
type detailMap map[Detail]detail

//go:generate stringer -type=Detail

const (
	DetailUnused Detail = iota
	DetailPlantaeRubiales
	DetailPlantaeOrchidales
	DetailAnimaliaMammals
	DetailAnimaliaAves
)

var detailPlantaeMap = detailMap{
	DetailPlantaeRubiales:   {0x00, "アカネ目"},
	DetailPlantaeOrchidales: {0x01, "ラン目"},
}

var detailAnimaliaMap = detailMap{
	DetailAnimaliaMammals: {0x10, "哺乳類"},
	DetailAnimaliaAves:    {0x11, "鳥類"},
}

//go:generate converter -type=Info

type Info struct {
	Type   Type   `label:"分類" byte:"1"`
	Detail Detail `label:"詳細分類" byte:"1" ref:"Type"`
}

func (t *Type) Encode() ([1]byte, error) {
	v, found := typeMap[*t]
	if !found {
		return [1]byte{0x00}, fmt.Errorf("invalid value %s Type Encode", t)
	}
	return [1]byte{v.Value}, nil
}

func (t *Type) Decode(b [1]byte) error {
	x := uint8(predefine.Bytes2uint(b[:]))
	for k, v := range typeMap {
		if v.Value == x {
			*t = k
			return nil
		}
	}
	return fmt.Errorf("invalid value %d Type Decode", x)
}

func (t Type) getMap() detailMap {
	switch t {
	case TypePlantae:
		return detailPlantaeMap
	case TypeAnimalia:
		return detailAnimaliaMap
	}
	return detailMap{}
}

func (d *Detail) Encode(t Type) ([1]byte, error) {
	v, found := t.getMap()[*d]
	if !found {
		return [1]byte{0x00}, fmt.Errorf("invalid value %s Type Encode", t)
	}
	return [1]byte{v.Value}, nil
}

func (d *Detail) Decode(b [1]byte, t Type) error {
	x := uint8(predefine.Bytes2uint(b[:]))
	for k, v := range t.getMap() {
		if v.Value == x {
			*d = k
			return nil
		}
	}
	return fmt.Errorf("invalid value %d Type Decode", x)
}
