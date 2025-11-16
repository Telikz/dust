package model

import (
	"math"
	"math/rand"
)

// Emitter creates particles with configurable properties
type Emitter struct {
	PositionX    float64 // Emitter X position
	PositionY    float64 // Emitter Y position
	SpreadRadius float64 // Radius to spread particles around emitter
	VelocityMin  float64 // Minimum velocity magnitude
	VelocityMax  float64 // Maximum velocity magnitude
	AngleMin     float64 // Minimum angle in radians
	AngleMax     float64 // Maximum angle in radians
	Lifetime     float64 // Particle lifetime in seconds
	Color        Color   // Particle color
	Size         float64 // Particle size/radius
	Acceleration Acceleration
	WithAccel    bool // Whether to apply acceleration
}

// NewEmitter creates a default emitter
func NewEmitter(x, y float64) *Emitter {
	return &Emitter{
		PositionX:    x,
		PositionY:    y,
		SpreadRadius: 1.0,
		VelocityMin:  10.0,
		VelocityMax:  50.0,
		AngleMin:     0,
		AngleMax:     2 * math.Pi,
		Lifetime:     2.0,
		Color:        Color{R: 1.0, G: 0.5, B: 0.2, A: 1.0},
		Size:         2.0,
		Acceleration: Acceleration{AX: 0, AY: -9.8}, // Gravity
		WithAccel:    true,
	}
}

// Emit creates a single particle with randomized properties
func (e *Emitter) Emit(world *World) Particle {
	// Create particle
	p := world.AddParticle()

	// Position with spread
	angle := rand.Float64() * 2 * math.Pi
	radius := rand.Float64() * e.SpreadRadius
	x := e.PositionX + radius*math.Cos(angle)
	y := e.PositionY + radius*math.Sin(angle)
	world.GetPositions().Set(p, Position{X: x, Y: y})

	// Velocity
	velocityMag := e.VelocityMin + rand.Float64()*(e.VelocityMax-e.VelocityMin)
	velocityAngle := e.AngleMin + rand.Float64()*(e.AngleMax-e.AngleMin)
	vx := velocityMag * math.Cos(velocityAngle)
	vy := velocityMag * math.Sin(velocityAngle)
	world.GetVelocities().Set(p, Velocity{VX: vx, VY: vy})

	// Lifetime
	world.GetLifetimes().Set(p, Lifetime{Current: 0, Max: e.Lifetime})

	// Color
	world.GetColors().Set(p, e.Color)

	// Size
	world.GetSizes().Set(p, Size{Radius: e.Size})

	// Acceleration (if enabled)
	if e.WithAccel {
		world.GetAccelerations().Set(p, e.Acceleration)
	}

	return p
}

// EmitBurst creates multiple particles at once
func (e *Emitter) EmitBurst(world *World, count int) []Particle {
	particles := make([]Particle, count)
	for i := 0; i < count; i++ {
		particles[i] = e.Emit(world)
	}
	return particles
}
