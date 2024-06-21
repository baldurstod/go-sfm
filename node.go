package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type Node struct {
	Name      string
	Transform *Transform
	Visible   bool
	children  []Element
}

func NewNode(name string) *Node {
	return &Node{
		Name:      name,
		Transform: NewTransform(""),
		Visible:   true,
		children:  make([]Element, 0),
	}
}

func (s *Node) AddChildren(child Element) {
	s.children = append(s.children, child)
}

func (s *Node) toDmElement(serializer *Serializer) *dmx.DmElement {
	e := dmx.NewDmElement("DmeDag")

	e.CreateStringAttribute("name", s.Name)
	e.CreateElementAttribute("transform", serializer.GetElement(s.Transform))
	e.CreateBoolAttribute("visible", s.Visible)

	children := e.CreateAttribute("children", dmx.AT_ELEMENT_ARRAY)
	for _, child := range s.children {
		children.PushElement(serializer.GetElement(child))
	}

	return e
}

/*
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
*/
