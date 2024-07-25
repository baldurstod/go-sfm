package sfm_test

import (
	"log"
	"math"
	"os"
	"path"
	"testing"

	dota "github.com/baldurstod/go-dota2"
	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
	"github.com/baldurstod/go-sfm/utils/dota2"
)

func initDota() error {
	buf, err := os.ReadFile(varFolder + "npc_heroes.txt")
	if err != nil {
		return err
	}
	err = dota.InitHeroes(buf)
	if err != nil {
		return err
	}

	buf, err = os.ReadFile(varFolder + "items_game.txt")
	if err != nil {
		return err
	}
	err = dota.InitItems(buf)
	if err != nil {
		return err
	}
	return nil
}

func TestHeroes(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := initDota(); err != nil {
		t.Error(err)
		return

	}

	c, err := dota2.NewCharacter("npc_dota_hero_dark_willow")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(c)
}

func TestSfmHeroes(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := initDota(); err != nil {
		t.Error(err)
		return
	}

	c, err := dota2.NewCharacter("npc_dota_hero_dark_willow")
	if err != nil {
		t.Error(err)
		return
	}

	session := sfm.NewSession()

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 150)

	c.CreateGameModel(shot1)

	session.WriteBinaryFile(path.Join(varFolder, "test_session.dmx"))
}
