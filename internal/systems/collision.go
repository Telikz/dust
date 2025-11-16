package systems

import (
	"math"

	"github.com/telikz/dust/internal/model"
)

type CollisionSystem struct{}

func NewCollisionSystem(cellSize float64) *CollisionSystem {
	return &CollisionSystem{}
}

func (cs *CollisionSystem) Update(world *model.World, deltaTime float64) {
	particles := make([]model.Particle, 0, len(world.Particles))
	for p := range world.Particles {
		particles = append(particles, p)
	}

	// Check particle-to-particle collisions - just separate, no bouncing
	for i := 0; i < len(particles); i++ {
		for j := i + 1; j < len(particles); j++ {
			cs.checkParticleCollision(world, particles[i], particles[j])
		}
	}

	// Check floor and boundary collisions
	for _, p := range particles {
		cs.checkBoundaryCollision(world, p, deltaTime)
	}
}

func (cs *CollisionSystem) checkParticleCollision(world *model.World, p1, p2 model.Particle) {
	comps1 := world.Components[p1]
	comps2 := world.Components[p2]

	pos1, ok1 := comps1["Position"].(*model.Position)
	pos2, ok2 := comps2["Position"].(*model.Position)
	size1, ok5 := comps1["Size"].(*model.Size)
	size2, ok6 := comps2["Size"].(*model.Size)

	if !ok1 || !ok2 || !ok5 || !ok6 {
		return
	}

	dx := pos2.X - pos1.X
	dy := pos2.Y - pos1.Y
	distSq := dx*dx + dy*dy
	minDist := size1.Radius + size2.Radius
	minDistSq := minDist * minDist

	// If overlapping, just push apart
	if distSq < minDistSq && distSq > 0.01 {
		dist := math.Sqrt(distSq)
		nx := dx / dist
		ny := dy / dist

		// Push apart equally
		overlap := minDist - dist
		push := overlap * 0.5
		pos1.X -= nx * push
		pos1.Y -= ny * push
		pos2.X += nx * push
		pos2.Y += ny * push
	}
}

func (cs *CollisionSystem) checkBoundaryCollision(world *model.World, particle model.Particle, deltaTime float64) {
	components := world.Components[particle]

	pos, okPos := components["Position"].(*model.Position)
	vel, okVel := components["Velocity"].(*model.Velocity)
	phys, okPhys := components["Physics"].(*model.Physics)

	if !okPos || !okVel || !okPhys {
		return
	}

	// Floor
	if pos.Y >= world.FloorY {
		pos.Y = world.FloorY
		vel.VY = -vel.VY * phys.Elasticity
		vel.VX *= (1.0 - phys.Friction*deltaTime)
		if vel.VY < 0.05 {
			vel.VY = 0
		}
	}

	// Ceiling
	if pos.Y < 0 {
		pos.Y = 0
		vel.VY = -vel.VY * phys.Elasticity
	}

	// Sides
	if pos.X < 0 {
		pos.X = 0
		vel.VX = -vel.VX * phys.Elasticity
	}
	if pos.X >= float64(world.Map.Width) {
		pos.X = float64(world.Map.Width - 1)
		vel.VX = -vel.VX * phys.Elasticity
	}
}
