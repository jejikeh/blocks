package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Block *Blocks
}

func newGame() *Game {
	return &Game{
		Block: NewBlocks(10),
	}
}

func (g *Game) init() {
	g.Block.Init()
}

func (g *Game) update() {
	g.Block.Update()
}

func (g *Game) draw() {
	rl.ClearBackground(rl.Black)

	g.Block.Draw()
}
