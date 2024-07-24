package sfm

import "fmt"

type ControlPoint struct {
	Name string
	*Node
	transformControl *TransformControl
	ControlModel     *GameModel
}

func (cp *ControlPoint) isNode() {
}

func (cp *ControlPoint) isExportable() bool {
	return true
}

func newControlPoint(id int) *ControlPoint {
	name := fmt.Sprintf("controlPoint%d", id)
	return &ControlPoint{
		Node: NewNode(name),
		Name: name,
	}
}
