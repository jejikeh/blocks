package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Inactive = iota
	Highlight
	Selected
	Active
)

const BlockWidth = 10
const BlockHeight = 10

type Block struct {
	Entity

	State int
}

func NewBlock(x, y float32) *Block {
	clothing, ok := TextureManagerInstance.Textures[Clothing]
	if !ok {
		return nil
	}

	return &Block{
		Entity: *NewEntity(
			&clothing,
			rl.NewVector2(x, y),
			rl.NewVector2(BlockWidth*2, BlockHeight*2),
			rl.NewRectangle(BlockWidth*float32(rand.Int31n(64)), BlockHeight*float32(rand.Int31n(64)), BlockWidth, BlockHeight)),
	}
}

func (b *Block) Update() {
	b.Entity.Update()
}

type Blocks struct {
	blocks []*Block

	row, col int
}

func NewBlocks(count int) *Blocks {
	b := &Blocks{
		blocks: make([]*Block, count*count),
		row:    count,
		col:    count,
	}

	for i := 0; i < b.row; i++ {
		for j := 0; j < b.col; j++ {
			x := float32(i*BlockWidth+16*i) + float32(RenderInstance.RendererWidth/2) - float32(b.col*BlockWidth/2)
			y := float32(j*BlockHeight+16*j) + float32(RenderInstance.RendererHeight/2) - float32(b.row*BlockHeight/2)

			b.blocks[i*b.col+j] = NewBlock(x, y)
		}
	}

	return b
}

func (b *Blocks) Update() {
	for _, block := range b.blocks {
		block.Update()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		r := rand.Intn(len(b.blocks))

		b.blocks[r].State = Selected
	}
}

func (b *Blocks) Draw() {
	for _, block := range b.blocks {
		block.Draw()
	}
}
