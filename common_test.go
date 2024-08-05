package sfm_test

import (
	"log"
	"os"

	dota "github.com/baldurstod/go-dota2"
	"github.com/baldurstod/go-source2-tools/repository"
	"github.com/baldurstod/go-source2-tools/vpk"
)

const varFolder = "./var/"

func initRepo() bool {
	repository.AddRepository("dota2", vpk.NewVpkFS("R:\\SteamLibrary\\steamapps\\common\\dota 2 beta\\game\\dota\\pak01_dir.vpk"))
	return true
}

var _ = initRepo()

var _ = func() bool { log.SetFlags(log.LstdFlags | log.Lshortfile); return true }()

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
