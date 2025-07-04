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
* - add help
* - plus block
* - L block
* - zigzag block
* - more depth on block grid
* - fix type size
* - gather points for destroying rows and cols
* - save ranking and display records
 */

const (
	TITLE                            = "GoBlocks"
	WIN_WIDTH                int32   = 960
	WIN_HEIGHT               int32   = 820
	FPS                      int32   = 60
	TOP_MARGIN               int32   = 20
	LEFT_MARGIN              int32   = 30
	RIGHT_MARGIN             int32   = LEFT_MARGIN
	BOTTOM_MARGIN            int32   = 200
	BLOCK_SIZE               int32   = 60
	GAME_WIDTH               int32   = WIN_WIDTH - LEFT_MARGIN - RIGHT_MARGIN
	GAME_HEIGHT              int32   = WIN_HEIGHT - TOP_MARGIN - BOTTOM_MARGIN
	GAME_START_X             int32   = LEFT_MARGIN
	GAME_END_X               int32   = GAME_START_X + GAME_WIDTH
	GAME_START_Y             int32   = TOP_MARGIN
	GAME_END_Y               int32   = GAME_START_Y + GAME_HEIGHT
	GAME_BLOCK_SCALE         float32 = 1
	AVAILABLE_BLOCKS_WIDTH   int32   = (WIN_WIDTH-LEFT_MARGIN)/2 + 2
	AVAILABLE_BLOCKS_HEIGHT  int32   = BOTTOM_MARGIN - 2*TOP_MARGIN
	AVAILABLE_BLOCKS_START_X int32   = LEFT_MARGIN
	AVAILABLE_BLOCKS_END_X   int32   = AVAILABLE_BLOCKS_START_X + AVAILABLE_BLOCKS_WIDTH
	AVAILABLE_BLOCKS_START_Y int32   = GAME_END_Y + TOP_MARGIN
	AVAILABLE_BLOCKS_END_Y   int32   = AVAILABLE_BLOCKS_START_Y + AVAILABLE_BLOCKS_HEIGHT
	AVAILABLE_BLOCK_SCALE    float32 = 1.0 / 2
	AVAILABLE_BLOCK_SIZE     int32   = int32(float32(BLOCK_SIZE) * AVAILABLE_BLOCK_SCALE)
	COLS                     int32   = GAME_WIDTH / BLOCK_SIZE
	ROWS                     int32   = GAME_HEIGHT / BLOCK_SIZE
	STROKE                   float32 = 2
)

var (
	logger   *log.Logger = log.New(os.Stdout, "goblocks:", log.LstdFlags)
	showHelp             = false
)

func randomColor() rl.Color {
	colors := []rl.Color{rl.Red, rl.Orange, rl.Yellow, rl.Green, rl.Blue, rl.Pink}
	colorIdx := rand.Int() % len(colors)
	return colors[colorIdx]
}

func drawBlock(posX, posY int32, color rl.Color, highlight bool, scale float32) {
	if scale == 0 {
		scale = 1
	}
	blockSize := int32(float32(BLOCK_SIZE) * scale)
	rl.DrawRectangle(posX, posY, blockSize, blockSize, color)
	borderColor := rl.Black
	if highlight {
		borderColor = rl.White
	}
	rl.DrawRectangleLines(posX, posY, blockSize, blockSize, borderColor)
}

type Position struct {
	col, row int32
}

type Block interface {
	draw(occupied [COLS][ROWS]bool, highlight bool)
	drawAsAvailable(idx int32, selected bool)
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

func (s SquareBlock) draw(occupied [COLS][ROWS]bool, highlight bool) {
	for xi := range s.size {
		x := (s.col+xi)*BLOCK_SIZE + LEFT_MARGIN
		for yi := range s.size {
			y := (s.row+yi)*BLOCK_SIZE + TOP_MARGIN
			if highlight || occupied[s.col+xi][s.row+yi] {
				drawBlock(x, y, s.color, highlight, GAME_BLOCK_SCALE)
			}
		}
	}
}

func (s SquareBlock) drawAsAvailable(idx int32, selected bool) {
	startX := AVAILABLE_BLOCKS_START_X + idx*AVAILABLE_BLOCKS_WIDTH/3 + AVAILABLE_BLOCKS_WIDTH/6 - s.size*AVAILABLE_BLOCK_SIZE/2.0 + idx
	startY := AVAILABLE_BLOCKS_START_Y + AVAILABLE_BLOCKS_HEIGHT/2 - s.size*AVAILABLE_BLOCK_SIZE/2.0
	for xi := range s.size {
		x := startX + xi*AVAILABLE_BLOCK_SIZE
		for yi := range s.size {
			y := startY + yi*AVAILABLE_BLOCK_SIZE
			drawBlock(x, y, s.color, selected, AVAILABLE_BLOCK_SCALE)
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

type LineOrientation string

const (
	LINE_HORIZONTAL LineOrientation = "horizontal"
	LINE_VERTICAL   LineOrientation = "vertical"
)

type LineBlock struct {
	Position
	size        int32
	color       rl.Color
	orientation LineOrientation
}

func newLineBlock(col, row, size int32) *LineBlock {
	l := LineBlock{
		Position{col, row},
		size,
		randomColor(),
		LINE_VERTICAL,
	}
	logger.Printf("[DEBUG] created new line=%+v", l)
	return &l
}

func (l LineBlock) draw(occupied [COLS][ROWS]bool, isNew bool) {
	if l.orientation == LINE_VERTICAL {
		for yi := range l.size {
			x := l.col*BLOCK_SIZE + LEFT_MARGIN
			y := (l.row+yi)*BLOCK_SIZE + TOP_MARGIN
			if isNew || occupied[l.col][l.row+yi] {
				drawBlock(x, y, l.color, isNew, GAME_BLOCK_SCALE)
			}
		}
	} else {
		for xi := range l.size {
			x := (l.col+xi)*BLOCK_SIZE + LEFT_MARGIN
			y := l.row*BLOCK_SIZE + TOP_MARGIN
			if isNew || occupied[l.col+xi][l.row] {
				drawBlock(x, y, l.color, isNew, GAME_BLOCK_SCALE)
			}
		}
	}
}

func (l LineBlock) drawAsAvailable(idx int32, selected bool) {
	startX := AVAILABLE_BLOCKS_START_X + idx*AVAILABLE_BLOCKS_WIDTH/3 + AVAILABLE_BLOCKS_WIDTH/6 - l.size*AVAILABLE_BLOCK_SIZE/2.0 + idx
	startY := AVAILABLE_BLOCKS_START_Y + AVAILABLE_BLOCKS_HEIGHT/2 - AVAILABLE_BLOCK_SIZE/2.0
	for xi := range l.size {
		x := startX + xi*AVAILABLE_BLOCK_SIZE
		y := startY
		drawBlock(x, y, l.color, selected, AVAILABLE_BLOCK_SCALE)
	}
}

func (l LineBlock) setOccupied(occupied *[COLS][ROWS]bool) {
	if l.orientation == LINE_VERTICAL {
		for row := range l.size {
			logger.Printf("[DEBUG] line(%+v).setOccupied(col=%d, row=%d)", l, l.col, l.row+row)
			occupied[l.col][l.row+row] = true
		}
	} else {
		for col := range l.size {
			logger.Printf("[DEBUG] line(%+v).setOccupied(col=%d, row=%d)", l, l.col+col, l.row)
			occupied[l.col+col][l.row] = true
		}
	}
}

func (l LineBlock) isOnOccupied(occupied [COLS][ROWS]bool) bool {
	if l.orientation == LINE_VERTICAL {
		for row := range l.size {
			if occupied[l.col][l.row+row] {
				logger.Printf("[DEBUG] line(%+v).isOnOccupied(col=%d, row=%d)", l, l.col, l.row+row)
				return true
			}
		}
	} else {
		for col := range l.size {
			if occupied[l.col+col][l.row] {
				logger.Printf("[DEBUG] line(%+v).isOnOccupied(col=%d, row=%d)", l, l.col+col, l.row)
				return true
			}
		}
	}
	return false
}

func (l *LineBlock) moveUp() {
	if l.row <= 0 {
		return
	}
	l.row--
}

func (l *LineBlock) moveDown() {
	if (l.orientation == LINE_VERTICAL && l.row+l.size < ROWS) || (l.orientation == LINE_HORIZONTAL && l.row < ROWS-1) {
		l.row++
	}
}

func (l *LineBlock) moveLeft() {
	if l.col > 0 {
		l.col--
	}
}

func (l *LineBlock) moveRight() {
	if (l.orientation == LINE_VERTICAL && l.col < COLS-1) || (l.orientation == LINE_HORIZONTAL && l.col+l.size < COLS) {
		l.col++
	}
}

func (l *LineBlock) rotate() {
	if l.orientation == LINE_VERTICAL {
		l.orientation = LINE_HORIZONTAL
	} else {
		l.orientation = LINE_VERTICAL
	}
}

func (l *LineBlock) mirror() {}

func drawBoard() {
	rl.DrawRectangle(LEFT_MARGIN, TOP_MARGIN, GAME_WIDTH, GAME_HEIGHT, rl.LightGray)
	rl.SetLineWidth(STROKE)
	for i := range COLS + 1 {
		xBase := int32(i*BLOCK_SIZE) + LEFT_MARGIN
		rl.DrawLine(xBase, TOP_MARGIN, xBase, TOP_MARGIN+GAME_HEIGHT, rl.Gray)
	}
	for i := range ROWS + 1 {
		yBase := int32(i*BLOCK_SIZE) + TOP_MARGIN
		rl.DrawLine(LEFT_MARGIN, yBase, LEFT_MARGIN+GAME_WIDTH, yBase, rl.Gray)
	}
}

func drawAvailable(blocks []Block, curBlockIdx int) {
	rl.DrawRectangle(AVAILABLE_BLOCKS_START_X, AVAILABLE_BLOCKS_START_Y, AVAILABLE_BLOCKS_WIDTH, AVAILABLE_BLOCKS_HEIGHT, rl.Gray)
	rl.SetLineWidth(1)
	sepStartX := AVAILABLE_BLOCKS_START_X + AVAILABLE_BLOCKS_WIDTH/3 + 1
	sepStartY := AVAILABLE_BLOCKS_START_Y
	sepEndY := sepStartY + AVAILABLE_BLOCKS_HEIGHT
	rl.DrawLine(sepStartX, sepStartY, sepStartX, sepEndY, rl.DarkGray)
	sepStartX = sepStartX + AVAILABLE_BLOCKS_WIDTH/3 + 1
	rl.DrawLine(sepStartX, sepStartY, sepStartX, sepEndY, rl.DarkGray)
	for idx, block := range blocks {
		selected := false
		if idx == curBlockIdx {
			selected = true
		}
		block.drawAsAvailable(int32(idx), selected)
	}
	rl.SetLineWidth(STROKE)
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

func isPrevBlockPressed() bool {
	return rl.IsKeyPressed(rl.KeyP)
}

func isNextBlockPressed() bool {
	return rl.IsKeyPressed(rl.KeyN)
}

func isHelpPressed() bool {
	if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) {
		return rl.IsKeyPressed(rl.KeySlash)
	}
	if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl) {
		return rl.IsKeyPressed(rl.KeyH)
	}
	return false
}

func handleInput(block Block, occupied [COLS][ROWS]bool, curBlockIdx *int) bool {
	if isHelpPressed() {
		logger.Println("[DEBUG] help pressed")
		showHelp = !showHelp
		return false
	}
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
	if isPrevBlockPressed() && *curBlockIdx > 0 {
		*curBlockIdx--
		logger.Printf("[DEBUG] selecting previous block, current idx=%d", *curBlockIdx)
		return false
	}
	if isNextBlockPressed() && *curBlockIdx < 2 {
		*curBlockIdx++
		logger.Printf("[DEBUG] selecting next block, current idx=%d", *curBlockIdx)
		return false
	}
	return false
}

func showHelpWindow() {
	rl.DrawRectangle(GAME_START_X, GAME_START_Y, GAME_WIDTH, GAME_HEIGHT, rl.Fade(rl.SkyBlue, 0.9))
	padding := int32(6)
	fontSize := int32(30)
	textPositionX := GAME_START_X + padding
	textPositionY := GAME_START_Y + padding
	rl.DrawText("Controls:", textPositionX, textPositionY, fontSize, rl.Black)
	lines := [][]string{
		{"Toggle help:", "Ctrl+H or ?"},
		{"Close game:", "ESCAPE"},
		{"Move up:", "W or ARROW_UP or K"},
		{"Move down:", "S or ARROW_DOWN or J"},
		{"Move left:", "A or ARROW_LEFT or H"},
		{"Move right:", "D or ARROW_RIGHT or L"},
		{"Place block:", "SPACE or ENTER"},
		{"Rotate block:", "R"},
		{"Mirror block:", "M"},
		{"Select next block:", "N"},
		{"Select prev block:", "P"},
	}
	fontSize = 24
	for _, line := range lines {
		textPositionY += fontSize + int32(padding)
		rl.DrawText(line[0], textPositionX, textPositionY, fontSize, rl.DarkGray)
		rl.DrawText(line[1], textPositionX+300, textPositionY, fontSize, rl.DarkGray)
	}
}

func randomBlock() Block {
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
	rl.InitWindow(WIN_WIDTH, WIN_HEIGHT, TITLE)
	defer rl.CloseWindow()
	rl.SetTargetFPS(FPS)

	occupied := [COLS][ROWS]bool{}
	resetOccupied(&occupied)

	availableBlocks := []Block{}
	for range 3 {
		availableBlocks = append(availableBlocks, randomBlock())
	}
	curBlockIdx := 0
	curBlock := availableBlocks[curBlockIdx]

	placedBlocks := []Block{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		drawBoard()
		drawAvailable(availableBlocks, curBlockIdx)

		for _, block := range placedBlocks {
			block.draw(occupied, false)
		}

		placed := handleInput(curBlock, occupied, &curBlockIdx)
		if showHelp {
			showHelpWindow()
		} else {
			curBlock = availableBlocks[curBlockIdx]
			curBlock.draw(occupied, true)
			if placed {
				logger.Println("[DEBUG] block can be placed")
				placedBlocks = append(placedBlocks, curBlock)
				logger.Printf("[DEBUG] placed blocks=%+v", placedBlocks)
				curBlock.setOccupied(&occupied)
				unoccupyFull(&occupied)
				availableBlocks[curBlockIdx] = randomBlock()
			}
		}

		rl.EndDrawing()
	}
}

func main() {
	gameLoop()
}
