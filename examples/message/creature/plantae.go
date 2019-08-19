package creature

import "github.com/t10471/converter/predefine"

//go:generate converter -type=Plantae

type Plantae struct {
	Info
	Number     predefine.Hex   `label:"番号"  byte:"2"`
	ExtendArea predefine.Blank `label:"拡張エリア"  byte:"10"`
}
