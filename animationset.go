package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type AnimationSet struct {
	Name             string
	controls         map[IControl]struct{}
	presetGroups     []*PresetGroup
	operators        []Operator
	RootControlGroup ControlGroup
	GameModel        *GameModel
}

func NewAnimationSet(name string) *AnimationSet {
	return &AnimationSet{
		Name:             name,
		controls:         make(map[IControl]struct{}),
		RootControlGroup: *NewControlGroup(""),
	}
}

func (as *AnimationSet) AddOperator(o Operator) {
	as.operators = append(as.operators, o)
}

func (as *AnimationSet) AddControl(c IControl) {
	as.controls[c] = struct{}{}
}

func (as *AnimationSet) CreateTransformControl(name string) *TransformControl {
	cg := as.RootControlGroup.GetSubGroup(name)
	c := cg.CreateTransformControl(name)
	as.AddControl(c)
	return c
}

func (as *AnimationSet) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(as.Name, "DmeAnimationSet")

	controls := e.CreateAttribute("controls", dmx.AT_ELEMENT_ARRAY)
	for c := range as.controls {
		controls.PushElement(serializer.GetElement(c))
	}

	presetGroups := e.CreateAttribute("presetGroups", dmx.AT_ELEMENT_ARRAY)
	for _, pg := range as.presetGroups {
		presetGroups.PushElement(serializer.GetElement(pg))
	}

	operators := e.CreateAttribute("operators", dmx.AT_ELEMENT_ARRAY)
	for _, o := range as.operators {
		operators.PushElement(serializer.GetElement(o))
	}

	e.CreateElementAttribute("rootControlGroup", serializer.GetElement(&as.RootControlGroup))

	if as.GameModel != nil {
		e.CreateElementAttribute("gameModel", serializer.GetElement(as.GameModel))
	} else {
		e.CreateElementAttribute("gameModel", nil)
	}

	/*
		e.CreateElementAttribute("timeFrame", serializer.GetElement(c.TimeFrame))
		e.CreateColorAttribute("color", c.Color)
		e.CreateStringAttribute("text", c.Text)
		e.CreateBoolAttribute("mute", c.Mute)

		trackGroups := e.CreateAttribute("trackGroups", dmx.AT_ELEMENT_ARRAY)
		for _, tg := range c.trackGroups {
			trackGroups.PushElement(serializer.GetElement(tg))
		}*/

	return e
}

/*
"DmeAnimationSet"
{
	"id" "elementid" "3492493c-da0d-439c-8b51-8eab47e775a9"
	"name" "string" "tiny_04_1"
	"controls" "element_array"
	[
		"element" "5bb1548e-9693-4cca-9e6d-4d59754e9c37",
		"element" "3183c4da-3306-43e4-89e0-3a18431134df",
		"element" "3a653059-5999-4df5-9f38-49703d6688fb",
		"element" "b57ee740-c139-4f1d-81c7-32093fd1ca94",
		"element" "33315bcd-e149-4fd2-afbc-11a7e1d3093d",
		"element" "300c3d33-0056-4c90-97b3-12fb5baabeb0",
		"element" "6c2b6edc-4187-4d7f-8150-0d3ae715979e",
		"element" "a50c189e-02fb-4166-b124-0c90cb0752c0",
		"element" "62b92bad-d55f-4309-ba32-a58744f37f45",
		"element" "eaad0e7c-9d06-4dbc-a1dc-3d39dedd1878",
		"element" "e85ccea1-1444-4184-926f-8584052976c3",
		"element" "d52cba8b-2f81-4391-91e9-8fb0f5d88bb6",
		"element" "8aa74020-1afb-4071-aa72-8fa6dfbd8e81",
		"element" "987aaf97-3039-4283-aa61-b0703d6520cd",
		"element" "31b09741-17d2-4db7-b6de-eaf862ab6243",
		"element" "a9132063-f1c9-4da3-9da7-e4959edf9b9e",
		"element" "69729a35-662f-41a0-9850-0205616bda9e",
		"element" "4ef7e953-0137-4d3b-a956-0aaaa722e043",
		"element" "7f4de600-810f-4cb3-8a50-e10bf29fce54",
		"element" "4cc44428-00f5-4313-b932-ca00e995c95b",
		"element" "3afb4b51-0867-4a4f-be18-33c677a41431",
		"element" "f0178e15-7731-446e-af85-5d391b03ecc2",
		"element" "a64ef061-ce74-4e78-b7e3-ba303c81f2bc",
		"element" "df57568f-8b1c-4b28-90cf-0ce0a9dfdbb3",
		"element" "fbebb4f0-8804-4dca-b1c0-d580da4db00b",
		"element" "ab6c3edf-81ca-4592-b34a-582735a4ced6",
		"element" "02cc2afd-800c-4606-ad94-b6e7f9706d60",
		"element" "eb7cdfb8-18ee-4af3-a5cd-0bbaa153a6d0",
		"element" "e9f7377f-1f74-4596-a876-6b7002d1ab24",
		"element" "7d8c9465-6548-43c2-b172-96b48bbbea7e",
		"element" "0d6e4934-1373-448e-9b74-270ddc37c4c0",
		"element" "861bc25b-41a7-470f-af89-4f0a56354904",
		"element" "c081a12f-f724-4db9-8911-b449a091c35a",
		"element" "e9614fab-9848-4af2-9cee-2a42ea75b1df",
		"element" "b468cee4-d3d5-4c3d-9268-0c84c3eec3d4",
		"element" "fe98eaed-1bda-49a0-b38d-b0669212eca6",
		"element" "0d6c166a-51a8-4c28-bd66-bf245e920b9a",
		"element" "77a2e0f0-3b02-4ca2-82d0-1e2c232838dc",
		"element" "c28d9218-b812-4f43-a8c0-e575bdd20217",
		"element" "fdf80db3-c912-45d6-b769-82be1bf1713c",
		"element" "db6ead32-99f5-492a-962c-66cb5460da18",
		"element" "76c435ac-ae58-427f-bc33-16f9f41bc8cd",
		"element" "550541aa-907c-4b51-9457-fecc78bd4379",
		"element" "6bc9aafc-70bf-4a18-bd50-cdf704ee6070",
		"element" "b7b04180-f929-4784-8bb8-5ddd4d4dc0f1",
		"element" "aee097af-7d10-4332-af2a-f6907d127629",
		"element" "520ef09a-0021-43e4-af6a-2336ef2c0faf",
		"element" "cf4b92f5-0989-4825-8d7c-5a394c200627",
		"element" "c7e68b8a-aa54-4753-85fa-fda088043501",
		"element" "a856da62-cb58-4ade-ac57-e41582db13a9",
		"element" "025d0a2e-2d30-4570-929b-53356df8cd35",
		"element" "f5eeabdc-9b33-4275-b6eb-9cf9383b339d",
		"element" "e973333e-bcba-4921-aa75-88fb936dcfe1",
		"element" "c7db4d82-b6be-4f5a-8080-b63945639921",
		"element" "312ea458-6722-4cd2-be32-37f8f7cc1059",
		"element" "d04eaa4e-5554-4787-9dcc-7cd01f99fb13",
		"element" "82906182-51ac-42f8-8e43-a3c3ec6d5d1a",
		"element" "37b251a4-f784-492c-bd60-23babedbc74a",
		"element" "7237cc8a-39c5-4581-a890-862a51e9e695",
		"element" "4fa6e8fa-8b5a-4c1c-b471-1f28f9082fe7",
		"element" "4f0e3335-ee03-402e-82ee-04086bd22316",
		"element" "4b116f4b-6ebc-44b3-8ede-f1b3882b1339",
		"element" "6b8412ba-ee89-4b09-be45-0a896ba3c167",
		"element" "91c0447c-9e7e-47ae-bd18-7a1567338d5c",
		"element" "12adfb64-7181-450f-aeb1-86097a919431",
		"element" "c4209f36-c97d-495e-90c6-b403b921167a",
		"element" "07edd5e2-f7ed-400c-92e3-a9d34ef2bc15",
		"element" "0c2af57c-4ba1-4392-9cd0-520daa496d5e",
		"element" "9d0916f5-6db1-4bd3-9019-7c9558ac50a8",
		"element" "35c42a30-eb73-454d-9d53-33ed22e5b321",
		"element" "896ab609-fdea-49db-88fd-647671b1fea1",
		"element" "2a2436f5-23bd-4ef4-bbab-7e6b64c7a8fc",
		"element" "ec7d2445-43fd-45f9-8555-233cb54bce5e",
		"element" "9818d033-658d-4284-9078-c75ee9fab674",
		"element" "0d768d9e-f246-4218-b4ad-8207fef55406",
		"element" "12c3ecfe-105c-4d25-aef9-a4dbc5ee14dd",
		"element" "a331ce41-77f6-4e06-9a49-5659d40cecdb",
		"element" "78a8e701-7077-473c-bb5c-54e112348667",
		"element" "c2dd2e8f-cadc-4523-b3a7-bee6a23fec38",
		"element" "3a9b80b8-7739-4a14-911c-33b090d87054",
		"element" "e6be2007-4616-4e0e-8809-7b0fa52a9f23"
	]
	"presetGroups" "element_array"
	[
		"element" "d947771e-82a3-4992-a8dd-ede0d26183df"
	]
	"phonememap" "element_array"
	[
	]
	"operators" "element_array"
	[
		"element" "195243ee-5973-42b9-9fd3-06b351e861b4"
	]
	"rootControlGroup" "DmeControlGroup"
	{
		"id" "elementid" "30ef327f-a746-42f4-be9f-89c50d944c1b"
		"children" "element_array"
		[
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
			"DmeControlGroup"
			{
				"id" "elementid" "c764fb69-f3a0-450f-9877-747cd04c896b"
				"name" "string" "Body"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "e9f7377f-1f74-4596-a876-6b7002d1ab24",
					"element" "7d8c9465-6548-43c2-b172-96b48bbbea7e",
					"element" "550541aa-907c-4b51-9457-fecc78bd4379",
					"element" "76c435ac-ae58-427f-bc33-16f9f41bc8cd",
					"element" "0d6e4934-1373-448e-9b74-270ddc37c4c0",
					"element" "861bc25b-41a7-470f-af89-4f0a56354904",
					"element" "c081a12f-f724-4db9-8911-b449a091c35a",
					"element" "e6be2007-4616-4e0e-8809-7b0fa52a9f23"
				]
				"groupColor" "color" "0 128 255 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			},
			"DmeControlGroup"
			{
				"id" "elementid" "790f3d3e-44f4-4f63-9a79-020461ccd44b"
				"name" "string" "Arms"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "77a2e0f0-3b02-4ca2-82d0-1e2c232838dc",
					"element" "c28d9218-b812-4f43-a8c0-e575bdd20217",
					"element" "fdf80db3-c912-45d6-b769-82be1bf1713c",
					"element" "db6ead32-99f5-492a-962c-66cb5460da18",
					"element" "e9614fab-9848-4af2-9cee-2a42ea75b1df",
					"element" "b468cee4-d3d5-4c3d-9268-0c84c3eec3d4",
					"element" "fe98eaed-1bda-49a0-b38d-b0669212eca6",
					"element" "0d6c166a-51a8-4c28-bd66-bf245e920b9a"
				]
				"groupColor" "color" "0 128 255 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			},
			"DmeControlGroup"
			{
				"id" "elementid" "e63f73a8-a5a7-4685-8b13-cd78656d503f"
				"name" "string" "Fingers"
				"children" "element_array"
				[
					"DmeControlGroup"
					{
						"id" "elementid" "4d9f8901-1211-45a6-b398-4c0834a01a29"
						"name" "string" "LeftFingers"
						"children" "element_array"
						[
						]
						"controls" "element_array"
						[
							"element" "e973333e-bcba-4921-aa75-88fb936dcfe1",
							"element" "c7db4d82-b6be-4f5a-8080-b63945639921",
							"element" "025d0a2e-2d30-4570-929b-53356df8cd35",
							"element" "f5eeabdc-9b33-4275-b6eb-9cf9383b339d",
							"element" "312ea458-6722-4cd2-be32-37f8f7cc1059",
							"element" "d04eaa4e-5554-4787-9dcc-7cd01f99fb13"
						]
						"groupColor" "color" "64 200 64 255"
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
			"DmeControlGroup"
			{
				"id" "elementid" "27748876-f190-4fb3-afcf-a76c3dc4dac5"
				"name" "string" "Legs"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "520ef09a-0021-43e4-af6a-2336ef2c0faf",
					"element" "cf4b92f5-0989-4825-8d7c-5a394c200627",
					"element" "6bc9aafc-70bf-4a18-bd50-cdf704ee6070",
					"element" "b7b04180-f929-4784-8bb8-5ddd4d4dc0f1",
					"element" "aee097af-7d10-4332-af2a-f6907d127629"
				]
				"groupColor" "color" "0 128 255 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			},
			"DmeControlGroup"
			{
				"id" "elementid" "a3a85ed0-7de5-4508-b8cd-b7c39a181d81"
				"name" "string" "Unknown"
				"children" "element_array"
				[
				]
				"controls" "element_array"
				[
					"element" "c7e68b8a-aa54-4753-85fa-fda088043501",
					"element" "a856da62-cb58-4ade-ac57-e41582db13a9",
					"element" "82906182-51ac-42f8-8e43-a3c3ec6d5d1a",
					"element" "37b251a4-f784-492c-bd60-23babedbc74a",
					"element" "7237cc8a-39c5-4581-a890-862a51e9e695",
					"element" "4fa6e8fa-8b5a-4c1c-b471-1f28f9082fe7",
					"element" "4f0e3335-ee03-402e-82ee-04086bd22316",
					"element" "4b116f4b-6ebc-44b3-8ede-f1b3882b1339",
					"element" "6b8412ba-ee89-4b09-be45-0a896ba3c167",
					"element" "91c0447c-9e7e-47ae-bd18-7a1567338d5c",
					"element" "12adfb64-7181-450f-aeb1-86097a919431",
					"element" "c4209f36-c97d-495e-90c6-b403b921167a",
					"element" "07edd5e2-f7ed-400c-92e3-a9d34ef2bc15",
					"element" "0c2af57c-4ba1-4392-9cd0-520daa496d5e",
					"element" "9d0916f5-6db1-4bd3-9019-7c9558ac50a8",
					"element" "35c42a30-eb73-454d-9d53-33ed22e5b321",
					"element" "896ab609-fdea-49db-88fd-647671b1fea1",
					"element" "2a2436f5-23bd-4ef4-bbab-7e6b64c7a8fc",
					"element" "ec7d2445-43fd-45f9-8555-233cb54bce5e",
					"element" "9818d033-658d-4284-9078-c75ee9fab674",
					"element" "0d768d9e-f246-4218-b4ad-8207fef55406",
					"element" "12c3ecfe-105c-4d25-aef9-a4dbc5ee14dd",
					"element" "a331ce41-77f6-4e06-9a49-5659d40cecdb",
					"element" "78a8e701-7077-473c-bb5c-54e112348667",
					"element" "c2dd2e8f-cadc-4523-b3a7-bee6a23fec38",
					"element" "3a9b80b8-7739-4a14-911c-33b090d87054"
				]
				"groupColor" "color" "0 128 255 255"
				"controlColor" "color" "200 200 200 255"
				"visible" "bool" "1"
				"selectable" "bool" "1"
				"snappable" "bool" "1"
			}
		]
		"controls" "element_array"
		[
		]
		"groupColor" "color" "200 200 200 255"
		"controlColor" "color" "200 200 200 255"
		"visible" "bool" "1"
		"selectable" "bool" "1"
		"snappable" "bool" "1"
	}

	"gameModel" "element" "94de5e75-487b-4aaa-81c2-8de83928762c"
},*/
