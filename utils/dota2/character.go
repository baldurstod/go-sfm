package dota2

import (
	"errors"

	"github.com/baldurstod/go-dota2"
	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
)

type Character struct {
	slots map[string]*dota2.Item
	hero  *dota2.Hero
}

func NewCharacter(name string) (*Character, error) {
	c := Character{
		slots: make(map[string]*dota2.Item),
	}

	h, err := dota2.GetHero(name)
	if err != nil {
		return nil, err
	}
	c.hero = h

	// Init slots
	for _, slot := range h.ItemSlots {
		c.slots[slot.SlotName] = nil
	}

	// Init base items
	for _, item := range dota2.GetBaseItems(name) {
		if _, ok := c.slots[item.ItemSlot]; !ok {
			return nil, errors.New("unknown slot : " + item.ItemSlot)
		}

		c.slots[item.ItemSlot] = item
	}

	return &c, nil
}

func (c *Character) CreateGameModel(clip *sfm.FilmClip) error {
	as, err := utils.AddModel(clip, c.hero.Entity, "dota2", c.hero.Model)
	if err != nil {
		return err
	}
	return nil

	for _, item := range c.slots {
		if item == nil || item.ModelPlayer == "" {
			continue
		}

		as2, err := utils.AddModel(clip, "Tiny", "dota2", item.ModelPlayer)
		if err != nil {
			return err
		}
		as2.GetGameModel().SetParentModel(as.GetGameModel())
	}

	return nil
}
