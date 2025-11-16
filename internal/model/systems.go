package model

// System interface represents a system that operates on particles
type System interface {
	Update(world *World, dt float64)
}

// MovementSystem updates particle positions based on velocity and acceleration
type MovementSystem struct {
	positions     *ComponentStore[Position]
	velocities    *ComponentStore[Velocity]
	accelerations *ComponentStore[Acceleration]
}

func NewMovementSystem(w *World) *MovementSystem {
	return &MovementSystem{
		positions:     w.GetPositions(),
		velocities:    w.GetVelocities(),
		accelerations: w.GetAccelerations(),
	}
}

func (ms *MovementSystem) Update(world *World, dt float64) {
	positions := ms.positions.All()
	velocities := ms.velocities.All()
	accelerations := ms.accelerations.All()

	for particle, pos := range positions {
		vel, hasVel := velocities[particle]
		if !hasVel {
			continue
		}

		// Apply acceleration if present
		if accel, hasAccel := accelerations[particle]; hasAccel {
			vel.VX += accel.AX * dt
			vel.VY += accel.AY * dt
			ms.velocities.Set(particle, vel)
		}

		// Update position
		pos.X += vel.VX * dt
		pos.Y += vel.VY * dt
		ms.positions.Set(particle, pos)
	}
}

// LifetimeSystem ages particles and removes dead ones
type LifetimeSystem struct {
	lifetimes *ComponentStore[Lifetime]
}

func NewLifetimeSystem(w *World) *LifetimeSystem {
	return &LifetimeSystem{
		lifetimes: w.GetLifetimes(),
	}
}

func (ls *LifetimeSystem) Update(world *World, dt float64) {
	lifetimes := ls.lifetimes.All()

	for particle, lifetime := range lifetimes {
		lifetime.Age(dt)

		if !lifetime.IsAlive() {
			world.RemoveParticle(particle)
			ls.lifetimes.Remove(particle)
		} else {
			ls.lifetimes.Set(particle, lifetime)
		}
	}
}
