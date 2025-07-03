package main

import (
	"log"
	"math/rand"
	"os"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
* TODO:
* - more depth on block grid
* - removing full row/col
* - display for available blocks
* - selecting prev/next block
* - rotating block
* - mirroring block
* - plus block
* - L block
* - zigzag block
* - fix type size
 */

const (
	TITLE                 = "GoBlocks"
	W_WIDTH       int32   = 960
	W_HEIGHT      int32   = 820
	FPS           int32   = 60
	TOP_MARGIN    int32   = 20
	LEFT_MARGIN   int32   = 30
	RIGHT_MARGIN  int32   = LEFT_MARGIN
	BOTTOM_MARGIN int32   = 200
	G_WIDTH       int32   = W_WIDTH - LEFT_MARGIN - RIGHT_MARGIN
	G_HEIGHT      int32   = W_HEIGHT - TOP_MARGIN - BOTTOM_MARGIN
	BLOCK_SIZE    int32   = 60
	COLS          int32   = G_WIDTH / BLOCK_SIZE
	ROWS          int32   = G_HEIGHT / BLOCK_SIZE
	STROKE        float32 = 2
)

var logger *log.Logger = log.New(os.Stdout, "goblocks:", log.LstdFlags)

func randomColor() rl.Color {
	colors := []rl.Color{rl.Red, rl.Orange, rl.Yellow, rl.Green, rl.Blue, rl.Pink}
	colorIdx := rand.Int() % len(colors)
	return colors[colorIdx]
}

func drawBlock(posX, posY int32, color rl.Color) {
	rl.DrawRectangle(posX, posY, BLOCK_SIZE, BLOCK_SIZE, color)
	rl.DrawRectangleLines(posX, posY, BLOCK_SIZE, BLOCK_SIZE, rl.Black)
}

type Position struct {
	col, row int32
}

type Block interface {
	draw(occupied [COLS][ROWS]bool, isNew bool)
	setOccupied(occupied *[COLS][ROWS]bool)
	isOnOccupied(occupied [COLS][ROWS]bool) bool
	moveUp()
	moveDown()
	moveLeft()
	moveRight()
	rotate()
	mirror()
}

type SquareBlock struct {
	Position

	size  int32
	color rl.Color
}

func newSquareBlock(col, row, size int32) *SquareBlock {
	s := SquareBlock{
		Position{col, row},
		size,
		randomColor(),
	}
	logger.Printf("[DEBUG] created new square=%+v", s)
	return &s
}

func (s SquareBlock) draw(occupied [COLS][ROWS]bool, isNew bool) {
	for xi := range s.size {
		x := (s.col+xi)*BLOCK_SIZE + LEFT_MARGIN
		for yi := range s.size {
			y := (s.row+yi)*BLOCK_SIZE + TOP_MARGIN
			if isNew || occupied[s.col+xi][s.row+yi] {
				drawBlock(x, y, s.color)
			}
		}
	}
}

func (s SquareBlock) setOccupied(occupied *[COLS][ROWS]bool) {
	for col := range s.size {
		for row := range s.size {
			logger.Printf("[DEBUG] square(%+v).setOccupied(col=%d, row=%d)", s, s.col+col, s.row+row)
			occupied[s.col+col][s.row+row] = true
		}
	}
}

func (s SquareBlock) isOnOccupied(occupied [COLS][ROWS]bool) bool {
	for col := range s.size {
		for row := range s.size {
			if occupied[s.col+col][s.row+row] {
				logger.Printf("[DEBUG] square(%+v).isOnOccupied(col=%d, row=%d)", s, col, row)
				return true
			}
		}
	}
	return false
}

func (s *SquareBlock) moveUp() {
	if s.row <= 0 {
		return
	}
	s.row--
}

func (s *SquareBlock) moveDown() {
	if s.row+s.size >= ROWS {
		return
	}
	s.row++
}

func (s *SquareBlock) moveLeft() {
	if s.col <= 0 {
		return
	}
	s.col--
}

func (s *SquareBlock) moveRight() {
	if s.col+s.size >= COLS {
		return
	}
	s.col++
}

func (s *SquareBlock) rotate() {}

func (s *SquareBlock) mirror() {}

type LineBlock struct {
	Position
	size  int32
	color rl.Color
}

func newLineBlock(col, row, size int32) *LineBlock {
	l := LineBlock{
		Position{col, row},
		size,
		randomColor(),
	}
	logger.Printf("[DEBUG] created new line=%+v", l)
	return &l
}

func (l LineBlock) draw(occupied [COLS][ROWS]bool, isNew bool) {
	for yi := range l.size {
		x := l.col*BLOCK_SIZE + LEFT_MARGIN
		y := (l.row+yi)*BLOCK_SIZE + TOP_MARGIN
		if isNew || occupied[l.col][l.row+yi] {
			drawBlock(x, y, l.color)
		}
	}
}

func (l LineBlock) setOccupied(occupied *[COLS][ROWS]bool) {
	// TODO: switch on orientation (horizontal/vertical)
	for row := range l.size {
		logger.Printf("[DEBUG] line(%+v).setOccupied(col=%d, row=%d)", l, l.col, l.row+row)
		occupied[l.col][l.row+row] = true
	}
}

func (l LineBlock) isOnOccupied(occupied [COLS][ROWS]bool) bool {
	// TODO: switch on orientation (horizontal/vertical)
	for row := range l.size {
		if occupied[l.col][l.row+row] {
			logger.Printf("[DEBUG] line(%+v).isOnOccupied(col=%d, row=%d)", l, l.col, row)
			return true
		}
	}
	return false
}

func (l *LineBlock) moveUp() {
	// TODO: switch on orientation (horizontal/vertical)
	if l.row <= 0 {
		return
	}
	l.row--
}

func (l *LineBlock) moveDown() {
	// TODO: switch on orientation (horizontal/vertical)
	if l.row+l.size >= ROWS {
		return
	}
	l.row++
}

func (l *LineBlock) moveLeft() {
	// TODO: switch on orientation (horizontal/vertical)
	if l.col <= 0 {
		return
	}
	l.col--
}

func (l *LineBlock) moveRight() {
	// TODO: switch on orientation (horizontal/vertical)
	if l.col >= COLS {
		return
	}
	l.col++
}

func (l *LineBlock) rotate() {
	// TODO: switch orientation
}

func (l *LineBlock) mirror() {}

func drawBoard() {
	rl.DrawRectangle(LEFT_MARGIN, TOP_MARGIN, G_WIDTH, G_HEIGHT, rl.LightGray)
	rl.SetLineWidth(STROKE)
	for i := range COLS + 1 {
		xBase := int32(i*BLOCK_SIZE) + LEFT_MARGIN
		rl.DrawLine(xBase, TOP_MARGIN, xBase, TOP_MARGIN+G_HEIGHT, rl.Gray)
	}
	for i := range ROWS + 1 {
		yBase := int32(i*BLOCK_SIZE) + TOP_MARGIN
		rl.DrawLine(LEFT_MARGIN, yBase, LEFT_MARGIN+G_WIDTH, yBase, rl.Gray)
	}
}

func resetOccupied(occupied *[COLS][ROWS]bool) {
	for col := range occupied {
		for row := range occupied[col] {
			occupied[col][row] = false
		}
	}
}

func isUpPressed() bool {
	return rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyK)
}

func isDownPressed() bool {
	return rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyJ)
}

func isLeftPressed() bool {
	return rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyH)
}

func isRightPressed() bool {
	return rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyL)
}

func isRotatePressed() bool {
	return rl.IsKeyPressed(rl.KeyR)
}

func isMirrorPressed() bool {
	return rl.IsKeyPressed(rl.KeyM)
}

func isPlacePressed() bool {
	return rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)
}

func handleInput(block Block, occupied [COLS][ROWS]bool) bool {
	if isUpPressed() {
		logger.Println("[DEBUG] UP pressed")
		block.moveUp()
		return false
	}
	if isDownPressed() {
		logger.Println("[DEBUG] DOWN pressed")
		block.moveDown()
		return false
	}
	if isLeftPressed() {
		logger.Println("[DEBUG] LEFT pressed")
		block.moveLeft()
		return false
	}
	if isRightPressed() {
		logger.Println("[DEBUG] RIGHT pressed")
		block.moveRight()
		return false
	}
	if isRotatePressed() {
		logger.Println("[DEBUG] ROTATE pressed")
		block.rotate()
		return false
	}
	if isMirrorPressed() {
		logger.Println("[DEBUG] MIRROR pressed")
		block.mirror()
		return false
	}
	if isPlacePressed() {
		logger.Println("[DEBUG] PLACE pressed")
		return !block.isOnOccupied(occupied)
	}
	return false
}

func nextBlock() Block {
	const nBlockTypes = 2
	blockIdx := rand.Int() % nBlockTypes
	var block Block
	switch blockIdx {
	case 0:
		size := int32(rand.Int()%3 + 1)
		block = newSquareBlock((COLS/2)-(size/2), (ROWS/2)-(size/2), size)
	case 1:
		sizes := []int32{2, 3, 4, 5}
		sizeIdx := rand.Int() % len(sizes)
		size := sizes[sizeIdx]
		// TODO: make dependent on orientation
		block = newLineBlock(COLS/2, (ROWS/2)-(size/2), size)
	}
	return block
}

func isColumnFull(occupied [COLS][ROWS]bool, col int32) bool {
	for row := range ROWS {
		if !occupied[col][row] {
			return false
		}
	}
	logger.Printf("[DEBUG] column %d is full", col)
	return true
}

func getFullColumns(occupied [COLS][ROWS]bool) []int32 {
	fullCols := []int32{}
	for col := range COLS {
		if isColumnFull(occupied, col) {
			fullCols = append(fullCols, col)
		}
	}
	return fullCols
}

func isRowFull(occupied [COLS][ROWS]bool, row int32) bool {
	for col := range COLS {
		if !occupied[col][row] {
			return false
		}
	}
	logger.Printf("[DEBUG] row %d is full", row)
	return true
}

func getFullRows(occupied [COLS][ROWS]bool) []int32 {
	fullRows := []int32{}
	for row := range ROWS {
		if isRowFull(occupied, row) {
			fullRows = append(fullRows, row)
		}
	}
	return fullRows
}

func unoccupyFull(occupied *[COLS][ROWS]bool) {
	fullCols := getFullColumns(*occupied)
	fullRows := getFullRows(*occupied)
	for col := range COLS {
		for row := range ROWS {
			if slices.Contains(fullCols, col) || slices.Contains(fullRows, row) {
				logger.Printf("[DEBUG] unoccupying col=%d row=%d", col, row)
				occupied[col][row] = false
			}
		}
	}
}

func gameLoop() {
	rl.InitWindow(W_WIDTH, W_HEIGHT, TITLE)
	defer rl.CloseWindow()
	rl.SetTargetFPS(FPS)

	occupied := [COLS][ROWS]bool{}
	resetOccupied(&occupied)

	availableBlocks := []Block{}
	for range 3 {
		availableBlocks = append(availableBlocks, nextBlock())
	}
	curBlockIdx := 0
	curBlock := availableBlocks[curBlockIdx]

	placedBlocks := []Block{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		drawBoard()

		for _, block := range placedBlocks {
			block.draw(occupied, false)
		}

		placed := handleInput(curBlock, occupied)
		curBlock.draw(occupied, true)
		if placed {
			logger.Println("[DEBUG] block can be placed")
			placedBlocks = append(placedBlocks, curBlock)
			logger.Printf("[DEBUG] placed blocks=%+v", placedBlocks)
			curBlock.setOccupied(&occupied)
			unoccupyFull(&occupied)
			availableBlocks[curBlockIdx] = nextBlock()
			curBlock = availableBlocks[curBlockIdx]
		}

		rl.EndDrawing()
	}
}

func main() {
	gameLoop()
}
