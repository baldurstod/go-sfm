package sfm

import (
	"github.com/baldurstod/go-dmx"
	"github.com/baldurstod/go-vector"
)

type ControlGroup struct {
	Name         string
	children     map[string]*ControlGroup
	controls     map[IControl]struct{}
	GroupColor   vector.Vector4[uint8]
	ControlColor vector.Vector4[uint8]
	Visible      bool
	Selectable   bool
	Snappable    bool
}

func NewControlGroup(name string) *ControlGroup {
	return &ControlGroup{
		Name:         name,
		children:     make(map[string]*ControlGroup),
		controls:     make(map[IControl]struct{}),
		GroupColor:   [4]uint8{200, 200, 200, 255},
		ControlColor: [4]uint8{200, 200, 200, 255},
		Visible:      true,
		Selectable:   true,
		Snappable:    true,
	}
}

func (cg *ControlGroup) CreateChildren(name string) *ControlGroup {
	child, ok := cg.children[name]
	if !ok {
		child = NewControlGroup(name)
		cg.children[name] = child
	}
	return child
}

func (cg *ControlGroup) AddControl(control IControl) {
	cg.controls[control] = struct{}{}
}

func (cg *ControlGroup) CreateControl(name string) *Control {
	control := NewControl(name)
	cg.AddControl(control)
	return control
}

func (cg *ControlGroup) CreateTransformControl(name string) *TransformControl {
	control := NewTransformControl(name)
	cg.AddControl(control)
	return control
}

func (cg *ControlGroup) GetSubGroup(controlName string) *ControlGroup {
	group := GetAnimationGroup(controlName)

	if group == nil {
		return cg
	}

	current := cg
	for _, v := range group.root {
		current = current.CreateChildren(v)
	}
	current.GroupColor = group.Color

	return current
}

func (cg *ControlGroup) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(cg.Name, "DmeControlGroup")

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range cg.children {
		children.PushElement(serializer.GetElement(child))
	}

	controls := e.CreateAttribute("controls", dmx.AT_ELEMENT_ARRAY)
	for control := range cg.controls {
		controls.PushElement(serializer.GetElement(control))
	}

	e.CreateColorAttribute("groupColor", cg.GroupColor)
	e.CreateColorAttribute("controlColor", cg.ControlColor)

	e.CreateBoolAttribute("visible", cg.Visible)
	e.CreateBoolAttribute("selectable", cg.Selectable)
	e.CreateBoolAttribute("snappable", cg.Snappable)

	return e
}

/*
	"DmeControlGroup"
	{
		"id" "elementid" "7559839e-0e28-4813-b0a4-7a7beccb3f70"
		"name" "string" "Face"
		"children" "element_array"
		[
			"DmeControlGroup"
			{
				"id" "elementid" "fe3871bc-298e-4b90-9056-adc150708e55"
				"name" "string" "Upper Face"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "5bb1548e-9693-4cca-9e6d-4d59754e9c37",
					"element" "3183c4da-3306-43e4-89e0-3a18431134df",
					"element" "3a653059-5999-4df5-9f38-49703d6688fb",
					"element" "33315bcd-e149-4fd2-afbc-11a7e1d3093d",
					"element" "b57ee740-c139-4f1d-81c7-32093fd1ca94"
				]
				"groupColor" "color" "255 128 32 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			},
			"DmeControlGroup"
			{
				"id" "elementid" "dfc7b835-cc01-4c2d-b988-5a5b7deb7679"
				"name" "string" "Mid Face"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "300c3d33-0056-4c90-97b3-12fb5baabeb0",
					"element" "6c2b6edc-4187-4d7f-8150-0d3ae715979e",
					"element" "a50c189e-02fb-4166-b124-0c90cb0752c0",
					"element" "62b92bad-d55f-4309-ba32-a58744f37f45"
				]
				"groupColor" "color" "255 128 32 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			},
			"DmeControlGroup"
			{
				"id" "elementid" "9f26dc0e-b697-4209-9e25-aeaed4bc5057"
				"name" "string" "Lower Face"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "ab6c3edf-81ca-4592-b34a-582735a4ced6",
					"element" "02cc2afd-800c-4606-ad94-b6e7f9706d60",
					"element" "df57568f-8b1c-4b28-90cf-0ce0a9dfdbb3",
					"element" "987aaf97-3039-4283-aa61-b0703d6520cd",
					"element" "69729a35-662f-41a0-9850-0205616bda9e",
					"element" "4ef7e953-0137-4d3b-a956-0aaaa722e043",
					"element" "7f4de600-810f-4cb3-8a50-e10bf29fce54",
					"element" "4cc44428-00f5-4313-b932-ca00e995c95b",
					"element" "3afb4b51-0867-4a4f-be18-33c677a41431",
					"element" "f0178e15-7731-446e-af85-5d391b03ecc2",
					"element" "a64ef061-ce74-4e78-b7e3-ba303c81f2bc",
					"element" "fbebb4f0-8804-4dca-b1c0-d580da4db00b",
					"element" "eb7cdfb8-18ee-4af3-a5cd-0bbaa153a6d0",
					"element" "eaad0e7c-9d06-4dbc-a1dc-3d39dedd1878",
					"element" "8aa74020-1afb-4071-aa72-8fa6dfbd8e81",
					"element" "31b09741-17d2-4db7-b6de-eaf862ab6243",
					"element" "a9132063-f1c9-4da3-9da7-e4959edf9b9e",
					"element" "e85ccea1-1444-4184-926f-8584052976c3",
					"element" "d52cba8b-2f81-4391-91e9-8fb0f5d88bb6"
				]
				"groupColor" "color" "255 128 32 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			}
		]
		"controls" "element_array"
		[
		]
		"groupColor" "color" "0 128 255 255"
		"controlColor" "color" "200 200 200 255"
		"visible" "bool" "1"
		"selectable" "bool" "1"
		"snappable" "bool" "1"
	},
*/
