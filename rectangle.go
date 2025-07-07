package main

import (
	"log"
)

type RectangleBlock struct {
	CommonBlock
	width int32
	height int32
}

func NewRectangleBlock(col, row, width, height int32) *RectangleBlock {
	r := RectangleBlock{
		newCommonBlock(col, row),
		width,
		height,
	}
	log.Printf("[DEBUG] created new rectangle=%+v", r)
	return &r
}

func (r RectangleBlock) Draw(b Board) {
	for xi := range r.width {
		x := (r.col+xi)*b.BlockSize() + b.StartX()
		for yi := range r.height {
			y := (r.row+yi)*b.BlockSize() + b.StartY()
			if r.selected || b.IsOccupied(r.col+xi, r.row+yi) {
				DrawBlock(x, y, r.color, b.BlockSize(), r.selected)
			}
		}
	}
}

func (r RectangleBlock) DrawAsAvailable(ab AvailableBlocks, idx int32) {
	startX := ab.StartX() + idx*ab.Width()/3 + ab.Width()/6 - r.width*ab.BlockSize()/2.0 + idx
	startY := ab.StartY() + ab.Height()/2 - r.height*ab.BlockSize()/2.0
	for xi := range r.width {
		x := startX + xi*ab.BlockSize()
		for yi := range r.height {
			y := startY + yi*ab.BlockSize()
			DrawBlock(x, y, r.color, ab.BlockSize(), r.selected)
		}
	}
}

func (r RectangleBlock) SetOccupied(b *Board) {
	for col := range r.width {
		for row := range r.height {
			log.Printf("[DEBUG] rectangle(%+v).setOccupied(col=%d, row=%d)", r, r.col+col, r.row+row)
			b.SetOccupied(r.col+col, r.row+row)
		}
	}
}

func (r RectangleBlock) IsOnOccupied(b Board) bool {
	for col := range r.width {
		for row := range r.height {
			if b.IsOccupied(r.col+col, r.row+row) {
				log.Printf("[DEBUG] rectangle(%+v).isOnOccupied(col=%d, row=%d)", r, col, row)
				return true
			}
		}
	}
	return false
}

func (r *RectangleBlock) MoveUp() {
	if r.row <= 0 {
		return
	}
	r.row--
}

func (r *RectangleBlock) MoveDown(maxRow int32) {
	if r.row+r.height >= maxRow {
		return
	}
	r.row++
}

func (r *RectangleBlock) MoveLeft() {
	if r.col <= 0 {
		return
	}
	r.col--
}

func (r *RectangleBlock) MoveRight(maxCol int32) {
	if r.col+r.width >= maxCol {
		return
	}
	r.col++
}

func (r *RectangleBlock) Rotate(b Board) {
	tmp := r.width
	r.width = r.height
	r.height = tmp
	if r.col+r.width >= b.cols {
		r.col = b.cols - r.width
	}
	if r.row+r.height >= b.rows {
		r.row = b.rows - r.height
	}
}

func (r *RectangleBlock) Mirror(b Board) {}
