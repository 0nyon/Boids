package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ! all constants here
const boidAmnt uint = 50
const boidSightRange float32 = 50
const tooCloseRange float32 = 30
const cohesionSteeringMult float32 = 0.15 //! has to be between 0-1
const alignmentSteeringMult float32 = 0.4 //! has to be between 0-1
const seperationSteeringMult float32 = 1  //! has to be between 0-1

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

func (this *boid) getAllboidsInSightRange(allBoids []boid) []boid {
	res := make([]boid, 0)
	for i := range allBoids {
		b := allBoids[i]
		if *this == b {
			continue
		}
		if this.posVec.Distance(b.posVec) <= boidSightRange {
			res = append(res, b)
		}
	}
	return res
}

func (this *boid) getAllboidsInTooCloseRange(allBoids []boid) []boid {
	res := make([]boid, 0)
	for i := range allBoids {
		b := allBoids[i]
		if *this == b {
			continue
		}
		if this.posVec.Distance(b.posVec) <= tooCloseRange {
			res = append(res, b)
		}
	}
	return res
}

func (b *boid) getLookingAngle() float32 {
	return 360 * (b.velocityVec.Angle(rl.Vector2{X: 1, Y: 0}))
}

