package main

import (
	"bytes"
	_ "embed"
	"image/png"

	"github.com/AllenDang/cimgui-go/imgui"
	rlig "github.com/anasrar/raylib-cimgui-go-starterkit/pkg/raylib_imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed ColorChecker.png
var colorChecker []byte

var showDemoWindow = false
var backgroundColor = [3]float32{0, 0, 0}
var cubeColor = [3]float32{0.902, 0.161, 0.216}
var wireframeColor = [3]float32{0.745, 0.129, 0.216}
var rotation = float32(0)
var autoRotate = true
var dummyString = "Welcome to the third dimension!"

func main() {
	imgColorChecker, err := png.Decode(bytes.NewReader(colorChecker))
	if err != nil {
		panic(err)
	}

	rl.InitWindow(int32(1200), int32(800), "Raylib Window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)

	rlig.Load()
	defer rlig.Unload()
	imgui.StyleColorsDark()

	rlImgColorChecker := rl.NewImageFromImage(imgColorChecker)
	defer rl.UnloadImage(rlImgColorChecker)
	texColorChecker := rl.LoadTextureFromImage(rlImgColorChecker)
	defer rl.UnloadTexture(texColorChecker)

	camera := rl.NewCamera3D(rl.NewVector3(0, 10, 10), rl.NewVector3(0, 0, 0), rl.NewVector3(0, 1, 0), 45, rl.CameraPerspective)

	for !rl.WindowShouldClose() {
		rlig.Update()

		if autoRotate {
			rotation += 10 * rl.GetFrameTime()
			if rotation > 360 {
				rotation = 0
			}
		}

		imgui.NewFrame()
		if showDemoWindow {
			imgui.ShowDemoWindowV(&showDemoWindow)
		}

		imgui.SetNextWindowPosV(imgui.NewVec2(12, 220), imgui.CondFirstUseEver, imgui.NewVec2(0, 0))
		imgui.SetNextWindowSizeV(imgui.NewVec2(320, 320), imgui.CondFirstUseEver)
		imgui.Begin("Window Test")
		imgui.Text("Hello World")
		imgui.Checkbox("Show Demo Window", &showDemoWindow)
		imgui.ColorEdit3V("Background", &(backgroundColor), imgui.ColorEditFlagsNoOptions)
		imgui.ColorEdit3V("Cube", &(cubeColor), imgui.ColorEditFlagsNoOptions)
		imgui.ColorEdit3V("Wireframe", &(wireframeColor), imgui.ColorEditFlagsNoOptions)
		imgui.SliderFloatV("Rotation", &rotation, 0.0, 360, "%.3f", imgui.SliderFlagsNone)
		imgui.Checkbox("Auto Rotate", &autoRotate)
		imgui.InputTextWithHint("Dummy String", "", &dummyString, imgui.InputTextFlagsNone, nil)
		imgui.Image(imgui.TextureID(texColorChecker.ID), imgui.NewVec2(100, 100))
		imgui.End()

		rl.BeginDrawing()
		rl.ClearBackground(
			rl.NewColor(
				uint8(backgroundColor[0]*0xFF),
				uint8(backgroundColor[1]*0xFF),
				uint8(backgroundColor[2]*0xFF),
				0xFF,
			),
		)

		rl.BeginMode3D(camera)
		rl.PushMatrix()
		rl.Rotatef(rotation, 0, 1, 0)
		rl.DrawCube(
			rl.Vector3Zero(),
			2,
			2,
			2,
			rl.NewColor(
				uint8(cubeColor[0]*0xFF),
				uint8(cubeColor[1]*0xFF),
				uint8(cubeColor[2]*0xFF),
				0xFF,
			),
		)
		rl.DrawCubeWires(
			rl.Vector3Zero(),
			2,
			2,
			2,
			rl.NewColor(
				uint8(wireframeColor[0]*0xFF),
				uint8(wireframeColor[1]*0xFF),
				uint8(wireframeColor[2]*0xFF),
				0xFF,
			),
		)
		rl.PopMatrix()
		rl.DrawGrid(10, 1)
		rl.EndMode3D()

		rl.DrawText(dummyString, 10, 40, 20, rl.Red)
		rl.DrawFPS(10, 10)

		rlig.Render()
		rl.EndDrawing()
	}
}
