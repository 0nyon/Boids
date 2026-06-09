package main

import (
	"math/rand"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ! all constants here
const boidAmnt uint = 150
const protectedRange uint = 10
const visableRange uint = 60
const avoidFactor float32 = 0.4
const matchingFactor float32 = 0.2
const centeringFactor float32 = 0.007
const maxSpeed float32 = 4
const minSpeed float32 = maxSpeed - 3
const turnFactor float32 = 0.3
const borderPadding uint = 150

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

func (b *boid) getRotation() float32 {
	//return 360 * b.velocityVec.Angle(rl.Vector2{X: 1, Y: 0})
	return 0
}

func (b *boid) getAllBoidsInProtectedRange(others []boid) []boid {
	len := len(others)
	res := make([]boid, 0)
	for i := 0; i < len; i++ {
		other := &others[i]

		if *b == *other {
			continue
		}

		if b.posVec.Distance(other.posVec) <= float32(protectedRange) {
			res = append(res, *other)
		}
	}
	return res
}

func (b *boid) getAllBoidsInVisableRange(others []boid) []boid {

	len := len(others)
	res := make([]boid, 0)
	for i := 0; i < len; i++ {
		other := &others[i]

		if *b == *other {
			continue
		}

		if b.posVec.Distance(other.posVec) <= float32(visableRange) {
			res = append(res, *other)
		}
	}
	return res
}

func (b *boid) getAllBoidsInVisableRangeEx(others []boid) []boid {
	allVis := b.getAllBoidsInVisableRange(others)
	allProt := b.getAllBoidsInProtectedRange(others)
	res := make([]boid, 0)
	for _, visable := range allVis {
		if slices.Contains(allProt, visable) {
			continue
		}
		res = append(res, visable)
	}
	return res
}

func (b *boid) getSeperationVector(inProtRange []boid) rl.Vector2 {
	len := len(inProtRange)
	res := rl.Vector2{X: 0, Y: 0} //* this would be the vector for summing the distance from "b" to all others in prot range

	for i := 0; i < len; i++ {
		res = res.Add(b.posVec.Subtract(inProtRange[i].posVec))
	}

	res.X = res.X * avoidFactor
	res.Y = res.Y * avoidFactor

	//* this result should be added to the boid's velocity vector
	return res
}

func (b *boid) getAlignmentVector(inExVisable []boid) rl.Vector2 {
	len := len(inExVisable)
	if len == 0 {
		return rl.Vector2{X: 0, Y: 0}
	}

	avgVel := rl.Vector2{X: 0, Y: 0} //* for finding avg velocity
	for i := 0; i < len; i++ {
		avgVel = avgVel.Add(inExVisable[i].velocityVec)
	}

	avgVel.X /= float32(len)
	avgVel.Y /= float32(len)

	res := avgVel.Subtract(b.velocityVec)
	res.X *= matchingFactor
	res.Y *= matchingFactor

	//* this result should be added to the boid's velocity vector
	return res
}

func (b *boid) getCohesionVector(inExVisable []boid) rl.Vector2 {
	len := len(inExVisable)
	if len == 0 {
		return rl.Vector2{X: 0, Y: 0}
	}

	avgPos := rl.Vector2{X: 0, Y: 0} //* for finding avg position
	for i := 0; i < len; i++ {
		avgPos = avgPos.Add(inExVisable[i].posVec)
	}

	avgPos.X /= float32(len)
	avgPos.Y /= float32(len)

	res := avgPos.Subtract(b.posVec)
	res.X *= centeringFactor
	res.Y *= centeringFactor

	//* this result should be added to the boid's velocity vector
	return res
}

func (b *boid) applySpeedRule() {
	//force the boids to be faster than minSpeed and slower than maxSpeed
	speed := b.velocityVec.Length()

	if speed > maxSpeed {
		b.velocityVec.X = (float32(b.velocityVec.X) / speed) * maxSpeed
		b.velocityVec.Y = (float32(b.velocityVec.Y) / speed) * maxSpeed
	} else if speed < minSpeed {
		b.velocityVec.X = (float32(b.velocityVec.X) / speed) * minSpeed
		b.velocityVec.Y = (float32(b.velocityVec.Y) / speed) * minSpeed
	}

}

func (b *boid) applyScreenEdgeRule() {
	//force the boids to be faster than minSpeed and slower than maxSpeed
	var top float32 = float32(borderPadding)
	var bottom float32 = float32(screenSize) - float32(borderPadding)
	var left float32 = float32(borderPadding)
	var right float32 = float32(screenSize) - float32(borderPadding)

	if b.posVec.X < left {
		b.velocityVec.X += turnFactor
	} else if b.posVec.X > right {
		b.velocityVec.X -= turnFactor
	}

	if b.posVec.Y < top {
		b.velocityVec.Y += turnFactor
	} else if b.posVec.Y > bottom {
		b.velocityVec.Y -= turnFactor
	}
}

func applyAllRules(all []boid) {
	len := len(all)
	vecsToAdd := make([]rl.Vector2, len)
	for i := 0; i < len; i++ {
		b := &all[i]
		vecToAdd := rl.Vector2{X: 0, Y: 0}

		vecToAdd = b.velocityVec.Add(b.getSeperationVector(b.getAllBoidsInProtectedRange(all)))

		vecToAdd = vecToAdd.Add(b.getAlignmentVector(b.getAllBoidsInVisableRangeEx(all)))

		vecToAdd = vecToAdd.Add((b.getCohesionVector(b.getAllBoidsInVisableRangeEx(all))))

		vecsToAdd[i] = vecToAdd
	}

	for i := 0; i < len; i++ {
		all[i].velocityVec = all[i].velocityVec.Add(vecsToAdd[i])
		all[i].posVec = all[i].posVec.Add(all[i].velocityVec)
		all[i].applySpeedRule()
		all[i].applyScreenEdgeRule()
	}
}
