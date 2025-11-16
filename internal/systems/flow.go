package systems

import "github.com/telikz/dust/internal/model"

type FlowSystem struct{}

func (fs *FlowSystem) Update(world *model.World, deltaTime float64) {
	particles := make([]model.Particle, 0, len(world.Particles))
	for p := range world.Particles {
		particles = append(particles, p)
	}

	// Apply flow behavior to fluids
	for _, p := range particles {
		fs.applyFlow(world, p, particles)
	}
}

func (fs *FlowSystem) applyFlow(world *model.World, particle model.Particle, allParticles []model.Particle) {
	comps := world.Components[particle]

	behavior, okBehavior := comps["Behavior"].(*model.Behavior)
	if !okBehavior || !behavior.CanFlow {
		return // Only apply to flowing particles
	}

	pos, okPos := comps["Position"].(*model.Position)
	vel, okVel := comps["Velocity"].(*model.Velocity)

	if !okPos || !okVel {
		return
	}

	// Make fluids spread sideways when stacked
	for _, other := range allParticles {
		if other == particle {
			continue
		}

		otherComps := world.Components[other]
		otherBehavior, ok := otherComps["Behavior"].(*model.Behavior)
		if !ok || !otherBehavior.CanFlow {
			continue
		}

		otherPos, ok := otherComps["Position"].(*model.Position)
		if !ok {
			continue
		}

		otherVel, ok := otherComps["Velocity"].(*model.Velocity)
		if !ok {
			continue
		}

		// If very close (stacked)
		dx := otherPos.X - pos.X
		dy := otherPos.Y - pos.Y

		// Particles touching vertically or nearly so
		if dy > -0.5 && dy < 0.5 && (dx > -1.5 && dx < 1.5) && dx*dx > 0.1 {
			// Push sideways to spread
			pushForce := 0.5
			if dx > 0 {
				vel.VX -= pushForce
				otherVel.VX += pushForce
			} else {
				vel.VX += pushForce
				otherVel.VX -= pushForce
			}
		}
	}

	// Dampen velocity for fluids (they settle faster)
	vel.VX *= 0.95
	vel.VY *= 0.98
}
