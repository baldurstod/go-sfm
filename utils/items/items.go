package items

import (
	"strings"

	"github.com/baldurstod/vdf"
)

type item struct {
	ig      *itemsGame
	id      string
	prefabs []*item
	kv      *vdf.KeyValue
}

func (i *item) init(ig *itemsGame, kv *vdf.KeyValue) bool {
	i.ig = ig
	i.id = kv.Key
	i.kv = kv

	return true
}

func (i *item) initPrefabs() {
	if s, err := i.kv.GetString("prefab"); err != nil {
		prefabs := strings.Split(s, " ")
		for _, prefabName := range prefabs {
			prefab := i.ig.getPrefab(prefabName)
			prefab.initPrefabs() //Ensure prefab is initialized
			i.prefabs = append(i.prefabs, prefab)
		}
	}
}

func (i *item) getUsedByHeroes() []string {
	ret := []string{}

	if usedByHeroes, err := i.kv.GetStringMap("used_by_heroes"); err != nil {
		for key, val := range *usedByHeroes {
			if val == "1" {
				ret = append(ret, key)
			}
		}
	}
	return ret
}

func (i *item) GetModelPlayer() (string, bool) {
	return i.getStringAttribute("model_player")
}

func (i *item) getStringAttribute(attributeName string) (string, bool) {
	if s, err := i.kv.GetString(attributeName); err != nil {
		return s, true
	}

	for _, prefab := range i.prefabs {
		if s, ok := prefab.getStringAttribute(attributeName); ok && s != "0" { //TODO: remove s != "0"
			return s, true
		}
	}
	return "", false
}
