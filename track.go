package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Track struct {
	Name         string
	children     []Element
	Collapsed    bool
	Mute         bool
	Synched      bool
	ClipType     DmeClipType
	Volume       float32
	DisplayScale float32
}

func NewTrack(name string, clipType DmeClipType) *Track {
	return &Track{
		Name:         name,
		Synched:      true,
		Volume:       1,
		DisplayScale: 1,
		ClipType:     clipType,
	}
}

func (t *Track) AddChildren(child Element) {
	t.children = append(t.children, child)
}

func (t *Track) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(t.Name, "DmeTrack")

	e.CreateBoolAttribute("synched", t.Synched)
	e.CreateIntAttribute("clipType", t.ClipType)
	e.CreateFloatAttribute("volume", t.Volume)
	e.CreateFloatAttribute("displayScale", t.DisplayScale)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range t.children {
		children.PushElement(serializer.GetElement(child))
	}

	return e
}

/*
		"DmeTrack"
		{
			"id" "elementid" "164d4b8c-5ddb-4a21-bc01-834680cc4a09"
			"name" "string" "Dialog"
			"children" "element_array"
			[
			]
			"collapsed" "bool" "0"
			"mute" "bool" "0"
			"synched" "bool" "1"
			"clipType" "int" "1"
			"volume" "float" "1"
			"displayScale" "float" "1"
		},
	"DmeTrack"
	{
		"id" "elementid" "a2e0ceb5-e2ec-41b3-acd3-22e262cef4c8"
		"name" "string" "Film"
		"children" "element_array"
		[
			"DmeFilmClip"
			{
				"id" "elementid" "488da5a2-c1ce-40b6-aae5-133e79d385f6"
				"name" "string" "shot1"
				"timeFrame" "DmeTimeFrame"
				{
					"id" "elementid" "549ada85-d3b2-4444-8a45-55f97ca5fd29"
					"start" "time" "0.0000"
					"duration" "time" "60.0000"
					"offset" "time" "0.0000"
					"scale" "float" "1"
				}

				"color" "color" "0 0 0 0"
				"text" "string" ""
				"mute" "bool" "0"
				"trackGroups" "element_array"
				[
					"DmeTrackGroup"
					{
						"id" "elementid" "fe3091df-7417-45c3-9a53-e1a0c4a350bf"
						"name" "string" "channelTrackGroup"
						"tracks" "element_array"
						[
							"DmeTrack"
							{
								"id" "elementid" "02e4d621-e3ea-4004-8aec-4915f18a1c8e"
								"name" "string" "animSetEditorChannels"
								"children" "element_array"
								[
									"DmeChannelsClip"
									{
										"id" "elementid" "e9068407-be37-4484-8540-8f7b9e78255b"
										"name" "string" "tiny_04_1"
										"timeFrame" "DmeTimeFrame"
										{
											"id" "elementid" "ee52a31e-01e0-47ee-9e64-e3cc595c60f0"
											"start" "time" "-5.0000"
											"duration" "time" "70.0000"
											"offset" "time" "0.0000"
											"scale" "float" "1"
										}

										"color" "color" "0 0 0 0"
										"text" "string" ""
										"mute" "bool" "0"
										"trackGroups" "element_array"
										[
										]
										"displayScale" "float" "1"
										"channels" "element_array"
										[
											"element" "58bf2a7d-8fff-425f-9fde-2241804fa0ad",
											"element" "4044f367-f423-4722-a9ee-c0f1ea0557f4",
											"element" "71fe1e20-c238-4c52-90ba-58fed307d17f",
											"element" "194777d2-742d-432f-a807-82dcc2174a6f",
											"element" "f219585f-7003-4d04-bc47-d114d05e4f76",
											"element" "3860cc07-d762-41f0-8723-d536da7aa19c",
											"element" "24b4a924-1a02-48a6-9e8f-c0128b03af3d",
											"element" "ea652206-241b-4ea1-8e6d-85b85f94c370",
											"element" "30fe924c-5562-4455-9413-cf6ae292a2ab",
											"element" "b648390f-1137-4245-94d8-3a0ed47dee98",
											"element" "c5416b9b-09e2-4e3a-a0d1-4788fece4260",
											"element" "c80e59e4-2e44-4f8f-aa9f-d2635af89cb0",
											"element" "db64c138-71a3-46a3-b8f4-ea1476afbe1a",
											"element" "84df4fd0-4f7b-4442-abca-392835d9dd44",
											"element" "853483dc-8d4b-4cc6-ae58-bf85a5292828",
											"element" "fbe895c0-83ed-486b-893f-16215dd194eb",
											"element" "d084053e-6980-445d-a249-814a51586a6b",
											"element" "1ca8f4cd-3bec-441a-9f0e-4c84e62e913f",
											"element" "8a699622-c653-4850-988c-20c62612b935",
											"element" "1268f8c2-0dbb-4cc7-af69-ef7d79517e7f",
											"element" "44758c95-9eb6-4047-904c-3ec904b677e7",
											"element" "cbb4c26d-0f1c-4cbe-95bb-137a4d88d4c2",
											"element" "2ea38434-3c8a-405b-b98a-6f1f02291714",
											"element" "7df0c7a6-5146-4b7a-8ec3-1f1ea9a5fdf7",
											"element" "3f3918b2-58d6-459b-b888-311a9e814b07",
											"element" "6ebe26e0-7f2b-437d-b31a-d89f5722c9ee",
											"element" "2575dd73-8e59-4487-b2ca-a74fb925b983",
											"element" "ebc0717e-53a0-4b72-9639-afe155fd2457",
											"element" "1ce0e6d2-4e40-4240-b9f9-e4db532702d7",
											"element" "12977884-3ccb-438d-b311-0c0f57336216",
											"element" "56086c17-7eaa-47f7-87fc-ba3443b56e26",
											"element" "fce2d298-3d96-4f2a-a56b-f4554df2dbd4",
											"element" "6c31d75a-5c91-4164-b5a2-e73f68707a7d",
											"element" "0044b5d2-ad78-4a69-9e96-4f168d0035d7",
											"element" "7eb665e0-06c1-4020-b17f-71594882c75a",
											"element" "9132e089-cb08-480e-b803-dba47b13cde9",
											"element" "0babd7be-3050-43b9-a5ed-7e2ce8e03185",
											"element" "0ddb22d4-30c1-4f02-bb9d-c38a5b5de346",
											"element" "24605135-8560-466c-a330-8a2844fdf07b",
											"element" "4d334ed3-3dd7-4889-9863-76cc32e928ad",
											"element" "edb32f56-7f16-44ec-9296-c748f30047c7",
											"element" "298456b6-ea8e-4be6-8358-ac24d03a0f95",
											"element" "1ad67fae-aec2-4c1f-985a-2d715039ffa9",
											"element" "16a3373b-3084-4a97-af5f-a308e59f8d6e",
											"element" "36985d52-0154-4d29-9af7-2b50f57f7002",
											"element" "24307d54-1bf4-493e-8d48-b7d26b40228a",
											"element" "7a811f77-2adb-4541-8787-49f04a7b0d00",
											"element" "b3aa5496-4c80-4ca5-821c-a3784aecac3c",
											"element" "f34e1700-d447-4f99-b60a-f2c456795365",
											"element" "054ec613-0006-4875-8868-51d2f532a26f",
											"element" "4bc8a8d0-f87b-4b03-af31-8f9c31e2faa5",
											"element" "c9c3c929-53f1-406a-90b6-a0a49f22a7eb",
											"element" "82149a55-7ec3-4f72-8fb5-57a4ddf7597f",
											"element" "c5889eec-e95c-4739-8927-3c70853a9554",
											"element" "a3b5eb3d-0cf9-403b-b76e-15dab14f6eea",
											"element" "0da710d0-358d-467a-a0d7-d496188123bd",
											"element" "93521db1-4db3-4494-aa90-b8fc996b93b8",
											"element" "964fcd7c-4760-4981-8447-db78cb479045",
											"element" "f1fee255-2aa1-445e-a9d6-4e4ab95e6137",
											"element" "753ae967-7dc1-4429-99b9-c3e28079eee4",
											"element" "fc0b5da8-b3bb-4966-95c7-ffcfb75b0f01",
											"element" "32959b56-6d45-4683-9715-284d2a6bb690",
											"element" "c6767bae-1225-474f-a93b-8536744eea99",
											"element" "c7f5145d-9323-46a6-bc26-de818caa351d",
											"element" "4ad621a3-f1e4-4666-976c-0136bd7258c6",
											"element" "928381ae-a334-4de1-86a8-d99c29c657ef",
											"element" "b837f8ed-238d-4d46-9820-2ff1593a49bd",
											"element" "23f3a876-c417-442f-ac9f-c17935093802",
											"element" "2cf511d5-f4da-4c8b-ae0c-299d428b4a5d",
											"element" "b14d9919-57b2-460a-8dd5-e11a918fd65a",
											"element" "d5c83abe-952e-41bd-b3bb-56c45457ee8a",
											"element" "d6234e70-f387-4cde-bd0b-5b71f19a95c8",
											"element" "7fa91b52-5920-4caf-a81e-dcba3892c2e9",
											"element" "9c9db7ac-081f-4101-8791-b183a5edccee",
											"element" "06b5d7f1-b6ec-4c2e-a661-4afa32fe6487",
											"element" "a03aec2b-b7fc-4399-8b97-462cee7cc7f0",
											"element" "47165812-a03b-442b-970f-4a19b9aa9497",
											"element" "ba8b9ef2-ee3d-4120-9fe4-d5baaa413788",
											"element" "f924b24f-478c-4166-8980-dc9344765d4d",
											"element" "a78bbb5a-0073-4705-8667-36d1cd4d0166",
											"element" "3913bd91-0594-4246-af3e-db6a87b9963c",
											"element" "232b140b-19d1-4d8b-9cb7-f081ace9dd59",
											"element" "9cf675c2-4135-4055-afd0-a109498a64df",
											"element" "cbc77550-6a98-4aa6-b685-ff77e11a00d2",
											"element" "b3873bf2-2879-4228-a6e0-c739ef1c92b4",
											"element" "c5223f76-05ac-4ab3-9cd8-877406673201",
											"element" "cfdafa0d-02dc-4f9e-a3eb-92b142f98fd1",
											"element" "53b46255-28fe-4bbf-8df4-cd2b94d55709",
											"element" "891faf96-6eb7-4073-9b1d-dd3c511effa0",
											"element" "47b5adff-1285-41c0-9540-fb2c36e33563",
											"element" "71c3cf7a-befa-41e4-a7dc-572176dbad7f",
											"element" "afb059cd-0d05-4552-babb-747d1b7a8641",
											"element" "714c9bc9-6ebc-49b9-adc3-97f6ae72dcaa",
											"element" "933054c3-f518-4ad1-a55f-0756ecae7e09",
											"element" "dae53886-9d43-476c-9d15-e81502b225fe",
											"element" "59004230-6354-4a9d-8334-dcea988d19be",
											"element" "a39abd6d-42f8-47b2-82f0-a2ee9cb27ea4",
											"element" "3c425454-707f-4e45-92cc-c235993d2be7",
											"element" "448229b0-1fc7-4a57-b2a7-21e73f676f05",
											"element" "118313f4-b3ac-4efa-816c-22042c6bdab7",
											"element" "c997b97f-a11e-4ad3-82c3-8b645f5e5f35",
											"element" "fd54c148-2a6f-4c4a-8fbc-d992aeac95c8",
											"element" "d405ead3-4bd4-48c8-b88d-524fe40a794e",
											"element" "53440d09-0a2e-4cf4-9dd1-c78b8d812455",
											"element" "b567c34d-330e-4de0-9aa8-2c240e41422b",
											"element" "cefe54dd-f596-4e52-9fd7-9fc8b60dc2dd",
											"element" "af67dc34-b219-4dc5-b4d9-e889fae26851",
											"element" "2a9d52be-91e1-4659-ba23-efd476e75321",
											"element" "fd152252-ff8b-4bdf-b480-efb2fc7abec1",
											"element" "597153e0-c450-4d52-9519-b8fe53ee49d3",
											"element" "d5bdf724-9da7-48c3-bf9e-af8ef214ee21",
											"element" "712931d1-9c60-4bb5-b227-4bb9df0f1f1b",
											"element" "8031bd67-9c49-4019-90c5-0f763b37b578",
											"element" "9685ae11-939b-4af0-89c0-fa728faf689f",
											"element" "6fb5af58-e8de-4b26-99f7-28cf5cea929c",
											"element" "8c464eeb-13d9-4d69-be0e-12a3b530653f",
											"element" "ab524ab4-7b5c-4f2d-9360-6780e1371b3b",
											"element" "1cb3beae-3785-45fd-aef1-29e18a66a9be",
											"element" "b0174f3f-0d7f-4841-bc61-cdc393ec29d9",
											"element" "de2a37dd-879a-4b64-91a8-c19bbf976a35",
											"element" "31a2a549-c75a-4f81-b78b-20f1b8808f5d",
											"element" "e421447b-e1ea-4c5f-8bd0-93ed6d3a2235",
											"element" "c5a88cd5-8d47-494a-afa7-768d2db17756",
											"element" "c55d70ac-afcb-4103-92d0-278cf1c32bcb",
											"element" "f5b854a8-3506-44e2-83f1-5546c6850f26",
											"element" "3b3a80a9-b2ca-4399-a03d-e60a0412626c",
											"element" "957f94bb-7ea9-4c67-bde8-655d81018ec0",
											"element" "18217967-5b41-4f25-ac41-544ac152982a",
											"element" "51842ca8-1be5-4d85-8446-f9109749f1c1",
											"element" "acd57536-6691-4394-8ec5-0db0c8f172ef",
											"element" "eb73b4a3-c2d1-4cdb-9310-b1d63cf81466",
											"element" "dbb09555-8c5b-405b-8cb7-8f164841954d",
											"element" "25f69b5e-cb37-44f0-945f-1b042bf6287c",
											"DmeChannel"
											{
												"id" "elementid" "13e948fe-1f0d-4732-9535-c7c451e291dd"
												"name" "string" "scaled_rootTransform_scale_channel"
												"fromElement" "element" "195243ee-5973-42b9-9fd3-06b351e861b4"
												"fromAttribute" "string" "result"
												"fromIndex" "int" "0"
												"toElement" "element" "9168f52d-a1ab-47f5-8e90-4b09e5c88eba"
												"toAttribute" "string" "scale"
												"toIndex" "int" "0"
												"mode" "int" "1"
												"log" "DmeFloatLog"
												{
													"id" "elementid" "a4215837-2231-43d8-8dcb-71ff5cf2bd1c"
													"name" "string" "float log"
													"layers" "element_array"
													[
														"DmeFloatLogLayer"
														{
															"id" "elementid" "71fe343d-53fe-4157-81ad-4b6d9540f0e4"
															"name" "string" "float log"
															"times" "time_array"
															[
															]
															"curvetypes" "int_array"
															[
															]
															"values" "float_array"
															[
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "float" "0"
													"bookmarks" "time_array"
													[
													]
												}

											}
										]
									},
									"DmeChannelsClip"
									{
										"id" "elementid" "80e189d2-592c-48ec-b50f-6cacebfcca8a"
										"name" "string" "tiny_ambient1"
										"timeFrame" "DmeTimeFrame"
										{
											"id" "elementid" "6c244dd3-61e7-4ce0-9b3c-1f5cd9f822da"
											"start" "time" "-5.0000"
											"duration" "time" "70.0000"
											"offset" "time" "0.0000"
											"scale" "float" "1"
										}

										"color" "color" "0 0 0 0"
										"text" "string" ""
										"mute" "bool" "0"
										"trackGroups" "element_array"
										[
										]
										"displayScale" "float" "1"
										"channels" "element_array"
										[
											"DmeChannel"
											{
												"id" "elementid" "4596ec34-e26c-4964-9316-7690feec30be"
												"name" "string" "emitting channel"
												"fromElement" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
												"fromAttribute" "string" "emitting"
												"fromIndex" "int" "0"
												"toElement" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
												"toAttribute" "string" "emitting"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "8ac7c91f-91ec-4cfe-a32f-731cea345975"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "aff6e15a-0956-40a7-84cb-3725a248367b"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "8d773beb-82e2-4cc8-8a26-2a025acff114"
												"name" "string" "visible channel"
												"fromElement" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
												"fromAttribute" "string" "visible"
												"fromIndex" "int" "0"
												"toElement" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
												"toAttribute" "string" "visible"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "f9af0026-1d24-42eb-b402-fac9f18320c3"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "a08c07ba-4566-41d1-b975-928e71fce85d"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "407ec0a5-a953-47e8-b564-2edcca299def"
												"name" "string" "simulating channel"
												"fromElement" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
												"fromAttribute" "string" "simulating"
												"fromIndex" "int" "0"
												"toElement" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
												"toAttribute" "string" "simulating"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "002b28b9-7574-45fd-8c61-3885c7b50762"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "1ea22bcf-78c2-4a91-b2d2-f5ff978a6161"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"element" "f23645c4-b517-4a6f-a1f4-8663f3048395",
											"element" "eabbe2c2-2c49-481d-b9bf-0287f226baf8",
											"element" "ef3957b4-c9df-4f1a-be84-5f63aaa242b3",
											"element" "7d488edf-8a65-451b-8103-c9c75ed5ec9f",
											"element" "42625eba-6fa9-43c4-abfe-7546f26cb753",
											"element" "da787765-a663-4f07-bf73-84ecd3e01a77"
										]
									},
									"DmeChannelsClip"
									{
										"id" "elementid" "d64ac2e1-069a-40da-b146-6465efe9b887"
										"name" "string" "tiny_ambient_lvl2_1"
										"timeFrame" "DmeTimeFrame"
										{
											"id" "elementid" "2b7e0d4f-a1c2-4371-82f7-a726ed192c96"
											"start" "time" "-5.0000"
											"duration" "time" "70.0000"
											"offset" "time" "0.0000"
											"scale" "float" "1"
										}

										"color" "color" "0 0 0 0"
										"text" "string" ""
										"mute" "bool" "0"
										"trackGroups" "element_array"
										[
										]
										"displayScale" "float" "1"
										"channels" "element_array"
										[
											"DmeChannel"
											{
												"id" "elementid" "170e93eb-193b-49a3-a13a-e30c3c5aa412"
												"name" "string" "emitting channel"
												"fromElement" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
												"fromAttribute" "string" "emitting"
												"fromIndex" "int" "0"
												"toElement" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
												"toAttribute" "string" "emitting"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "2bfa1cf1-dbb7-4371-a676-d277925e299f"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "b7f58599-dba3-43f6-95ff-9e27264f393b"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "f600ea37-f1d0-455c-89f8-ca4276f8027c"
												"name" "string" "visible channel"
												"fromElement" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
												"fromAttribute" "string" "visible"
												"fromIndex" "int" "0"
												"toElement" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
												"toAttribute" "string" "visible"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "f9450236-70dd-4fcb-9cba-2f4725fef372"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "258b8184-a4a4-43ff-9b11-042264daf37d"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "b28d43c8-0c50-4216-ad1d-7c05d715cddd"
												"name" "string" "simulating channel"
												"fromElement" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
												"fromAttribute" "string" "simulating"
												"fromIndex" "int" "0"
												"toElement" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
												"toAttribute" "string" "simulating"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "c74587f8-6008-49b8-8183-849eb1b349dd"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "2b714eb7-e23f-4bd7-baac-e2caa80404c3"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"element" "c9105294-7219-45c9-9745-1219d20ced84",
											"element" "c790ff00-2a10-4ed0-9f2f-71852ab50fcb",
											"element" "464ca423-957b-4388-b3ee-2ca12db3af42",
											"element" "6e16ed07-b512-4374-af30-3888a60fff49",
											"element" "28e77bd4-e6eb-4103-977b-5f096cc88b2a",
											"element" "d347f4a6-4d0e-4714-8739-f461758c8c67"
										]
									},
									"DmeChannelsClip"
									{
										"id" "elementid" "6d635785-82de-4d65-9cbd-4007fcf7ee07"
										"name" "string" "tiny_ambient_lvl3_1"
										"timeFrame" "DmeTimeFrame"
										{
											"id" "elementid" "b1595cec-3ccd-4511-b01d-5a833c34bd9e"
											"start" "time" "-5.0000"
											"duration" "time" "70.0000"
											"offset" "time" "0.0000"
											"scale" "float" "1"
										}

										"color" "color" "0 0 0 0"
										"text" "string" ""
										"mute" "bool" "0"
										"trackGroups" "element_array"
										[
										]
										"displayScale" "float" "1"
										"channels" "element_array"
										[
											"DmeChannel"
											{
												"id" "elementid" "ceb2129b-d279-4a28-86ab-7c7b6558a84c"
												"name" "string" "emitting channel"
												"fromElement" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
												"fromAttribute" "string" "emitting"
												"fromIndex" "int" "0"
												"toElement" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
												"toAttribute" "string" "emitting"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "3a2cb73f-02d9-458e-ba77-d11c13043c16"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "365f7acd-a012-414c-a4f1-e6818df7c0c6"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "4c15031f-4c26-40a6-b337-fe635016e79e"
												"name" "string" "visible channel"
												"fromElement" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
												"fromAttribute" "string" "visible"
												"fromIndex" "int" "0"
												"toElement" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
												"toAttribute" "string" "visible"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "7b30532c-d137-48d8-b86d-69accb4db458"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "6e14c058-3d9b-42b9-8d91-8a478b0a43a1"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "900d8acc-aebf-4714-b8e6-ab797bcecd98"
												"name" "string" "simulating channel"
												"fromElement" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
												"fromAttribute" "string" "simulating"
												"fromIndex" "int" "0"
												"toElement" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
												"toAttribute" "string" "simulating"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "9061d008-35eb-443a-aabc-a09ad94a3880"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "0aa1aa75-9976-476d-9fa2-beb40cd63424"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"element" "e060083f-ad5b-40c2-8eb3-bc7691410ef8",
											"element" "2a7aa308-e400-47f5-9eca-5e034df3ee9a",
											"element" "96701921-9891-4d83-b4cd-858fb391b5b3",
											"element" "4e82515d-2322-4544-afbb-5e45859496f3",
											"element" "2642303b-faea-41e7-99f5-a9fdd99985f4",
											"element" "da6d13be-38f7-4fad-a345-c5f120c1f5c9"
										]
									},
									"DmeChannelsClip"
									{
										"id" "elementid" "463c95e5-e464-4d89-b74a-b1bd95b340f4"
										"name" "string" "tiny_ambient_lvl4_1"
										"timeFrame" "DmeTimeFrame"
										{
											"id" "elementid" "e5311bb2-d703-41ba-9650-ef79616c5aa5"
											"start" "time" "-5.0000"
											"duration" "time" "70.0000"
											"offset" "time" "0.0000"
											"scale" "float" "1"
										}

										"color" "color" "0 0 0 0"
										"text" "string" ""
										"mute" "bool" "0"
										"trackGroups" "element_array"
										[
										]
										"displayScale" "float" "1"
										"channels" "element_array"
										[
											"DmeChannel"
											{
												"id" "elementid" "6cd51a8d-8389-4cb3-849d-fd587965a6af"
												"name" "string" "emitting channel"
												"fromElement" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
												"fromAttribute" "string" "emitting"
												"fromIndex" "int" "0"
												"toElement" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
												"toAttribute" "string" "emitting"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "07c20439-bb14-4667-91e1-607f6c977ead"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "1a3a7437-5ad1-4b3e-83e6-d7a8bd6b7b17"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "5407dc6e-2fb9-487e-ae57-0b51f4e7a52e"
												"name" "string" "visible channel"
												"fromElement" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
												"fromAttribute" "string" "visible"
												"fromIndex" "int" "0"
												"toElement" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
												"toAttribute" "string" "visible"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "b69675bd-abef-4373-aed1-e5a923e35933"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "782e7a4c-0691-4dd7-a903-07412f2e0416"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"DmeChannel"
											{
												"id" "elementid" "0714c17e-f149-4332-8a01-6960eeeca9e0"
												"name" "string" "simulating channel"
												"fromElement" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
												"fromAttribute" "string" "simulating"
												"fromIndex" "int" "0"
												"toElement" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
												"toAttribute" "string" "simulating"
												"toIndex" "int" "0"
												"mode" "int" "3"
												"log" "DmeBoolLog"
												{
													"id" "elementid" "97f0f749-ca60-4881-9c70-5e0459ea5d60"
													"name" "string" "bool log"
													"layers" "element_array"
													[
														"DmeBoolLogLayer"
														{
															"id" "elementid" "388d1e20-69e3-4e07-a488-0f9586506aba"
															"name" "string" "bool log"
															"times" "time_array"
															[
																"4.9999",
																"5.0000",
																"65.0000",
																"65.0001"
															]
															"curvetypes" "int_array"
															[
															]
															"values" "bool_array"
															[
																"0",
																"1",
																"1",
																"0"
															]
															"compressed" "binary"
															"
															"
														}
													]
													"curveinfo" "element" ""
													"usedefaultvalue" "bool" "0"
													"defaultvalue" "bool" "0"
													"bookmarks" "time_array"
													[
													]
												}

											},
											"element" "aa2a8e2b-1af9-4ca9-ba55-83efd3209ada",
											"element" "cf123212-f199-4266-8ea8-155f4b73d9f2",
											"element" "a2e04abf-d930-4621-a0b3-293109049cc9",
											"element" "57879db1-a97e-4169-ae6d-47be07b78027",
											"element" "e1c8b939-e873-4fdd-b62f-4d08383f2eae",
											"element" "7db51dd7-e3d8-462f-8c58-bdd1af4ea1d5"
										]
									},
									"DmeChannelsClip"
									{
										"id" "elementid" "525ddab3-9b41-4a85-a053-885d863d0847"
										"name" "string" "tiny_astral_order_weapon1"
										"timeFrame" "DmeTimeFrame"
										{
											"id" "elementid" "66fedcd7-68aa-4376-a5a3-65f74cfdc84f"
											"start" "time" "-5.0000"
											"duration" "time" "70.0000"
											"offset" "time" "0.0000"
											"scale" "float" "1"
										}

										"color" "color" "0 0 0 0"
										"text" "string" ""
										"mute" "bool" "0"
										"trackGroups" "element_array"
										[
										]
										"displayScale" "float" "1"
										"channels" "element_array"
										[
											"element" "81aab563-bff3-46e9-8ce0-3b4ad6e18af9",
											"element" "c9ba8722-8019-49f5-b925-967d3bc8f6b4",
											"element" "b45cf80f-5cd3-4f88-bf03-ac9f7edd1f95",
											"element" "4cd2d622-dc89-47d7-85d8-87802f11914e",
											"element" "596f1f13-8dc9-4a0a-873a-0cc2e7db1543",
											"element" "ee51bb57-8e2a-42d6-8043-e6c8322be985",
											"element" "9322f48f-e1b2-4cff-8d19-cced73e54688",
											"element" "cfaeae32-d0bf-42fc-ad03-25b4c0098583",
											"element" "bd3f8682-e0fd-45a5-a611-c96289a9cdbd",
											"element" "8de4e74e-db6a-4c11-8093-fe038d0958ff",
											"element" "f2185719-27af-4758-a374-e6c3eba40271",
											"element" "fa103cd3-7939-4390-a015-b11cb5f05130",
											"element" "42534b50-88e7-44bf-bcff-69e80c2a94a8",
											"element" "21aac465-9970-4798-8931-fc9ab8333dcd",
											"element" "5caab578-a954-4df7-94e1-85240eb9eb97",
											"element" "766e97c0-4091-46dd-9873-822c829438e9",
											"element" "9c711cba-f14d-4cb6-9fd0-07f0d1227e75",
											"element" "6f10739d-6fcd-4946-b6f3-c63e3020ecec",
											"element" "a483bcbf-0af0-41f7-b1aa-8bb0f79efd29",
											"element" "ccd49d4a-f680-4dd5-9b58-ab7acbc4ba6f",
											"element" "eec521a0-2120-484a-b0c4-e045141b37fb",
											"element" "9ff7c6d1-ee31-4ff8-ad9e-6e9c600c10ae",
											"element" "e66f6c54-ea31-46a9-9aa5-2b262bb2d951",
											"element" "d5a1ec37-801e-4896-ae6c-bfc4513992c8"
										]
									}
								]
								"collapsed" "bool" "1"
								"mute" "bool" "0"
								"synched" "bool" "1"
								"clipType" "int" "0"
								"volume" "float" "1"
								"displayScale" "float" "1"
							}
						]
						"visible" "bool" "1"
						"mute" "bool" "0"
						"displayScale" "float" "1"
						"minimized" "bool" "0"
						"volume" "float" "1"
						"forcemultitrack" "bool" "0"
					}
				]
				"displayScale" "float" "1"
				"mapname" "string" ""
				"camera" "DmeCamera"
				{
					"id" "elementid" "29e5cfcb-d276-4dff-8a8c-c0095f8d8884"
					"name" "string" "camera1"
					"transform" "DmeTransform"
					{
						"id" "elementid" "a3d8190a-4cc1-499d-82ec-e54bcb44105a"
						"position" "vector3" "-1657.3859863281 -1764.8508300781 1462.5373535156"
						"orientation" "quaternion" "-0.0299309008 0.0761209726 0.3647043407 0.9275238514"
					}

					"shape" "element" ""
					"visible" "bool" "1"
					"children" "element_array"
					[
					]
					"fieldOfView" "float" "45"
					"znear" "float" "12"
					"zfar" "float" "24000"
					"micDistance" "float" "0"
					"eyeOffset" "float" "0"
					"focalDistance" "float" "72"
					"aperture" "float" "0.200000003"
					"shutterSpeed" "time" "0.0208"
					"toneMapScale" "float" "1"
					"SSAOBias" "float" "0.5"
					"SSAOStrength" "float" "2"
					"SSAORadius" "float" "30"
					"depthOfFieldSamples" "int" "64"
					"motionBlurSamples" "int" "8"
				}

				"scene" "DmeDag"
				{
					"id" "elementid" "6fdc315b-3e74-4768-b9a8-200653a67dc7"
					"name" "string" "scene"
					"transform" "DmeTransform"
					{
						"id" "elementid" "75b31d5e-73c6-4ab3-997e-bed4d7d177ce"
						"position" "vector3" "0 0 0"
						"orientation" "quaternion" "0 0 0 1"
					}

					"shape" "element" ""
					"visible" "bool" "1"
					"children" "element_array"
					[
						"DmeDag"
						{
							"id" "elementid" "33f8f763-9c36-4bb3-9afe-6666b25fae61"
							"name" "string" "tiny_01_1"
							"transform" "DmeTransform"
							{
								"id" "elementid" "31b73e55-c061-49a8-a5b9-1ee7ccf174eb"
								"position" "vector3" "0 0 0"
								"orientation" "quaternion" "0 0 0 1"
							}

							"shape" "element" ""
							"visible" "bool" "1"
							"children" "element_array"
							[
							]
						},
						"DmeDag"
						{
							"id" "elementid" "6cc48e98-c48c-4fb4-95c4-4ee93cb2d2a4"
							"name" "string" "tiny_04_1"
							"transform" "DmeTransform"
							{
								"id" "elementid" "6b88d339-59a4-47cb-a204-2bcd3bd7f4b9"
								"position" "vector3" "0 0 0"
								"orientation" "quaternion" "0 0 0 1"
							}

							"shape" "element" ""
							"visible" "bool" "1"
							"children" "element_array"
							[
								"DmeDag"
								{
									"id" "elementid" "d28da9cb-5f7d-4b3e-ab5c-ebcff30e4418"
									"name" "string" "tiny_04_1"
									"transform" "DmeTransform"
									{
										"id" "elementid" "e8b6ca76-0795-4489-a5c3-b185f56e5ef0"
										"position" "vector3" "0 0 0"
										"orientation" "quaternion" "0 0 0 1"
									}

									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
										"element" "94de5e75-487b-4aaa-81c2-8de83928762c"
									]
								},
								"DmeDag"
								{
									"id" "elementid" "9b607e9b-f666-477f-936f-3c419d9a46db"
									"name" "string" "tiny_ambient1"
									"transform" "DmeTransform"
									{
										"id" "elementid" "c5672288-97e5-4f99-b08a-d73d44a1b9a1"
										"position" "vector3" "0 0 0"
										"orientation" "quaternion" "0 0 0 1"
									}

									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
										"element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
									]
								},
								"DmeDag"
								{
									"id" "elementid" "0463e5c2-7bb9-4dc1-a812-51299e795773"
									"name" "string" "tiny_ambient_lvl2_1"
									"transform" "DmeTransform"
									{
										"id" "elementid" "32436532-63dd-423a-bc5b-93fda67356c4"
										"position" "vector3" "0 0 0"
										"orientation" "quaternion" "0 0 0 1"
									}

									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
										"element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
									]
								},
								"DmeDag"
								{
									"id" "elementid" "28a2140e-0667-4c60-9921-96c0bfe05856"
									"name" "string" "tiny_ambient_lvl3_1"
									"transform" "DmeTransform"
									{
										"id" "elementid" "a430bcce-f9ee-4cf5-a31c-defe5d235b2d"
										"position" "vector3" "0 0 0"
										"orientation" "quaternion" "0 0 0 1"
									}

									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
										"element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
									]
								},
								"DmeDag"
								{
									"id" "elementid" "1318c845-674d-4299-b327-320a1e0378c2"
									"name" "string" "tiny_ambient_lvl4_1"
									"transform" "DmeTransform"
									{
										"id" "elementid" "4f282707-b0da-4cd2-820d-0f2d00bedaa2"
										"position" "vector3" "0 0 0"
										"orientation" "quaternion" "0 0 0 1"
									}

									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
										"element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
									]
								}
							]
						},
						"DmeDag"
						{
							"id" "elementid" "b79a7057-afb5-4225-8c69-2d28085b12d3"
							"name" "string" "tiny_astral_order_weapon1"
							"transform" "DmeTransform"
							{
								"id" "elementid" "95dd32db-8a34-4f37-a56a-103eb6be9b9e"
								"position" "vector3" "0 0 0"
								"orientation" "quaternion" "0 0 0 1"
							}

							"shape" "element" ""
							"visible" "bool" "1"
							"children" "element_array"
							[
								"element" "4df12148-2b26-40f0-866d-4f65cc381789"
							]
						}
					]
				}

				"globalState" "element" ""
				"fadeIn" "time" "0.0000"
				"fadeOut" "time" "0.0000"
				"backgroundColor" "color" "64 64 64 255"
				"backgroundFXClip" "element" ""
				"operators" "element_array"
				[
				]
				"animationSets" "element_array"
				[
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
					},
					"DmeAnimationSet"
					{
						"id" "elementid" "20de8e21-d221-45dc-8cb0-6419ee7f1610"
						"name" "string" "tiny_ambient1"
						"controls" "element_array"
						[
							"element" "115ca4ce-aa44-411b-99aa-a7b8b9d20ba6",
							"element" "47864928-6ec7-46e5-8fdf-65534f682f4f",
							"element" "a57a2c95-5640-40da-b4df-c787e402a6ae"
						]
						"presetGroups" "element_array"
						[
						]
						"phonememap" "element_array"
						[
						]
						"operators" "element_array"
						[
						]
						"rootControlGroup" "DmeControlGroup"
						{
							"id" "elementid" "585100f1-9721-4199-8986-72801dccc99e"
							"children" "element_array"
							[
								"DmeControlGroup"
								{
									"id" "elementid" "72b04cfa-1e8c-4af7-8076-b86f2e01eb27"
									"name" "string" "all"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "115ca4ce-aa44-411b-99aa-a7b8b9d20ba6",
										"element" "47864928-6ec7-46e5-8fdf-65534f682f4f",
										"element" "a57a2c95-5640-40da-b4df-c787e402a6ae"
									]
									"groupColor" "color" "200 200 200 255"
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

						"particle system" "element" "2dbc7f72-6fcc-4ad6-bd0e-30c02b301258"
					},
					"DmeAnimationSet"
					{
						"id" "elementid" "ebba9276-c227-479e-a519-ff66c4c5dc1e"
						"name" "string" "tiny_ambient_lvl2_1"
						"controls" "element_array"
						[
							"element" "fe51a9c4-403e-41bb-bac6-d6ae651df17d",
							"element" "36ee3ec0-cad2-413d-9bbf-de08e91802ed",
							"element" "02e96aeb-ebd5-4e8b-998b-23c3785eae9d"
						]
						"presetGroups" "element_array"
						[
						]
						"phonememap" "element_array"
						[
						]
						"operators" "element_array"
						[
						]
						"rootControlGroup" "DmeControlGroup"
						{
							"id" "elementid" "935be86b-2250-4483-bcc5-aab00ded945e"
							"children" "element_array"
							[
								"DmeControlGroup"
								{
									"id" "elementid" "ed0e26ab-4aad-43b0-89f8-de903aa1a2b1"
									"name" "string" "all"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "fe51a9c4-403e-41bb-bac6-d6ae651df17d",
										"element" "36ee3ec0-cad2-413d-9bbf-de08e91802ed",
										"element" "02e96aeb-ebd5-4e8b-998b-23c3785eae9d"
									]
									"groupColor" "color" "200 200 200 255"
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

						"particle system" "element" "30f072bd-d2fe-4763-8cc3-33f2af2b9375"
					},
					"DmeAnimationSet"
					{
						"id" "elementid" "74830dbe-2811-451f-9811-383bc301b02f"
						"name" "string" "tiny_ambient_lvl3_1"
						"controls" "element_array"
						[
							"element" "d2669d37-d613-4c3c-8c82-d7bb18c59f40",
							"element" "a0e50c23-bced-46ff-97f0-bed0c1a47da3",
							"element" "2bbe66b9-27bb-4c5f-bfea-455b4abe1e8e"
						]
						"presetGroups" "element_array"
						[
						]
						"phonememap" "element_array"
						[
						]
						"operators" "element_array"
						[
						]
						"rootControlGroup" "DmeControlGroup"
						{
							"id" "elementid" "988740da-d33e-4658-9b33-f27355626461"
							"children" "element_array"
							[
								"DmeControlGroup"
								{
									"id" "elementid" "eeb64c53-3cc1-4bd6-bb13-887c8b4d6244"
									"name" "string" "all"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "d2669d37-d613-4c3c-8c82-d7bb18c59f40",
										"element" "a0e50c23-bced-46ff-97f0-bed0c1a47da3",
										"element" "2bbe66b9-27bb-4c5f-bfea-455b4abe1e8e"
									]
									"groupColor" "color" "200 200 200 255"
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

						"particle system" "element" "beff30f4-b4a5-4c15-ae6e-e22ab609e085"
					},
					"DmeAnimationSet"
					{
						"id" "elementid" "4fb2e2ff-236f-4539-8ecf-663dd2a8f94a"
						"name" "string" "tiny_ambient_lvl4_1"
						"controls" "element_array"
						[
							"element" "90457517-59ed-453e-85d3-3ec1dbfb39c1",
							"element" "a64c1456-090b-4eca-9d86-cd6c33d97948",
							"element" "376b3754-0031-40d1-abe5-d0de10edf682"
						]
						"presetGroups" "element_array"
						[
						]
						"phonememap" "element_array"
						[
						]
						"operators" "element_array"
						[
						]
						"rootControlGroup" "DmeControlGroup"
						{
							"id" "elementid" "c16b998a-ad51-4174-ac1c-15e84cef34a1"
							"children" "element_array"
							[
								"DmeControlGroup"
								{
									"id" "elementid" "e32e398d-3d52-40f6-8e75-bee9fc4e2e60"
									"name" "string" "all"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "90457517-59ed-453e-85d3-3ec1dbfb39c1",
										"element" "a64c1456-090b-4eca-9d86-cd6c33d97948",
										"element" "376b3754-0031-40d1-abe5-d0de10edf682"
									]
									"groupColor" "color" "200 200 200 255"
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

						"particle system" "element" "951d2dd6-edff-47b0-a7c5-69ffd4fc9955"
					},
					"DmeAnimationSet"
					{
						"id" "elementid" "bdfae789-9d2f-4ef7-a4be-a1da7f37253f"
						"name" "string" "tiny_astral_order_weapon1"
						"controls" "element_array"
						[
							"element" "0a710bed-a070-4209-9ad8-9d13858e0cf4",
							"element" "595ffc46-f201-48fe-9339-2c13b36f1b41",
							"element" "fa47d9ee-2d3b-4410-82f9-9a51a28a9300",
							"element" "59081215-232c-48fc-9bec-2c7c3078665f",
							"element" "07ec8e9f-466f-4c4a-94ff-2f36999f7318",
							"element" "03caa148-c706-4e06-8a30-718a92204f37",
							"element" "c88657a5-8fd9-443c-8556-06a2bda07e60",
							"element" "8f083edf-7f98-4cd5-8fc4-fa93515bd5fc",
							"element" "3576ac1e-6ed9-4dd4-8f05-dddd37a628e4",
							"element" "91fa4141-74ff-432a-be27-58b6dd0a770f",
							"element" "ca39033d-c0c8-4360-be5d-abf63340a793",
							"element" "a057f32f-91c1-415d-ae96-87b2e86805a4"
						]
						"presetGroups" "element_array"
						[
						]
						"phonememap" "element_array"
						[
						]
						"operators" "element_array"
						[
						]
						"rootControlGroup" "DmeControlGroup"
						{
							"id" "elementid" "ed748c88-1446-4912-9ff1-d606df2bd95f"
							"children" "element_array"
							[
								"DmeControlGroup"
								{
									"id" "elementid" "19bcf18c-fd7e-4eb4-a0d0-fdae1f750830"
									"name" "string" "Body"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "0a710bed-a070-4209-9ad8-9d13858e0cf4"
									]
									"groupColor" "color" "0 128 255 255"
									"controlColor" "color" "200 200 200 255"
									"visible" "bool" "1"
									"selectable" "bool" "1"
									"snappable" "bool" "1"
								},
								"DmeControlGroup"
								{
									"id" "elementid" "24fd0c56-8f24-44a1-9008-fc650b8a249b"
									"name" "string" "Unknown"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "595ffc46-f201-48fe-9339-2c13b36f1b41",
										"element" "fa47d9ee-2d3b-4410-82f9-9a51a28a9300",
										"element" "07ec8e9f-466f-4c4a-94ff-2f36999f7318",
										"element" "03caa148-c706-4e06-8a30-718a92204f37",
										"element" "c88657a5-8fd9-443c-8556-06a2bda07e60",
										"element" "8f083edf-7f98-4cd5-8fc4-fa93515bd5fc",
										"element" "3576ac1e-6ed9-4dd4-8f05-dddd37a628e4",
										"element" "ca39033d-c0c8-4360-be5d-abf63340a793",
										"element" "a057f32f-91c1-415d-ae96-87b2e86805a4"
									]
									"groupColor" "color" "0 128 255 255"
									"controlColor" "color" "200 200 200 255"
									"visible" "bool" "1"
									"selectable" "bool" "1"
									"snappable" "bool" "1"
								},
								"DmeControlGroup"
								{
									"id" "elementid" "d5fbedc8-661f-44d2-9178-8004d7ed7a1a"
									"name" "string" "Procedural"
									"children" "element_array"
									[
									]
									"controls" "element_array"
									[
										"element" "59081215-232c-48fc-9bec-2c7c3078665f",
										"element" "91fa4141-74ff-432a-be27-58b6dd0a770f"
									]
									"groupColor" "color" "0 128 255 255"
									"controlColor" "color" "200 200 200 255"
									"visible" "bool" "0"
									"selectable" "bool" "0"
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

						"gameModel" "element" "4df12148-2b26-40f0-866d-4f65cc381789"
					}
				]
				"bookmarkSets" "element_array"
				[
					"DmeBookmarkSet"
					{
						"id" "elementid" "dc531fdc-54b3-43f6-9130-7fd4cfc1a18a"
						"name" "string" "Default Set"
						"bookmarks" "element_array"
						[
						]
					}
				]
				"activeBookmarkSet" "int" "0"
				"subClipTrackGroup" "element" ""
				"volume" "float" "1"
				"concommands" "string_array"
				[
				]
				"convars" "string_array"
				[
				]
			}
		]
		"collapsed" "bool" "0"
		"mute" "bool" "0"
		"synched" "bool" "1"
		"clipType" "int" "3"
		"volume" "float" "1"
		"displayScale" "float" "1"
	}

*/
