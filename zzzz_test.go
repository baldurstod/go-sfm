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
	"github.com/baldurstod/go-sfm/utils/items"
	"github.com/baldurstod/go-source2-tools/repository"
	"github.com/baldurstod/go-source2-tools/vpk"
	"github.com/baldurstod/go-vector"
)

const varFolder = "./var/"

func initRepo() bool {
	repository.AddRepository("dota2", vpk.NewVpkFS("R:\\SteamLibrary\\steamapps\\common\\dota 2 beta\\game\\dota\\pak01_dir.vpk"))
	return true
}

var _ = initRepo()

func TestSession(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	session := sfm.NewSession()

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	camera := shot1.Camera

	_, err = utils.AddModel(shot1, "Tiny", "dota2", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = utils.AddModel(shot1, "Marci", "dota2", "models/heroes/marci/marci_base.vmdl")
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

func TestMovement(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	session := sfm.NewSession()

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 150)

	filename := "models/heroes/shadow_fiend/shadow_fiend"
	filename = "models/heroes/tiny/tiny_01/tiny_01"
	filename = "models/heroes/dawnbreaker/dawnbreaker"
	filename = "models/heroes/slardar/slardar"
	items := [...]string{
		//"models/heroes/dawnbreaker/dawnbreaker_head",
		//"models/heroes/dawnbreaker/dawnbreaker_weapon",
		//"models/heroes/dawnbreaker/dawnbreaker_arms",
		//"models/heroes/dawnbreaker/dawnbreaker_armor",
		//"models/items/dawnbreaker/dawnbreaker_astral_angel_armor/dawnbreaker_astral_angel_armor.vmdl_c",
		"models/items/slardar/takoyaki/slardar_takoyaki.vmdl_c",
	}

	as, err := utils.AddModel(shot1, "Tiny", "dota2", filename)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		as2, err := utils.AddModel(shot1, "Tiny", "dota2", item)
		if err != nil {
			t.Error(err)
			return
		}
		as2.GetGameModel().SetParentModel(as.GetGameModel())

		model, err := utils.GetModel("dota2", item)
		if err != nil {
			t.Error(err)
			return
		}

		seq, err := model.GetSequence("ACT_DOTA_IDLE", nil)
		if err != nil {
			t.Error(err)
			continue
		}

		frames := seq.GetFrameCount()
		fps := seq.GetFps()

		for frameId := 0; frameId < frames; frameId++ {
			frame, err := seq.GetFrame(frameId)
			frameTime := float32(frameId) / float32(fps)
			if err != nil {
				t.Error(err)
				return
			}

			as2.SetFrame(frameTime, frame, nil)
		}

	}

	tiny, err := utils.GetModel("dota2", filename)
	if err != nil {
		t.Error(err)
		return
	}

	modifiers := make(map[string]struct{})
	modifiers["tako"] = struct{}{}
	seq, err := tiny.GetSequence("ACT_DOTA_IDLE", modifiers)
	if err != nil {
		t.Error(err)
		return
	}
	/*
		modifiers := make(map[string]struct{})
		modifiers["PostGameIdle"] = struct{}{}
		seq, err = tiny.GetSequence("ACT_DOTA_LOADOUT", modifiers)
		if err != nil {
			t.Error(err)
			return
		}
	*/

	frames := seq.GetFrameCount()
	fps := seq.GetFps()

	flexes, err := tiny.GetFlexes()
	if err != nil {
		t.Error(err)
		return
	}

	for frameId := 0; frameId < frames; frameId++ {
		frame, err := seq.GetFrame(frameId)
		frameTime := float32(frameId) / float32(fps)
		if err != nil {
			t.Error(err)
			return
		}

		as.SetFrame(frameTime, frame, flexes)
		//log.Println(frame)
	}

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().Serialize(session))

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func DisbaledTestMovement(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	session := sfm.NewSession()

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 150)

	filename := "models/heroes/shadow_fiend/shadow_fiend"
	filename = "models/heroes/tiny/tiny_01/tiny_01"
	filename = "models/heroes/dawnbreaker/dawnbreaker"
	/*items := [...]string{
		"models/heroes/dawnbreaker/dawnbreaker_head",
		"models/heroes/dawnbreaker/dawnbreaker_weapon",
		"models/heroes/dawnbreaker/dawnbreaker_arms",
		"models/heroes/dawnbreaker/dawnbreaker_armor",
	}*/

	as, err := utils.AddModel(shot1, "Tiny", "dota2", filename)
	if err != nil {
		t.Error(err)
		return
	}

	buf1, err := os.ReadFile("./var/items_game.txt")
	if err != nil {
		panic(err)
	}

	ig := items.NewItemsGame(buf1)
	items := ig.GetItemsPerHero("npc_dota_hero_dawnbreaker")
	for _, item := range items {
		log.Println(item.GetModelPlayer())

	}

	for _, item := range items {
		modelPlayer, ok := item.GetModelPlayer()
		if !ok {
			continue
		}

		as2, err := utils.AddModel(shot1, "Tiny", "dota2", modelPlayer)
		if err != nil {
			t.Error(err)
			return
		}
		as2.GetGameModel().SetParentModel(as.GetGameModel())
	}

	tiny, err := utils.GetModel("dota2", filename)
	if err != nil {
		t.Error(err)
		return
	}

	seq, err := tiny.GetSequence("ACT_DOTA_IDLE_RARE", nil)
	if err != nil {
		t.Error(err)
		return
	}
	/*
		modifiers := make(map[string]struct{})
		modifiers["PostGameIdle"] = struct{}{}
		seq, err = tiny.GetSequence("ACT_DOTA_LOADOUT", modifiers)
		if err != nil {
			t.Error(err)
			return
		}
	*/

	frames := seq.GetFrameCount()
	fps := seq.GetFps()

	flexes, err := tiny.GetFlexes()
	if err != nil {
		t.Error(err)
		return
	}

	for frameId := 0; frameId < frames; frameId++ {
		frame, err := seq.GetFrame(frameId)
		frameTime := float32(frameId) / float32(fps)
		if err != nil {
			t.Error(err)
			return
		}
		//log.Println(frame)
		positionChannel := frame.GetChannel("BoneChannel", "Position")
		for _, element := range positionChannel.Datas {

			tc := as.GetTransformControl(element.Name)
			if tc != nil {
				layer := any(tc.PositionChannel.Log.GetLayer("vector3 log")).(*sfm.LogLayer[vector.Vector3[float32]])
				layer.SetValue(frameTime, element.Datas.(vector.Vector3[float32]))
			}
		}
		orientationChannel := frame.GetChannel("BoneChannel", "Angle")
		for _, element := range orientationChannel.Datas {

			tc := as.GetTransformControl(element.Name)
			if tc != nil {
				layer := any(tc.OrientationChannel.Log.GetLayer("quaternion log")).(*sfm.LogLayer[vector.Quaternion[float32]])
				layer.SetValue(frameTime, (element.Datas.(vector.Quaternion[float32])))
			}
		}

		morphChannel := frame.GetChannel("MorphChannel", "data")
		for _, element := range morphChannel.Datas {

			value := float32(0)
			for _, flex := range flexes {
				if flex.Name == element.Name {
					value = flex.GetControllerValue((element.Datas.(float32))) //((element.Datas.(float32)) - flex.Min) / (flex.Max - flex.Min)
					break
				}
			}

			tc := as.GetControl(element.Name)
			if tc != nil {
				layer := any(tc.Channel.Log.GetLayer("float log")).(*sfm.LogLayer[float32])
				layer.SetValue(frameTime, value)
			}
		}
	}

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().Serialize(session))

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func TestAnimationGroups(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	group := sfm.GetAnimationGroup("bigToe_R_1")

	log.Println(group)
}

func TestItems(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	buf, err := os.ReadFile("./var/items_game.txt")
	if err != nil {
		panic(err)
	}
	/*
		vdf := vdf.VDF{}
		items := vdf.Parse(buf)
		log.Println(items)*/
	ig := items.NewItemsGame(buf)
	items := ig.GetItemsPerHero("npc_dota_hero_dawnbreaker")
	for _, item := range items {
		log.Println(item.GetModelPlayer())

	}
}
