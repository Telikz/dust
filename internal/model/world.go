// Package model contains the core data structures.
package model

type Map struct {
	Height, Width int
}

type World struct {
	FloorY     float64
	Map        Map
	Particles  map[Particle]bool
	Components map[Particle]map[string]Component
}

func NewWorld() *World {
	return &World{
		FloorY:     0,
		Map:        Map{Height: 100, Width: 100},
		Particles:  make(map[Particle]bool),
		Components: make(map[Particle]map[string]Component),
	}
}

func (world *World) Update(systems []System, deltaTime float64) {
	for _, system := range systems {
		system.Update(world, deltaTime)
	}
}
