package sfm_test

import (
	"bytes"
	"log"
	"math"
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

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	camera := shot1.Camera

	err = utils.AddModel(shot1, "models/heroes/tiny/tiny_01/tiny_01.vmdl", "tiny_01.vmdl_c")
	if err != nil {
		t.Error(err)
		return
	}

	camera.Transform.Orientation.RotateZ(math.Pi)
	camera.Transform.Position.Set(200, 0, 60)

	//log.Println(buf)
	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().GetElement(session))

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func TestAnimationGroups(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	group := sfm.GetAnimationGroup("bigToe_R_1")
	if group == nil {
		t.Error("group not found")
		return
	}

	log.Println(group)
}
