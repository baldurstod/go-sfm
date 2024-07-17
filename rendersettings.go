package sfm

import (
	"github.com/baldurstod/go-dmx"
)

type RenderSettings struct {
	Name                     string
	Width                    int32
	Height                   int32
	FrameRate                float32
	AmbientOcclusionMode     int32
	ShowAmbientOcclusion     int32
	DrawGameRenderablesMask  int32
	DrawToolRenderablesMask  int32
	ToneMapScale             float32
	GlobalLightShadowEnabled bool
	PlaybackClampFrameCount  int32
	IgnoreAlphaFade          bool
	ProgressiveRefinement    *ProgressiveRefinementSettings
}

func NewRenderSettings() *RenderSettings {
	return &RenderSettings{
		Name:                     "renderSettings",
		Width:                    1280,
		Height:                   720,
		FrameRate:                24,
		AmbientOcclusionMode:     1,
		DrawGameRenderablesMask:  63,
		DrawToolRenderablesMask:  31,
		ToneMapScale:             1,
		GlobalLightShadowEnabled: true,
		PlaybackClampFrameCount:  5,
		IgnoreAlphaFade:          true,
		ProgressiveRefinement:    NewProgressiveRefinementSettings(),
	}
}

func (rs *RenderSettings) createDmElement(serializer *Serializer) *dmx.DmElement {
	return dmx.NewDmElement(rs.Name, "DmElement")
}

func (rs *RenderSettings) isExportable() bool {
	return true
}

func (rs *RenderSettings) toDmElement(serializer *Serializer, e *dmx.DmElement) {
	e.CreateIntAttribute("width", rs.Width)
	e.CreateIntAttribute("height", rs.Height)
	e.CreateFloatAttribute("frameRate", rs.FrameRate)
	e.CreateIntAttribute("ambientOcclusionMode", rs.AmbientOcclusionMode)
	e.CreateIntAttribute("showAmbientOcclusion", rs.ShowAmbientOcclusion)
	e.CreateIntAttribute("drawGameRenderablesMask", rs.DrawGameRenderablesMask)
	e.CreateIntAttribute("drawToolRenderablesMask", rs.DrawToolRenderablesMask)
	e.CreateFloatAttribute("toneMapScale", rs.ToneMapScale)
	e.CreateBoolAttribute("globalLightShadowEnabled", rs.GlobalLightShadowEnabled)
	e.CreateIntAttribute("playbackClampFrameCount", rs.PlaybackClampFrameCount)
	e.CreateBoolAttribute("ignoreAlphaFade", rs.IgnoreAlphaFade)
	e.CreateElementAttribute("ProgressiveRefinement", serializer.GetElement(rs.ProgressiveRefinement))
}

/*
	"renderSettings" "DmElement"
	{
		"id" "elementid" "d377d7da-f28a-45a4-92b0-4ecf4bdcfcac"
		"name" "string" "renderSettings"
		"width" "int" "1280"
		"height" "int" "720"
		"frameRate" "float" "24"
		"ambientOcclusionMode" "int" "1"
		"showAmbientOcclusion" "int" "0"
		"drawGameRenderablesMask" "int" "63"
		"drawToolRenderablesMask" "int" "31"
		"toneMapScale" "float" "1"
		"globalLightShadowEnabled" "bool" "1"
		"playbackClampFrameCount" "int" "5"
		"ignoreAlphaFade" "bool" "1"
		"ProgressiveRefinement" "DmElement"
		{
			"id" "elementid" "b2add799-236e-4b46-980a-bac07c39679c"
			"name" "string" "ProgressiveRefinementSettings"
			"on" "bool" "1"
			"useDepthOfField" "bool" "1"
			"overrideDepthOfFieldSamples" "bool" "0"
			"overrideDepthOfFieldSamplesValue" "int" "64"
			"useMotionBlur" "bool" "1"
			"overrideMotionBlurSamples" "bool" "0"
			"overrideMotionBlurSamplesValue" "int" "8"
			"useAntialiasing" "bool" "1"
			"overrideShutterSpeed" "bool" "0"
			"overrideShutterSpeedValue" "float" "0.020833334"
		}

	}

*/
