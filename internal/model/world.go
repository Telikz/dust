package model

import "sync"

type Map struct {
	height, width int
}

type World struct {
	Map              Map
	Particles        map[Particle]bool
	nextParticleID   Particle
	mu               sync.RWMutex
	componentStores  map[string]any
	componentStoreMu sync.RWMutex
}

func NewWorld() *World {
	return &World{
		Particles:       make(map[Particle]bool),
		nextParticleID:  1,
		componentStores: make(map[string]any),
	}
}

func (world *World) GetPositions() *ComponentStore[Position] {
	world.componentStoreMu.Lock()
	defer world.componentStoreMu.Unlock()

	const typeName = "Position"
	if store, exists := world.componentStores[typeName]; exists {
		return store.(*ComponentStore[Position])
	}

	store := NewComponentStore[Position]()
	world.componentStores[typeName] = store
	return store
}

func (world *World) GetVelocities() *ComponentStore[Velocity] {
	world.componentStoreMu.Lock()
	defer world.componentStoreMu.Unlock()

	const typeName = "Velocity"
	if store, exists := world.componentStores[typeName]; exists {
		return store.(*ComponentStore[Velocity])
	}

	store := NewComponentStore[Velocity]()
	world.componentStores[typeName] = store
	return store
}

func (world *World) GetAccelerations() *ComponentStore[Acceleration] {
	world.componentStoreMu.Lock()
	defer world.componentStoreMu.Unlock()

	const typeName = "Acceleration"
	if store, exists := world.componentStores[typeName]; exists {
		return store.(*ComponentStore[Acceleration])
	}

	store := NewComponentStore[Acceleration]()
	world.componentStores[typeName] = store
	return store
}

func (world *World) GetLifetimes() *ComponentStore[Lifetime] {
	world.componentStoreMu.Lock()
	defer world.componentStoreMu.Unlock()

	const typeName = "Lifetime"
	if store, exists := world.componentStores[typeName]; exists {
		return store.(*ComponentStore[Lifetime])
	}

	store := NewComponentStore[Lifetime]()
	world.componentStores[typeName] = store
	return store
}

func (world *World) GetColors() *ComponentStore[Color] {
	world.componentStoreMu.Lock()
	defer world.componentStoreMu.Unlock()

	const typeName = "Color"
	if store, exists := world.componentStores[typeName]; exists {
		return store.(*ComponentStore[Color])
	}

	store := NewComponentStore[Color]()
	world.componentStores[typeName] = store
	return store
}

func (world *World) GetSizes() *ComponentStore[Size] {
	world.componentStoreMu.Lock()
	defer world.componentStoreMu.Unlock()

	const typeName = "Size"
	if store, exists := world.componentStores[typeName]; exists {
		return store.(*ComponentStore[Size])
	}

	store := NewComponentStore[Size]()
	world.componentStores[typeName] = store
	return store
}
