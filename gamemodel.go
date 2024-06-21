package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type GameModel struct {
	*Node
	ModelName               string
	Skin                    int32
	FlexWeights             []float32
	MeshGroupMask           uint64
	bones                   []*Bone
	globalFlexControllers   []*GlobalFlexControllerOperator
	ComputeBounds           bool
	EvaluateProceduralBones bool
}

func NewGameModel(name string, modelName string) *GameModel {
	return &GameModel{
		Node:                    NewNode(name),
		ModelName:               modelName,
		Skin:                    0,
		FlexWeights:             make([]float32, 0),
		bones:                   make([]*Bone, 0),
		MeshGroupMask:           0xffffffffffffffff,
		globalFlexControllers:   make([]*GlobalFlexControllerOperator, 0),
		ComputeBounds:           true,
		EvaluateProceduralBones: true,
	}
}

func (gm *GameModel) isNode() {
}

func (gm *GameModel) AddChildren(child INode) {
	gm.children = append(gm.children, child)
}

func (gm *GameModel) AddBone(bone *Bone) {
	gm.bones = append(gm.bones, bone)
}

func (gm *GameModel) CreateBone(name string) *Bone {
	bone := NewBone(name)
	gm.bones = append(gm.bones, bone)
	return bone
}

func (gm *GameModel) AddGlobalFlexControllerOperator(o *GlobalFlexControllerOperator) {
	gm.globalFlexControllers = append(gm.globalFlexControllers, o)
}

func (gm *GameModel) CreateGlobalFlexControllerOperator(name string, flexWeight float32) *GlobalFlexControllerOperator {
	o := NewGlobalFlexControllerOperator(name, flexWeight, gm)
	gm.globalFlexControllers = append(gm.globalFlexControllers, o)
	return o
}

func (gm *GameModel) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement(gm.Name, "DmeGameModel")

	e.CreateStringAttribute("modelName", gm.ModelName)
	e.CreateIntAttribute("skin", gm.Skin)
	e.CreateElementAttribute("transform", serializer.GetElement(gm.Transform))
	e.CreateBoolAttribute("visible", gm.Visible)
	e.CreateUint64Attribute("meshGroupMask", gm.MeshGroupMask)
	e.CreateBoolAttribute("computeBounds", gm.ComputeBounds)
	e.CreateBoolAttribute("evaluateProceduralBones", gm.EvaluateProceduralBones)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range gm.children {
		children.PushElement(serializer.GetElement(child))
	}
	for _, bone := range gm.bones {
		children.PushElement(serializer.GetElement(bone))
	}

	flexWeights := e.CreateAttribute("flexWeights", dmx.AT_FLOAT_ARRAY)
	for _, weights := range gm.FlexWeights {
		flexWeights.PushFloat(weights)
	}

	bones := e.CreateAttribute("bones", dmx.AT_ELEMENT_ARRAY)
	for _, bone := range gm.bones {
		bones.PushElement(serializer.GetElement(bone.Transform))
	}


	globalFlexControllers := e.CreateAttribute("globalFlexControllers", dmx.AT_ELEMENT_ARRAY)
	for _, gfc := range gm.globalFlexControllers {
		globalFlexControllers.PushElement(serializer.GetElement(gfc))
	}
	/*
		controls := e.CreateAttribute("controls", dmx.AT_ELEMENT_ARRAY)
		for _, c := range gm.controls {
			controls.PushElement(serializer.GetElement(c))
		}

		presetGroups := e.CreateAttribute("presetGroups", dmx.AT_ELEMENT_ARRAY)
		for _, pg := range gm.presetGroups {
			presetGroups.PushElement(serializer.GetElement(pg))
		}

		operators := e.CreateAttribute("operators", dmx.AT_ELEMENT_ARRAY)
		for _, o := range gm.operators {
			operators.PushElement(serializer.GetElement(o))
		}*/

	return e
}

/*
"DmeGameModel"
{
	"id" "elementid" "94de5e75-487b-4aaa-81c2-8de83928762c"
	"name" "string" "tiny_04"
	"transform" "element" "9168f52d-a1ab-47f5-8e90-4b09e5c88eba"
	"shape" "element" ""
	"visible" "bool" "1"
	"children" "element_array"
	[
		"DmeDag"
		{
			"id" "elementid" "26706960-e37b-4c0b-9606-04439b818cf0"
			"name" "string" "bone 0 (root)"
			"transform" "element" "14af0a26-8b83-4723-89f6-7201842ec82a"
			"shape" "element" ""
			"visible" "bool" "1"
			"children" "element_array"
			[
				"DmeDag"
				{
					"id" "elementid" "eca426e9-2208-4833-9d34-baa7f4b81686"
					"name" "string" "bone 1 (spine1)"
					"transform" "element" "5f8bec3f-5b2e-4b6a-bde3-f768a751b8aa"
					"shape" "element" ""
					"visible" "bool" "1"
					"children" "element_array"
					[
						"DmeDag"
						{
							"id" "elementid" "43381f59-2379-419d-9e2e-6d4b5b516e5c"
							"name" "string" "bone 2 (spine2)"
							"transform" "element" "eb7b46a9-9fb7-4329-a984-f538ce3d90a4"
							"shape" "element" ""
							"visible" "bool" "1"
							"children" "element_array"
							[
								"DmeDag"
								{
									"id" "elementid" "824ac100-bae4-4a9e-a6be-c0432c12ea62"
									"name" "string" "bone 3 (spine3)"
									"transform" "element" "4c50b4f7-1044-4e77-ac25-a6ddd96c77cb"
									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
										"DmeDag"
										{
											"id" "elementid" "7e278023-0324-4fec-a849-75fc3f8dca23"
											"name" "string" "bone 4 (clavicle_R)"
											"transform" "element" "1bd1db67-abfc-4ec2-9603-c7ec30dc1f85"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
												"DmeDag"
												{
													"id" "elementid" "495bfe8b-56de-4794-a078-292080371ca7"
													"name" "string" "bone 5 (bicep_R)"
													"transform" "element" "b0e37774-7d9c-469d-8bc6-25b244aa53b6"
													"shape" "element" ""
													"visible" "bool" "1"
													"children" "element_array"
													[
														"DmeDag"
														{
															"id" "elementid" "9b7d7aca-0e24-42ab-a4f2-d13dad0e000f"
															"name" "string" "bone 6 (elbow_R)"
															"transform" "element" "d94d28cb-65eb-4f35-a986-c8625d80f70d"
															"shape" "element" ""
															"visible" "bool" "1"
															"children" "element_array"
															[
																"DmeDag"
																{
																	"id" "elementid" "e86f9729-fe21-4eef-972d-6fee23d20e1e"
																	"name" "string" "bone 7 (wrist_R)"
																	"transform" "element" "022886fe-7a8c-44f2-b403-be340dd6273d"
																	"shape" "element" ""
																	"visible" "bool" "1"
																	"children" "element_array"
																	[
																	]
																}
															]
														}
													]
												}
											]
										},
										"DmeDag"
										{
											"id" "elementid" "09cc1692-5087-461e-beb1-cb9e3a6610fb"
											"name" "string" "bone 8 (clavicle_L)"
											"transform" "element" "20ea02f4-ac22-4b3b-9125-dc073d4cac0d"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
												"DmeDag"
												{
													"id" "elementid" "b6b9da4d-b95a-4ec4-aa0b-5c4e7711f378"
													"name" "string" "bone 9 (bicep_L)"
													"transform" "element" "25a8de6a-bc88-4923-8190-f1e203021b38"
													"shape" "element" ""
													"visible" "bool" "1"
													"children" "element_array"
													[
														"DmeDag"
														{
															"id" "elementid" "63ead905-2838-46fa-b649-cc4eb5e0150a"
															"name" "string" "bone 10 (elbow_L)"
															"transform" "element" "ed56a90e-e1d6-4a11-8b00-399e1b194660"
															"shape" "element" ""
															"visible" "bool" "1"
															"children" "element_array"
															[
																"DmeDag"
																{
																	"id" "elementid" "e9b8e299-520a-4f74-bdff-933fb5920932"
																	"name" "string" "bone 11 (wrist_L)"
																	"transform" "element" "433e45d1-77a7-41ba-879e-00a3a161d353"
																	"shape" "element" ""
																	"visible" "bool" "1"
																	"children" "element_array"
																	[
																		"DmeDag"
																		{
																			"id" "elementid" "0629a217-2daa-462a-acca-38fec79e5799"
																			"name" "string" "bone 21 (index_0_L)"
																			"transform" "element" "0ede7ef0-e71a-4119-9229-fcf7acd8f95d"
																			"shape" "element" ""
																			"visible" "bool" "1"
																			"children" "element_array"
																			[
																				"DmeDag"
																				{
																					"id" "elementid" "41a8dbc5-58cc-4ef2-b130-ceb75343e6c8"
																					"name" "string" "bone 22 (index_1_L)"
																					"transform" "element" "b61478a0-f1e5-4e46-9544-0e22372a0679"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																					]
																				}
																			]
																		},
																		"DmeDag"
																		{
																			"id" "elementid" "fa1252ee-adfb-40e6-8787-bb01ff2a9d4a"
																			"name" "string" "bone 23 (thumb_0_L)"
																			"transform" "element" "75290f10-4cb4-4248-a6e3-a9bcbd7da593"
																			"shape" "element" ""
																			"visible" "bool" "1"
																			"children" "element_array"
																			[
																				"DmeDag"
																				{
																					"id" "elementid" "b90f9930-6585-41a3-a74c-cc7a384828cb"
																					"name" "string" "bone 24 (thumb_1_L)"
																					"transform" "element" "145c9be7-5438-4dc2-b7c5-dcaff8d9f033"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																					]
																				}
																			]
																		},
																		"DmeDag"
																		{
																			"id" "elementid" "377f979b-1fc1-4089-ad5a-295a7573e1b5"
																			"name" "string" "bone 25 (mid_0_L)"
																			"transform" "element" "37fb7a38-2c06-4b18-b42d-34ff407d39b5"
																			"shape" "element" ""
																			"visible" "bool" "1"
																			"children" "element_array"
																			[
																				"DmeDag"
																				{
																					"id" "elementid" "529e204f-36e3-4af7-b2c7-31b98f3e1f43"
																					"name" "string" "bone 26 (mid_1_L)"
																					"transform" "element" "c52e6406-7a5a-4e71-85f5-6203450b69f8"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																					]
																				}
																			]
																		},
																		"DmeDag"
																		{
																			"id" "elementid" "d1d8bdf1-0b03-4a25-a13f-cf7e59c33311"
																			"name" "string" "bone 27 (Tree_1)"
																			"transform" "element" "1b02b518-0769-4523-82c6-b949fa794b73"
																			"shape" "element" ""
																			"visible" "bool" "1"
																			"children" "element_array"
																			[
																				"DmeDag"
																				{
																					"id" "elementid" "bfe2139e-80c4-4e2f-9c57-1e936f2baaa1"
																					"name" "string" "bone 28 (RootA2_0)"
																					"transform" "element" "0c85c804-e5f3-44f1-9c1a-cc8119f97d46"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																						"DmeDag"
																						{
																							"id" "elementid" "57c6a3bc-08c0-434b-8b15-13d5f1e8cfde"
																							"name" "string" "bone 29 (RootA2_1)"
																							"transform" "element" "4859267e-9321-405b-9f78-f1cb9732b0cd"
																							"shape" "element" ""
																							"visible" "bool" "1"
																							"children" "element_array"
																							[
																							]
																						}
																					]
																				},
																				"DmeDag"
																				{
																					"id" "elementid" "e43916d1-859d-481d-8e84-03b2527b3f85"
																					"name" "string" "bone 30 (RootB3_0)"
																					"transform" "element" "7359dcfd-4a31-40de-8079-b3c671e24101"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																					]
																				},
																				"DmeDag"
																				{
																					"id" "elementid" "f2deee9e-cbbb-40a2-9367-b96c34242b05"
																					"name" "string" "bone 31 (RootC4_0)"
																					"transform" "element" "cc4caaf1-f884-48d3-8255-e92c6da515c4"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																					]
																				},
																				"DmeDag"
																				{
																					"id" "elementid" "a7c969c3-9679-4504-9b30-e85c0dcb42be"
																					"name" "string" "bone 32 (RootD5_0)"
																					"transform" "element" "6c14dd09-fcf0-41f4-baa0-c9d9201cd7a5"
																					"shape" "element" ""
																					"visible" "bool" "1"
																					"children" "element_array"
																					[
																					]
																				},
																				"element" "822213a9-7fa8-4491-b772-eca29d0c0d3f",
																				"element" "74de7421-45ac-4b0f-b9f3-d19fd55d2416"
																			]
																		}
																	]
																}
															]
														}
													]
												},
												"DmeDag"
												{
													"id" "elementid" "aacc3eb0-b9a6-4a57-9f34-74a99f279c75"
													"name" "string" "bone 20 (shoulder_rock_JNT)"
													"transform" "element" "60370b08-228a-4432-b002-55c4251a1c3f"
													"shape" "element" ""
													"visible" "bool" "1"
													"children" "element_array"
													[
													]
												}
											]
										},
										"DmeDag"
										{
											"id" "elementid" "e3ee487e-d0fc-4476-bf71-196d37833be6"
											"name" "string" "bone 12 (neck1)"
											"transform" "element" "9b7a839d-c4a9-41d4-b0a8-68023c83054b"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
												"DmeDag"
												{
													"id" "elementid" "a32eff05-7e68-4724-bb54-763cfb86c242"
													"name" "string" "bone 13 (head)"
													"transform" "element" "aee9b8e3-fddc-4fb6-bcee-0989f111291e"
													"shape" "element" ""
													"visible" "bool" "1"
													"children" "element_array"
													[
														"DmeDag"
														{
															"id" "elementid" "0921f148-96dc-4cfd-bddd-376e9cac7c36"
															"name" "string" "bone 40 (jaw0_0)"
															"transform" "element" "4b245b33-5a67-48ff-bfbb-f08026f816d3"
															"shape" "element" ""
															"visible" "bool" "1"
															"children" "element_array"
															[
																"DmeDag"
																{
																	"id" "elementid" "81a16f2b-a3bc-4a38-a514-822258270b5d"
																	"name" "string" "bone 41 (beard_LF)"
																	"transform" "element" "8fc1c63a-b241-45b7-b22c-3b01082286e5"
																	"shape" "element" ""
																	"visible" "bool" "1"
																	"children" "element_array"
																	[
																	]
																},
																"DmeDag"
																{
																	"id" "elementid" "28c36fff-7a76-427b-9ee5-ef4ef28fb966"
																	"name" "string" "bone 42 (beard_RT)"
																	"transform" "element" "e5a0880d-9887-4202-a0ea-dce93802c50d"
																	"shape" "element" ""
																	"visible" "bool" "1"
																	"children" "element_array"
																	[
																	]
																}
															]
														}
													]
												}
											]
										},
										"DmeDag"
										{
											"id" "elementid" "08e5e2ad-d2a6-4a47-9379-4d5fa57f203f"
											"name" "string" "bone 19 (mountain_A)"
											"transform" "element" "f11d4724-c59d-4e1f-806f-f1d64307af5e"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "14781b04-3d44-4748-859c-d7d7730a0ab9"
											"name" "string" "bone 43 (mountain_B)"
											"transform" "element" "44e34340-0e78-449b-beab-b49c6d372435"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "c9cfba2f-b9e4-4b59-b128-b9ba3d31a3e7"
											"name" "string" "bone 44 (mountain_C)"
											"transform" "element" "b9024163-f802-4f65-94f4-e2195a03ed46"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "2c0d3bdc-b861-4565-a300-e9510f40a822"
											"name" "string" "bone 45 (mountain_D)"
											"transform" "element" "df0c5d14-9074-4fbd-af87-11d86357447e"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "54638c3f-14fc-41ea-8e59-749b2042c1d3"
											"name" "string" "bone 46 (mountain_E)"
											"transform" "element" "5e424695-6946-4e0a-825a-b930424c6825"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "a3ea6cc5-a030-4f68-95c7-4e6c6da9c05d"
											"name" "string" "bone 47 (mountain_F)"
											"transform" "element" "d80bd096-fb51-4463-97ac-9f63b4afbab1"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "b7063b4d-22a6-494c-9652-4d0a68d2d285"
											"name" "string" "bone 48 (mountain_G)"
											"transform" "element" "74f6db80-3015-4bef-86bb-f9eaf8bf222f"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "3f603b62-5ed5-4cd1-a8f9-12a1ba2e192f"
											"name" "string" "bone 49 (mountain_H)"
											"transform" "element" "07ffab5d-ab03-4610-812b-e4714bb1e62e"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										},
										"DmeDag"
										{
											"id" "elementid" "09e7541b-6442-4949-9946-d67cc7664e53"
											"name" "string" "bone 50 (mountain_I)"
											"transform" "element" "32bf4039-257e-4f00-904b-c06b921cf528"
											"shape" "element" ""
											"visible" "bool" "1"
											"children" "element_array"
											[
											]
										}
									]
								}
							]
						}
					]
				},
				"DmeDag"
				{
					"id" "elementid" "5b2f25e9-c9f5-4cf7-8489-fd7bf2fa07a1"
					"name" "string" "bone 14 (thigh_R)"
					"transform" "element" "7069c4c0-1c14-42cc-922a-f769f6c0dac8"
					"shape" "element" ""
					"visible" "bool" "1"
					"children" "element_array"
					[
						"DmeDag"
						{
							"id" "elementid" "af5bca71-cdd7-4ae8-ba04-8255eb05cb88"
							"name" "string" "bone 15 (knee_R)"
							"transform" "element" "26d99d41-a164-4dfb-9af3-21a7a8c52c4a"
							"shape" "element" ""
							"visible" "bool" "1"
							"children" "element_array"
							[
								"DmeDag"
								{
									"id" "elementid" "2172c0da-95cc-46b9-a85d-02daf89f8a55"
									"name" "string" "bone 16 (ankle_R)"
									"transform" "element" "145b62e1-eb79-4369-b622-177f5bb2652f"
									"shape" "element" ""
									"visible" "bool" "1"
									"children" "element_array"
									[
									]
								}
							]
						}
					]
				},
				"DmeDag"
				{
					"id" "elementid" "f970348d-e93d-497d-9dad-917f01384750"
					"name" "string" "bone 17 (thigh_L)"
					"transform" "element" "f8702ebb-5f39-4789-969a-e96b6ecffaa7"
					"shape" "element" ""
					"visible" "bool" "1"
					"children" "element_array"
					[
						"DmeDag"
						{
							"id" "elementid" "ec250164-7f84-44bb-9614-72810c7479e0"
							"name" "string" "bone 18 (knee_L)"
							"transform" "element" "75087a6a-7536-45fe-b75b-fe53814cb0fa"
							"shape" "element" ""
							"visible" "bool" "1"
							"children" "element_array"
							[
							]
						}
					]
				}
			]
		}
	]
	"flexWeights" "float_array"
	[
		"0",
		"0",
		"0",
		"0.5",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0.5",
		"0",
		"0",
		"0.5",
		"0",
		"0.5",
		"0.5",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0.5",
		"0",
		"0.5",
		"0.5",
		"0"
	]
	"modelName" "string" "models/heroes/tiny/tiny_04/tiny_04.vmdl"
	"skin" "int" "0"
	"meshGroupMask" "uint64" "0xffffffffffffffff"
	"bones" "element_array"
	[
		"element" "14af0a26-8b83-4723-89f6-7201842ec82a",
		"element" "5f8bec3f-5b2e-4b6a-bde3-f768a751b8aa",
		"element" "eb7b46a9-9fb7-4329-a984-f538ce3d90a4",
		"element" "4c50b4f7-1044-4e77-ac25-a6ddd96c77cb",
		"element" "1bd1db67-abfc-4ec2-9603-c7ec30dc1f85",
		"element" "b0e37774-7d9c-469d-8bc6-25b244aa53b6",
		"element" "d94d28cb-65eb-4f35-a986-c8625d80f70d",
		"element" "022886fe-7a8c-44f2-b403-be340dd6273d",
		"element" "20ea02f4-ac22-4b3b-9125-dc073d4cac0d",
		"element" "25a8de6a-bc88-4923-8190-f1e203021b38",
		"element" "ed56a90e-e1d6-4a11-8b00-399e1b194660",
		"element" "433e45d1-77a7-41ba-879e-00a3a161d353",
		"element" "9b7a839d-c4a9-41d4-b0a8-68023c83054b",
		"element" "aee9b8e3-fddc-4fb6-bcee-0989f111291e",
		"element" "7069c4c0-1c14-42cc-922a-f769f6c0dac8",
		"element" "26d99d41-a164-4dfb-9af3-21a7a8c52c4a",
		"element" "145b62e1-eb79-4369-b622-177f5bb2652f",
		"element" "f8702ebb-5f39-4789-969a-e96b6ecffaa7",
		"element" "75087a6a-7536-45fe-b75b-fe53814cb0fa",
		"element" "f11d4724-c59d-4e1f-806f-f1d64307af5e",
		"element" "60370b08-228a-4432-b002-55c4251a1c3f",
		"element" "0ede7ef0-e71a-4119-9229-fcf7acd8f95d",
		"element" "b61478a0-f1e5-4e46-9544-0e22372a0679",
		"element" "75290f10-4cb4-4248-a6e3-a9bcbd7da593",
		"element" "145c9be7-5438-4dc2-b7c5-dcaff8d9f033",
		"element" "37fb7a38-2c06-4b18-b42d-34ff407d39b5",
		"element" "c52e6406-7a5a-4e71-85f5-6203450b69f8",
		"element" "1b02b518-0769-4523-82c6-b949fa794b73",
		"element" "0c85c804-e5f3-44f1-9c1a-cc8119f97d46",
		"element" "4859267e-9321-405b-9f78-f1cb9732b0cd",
		"element" "7359dcfd-4a31-40de-8079-b3c671e24101",
		"element" "cc4caaf1-f884-48d3-8255-e92c6da515c4",
		"element" "6c14dd09-fcf0-41f4-baa0-c9d9201cd7a5",
		"element" "65e647e3-ab93-4eff-8e6e-93920a042158",
		"element" "01b25275-e5ed-4b16-b0a3-75899060f103",
		"element" "0b8aac9c-5f3b-4427-a0c2-0e141e8e5503",
		"element" "5e0ad52f-2c63-4ca4-b539-ec58cb9966fa",
		"element" "8355b355-19f6-4f2c-b785-245a3bc3855f",
		"element" "672e38aa-c999-43e7-a99f-b51fef31d435",
		"element" "10c8dc0a-a1e3-479c-8f3d-79898f524c13",
		"element" "4b245b33-5a67-48ff-bfbb-f08026f816d3",
		"element" "8fc1c63a-b241-45b7-b22c-3b01082286e5",
		"element" "e5a0880d-9887-4202-a0ea-dce93802c50d",
		"element" "44e34340-0e78-449b-beab-b49c6d372435",
		"element" "b9024163-f802-4f65-94f4-e2195a03ed46",
		"element" "df0c5d14-9074-4fbd-af87-11d86357447e",
		"element" "5e424695-6946-4e0a-825a-b930424c6825",
		"element" "d80bd096-fb51-4463-97ac-9f63b4afbab1",
		"element" "74f6db80-3015-4bef-86bb-f9eaf8bf222f",
		"element" "07ffab5d-ab03-4610-812b-e4714bb1e62e",
		"element" "32bf4039-257e-4f00-904b-c06b921cf528"
	]
	"globalFlexControllers" "element_array"
	[
		"element" "66a1d5a0-8425-4d1f-b344-fdd442dcce4e",
		"element" "b731cefe-6405-4e37-a35f-aedc5b32015d",
		"element" "0cd6d209-1492-4f87-8eeb-11e43c9f6d70",
		"element" "d2c8abd7-a623-4792-8c05-b0eb0a60c45a",
		"element" "7d083c2b-89ea-4712-9328-2667983f826a",
		"element" "d11befc1-b654-46d3-9f0b-29d6dd24f7f0",
		"element" "3edba750-1de4-4a25-8cb7-69bcb1daf72a",
		"element" "56ccc05e-b197-406a-8f2d-c59d5cad1965",
		"element" "1e997967-f8db-49b2-a597-df8a30e34ad1",
		"element" "d0c2cd57-64a9-4fc5-bece-231604d2ee2c",
		"element" "91811d58-6541-4ae7-a4db-84a8907ea4c8",
		"element" "e23225b1-0be5-47a7-b061-293266e162d0",
		"element" "5f7a7a1a-5165-4690-a81f-8d262d8a1a62",
		"element" "7bc5a925-9428-42c4-bdf6-ee6535797113",
		"element" "ca0a6c8b-3a98-4ac6-87e3-91718541340a",
		"element" "10fc4ae8-00b5-41fd-9d28-1bc500b49670",
		"element" "007f9010-458c-4dce-982b-fae85c4b474a",
		"element" "535eeb76-3516-47b0-9617-1e2b576611cd",
		"element" "f6acff48-e0ad-4116-b7ef-a1151e78fc54",
		"element" "f87f6e7c-de11-4ec0-bf8c-f9bdae4837dd",
		"element" "0fea9b12-133e-4c00-9672-822d42e0579f",
		"element" "fba76b45-5b9a-4628-85f0-b663995b32a0",
		"element" "d13126ad-c0cb-4278-9740-27029d2d6b03",
		"element" "9a4ac8cc-5785-4471-bc84-3405693699e6",
		"element" "95dbbd87-7c04-455b-8b6c-0f777d45f373",
		"element" "cfd8096e-b9ec-438d-828e-f624ea269767",
		"element" "2de3a140-f7df-4d82-830e-d95d217a6e17",
		"element" "40d5dd4c-5d7c-4c5c-8748-9fb7215009ce"
	]
	"computeBounds" "bool" "1"
	"evaluateProceduralBones" "bool" "1"
}
*/
