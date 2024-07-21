package dota2

import (
	"strings"

	"github.com/baldurstod/vdf"
)

type Persona struct {
	Name  string
	Model string
}

func (p *Persona) initFromData(data *vdf.KeyValue) error {
	var err error

	if p.Name, err = data.GetString("name"); err != nil {
		return err
	}

	if p.Model, err = data.GetString("Model"); err != nil {
		return err
	}

	return nil
}

func (p *Persona) String() string {
	var sb strings.Builder

	sb.WriteString("Name: " + p.Name + "\n")
	sb.WriteString("Model: " + p.Model + "\n")

	return sb.String()
}

/*
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
*/
