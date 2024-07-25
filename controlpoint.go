package sfm

import "fmt"

type ControlPoint struct {
	Name string
	*Node
	transformControl *TransformControl
	ControlModel     *GameModel
	AttachType       string
	AttachmentName   string
	EntityName       string
}

func (cp *ControlPoint) isNode() {
}

func (cp *ControlPoint) isExportable() bool {
	return true
}

func newControlPoint(id uint) *ControlPoint {
	name := fmt.Sprintf("controlPoint%d", id)
	return &ControlPoint{
		Node: NewNode(name),
		Name: name,
	}
}
