package main

type element struct {
	startX int32
	startY int32
	endX   int32
	endY   int32
}

func newPane(startX, startY, endX, endY int32) element {
	return element{
		startX,
		startY,
		endX,
		endY,
	}
}

func (p element) StartX() int32 {
	return p.startX
}

func (p element) StartY() int32 {
	return p.startY
}

func (p element) EndX() int32 {
	return p.endX
}

func (p element) EndY() int32 {
	return p.endY
}

func (p element) Width() int32 {
	return p.endX - p.startX
}

func (p element) Height() int32 {
	return p.endY - p.startY
}

type PaneWithBlocks struct {
	element
	blockSize int32
	blocks    []Block
}

func newPaneWithBlocks(startX, startY, endX, endY, blockSize int32, blocks []Block) PaneWithBlocks {
	return PaneWithBlocks{
		newPane(startX, startY, endX, endY),
		blockSize,
		blocks,
	}
}

func (p PaneWithBlocks) BlockSize() int32 {
	return p.blockSize
}
