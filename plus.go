package main

import "log"

type PlusOrientation string

const (
	PLUS_UP    PlusOrientation = "up"
	PLUS_RIGHT PlusOrientation = "right"
	PLUS_DOWN  PlusOrientation = "down"
	PLUS_LEFT  PlusOrientation = "left"
)

type PlusBlock struct {
	CommonBlock
	orientation PlusOrientation
}

func NewPlusBlock(col, row int32, orientation PlusOrientation) *PlusBlock {
	p := PlusBlock{
		newCommonBlock(col, row),
		orientation,
	}
	log.Printf("[DEBUG] created new plus=%+v", p)
	return &p
}

func (p PlusBlock) Draw(b Board) {
	col := p.col
	row := p.row
	color := p.color
	blockSize := b.BlockSize()
	selected := p.selected
	startX := b.StartX()
	startY := b.StartY()
	switch p.orientation {
	case PLUS_UP:
		x := col*blockSize + startX
		y := row*blockSize + startY
		DrawBlock(x-blockSize, y, color, blockSize, selected)
		DrawBlock(x, y-blockSize, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x+blockSize, y, color, blockSize, selected)
	case PLUS_RIGHT:
		x := col*blockSize + startX
		y := row*blockSize + startY
		DrawBlock(x, y-blockSize, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x+blockSize, y, color, blockSize, selected)
		DrawBlock(x, y+blockSize, color, blockSize, selected)
	case PLUS_DOWN:
		x := col*blockSize + startX
		y := row*blockSize + startY
		DrawBlock(x-blockSize, y, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x, y+blockSize, color, blockSize, selected)
		DrawBlock(x+blockSize, y, color, blockSize, selected)
	case PLUS_LEFT:
		x := col*blockSize + startX
		y := row*blockSize + startY
		DrawBlock(x, y-blockSize, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x-blockSize, y, color, blockSize, selected)
		DrawBlock(x, y+blockSize, color, blockSize, selected)
	default:
		panic("unreachable")
	}
}

func (p PlusBlock) DrawAsAvailable(ab AvailableBlocks, idx int32) {
	color := p.color
	blockSize := ab.BlockSize()
	selected := p.selected
	startX := ab.StartX()
	startY := ab.StartY()
	switch p.orientation {
	case PLUS_UP:
		x := startX + idx*ab.Width()/3.0 + ab.Width()/6.0 - blockSize/2.0 + idx
		y := startY + ab.Height()/2
		DrawBlock(x-blockSize, y, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x, y-blockSize, color, blockSize, selected)
		DrawBlock(x+blockSize, y, color, blockSize, selected)
	case PLUS_RIGHT:
		x := startX + idx*ab.Width()/3.0 + ab.Width()/6.0 - blockSize + idx
		y := startY + ab.Height()/2 - blockSize/2.0
		DrawBlock(x, y-blockSize, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x+blockSize, y, color, blockSize, selected)
		DrawBlock(x, y+blockSize, color, blockSize, selected)
	case PLUS_DOWN:
		x := startX + idx*ab.Width()/3.0 + ab.Width()/6.0 - blockSize/2.0 + idx
		y := startY + ab.Height()/2 - blockSize
		DrawBlock(x-blockSize, y, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x, y+blockSize, color, blockSize, selected)
		DrawBlock(x+blockSize, y, color, blockSize, selected)
	case PLUS_LEFT:
		x := startX + idx*ab.Width()/3.0 + ab.Width()/6.0 + idx
		y := startY + ab.Height()/2 - blockSize/2.0
		DrawBlock(x, y-blockSize, color, blockSize, selected)
		DrawBlock(x, y, color, blockSize, selected)
		DrawBlock(x-blockSize, y, color, blockSize, selected)
		DrawBlock(x, y+blockSize, color, blockSize, selected)
	default:
		panic("unreachable")
	}
}

func (p PlusBlock) SetOccupied(b *Board) {
	col := p.col
	row := p.row
	switch p.orientation {
	case PLUS_UP:
		b.SetOccupied(col-1, row)
		b.SetOccupied(col, row)
		b.SetOccupied(col, row-1)
		b.SetOccupied(col+1, row)
	case PLUS_RIGHT:
		b.SetOccupied(col, row-1)
		b.SetOccupied(col, row)
		b.SetOccupied(col+1, row)
		b.SetOccupied(col, row+1)
	case PLUS_DOWN:
		b.SetOccupied(col-1, row)
		b.SetOccupied(col, row)
		b.SetOccupied(col, row+1)
		b.SetOccupied(col+1, row)
	case PLUS_LEFT:
		b.SetOccupied(col, row-1)
		b.SetOccupied(col, row)
		b.SetOccupied(col-1, row)
		b.SetOccupied(col, row+1)
	default:
		panic("unreachable")
	}
}

func (p PlusBlock) IsOnOccupied(b Board) bool {
	col := p.col
	row := p.row
	switch p.orientation {
	case PLUS_UP:
		return b.IsOccupied(col-1, row) || b.IsOccupied(col, row) || b.IsOccupied(col, row-1) || b.IsOccupied(col+1, row)
	case PLUS_RIGHT:
		return b.IsOccupied(col, row-1) || b.IsOccupied(col, row) || b.IsOccupied(col+1, row) || b.IsOccupied(col, row+1)
	case PLUS_DOWN:
		return b.IsOccupied(col-1, row) || b.IsOccupied(col, row) || b.IsOccupied(col, row+1) || b.IsOccupied(col+1, row)
	case PLUS_LEFT:
		return b.IsOccupied(col, row-1) || b.IsOccupied(col, row) || b.IsOccupied(col-1, row) || b.IsOccupied(col, row+1)
	}
	panic("unreachable")
}

func (p *PlusBlock) MoveUp() {
	switch p.orientation {
	case PLUS_UP:
		fallthrough
	case PLUS_RIGHT:
		fallthrough
	case PLUS_LEFT:
		if p.row - 1 > 0 {
			p.row--
		}
	case PLUS_DOWN:
		if p.row > 0 {
			p.row--
		}
	default:
		panic("unreachable")
	}
}

func (p *PlusBlock) MoveDown(maxRow int32) {
	switch p.orientation {
	case PLUS_RIGHT:
		fallthrough
	case PLUS_LEFT:
		fallthrough
	case PLUS_DOWN:
		if p.row + 2 < maxRow {
			p.row++
		}
	case PLUS_UP:
		if p.row + 1 < maxRow {
			p.row++
		}
	default:
		panic("unreachable")
	}
}

func (p *PlusBlock) MoveLeft() {
	switch p.orientation {
	case PLUS_UP:
		fallthrough
	case PLUS_LEFT:
		fallthrough
	case PLUS_DOWN:
		if p.col - 1 > 0 {
			p.col--
		}
	case PLUS_RIGHT:
		if p.col > 0 {
			p.col--
		}
	default:
		panic("unreachable")
	}
}

func (p *PlusBlock) MoveRight(maxCol int32) {
	switch p.orientation {
	case PLUS_UP:
		fallthrough
	case PLUS_RIGHT:
		fallthrough
	case PLUS_DOWN:
		if p.col < maxCol - 2 {
			p.col++
		}
	case PLUS_LEFT:
		if p.col < maxCol - 1 {
			p.col++
		}
	default:
		panic("unreachable")
	}
}

func (p *PlusBlock) Rotate(b Board) {
	switch p.orientation {
	case PLUS_UP:
		p.orientation = PLUS_RIGHT
		if p.row + 1 >= b.rows {
			p.row--
		}
	case PLUS_RIGHT:
		p.orientation = PLUS_DOWN
		if p.col == 0 {
			p.col++
		} else if p.col + 1 >= b.cols {
			p.col--
		}
	case PLUS_DOWN:
		p.orientation = PLUS_LEFT
		if p.row == 0 {
			p.row++
		}
	case PLUS_LEFT:
		p.orientation = PLUS_UP
		if p.col + 1 == b.cols {
			p.col--
		}
	default:
		panic("unreachable")
	}
	log.Printf("[DEBUG] orientation=%s", p.orientation)
}

func (p *PlusBlock) Mirror(b Board) {
	switch p.orientation {
	case PLUS_UP:
		p.orientation = PLUS_DOWN
		if p.row + 1 >= b.rows {
			p.row--
		} 
	case PLUS_RIGHT:
		p.orientation = PLUS_LEFT
		if p.col == 0 {
			p.col++
		}
	case PLUS_DOWN:
		p.orientation = PLUS_UP
		if p.row == 0 {
			p.row++
		}
	case PLUS_LEFT:
		p.orientation = PLUS_RIGHT
		if p.col >= b.cols - 1 {
			p.col--
		}
	default:
		panic("unreachable")
	}
}
