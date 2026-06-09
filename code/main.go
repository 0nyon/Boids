package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ! running constants here
const shouldCloseAfter5Sec bool = false
const screenSize uint = 800
const textureScale float32 = 1.7
const backgroundTextureScale float32 = 8

var boidTexture rl.Texture2D
var backgroundTexture rl.Texture2D

func main() {
	boidTexture = rl.LoadTexture("../assets/bird.png")
	backgroundTexture = rl.LoadTexture("../assets/background.png")
	defer rl.UnloadTexture(boidTexture)
	defer rl.UnloadTexture(backgroundTexture)

	applyTimeout(shouldCloseAfter5Sec)

	all := createAllBoids(&boidTexture)

	for !rl.WindowShouldClose() {

		applyAllRules(all)
		rl.ClearBackground(rl.White)
		rl.DrawTextureEx(backgroundTexture, rl.Vector2{X: 0, Y: 0}, 0, backgroundTextureScale, rl.White)
		renderAllBoids(all, &boidTexture)
	}
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

func renderAllBoids(all []boid, texture *rl.Texture2D) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	len := len(all)
	for i := 0; i < len; i++ {
		rl.DrawTextureEx(*texture, all[i].posVec, all[i].getRotation(), textureScale, rl.White)
	}
}
