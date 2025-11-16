package model

type Particle uint64

var NextParticleID Particle = 1

func NewParticle() Particle {
	id := NextParticleID
	NextParticleID++
	return id
}

func CreateSand(world *World, x, y float64) Particle {
	particle := NewParticle()
	world.Particles[particle] = true
	world.Components[particle] = make(map[string]Component)
	world.Components[particle]["Position"] = &Position{X: x, Y: y}
	world.Components[particle]["Velocity"] = &Velocity{VX: 0, VY: 0}
	world.Components[particle]["Physics"] = &Physics{Mass: 1.0, Friction: 0.2, Elasticity: 0.1}
	world.Components[particle]["Behavior"] = &Behavior{CanFlow: false, CanBurn: false, Density: 2.5}
	world.Components[particle]["Material"] = &Material{Name: "Sand"}
	world.Components[particle]["Color"] = &Color{R: 0.76, G: 0.7, B: 0.5, A: 1.0}
	world.Components[particle]["Size"] = &Size{Radius: 0.3}
	return particle
}

func CreateWater(world *World, x, y float64) Particle {
	p := NewParticle()
	world.Particles[p] = true
	world.Components[p] = make(map[string]Component)
	world.Components[p]["Position"] = &Position{X: x, Y: y}
	world.Components[p]["Velocity"] = &Velocity{VX: 0, VY: 0}
	world.Components[p]["Physics"] = &Physics{Mass: 1.0, Friction: 0.1, Elasticity: 0.0}
	world.Components[p]["Behavior"] = &Behavior{CanFlow: true, CanBurn: false, Density: 1.0}
	world.Components[p]["Material"] = &Material{Name: "Water"}
	world.Components[p]["Color"] = &Color{R: 0.2, G: 0.6, B: 0.9, A: 1.0}
	world.Components[p]["Size"] = &Size{Radius: 0.4}
	return p
}

func CreateOil(world *World, x, y float64) Particle {
	p := NewParticle()
	world.Particles[p] = true
	world.Components[p] = make(map[string]Component)
	world.Components[p]["Position"] = &Position{X: x, Y: y}
	world.Components[p]["Velocity"] = &Velocity{VX: 0, VY: 0}
	world.Components[p]["Physics"] = &Physics{Mass: 0.8, Friction: 0.05, Elasticity: 0.0}
	world.Components[p]["Behavior"] = &Behavior{CanFlow: true, CanBurn: true, Density: 0.9}
	world.Components[p]["Material"] = &Material{Name: "Oil"}
	world.Components[p]["Color"] = &Color{R: 0.3, G: 0.3, B: 0.1, A: 1.0}
	world.Components[p]["Size"] = &Size{Radius: 0.45}
	return p
}
