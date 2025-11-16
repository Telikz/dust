// Package systems implements various systems for the ECS architecture.
package systems

import "github.com/telikz/dust/internal/model"

type GravitySystem struct {
	Gravity float64
}

func (gs *GravitySystem) Update(world *model.World, deltaTime float64) {
	for particle := range world.Particles {
		components := world.Components[particle]
		if velComp, ok := components["Velocity"].(*model.Velocity); ok {
			velComp.VY += gs.Gravity * deltaTime
		}
	}
}
