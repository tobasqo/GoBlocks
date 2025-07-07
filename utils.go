package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RandomColor() rl.Color {
	colors := []rl.Color{rl.Red, rl.Orange, rl.Yellow, rl.Green, rl.Blue, rl.Pink}
	colorIdx := rand.Int() % len(colors)
	return colors[colorIdx]
}

func DrawBlock(posX, posY int32, color rl.Color, blockSize int32, selected bool) {
	rl.DrawRectangle(posX, posY, blockSize, blockSize, color)
	borderColor := rl.Black
	if selected {
		borderColor = rl.White
	}
	rl.DrawRectangleLines(posX, posY, blockSize, blockSize, borderColor)
}

func RandomBlock(cols, rows int32) Block {
	const nBlockTypes = 2
	blockIdx := rand.Int() % nBlockTypes
	switch blockIdx {
	case 0:
		size := int32(rand.Int()%3 + 1)
		return NewSquareBlock((cols/2)-(size/2), (rows/2)-(size/2), size)
	case 1:
		sizes := []int32{2, 3, 4, 5}
		sizeIdx := rand.Int() % len(sizes)
		size := sizes[sizeIdx]
		return NewLineBlock(cols/2, (rows/2)-(size/2), size)
	}
	panic("rnd block idx out of range")
}
