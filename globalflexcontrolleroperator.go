package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type GlobalFlexControllerOperator struct {
	Name       string
	FlexWeight float32
	GameModel  *GameModel
}

func NewGlobalFlexControllerOperator(name string, flexWeight float32, gameModel *GameModel) *GlobalFlexControllerOperator {
	return &GlobalFlexControllerOperator{
		Name:       name,
		FlexWeight: flexWeight,
		GameModel:  gameModel,
	}
}

func (o *GlobalFlexControllerOperator) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(o.Name, "DmeGlobalFlexControllerOperator")
}

func (o *GlobalFlexControllerOperator) isExportable() bool {
	return true
}

func (o *GlobalFlexControllerOperator) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateFloatAttribute("flexWeight", o.FlexWeight)
	e.CreateElementAttribute("gameModel", serializer.GetElement(o.GameModel))
}

/*
"DmeGlobalFlexControllerOperator"
{
	"id" "elementid" "66a1d5a0-8425-4d1f-b344-fdd442dcce4e"
	"name" "string" "innerBrowRaiser"
	"flexWeight" "float" "0"
	"gameModel" "element" "94de5e75-487b-4aaa-81c2-8de83928762c"
}

*/
