package model

// Position represents a particle's location in 2D space
type Position struct {
	X, Y float64
}

// Velocity represents a particle's velocity in 2D space
type Velocity struct {
	VX, VY float64
}

// Lifetime represents how long a particle has existed and its max lifetime
type Lifetime struct {
	Current float64 // Current age in seconds
	Max     float64 // Max lifetime in seconds
}

// IsAlive returns true if the particle hasn't exceeded its max lifetime
func (lt *Lifetime) IsAlive() bool {
	return lt.Current < lt.Max
}

// Age increments the particle's age
func (lt *Lifetime) Age(dt float64) {
	lt.Current += dt
}

// Alpha returns a value from 1.0 (birth) to 0.0 (death) for fade-out
func (lt *Lifetime) Alpha() float64 {
	if lt.Max == 0 {
		return 1.0
	}
	return 1.0 - (lt.Current / lt.Max)
}

// Color represents RGBA color
type Color struct {
	R, G, B, A float64
}

// Acceleration represents constant acceleration applied to velocity
type Acceleration struct {
	AX, AY float64
}

// Size represents particle size/radius
type Size struct {
	Radius float64
}
