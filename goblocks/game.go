package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	block *Blocks
}

func newGame() *Game {
	return &Game{
		block: NewBlocks(5),
	}
}

func (g *Game) init() {
}

func (g *Game) update() {
	g.block.Update()
}

func (g *Game) draw() {
	rl.ClearBackground(rl.Black)

	g.block.Draw()
}
