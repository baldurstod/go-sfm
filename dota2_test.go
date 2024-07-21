package sfm_test

import (
	"log"
	"os"
	"testing"

	dota "github.com/baldurstod/go-dota2"
	"github.com/baldurstod/go-sfm/utils/dota2"
)

func TestHeroes(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	buf, err := os.ReadFile(varFolder + "npc_heroes.txt")
	if err != nil {
		t.Error(err)
		return
	}
	err = dota.InitHeroes(buf)
	if err != nil {
		t.Error(err)
		return
	}

	buf, err = os.ReadFile(varFolder + "items_game.txt")
	if err != nil {
		t.Error(err)
		return
	}
	err = dota.InitItems(buf)
	if err != nil {
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
