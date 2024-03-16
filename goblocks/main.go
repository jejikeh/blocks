package main

var TextureManagerInstance *TextureManager
var RenderInstance *Renderer

func main() {
	RenderInstance = newRenderer("Blocks", 1280, 720)

	TextureManagerInstance = NewTextureManager()

	g := newGame()

	RenderInstance.run(g)
}
