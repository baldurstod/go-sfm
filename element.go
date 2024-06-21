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
