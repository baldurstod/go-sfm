package sfm_test

import (
	"github.com/baldurstod/go-sfm"
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	session := sfm.NewSession()

	log.Println(session)
}
