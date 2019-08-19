package creature

import "github.com/t10471/converter/predefine"

//go:generate converter -type=Animalia

type Animalia struct {
	Info
	Name       predefine.Ebcdic `label:"名前"  byte:"7"`
	Age        predefine.Hex    `label:"年齢"  byte:"2"`
	ExtendArea predefine.Blank  `label:"拡張エリア"  byte:"10"`
}
