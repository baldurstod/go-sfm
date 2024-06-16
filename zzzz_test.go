package sfm_test

import (
	"github.com/baldurstod/go-sfm"
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	session := sfm.NewSession()
	e := session.ToDmElement()

	log.Println(session, e)
}

func TestCamera(t *testing.T) {
	camera := sfm.Camera{}

	log.Println(camera)
}
