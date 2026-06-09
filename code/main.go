package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ! running constants here
const shouldCloseAfter5Sec bool = true
const screenSize uint = 800
const textureScale float32 = 2

var boidTexture rl.Texture2D

func main() {
	boidTexture = rl.LoadTexture("../assets/clown-fish.png")
	applyTimeout(shouldCloseAfter5Sec)

	rl.UnloadTexture(boidTexture)
}

// in go the init function is always called before main for setup stuff
func init() {
	rl.InitWindow(int32(screenSize), int32(screenSize), "boids")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
	}
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
