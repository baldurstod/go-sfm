package sfm_test

import (
	"math"
	"path"
	"testing"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
	"github.com/baldurstod/go-sfm/utils/dota2"
	"github.com/baldurstod/go-source2-tools/model"
)

func TestParticles(t *testing.T) {
	if err := initDota(); err != nil {
		t.Error(err)
		return
	}
	/*
		npc_dota_hero_drow_ranger
		npc_dota_hero_windrunner
		npc_dota_hero_ogre_magi
		npc_dota_hero_axe
		npc_dota_hero_dawnbreaker
	*/

	c, err := dota2.NewCharacter("npc_dota_hero_axe")
	if err != nil {
		t.Error(err)
		return
	}

	c.EquipItem("12964")

	session := sfm.NewSession()

	shot1 := utils.CreateClip(session)

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 150)

	as, err := c.CreateGameModel(shot1)
	if err != nil {
		t.Error(err)
		return
	}

	err = c.CreateItemModels(shot1)
	if err != nil {
		t.Error(err)
		return
	}

	err = utils.PlayActivity(as, model.NewActivity("ACT_DOTA_IDLE"), 0.1)
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
