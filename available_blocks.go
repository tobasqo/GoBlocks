package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AvailableBlocks struct {
	PaneWithBlocks
	curBlockIdx int32
}

func NewAvailableBlocks(startX, startY, endX, endY, blockSize, cols, rows int32) *AvailableBlocks {
	blocks := []Block{}
	for range 3 {
		blocks = append(blocks, RandomBlock(cols, rows))
	}
	curBlockIdx := 1
	blocks[curBlockIdx].Select()
	return &AvailableBlocks{
		newPaneWithBlocks(startX, startY, endX, endY, blockSize, blocks),
		int32(curBlockIdx),
	}
}

func (ab AvailableBlocks) drawPane() {
	rl.DrawRectangle(ab.startX, ab.startY, ab.Width(), ab.Height(), rl.Gray)
	sepStartX := ab.startX + ab.Width()/3 + 1
	sepStartY := ab.startY
	sepEndY := sepStartY + ab.Height()
	rl.DrawLine(sepStartX, sepStartY, sepStartX, sepEndY, rl.DarkGray)
	sepStartX = sepStartX + ab.Width()/3 + 1
	rl.DrawLine(sepStartX, sepStartY, sepStartX, sepEndY, rl.DarkGray)
}

func (ab AvailableBlocks) drawBlocks() {
	rl.SetLineWidth(1)
	for idx, block := range ab.blocks {
		block.DrawAsAvailable(ab, int32(idx))
	}
}

func (ab AvailableBlocks) Draw() {
	rl.SetLineWidth(1)
	ab.drawPane()
	ab.drawBlocks()
}

func (ab AvailableBlocks) SelectedBlock() Block {
	return ab.blocks[ab.curBlockIdx]
}

func (ab *AvailableBlocks) SelectPrev() {
	if ab.curBlockIdx > 0 {
		ab.blocks[ab.curBlockIdx].Deselect()
		ab.curBlockIdx--
		ab.blocks[ab.curBlockIdx].Select()
	}
}

func (ab *AvailableBlocks) SelectNext() {
	if ab.curBlockIdx < 2 {
		ab.blocks[ab.curBlockIdx].Deselect()
		ab.curBlockIdx++
		ab.blocks[ab.curBlockIdx].Select()
	}
}

func (ab *AvailableBlocks) PopSelected(cols, rows int32) Block {
	block := ab.SelectedBlock()
	block.Deselect()
	ab.blocks[ab.curBlockIdx] = RandomBlock(cols, rows)
	ab.blocks[ab.curBlockIdx].Select()
	return block
}
