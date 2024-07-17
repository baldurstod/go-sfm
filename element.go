package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Element interface {
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
	isExportable() bool
}

type Operator interface {
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
	isOperator()
	isExportable() bool
}

type INode interface {
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
	isNode()
	AddChildren(child INode)
	isExportable() bool
}

type IControl interface {
	isControl()
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
	isExportable() bool
}

type IClip interface {
	isClip()
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
	isExportable() bool
}

type ILog interface {
	isLog()
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
	AddLayer(name string) ILogLayer
	GetLayer(name string) ILogLayer
	isExportable() bool
}

type ILogLayer interface {
	isLogLayer()
	toDmElement(*Serializer, *dmx.DmElement)
	createDmElement(*Serializer) *dmx.DmElement
}
