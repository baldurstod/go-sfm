package sfm_test

import (
	"log"
	"math"
	"path"
	"testing"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
	"github.com/baldurstod/go-sfm/utils/dota2"
	"github.com/baldurstod/go-vector"
)

func TestHeroes(t *testing.T) {
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

	shot1 := utils.CreateClip(session)

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 150)

	as, err := c.CreateGameModel(shot1)
	if err != nil {
		t.Error(err)
		return
	}

	err = utils.PlaySequence(as, "idle", shot1.GetDuration())
	if err != nil {
		t.Error(err)
		return
	}

	err = session.WriteBinaryFile(path.Join(varFolder, "test_session.dmx"))
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRazorArcana(t *testing.T) {
	if err := initDota(); err != nil {
		t.Error(err)
		return
	}

	c, err := dota2.NewCharacter("npc_dota_hero_razor")
	if err != nil {
		t.Error(err)
		return
	}

	if item, err := c.EquipItem("23100"); err != nil {
		t.Error(err)
		return
	} else {
		item.Style = 1
	}

	session := sfm.NewSession()

	shot1 := utils.CreateClip(session)

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 150)

	as, err := c.CreateGameModel(shot1)
	if err != nil {
		t.Error(err)
		return
	}

	tc := as.GetTransformControl(sfm.ROOT_TRANSFORM)
	layer := tc.PositionChannel.Log.GetLayer("vector3 log").(*sfm.LogLayer[vector.Vector3[float32]])
	layer.SetValue(0, vector.Vector3[float32]{500, 0, 0})

	err = c.CreateItemModels(shot1)
	if err != nil {
		t.Error(err)
		return
	}

	err = utils.PlayActivity(as, c.CreateActivity("ACT_DOTA_IDLE"), shot1.GetDuration())
	if err != nil {
		t.Error(err)
		return
	}

	err = session.WriteTextFile(path.Join(varFolder, "test_session.dmx"))
	if err != nil {
		t.Error(err)
		return
	}
}
