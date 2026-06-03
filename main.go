package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const shouldCloseAfter5Sec bool = true

func main() {

	if shouldCloseAfter5Sec {
		go func() {
			after := time.After(time.Second * 5)
			select {
			case <-after:
				panic("enough time")
			}
		}()
	}

	rl.InitWindow(500, 500, "boids")
	rl.SetTargetFPS(60)

	fih := rl.LoadTexture("assets/clown-fish.png")
	defer rl.UnloadTexture(fih)

	for !rl.WindowShouldClose() {
		//this is how you render a fish img, raylib provided a vector struct which is nice, i'll implement all the fun boids stuff tommorow, for now this
		//remains here as an exmaple
		position := rl.NewVector2(100, 100)

		rotation := float32(0.0)

		scale := float32(2.0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTextureEx(fih, position, rotation, scale, rl.White)
		rl.DrawTextureEx(fih, rl.NewVector2(100, 110), rotation, scale, rl.White)
		rl.EndDrawing()
	}
}
