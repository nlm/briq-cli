package utils

// Set is an object in which you can store items
// and efficiently check if an item is contained in it.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet returns a new Set with the provided items.
func NewSet[T comparable](items ...T) *Set[T] {
	s := Set[T]{}
	s.Add(items...)
	return &s
}

// Add adds an item to the set if it's not already present.
func (s *Set[T]) Add(items ...T) {
	if s == nil {
		s = &Set[T]{}
	}
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

// Remove removes an item from the set if it's present.
func (s *Set[T]) Remove(items ...T) {
	if s == nil || s.m == nil {
		return
	}
	for _, item := range items {
		delete(s.m, item)
	}
}

// Contains indicates if an item is present in the Set.
func (s *Set[T]) Contains(item T) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[item]
	return ok
}

// Len returns the number of items in the Set.
func (s *Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.m)
}

// Items returns a slice with copies of the items of the Set.
func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.m))
	for k := range s.m {
		items = append(items, k)
	}
	return items
}
