package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type PresetGroup struct {
	Name     string
	presets  []*Preset
	Visible  bool
	Readonly bool
}

func NewPresetGroup(name string) *PresetGroup {
	return &PresetGroup{
		Name:    name,
		presets: make([]*Preset, 0),
	}
}

func (pg *PresetGroup) AddPreset(p *Preset) {
	pg.presets = append(pg.presets, p)
}

func (pg *PresetGroup) CreatePreset(name string) *Preset {
	p := NewPreset(name)
	pg.presets = append(pg.presets, p)
	return p
}

func (pg *PresetGroup) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeChannel")

	e.CreateStringAttribute("name", pg.Name)

	presets := e.CreateAttribute("presets", dmx.AT_ELEMENT_ARRAY)
	for _, p := range pg.presets {
		presets.PushElement(serializer.GetElement(p))
	}

	return e
}

/*
"DmePresetGroup"
{
	"id" "elementid" "d947771e-82a3-4992-a8dd-ede0d26183df"
	"name" "string" "phoneme"
	"presets" "element_array"
	[
		"DmePreset"
		{
			"id" "elementid" "45aecc03-3eba-4eb8-827f-9de57ea0472b"
			"name" "string" "er"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "568d201f-9880-4e3a-9807-a9f3c0aa993d"
					"name" "string" "dimpler"
					"value" "float" "0.1400000006"
				},
				"DmElement"
				{
					"id" "elementid" "4649ed39-685e-4a08-bffa-b5c835a096ad"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "fa52996d-a0b7-430f-8d01-2221db3fd2ad"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "fa45a511-3177-4f39-84d1-339bee6fc818"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "67e153dd-40cd-4ccf-b5d8-ffdfd6ce3f48"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8160f1de-5e1c-48ef-bfd0-d03701097617"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "48559f42-8632-4b21-8994-a5dab547d01d"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7a0e3fae-f832-4df8-b802-75d99ec66555"
					"name" "string" "lipPuckerer"
					"value" "float" "0.4799999893"
				},
				"DmElement"
				{
					"id" "elementid" "59957dc5-5818-4305-812a-4dd6c31b1f26"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.5400000215"
				},
				"DmElement"
				{
					"id" "elementid" "802ab04c-382d-4deb-9be8-fa98f311c653"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.6299999952"
				},
				"DmElement"
				{
					"id" "elementid" "6c34119c-0206-40e4-a163-c6e713af5552"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e51e4d3a-d971-4bec-b96b-1dc0bbbd4b75"
					"name" "string" "jawOpen"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "886dc5a9-ac63-4db0-bf31-7310b7ffe2dc"
					"name" "string" "alt_dimpler"
					"value" "float" "0.1400000006"
				},
				"DmElement"
				{
					"id" "elementid" "afa78749-c7df-4d68-8b20-ef8d59c9c7b1"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "952d07a3-d777-47d8-8ad6-0122451a155a"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "e9cd57d4-3f0d-4554-977d-ffd6af92054f"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fa4f63b8-eba9-4452-b57e-618231dedace"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "70bfc919-458c-4f99-9082-7f8c6419e12d"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ca769bbb-003c-4755-b545-010219dc7f8e"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3e7ed4f5-3aa4-4f5e-83ce-db751d1ae429"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.4799999893"
				},
				"DmElement"
				{
					"id" "elementid" "bd0511f0-0158-451f-844b-564f4232b8a4"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.5400000215"
				},
				"DmElement"
				{
					"id" "elementid" "f4eedeb1-d52d-46fc-8766-e65756b32dc0"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.6299999952"
				},
				"DmElement"
				{
					"id" "elementid" "a4f877f4-3363-4807-a6fc-d17b4c1ed47d"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "17fb79e8-ca52-41bd-b0b4-aecad9df8ce1"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "b890c642-9bc0-403c-a55c-555734eef460"
					"name" "string" "tongueUp"
					"value" "float" "0.8000000119"
				}
			]
			"description" "string" "URn : rhotacized schwa"
		},
		"DmePreset"
		{
			"id" "elementid" "4a4a8ee9-e323-4d7a-90db-d32f9abcb4bd"
			"name" "string" "m"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "5fde6a4c-14fe-4284-8705-592b020b93fc"
					"name" "string" "phonemeBMP"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "5446b95e-c151-425d-b0ab-239b8fc0e1db"
					"name" "string" "alt_phonemeBMP"
					"value" "float" "1"
				}
			]
			"description" "string" "Mat : voiced bilabial nasal"
		},
		"DmePreset"
		{
			"id" "elementid" "0c97c95a-e570-4787-bc65-ed94367b2df5"
			"name" "string" "p"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "f4411fc5-beaa-4aa5-9266-d8cf5bf09be6"
					"name" "string" "phonemeBMP"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "61cfad73-c99a-4a75-bf65-3473a807fa56"
					"name" "string" "alt_phonemeBMP"
					"value" "float" "1"
				}
			]
			"description" "string" "Put; voiceless alveolar stop"
		},
		"DmePreset"
		{
			"id" "elementid" "98c35ca2-71f7-4f1b-b308-8bf1e5213fb7"
			"name" "string" "w"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "9128d8fa-868f-43cf-88b4-d2648010ccfb"
					"name" "string" "dimpler"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "e3af8671-f003-476c-ba19-ae61e34dabb3"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "d915068e-5682-48d8-a3e4-3b88345ac756"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2fe53023-936e-4785-8073-58799245b146"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b76ebc19-046e-4565-9618-15cd91754498"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "571ff245-63f7-45b6-ade6-975b92494fec"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1be17593-03a3-47d3-a478-1adfa486d454"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9351aab3-6d2a-45bf-b744-3a55e6e0e2c1"
					"name" "string" "lipPuckerer"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "3b7a0865-3d8c-41aa-9a88-3c28fd89a233"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "99c2b124-7775-4b58-ae2e-52e3163616c2"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.6600000262"
				},
				"DmElement"
				{
					"id" "elementid" "f112ce66-1490-4f4b-86a0-bb3f0cd68b57"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "710e8368-1ff1-4d01-9a1c-8823e2ed14e8"
					"name" "string" "jawOpen"
					"value" "float" "0.0799999982"
				},
				"DmElement"
				{
					"id" "elementid" "1aefb369-788c-4189-8828-f0e81baa8c20"
					"name" "string" "alt_dimpler"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "a2ac3150-7c34-4385-ae97-16fe25b92f2c"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "726bc931-e635-4228-a1c3-7c0482c722f8"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "79220109-9830-41b9-ba69-ad176cabd069"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b655d401-6a22-4085-ba8d-66a75c369b1d"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7ede2aa9-4e48-4d24-8f5f-9e36175556f3"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "64bdf36a-cae0-4397-975b-cfc4f34d889a"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d03a79b1-deba-4de0-bc1e-5bd6c15a613e"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "eb76cc34-6865-466b-a8c6-4ada98a364ab"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "fb397d15-b713-4c19-9241-579cfbcbb9f7"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.6600000262"
				},
				"DmElement"
				{
					"id" "elementid" "9f284f7b-76a0-468b-b0f2-05296fd5a13e"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "93f9044a-4fb6-4664-b6d2-56c7ecb8bf5c"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0799999982"
				}
			]
			"description" "string" "With : voiced labial-velar approximant"
		},
		"DmePreset"
		{
			"id" "elementid" "1b092cfd-3149-4e2f-8c3c-e1c597d338b3"
			"name" "string" "f"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "079e84cf-5d8f-4187-8ba2-8e49379dd897"
					"name" "string" "phonemeFV"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "fb765f2d-4c6d-45fe-a86a-e97d47afc2ab"
					"name" "string" "alt_phonemeFV"
					"value" "float" "1"
				}
			]
			"description" "string" "Fork : voiceless labiodental fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "83df5d30-e128-4bc9-8616-d7806a7a3884"
			"name" "string" "v"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "41e46e0b-e304-4325-b4b3-2be0a36d31ee"
					"name" "string" "phonemeFV"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "861363eb-e6b7-4bf8-a5e5-d118c3cd82d7"
					"name" "string" "alt_phonemeFV"
					"value" "float" "1"
				}
			]
			"description" "string" "Val : voiced labialdental fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "ddfbdac1-3fb4-4d3b-a129-a8a9880306a7"
			"name" "string" "r"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "621fa343-3449-48bc-aee6-2ad444cc3e70"
					"name" "string" "dimpler"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "150812bf-90be-4d22-a468-3c5c39a6d7f5"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "15d3e10e-8f96-4338-b91d-e2552eb9083a"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "37affbff-6395-4c5f-9cef-dc7a932a3e7a"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "45331a2f-b1b9-4aa3-9b90-fc1174a521b1"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "73bc53f0-27b4-4fd7-b7ba-fce2748aebaa"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e5d4a8ed-0bd9-4725-a782-2a885ea0410c"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d80b8d59-2dfa-4570-831f-91b7fe8eae78"
					"name" "string" "lipPuckerer"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "b08d508e-7568-4e2a-b679-1bb013c39430"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "6489ade5-d422-40d7-a6e5-4f19088de5e3"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.6600000262"
				},
				"DmElement"
				{
					"id" "elementid" "4767f228-2235-4e26-8c66-bdd7527044da"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ade8403e-7bfa-4508-923d-dee6a4096e47"
					"name" "string" "jawOpen"
					"value" "float" "0.0799999982"
				},
				"DmElement"
				{
					"id" "elementid" "2b8bc989-200b-49a6-8366-79839e8b1d5d"
					"name" "string" "alt_dimpler"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "af454331-9e38-4ce1-b520-fec9b9bfe79a"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "0c8df03c-297a-4dc4-b4de-9c9f68e92211"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "77930f2a-b969-422c-a11a-e6d362bf601c"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bd0f1e12-0c43-4cf3-af1a-4b12f1af0bb3"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "996ef442-3ae2-4e5b-a950-243945e6be7a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4efdea94-8ab8-421f-b780-f6ba94511a46"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ad495182-6bef-447a-9b3c-e697d1eb9b66"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "9133c988-f7a2-40ca-afde-ec218661e86f"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "32c83b8c-c64a-4600-84c1-0e9f89903075"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.6600000262"
				},
				"DmElement"
				{
					"id" "elementid" "e1f5c109-f2ac-4743-a9d9-a536fafe845c"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dd26257e-55e0-48dc-a2ef-e36739f76d26"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0799999982"
				},
				"DmElement"
				{
					"id" "elementid" "87f3cb5e-f026-4ee6-927e-8de194148e33"
					"name" "string" "tongueUp"
					"value" "float" "0.8000000119"
				},
				"DmElement"
				{
					"id" "elementid" "3c92713d-5e0d-4acd-b1c0-6f2ab9735fda"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.1000000015"
				}
			]
			"description" "string" "Red : voiced alveolar approximant"
		},
		"DmePreset"
		{
			"id" "elementid" "0b8e5bf5-4a1a-4df1-b9c8-975ba735842e"
			"name" "string" "r2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "5e8403f6-4518-4541-90f5-d57870d04722"
					"name" "string" "dimpler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "8741a88c-31ab-4bc2-b71b-0ab552fb8cfb"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.4799999893"
				},
				"DmElement"
				{
					"id" "elementid" "87a44f8b-e1c8-42ab-8c87-64bafebc08cf"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "d9041304-90c9-42cd-80c7-f7565eb6a6e9"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b3a2edb8-88fa-43ba-809c-d8543d2ba868"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2144084f-cab0-451d-9b76-ede3888e972b"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "383efc48-2ebb-4e28-b754-158f2dfb4692"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "856b021f-2b81-489f-a2de-fa131b08f3b2"
					"name" "string" "lipPuckerer"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "1d66387f-a945-4649-b217-564dcdb344c8"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.8899999857"
				},
				"DmElement"
				{
					"id" "elementid" "f996f2af-e5c8-4c1e-a984-e22998dd03f5"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.8299999833"
				},
				"DmElement"
				{
					"id" "elementid" "60bf97e2-4a27-4a4c-bae1-5c665d9118a3"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8e92e843-5f34-497b-a77e-6013c5a769f0"
					"name" "string" "jawOpen"
					"value" "float" "0.150000006"
				},
				"DmElement"
				{
					"id" "elementid" "a3cd6e8b-2dca-4512-aa5d-d14d6e8e8c2b"
					"name" "string" "alt_dimpler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "74faf8d4-5ae7-4388-a514-39c4f6e60f73"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.4799999893"
				},
				"DmElement"
				{
					"id" "elementid" "3cb4752a-c724-44dc-817a-edfcad8590d4"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "b8a6afa6-76b1-4590-bcd3-509e3cca146c"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "badf7a63-905a-4dd6-b40f-e9d2f3d6a879"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7e3a8a1-c4ec-4e6d-9bb2-f6ec3171bf34"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7c8a4521-1339-46a8-a020-176473838bf2"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f5be3077-4453-4de8-aa92-268ddacb2402"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "43646344-37c3-48fd-afe4-7c4a4043835e"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.8899999857"
				},
				"DmElement"
				{
					"id" "elementid" "232ff8c9-7268-433e-83c4-2ba174f173b9"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.8299999833"
				},
				"DmElement"
				{
					"id" "elementid" "b112dae7-7d4f-41c1-81e0-7aaf174c7f90"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e86fb424-12de-429d-b3a1-cf105bc96f14"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.150000006"
				},
				"DmElement"
				{
					"id" "elementid" "18f77a27-1407-40a6-96c2-b69904497a9c"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.1000000015"
				}
			]
			"description" "string" "Red : voiced alveolar trill"
		},
		"DmePreset"
		{
			"id" "elementid" "66036ecf-4394-429b-bf0b-ed5a5db89a0f"
			"name" "string" "r3"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "b4ece02f-06c6-4da7-8475-c785e30ea380"
					"name" "string" "dimpler"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "8631230c-8237-4337-b1eb-f67163462960"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "34f83847-4dfe-457e-9428-0e9720d34668"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d544313b-d41f-41a5-af3f-3faf306f31b9"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dca79c33-06aa-4d95-b97e-5f8a2ce4d4be"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4ee03840-a222-49a4-b395-a010d3268b3e"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "98d46148-c8c3-4627-8718-c5970fabf7d3"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d1df9d7c-b654-4d99-bf5f-2f18ae264a8a"
					"name" "string" "lipPuckerer"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "fe9ba1f6-768e-4512-9f68-7c6e761c7a78"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "c753f6a2-b1fd-495c-8d2b-cc7100d310cc"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.6600000262"
				},
				"DmElement"
				{
					"id" "elementid" "c4ab25e1-cf43-4505-bafc-b906695e8047"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "20eb7565-5dbc-40bf-b59d-7cc108c4de39"
					"name" "string" "jawOpen"
					"value" "float" "0.0799999982"
				},
				"DmElement"
				{
					"id" "elementid" "df243883-c354-4473-97ce-4033788c2b15"
					"name" "string" "alt_dimpler"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "74e86602-d7f9-4d73-ae04-59c3fd32f0de"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "1a5330f0-4cba-43c6-95d6-e62b1d925af2"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "66e60cf1-6f88-49ba-a461-160a380229d0"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b93ed345-8ed3-4d56-b24d-364c3d49d949"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e7eadbed-78b5-4abb-9ce3-d2ea6d00447c"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0299738b-9ce0-4657-9a0d-d447ff403f67"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6bcdace5-01fc-4241-9e22-2564cb2f3156"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "40ce1d58-b6be-41ec-aa96-9a0279a10088"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "ef1afed1-9c89-4121-8655-c70b4340bec6"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.6600000262"
				},
				"DmElement"
				{
					"id" "elementid" "ff55dafa-71b1-4fbd-b6d9-9930946616d0"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "aacda5b8-ce1c-49bb-89d7-6a4499ebb059"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0799999982"
				}
			]
			"description" "string" "Red : voiced retroflex approximant"
		},
		"DmePreset"
		{
			"id" "elementid" "257db13d-4f73-48df-b52e-2577e1d3a7c3"
			"name" "string" "<sil>"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "9b8f8156-240d-4642-a95d-67a867266260"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9e94e171-541d-4262-bbf4-44ff0fc4040d"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5e67d3a6-ccd2-4776-a646-e778fc72d32a"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3898adfe-071c-4f5a-9bc8-72b2c83fcf7c"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "29705bc4-b14a-4cf7-b22c-6183486a96f7"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fddcf4ed-0eb6-45b1-ab32-e44e599e56b9"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "20b125dd-16d4-4d32-8c57-358b335e0428"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7807e538-75bd-497b-9b9b-0a494b6f83ba"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "eca397df-baa8-49b0-a8c4-40386e514ec2"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "93d358be-1d33-4c36-8f78-5fa1297bcc33"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f2b88209-6123-4215-b2ac-b638573c3317"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "534a89ed-340a-441e-bf54-69765421acc9"
					"name" "string" "jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "d8bd3d5a-be90-4d5f-b569-b6eca3dade2d"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "acaddf09-d811-4e33-9339-8a9ce953a681"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c01d02e9-5525-4214-bdfc-8a1955ecf99e"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6cc6dd4f-1888-4ab1-92d1-24d77fbcdfc0"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f15b5405-766a-4364-a2fe-48f04c4fc2b9"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "85de69d7-7119-4ca0-beb0-50e2fcb5b029"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6c527a78-7ff0-4ee3-9b98-3a7005007164"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "35324e18-8260-468f-beb8-00e4da5222b6"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "365bd7f3-e302-49e0-a234-2f454e19b07f"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5f0f8d09-cf2c-4138-9240-1ca72de8a591"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d26bfa24-2c9f-4c11-8196-72f3659c4bf2"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7f6e5d53-ae02-4e64-99e0-7bba754e7338"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1000000015"
				}
			]
			"description" "string" "Silence"
		},
		"DmePreset"
		{
			"id" "elementid" "30f17937-fe99-433d-bd87-6bbd86b358eb"
			"name" "string" "er2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "ec16704c-892b-43f3-8b1c-2028beacd22b"
					"name" "string" "dimpler"
					"value" "float" "0.6299999952"
				},
				"DmElement"
				{
					"id" "elementid" "ac205926-3d33-4b46-8607-d3d71bc146f2"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "a36786f2-cfff-4f2d-861d-57cc4bf9b029"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "-0.4199999869"
				},
				"DmElement"
				{
					"id" "elementid" "70849b33-b6b4-4a18-bf38-842903c5830d"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c5378ced-bf94-4320-a7af-ff7bbe33d7b0"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a1851ad1-1dd0-4fb5-b351-53f6f3414cc0"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e34c006f-0f57-48e7-9f80-f54b7c95cde1"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "34daeb43-2cbd-4bdd-9b66-35479b788320"
					"name" "string" "lipPuckerer"
					"value" "float" "0.6000000238"
				},
				"DmElement"
				{
					"id" "elementid" "67699f2c-eb92-49fd-bf68-acd656718a6e"
					"name" "string" "upperLipFunneler"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "7a58cc58-e979-48d3-a0fb-b33153693623"
					"name" "string" "lowerLipFunneler"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "07a43c5e-c2d0-4107-88c0-9656a35c941c"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "baeb86e3-7cfb-4313-9c87-c399b5ee0925"
					"name" "string" "jawOpen"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "e74988e8-3509-4767-becc-87dcb4fc0ed3"
					"name" "string" "alt_dimpler"
					"value" "float" "0.6299999952"
				},
				"DmElement"
				{
					"id" "elementid" "0d3f02ee-f0a6-470d-93d2-1e53aae6e6fb"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "cbd93fd2-099c-4231-922a-a3708385796d"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "-0.4199999869"
				},
				"DmElement"
				{
					"id" "elementid" "6e6b25e2-034d-4b93-bba1-2a5e7e51fa8b"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3a3d4b26-7600-4d58-aaf2-21152a0f8aaa"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "67c63504-626b-4b70-b0e9-5eba58457181"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5263b197-576c-4d98-a5fc-f33a5d05339d"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3002d1b8-56a4-4d9b-8d72-3cb1af79e2e6"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.6000000238"
				},
				"DmElement"
				{
					"id" "elementid" "9acb0fe5-33f3-4ef8-86d0-fd2bbb2b22d9"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "94b4b2f1-e642-4a9b-80cf-9726ff39532f"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "8b38cf38-1fd2-4885-a2c3-d28fd365f354"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ccf0a466-7de3-4f32-ac1e-dfbfc6fd03a6"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2300000042"
				}
			]
			"description" "string" "URn : rhotacized lower-mid central vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "d95c057b-6b1a-4e14-b307-fe09f42e0b96"
			"name" "string" "dh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "156424ec-b35b-40b8-bc0b-883694ae420c"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "85b160fd-5fdf-42f8-be82-3bac8225dc4d"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f9c62c67-4dfa-49b1-b988-68824ffd743c"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "725ab383-6193-466e-ac0d-beb19f7fbdb7"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e48a02ed-c6af-4ad1-94f2-093ddb798273"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e11e7ee6-b30f-401d-b541-2fb544edcb7a"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f45c60da-8d8a-49f9-b1d9-c47ed52b67bc"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6262fbf5-c975-42cf-ae8f-eb4f7687028c"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b8625322-7157-4dc7-883b-395bf9714a44"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "295a6d36-65a6-4778-9153-457b5b0f9365"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "913186d0-ab35-4bc0-841d-afb24d16d361"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "9f5c5817-d397-41e7-b792-d56c372c57bd"
					"name" "string" "jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "5ecc7947-b6ad-47cf-80da-752b222f686b"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f449dbc5-0186-46d6-9674-96da17e935d1"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d1b49b1d-ef67-460e-9492-e420654e1eb2"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "aa5552a9-45bb-45fe-a280-867a9c07598b"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f90f2f0f-d80f-4a4b-8c8f-ba7f6f551720"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "96255c2d-2908-428d-8e24-f98028527fab"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e4dbd19f-c372-415e-bfdd-dc7ed72e4370"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8dbc62d1-c88f-4805-9bc5-11cf1348c37f"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f71c6bb3-73ea-4234-a7d9-d2ea0fa8b767"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8a386726-4cfa-41b8-92f8-f1b60e73c0de"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cb217b47-af73-4438-ad0c-6886467852c6"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "6f2ead69-2e00-485e-9f51-8076ae9c6e8d"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "b7cf2c13-ce9f-4e35-bc0c-d135371a1d9b"
					"name" "string" "tongueUpThrust"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "505ce208-bd98-496c-a60f-095977969a4e"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "THen : voiced dental fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "97b1db6a-ca70-4d64-9893-3f524a52b62e"
			"name" "string" "th"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "3f457e42-ed65-46a9-88c2-e41b889f6c08"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "50c36792-37a5-4883-9212-3574093bbb75"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7de649d5-0726-4108-8873-51399eec391f"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "c1f9018d-5eb9-47a7-98e2-5ea7017dc171"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b9595460-bafe-43b0-bcf7-77fa83d3d3a9"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a152a3b7-5ee2-40f0-b4d4-ebe3fb2a0641"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c155217d-e843-4e61-94dc-0ff9de5afc7b"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "867691c1-bc11-4a79-a1a9-b5c577d575cf"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "24863a70-af35-439e-8c3c-76c2e90f513c"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e53c39e3-06d7-413c-9a11-a642d0aa0ffc"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8e07c9f5-2747-4046-a15d-6358f7e19ebf"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "615ebce4-d6df-4492-8885-0877c2ed4740"
					"name" "string" "jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "4ffe1a71-73e6-4019-82f2-cb25ac7b974d"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7e15ab40-8657-4c01-acd7-2296d0e17540"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f9528ab8-a490-4c13-8dcd-799f0529972f"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "d9ebdb17-7128-4e7c-bd08-0cf6d568b542"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d24ad7ff-1ecd-428f-96e9-e1694fc3e1a5"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "264fc743-2d0d-498c-91a2-25c2835435d0"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "18823558-a0df-4265-8d6a-14ca90046db9"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "717e621c-d043-4ee5-814e-02ff492f7611"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "57fba446-27a9-4024-92d2-b5f30bd128c7"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c9ab14c5-d934-4a4b-a75e-2adfeff55ec8"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1cdc64ec-c65e-410d-b7db-de7fcde0b17f"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3daec2de-514f-4424-b54b-8caa1bbb59e8"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "83f71da3-a6de-4f06-93f7-6aec4735c2e4"
					"name" "string" "tongueUpThrust"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "9d141602-9cba-4468-a60a-74f13aa3a5ce"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "THin : voiceless dental fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "c137d07a-9d97-4a8c-a52b-dc40244f2ec0"
			"name" "string" "sh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "1d855868-06e3-4ede-9d57-ebafa0c0352f"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fc3eb000-9994-4fca-b2d6-935cf7c7ec49"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "5e36232a-cbc6-4697-963d-706e8105b1d2"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "c5949b7a-9434-4999-bbb2-a08d04602e69"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a9fbd8f1-27da-4f55-8295-ab1b3ac1004b"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4f623393-824b-4d2b-8f51-812a11bc664c"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e1df47d4-2f2c-411c-9a69-3037b3e13fd6"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fe3b5ffe-7c84-42b3-8638-6c38f39f243b"
					"name" "string" "lipPuckerer"
					"value" "float" "0.5"
				},
				"DmElement"
				{
					"id" "elementid" "03736889-36d1-43a9-8907-b6fe51898503"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "101e5d9a-f419-4e1b-9c2a-941d852c25b2"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.4399999976"
				},
				"DmElement"
				{
					"id" "elementid" "cd0a44d7-a989-4776-931d-ffede0f7bc25"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8d6d9d25-973e-4266-8d53-d566afe05fee"
					"name" "string" "jawOpen"
					"value" "float" "0.0099999998"
				},
				"DmElement"
				{
					"id" "elementid" "79487c01-0267-4b22-8869-2f6453c0338d"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5c991912-3cb2-492c-a714-aa8e9d21fcb0"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "63721e5d-aa6a-4b1c-8d8e-eefc45ef784b"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "f84a4245-f11b-4eba-b5c2-0e1eea464f48"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "92362cb4-c83b-4ad1-95f4-ecd4813d7270"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "19051f1f-1c9e-4a7b-97b9-3e44184b2c0a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "63505115-64bc-4d5b-841a-ac045046c48b"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1240c8ae-9bd0-4ced-b7ce-1667341fe430"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.5"
				},
				"DmElement"
				{
					"id" "elementid" "9874c136-0ee4-4da5-8313-abaa58e344cd"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.6399999857"
				},
				"DmElement"
				{
					"id" "elementid" "fe3fb7dd-f6f4-41e9-adb8-5418266a595b"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.4399999976"
				},
				"DmElement"
				{
					"id" "elementid" "85e99efe-6692-4878-8c37-db9f599f559c"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5f0b7d8c-6c49-4e6d-bad0-f1ed1f27103b"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0099999998"
				},
				"DmElement"
				{
					"id" "elementid" "dd31a739-5fd1-44d0-ae21-b06816e82cf9"
					"name" "string" "tongueUp"
					"value" "float" "0.8000000119"
				},
				"DmElement"
				{
					"id" "elementid" "2f6170c4-b7a7-4c6d-a33e-10b4636108a5"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "SHe : voiceless postalveolar fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "d9ca530d-f025-4a56-969e-7e4e3bcfa3a7"
			"name" "string" "jh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "d70c4a56-549d-47c9-a602-fceba1431ad1"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "85cf3ebf-f0e1-4866-828f-1c63ae0c87a6"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.3000000119"
				},
				"DmElement"
				{
					"id" "elementid" "1a800a09-0e70-4379-8214-7afe51bbc3e7"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "fc19e2fa-cb89-4a25-ad34-ebe0a8f3e825"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fb47579c-636a-40fe-8d3e-3652c951cdd7"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3c30ff40-7aec-40a2-8e18-0c44dd8b0039"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "96ec31d4-9c0e-4a57-83f7-957cae673d2d"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0efc5bb4-aa80-4798-be96-ba94780cf85a"
					"name" "string" "lipPuckerer"
					"value" "float" "0.2899999917"
				},
				"DmElement"
				{
					"id" "elementid" "8871be7c-e0fd-4b26-937e-30f12254734f"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.2099999934"
				},
				"DmElement"
				{
					"id" "elementid" "e38f681e-040c-4cf6-84df-11d82c0de61a"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "d0463458-1802-44e1-91af-fcffbd301faf"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "eb4bef7e-f7b3-43ed-84fa-9002e230e853"
					"name" "string" "jawOpen"
					"value" "float" "0.0900000036"
				},
				"DmElement"
				{
					"id" "elementid" "a1900d7c-53b7-49a7-b330-e47b8d513b0e"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "db953228-8c54-4805-aaea-333f49ad956c"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.3000000119"
				},
				"DmElement"
				{
					"id" "elementid" "8f097c5d-0be2-4c3f-a917-e142cec60d40"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "f01b6870-33ac-4f20-88d6-c4f246f55fdb"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9da5eb83-28be-4ba7-9c09-649444b9abd7"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9bcdf309-26b9-48af-acc4-cb9063498b5a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "220d1672-c99e-495a-bd5f-b1589e240509"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "db637f52-dab5-4e23-9f19-1015e862efd6"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.2899999917"
				},
				"DmElement"
				{
					"id" "elementid" "55ceb9a5-c593-4312-a99f-241e27d3d924"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.2099999934"
				},
				"DmElement"
				{
					"id" "elementid" "cf89663a-6f3c-4bee-9e4a-e01a673b39b4"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "31da144e-295e-40b1-ba2f-6441ff99e5c8"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0e9b1591-ed8d-4315-a565-b61d8869e49f"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0900000036"
				},
				"DmElement"
				{
					"id" "elementid" "b784ebd1-55a6-4978-a1af-88237b73daa3"
					"name" "string" "tongueUp"
					"value" "float" "0.8000000119"
				},
				"DmElement"
				{
					"id" "elementid" "4e826dc0-324d-4908-a021-0e9f1cb75f64"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "Joy : voiced postalveolar afficate"
		},
		"DmePreset"
		{
			"id" "elementid" "38188ae8-0bf8-4f32-bb27-fe6a9ef5c59d"
			"name" "string" "ch"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "f58189a0-aee9-4369-850d-2a88992e8cc9"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c61e57dc-3d6b-4b9b-b03e-c2098e2d3412"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "d613cc0f-5f76-418d-a3bd-db0b7c3ae405"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "397bfaab-eefd-4352-aa97-468cc63f1435"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7bfcfe4f-94a5-403e-814b-bf31d523c7b1"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9b3b199b-ec2f-4397-846d-dd3bf68cae30"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a5f205f1-a375-409a-b30d-aafdc96ee8c9"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "11d7b862-03bf-418e-b560-95b55522044a"
					"name" "string" "lipPuckerer"
					"value" "float" "0.1400000006"
				},
				"DmElement"
				{
					"id" "elementid" "e2f24ba5-9bb8-4e2e-906b-c9bc6f3e06f3"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "f2e81af4-531d-487d-8cf9-db5916af65ad"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.2899999917"
				},
				"DmElement"
				{
					"id" "elementid" "2c5ab620-7a5f-4ebe-99bc-33a1d8791359"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f8bbaf8d-b6ea-44d8-bc40-7cddf94d633e"
					"name" "string" "jawOpen"
					"value" "float" "0.0099999998"
				},
				"DmElement"
				{
					"id" "elementid" "96bda9da-3f70-4c9b-b71e-52d33aecf80b"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "55805d64-0e46-4cfd-9f98-187eda2de006"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "fa926d69-2528-49aa-b9e1-6455fb26fe4b"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "b22b1ffc-c26e-42d9-a0fc-4c3def50f25e"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1d3a8345-43d8-43a3-8ba1-b1cd8eb4aff0"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5d1d3fde-233f-49e1-a42a-9852430285cf"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f97ab60e-894e-4f80-b2e6-b711086f5b9c"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "091d10d5-9d97-4ce1-80e5-a8b31632c99c"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.1400000006"
				},
				"DmElement"
				{
					"id" "elementid" "64b72d52-74ba-403a-86f1-5035bd8fea1a"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "519c3a00-7715-4168-b72a-d651f97769af"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.2899999917"
				},
				"DmElement"
				{
					"id" "elementid" "422379fe-4502-419a-97c4-c8776dad8e3c"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "17791cb9-4458-477f-a9f8-bb9133f83004"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0099999998"
				},
				"DmElement"
				{
					"id" "elementid" "80aea9c0-30ce-46e3-a380-104a2dcd6665"
					"name" "string" "tongueUp"
					"value" "float" "0.8000000119"
				},
				"DmElement"
				{
					"id" "elementid" "33081fee-2d3d-4fcb-984a-5d677ce34a78"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "CHin : voiceless postalveolar affricate"
		},
		"DmePreset"
		{
			"id" "elementid" "a049293f-b45b-4db7-b3e8-2510d8f23a39"
			"name" "string" "s"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "20957914-b909-44ac-9639-5fe64cc87d8b"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fadd4169-3e95-404c-9468-0af4652b41f4"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d58cad02-9da3-4ee3-bdda-4a2fcfeb3150"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4000d397-6e02-4434-86f9-724424b019bb"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2199999988"
				},
				"DmElement"
				{
					"id" "elementid" "67319c54-08b0-4368-b450-9cb4acd22269"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "118fd888-e6da-4516-aaa2-f2a6d93d914b"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c3b9ee29-b8fd-4dc5-9daa-ba6adaab3a92"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "04acbdb1-bdac-46a4-992d-ebc19ec4c409"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fb98873a-1e03-4150-b618-74a4be459734"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "14148801-3fdc-41fa-8548-00735bc910ae"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c53831c3-4786-4039-b920-c2c3a312f8cb"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1165b789-e22c-4224-95b1-ed0975df31e7"
					"name" "string" "jawOpen"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d39a43d7-15cc-4a73-a4cb-f221bf40db26"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bd6630f8-537f-47fc-8041-e732e84bab3a"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "04be232d-a16a-42f4-a277-79f20009e66b"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7aeb1dbc-7659-4402-bdfb-f9921eca53b7"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2199999988"
				},
				"DmElement"
				{
					"id" "elementid" "f07f6b9f-3d08-4ead-b537-5d3f2b022f26"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d6517b0b-e8ff-4f03-8a13-72f1a48042ab"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6ead7a9c-e7d4-41ff-8af2-f42551a67fc4"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ef98df42-058d-45d8-a7dc-bc2e9abb92b9"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2c346765-536f-4593-974f-d1ba7f4bb1f8"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c30a7e8f-ac3f-47ad-b8c1-6be7e49024e3"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "db53d47f-f228-43c5-88bc-175b680e21e0"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "937aa8e9-6a4b-44a8-9241-814636331e3d"
					"name" "string" "alt_jawOpen"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "50574827-075b-48b4-8514-830c3156bcf2"
					"name" "string" "tongueUp"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "29383fcd-55ec-4bd4-8715-4158dc40b5e9"
					"name" "string" "tongueUpThrust"
					"value" "float" "0.3000000119"
				},
				"DmElement"
				{
					"id" "elementid" "d5acccbc-c6c8-4b68-95c7-f5f4110868f7"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "Sit : voiceless alveolar fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "8834b508-7ac1-4b82-99ee-183020f168f9"
			"name" "string" "z"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "566de80b-49a0-4df0-905e-f787398e31cc"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "77a4d9d1-2479-4a1a-a337-694028763e02"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4f0e4798-7543-46b9-b2ca-443f6c1d970b"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dd3d69e9-bd59-431a-b028-e1d3e7ea8c4b"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2199999988"
				},
				"DmElement"
				{
					"id" "elementid" "a04e3a06-69c1-43b2-b1de-416f21068857"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "263c93ca-cfac-4388-ace9-efffc9b103d1"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ff47f18e-da21-4763-b85c-badb17328366"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b55ce362-9b92-4473-906a-7851b7bcddc2"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2f781eff-c21d-41b0-ab43-18c269478b21"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "723eb27f-9459-4305-b9c1-9f8c8e7ba6cb"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c7d63f5b-fb68-40d9-8228-e11614cee20a"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "878b742f-2d4c-4bea-9d86-311ac450722a"
					"name" "string" "jawOpen"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "13fadfc6-556e-454f-8bc6-bd068cb19968"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "127e42f5-64e4-4f1d-a7ee-41d4e6cea764"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fbee27d7-b25c-4bb6-bfb3-155db3105f44"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9671130d-1365-4d51-9b4c-c4f147dcbc18"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2199999988"
				},
				"DmElement"
				{
					"id" "elementid" "90fa3eca-320a-427f-8cf3-4cebefaada83"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e0deed46-f4a3-44bc-afdf-5217a7e5a417"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b993b019-5116-4cd5-9adc-c5cdb43a9a29"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d6c9a8ca-56f3-49f3-9fd4-051f2be5d950"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "af42d9d9-fb62-49bb-aea4-fd4a2a5c9436"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7c94f4b0-69b7-4ed3-9ef4-275a3ae57be7"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "84f2b714-2edb-4696-ba05-7116d5f7b780"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8e295eae-1fda-4500-ad48-3c96a6ded602"
					"name" "string" "alt_jawOpen"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "22917750-345f-4ff7-884b-2e6300f0b657"
					"name" "string" "tongueUp"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "0ff39668-998d-4b9e-a8be-44f3e4041af7"
					"name" "string" "tongueUpThrust"
					"value" "float" "0.3000000119"
				},
				"DmElement"
				{
					"id" "elementid" "c897b3f4-74f9-4dee-aef6-fdb246b9e628"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "Zap : voiced alveolar fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "d0840393-c622-458c-b7fe-5f8246c57075"
			"name" "string" "d"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "4601e542-6f43-4fe8-973e-ac0675deac5b"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4536754d-8f21-45f8-ba60-a3c9554a4d3e"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.1199999973"
				},
				"DmElement"
				{
					"id" "elementid" "0ebe92f5-d090-4ea7-9bc2-bb404faf2742"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.2599999905"
				},
				"DmElement"
				{
					"id" "elementid" "6e66eb79-6f35-4c57-a557-2c6910200701"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dae9a232-43f1-45d4-adef-e1b87842a2b7"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3eff8e30-646f-495f-8db8-742839bbfc02"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "41d2c935-c9d0-4aa7-a4db-7663e9ee9b24"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bf6c0f06-b9f2-4a7f-98e0-f3e880446364"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3d8c2caa-9976-4a5d-884f-52579ae1159b"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cc1463aa-daa2-476a-9345-4bde49cde59b"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "48357f1a-74ef-46c8-be59-9e1ce6be23cd"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a8907991-9a47-4756-a7b4-521520062c3e"
					"name" "string" "jawOpen"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "66a1b774-f631-4c07-a52b-ea4f0a0937dd"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ef622696-5e92-476c-ab83-3ce76c6114d5"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.1199999973"
				},
				"DmElement"
				{
					"id" "elementid" "39c8d6e3-fa51-44cc-8cbb-c2aa970bf76f"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.2599999905"
				},
				"DmElement"
				{
					"id" "elementid" "fb2fe19a-80c2-4c56-9fa8-75554e6c3b2a"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "071bdf2f-bef4-414d-b6ab-67585c742546"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b809254e-596b-443a-86c5-a37b5b3b90cc"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8b83aebf-5c97-435a-a918-5245f3f8c6d2"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "903c50b3-567e-44e5-a3ae-5a0f3f65c49e"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8f58290d-b78b-4be1-9e32-2544e37e9f7d"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3d17478c-7ea0-4bb8-a3f2-320f176c2535"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d9a76e39-1202-4009-ae69-debff1d1e48d"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b45c9d5b-f5d4-4ab2-9b48-3a691ecea4e2"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "4da36d83-5857-4aa6-8586-9b4f2788bec0"
					"name" "string" "tongueUpThrust"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "75f8e5f9-38ee-41bb-96a4-ebdef3ec2d07"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "Dig : voiced bilabial stop"
		},
		"DmePreset"
		{
			"id" "elementid" "88022115-f84d-4667-9725-060b55b4448d"
			"name" "string" "d2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "5a307565-4903-4cee-93b2-66e8da669cf2"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b0c52983-6dbc-492c-bc89-1c353644b74d"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1420286a-f8fa-48b1-8891-d250e4740c34"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "431b4d36-4df5-4d91-a254-fb7400e464cd"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e759d175-f828-4b89-9116-40b018d930fd"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4e3e285a-7dad-40f4-a357-53f3a435dbfb"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bf304cbb-ff15-4ec0-b958-9641ec7ee485"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9c109948-8743-49a8-9d32-0fec27ded4c2"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "15b93643-a14a-489c-9127-c20410219d6e"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2f01c44e-5912-421b-85a1-df5ab13e2014"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ee15f698-d3ff-4522-8426-36f163e89aa3"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "4e3b92e3-386d-4b99-aa7a-1e2d15070d2c"
					"name" "string" "jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "5abc33a6-fd40-41c8-8880-5942e3b13c13"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "89b83fcc-acd2-4fe1-beff-7cc31b8a89bf"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "624f0d50-9e7e-4ac2-b938-005dec0f4a2a"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4363fc52-9454-4497-8ba5-e52a6d1cc123"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c737e8b4-c6b7-4d3d-b8e6-a7b9b19fa446"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6cecadfd-d69c-4a0f-a962-ae4bd8dffc5a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "39462034-79d6-4bf8-bda1-97faf74ef5a3"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "80954010-d98e-452c-bfeb-45b83decd710"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "88a11b3b-5e0d-44c9-ad04-ac503bc0b477"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d7a1552c-fe68-43c8-847b-ea8152962a3d"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bd54e523-ccc7-422f-b69f-d69bd819976b"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "70c4c9ae-0b0f-4a7b-831c-392e5f7b36d3"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1599999964"
				}
			]
			"description" "string" "Dig : voiced alveolar flap or tap"
		},
		"DmePreset"
		{
			"id" "elementid" "c480ef82-d052-4a8a-a432-c5813fa0feec"
			"name" "string" "l"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "199f51ee-6f49-4b0c-8250-25e15b11dc34"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a1716c33-e7b2-418c-8e77-9a248f63f628"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.1199999973"
				},
				"DmElement"
				{
					"id" "elementid" "91ffcf9a-6b59-4a53-8192-05693abf1c86"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.2599999905"
				},
				"DmElement"
				{
					"id" "elementid" "0deb2a9c-1045-460d-8c0c-021f3b05ad4a"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2d619916-eb75-4c91-b2f7-215b9744b6af"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3c6290dc-ad47-4e8f-938b-ffcf8bef55b4"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5dc437be-970d-47c9-995b-048354c61eaf"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0742af00-2002-4e38-90ae-842e7e2c7ae2"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d62ecc7f-580d-4520-b3b9-5f2386c061bd"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "84a3fc1e-09a1-45ac-a123-49059371bbb7"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2df5630e-9ea8-4994-9c8f-646f1b90bebd"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a64a0591-87ac-42d2-b1f3-02551be94821"
					"name" "string" "jawOpen"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "cf63c364-bec4-4b7b-9a0a-831a1bd8a67e"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d1fde63e-7415-439a-a948-14d605748e2a"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.1199999973"
				},
				"DmElement"
				{
					"id" "elementid" "3ef07aa7-9025-4ee7-a393-5b295dce500b"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.2599999905"
				},
				"DmElement"
				{
					"id" "elementid" "d207a4db-289a-4d71-8d76-9260ce492039"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b704ffee-3642-419e-92e0-6bc3d97995e0"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6a22e9b9-c501-4ed1-a04b-d4efadf87ba5"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "33ae62d2-e59d-4160-b40b-4a3d37466259"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7150c057-3f46-40ae-a4c6-9d71d337e9f8"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "22f030c7-95b6-4c14-9aa5-437b9c6156bb"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "084b0313-f28f-4ec0-9c6b-16e38248576f"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4751b901-5ad9-4d33-bc37-097abf183d2b"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "68a83882-1c3d-42cb-85f5-19201204d6e5"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "76641128-6cf2-47fc-a147-0b620bc9db1e"
					"name" "string" "tongueUpThrust"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "e4fc5344-c922-4711-aa65-17a69d762f2b"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "Lid : voiced alveolar lateral approximant"
		},
		"DmePreset"
		{
			"id" "elementid" "d773d7b5-13a3-4386-9d2b-84298a076f09"
			"name" "string" "l2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "b3142dfc-fc48-4955-bfc8-6acd0fbe7248"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "158be726-deca-4ff7-8bf7-e144dcdd1583"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.1199999973"
				},
				"DmElement"
				{
					"id" "elementid" "54b0d9eb-ac4c-4a70-97d0-978def7a57be"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.2599999905"
				},
				"DmElement"
				{
					"id" "elementid" "ee795d70-9dfd-4da2-8c8d-04cd3daeea43"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e457458b-7a28-4739-b1b3-54eeae9ab13b"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "54c35f52-ace8-4512-9365-da4c0657c7f3"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "37537f46-5831-45bd-85a4-1632411affff"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c27efa13-59b4-440e-8315-6f6010b172ed"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "98d3af71-6620-417b-9ce5-47dca016bad1"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c30eeb91-a61f-493f-a291-c7a99c06f576"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2d860098-6d29-4964-a3fe-6ce15c82ba68"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "86d7fb89-7b9a-4987-8491-2eaa60eb23d1"
					"name" "string" "jawOpen"
					"value" "float" "0.0900000036"
				},
				"DmElement"
				{
					"id" "elementid" "d24dd811-9472-4114-8e35-efdabd6e1b64"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4100d61b-8d52-42cd-9e64-a5e64bd0fea4"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.1199999973"
				},
				"DmElement"
				{
					"id" "elementid" "7a68314e-0cff-4660-9ada-2960c82814b8"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.2599999905"
				},
				"DmElement"
				{
					"id" "elementid" "02ed3866-ec4f-4dfd-aa72-95525b3fb383"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b890d552-e89b-4a62-9b7a-dae88c6d8dc0"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "159476c3-cade-40e8-8a1f-8b645cba4329"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dc8ab734-6dc3-4d86-8c8f-c2f2d372f14a"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3bae0d88-af00-429e-841b-3e7ed9aa686a"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d8319d56-7e61-4832-b45e-7a7c62b21384"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6bd56d20-e463-4951-84de-7d3a51465eff"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c5318d3e-12be-4042-ad5c-3dfa8390bba8"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d2bb21f5-55e2-4961-b136-6df093b12403"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0900000036"
				}
			]
			"description" "string" "Lid : velarized voiced alveolar lateral approximant"
		},
		"DmePreset"
		{
			"id" "elementid" "3a95ee70-2705-4391-b480-c65675899106"
			"name" "string" "n"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "a6a6668d-c804-4a3c-8ea9-836edb65e19f"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e9caf584-8075-46a6-b84b-92469a4f5ce7"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a20b3fd5-6dcf-43fe-8b7a-cf45e2af89a1"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "df7463fd-e85d-4035-b96d-c6b0f6ae1066"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d6b66058-66f0-4874-a6d6-5d4fd20d0950"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e0f3dc5c-b4c0-4422-89a2-36178c7a9860"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "89bdedc0-898d-4cd7-8873-987904f1abbd"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6df7b46b-efb1-411f-b8ec-646675cca285"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e6adfc8a-2bbf-4bba-8e14-bcac7cd9867b"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "aa0c5372-842e-4573-aad0-ca7582a6a30b"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2a236487-f9dc-46ff-9d1b-12c9c2650f70"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5b1e1d55-b2ce-4026-91f2-37b86aa93974"
					"name" "string" "jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "62986025-cf23-49a3-a086-e7e6ce5586a6"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f76249a7-d31e-430c-b4f1-e1cb1285e200"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b1d067cc-a8c9-496d-be31-e215b12ca91e"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a7db35f4-59b9-4957-a10e-26d13c2220c2"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1646bfff-bf17-48c1-a3fc-df5a02a248d4"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "039864b6-eb05-479d-a096-f7aaeb3cc484"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2419da5f-54ab-4868-af0a-e1a8c454f52f"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "91af6c68-ff43-47ab-ac4c-fb9c9258043a"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3e04d61e-da74-4dfb-8a6a-38ddfa8b71e3"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b96a8183-e3e8-46b4-8c96-5f846950e75d"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a987cfa5-4c92-4719-abd4-b94f48175070"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b6b28f66-f910-416d-bcfc-dca3f3e45b6b"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "3a91181c-ecdb-439f-b228-facb506ee77c"
					"name" "string" "tongueUpThrust"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "26219723-eb37-43ef-b063-45d4cd78b03a"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "No : voiced alveolar nasal"
		},
		"DmePreset"
		{
			"id" "elementid" "1703f5bf-b2f6-4210-a619-34f7d83390a4"
			"name" "string" "t"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "af401de9-2eea-41e9-a55d-d721566eda14"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "812a480f-846c-45b1-a77b-f681d47c496f"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "46bb2c2f-f57d-4ce0-909a-3bbba6d355fc"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "564c256f-084b-448b-89a7-8a36508de8e3"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "109109e2-3ee2-470a-95fb-4793ed025682"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7abbf1b1-3b1a-4f24-b0e5-50463b53be0e"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1463c6af-cad6-4a33-9972-16f1699dc7c1"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f13cf6d7-5ac8-413a-b300-b758b0b3fbc3"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4caf711c-d1c3-489f-aec5-89eeebebc208"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "21b44a68-d72a-4098-be60-559ed04af0f6"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0a9959c8-52db-4734-91d7-6e5b43a1def7"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "19020ff9-7b68-4272-afde-060b0177511c"
					"name" "string" "jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "f862b366-1439-4818-9b25-5286bd412314"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3c36d915-a71b-4d0a-bf66-e9b0aa0dd486"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "831d7b22-4066-452c-92d1-15224f5d5e25"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c8df08a3-853d-44db-b0ab-7601ca0c1fba"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8ddd19d8-18f2-443b-9b3c-abc0907ccf97"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e5ef83fc-e815-44d2-b4a1-cf8da9dd180c"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "042a13ff-a8f0-4295-b1bb-dbb3ef4de1b5"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e5afac4c-7a04-4cdc-ad3b-cdd35e1f322d"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0f1af80d-290a-466f-abbf-0f14b70d17b1"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "151af9e3-db4b-473d-97be-71ec92c836f3"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "370a9377-3d85-44d6-a928-68035225e1ee"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "62b29cb4-5f1d-43b7-8358-43b8c31eab46"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "8e3586e8-340d-4a5f-91d6-8419540d6cac"
					"name" "string" "tongueUpThrust"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "d844504b-f009-4718-9bf6-6538f9f29ef2"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "Talk : voiceless bilabial stop"
		},
		"DmePreset"
		{
			"id" "elementid" "249b7fc3-07fd-4f2e-bf8d-d961afa6012f"
			"name" "string" "ow"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "41f0aae4-495d-43d6-86a4-e56f62312366"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8f9e0d3a-cb18-4071-ab5e-bf77c0484ae4"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6c3feaab-653c-4dad-bede-e1be9d692a72"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "15776a02-d943-4ab0-a0f0-0ab4b38bb0fa"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e60873e0-f524-4025-af03-65b54cd034bd"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e3a69ee6-3604-4304-ac6f-88e9cdb019d8"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "975420af-f79e-4e7e-abe2-ad4c1e0ef5c7"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "94115dd6-3c10-4e00-8101-dcf71e349dd5"
					"name" "string" "lipPuckerer"
					"value" "float" "0.7900000215"
				},
				"DmElement"
				{
					"id" "elementid" "fc573c17-00aa-43e0-8de1-e80f2f34cd1c"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.3100000024"
				},
				"DmElement"
				{
					"id" "elementid" "e9543783-9ced-4dc7-b8ea-013909876c90"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.2800000012"
				},
				"DmElement"
				{
					"id" "elementid" "35a5c3f4-404f-46d8-9076-732d9a1213a4"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "666f8adf-b2d2-40d7-b9da-3b3c51fccadb"
					"name" "string" "jawOpen"
					"value" "float" "0.3100000024"
				},
				"DmElement"
				{
					"id" "elementid" "9d9a89a3-5404-43c3-b415-3b0d6e08d662"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3dd8105e-6bab-4526-b72e-194388e7ac92"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "403e9d0b-b885-4d17-bd43-338e7008fde1"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c067a363-dc8e-42ac-8d25-c224485efb19"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c0f05784-04d3-4c65-a302-a6a194424384"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9eb6f5bf-d316-4e53-92cd-29e59ebd6078"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2a457af8-2b59-4325-8259-12d74960637e"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "97f5e536-83f8-49f5-885c-47789369a611"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.7900000215"
				},
				"DmElement"
				{
					"id" "elementid" "f48260ee-da19-4288-9076-7468b44dc0d7"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.3100000024"
				},
				"DmElement"
				{
					"id" "elementid" "1dcee821-fa8c-4f4c-9aa1-438b8bdaabb9"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.2800000012"
				},
				"DmElement"
				{
					"id" "elementid" "babf740f-23e6-4fdc-b389-5ea41fa04343"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "43eed1dc-5726-498c-9448-75ea72c7c823"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.3100000024"
				}
			]
			"description" "string" "gO : upper-mid back rounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "2f4cfcd1-9270-4c9b-ae40-488382b10976"
			"name" "string" "uw"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "e20d0435-2404-4b81-92e4-2f47f6d6c751"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5f288f46-6a2c-43b7-a76d-eb52755280c7"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "-0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "75a6887c-aeed-4318-acad-c2c00264bd46"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "-0.3199999928"
				},
				"DmElement"
				{
					"id" "elementid" "a2b120c5-7439-481f-b7d2-2b9a1131a91d"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7372c108-f602-4628-b470-995145cedd63"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9ceb9bf5-dcdc-4777-b341-96d716bee586"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3a4cb23e-a3fb-48a9-96e1-55372c49a0d5"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "434833dd-d86f-4566-8748-94b3ff7d7b50"
					"name" "string" "lipPuckerer"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "c13d7163-b869-4bea-a198-080449939fdf"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "845354bc-fd0a-49d6-ba5d-732e1028a951"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.3899999857"
				},
				"DmElement"
				{
					"id" "elementid" "99ea0e46-bc26-489c-9c83-978309f9dea6"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a1d13668-1261-4d76-aa2e-44abf9a2d913"
					"name" "string" "jawOpen"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "33d97a14-dbe7-4fc5-b0c0-a10e075bcc44"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ffa1d26d-28df-40b8-b9d0-84149e016555"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "-0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "e35399bb-fcb5-49ac-a868-99b4f17e9d89"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "-0.3199999928"
				},
				"DmElement"
				{
					"id" "elementid" "9bb468b5-4453-4458-95e2-0ae9e26bd3e2"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "91d681d7-7f8f-4998-8a60-f76d90d77d7a"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "46e950f9-1953-4803-90de-fa592a55fc6b"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c6a1c4f1-5bc9-4669-bd9c-bc08859a76a6"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ce1a003e-9131-4fa5-b51f-a954868646c9"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "b9a9037c-5277-400c-89da-e0948a4fb6fd"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "31bfd9c8-03cf-4714-8d6e-023343f5a617"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.3899999857"
				},
				"DmElement"
				{
					"id" "elementid" "84c6b4bb-acca-431d-8507-2690ea1456e1"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9daed642-2bf6-4f1f-a4db-33bbe310da88"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2300000042"
				}
			]
			"description" "string" "tOO : high back rounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "6aa68bab-7e59-484d-a6f4-6d7bddc22f6e"
			"name" "string" "ey"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "adb52caa-f78b-4a6a-a886-faa712739054"
					"name" "string" "dimpler"
					"value" "float" "0.3300000131"
				},
				"DmElement"
				{
					"id" "elementid" "45be6577-93dc-4cf0-b593-562f283f109d"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "9a065da1-8d59-4f4c-a48e-94748cb6386e"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c5b88049-6e01-4239-b011-5df03978112f"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "af7bc5f4-8f08-4f25-b682-91ae66b799db"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "df807e52-0113-4165-8563-9f7f6d90b3ee"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "78f577cf-cc80-48fc-ba7c-b68ef489ba26"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "19702c36-55ad-411b-843d-1e7ceaeca782"
					"name" "string" "lipPuckerer"
					"value" "float" "0.1299999952"
				},
				"DmElement"
				{
					"id" "elementid" "09637569-139c-4a47-98fa-31dce7a3701c"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3368242f-6c37-4200-8e67-f39f6e05a6c8"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6f5246c2-ef28-4e4f-81b4-012162c9b56a"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "859588ca-d294-4211-96f5-aedfff2c31f4"
					"name" "string" "jawOpen"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "914e5fad-90a7-4222-a618-d88f268d612c"
					"name" "string" "alt_dimpler"
					"value" "float" "0.3300000131"
				},
				"DmElement"
				{
					"id" "elementid" "34645bde-5201-4c78-9792-cc382b6d0383"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "b4a7378c-8f55-401f-a8a4-1ad3f8f35e62"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bfdc0507-5c49-4890-aa96-59263d277b4a"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "04382d67-c68d-4959-a968-26308abdcaf9"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "91fd50c7-5a2b-407f-a902-0699653a8fad"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "897d0182-19ed-4c3a-8e34-37e3c557a2b2"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "65ca62a8-9705-405c-9269-2eadd4e74840"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.1299999952"
				},
				"DmElement"
				{
					"id" "elementid" "0b6da624-d41d-455e-9da2-d8c1db1691be"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cdd70662-57f8-45e5-867a-80a83ddd65d7"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "11bb3689-d9f6-40f1-9648-dd77daf88094"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f943b74a-dcf1-4466-b3a5-7f6dc0e2507b"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2399999946"
				}
			]
			"description" "string" "Ate : upper-mid front unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "282ad927-3a7b-4418-8a09-94fe532f903b"
			"name" "string" "ae"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "5854cd1a-daa8-4972-a68f-6132af0a1b3c"
					"name" "string" "dimpler"
					"value" "float" "0.6299999952"
				},
				"DmElement"
				{
					"id" "elementid" "b88f9395-6813-4aaa-8968-ead2320a1234"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0685fb5b-f625-4afa-8432-3f0058c90851"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "463181e8-0c8f-47a6-a3de-d680a257b4ef"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ad4bed8e-d75e-450f-8a4c-15bf478eabdf"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7fb9ee2a-0a5e-4096-8db8-638a57e072b0"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "331cd76a-e98a-43c5-9d25-d5d6015a88f0"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2bd3c2d0-a476-49ef-9b7a-fa40389541b1"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e4745fab-191f-4e29-abe2-e8bd313ea234"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b71ae419-5b7e-495e-a72a-1153b2c651a6"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "22a6f6bc-f3dc-4003-9e42-a14a850f2dd0"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "64e1c861-4eaa-4263-9ee7-283c5ef45f84"
					"name" "string" "jawOpen"
					"value" "float" "0.4499999881"
				},
				"DmElement"
				{
					"id" "elementid" "fd0b8b2c-0684-4aa1-9e1e-b670ae8ca475"
					"name" "string" "alt_dimpler"
					"value" "float" "0.6299999952"
				},
				"DmElement"
				{
					"id" "elementid" "705aa2bc-7489-48b0-a646-08c1f4db4692"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "11bf3c0b-dd8d-4e86-8070-252435c856ab"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3879ef6d-0f97-4898-be5b-8e85773a291d"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1f90a17d-f304-4b35-aef6-2555e1f65032"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "95296b1d-3c90-49c4-9dd7-855607136994"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "af1bfcb0-28a9-4b0b-9429-814c4bb003bc"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a3c12fad-56cf-4ddd-9b6a-6fdfa0fb5fb0"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "54146a1a-5708-4fa7-a552-a13b20af449d"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f92fe04a-5e63-433a-9f5d-88d282213064"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f58f8b71-0750-43bd-8ed5-203d662caee4"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5c55e288-586c-462e-9c31-0ff962ebfccd"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.4499999881"
				}
			]
			"description" "string" "cAt : semi-low front unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "0f416192-e239-4cbf-8e71-7ffdab9dd5d9"
			"name" "string" "aa"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "265525dc-b4ac-4ec5-a56f-1b9aaa343c30"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "21e4590c-2d69-4ade-ab8b-d9a197e38e19"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9f8e3bb0-1447-4d2d-931f-d1e96d25c157"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1ccd9c8b-fbc4-4a82-b795-1c2a9ace0085"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b9838e83-0a5e-49c9-af03-99c237fb5055"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7391cf56-75f9-4185-94ff-cc38d7576f07"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9ccd4a40-a325-4576-96ee-27a16cb8af2a"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a24861c6-b33d-4217-a321-6a9195b36b0d"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9127ce15-ccb1-41bd-9578-e852050265d1"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "21b1ee2e-e427-4e15-8025-e83b2e93d63b"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3fe8962b-9a87-421c-ae09-0c9b87e54aa6"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "42000c52-35d7-4ccf-a5d5-a168e3312532"
					"name" "string" "jawOpen"
					"value" "float" "0.4600000083"
				},
				"DmElement"
				{
					"id" "elementid" "5cca67ac-8dd9-402b-967a-bd573903650b"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "79fa7fd7-4c11-4aa7-8269-727af12fab7c"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a26c8aa7-7726-460b-b883-5a1520ba3c2f"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a3f7ecf8-59b3-4324-87b8-b6682f445e30"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b2e07e9e-2845-4df7-9e6d-d917f23157bb"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "414349de-aa9e-4c5d-be42-a6cee5dab22d"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f5e56620-f5ca-4aa7-96ef-06741ae19f18"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "77923d0f-69b4-4e9f-a735-130faebc291f"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fac30106-8fda-4f79-a5d9-79d55b6d1aeb"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "05332e63-88c4-427c-8c42-9733e1e33bce"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5bbb9f98-5d31-4317-b29b-4287ac7046f8"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a0787214-2daf-4e19-9e44-13df9a3c97e9"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.4600000083"
				}
			]
			"description" "string" "fAther : low back unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "5c2d3fba-0306-4582-bb44-4a4d96db4b7f"
			"name" "string" "aa2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "3d049df8-df7f-474e-98ad-022f0f47b88e"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "89d86baf-f231-41e5-9b56-540730169ada"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a4f8b4fa-9c13-4854-a98b-dc1eb2997722"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "967760b2-503c-4235-8db2-32373e7ddc15"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8d3a3320-c6e6-4b37-9f58-1d1197fac03c"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "23141e32-1ebd-4eba-a094-67fbd98962f8"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ba6eb36a-cebe-4bb2-bce9-700b3c7ca364"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "48906b52-f2be-4123-851c-474a911424b5"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cccb48a3-8b3f-4c44-890f-84db16cfee05"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b958cdf7-4682-4751-8af8-b16978131cd1"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2c77798f-eb31-4875-a536-fcf6a9dcaacc"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5881e5f7-3b33-4235-89c6-ae4749485a25"
					"name" "string" "jawOpen"
					"value" "float" "0.4499999881"
				},
				"DmElement"
				{
					"id" "elementid" "8f3064cd-2a8f-4410-ad5c-9c5c60001e4a"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "08eb01eb-16ff-4eb1-ae7c-90955d0b629b"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ae031627-87ac-4115-98d8-4e26b9b8ef86"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "96c71a07-ef6f-4738-b1fc-83e3d69e20b8"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6e0b4046-d05b-4472-87e3-20f96e0c165e"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bcf3263e-9acc-4978-a93a-b905a58e146a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ba7a1d12-8922-4f9e-b17f-48939c3b004e"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1fc85e5f-d509-4e30-b8f8-b79c817d4ba4"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "53a92b2d-d1c6-4bfa-8fb5-be46f34e930b"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "67069c7a-0f57-4a15-9a68-33b5c053583f"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "041a32a0-54e0-4b9e-bca1-f113ee7e6630"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8bc0413a-1d05-4588-afca-e856093ba235"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.4499999881"
				}
			]
			"description" "string" "fAther : low front unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "0d089f55-a259-4900-b8f3-c89411eb5028"
			"name" "string" "iy"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "5515e7a5-47c0-4089-9e89-59c0e474c367"
					"name" "string" "dimpler"
					"value" "float" "0.4499999881"
				},
				"DmElement"
				{
					"id" "elementid" "795f6b13-0233-46a6-9138-33c43425535b"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a96bbf68-f3fa-49fa-9902-9c41d9b0340a"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e65c85a1-50cb-4d93-b39b-78f2dd377383"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "b36ef0ca-491c-4a6b-af44-349bf98b43c6"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.2800000012"
				},
				"DmElement"
				{
					"id" "elementid" "4ea7d5c2-3562-422b-8b2b-b617e74951b8"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f8d7c837-ab92-44d2-a3e3-617540244835"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bcb89b00-0b7f-4196-820d-b9580da6272f"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ad019f4f-ecf0-482b-8925-3a70cbc4e69a"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "39535053-0764-42f3-9e60-12c245c7ef66"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3a6f4014-757e-4e70-aceb-6e9a3220356f"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cb51c2fc-28aa-43cb-9106-22d34c2be3e8"
					"name" "string" "jawOpen"
					"value" "float" "0.0799999982"
				},
				"DmElement"
				{
					"id" "elementid" "b470ec1e-37f4-4149-926d-ccf21c9d7aef"
					"name" "string" "alt_dimpler"
					"value" "float" "0.4499999881"
				},
				"DmElement"
				{
					"id" "elementid" "2b59b6ff-de81-4d37-a64d-17a0de94f6ab"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ff98db8d-9bad-4016-a231-f81ef30ee500"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e927544e-126a-4ec9-9592-a2bc542f6c37"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "708372f9-2849-43c6-8b88-8d17a8051f14"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.2800000012"
				},
				"DmElement"
				{
					"id" "elementid" "7f23f8cb-eb67-40d2-9ad4-dea66b80a2df"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6f34880e-aab8-4095-b819-0f146a28bb19"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9f8b0d3b-9c0b-4e38-a1e8-bca451f0d53b"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "27e88a85-763a-4a04-b13d-fb88c637fdc8"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "879f9e68-08c4-4b03-a7d1-deb6c4cbca3e"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "09ff2cc1-28cd-4468-8be5-4b4b5fc2cd60"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "af87faf2-e5ee-4f66-8c2c-9d7502d53d23"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0799999982"
				}
			]
			"description" "string" "fEEl : high front unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "7c30caa7-083e-458c-9abc-392c070cf1c4"
			"name" "string" "y"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "16a2c9c5-1493-4cd9-92f0-9a572a0c3743"
					"name" "string" "dimpler"
					"value" "float" "0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "872eb53b-61fd-4df0-b619-489df2d98e2b"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8d287f1e-d5a1-4ea5-aa04-a770104f9f2e"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "47b84463-bc12-4569-b1a5-c781f3d63c59"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "10dec3c3-570b-4ab3-b49e-ccc8002f8a49"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dfc7400c-bc2d-41dc-adfd-2c4d66415f52"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "06f2a917-a6a2-42cc-b483-1e477c73146a"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a988feff-ddba-406a-9cc7-3d8544a9a79e"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dfa544d7-9893-449c-bc7e-3520e9b22f07"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7eac155a-dc7e-487b-8236-9ef424349e26"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "96f2062e-8aa5-4735-a7c1-9757ddb8eccf"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "92921d3d-2f66-4a09-8c47-e9067d2ceb67"
					"name" "string" "jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "4d42bfb5-6206-4eb7-b6a0-46bd9110284e"
					"name" "string" "alt_dimpler"
					"value" "float" "0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "a566a650-4e77-4852-9920-49f39db92111"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7997c990-93a2-4728-80d4-f1556dbf23b8"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4ab1f423-01e5-4f80-9c8b-f80ec48a64bf"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "02139a4d-9c7d-4e9b-a45b-d63e877604cf"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1e834645-c95b-4acb-a0fc-28a116a379d9"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "58213e2c-8fde-4360-b19f-8caa187a70de"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "472a4d74-29a4-4da0-98c5-54e23ee2b1a5"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bd7a6453-6c61-4236-a5dd-3f5c75f86fb9"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f14734a3-379d-4fb9-a7fd-1447791295ba"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7cba17a-51ed-4dbb-b95b-c83762b98217"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "3f65a295-473a-4e9c-bcba-b6409526f75a"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "2e5c3eb8-d190-4adc-8b23-a9c2da219a8d"
					"name" "string" "tongueUp"
					"value" "float" "0"
					"weight" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "739246de-f6eb-4d86-9d85-2614d3475b93"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "0.5"
				}
			]
			"description" "string" "Yacht : voiced palatal approximant"
		},
		"DmePreset"
		{
			"id" "elementid" "a8b1b5c0-4ab3-479d-9da3-f9339c638ff7"
			"name" "string" "ah"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "89254da3-37af-478e-a366-73f1c47c6e1a"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7551a79e-5e51-4a50-87b5-ab5f05af4a68"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c7093b23-5876-4faa-9e93-283a922859eb"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "adc36786-6f35-45c7-a859-ca687ef23d3f"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0dac6b98-3595-4d0b-af28-2db7f09d1b6c"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "aba1508b-96b4-4364-a663-ff15c6fed83d"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "36837ebd-8b08-4de4-8d37-59a4a687d4ef"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "18085302-34b9-42da-9de0-4f249e9b58e4"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "71626468-f732-49e7-9b72-43acda53433d"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6d5c5047-8839-43b0-9b52-c8e4c9a83ccd"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4e5617cd-8a8f-4ba8-a6b7-328a525b76be"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "77dd5373-da27-418b-b8d1-9b4869109aa5"
					"name" "string" "jawOpen"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "7b6f9988-5d5a-4708-8253-c67e39fdb619"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3a689202-01b6-41c4-8018-12b6caf7ec47"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a6111150-2c3c-413f-abd8-929b2343ac20"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "773768b3-e61b-4f29-b4c9-8661df2e3d9e"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "50aa69b4-64b7-4138-88ec-fa7c15c9c142"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f1d16170-3fab-4a4c-9933-f1e4d91bd54a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "72c01252-46b3-4b7b-9501-37c3ee829b74"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8f97f2e5-947d-4d81-9e1f-cf333eea3800"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7f87e34-21f9-46b1-a4ba-707786f17a85"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e2345691-1865-4c82-ad24-a37309182539"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "caca8754-a20a-47be-b094-2081e26aa825"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "034ec051-11f7-449a-b754-dcff835005e8"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.400000006"
				}
			]
			"description" "string" "cUt : lower-mid back unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "eab11f93-9145-4532-8821-67c0ab674670"
			"name" "string" "ao"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "f365f068-c5b1-4fba-a2e3-cc887c379e34"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "939ab0ff-6548-4bb8-9b09-ba06f0d31ee5"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e5e22c9b-5abd-4929-8f94-caf6229ef9e0"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2c6c1d1d-8e45-413c-b7b0-fcc313e3d465"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "92f234ff-075c-4b91-b1a7-26a320958aac"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "add1d847-b1bc-426e-92b9-eb69c040aa5e"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "35aaaf2c-858c-4537-8cde-02cb26052825"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6e214696-616c-4b9f-b9b2-4654f3f383fc"
					"name" "string" "lipPuckerer"
					"value" "float" "0.349999994"
				},
				"DmElement"
				{
					"id" "elementid" "19a34ab4-edaa-4ae6-a3f5-c740c6adbcad"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9ccb95f1-9c7a-46c4-a218-1b85a287faac"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7a8f2b34-1410-465c-b4a5-96c58c826da6"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8cb8ad19-9f15-4927-823f-ce5db01a8b17"
					"name" "string" "jawOpen"
					"value" "float" "0.4600000083"
				},
				"DmElement"
				{
					"id" "elementid" "8fb6040d-7e2a-47db-9efd-d7e7c698cf58"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ad6dfc18-f1ec-4450-8d8e-517c26ea15ff"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "48191c8b-c008-42f5-b804-f82960343870"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c202dceb-5844-4cbc-87b5-edc0201a4bbe"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "88a977ef-e725-4cdf-ac3a-33053067ac3d"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3ff1cf86-60bb-4856-a8a8-9a06025fef80"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "984cf41d-c67e-4e04-b5b6-4bb3f21bda1c"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ec1b6498-3ab5-4205-a622-5b81fd612839"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.349999994"
				},
				"DmElement"
				{
					"id" "elementid" "9e57a2ac-8a55-425c-89c1-d045872c81de"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4d4a3085-293d-4e50-a847-0b669caab969"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d5d49824-868d-4d39-9166-587b6e0157aa"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "834fd0ba-871d-4b73-880d-3829e9a0ff1a"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.4600000083"
				}
			]
			"description" "string" "dOg : lower-mid back rounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "d58047ec-d38a-47dc-89df-c0df001542f2"
			"name" "string" "ax"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "e40f5eda-2c56-4c20-b087-abbccded4359"
					"name" "string" "dimpler"
					"value" "float" "0.0500000007"
				},
				"DmElement"
				{
					"id" "elementid" "44f9b468-9266-43d2-8a3f-79e53c60122c"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "320ded74-fdf4-45de-a336-155e127d2fe0"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "-0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "be9393f4-b60c-4fd8-9d93-399c10b02ce9"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "97a67c71-e1b3-4328-99ca-414f6a92c7a5"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "2def6300-cb7b-41c0-b5f0-82c7f1fac995"
					"name" "string" "lipCornerPuller"
					"value" "float" "0.1299999952"
				},
				"DmElement"
				{
					"id" "elementid" "74527823-cf17-41a9-a5ce-6e0959bf809f"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6dd9195e-3ae4-4542-af74-60e338791507"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3e6a6ee1-4ddf-4d7b-a8cb-21c32ff83bca"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3bccbf75-1839-4620-89fe-8d03c4219f8a"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5c91ed50-08eb-4e7a-a312-ecc3533a14c7"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "94957671-2ef1-4969-a0dc-75924103a818"
					"name" "string" "jawOpen"
					"value" "float" "0.3600000143"
				},
				"DmElement"
				{
					"id" "elementid" "11bc1710-f204-4f45-b42d-dda85e0ff820"
					"name" "string" "alt_dimpler"
					"value" "float" "0.0500000007"
				},
				"DmElement"
				{
					"id" "elementid" "cc09c050-9285-43dd-b888-313a6a791af4"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "f089af72-4322-4442-9068-6586671f8cce"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "-0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "e2b5122e-0f4b-464a-accd-b484e751f287"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "c0509626-9246-4746-a5bc-edc686133ed9"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "80683cb5-c7b3-4b67-a969-974f272e0c3f"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0.1299999952"
				},
				"DmElement"
				{
					"id" "elementid" "455496c1-aa2f-42e4-b5f5-f7d5b2f43150"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f6ec7863-ce3f-4b10-b150-824b7a9466fc"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bfc9ce15-3d7e-4530-9a2f-f32388606708"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9936aa59-fd9c-49f4-b476-6689e1da09ae"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "268e74ac-7144-426b-84fa-13002660fa25"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9379c548-dd38-4950-b2f9-875fd6a5062d"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.3600000143"
				}
			]
			"description" "string" "Ago : mid-central unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "36224379-b57e-45c3-ae72-3631fee773ed"
			"name" "string" "ax2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "52eb8171-83ce-4836-805e-acb0e56cdcea"
					"name" "string" "dimpler"
					"value" "float" "0.5500000119"
				},
				"DmElement"
				{
					"id" "elementid" "c2c5cf80-75fd-47f0-a502-c5f76e9ea40d"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "c596fd07-de74-4ddb-837b-e2262167b1a4"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "-0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "2bc878ba-f49f-46c5-a65a-743bcc1ccf6c"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.5"
				},
				"DmElement"
				{
					"id" "elementid" "08aa2bb6-76ba-4966-b64b-f4045625b402"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.5799999833"
				},
				"DmElement"
				{
					"id" "elementid" "de3b7106-5c65-4a00-b3c8-86f7da29638d"
					"name" "string" "lipCornerPuller"
					"value" "float" "0.3300000131"
				},
				"DmElement"
				{
					"id" "elementid" "4285f0d2-8ba2-489c-a022-2fcccec1af3e"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ed4b3ae2-51bf-4650-91ab-3e51d8ac4ed7"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4148e2ff-7840-4a2e-936e-c4e4e135b6fe"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f9b03cdc-f15a-465c-a310-4477a1783800"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fb046225-ef45-493f-8d06-3b5554c1331e"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "69929cef-dcbc-4f8f-88d7-07b93a611d31"
					"name" "string" "jawOpen"
					"value" "float" "0.4300000072"
				},
				"DmElement"
				{
					"id" "elementid" "373dee3e-e1c5-4fa6-96f6-8bad640759c5"
					"name" "string" "alt_dimpler"
					"value" "float" "0.5500000119"
				},
				"DmElement"
				{
					"id" "elementid" "43234a96-53ad-4f74-af5c-343691d541f0"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "64bea040-728c-41e3-8be8-47aa9380a1a6"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "-0.0199999996"
				},
				"DmElement"
				{
					"id" "elementid" "7a625867-5525-42f5-888c-26bd77fbfa99"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.5"
				},
				"DmElement"
				{
					"id" "elementid" "581c6d32-1afe-4cba-b7f2-ed4a9c54a4a3"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.5799999833"
				},
				"DmElement"
				{
					"id" "elementid" "212f19f7-e763-4e0c-a045-356323f7cd51"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0.3300000131"
				},
				"DmElement"
				{
					"id" "elementid" "7cf37eb2-20a0-47c3-ae67-6e95ec3529ad"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "44e935d8-4cbc-4a45-9afc-7ffaa37fa8ac"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ede7a3aa-8bca-43ab-99be-59249e75d32f"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ddce6cd6-3c60-4447-bce5-37d39ecbfdca"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cc4d4a48-4bbc-4b95-a930-b7002b52172e"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "56e86d79-f9fd-403d-8ca8-c20baf1c6bf5"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.4300000072"
				}
			]
			"description" "string" "Ago : lower-mid central unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "a219b91b-3b7e-4b57-adbf-2ac7da84be62"
			"name" "string" "eh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "4c82252e-cac3-41d8-b18b-4ee4ab60639f"
					"name" "string" "dimpler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "50ae822d-e108-45bf-add5-b22e12f488af"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fe264779-7ac8-49ac-8c01-b75912fc71c6"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6ae04870-c0ba-4b7b-a8f7-ace6aff3fbdd"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "ce17369f-a737-47d7-a30e-fbdfffc5e46a"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9792f0cd-df33-4476-8bef-8a1e667500c8"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c1068593-5cee-4619-8be1-7bfb4e60396f"
					"name" "string" "lipStretcher"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "a74b0e35-77db-4475-a6ec-36cda3adeddf"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ffc8383a-f073-4b02-8c6d-9f37f1d35ad8"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4fd49fb3-3752-4da1-ab3b-c13fee2f6e6b"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "31d0ccc7-1c39-4f10-a6f0-a830b2bc23d0"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "12d56f5d-dcbd-4585-92fe-1774dd6b7830"
					"name" "string" "jawOpen"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "9bf58c67-3bf0-4ddc-b860-fa3d8294c5e8"
					"name" "string" "alt_dimpler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "dd7ed349-70cc-4276-a7eb-90adf5bd5165"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "898d0483-e9f3-4e0e-b419-b01171706b53"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f6c7b15b-f2ff-4615-8dd3-35bce9f0f3cb"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "4bd6f8ab-0a28-421e-a911-7c25b7e75def"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2aa30329-00c2-4f73-b7ac-9d46a7ff3b33"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7dbb5457-c190-49f0-9304-71bb3732c501"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "971e52c4-bbdf-47e9-810c-fe3e550121db"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3819f384-0638-41d1-beaa-b54a121c1903"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "320b0a70-a05a-49e3-afb3-d76f97df76e8"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "19903e1c-0181-4993-8cfb-38c84bf75733"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ceda891d-d57b-48ef-9ed1-9a7484aa89ec"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "pEt : lower-mid front unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "31193cc6-244d-4c4a-9b64-dc3ec30e80a4"
			"name" "string" "ih"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "33a5cee1-ff09-40ac-984d-f0426dac0b31"
					"name" "string" "dimpler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "c4966641-48b0-469e-87d9-0ecb66599b38"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "576cb0b5-5a2d-4733-abbc-ff81c8424547"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "807a837c-a1cc-465d-b656-385a779a9e52"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "4e60afee-bb21-4ca5-9c27-69382a543285"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e6a20ccd-3c3c-4b3f-8e7c-eb0d99c48344"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a1df3e34-51d3-4e7a-85b3-ee204c8a16fa"
					"name" "string" "lipStretcher"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "9f54e201-dcee-4476-bf29-1d7d2a80a20a"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c76fe24e-c316-43da-b3e5-367a98912ac7"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5b4a3f2e-9b64-4147-afd1-e27cc959c7ac"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6e4e466d-6fdb-4ac4-912e-ef539141e6a1"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9319da65-a914-44ca-a54f-32be48c7a5e1"
					"name" "string" "jawOpen"
					"value" "float" "0.25"
				},
				"DmElement"
				{
					"id" "elementid" "75efcf83-707b-4a5d-a39d-9913ca3cd9f0"
					"name" "string" "alt_dimpler"
					"value" "float" "0.6899999976"
				},
				"DmElement"
				{
					"id" "elementid" "b7323b2a-5da0-4deb-ab27-b5f0a1e8c45a"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "159b1c1e-0644-4186-a3b7-63f794a0a861"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1b6e987f-286e-4314-95a7-8aac1922501e"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "319be59d-218a-45aa-9b4a-dff108c8d7dd"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a70554d9-2dc2-436d-b674-5cbe44c00bc8"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "02ce4b86-5d0e-4568-a1ae-2b465704b92e"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "67ee8f19-d30c-4d70-a11a-6af738fab73a"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "32bacfcc-0e69-4b72-9ae3-d9dc0543025c"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e2b3a594-e40d-495c-a127-d8bde319937f"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1c0874f6-e7ef-4e3d-8a80-5132c5498920"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3eb5a720-221f-4cfb-b9cb-90710f070355"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.25"
				}
			]
			"description" "string" "fIll : semi-high front unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "6e532bf7-7a7a-4747-af14-ad5bbde67c22"
			"name" "string" "ih2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "1c32b82a-6b0f-4517-a28d-d6e2c6e9a635"
					"name" "string" "dimpler"
					"value" "float" "0.25"
				},
				"DmElement"
				{
					"id" "elementid" "742707fd-4e9a-4697-a844-700d9d05d25f"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b0ee442c-407a-467a-a2b8-616f3226f901"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2ce4ba78-24bd-40d2-9235-6cd941a006a4"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "efc40b26-5dd2-46ec-893e-9fa5b414fb0b"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "97aa542b-b617-4ffb-bf30-c49399209153"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "928438aa-668d-4ca8-b539-4ea89d8ae468"
					"name" "string" "lipStretcher"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "21b6be78-6368-4407-ba7e-8a21fb262681"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0f44a206-2838-43c9-b773-49d4f2058f4b"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "391f38c7-65e5-49f6-b49b-699dca5c49de"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "19122a0e-4539-470d-80c4-9cd206ead9eb"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4dabf8b9-e50e-4bc1-966a-a007569f8713"
					"name" "string" "jawOpen"
					"value" "float" "0.150000006"
				},
				"DmElement"
				{
					"id" "elementid" "3861b993-d82e-4c0e-8e01-5f7a76f7300b"
					"name" "string" "alt_dimpler"
					"value" "float" "0.25"
				},
				"DmElement"
				{
					"id" "elementid" "21e6ea15-c123-4999-8eb1-df454877f306"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9a2714e8-a3fa-4aad-9f78-37ff1ab387a4"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ffd7755c-a633-4eb9-9959-e1dab3634628"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "ef7241e4-08d0-4c1a-85ad-2d3773f36cca"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0.3799999952"
				},
				"DmElement"
				{
					"id" "elementid" "b34f9cdc-3000-4f12-bc86-ca9459117ddc"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "91f6ffed-b195-4e56-8657-7d57ae6d0955"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "d2b85d28-cac9-43fd-a7b8-d398f5b1b527"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5f6d26da-d217-44c8-8447-28774adccd51"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "562662df-8f5a-4a7c-8a82-07ebfe70c44c"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "82bfc210-299b-49de-a332-53d8db179899"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "390531ef-1fc5-4aa3-a28b-e4bbab5eae95"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.150000006"
				}
			]
			"description" "string" "fIll : high central unrounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "0329ec7d-a914-4eb7-97d3-5888425c772c"
			"name" "string" "uh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "d0214c2a-8c7f-4934-9445-b6c14e71be0b"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "46013d4a-1e86-495d-bc42-ee2ff36fbf03"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7f9ed8e1-3b0a-4c75-9260-74e1cda4d519"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c043a803-740b-4e5f-b607-0cc804621b28"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "59d0a43d-1d7c-4b3a-8052-3a2c2b6f7fbb"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2a1d9e7e-b842-40c5-9e8b-5ed258db9857"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0473238b-c173-4698-be45-43aa4ac07994"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8c32ce88-4ff0-4db3-89b7-76401eb10928"
					"name" "string" "lipPuckerer"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "ae8cebf6-6431-4206-8ac5-6cb3e464405e"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "288529fd-60df-413b-88d0-3efd021af4a4"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "e2006de5-f910-42c5-8129-128a1ada4406"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7b7d930-4e54-4146-9b43-33bf882bc9dc"
					"name" "string" "jawOpen"
					"value" "float" "0.2099999934"
				},
				"DmElement"
				{
					"id" "elementid" "3d0b8586-adfa-419f-aab2-cc32ef0beb4d"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "94a9e099-8dad-4175-bd4f-5e941c1c6e93"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "673b9456-45a7-4f33-9bae-e3c585619c03"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "16b6fa15-4d0a-4964-b4ce-941cd3e19fc0"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7f164aa0-1bd3-4eaf-b070-f4946ec8ff64"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "afe662b2-45a7-4700-853b-23e45edc6f77"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3265970e-d457-4201-805b-cf71d65f6186"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7de2e17c-8314-4762-a4a0-2f0b44afe572"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "e55788d7-a066-4f52-abd2-ac791f1826ae"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "1a12bf52-de34-4113-b6e8-5ded5db5b079"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.2300000042"
				},
				"DmElement"
				{
					"id" "elementid" "528eef30-a982-4bf4-8326-20c81df28ee5"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2f3d3d75-3e3a-48d5-ac99-c9b1dcd77a2c"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2099999934"
				}
			]
			"description" "string" "bOOk : semi-high back rounded vowel"
		},
		"DmePreset"
		{
			"id" "elementid" "1ed2389e-25ae-44be-9749-549cb12bf856"
			"name" "string" "g"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "828f837d-6943-4177-9e2e-ef275fb7f531"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4f119b8b-50ea-4e36-9a73-9592d742bcd0"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d670989c-ad2b-4d8d-af84-5330cad797ec"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "36a68f59-2875-41e6-830e-46c9a03898df"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "cef46845-5d2c-456a-91c2-1598a5643d9e"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "01b98877-93bb-469f-bcb1-aa3a6af5a252"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "bf537058-e754-404a-993a-ed2ee101f554"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7bcd8d41-73ca-4589-a677-19be4c8dcf96"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0e3c9d4d-7d9c-4a1d-95b9-292db3be0255"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "25f68354-3426-49a0-8f22-db4feed5c49a"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "81f31991-555f-4cd5-926f-10da18c0b99d"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "5ece3550-08a5-490f-a49d-ba6056390e6c"
					"name" "string" "jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "b45b55f0-778b-4d98-8fbb-0889b3361e55"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "54b31468-4c14-4edd-b175-d33b96833d5c"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4c68f4f1-f41e-4ff2-8308-d7d18dccf418"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "762bc5b7-9fc8-4de5-b8c5-22fe9a13484b"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b0607ad9-bf91-456a-a69d-857d71c8148a"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "3e2ea63c-6ef0-4b99-b854-0c102bc811a9"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "07e81044-7058-4bec-acd4-75d8693ef6a3"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "441f9d4b-d4d9-4e7b-af6b-2c47e55cbb13"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fd0cd39e-920d-4689-8a08-36de0abcdb23"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e3c779b5-203f-4588-bb6e-0f4b81cc2c99"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "547800d0-5367-4e50-8011-095657ea0f29"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "fd667446-c748-4cf4-82f5-03eb64218ec7"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1000000015"
				},
				"DmElement"
				{
					"id" "elementid" "25095a0b-892c-47bd-b6b1-82b1cfbd6b55"
					"name" "string" "tongueUp"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "taG : voiced velar stop"
		},
		"DmePreset"
		{
			"id" "elementid" "3a4e840a-0879-4b03-9c43-96fea488d7b0"
			"name" "string" "g2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "033e0e40-df4f-460b-8aea-a217705a43fd"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "802647b6-ff28-46ee-ae02-12981bd75dc1"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b69a25b0-3ae6-4109-b49e-d1b5fb05a433"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "815dc1e7-5d0c-43c3-b460-cb5abc8f849e"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a99e655a-1016-49f7-b233-cf69fed7bb2c"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5734fc40-f112-4666-a5fc-3c2826cac561"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "7e797a80-6db4-4cf5-8c8a-0e8d6fc37c25"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4ca988c7-4063-4537-8061-51f91287229b"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c13f2aa8-d812-4e5b-a772-2291da75a547"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ae372c8d-0880-46bc-a16c-64923217cfb8"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6c0e35f1-578b-4dfc-9f50-471544ced28d"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "f33a2f46-5e86-4d2c-8a80-d0b5d8db378a"
					"name" "string" "jawOpen"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "da39bc75-9383-472d-8b1a-0d09854145b9"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b0fefa79-b2dc-4818-96d2-c4cea84d7f28"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7508e13-b4fa-4d51-a155-2b4e7f228489"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5cf9a760-ee5c-401e-a891-a4b892008238"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7bed139-0e43-4294-af3d-4a08e0b592b6"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ae37fa22-f1ae-4ccb-9ec3-b9de9050003a"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2f7b5ee7-85d1-4512-b75a-ab9069a24098"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0ccd4b0f-324d-4d24-8bb6-00ecfb8f67dc"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "948e48a8-c909-46c4-aef2-4da1bc14866c"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "622decdb-3ac3-4137-a32d-60556bde08de"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "dc49a468-6f04-4285-89d4-6aaa475a3ca2"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.200000003"
				},
				"DmElement"
				{
					"id" "elementid" "6f87d548-c390-4445-b1f6-db47e952c7e9"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2399999946"
				}
			]
			"description" "string" "taG : voiced velar stop"
		},
		"DmePreset"
		{
			"id" "elementid" "57a1e26f-9233-4f0d-974a-8ff9800be391"
			"name" "string" "hh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "5d5952cc-863e-4bb6-ac24-b052456c02d0"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5f9b24ea-5a26-4147-9ccd-93d75673c996"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "596d16bd-c250-4803-957b-f571ad9980c8"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "e5637250-af0f-4c92-b2d0-8e754eeca130"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a96ffb45-6b34-43f8-a5cb-cf53a5526962"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ac7e96c2-03d4-45db-9f7d-87f63e0a238c"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "37f2023c-45fc-429e-9816-11263d1ec7fe"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6ae1b6aa-2473-474f-aeaf-70ed596c6992"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "83ba4665-651d-41a7-b5e0-47af45ddad5a"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "109da714-4ed2-4f3d-a488-ad94b7671d4b"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "17017cfa-0fde-4c25-9d18-6ff3929b890c"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0f25d731-c1ee-4c36-8c83-b31bb5f9f1ba"
					"name" "string" "jawOpen"
					"value" "float" "0.1800000072"
				},
				"DmElement"
				{
					"id" "elementid" "446e9e80-8f8d-4aed-9c74-29ccc8d09a31"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "575f6bc3-f992-47ad-9bf1-18eeac589af7"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0d46783c-8158-4c4d-a9f2-62beb9da1c89"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "38d400c3-d7a3-43c2-8245-d0593eb99baf"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "41737a1d-099a-43a6-8344-bd58b93a5ebc"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "51c2d52f-4563-49b0-a03a-6361353f2fb8"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "92028378-be92-497a-8580-6ba631c280e1"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "10ea4b03-868c-43f9-b498-4a7741f63795"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2fddc87a-119f-4294-b3f4-998f222e5fbd"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0512c95a-851f-4c7d-bbb4-bf862d1a3b5b"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b7b0fe8a-011b-4f2b-b697-4ad9d7546214"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1a102ffc-be1f-4218-8524-fc21ed2b337f"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1800000072"
				}
			]
			"description" "string" "Help : voiceless glottal fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "5247d63f-2ff0-453e-b861-e3f06c551fec"
			"name" "string" "hh2"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "a52c27d1-5d16-4db7-8e68-bd6c1d14baa5"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a505d0c1-732a-41b3-957b-7e342be59264"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0a56cc5c-e682-4393-a7e1-1e73d7d16fef"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1e8037cd-1aed-429e-ad59-9ad98f0f2a9f"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a984a875-83e5-4ddb-bc82-37171cd2ffde"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1473d0c2-52ad-48d4-bb54-d70caef8a8ec"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8318e9fc-64e5-417d-acd8-9ceb6e47330a"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4f67e1f3-e787-4f2d-8ec4-d97be0c0bb1f"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "c4778bda-a6c8-4210-89a9-625b24434a39"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d9242488-0f39-40b4-990f-796f77b93e0d"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5b40023b-d3af-4bdf-a679-048198c08d5e"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "f4627638-7856-4936-8218-1b07e3d698e4"
					"name" "string" "jawOpen"
					"value" "float" "0.2899999917"
				},
				"DmElement"
				{
					"id" "elementid" "57846e4e-0a6d-419d-b0d9-9f8c292c1db5"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "45a791ea-26ed-4f24-842b-5f3d7d2e96a7"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "11e08f20-c688-4bf4-b2b0-ac301808649b"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0aeab2b6-e783-4cd8-90d7-7205bdb93f84"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d7ff6b21-a543-48d0-9414-6ec30597824d"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5a68b6d3-1b24-4a47-bc4a-cb11afc6c68c"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9d33a08b-42e7-48a6-8caf-5611630c4a95"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "35ef9fa1-3dd5-4d97-91e5-8e3e5871fe2b"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a3eb9784-aa1c-460c-9ef4-f0b439d48b66"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e0f6cb6d-859f-4395-a16a-32b76b756b4d"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d702ca72-7fcf-4ef8-a12a-7be9283bebb0"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9cd503d9-350a-4348-844c-4451e862572f"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2899999917"
				}
			]
			"description" "string" "Help : breathy-voiced glottal fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "f7252720-bd59-46b2-b8bd-114815f8095c"
			"name" "string" "c"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "379b6283-e5aa-4764-a377-f664687d1730"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b334bc8f-f8f6-4bf5-ac5d-91507353ebb8"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0a364ad6-e047-46c0-b168-0170c53d7cc4"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "1d68b3c6-708c-48da-befe-7028b58173d6"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5e092124-d719-450c-863b-48eb4c5ba499"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "6a9b43ab-dfcc-4e27-a414-86bcf8ce7b25"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "2197336b-b77b-445f-9efe-eb6ea723d8ed"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8e89dbdc-46cc-4544-a4a1-5160ee5af63e"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d7c7d1e0-9f84-4b30-bcc3-59118cb1d992"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b1bc4db4-db73-468f-a967-560ee12722b1"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5e985738-f14a-455b-8159-89d082479d1a"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.2199999988"
				},
				"DmElement"
				{
					"id" "elementid" "adce3990-0e19-40cb-8010-0bf17c5fd0f1"
					"name" "string" "jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "c127dffa-9e96-430a-9f9e-0af0c85474d8"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "ac7a0780-fbe0-4dbe-b22f-65a7264279bc"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e3bd4b1a-2883-4ff9-b89f-461469812928"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d0b53bc6-ea9e-41dc-acc2-a08a52f867c2"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a7b44bb2-a48b-4125-bf0a-ff409706a6f4"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "752997c1-d3cc-4842-a654-0f20d0db2da0"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "83521672-3cd3-4eb7-9c48-a7ec3e0a6a8a"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9d2ce95f-be12-4295-a9e1-f649173924c5"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5d7b2e24-4d3e-494d-830c-f5e7a6c4a104"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "db1f8d27-b089-490d-ab03-c9da488ee13d"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5eac0f03-226b-47c8-af06-895b4354611c"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.2199999988"
				},
				"DmElement"
				{
					"id" "elementid" "5bb5e668-5913-4d46-a48c-c9c7559e9f98"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.1599999964"
				},
				"DmElement"
				{
					"id" "elementid" "da1c36ba-42de-4592-89ab-56650fa6d266"
					"name" "string" "tongueUp"
					"value" "float" "1"
					"weight" "float" "0.200000003"
				}
			]
			"description" "string" "Cut : voiceless velar stop"
		},
		"DmePreset"
		{
			"id" "elementid" "d070a288-3d1c-4d22-ae36-bbf1cbbb85c7"
			"name" "string" "nx"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "bfce5418-0884-4f9f-ae4a-67750bf197ad"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fbc6da36-3c3e-4f70-9856-89dfdc45e537"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9d383613-4011-4cf5-b7a2-24a50ac76dfc"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "883a625a-76ee-4838-9468-cecc794fa77c"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "fb7f19f7-f0ea-4cd0-beec-56bff7fd6e52"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "886e8650-f291-4ec8-b940-72b3f93b9429"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "35f8e187-9546-4ae2-8d7a-e590951f83f2"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "9df2686a-41d3-4c3e-a908-12cb88b155b9"
					"name" "string" "lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d45d0df5-acad-430a-a1b1-e6c733be7055"
					"name" "string" "upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "da9b0fac-6b51-4ca5-87ce-a56be1d94020"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "5f2a2cef-ee7d-49c1-8965-7f330c0366fb"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "5648b8cb-e6b5-432c-a203-58a751dfc6f1"
					"name" "string" "jawOpen"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "1971dd83-a48d-433f-873e-9e3c2916f87e"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "627d0462-58ba-4a89-bedb-ef0305ba0f5d"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "a074c5ae-3438-41a8-ad3c-4379f961472f"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0650271d-e6e6-40e6-a6de-155f76679800"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "d02cd084-e0a1-4946-aa36-f0d3aedb151e"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "8f886b67-ec35-416d-99d9-dbb125b77bc6"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "80c84241-706a-4f81-b8d3-16d25837065b"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "60c3ec7b-ed11-46cf-907a-0404f9391e74"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "b50e8af7-cd1b-4076-978c-6eeff4c6ace4"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0d70806d-736b-42df-9a9f-be68c1d5f2d5"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "94fe4205-89b4-4a47-a8a7-ecc03e6f98eb"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "-0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "9ec62e9f-19e1-4b22-9216-ed402b1b4d4c"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.2399999946"
				},
				"DmElement"
				{
					"id" "elementid" "25036301-5a71-49c0-a9f5-f90b06dd6519"
					"name" "string" "tongueUp"
					"value" "float" "0.200000003"
				}
			]
			"description" "string" "siNG : voiced velar nasal"
		},
		"DmePreset"
		{
			"id" "elementid" "693f0cdc-f1ee-427a-8e66-96da9137e7b9"
			"name" "string" "zh"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "1facb821-4835-41c8-8d68-78d3f8e01b36"
					"name" "string" "dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4b74dc5d-6fe0-4ed2-9e81-21f2ac871d99"
					"name" "string" "upperLipsTowardAndPart"
					"value" "float" "0.3000000119"
				},
				"DmElement"
				{
					"id" "elementid" "4bf34bce-68a7-41f0-9ce9-b70096b3675a"
					"name" "string" "lowerLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "790a1a20-0b87-4893-97e2-9461c6c7e280"
					"name" "string" "lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "50a5f13e-2cfa-4fad-bae9-3a574389178d"
					"name" "string" "lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "35e1399a-a96b-4c19-b843-27fe0397f3e7"
					"name" "string" "lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "40664690-f6a3-416c-b7e6-a1389013e017"
					"name" "string" "lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "946a2a0b-2aec-4db0-83f3-c7244ca2051e"
					"name" "string" "lipPuckerer"
					"value" "float" "0.4099999964"
				},
				"DmElement"
				{
					"id" "elementid" "6ef36135-52a6-4ef2-ae88-243cbe6025bd"
					"name" "string" "upperLipFunneler"
					"value" "float" "0.4300000072"
				},
				"DmElement"
				{
					"id" "elementid" "dd9aa93a-5bea-41d3-bf98-d29352b0643d"
					"name" "string" "lowerLipFunneler"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "cb39a963-d5dd-445f-b026-8684a224885a"
					"name" "string" "jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "db956360-5aa2-48b2-9340-b4015ccaaaaa"
					"name" "string" "jawOpen"
					"value" "float" "0.0900000036"
				},
				"DmElement"
				{
					"id" "elementid" "ed56d14d-2d90-4a62-b621-7cd58bc00100"
					"name" "string" "alt_dimpler"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "39d96836-b5b3-4c67-bf05-c118041d4697"
					"name" "string" "alt_upperLipsTowardAndPart"
					"value" "float" "0.3000000119"
				},
				"DmElement"
				{
					"id" "elementid" "6cad7c48-605f-48aa-aa92-1b86395af636"
					"name" "string" "alt_lowerLipsTowardAndPart"
					"value" "float" "0.0399999991"
				},
				"DmElement"
				{
					"id" "elementid" "1b5de3a1-30c5-402a-9fdc-9c00e6b402ee"
					"name" "string" "alt_lowerLipDepressorAndChinRaiser"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "90ce589b-9f2d-4121-9d82-7185fb1e179c"
					"name" "string" "alt_lipCornerDepressorAndSharpLipPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "4d914fab-3b18-4ba4-9763-5ebd184390b9"
					"name" "string" "alt_lipCornerPuller"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "0a148060-6dc7-4e47-9a88-b32f4b2c2520"
					"name" "string" "alt_lipStretcher"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "801c9e6b-375b-4d18-befa-9103b552c5af"
					"name" "string" "alt_lipPuckerer"
					"value" "float" "0.4099999964"
				},
				"DmElement"
				{
					"id" "elementid" "429f05f5-56fa-4c29-b31a-b7e454edbbc4"
					"name" "string" "alt_upperLipFunneler"
					"value" "float" "0.4300000072"
				},
				"DmElement"
				{
					"id" "elementid" "294ce243-4dbd-4a9b-bf7b-27ce87147017"
					"name" "string" "alt_lowerLipFunneler"
					"value" "float" "0.400000006"
				},
				"DmElement"
				{
					"id" "elementid" "58e0f38a-99f7-49f8-a663-59f2ee13eabb"
					"name" "string" "alt_jawSuckAndThrust"
					"value" "float" "0"
				},
				"DmElement"
				{
					"id" "elementid" "e9d7c057-b0b0-4c8a-b583-3ac2d618b653"
					"name" "string" "alt_jawOpen"
					"value" "float" "0.0900000036"
				},
				"DmElement"
				{
					"id" "elementid" "b87689a8-518b-46f5-bb7b-f79f049a33de"
					"name" "string" "tongueUp"
					"value" "float" "0.8000000119"
				},
				"DmElement"
				{
					"id" "elementid" "8943f6f2-74b9-4766-9f38-4c8022d5080c"
					"name" "string" "phonemeCloseJaw"
					"value" "float" "1"
				}
			]
			"description" "string" "aZure : voiced postalveolar fricative"
		},
		"DmePreset"
		{
			"id" "elementid" "2529bbed-3283-4f7f-b4ff-a0f47cffedbc"
			"name" "string" "b"
			"controlValues" "element_array"
			[
				"DmElement"
				{
					"id" "elementid" "a194944f-3df8-4345-a068-3434959464ef"
					"name" "string" "phonemeBMP"
					"value" "float" "1"
				},
				"DmElement"
				{
					"id" "elementid" "0950daa5-b825-4254-804b-c99323d123bb"
					"name" "string" "alt_phonemeBMP"
					"value" "float" "1"
				}
			]
			"description" "string" "Big : voiced alveolar stop"
		}
	]
	"visible" "bool" "1"
	"readonly" "bool" "0"
}


*/
