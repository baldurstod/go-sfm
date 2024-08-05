package sfm_test

import (
	"math"
	"path"
	"testing"

	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
	"github.com/baldurstod/go-sfm/utils/dota2"
)

func TestParticles(t *testing.T) {
	if err := initDota(); err != nil {
		t.Error(err)
		return
	}

	c, err := dota2.NewCharacter("npc_dota_hero_ogre_magi")
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

	err = c.CreateItemModels(shot1)
	if err != nil {
		t.Error(err)
		return
	}

	err = utils.PlaySequence(as, "idle", shot1.GetDuration())
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
