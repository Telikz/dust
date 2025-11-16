package model

type System interface {
	Update(world *World, deltaTime float64)
}
