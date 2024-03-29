package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	LOGO int = iota
	TITLE
	GAMEPLAY
	ENDING
)

func main() {
	const screenWidth = 800
	const screenHeight = 450

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic screen manager")

	currentScreen := LOGO

	framesCounter := 0  // Useful to count frames
	rl.SetTargetFPS(60) // Set desired framerate (frames-per-second)

	for !rl.WindowShouldClose() {
		switch currentScreen {
		case LOGO:
			framesCounter++ // Count frames
			if framesCounter > 120 {
				currentScreen = TITLE
			}
			break
		case TITLE:
			if rl.IsKeyPressed(rl.KeyEnter) || rl.IsGestureDetected(rl.GestureTap) {
				currentScreen = GAMEPLAY
			}
			break
		case GAMEPLAY:
			if rl.IsKeyPressed(rl.KeyEnter) || rl.IsGestureDetected(rl.GestureTap) {
				currentScreen = ENDING
			}
			break
		case ENDING:
			if rl.IsKeyPressed(rl.KeyEnter) || rl.IsGestureDetected(rl.GestureTap) {
				currentScreen = TITLE
			}
			break
		default:
			break
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		switch currentScreen {
		case LOGO:
			rl.DrawText("LOGO SCREEN", 20, 20, 40, rl.LightGray)
			rl.DrawText("WAIT for 2 SECONDS...", 290, 220, 20, rl.Gray)
			break
		case TITLE:
			rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Green)
			rl.DrawText("TITLE SCREEN", 20, 20, 40, rl.DarkGreen)
			rl.DrawText("PRESS ENTER or TAP to JUMP to GAMEPLAY SCREEN", 120, 220, 20, rl.DarkGreen)

			break
		case GAMEPLAY:
			rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Purple)
			rl.DrawText("GAMEPLAY SCREEN", 20, 20, 40, rl.Maroon)
			rl.DrawText("PRESS ENTER or TAP to JUMP to ENDING SCREEN", 130, 220, 20, rl.Maroon)

			break
		case ENDING:
			rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Blue)
			rl.DrawText("ENDING SCREEN", 20, 20, 40, rl.DarkBlue)
			rl.DrawText("PRESS ENTER or TAP to RETURN to TITLE SCREEN", 120, 220, 20, rl.DarkBlue)

			break
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
