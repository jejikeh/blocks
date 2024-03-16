package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Renderer struct {
	Title     string
	TargetFPS int32

	WindowWidth  int32
	WindowHeight int32

	RendererWidth  int32
	RendererHeight int32

	Textrure rl.RenderTexture2D
}

var MouseScale rl.Vector2

func newRenderer(title string, x, y int32) *Renderer {
	r := &Renderer{
		Title:          title,
		TargetFPS:      60,
		WindowWidth:    x,
		WindowHeight:   y,
		RendererWidth:  640,
		RendererHeight: 480,
	}

	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(r.WindowWidth, r.WindowHeight, r.Title)

	rl.SetTargetFPS(r.TargetFPS)

	r.Textrure = rl.LoadRenderTexture(int32(r.RendererWidth), int32(r.RendererHeight))
	rl.SetTextureFilter(r.Textrure.Texture, rl.TextureFilterLinear)

	return r
}

func (r *Renderer) run(g *Game) {
	calculateDestinationRectangle := func() rl.Rectangle {
		screenHeight := float32(rl.GetScreenHeight())
		screenWidth := float32(rl.GetScreenWidth())

		// if rl.IsWindowFullscreen() {
		// 	monitor := rl.GetCurrentMonitor()
		// 	screenHeight = float32(rl.GetMonitorHeight(monitor))
		// 	screenWidth = float32(rl.GetMonitorWidth(monitor))
		// }

		scale := screenHeight / float32(r.RendererHeight)

		x0 := (screenWidth - float32(r.RendererWidth)*scale) / 2
		var y0 float32 = 0.0

		x1 := float32(r.RendererWidth) * scale
		y1 := screenHeight

		rl.SetMouseOffset(-int(x0), -int(y0))
		rl.SetMouseScale(1/scale, 1/scale)
		MouseScale = rl.NewVector2(scale, scale)

		return rl.NewRectangle(x0, y0, x1, y1)
	}

	g.init()

	for !rl.WindowShouldClose() {
		g.update()

		rl.BeginTextureMode(r.Textrure)
		rl.ClearBackground(rl.Black)

		g.draw()

		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTexturePro(
			r.Textrure.Texture,
			rl.NewRectangle(0, 0, float32(r.Textrure.Texture.Width), -float32(r.Textrure.Texture.Height)),
			calculateDestinationRectangle(),
			rl.NewVector2(0, 0),
			0,
			rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
