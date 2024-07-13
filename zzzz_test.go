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
	"github.com/baldurstod/go-vector"
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
	initRepo()
	session := sfm.NewSession()

	shot1, err := utils.CreateClip(session)
	if err != nil {
		t.Error(err)
		return
	}

	shot1.Camera.Transform.Orientation.RotateZ(math.Pi)
	shot1.Camera.Transform.Position.Set(200, 0, 60)

	as, err := utils.AddModel(shot1, "Tiny", "dota2", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	if err != nil {
		t.Error(err)
		return
	}

	tiny, err := utils.GetModel("dota2", "models/heroes/tiny/tiny_01/tiny_01.vmdl")
	if err != nil {
		t.Error(err)
		return
	}

	seq, err := tiny.GetSequence("ACT_DOTA_RUN", nil)
	if err != nil {
		t.Error(err)
		return
	}

	frames := seq.GetFrameCount()
	//fps := seq.GetFps()

	for frameId := 0; frameId < frames; frameId++ {
		frame, err := seq.GetFrame(frameId)
		if err != nil {
			t.Error(err)
			return
		}
		//log.Println(frame)
		positionChannel := frame.GetChannel("BoneChannel", "Position")
		for _, element := range positionChannel.Datas {

			tc := as.GetTransformControl(element.Name)
			if tc != nil {
				//log.Println(tc)
				layer := any(tc.PositionChannel.Log.GetLayer("quaternion log")).(*sfm.LogLayer[vector.Vector3[float32]])
				//log.Println(layer)

				//layer.AddValue(0, vector.Vector3[float32]{})
				layer.AddValue(int32(frameId), element.Datas.(vector.Vector3[float32]))
			}
		}
		orientationChannel := frame.GetChannel("BoneChannel", "Angle")
		for _, element := range orientationChannel.Datas {

			tc := as.GetTransformControl(element.Name)
			if tc != nil {
				//log.Println(tc)
				layer := any(tc.OrientationChannel.Log.GetLayer("vector3 log")).(*sfm.LogLayer[vector.Quaternion[float32]])
				//log.Println(layer)

				//layer.AddValue(0, vector.Quaternion[float32]{})
				layer.AddValue(int32(frameId), (element.Datas.(vector.Quaternion[float32])))
			}
		}
	}
	/*
		tc := as.GetTransformControl("clavicle_R")
		if tc != nil {
			log.Println(tc)
			layer := any(tc.PositionChannel.Log.AddLayer()).(*sfm.LogLayer[vector.Vector3[float32]])
			log.Println(layer)

			for i := 0; i < 100; i++ {
				layer.AddValue(int32(i), vector.Vector3[float32]{float32(i), 0, 0})
			}
		}*/

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, sfm.NewSerializer().Serialize(session))

	os.WriteFile(path.Join(varFolder, "test_session.dmx"), buf.Bytes(), 0666)
}

func TestAnimationGroups(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	group := sfm.GetAnimationGroup("bigToe_R_1")

	log.Println(group)
}
