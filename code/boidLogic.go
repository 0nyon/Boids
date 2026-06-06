package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ! all constants here
const boidAmnt uint = 20
const boidSightRange float32 = 20

type boid struct {
	posVec      rl.Vector2
	velocityVec rl.Vector2
	texture     *rl.Texture2D
}

func createAllBoids(tex *rl.Texture2D) []boid {
	res := make([]boid, boidAmnt)
	for i := uint(0); i < boidAmnt; i++ {
		res[i] = boid{
			posVec: rl.Vector2{
				X: rand.Float32() * (float32(screenSize) - 1),
				Y: rand.Float32() * (float32(screenSize) - 1),
			},
			velocityVec: rl.Vector2{
				X: rand.Float32()*2 - 1,
				Y: rand.Float32()*2 - 1,
			}.Normalize(),
			texture: tex,
		}

	}

	return res
}

func applySeperationToAllBoids(boids []boid) {
	for i := range boids {
		boids[i].applySeperationToBoid(boids[i].getAllboidsInSightRange(boids))
	}
}

// * this function can take the slice of all boids, and one boid inside it
func (this *boid) applySeperationToBoid(boidsInRange []boid) {
	for _, b := range boidsInRange {
		this.velocityVec = this.velocityVec.Add(this.velocityVec.Add(b.velocityVec.Subtract(this.velocityVec)))
	}

	this.velocityVec = this.velocityVec.Normalize()
}

// * this function can take the slice of all boids, and one boid inside it
func (this *boid) getAllboidsInSightRange(allBoids []boid) []boid {
	res := make([]boid, 0)
	for _, b := range allBoids {
		if *this == b {
			continue
		}
		if this.posVec.Distance(b.posVec) <= boidSightRange {
			res = append(res, b)
		}
	}
	return res
}

func (b *boid) getLookingAngle() float32 {
	return 360 * (b.velocityVec.Angle(rl.Vector2{X: 1, Y: 0}))
}

func applyAllRules(boids []boid) {
	//! fun thing about go, a for range loop that refers to a value like so for _, v:= range ...
	//! actually makes a copy and doesn't mutate the refernces, keep that in mind
	applySeperationToAllBoids(boids)
}
