package main

var TextureManagerInstance *TextureManager
var RenderInstance *Renderer
var GameInstance *Game

func main() {
	RenderInstance = newRenderer("Blocks", 1280, 720)

	TextureManagerInstance = NewTextureManager()

	GameInstance = newGame()

	RenderInstance.run(GameInstance)
}
