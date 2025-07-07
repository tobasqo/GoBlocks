package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Block interface {
	Draw(b Board)
	DrawAsAvailable(ab AvailableBlocks, idx int32)
	SetOccupied(b *Board)
	IsOnOccupied(b Board) bool
	MoveUp()
	MoveDown(maxRow int32)
	MoveLeft()
	MoveRight(maxCol int32)
	Rotate(b Board)
	Mirror(b Board)
	Select()
	Deselect()
	Selected() bool
}

type Position struct {
	col, row int32
}

type CommonBlock struct {
	Position
	color    rl.Color
	selected bool
}

func (b *CommonBlock) Select() {
	b.selected = true
}

func (b *CommonBlock) Deselect() {
	b.selected = false
}

func (b CommonBlock) Selected() bool {
	return b.selected
}
