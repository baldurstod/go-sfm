package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Element interface {
	toDmElement(*Serializer) *dmx.DmElement
}

type Operator interface {
	toDmElement(*Serializer) *dmx.DmElement
	isOperator()
}

type INode interface {
	toDmElement(*Serializer) *dmx.DmElement
	isNode()
	AddChildren(child INode)
}

type IControl interface {
	isControl()
	toDmElement(*Serializer) *dmx.DmElement
}

type IClip interface {
	isClip()
	toDmElement(*Serializer) *dmx.DmElement
}
