package dota2

import (
	"errors"
	"strconv"

	"github.com/baldurstod/vdf"
)

type Hero struct {
	entity      string
	Name        string
	HeroID      int
	HeroOrderID int
	Model       string
	Personas    []Persona
	ItemSlots   []ItemSlot
}

func NewHero(entity string) *Hero {
	return &Hero{
		entity:    entity,
		Personas:  make([]Persona, 0),
		ItemSlots: make([]ItemSlot, 0, 5),
	}
}

func (h *Hero) initFromData(data *vdf.KeyValue) error {
	var err error
	if h.Model, err = data.GetString("Model"); err != nil {
		return errors.New("can't find Model key")
	}

	if h.HeroID, err = data.GetInt("HeroID"); err != nil {
		return errors.New("can't find HeroID key")
	}

	return nil
}

func (h *Hero) String() string {
	return "Model: " + h.Model + "HeroID" + strconv.Itoa(h.HeroID)
}

/*
	//=================================================================================================================
	// HERO: Base
	// Note: This is loaded and overriden/added to by values in the specific heroes chunks.
	//=================================================================================================================
	"npc_dota_hero_base"
	{
		// General
		//-------------------------------------------------------------------------------------------------------------
		"Model"		"models/dev/error.vmdl"
		"SoundSet"		"0"
		"Enabled"		"0"
		"Level"		"1"
		"BotImplemented"		"0"
		"NewHero"		"0"
		"HeroPool1"		"0"
		"HeroUnlockOrder"		"1"
		"CMEnabled"		"0"
		"new_player_enable"		"0"
		"Adjectives"
		{
			"Wings"			"0"
			"Horns"			"0"
			"Legs"			"2"
			"Steed"			"0"
			"Nose"			"1"
			"Fuzzy"			"0"
			"Bearded"		"0"
			"Female"		"0"
			"BadTeeth"		"0"
			"Cape"			"0"
			"NicePecs"		"0"
			"Potbelly"		"0"
			"Parent"		"0"
			"Arachnophobic"	"0"
			"Undead"		"0"
			"Aquatic"		"0"
			"Demon"			"0"
			"Spirit"		"0"
			"Flying"		"0"
			"Cute"			"0"
			"Fiery"			"0"
			"Icy"			"0"
			"Blue"			"0"
			"Red"			"0"
			"Green"			"0"
		}

		// Abilities
		//-------------------------------------------------------------------------------------------------------------
		"Ability1"					""										// Ability 1.
		"Ability2"					""										// Ability 2.
		"Ability3"					""										// Ability 3.
		"Ability4"					""										// Ability 4.
		"Ability5"					""
		"Ability6"					""										// Ability 6 - Extra.
		"Ability7"					""										// Ability 7 - Extra.
		"Ability8"					""										// Ability 8 - Extra.
		"Ability9"					""
		"Ability25"		"special_bonus_attributes"

		"AbilityTalentStart"		"10"

		// Armor
		//-------------------------------------------------------------------------------------------------------------
		"ArmorPhysical"		"-1"
		"MagicalResistance"		"25"

		// Attack
		//-------------------------------------------------------------------------------------------------------------
		"AttackCapabilities"		"DOTA_UNIT_CAP_RANGED_ATTACK"
		"BaseAttackSpeed"		"100"
		"AttackDamageMin"		"1"
		"AttackDamageMax"		"1"
		"AttackDamageType"		"DAMAGE_TYPE_ArmorPhysical"
		"AttackRate"		"1.700000"
		"AttackAnimationPoint"		"0.750000"
		"AttackAcquisitionRange"		"800"
		"AttackRange"		"600"
		"ProjectileModel"		"particles/base_attacks/ranged_hero.vpcf"
		"ProjectileSpeed"		"900"

		// Attributes
		//-------------------------------------------------------------------------------------------------------------
		"AttributePrimary"		"DOTA_ATTRIBUTE_STRENGTH"
		"AttributeBaseStrength"		"0"
		"AttributeStrengthGain"		"0"
		"AttributeBaseIntelligence"		"0"
		"AttributeIntelligenceGain"		"0"
		"AttributeBaseAgility"		"0"
		"AttributeAgilityGain"		"0"

		// Bounty
		//-------------------------------------------------------------------------------------------------------------
		"BountyXP"		"62"
		"BountyGoldMin"		"0"
		"BountyGoldMax"		"0"

		// Bounds
		//-------------------------------------------------------------------------------------------------------------
		"BoundsHullName"		"DOTA_HULL_SIZE_HERO"
		"RingRadius"		"70"

		// Movement
		//-------------------------------------------------------------------------------------------------------------
		"MovementCapabilities"		"DOTA_UNIT_CAP_MOVE_GROUND"
		"MovementSpeed"		"300"
		"MovementTurnRate"		"0.600000"
		"HasAggressiveStance"		"0"

		// Status
		//-------------------------------------------------------------------------------------------------------------
		"StatusHealth"		"120"
		"StatusMana"		"75"
		"StatusManaRegen"	"0"
		"StatusHealthRegen"		"0.25"

		// Team
		//-------------------------------------------------------------------------------------------------------------
		"TeamName"		"DOTA_TEAM_GOODGUYS"
		"CombatClassAttack"		"DOTA_COMBAT_CLASS_ATTACK_HERO"
		"CombatClassDefend"		"DOTA_COMBAT_CLASS_DEFEND_HERO"
		"UnitRelationshipClass"		"DOTA_NPC_UNIT_RELATIONSHIP_TYPE_HERO"

		// Vision
		//-------------------------------------------------------------------------------------------------------------
		"VisionDaytimeRange"		"1800"
		"VisionNighttimeRange"		"800"

		//Inventory
		"HasInventory"		"1"

		//Voice
		"VoiceBackgroundSound"		""
		"HealthBarOffset"		"200"
		"IdleExpression"		"scenes/default_idle.vcd"
		"IdleSoundLoop"				""
		"ARDMDisabled"		"0"
		"HUD"
		{
			"StatusHUD"
			{
				"StatusStrength"
				{
					"LocalizeToken"		"#DOTA_StrengthStatus"
					"Parameters"		"Traits:Strength"
					"HUDName"		"unitstrength"
				}
				"StatusAgility"
				{
					"LocalizeToken"		"#DOTA_AgilityStatus"
					"Parameters"		"Traits:Agility"
					"HUDName"		"unitagility"
				}
				"StatusIntellect"
				{
					"LocalizeToken"		"#DOTA_IntellectStatus"
					"Parameters"		"Traits:Intellect"
					"HUDName"		"unitintellect"
				}
			}
		}
	}
*/

/*

	//=================================================================================================================
	// HERO: Antimage
	//=================================================================================================================
	"npc_dota_hero_antimage"
	{
		// General
		//-------------------------------------------------------------------------------------------------------------
		"Model"		"models/heroes/antimage/antimage.vmdl"
		"SoundSet"		"Hero_Antimage"
		"IdleExpression"		"scenes/antimage/antimage_exp_idle_01.vcd"
		"HeroID"		"1"
		"Enabled"		"1"
		"HeroUnlockOrder"		"1"
		"Role"			"Carry,Escape,Nuker"
		"Rolelevels"	"3,3,1"
		"Complexity"	"1"
		"Team"		"Good"
		"ModelScale"		"0.900000"
		"VersusScale"		"0.9"
		"HeroGlowColor"		"120 64 148"
		"PickSound"		"antimage_anti_spawn_01"
		"BanSound"		"antimage_anti_anger_04"
		"CMEnabled"		"1"
		"NameAliases"		"am;wei"
		"workshop_guide_name"		"Anti-Mage"
		"LastHitChallengeRival"		"npc_dota_hero_bounty_hunter"
		"HeroSelectSoundEffect"		"Hero_Antimage.Pick"
		"GibType"		"default"
		"new_player_enable"		"1"
		"SimilarHeroes"	"109,12,94"
		"Adjectives"
		{
			"NicePecs"		"1"
			"Parent"		"1"
		}
		"Persona"
		{
			"1"
			{
				"name"					"npc_dota_hero_antimage_persona1"
				"token"					"#npc_dota_hero_antimage_persona1"
				"token_english"			"#npc_dota_hero_antimage_persona1__en"
				"Model"					"models/heroes/antimage_female/antimage_female.vmdl"	// For tools only
				"ModelScale"		"1.05"
			}
		}
		"HeroOrderID"	"1"


		// Abilities
		//-------------------------------------------------------------------------------------------------------------
		"Ability1"		"antimage_mana_break"
		"Ability2"		"antimage_blink"
		"Ability3"		"antimage_counterspell"
		"Ability4"		"antimage_mana_overload"
		"Ability5"		"antimage_counterspell_ally"
		"Ability6"		"antimage_mana_void"

		"AbilityDraftAbilities"
		{
			"Ability1"	"antimage_mana_break"
			"Ability2"	"antimage_blink"
			"Ability3"	"antimage_counterspell"
			"Ability4"	"antimage_mana_void"
		}

		"Ability10"		"special_bonus_unique_antimage_4"
		"Ability11"		"special_bonus_unique_antimage_manavoid_aoe"
		"Ability12"		"special_bonus_unique_antimage_7"
		"Ability13"		"special_bonus_unique_antimage_5"
		"Ability14"		"special_bonus_unique_antimage_6"
		"Ability15"		"special_bonus_unique_antimage"
		"Ability16"		"special_bonus_unique_antimage_3"
		"Ability17"		"special_bonus_unique_antimage_2"
		"Facets"
		{
			"antimage_magebanes_mirror"
			{
				"Icon"				"ricochet"
				"Color"				"Purple"
				"GradientID"		"1"
			}
			"antimage_mana_thirst"
			{
				"Icon"				"mana"
				"Color"				"Blue"
				"GradientID"		"3"
			}
		}

		// Armor
		//-------------------------------------------------------------------------------------------------------------
		"ArmorPhysical"		"1"

		// Attack
		//-------------------------------------------------------------------------------------------------------------
		"AttackCapabilities"		"DOTA_UNIT_CAP_MELEE_ATTACK"
		"AttackDamageMin"		"29"
		"AttackDamageMax"		"33"
		"AttackRate"		"1.400000"
		"AttackAnimationPoint"		"0.300000"
		"AttackAcquisitionRange"		"600"
		"AttackRange"		"150"
		"ProjectileSpeed"		"0"
		"AttackSpeedActivityModifiers"
		{
			"fast"		"180"
			"faster"	"300"
		}

		// Attributes
		//-------------------------------------------------------------------------------------------------------------
		"AttributePrimary"		"DOTA_ATTRIBUTE_AGILITY"
		"AttributeBaseStrength"		"19"
		"AttributeStrengthGain"		"1.600000"
		"AttributeBaseIntelligence"		"12"
		"AttributeIntelligenceGain"		"1.800000"
		"AttributeBaseAgility"		"24"
		"AttributeAgilityGain"		"2.8"
		"StatusHealthRegen"		"0.75"

		// Movement
		//-------------------------------------------------------------------------------------------------------------
		"MovementSpeed"		"310"
		"BoundsHullName"		"DOTA_HULL_SIZE_HERO"
		"particle_folder"		"particles/units/heroes/hero_antimage"
		"GameSoundsFile"		"soundevents/game_sounds_heroes/game_sounds_antimage.vsndevts"
		"VoiceFile"		"soundevents/voscripts/game_sounds_vo_antimage.vsndevts"

		// Additional data needed to render the out of game portrait
		"RenderablePortrait"
		{
			"Particles"
			{
				"particles/units/heroes/hero_antimage/antimage_loadout.vpcf"		"loadout"
			}
		}

		// Items
		//-------------------------------------------------------------------------------------------------------------
		"ItemSlots"
		{
			"0"
			{
				"SlotIndex"		"0"
				"SlotName"		"weapon"
				"SlotText"		"#LoadoutSlot_Weapon"
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"2500"
				"MaxPolygonsLOD1"		"1000"
			}
			"1"
			{
				"SlotIndex"		"1"
				"SlotName"		"offhand_weapon"
				"SlotText"		"#LoadoutSlot_Offhand_Weapon"
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"2500"
				"MaxPolygonsLOD1"		"1000"
			}
			"2"
			{
				"SlotIndex"		"2"
				"SlotName"		"head"
				"SlotText"		"#LoadoutSlot_Head_Accessory"
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"3000"
				"MaxPolygonsLOD1"		"1200"
			}
			"3"
			{
				"SlotIndex"		"3"
				"SlotName"		"armor"
				"SlotText"		"#LoadoutSlot_Armor"
				"TextureWidth"		"512"
				"TextureHeight"		"512"
				"MaxPolygonsLOD0"		"2500"
				"MaxPolygonsLOD1"		"1000"
			}
			"4"
			{
				"SlotIndex"		"4"
				"SlotName"		"arms"
				"SlotText"		"#LoadoutSlot_Arms"
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"1500"
				"MaxPolygonsLOD1"		"600"
			}
			"5"
			{
				"SlotIndex"		"5"
				"SlotName"		"belt"
				"SlotText"		"#LoadoutSlot_Belt"
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"1500"
				"MaxPolygonsLOD1"		"600"
			}
			"6"
			{
				"SlotIndex"		"6"
				"SlotName"		"shoulder"
				"SlotText"		"#LoadoutSlot_Shoulder"
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"2000"
				"MaxPolygonsLOD1"		"800"
			}
			"7"
			{
				"SlotIndex"		"7"
				"SlotName"		"taunt"
				"SlotText"		"#LoadoutSlot_Taunt"
			}
			"8"
			{
				"SlotIndex"		"8"
				"SlotName"		"weapon_persona_1"
				"SlotText"		"#LoadoutSlot_Weapon_Antimage_Persona"

				//
				// NOTE: this slot's budget is overridden via a workshop template!
				//
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"2000"
				"MaxPolygonsLOD1"		"1000"
				"no_import"		 "1"
			}
			"9"
			{
				"SlotIndex"		"9"
				"SlotName"		"offhand_weapon_persona_1"
				"SlotText"		"#LoadoutSlot_Offhand_Weapon_Antimage_Persona"

				//
				// NOTE: this slot's budget is overridden via a workshop template!
				//
				"TextureWidth"		"256"
				"TextureHeight"		"256"
				"MaxPolygonsLOD0"		"2000"
				"MaxPolygonsLOD1"		"1000"
				"no_import"		 "1"
			}
			"10"
			{
				"SlotIndex"		"10"
				"SlotName"		"head_persona_1"
				"SlotText"		"#LoadoutSlot_Head_Antimage_Persona"

				//
				// NOTE: this slot's budget is overridden via a workshop template!
				//
				"TextureWidth"		"512"
				"TextureHeight"		"512"
				"MaxPolygonsLOD0"		"2500"
				"MaxPolygonsLOD1"		"1250"
				"no_import"		 "1"
			}
			"11"
			{
				"SlotIndex"		"11"
				"SlotName"		"armor_persona_1"
				"SlotText"		"#LoadoutSlot_Armor_Antimage_Persona"

				//
				// NOTE: this slot's budget is overridden via a workshop template!
				//
				"TextureWidth"		"512"
				"TextureHeight"		"512"
				"MaxPolygonsLOD0"		"4000"
				"MaxPolygonsLOD1"		"2000"
				"no_import"		 "1"
			}
			"12"
			{
				"SlotIndex"		"12"
				"SlotName"		"taunt_persona_1"
				"SlotText"		"#LoadoutSlot_Taunt_Antimage_Persona"
				"DisplayInLoadout" "1"
			}
			"13"
			{
				"SlotIndex" "13"
				"SlotName" "ability_effects_1"
				"SlotText" "antimage_mana_break"
				"DisplayInLoadout" "0"
			}
			"14"
			{
				"SlotIndex" "14"
				"SlotName" "ability_effects_2"
				"SlotText" "antimage_blink"
				"DisplayInLoadout" "0"
			}
			"15"
			{
				"SlotIndex" "15"
				"SlotName" "ability_effects_3"
				"SlotText" "antimage_counterspell"
				"DisplayInLoadout" "0"
			}
			"16"
			{
				"SlotIndex" "16"
				"SlotName" "ability_effects_4"
				"SlotText" "antimage_mana_void"
				"DisplayInLoadout" "0"
			}
			// NOTE: PERSONAL SELECTOR MUST BE LAST!
			"17"
			{
				"SlotIndex"		"17"
				"SlotName"		"persona_selector"
				"SlotText"		"#LoadoutSlot_Persona_Selector"
			}
		}
		"Bot"
		{
			"HeroType"		"DOTA_BOT_HARD_CARRY"
			"LaningInfo"
			{
				"SoloDesire"		"1"
				"RequiresBabysit"		"2"
				"ProvidesBabysit"		"0"
				"SurvivalRating"		"2"
				"RequiresFarm"		"2"
				"ProvidesSetup"		"0"
				"RequiresSetup"		"1"
			}
		}
		"precache"
		{
			"particle"			"particles/econ/events/anniversary_10th/anniversary_10th_hat_ambient_npc_dota_hero_antimage.vpcf"
		}
		"party_hat_effect"			"particles/econ/events/anniversary_10th/anniversary_10th_hat_ambient_npc_dota_hero_antimage.vpcf"
		"party_hat_effect_persona"	"particles/econ/events/anniversary_10th/anniversary_10th_hat_ambient_npc_dota_hero_antimage.vpcf"
		"showcase_attachments"
		{
			"attach_head"		"1"
			"attach_attack1"	"2"
			"attach_attack2"	"3"
		}
	}*/
