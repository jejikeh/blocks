package main

import (
	"fmt"
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

var Score int

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
		// rl.NewRectangle(BlockWidth*10, BlockHeight*6, BlockWidth, BlockHeight)),
	}
}

func (b *Block) Update() {
	b.Entity.Update()

	if b.State == Selected {
		b.Color = LerpColor(b.Color, rl.Color{255, 0, 0, 92}, 0.1)
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), b.Entity.GetCollisionRec()) {
		if b.State == Inactive {
			return
		}

		if b.State == Selected {
			Score--
		} else {
			Score++
		}

		b.State = Selected

		b.PlayAnimation(Squish)

		b.InitialColor = rl.Red

		b.Position = rl.Vector2Lerp(b.Position, rl.GetMousePosition(), 0.1)

		neighbors := GameInstance.Block.GetNearBlocks(b)

		// remove highlight before highlight
		for _, bl := range GameInstance.Block.Blocks {
			if bl.State == Highlight {
				bl.State = Inactive

				bl.Color = bl.InitialColor
			}
		}

		for _, neighbor := range neighbors {
			if neighbor.State == Selected || neighbor.State == Active {
				continue
			}

			neighbor.Color = rl.Green

			neighbor.State = Highlight
		}
	}
}

func (b *Blocks) GetNearBlocks(bl *Block) []*Block {
	blocks := make([]*Block, 0)

	blIndex := 0

	for i, block := range b.Blocks {
		if block == bl {
			blIndex = i
			continue
		}
	}

	possibleMoves := [][2]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

	for _, move := range possibleMoves {
		x := move[0]
		y := move[1]

		if blIndex%b.col+x < 0 || blIndex%b.col+x >= b.col || blIndex/b.col+y < 0 || blIndex/b.col+y >= b.row {
			continue
		}

		blocks = append(blocks, b.Blocks[(blIndex/b.col+y)*b.col+blIndex%b.col+x])
	}

	return blocks
}

type Blocks struct {
	Blocks []*Block

	row, col int
}

func NewBlocks(count int) *Blocks {
	b := &Blocks{
		Blocks: make([]*Block, count*count),
		row:    count,
		col:    count,
	}

	for i := 0; i < b.col; i++ {
		for j := 0; j < b.row; j++ {
			x := float32(i*BlockWidth+16*i) + float32(RenderInstance.RendererWidth/4) - float32(b.col*BlockWidth/2)

			y := float32(j*BlockHeight+16*j) + float32(RenderInstance.RendererHeight/4) - float32(b.row*BlockHeight/2)

			b.Blocks[i*b.col+j] = NewBlock(y, x)
		}
	}

	return b
}

func (b *Blocks) Init() {
	r := rand.Intn(len(b.Blocks))

	b.Blocks[r].State = Selected
	b.Blocks[r].Color = rl.Red
	b.Blocks[r].InitialColor = rl.Red

	neighbors := GameInstance.Block.GetNearBlocks(b.Blocks[r])

	// remove highlight before highlight
	for _, bl := range GameInstance.Block.Blocks {
		if bl.State == Highlight {
			bl.State = Inactive

			bl.Color = bl.InitialColor
		}
	}

	for _, neighbor := range neighbors {
		if neighbor.State == Selected || neighbor.State == Active {
			continue
		}

		neighbor.Color = rl.Green

		neighbor.State = Highlight
	}

	Score = 0
}

func (b *Blocks) Update() {
	for _, block := range b.Blocks {
		block.Update()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		for _, block := range b.Blocks {
			if block.State == Highlight {
				block.State = Selected
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		// b.blocks = make([]*Block, b.row*b.col)

		for i := 0; i < b.col; i++ {
			for j := 0; j < b.row; j++ {
				x := float32(i*BlockWidth+16*i) + float32(RenderInstance.RendererWidth/4) - float32(b.col*BlockWidth/2)

				y := float32(j*BlockHeight+16*j) + float32(RenderInstance.RendererHeight/4) - float32(b.row*BlockHeight/2)

				b.Blocks[i*b.col+j] = NewBlock(y, x)
			}
		}

		b.Init()
	}
}

func (b *Blocks) Draw() {
	for _, block := range b.Blocks {
		block.Draw()
	}

	rl.DrawText(fmt.Sprintf("Score: %d", Score), 10, 10, 20, rl.Red)
}
