# raylib-cimgui-go-starterkit

Minimal starterkit for https://github.com/gen2brain/raylib-go/ and https://github.com/AllenDang/cimgui-go .

![Preview](https://github.com/user-attachments/assets/f55a9c37-2662-40c3-936d-130e925a41c2)

## Integrate Existing project

Copy `pkg/raylib_imgui/main.go`. minimal usage:

```go
func main() {
	// ..

	// NOTE: load imgui context
	rlig.Load()
	defer rlig.Unload()
	imgui.StyleColorsDark()

	// ..

	for !rl.WindowShouldClose() {
		// NOTE: update imgui input
		rlig.Update()

		// NOTE: create new framebuffer, need to start add imgui widget
		imgui.NewFrame()

		imgui.ShowDemoWindow()

		rl.BeginDrawing()
		rl.ClearBackground(
			rl.NewColor(
				0,
				0,
				0,
				0xFF,
			),
		)
		// NOTE: render imgui, need to call inside drawing scope
		rlig.Render()
		rl.EndDrawing()
	}
}
}
```
