package main

import (
	"log"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	PaneWithBlocks
	occupied [][]bool
	cols     int32
	rows     int32
}

func NewBoard(startX, startY, endX, endY, blockSize, cols, rows int32) *Board {
	occupied := make([][]bool, cols)
	for col := range occupied {
		occupied[col] = make([]bool, rows)
	}
	return &Board{
		newPaneWithBlocks(startX, startY, endX, endY, blockSize, []Block{}),
		occupied,
		cols,
		rows,
	}
}

func (b Board) drawGrid() {
	rl.DrawRectangle(b.startX, b.startY, b.Width(), b.Height(), rl.LightGray)
	for col := range b.cols + 1 {
		x := int32(col*b.blockSize) + b.startX
		rl.DrawLine(x, b.startY, x, b.endY, rl.Gray)
	}
	for row := range b.rows + 1 {
		y := int32(row*b.blockSize) + b.startY
		rl.DrawLine(b.startX, y, b.endX, y, rl.Gray)
	}
}

func (b Board) drawBlocks() {
	for _, block := range b.blocks {
		block.Draw(b)
	}
}

func (b Board) Draw() {
	rl.SetLineWidth(2)
	b.drawGrid()
	b.drawBlocks()
}

func (b *Board) AddBlock(block Block) {
	b.blocks = append(b.blocks, block)
}

func (b *Board) ResetOccupied() {
	for col := range b.cols {
		for row := range b.rows {
			b.occupied[col][row] = false
		}
	}
}

func (b Board) IsOccupied(col, row int32) bool {
	return b.occupied[col][row]
}

func (b *Board) SetOccupied(col, row int32) {
	b.occupied[col][row] = true
}

func (b Board) isColumnFull(col int32) bool {
	for row := range b.rows {
		if !b.IsOccupied(col, row) {
			return false
		}
	}
	log.Printf("[DEBUG] column %d is full", col)
	return true
}

func (b Board) getFullColumns() []int32 {
	fullCols := []int32{}
	for col := range b.cols {
		if b.isColumnFull(col) {
			fullCols = append(fullCols, col)
		}
	}
	return fullCols
}

func (b Board) isRowFull(row int32) bool {
	for col := range b.cols {
		if !b.IsOccupied(col, row) {
			return false
		}
	}
	log.Printf("[DEBUG] row %d is full", row)
	return true
}

func (b Board) getFullRows() []int32 {
	fullRows := []int32{}
	for row := range b.rows {
		if b.isRowFull(row) {
			fullRows = append(fullRows, row)
		}
	}
	return fullRows
}

func (b *Board) UnoccupyFull() {
	fullCols := b.getFullColumns()
	fullRows := b.getFullRows()
	for col := range b.cols {
		for row := range b.rows {
			if slices.Contains(fullCols, col) || slices.Contains(fullRows, row) {
				log.Printf("[DEBUG] unoccupying col=%d row=%d", col, row)
				b.occupied[col][row] = false
			}
		}
	}
}
