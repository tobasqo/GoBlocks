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
	const nBlockTypes = 3
	blockIdx := rand.Int() % nBlockTypes
	switch blockIdx {
	case 0:
		oneByOne := rand.Int()%6 == 1
		var width int32
		var height int32
		if oneByOne {
			width = 1
			height = 1
		} else {
			width = int32(rand.Int()%2 + 2)
			height = int32(rand.Int()%2 + 2)
		}
		return NewRectangleBlock((cols/2)-(width/2), (rows/2)-(height/2), width, height)
	case 1:
		sizes := []int32{2, 3, 4, 5}
		sizeIdx := rand.Int() % len(sizes)
		size := sizes[sizeIdx]
		horizontal := rand.Int() % 2 == 0
		var orientation LineOrientation
		if horizontal {
			orientation = LINE_HORIZONTAL
		} else {
			orientation = LINE_VERTICAL
		}
		return NewLineBlock(cols/2, (rows/2)-(size/2), size, orientation)
	case 2:
		orientationIdx := rand.Int() % 4
		var orientation PlusOrientation
		switch orientationIdx {
		case 0:
			orientation = PLUS_UP
		case 1:
			orientation = PLUS_LEFT
		case 2:
			orientation = PLUS_DOWN
		case 3:
			orientation = PLUS_RIGHT
		}
		return NewPlusBlock(cols/2, rows/2, orientation)
	}
	panic("rnd block idx out of range")
}
