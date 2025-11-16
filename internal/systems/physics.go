package systems

import "github.com/telikz/dust/internal/model"

type PhysicsSystem struct{}

func (ps *PhysicsSystem) Update(world *model.World, deltaTime float64) {
	for particle := range world.Particles {
		components := world.Components[particle]

		if posComp, ok := components["Position"].(*model.Position); ok {
			if velComp, ok := components["Velocity"].(*model.Velocity); ok {
				posComp.X += velComp.VX * deltaTime
				posComp.Y += velComp.VY * deltaTime
			}
		}
	}
}
