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
	"github.com/baldurstod/go-source2-tools/repository"
	"github.com/baldurstod/go-source2-tools/vpk"
)

const varFolder = "./var/"

func initRepo() {
	repository.AddRepository("dota2", vpk.NewVpkFS("R:\\SteamLibrary\\steamapps\\common\\dota 2 beta\\game\\dota\\pak01_dir.vpk"))
}

func TestSession(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initRepo()
	session := sfm.NewSession()

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	camera := shot1.Camera

	err = utils.AddModel(shot1, "Tiny", "dota2", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	if err != nil {
		t.Error(err)
		return
	}
	err = utils.AddModel(shot1, "Marci", "dota2", "models/heroes/marci/marci_base.vmdl")
	if err != nil {
		t.Error(err)
		return
	}

	camera.Transform.Orientation.RotateZ(math.Pi)
	camera.Transform.Position.Set(200, 0, 60)

	//log.Println(buf)
	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().Serialize(session))

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func TestAnimationGroups(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	group := sfm.GetAnimationGroup("bigToe_R_1")

	log.Println(group)
}
