package baggage

import (
	"github.com/t10471/converter/examples/message/creature"
	"github.com/t10471/converter/predefine"
)

//go:generate converter -type=List

type List struct {
	Name          predefine.Ebcdic        `label:"名前" byte:"7"`
	PlantaeCount  predefine.Hex           `label:"植物数" byte:"2"`
	PlantaeList   []creature.Plantae      `label:"植物リスト" length:"10"`
	AnimaliaCount predefine.Hex           `label:"動物数" byte:"2"`
	AnimaliaList  []creature.Animalia     `label:"動物リスト" length:"10"`
	Cost          predefine.PackedDecimal `label:"金額" byte:"7"`
}
