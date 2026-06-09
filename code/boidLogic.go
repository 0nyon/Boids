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

// * this function can take the slice of all boids, and one boid inside it
func (this *boid) applySeperationToBoid(boidsInRange []boid) rl.Vector2 {
	if len(boidsInRange) == 0 {
		return rl.Vector2{X: 0, Y: 0}
	}

	avgVec := rl.Vector2{X: 0, Y: 0}
	for i := range boidsInRange {
		b := boidsInRange[i]
		avgVec = avgVec.Add(this.posVec.Subtract(b.posVec))
	}

	avgVec.X /= float32(len(boidsInRange))
	avgVec.Y /= float32(len(boidsInRange))

	desiredVel := avgVec.Normalize()

	steeringVec := desiredVel.Subtract(this.velocityVec)
	steeringVec.X *= seperationSteeringMult
	steeringVec.Y *= seperationSteeringMult

	return steeringVec
}

// * this function can take the slice of all boids, and one boid inside it
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

func (b *boid) applyAlignmentToBoid(visableBoids []boid) rl.Vector2 {
	if len(visableBoids) == 0 {
		return rl.Vector2{X: 0, Y: 0}
	}
	sum := rl.Vector2{X: 0, Y: 0}
	for i := range visableBoids {
		visable := visableBoids[i]
		sum = sum.Add(visable.velocityVec)
	}

	sum.X /= float32(len(visableBoids))
	sum.Y /= float32(len(visableBoids))

	steeringVec := sum.Subtract(b.velocityVec)
	steeringVec.X *= alignmentSteeringMult
	steeringVec.Y *= alignmentSteeringMult

	return steeringVec
}

func (b *boid) applyCohesionToBoid(visableBoids []boid) rl.Vector2 {
	if len(visableBoids) == 0 {
		return rl.Vector2{X: 0, Y: 0}
	}

	avgPos := rl.Vector2{X: 0, Y: 0}
	for i := range visableBoids {
		visable := visableBoids[i]

		avgPos = avgPos.Add(visable.posVec)
	}

	avgPos.X /= float32(len(visableBoids))
	avgPos.Y /= float32(len(visableBoids))

	thisToAvgPos := avgPos.Subtract(b.posVec).Normalize()
	steeringVec := thisToAvgPos.Subtract(b.velocityVec)

	steeringVec.X *= cohesionSteeringMult
	steeringVec.Y *= cohesionSteeringMult

	return steeringVec
}

func applyAllRules(boids []boid) {
	//! fun thing about go, a for range loop that refers to a value, so: for _, v:= range ...
	//! actually makes a copy (of v) and doesn't mutate the refernces, keep that in mind
	steeringForEachBoid := make([]rl.Vector2, len(boids))

	for i := range boids {
		b := &boids[i]
		steeringForEachBoid[i] = b.applySeperationToBoid(boids[i].getAllboidsInTooCloseRange(boids))
		steeringForEachBoid[i] = steeringForEachBoid[i].Add(b.applyAlignmentToBoid(boids[i].getAllboidsInSightRange(boids)))
		steeringForEachBoid[i] = steeringForEachBoid[i].Add(b.applyCohesionToBoid(boids[i].getAllboidsInSightRange(boids)))
	}

	for i := range boids {
		b := &boids[i]
		b.velocityVec = steeringForEachBoid[i].Add(b.velocityVec).Normalize()
	}
}

//this should be only in rewrite
