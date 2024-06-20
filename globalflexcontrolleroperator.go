package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type GlobalFlexControllerOperator struct {
	Name       string
	FlexWeight float64
	GameModel  *GameModel
}

func NewGlobalFlexControllerOperator(name string, flexWeight float64, gameModel *GameModel) *GlobalFlexControllerOperator {
	return &GlobalFlexControllerOperator{
		Name:       name,
		FlexWeight: flexWeight,
		GameModel:  gameModel,
	}
}

func (o *GlobalFlexControllerOperator) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeGlobalFlexControllerOperator")

	e.CreateStringAttribute("name", o.Name)
	e.CreateFloatAttribute("position", o.FlexWeight)
	e.CreateElementAttribute("gameModel", serializer.GetElement(o.GameModel))

	return e
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
