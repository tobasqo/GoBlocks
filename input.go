package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func IsUpPressed() bool {
	return rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyK)
}

func IsDownPressed() bool {
	return rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyJ)
}

func IsLeftPressed() bool {
	return rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyH)
}

func IsRightPressed() bool {
	return rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyL)
}

func IsRotatePressed() bool {
	return rl.IsKeyPressed(rl.KeyR)
}

func IsMirrorPressed() bool {
	return rl.IsKeyPressed(rl.KeyM)
}

func IsPlacePressed() bool {
	return rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)
}

func IsPrevBlockPressed() bool {
	return rl.IsKeyPressed(rl.KeyP)
}

func IsNextBlockPressed() bool {
	return rl.IsKeyPressed(rl.KeyN)
}

func IsHelpPressed() bool {
	if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) {
		return rl.IsKeyPressed(rl.KeySlash)
	}
	if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl) {
		return rl.IsKeyPressed(rl.KeyH)
	}
	return false
}
