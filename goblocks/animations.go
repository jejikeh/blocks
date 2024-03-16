package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Squish = iota
	Float
	LerpBackToInitialSize
)

type Animation struct {
	CurrentFrame int
	TotalFrames  int
	Looping      bool

	Tag int

	FrameTime          float32
	TimeSinceLastFrame float32

	IsPlaying bool

	Update func(a *Animation, en *Entity)
}

func NewFloatAnimation() *Animation {
	return &Animation{
		TotalFrames: 20,
		Looping:     true,
		Tag:         Float,
		FrameTime:   0.05,
		IsPlaying:   true,
		Update: func(a *Animation, e *Entity) {
			e.Position.X += float32(math.Sin(float64(rl.GetTime()+e.TimeMod))) * 0.5 / e.Direction
			e.Position.Y += float32(math.Cos(float64(rl.GetTime()+e.TimeMod))) * 0.5 / e.Direction

			e.Size.X += float32(math.Sin(float64(rl.GetTime()+e.TimeMod))) * 0.1 / e.Direction
			e.Size.Y += float32(math.Cos(float64(rl.GetTime()+e.TimeMod))) * 0.1 / e.Direction
		},
	}
}

func NewSquishAnimation() *Animation {
	return &Animation{
		TotalFrames: 6,
		Looping:     false,
		Tag:         Squish,
		FrameTime:   0.03,
		IsPlaying:   false,
		Update: func(a *Animation, e *Entity) {
			e.Size.X += float32(math.Sin(float64(rl.GetTime()+rand.Float64()))) * 3.0
			e.Size.Y += float32(math.Cos(float64(rl.GetTime()+rand.Float64()))) * 3.0

			// last frame
			if a.CurrentFrame == a.TotalFrames-1 {
				e.PlayAnimation(LerpBackToInitialSize)
			}
		},
	}
}

func NewLerpBackToInitialSizeAnimation() *Animation {
	return &Animation{
		TotalFrames: 90,
		Looping:     false,
		Tag:         LerpBackToInitialSize,
		FrameTime:   0.03,
		IsPlaying:   false,
		Update: func(a *Animation, e *Entity) {
			e.Size = rl.Vector2Lerp(e.Size, e.InitialSize, 0.1)

			// last frame
			if a.CurrentFrame == a.TotalFrames-1 {
				e.Size = e.InitialSize
			}
		},
	}
}

func NewAnimation(tag int) *Animation {
	switch tag {
	case Squish:
		return NewSquishAnimation()
	case Float:
		return NewFloatAnimation()
	case LerpBackToInitialSize:
		return NewLerpBackToInitialSizeAnimation()
	}

	return nil
}

func (e *Entity) UpdateAnimations() {
	for _, a := range e.Animations {
		if a.IsPlaying {
			a.TimeSinceLastFrame += float32(rl.GetFrameTime())

			if a.TimeSinceLastFrame >= a.FrameTime {
				a.TimeSinceLastFrame = 0

				a.CurrentFrame = (a.CurrentFrame + 1) % a.TotalFrames

				a.Update(a, e)
			} else {
				if !a.Looping && a.CurrentFrame == a.TotalFrames-1 {
					a.IsPlaying = false
					a.CurrentFrame = 0
					a.TimeSinceLastFrame = 0
				}
			}
		}
	}
}

func (e *Entity) AddAnimation(tag int) {
	e.Animations[tag] = NewAnimation(tag)
}

func (e *Entity) PlayAnimation(tag int) {
	e.Animations[tag].IsPlaying = true
}

func (e *Entity) RemoveAnimation(tag int) {
	delete(e.Animations, tag)
}
