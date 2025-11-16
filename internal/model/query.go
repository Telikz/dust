package model

// Query represents a query for entities with specific components
type Query struct {
	world      *World
	predicates []func(Particle) bool
}

// Query creates a new query builder
func (world *World) Query() *Query {
	return &Query{
		world:      world,
		predicates: []func(Particle) bool{},
	}
}

// With adds a requirement for a component type to the query
func (q *Query) With(store *ComponentStore[any]) *Query {
	q.predicates = append(q.predicates, func(p Particle) bool {
		return store.Has(p)
	})
	return q
}

// Execute returns all particles matching the query predicates
func (q *Query) Execute() []Particle {
	q.world.mu.RLock()
	defer q.world.mu.RUnlock()

	var results []Particle
	for particle := range q.world.Particles {
		matches := true
		for _, predicate := range q.predicates {
			if !predicate(particle) {
				matches = false
				break
			}
		}
		if matches {
			results = append(results, particle)
		}
	}
	return results
}

// ForEach executes a function for each particle matching the query
func (q *Query) ForEach(fn func(Particle)) {
	for _, particle := range q.Execute() {
		fn(particle)
	}
}
