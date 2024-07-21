package items

import (
	"github.com/baldurstod/vdf"
)

type itemMap = map[string]*item

type itemsGame struct {
	vdf     *vdf.KeyValue
	prefabs itemMap
	items   itemMap
}

func NewItemsGame(dat []byte) *itemsGame {
	ig := itemsGame{}
	ig.init(dat)
	return &ig
}

func (ig *itemsGame) init(dat []byte) {
	vdf := vdf.VDF{}
	root := vdf.Parse(dat)
	ig.vdf, _ = root.Get("items_game")
	ig.prefabs = make(itemMap)
	ig.items = make(itemMap)

	if prefabs, err := ig.vdf.Get("prefabs"); err != nil {
		for _, val := range prefabs.GetChilds() {
			var it = item{}
			if it.init(ig, val) {
				ig.prefabs[it.id] = &it
			}
		}
	}

	if items, err := ig.vdf.Get("items"); err != nil {
		for _, val := range items.GetChilds() {
			var it = item{}
			if it.init(ig, val) {
				ig.items[it.id] = &it
			}
		}
	}

	for _, it := range ig.items {
		it.initPrefabs()
	}
}

func (ig *itemsGame) getPrefab(prefabName string) *item {
	return ig.prefabs[prefabName]
}

func (ig *itemsGame) GetItemsPerHero(npc string) []*item {
	items := []*item{}

	for _, item := range ig.items {
		for _, n := range item.getUsedByHeroes() {
			if npc == n {
				items = append(items, item)
			}
		}
	}
	return items
}
