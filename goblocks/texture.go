package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Texture string

const (
	Clothing Texture = "assets/clothing.png"
)

type TextureManager struct {
	Textures map[Texture]rl.Texture2D
}

func NewTextureManager() *TextureManager {
	tm := &TextureManager{
		Textures: make(map[Texture]rl.Texture2D),
	}

	tm.LoadTexture(Clothing)

	return tm
}

func (tm *TextureManager) LoadTexture(texture Texture) rl.Texture2D {
	tm.Textures[texture] = rl.LoadTexture(string(texture))

	return tm.Textures[texture]
}
