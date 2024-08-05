package dota2

import (
	"errors"
	"log"

	"github.com/baldurstod/go-dota2"
	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
)

type Character struct {
	//slots map[string]*dota2.Item
	hero *dota2.Hero

	animationSet *sfm.AnimationSet
	dag          *sfm.Node
}

func NewCharacter(name string) (*Character, error) {
	c := Character{
		//slots: make(map[string]*dota2.Item),
	}

	h, err := dota2.GetHero(name)
	if err != nil {
		return nil, err
	}
	c.hero = h

	// Init slots
	/*
		for _, slot := range h.ItemSlots {
			c.slots[slot.SlotName] = nil
		}
	*/

	// Init base items
	/*
		for _, item := range h.GetItems() {
			if _, ok := c.slots[item.ItemSlot]; !ok {
				return nil, errors.New("unknown slot : " + item.ItemSlot)
			}

			c.slots[item.ItemSlot] = item
		}*/

	return &c, nil
}

func (c *Character) CreateGameModel(clip *sfm.FilmClip) (*sfm.AnimationSet, error) {
	if c.hero == nil {
		return nil, errors.New("character doesn't have a hero")
	}

	entity := c.hero.GetEntity()

	c.dag = sfm.NewNode(entity)
	clip.Scene.AddChildren(c.dag)

	var err error
	c.animationSet, err = utils.AddModel(clip, entity, "dota2", c.hero.GetModel(), c.dag)
	if err != nil {
		return nil, err
	}
	/*
		for _, item := range c.hero.GetItems() {
			if item == nil {
				continue
			}

			if item.ModelPlayer != "" {
				as2, err := utils.AddModel(clip, item.Name, "dota2", item.ModelPlayer, dag)
				if err != nil {
					return nil, err
				}
				as2.GetGameModel().SetParentModel(as.GetGameModel())
			}

			modifiers := item.GetAssetModifiers(0)
			for _, modifier := range modifiers {
				log.Println(modifier)
				switch modifier.Type {
				case dota2.MODIFIER_PARTICLE_CREATE:
					as2, err := utils.AddParticleSystem(clip, item.Name, "dota2", modifier.Modifier, dag)
					if err != nil {
						return nil, err
					}
					as2.GetParticleSystem().SetParentModel(as.GetGameModel())

				}

			}
		}
	*/
	return c.animationSet, nil
}

func (c *Character) CreateItemModels(clip *sfm.FilmClip) error {
	if c.hero == nil {
		return errors.New("character doesn't have a hero")
	}

	for _, item := range c.hero.GetItems() {
		if item == nil {
			continue
		}

		itemModelPlayer := item.GetModelPlayer()
		itemName := item.GetName()
		var itemModel *sfm.GameModel
		if itemModelPlayer != "" {
			as2, err := utils.AddModel(clip, itemName, "dota2", itemModelPlayer, c.dag)
			if err != nil {
				return err
			}
			itemModel = as2.GetGameModel()
			if itemModel != nil {
				itemModel.SetParentModel(c.animationSet.GetGameModel())
				itemModel.Skin = int32(item.GetSkin())
			}
		}

		modifiers := item.GetAssetModifiers(0)
		for _, modifier := range modifiers {
			log.Println(modifier)
			switch modifier.Type {
			case dota2.MODIFIER_PARTICLE_CREATE:
				as2, err := utils.AddParticleSystem(clip, itemName, "dota2", modifier.Modifier, c.dag)
				if err != nil {
					return err
				}
				as2.GetParticleSystem().SetParentModel(itemModel)
			}

		}
	}
	return nil
}

func (c *Character) GetEntity() string {
	if c.hero == nil {
		return ""
	}
	return c.hero.GetEntity()
}

func (c *Character) EquipItem(index string) error {
	if c.hero == nil {
		return errors.New("character doesn't have a hero")
	}

	_, err := c.hero.EquipItem(index)
	return err
}
