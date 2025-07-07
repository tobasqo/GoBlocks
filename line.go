package main

import (
	"log"
)

type LineOrientation string

const (
	LINE_HORIZONTAL LineOrientation = "horizontal"
	LINE_VERTICAL   LineOrientation = "vertical"
)

type LineBlock struct {
	CommonBlock
	size        int32
	orientation LineOrientation
}

func NewLineBlock(col, row, size int32, orientation LineOrientation) *LineBlock {
	l := LineBlock{
		newCommonBlock(col, row),
		size,
		orientation,
	}
	log.Printf("[DEBUG] created new line=%+v", l)
	return &l
}

func (l LineBlock) Draw(b Board) {
	if l.orientation == LINE_VERTICAL {
		for yi := range l.size {
			x := l.col*b.BlockSize() + b.StartX()
			y := (l.row+yi)*b.BlockSize() + b.StartY()
			if l.selected || b.IsOccupied(l.col, l.row+yi) {
				DrawBlock(x, y, l.color, b.BlockSize(), l.selected)
			}
		}
	} else {
		for xi := range l.size {
			x := (l.col+xi)*b.BlockSize() + b.StartX()
			y := l.row*b.BlockSize() + b.StartY()
			if l.selected || b.IsOccupied(l.col+xi, l.row) {
				DrawBlock(x, y, l.color, b.BlockSize(), l.selected)
			}
		}
	}
}

func (l LineBlock) DrawAsAvailable(ab AvailableBlocks, idx int32) {
	if l.orientation == LINE_HORIZONTAL {
		startX := ab.StartX() + idx*ab.Width()/3 + ab.Width()/6 - l.size*ab.BlockSize()/2.0 + idx
		startY := ab.StartY() + ab.Height()/2 - ab.BlockSize()/2.0
		for xi := range l.size {
			x := startX + xi*ab.BlockSize()
			y := startY
			DrawBlock(x, y, l.color, ab.BlockSize(), l.selected)
		}
	} else {
		startX := ab.StartX() + idx*ab.Width()/3 + ab.Width()/6 - ab.BlockSize()/2 + idx
		startY := ab.StartY() + ab.Height()/2 - l.size*ab.BlockSize()/2.0
		for yi := range l.size {
			x := startX
			y := startY + yi*ab.BlockSize()
			DrawBlock(x, y, l.color, ab.BlockSize(), l.selected)
		}
	}
}

func (l LineBlock) SetOccupied(b *Board) {
	if l.orientation == LINE_VERTICAL {
		for row := range l.size {
			log.Printf("[DEBUG] line(%+v).setOccupied(col=%d, row=%d)", l, l.col, l.row+row)
			b.SetOccupied(l.col, l.row+row)
		}
	} else {
		for col := range l.size {
			log.Printf("[DEBUG] line(%+v).setOccupied(col=%d, row=%d)", l, l.col+col, l.row)
			b.SetOccupied(l.col+col, l.row)
		}
	}
}

func (l LineBlock) IsOnOccupied(b Board) bool {
	if l.orientation == LINE_VERTICAL {
		for row := range l.size {
			if b.IsOccupied(l.col, l.row+row) {
				log.Printf("[DEBUG] line(%+v).isOnOccupied(col=%d, row=%d)", l, l.col, l.row+row)
				return true
			}
		}
	} else {
		for col := range l.size {
			if b.IsOccupied(l.col+col, l.row) {
				log.Printf("[DEBUG] line(%+v).isOnOccupied(col=%d, row=%d)", l, l.col+col, l.row)
				return true
			}
		}
	}
	return false
}

func (l *LineBlock) MoveUp() {
	if l.row <= 0 {
		return
	}
	l.row--
}

func (l *LineBlock) MoveDown(maxRow int32) {
	if (l.orientation == LINE_VERTICAL && l.row+l.size < maxRow) || (l.orientation == LINE_HORIZONTAL && l.row < maxRow-1) {
		l.row++
	}
}

func (l *LineBlock) MoveLeft() {
	if l.col > 0 {
		l.col--
	}
}

func (l *LineBlock) MoveRight(maxCol int32) {
	if (l.orientation == LINE_VERTICAL && l.col < maxCol-1) || (l.orientation == LINE_HORIZONTAL && l.col+l.size < maxCol) {
		l.col++
	}
}

func (l *LineBlock) Rotate(b Board) {
	if l.orientation == LINE_VERTICAL {
		l.orientation = LINE_HORIZONTAL
		if l.col + l.size >= b.cols {
			l.col = b.cols - l.size
		}
	} else {
		l.orientation = LINE_VERTICAL
		if l.row + l.size >= b.rows {
			l.row = b.rows - l.size
		}
	}
}

func (l *LineBlock) Mirror(b Board) {}
