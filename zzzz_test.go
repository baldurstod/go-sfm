package sfm_test

import (
	"bytes"
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-sfm"
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	session := sfm.NewSession()
	session.CreateClip()

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, session.ToDmElement())
	log.Println(buf)
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
