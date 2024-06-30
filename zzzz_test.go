package sfm_test

import (
	"bytes"
	"log"
	"os"
	"path"
	"testing"

	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
)

const varFolder = "./var/"

func TestSession(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	session := sfm.NewSession()

	_, scene, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	err = utils.AddModel(scene, "models/heroes/tiny/tiny_01/tiny_01.vmdl", "tiny_01.vmdl_c")
	if err != nil {
		t.Error(err)
		return
	}

	//log.Println(buf)
	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func TestCamera(t *testing.T) {
	//camera := sfm.Camera{}

	//log.Println(camera)
}
