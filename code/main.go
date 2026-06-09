package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ! running constants here
const shouldCloseAfter5Sec bool = true
const screenSize uint = 500
const textureScale float32 = 2

var boidTexture rl.Texture2D

func main() {
	boidTexture = rl.LoadTexture("assets/clown-fish.png")
	applyTimeout(shouldCloseAfter5Sec)

	allBoids := createAllBoids(&boidTexture)

	for !rl.WindowShouldClose() {
		moveAllBoids(allBoids)
		renderAllBoids(allBoids)
	}

	rl.UnloadTexture(boidTexture)
}

// in go the init function is always called before main for setup stuff
func init() {
	rl.InitWindow(int32(screenSize), int32(screenSize), "boids")
	rl.SetTargetFPS(60)

}

func applyTimeout(yes bool) {
	if yes {

		go func() {
			after := time.After(time.Second * 5)
			select {
			case <-after:
				panic("enough time")
			}
		}()
	}
}

func renderAllBoids(boids []boid) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.White)
	for i := range boids {
		//b := boids[i]
		//rl.DrawTextureEx(boidTexture, b.posVec, b.getLookingAngle(), textureScale, rl.White)
		rl.DrawRectangle(int32(boids[i].posVec.X), int32(boids[i].posVec.Y), 10, 10, rl.Blue)
	}
}

func moveAllBoids(boids []boid) {
	applyAllRules(boids)
	for i := range boids {
		//todo: apply bounds checking
		boids[i].posVec = boids[i].posVec.Add(boids[i].velocityVec)
	}
}
