package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	Position rl.Vector2
	Rotation float32

	InitialSize rl.Vector2

	Size rl.Vector2

	Color rl.Color

	Texture       *rl.Texture2D
	SourceTexture rl.Rectangle

	Direction float32
	TimeMod   float64

	Animations map[int]*Animation
}

func NewEntity(t *rl.Texture2D, pos, size rl.Vector2, s rl.Rectangle) *Entity {
	en := &Entity{
		Position:      pos,
		InitialSize:   size,
		Size:          size,
		Texture:       t,
		SourceTexture: s,
		Color:         rl.White,
		Direction:     rand.Float32()*math.Pi + 2,
		TimeMod:       rand.Float64()*0.5 + 0.5,
		Animations:    make(map[int]*Animation),
	}

	en.AddAnimation(Squish)
	en.AddAnimation(Float)
	en.AddAnimation(LerpBackToInitialSize)

	return en
}

func (e *Entity) Update() {
	e.UpdateAnimations()

	if rl.IsKeyPressed(rl.KeySpace) {
		e.PlayAnimation(Squish)
	}
}

func (e *Entity) Draw() {
	rl.DrawTexturePro(*e.Texture, e.SourceTexture, rl.NewRectangle(e.Position.X, e.Position.Y, e.Size.X, e.Size.Y), rl.Vector2{}, e.Rotation, e.Color)
}
