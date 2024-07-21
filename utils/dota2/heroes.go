package dota2

import (
	"errors"
	"log"
	"strings"

	"github.com/baldurstod/vdf"
)

func createHeroes() map[string]*Hero {
	return map[string]*Hero{}
}

var heroes = createHeroes()

func InitHeroes(buf []byte) error {
	vdf := vdf.VDF{}
	root := vdf.Parse(buf)
	//log.Println(root.Get("DOTAHeroes"))
	heroes, err := root.Get("DOTAHeroes")

	if err != nil {
		return err
	}

	for _, hero := range heroes.GetChilds() {
		if strings.HasPrefix(hero.Key, "npc_") {
			addHero(hero)
		}
	}

	log.Println(heroes)

	return nil
}

func addHero(datas *vdf.KeyValue) (*Hero, error) {
	hero := NewHero(datas.Key)
	hero.initFromData(datas)

	heroes[hero.entity] = hero
	log.Println(hero)

	return hero, nil
}

func GetHero(entity string) (*Hero, error) {
	h, ok := heroes[entity]
	if !ok {
		return nil, errors.New("hero not found " + entity)
	}
	return h, nil
}
