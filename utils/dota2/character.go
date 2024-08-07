package dota2

import (
	"errors"
	"log"

	"github.com/baldurstod/go-dota2"
	"github.com/baldurstod/go-sfm"
	"github.com/baldurstod/go-sfm/utils"
	"github.com/baldurstod/go-vector"
)

type Character struct {
	//slots map[string]*dota2.Item
	hero *dota2.Hero

	animationSet *sfm.AnimationSet
	dag          *sfm.Node

	PosLayer *sfm.LogLayer[vector.Vector3[float32]]
	RotLayer *sfm.LogLayer[vector.Quaternion[float32]]
}

func NewCharacter(name string) (*Character, error) {
	c := Character{
		PosLayer: sfm.NewLogLayer[vector.Vector3[float32]](),
		RotLayer: sfm.NewLogLayer[vector.Quaternion[float32]](),
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

	c.animationSet.GetGameModel().Skin = c.hero.GetSkin()

	tc := c.animationSet.GetTransformControl(sfm.ROOT_TRANSFORM)
	tc.PositionChannel.Log.SetLayer("vector3 log", c.PosLayer)
	tc.OrientationChannel.Log.SetLayer("quaternion log", c.RotLayer)

	return c.animationSet, nil
}

func (c *Character) CreateItemModels(clip *sfm.FilmClip) error {
	if c.hero == nil {
		return errors.New("character doesn't have a hero")
	}

	heroModel := c.animationSet.GetGameModel()
	var itemModel, itemOrHeroModel *sfm.GameModel
	heroItems := c.hero.GetItems()
	modelReplacement := make(map[string]string)

	getReplacement := func(in string) string {
		if replacement, ok := modelReplacement[in]; ok {
			return replacement
		}
		return in
	}

	for _, item := range heroItems {
		modifiers := item.GetAssetModifiers()
		for _, modifier := range modifiers {
			switch modifier.Type {
			case dota2.MODIFIER_MODEL, dota2.MODIFIER_PARTICLE:
				modelReplacement[modifier.Asset] = modifier.Modifier
			}
		}
	}

	for _, item := range heroItems {
		if item == nil {
			continue
		}

		itemModelPlayer := item.GetModelPlayer()

		itemName := item.GetName()
		if itemModelPlayer != "" {
			as2, err := utils.AddModel(clip, itemName, "dota2", getReplacement(itemModelPlayer), c.dag)
			if err != nil {
				return err
			}
			itemModel = as2.GetGameModel()
			if itemModel != nil {
				itemModel.SetParentModel(heroModel)
				itemModel.Skin = item.GetSkin()
			}
		}

		if itemModel != nil {
			itemOrHeroModel = itemModel
		} else {
			itemOrHeroModel = heroModel
		}

		modifiers := item.GetAssetModifiers()
		for _, modifier := range modifiers {
			log.Println("modifier: ", modifier)
			switch modifier.Type {
			case dota2.MODIFIER_PARTICLE_CREATE:
				as2, err := utils.AddParticleSystem(clip, itemName, "dota2", getReplacement(modifier.Modifier), c.dag)
				if err != nil {
					return err
				}
				as2.GetParticleSystem().SetParentModel(itemOrHeroModel)
			case dota2.MODIFIER_MODEL:
				//TODO: replace model
			case dota2.MODIFIER_ABILITY_ICON, dota2.MODIFIER_ANNOUCER_PREVIEW, dota2.MODIFIER_ACTIVITY, dota2.MODIFIER_PARTICLE, dota2.MODIFIER_SOUND, dota2.MODIFIER_RESPONSE_CRITERIA:
			case dota2.MODIFIER_ICON_REPLACEMENT_HERO, dota2.MODIFIER_ICON_REPLACEMENT_HERO_MINIMAP:
			case dota2.MODIFIER_BUFF_MODIFIER, dota2.MODIFIER_CUSTOM_KILL_EFFECT:
			case dota2.MODIFIER_ARCANA_LEVEL, dota2.MODIFIER_INVENTORY_ICON, dota2.MODIFIER_CHATWHEEL:
			case dota2.MODIFIER_ENTITY_MODEL, dota2.MODIFIER_HERO_SCALE, dota2.MODIFIER_MODEL_SKIN:
				//already used for hero model
			case dota2.MODIFIER_ADDITIONAL_WEARABLE:
				asWearable, err := utils.AddModel(clip, itemName+" additional wearable", "dota2", getReplacement(modifier.Asset), c.dag)
				if err != nil {
					return err
				}
				itemModel = asWearable.GetGameModel()
				if itemModel != nil {
					itemModel.SetParentModel(heroModel)
				}
			default:
				log.Println("unknow modifier", modifier.Type)
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

func (c *Character) EquipItem(index string, replaceExisting ...bool) (*dota2.Item, error) {
	if c.hero == nil {
		return nil, errors.New("character doesn't have a hero")
	}

	replace := true
	if len(replaceExisting) == 1 {
		replace = replaceExisting[0]
	}

	item, err := c.hero.EquipItem(index, replace)
	return item, err
}
