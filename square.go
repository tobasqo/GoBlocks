package main

import (
	"log"
)

type SquareBlock struct {
	CommonBlock
	size int32
}

func NewSquareBlock(col, row, size int32) *SquareBlock {
	s := SquareBlock{
		CommonBlock{
			Position{col, row},
			RandomColor(),
			false,
		},
		size,
	}
	log.Printf("[DEBUG] created new square=%+v", s)
	return &s
}

func (s SquareBlock) Draw(b Board) {
	for xi := range s.size {
		x := (s.col+xi)*b.BlockSize() + b.StartX()
		for yi := range s.size {
			y := (s.row+yi)*b.BlockSize() + b.StartY()
			if s.selected || b.IsOccupied(s.col+xi, s.row+yi) {
				DrawBlock(x, y, s.color, b.BlockSize(), s.selected)
			}
		}
	}
}

func (s SquareBlock) DrawAsAvailable(ab AvailableBlocks, idx int32) {
	startX := ab.StartX() + idx*ab.Width()/3 + ab.Width()/6 - s.size*ab.BlockSize()/2.0 + idx
	startY := ab.StartY() + ab.Height()/2 - s.size*ab.BlockSize()/2.0
	for xi := range s.size {
		x := startX + xi*ab.BlockSize()
		for yi := range s.size {
			y := startY + yi*ab.BlockSize()
			DrawBlock(x, y, s.color, ab.BlockSize(), s.selected)
		}
	}
}

func (s SquareBlock) SetOccupied(b *Board) {
	for col := range s.size {
		for row := range s.size {
			log.Printf("[DEBUG] square(%+v).setOccupied(col=%d, row=%d)", s, s.col+col, s.row+row)
			b.SetOccupied(s.col+col, s.row+row)
		}
	}
}

func (s SquareBlock) IsOnOccupied(b Board) bool {
	for col := range s.size {
		for row := range s.size {
			if b.IsOccupied(s.col+col, s.row+row) {
				log.Printf("[DEBUG] square(%+v).isOnOccupied(col=%d, row=%d)", s, col, row)
				return true
			}
		}
	}
	return false
}

func (s *SquareBlock) MoveUp() {
	if s.row <= 0 {
		return
	}
	s.row--
}

func (s *SquareBlock) MoveDown(maxRow int32) {
	if s.row+s.size >= maxRow {
		return
	}
	s.row++
}

func (s *SquareBlock) MoveLeft() {
	if s.col <= 0 {
		return
	}
	s.col--
}

func (s *SquareBlock) MoveRight(maxCol int32) {
	if s.col+s.size >= maxCol {
		return
	}
	s.col++
}

func (s *SquareBlock) Rotate() {}

func (s *SquareBlock) Mirror() {}
