package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_COLUMNS = 20

func main() {
	const screenWidth = 800
	const screenHeight = 450

	rl.InitWindow(screenWidth, screenHeight, "raylib example - 3d camera")

	camera := rl.Camera3D{}

	camUp := rl.Vector3{
		X: 0.0,
		Y: 1.0,
		Z: 0.0,
	}

	camera.Position = rl.Vector3{X: 0.0, Y: 2.0, Z: 4.0}
	camera.Target = rl.Vector3{X: 0.0, Y: 2.0, Z: 0.0}
	camera.Up = camUp
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	cameraMode := rl.CameraFirstPerson
	heights := make([]float32, MAX_COLUMNS)
	positions := make([]rl.Vector3, MAX_COLUMNS)
	colors := make([]rl.Color, MAX_COLUMNS)

	for i := 0; i < MAX_COLUMNS; i++ {
		heights[i] = float32(rl.GetRandomValue(1, 12))
		positions[i] = rl.Vector3{
			X: float32(rl.GetRandomValue(-15, 15)),
			Y: heights[i] / 2,
			Z: float32(rl.GetRandomValue(-15, 15)),
		}
		colors[i] = rl.Color{
			R: uint8(rl.GetRandomValue(20, 255)),
			G: uint8(rl.GetRandomValue(10, 55)),
			B: 30,
			A: 255,
		}
	}

	rl.DisableCursor()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyOne) {
			cameraMode = rl.CameraFree
			camera.Up = camUp
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			cameraMode = rl.CameraFirstPerson
			camera.Up = camUp
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			cameraMode = rl.CameraThirdPerson
			camera.Up = camUp
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			cameraMode = rl.CameraOrbital
			camera.Up = camUp
		}

		if rl.IsKeyPressed(rl.KeyP) {
			if camera.Projection == rl.CameraPerspective {
				cameraMode = rl.CameraThirdPerson
				camera.Position = rl.Vector3{X: 0.0, Y: 2.0, Z: -100.0}
				camera.Target = rl.Vector3{X: 0.0, Y: 2.0, Z: 0.0}
				camera.Up = camUp
				camera.Projection = rl.CameraOrthographic
				camera.Fovy = 20.0 // near plane width in CAMERA_ORTHOGRAPHIC
				rl.CameraYaw(&camera, -135*rl.Deg2rad, 1)
				rl.CameraPitch(&camera, -45*rl.Deg2rad, 1, 1, 0)
			} else if camera.Projection == rl.CameraOrthographic {
				cameraMode = rl.CameraThirdPerson
				camera.Position = rl.Vector3{X: 0.0, Y: 2.0, Z: 10.0}
				camera.Target = rl.Vector3{X: 0.0, Y: 2.0, Z: 0.0}
				camera.Up = camUp
				camera.Projection = rl.CameraPerspective
				camera.Fovy = 60.0
			}
		}

		rl.UpdateCamera(&camera, cameraMode)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawPlane(
			rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
			rl.Vector2{X: 32.0, Y: 32.0},
			rl.LightGray,
		) // Draw ground

		rl.DrawCube(
			rl.Vector3{X: -16.0, Y: 2.5, Z: 0.0},
			1.0,
			5.0,
			32.0,
			rl.Blue,
		) // Draw a blue wall
		rl.DrawCube(
			rl.Vector3{X: 16.0, Y: 2.5, Z: 0.0},
			1.0,
			5.0,
			32.0,
			rl.Lime,
		) // Draw a green wall
		rl.DrawCube(
			rl.Vector3{X: 0.0, Y: 2.5, Z: 16.0},
			32.0,
			5.0,
			1.0,
			rl.Gold,
		) // Draw a yellow wall

		// Draw some cubes around
		for i := 0; i < MAX_COLUMNS; i++ {
			rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
			rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Maroon)
		}

		// Draw player cube
		if cameraMode == rl.CameraThirdPerson {
			rl.DrawCube(camera.Target, 0.5, 0.5, 0.5, rl.Green)
			rl.DrawCubeWires(camera.Target, 0.5, 0.5, 0.5, rl.DarkGreen)
		}

		rl.EndMode3D()

		// Draw info boxes
		rl.DrawRectangle(5, 5, 330, 100, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(5, 5, 330, 100, rl.Blue)
		//
		rl.DrawText("Camera controls:", 15, 15, 10, rl.Black)
		rl.DrawText("- Move keys: W, A, S, D, Space, Left-Ctrl", 15, 30, 10, rl.Black)
		rl.DrawText("- Look around: arrow keys or mouse", 15, 45, 10, rl.Black)
		rl.DrawText("- Camera mode keys: 1, 2, 3, 4", 15, 60, 10, rl.Black)
		rl.DrawText("- Zoom keys: num-plus, num-minus or mouse scroll", 15, 75, 10, rl.Black)
		rl.DrawText("- Camera projection key: P", 15, 90, 10, rl.Black)

		rl.DrawRectangle(600, 5, 195, 100, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(600, 5, 195, 100, rl.Blue)

		rl.DrawText("Camera status:", 610, 15, 10, rl.Black)

		var camModeStr string
		var camPersStr string

		switch cameraMode {
		case rl.CameraFirstPerson:
			camModeStr = "FIRST_PERSON"
			break
		case rl.CameraFree:
			camModeStr = "FREE"
			break
		case rl.CameraThirdPerson:
			camModeStr = "THIRD_PERSON"
			break
		case rl.CameraOrbital:
			camModeStr = "ORBITAL"
			break
		}

		switch camera.Projection {
		case rl.CameraPerspective:
			camPersStr = "PERSPECTIVE"
			break
		case rl.CameraOrthographic:
			camPersStr = "ORTHOGRAPHIC"
			break
		}

		rl.DrawText(fmt.Sprintf("- Mode: %s", camModeStr), 610, 30, 10, rl.Black)
		rl.DrawText(fmt.Sprintf("- Projection: %s", camPersStr), 610, 45, 10, rl.Black)
		rl.DrawText(
			fmt.Sprintf(
				"- Position: (%06.3f, %06.3f, %06.3f)",
				camera.Position.X,
				camera.Position.Y,
				camera.Position.Z,
			),
			610,
			60,
			10,
			rl.Black,
		)
		rl.DrawText(
			fmt.Sprintf(
				"- Target: (%06.3f, %06.3f, %06.3f)",
				camera.Target.X,
				camera.Target.Y,
				camera.Target.Z,
			),
			610,
			75,
			10,
			rl.Black,
		)
		rl.DrawText(
			fmt.Sprintf("- Up: (%06.3f, %06.3f, %06.3f)", camera.Up.X, camera.Up.Y, camera.Up.Z),
			610,
			90,
			10,
			rl.Black,
		)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
