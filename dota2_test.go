package sfm_test

import (
	"log"
	"os"
	"testing"

	"github.com/baldurstod/go-sfm/utils/dota2"
)

func TestHeroes(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("start heroes")

	buf, err := os.ReadFile(varFolder + "npc_heroes.txt")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("start parse")
	dota2.InitHeroes(buf)
	log.Println("end heroes")
}
