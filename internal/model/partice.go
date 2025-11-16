package model

import "sync"

type Particle uint64

// ComponentStore provides type-safe storage for a specific component type
type ComponentStore[T any] struct {
	data map[Particle]T
	mu   sync.RWMutex
}

func NewComponentStore[T any]() *ComponentStore[T] {
	return &ComponentStore[T]{
		data: make(map[Particle]T),
	}
}

func (cs *ComponentStore[T]) Set(particle Particle, component T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.data[particle] = component
}

func (cs *ComponentStore[T]) Get(particle Particle) (T, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	val, ok := cs.data[particle]
	return val, ok
}

func (cs *ComponentStore[T]) Has(particle Particle) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	_, ok := cs.data[particle]
	return ok
}

func (cs *ComponentStore[T]) Remove(particle Particle) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	delete(cs.data, particle)
}

func (cs *ComponentStore[T]) All() map[Particle]T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	// Return a copy to prevent external mutations
	result := make(map[Particle]T, len(cs.data))
	for k, v := range cs.data {
		result[k] = v
	}
	return result
}

func (world *World) AddParticle() Particle {
	world.mu.Lock()
	defer world.mu.Unlock()
	particle := world.nextParticleID
	world.nextParticleID++
	world.Particles[particle] = true
	return particle
}

func (world *World) RemoveParticle(particle Particle) {
	world.mu.Lock()
	defer world.mu.Unlock()
	delete(world.Particles, particle)
}

func (world *World) HasParticle(particle Particle) bool {
	world.mu.RLock()
	defer world.mu.RUnlock()
	_, exists := world.Particles[particle]
	return exists
}
