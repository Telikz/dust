package model

type Component any

type Physics struct {
	Mass       float64
	Friction   float64
	Elasticity float64
}

type Behavior struct {
	CanBurn      bool
	CanFreeze    bool
	CanMelt      bool
	CanEvaporate bool
	CanFlow      bool
	Density      float64
}

type Material struct {
	Name string
}

type Position struct {
	X, Y float64
}

type Velocity struct {
	VX, VY float64
}

type Color struct {
	R, G, B, A float64
}

type Size struct {
	Radius float64
}
