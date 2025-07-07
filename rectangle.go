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
	s := RectangleBlock{
		CommonBlock{
			Position{col, row},
			RandomColor(),
			false,
		},
		width,
		height,
	}
	log.Printf("[DEBUG] created new square=%+v", s)
	return &s
}

func (s RectangleBlock) Draw(b Board) {
	for xi := range s.width {
		x := (s.col+xi)*b.BlockSize() + b.StartX()
		for yi := range s.height {
			y := (s.row+yi)*b.BlockSize() + b.StartY()
			if s.selected || b.IsOccupied(s.col+xi, s.row+yi) {
				DrawBlock(x, y, s.color, b.BlockSize(), s.selected)
			}
		}
	}
}

func (s RectangleBlock) DrawAsAvailable(ab AvailableBlocks, idx int32) {
	startX := ab.StartX() + idx*ab.Width()/3 + ab.Width()/6 - s.width*ab.BlockSize()/2.0 + idx
	startY := ab.StartY() + ab.Height()/2 - s.height*ab.BlockSize()/2.0
	for xi := range s.width {
		x := startX + xi*ab.BlockSize()
		for yi := range s.height {
			y := startY + yi*ab.BlockSize()
			DrawBlock(x, y, s.color, ab.BlockSize(), s.selected)
		}
	}
}

func (s RectangleBlock) SetOccupied(b *Board) {
	for col := range s.width {
		for row := range s.height {
			log.Printf("[DEBUG] square(%+v).setOccupied(col=%d, row=%d)", s, s.col+col, s.row+row)
			b.SetOccupied(s.col+col, s.row+row)
		}
	}
}

func (s RectangleBlock) IsOnOccupied(b Board) bool {
	for col := range s.width {
		for row := range s.height {
			if b.IsOccupied(s.col+col, s.row+row) {
				log.Printf("[DEBUG] square(%+v).isOnOccupied(col=%d, row=%d)", s, col, row)
				return true
			}
		}
	}
	return false
}

func (s *RectangleBlock) MoveUp() {
	if s.row <= 0 {
		return
	}
	s.row--
}

func (s *RectangleBlock) MoveDown(maxRow int32) {
	if s.row+s.height >= maxRow {
		return
	}
	s.row++
}

func (s *RectangleBlock) MoveLeft() {
	if s.col <= 0 {
		return
	}
	s.col--
}

func (s *RectangleBlock) MoveRight(maxCol int32) {
	if s.col+s.width >= maxCol {
		return
	}
	s.col++
}

func (s *RectangleBlock) Rotate(b Board) {
	tmp := s.width
	s.width = s.height
	s.height = tmp
	if s.col+s.width >= b.cols {
		s.col = b.cols - s.width
	}
	if s.row+s.height >= b.rows {
		s.row = b.rows - s.height
	}
}

func (s *RectangleBlock) Mirror(b Board) {}
