package dota2

import (
	"errors"
	"log"

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
	dag := sfm.NewNode(c.hero.Entity)
	clip.Scene.AddChildren(dag)

	as, err := utils.AddModel(clip, c.hero.Entity, "dota2", c.hero.Model, dag)
	if err != nil {
		return err
	}

	for _, item := range c.slots {
		if item == nil {
			continue
		}

		if item.ModelPlayer != "" {
			as2, err := utils.AddModel(clip, item.Name, "dota2", item.ModelPlayer, dag)
			if err != nil {
				return err
			}
			as2.GetGameModel().SetParentModel(as.GetGameModel())
		}

		modifiers := item.GetAssetModifiers(0)
		for _, modifier := range modifiers {
			log.Println(modifier)
			switch modifier.Type {
			case "particle_create":
				as2, err := utils.AddParticleSystem(clip, item.Name, "dota2", modifier.Modifier, dag)
				if err != nil {
					return err
				}
				as2.GetParticleSystem().SetParentModel(as.GetGameModel())

			}

		}
	}

	return nil
}
