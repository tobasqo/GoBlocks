package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameWindow struct {
	title           string
	width           int32
	height          int32
	board           *Board
	availableBlocks *AvailableBlocks
}

type Game struct {
	window   gameWindow
	fps      int32
	showHelp bool
}

func NewGame() Game {
	var width int32 = 960
	var height int32 = 820
	var topMargin int32 = 20
	var leftMargin int32 = 30
	rightMargin := leftMargin
	bottomMargin := topMargin
	var cols int32 = 15
	var rows int32 = 10

	boardWidth := width - leftMargin - rightMargin
	blockSize := boardWidth / cols
	board := NewBoard(leftMargin, topMargin, leftMargin+boardWidth, height-topMargin-2*bottomMargin-140, blockSize, cols, rows)

	availableBlocksStartX := leftMargin
	availableBlocksStartY := board.EndY() + bottomMargin
	availableBlocksEndX := leftMargin + boardWidth/2 + 2
	availableBlocksEndY := height - bottomMargin
	availableBlocks := NewAvailableBlocks(availableBlocksStartX, availableBlocksStartY, availableBlocksEndX, availableBlocksEndY, blockSize/2, cols, rows)

	return Game{
		gameWindow{
			"GoBlocks",
			width,
			height,
			board,
			availableBlocks,
		},
		60,
		false,
	}
}

func (g Game) Start() {
	w := g.window
	b := w.board
	ab := w.availableBlocks

	rl.InitWindow(w.width, w.height, w.title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(g.fps)

	b.ResetOccupied()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		b.Draw()
		ab.Draw()

		block := ab.SelectedBlock()
		block.Draw(*b)
		g.HandleInput()

		rl.EndDrawing()
	}
}

func (g Game) ShowHelpWindow() {
	b := g.window.board
	rl.DrawRectangle(b.StartX(), b.StartY(), b.Width(), b.Height(), rl.Fade(rl.SkyBlue, 0.9))
	padding := int32(6)
	fontSize := int32(30)
	textPositionX := b.StartX() + padding
	textPositionY := b.StartY() + padding
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

func (g *Game) HandleInput() {
	if IsHelpPressed() {
		log.Println("[DEBUG] HELP pressed")
		g.showHelp = !g.showHelp
		g.ShowHelpWindow()
		return
	}
	b := g.window.board
	ab := g.window.availableBlocks
	block := ab.blocks[ab.curBlockIdx]
	if IsUpPressed() {
		log.Println("[DEBUG] UP pressed")
		block.MoveUp()
		return
	}
	if IsDownPressed() {
		log.Println("[DEBUG] DOWN pressed")
		block.MoveDown(b.rows)
		return
	}
	if IsLeftPressed() {
		log.Println("[DEBUG] LEFT pressed")
		block.MoveLeft()
		return
	}
	if IsRightPressed() {
		log.Println("[DEBUG] RIGHT pressed")
		block.MoveRight(b.cols)
		return
	}
	if IsRotatePressed() {
		log.Println("[DEBUG] ROTATE pressed")
		block.Rotate(*b)
		return
	}
	if IsMirrorPressed() {
		log.Println("[DEBUG] MIRROR pressed")
		block.Mirror(*b)
		return
	}
	if IsPlacePressed() {
		log.Println("[DEBUG] PLACE pressed")
		if !block.IsOnOccupied(*b) {
			log.Println("[DEBUG] block can be placed")
			block = ab.PopSelected(b.cols, b.rows)
			b.AddBlock(block)
			log.Printf("[DEBUG] placed blocks=%+v", b.blocks)
			block.SetOccupied(b)
			b.UnoccupyFull()
		}
	}
	if IsPrevBlockPressed() {
		log.Println("[DEBUG] PREV pressed")
		ab.SelectPrev()
		return
	}
	if IsNextBlockPressed() {
		log.Println("[DEBUG] NEXT pressed")
		ab.SelectNext()
		return
	}
}
