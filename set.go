package set

func Id[T comparable](x T) T {
	return x
}

// A Set implementation. T is the element type, U is the hash type.
type Set[U comparable, T any] struct {
	Elems  map[U]T
	hasher func(T) U
}

// Add element to the set
func NewSet[U comparable, T any](elems []T, hasher func(T) U) Set[U, T] {
	s := Set[U, T]{Elems: make(map[U]T), hasher: hasher}
	for _, elem := range elems {
		s.Add(elem)
	}
	return s
}

func NewEmptySet[U comparable, T any](hasher func(T) U) Set[U, T] {
	return NewSet([]T{}, hasher)
}

// NewSimpleSet creates a new Set where the hash function is the identity function.
func NewSimpleSet[U comparable]() Set[U, U] {
	return NewSet([]U{}, Id[U])
}

func (s Set[U, T]) Len() int {
	return len(s.Elems)
}

// Add element to the set
func (s Set[U, T]) Add(elem T) {
	s.Elems[s.hasher(elem)] = elem
}

// Remove element from the set
func (s Set[U, T]) Remove(elem T) {
	delete(s.Elems, s.hasher(elem))
}

// Check element is in the set
func (s Set[U, T]) Has(elem T) bool {
	_, ok := s.Elems[s.hasher(elem)]
	return ok
}

// Convert the set to a slice
func (s Set[U, T]) ToList() []T {
	list := []T{}
	for _, val := range s.Elems {
		list = append(list, val)
	}
	return list
}

// Creates a copy of the set
func (s Set[U, T]) Clone() Set[U, T] {
	clone := NewEmptySet(s.hasher)
	for k, v := range s.Elems {
		clone.Elems[k] = v
	}
	return clone
}

// Update the set by taking the union with the other set
func (s Set[U, T]) Union(other Set[U, T]) {
	for _, v := range other.Elems {
		s.Add(v)
	}
}

// Compute the union of s1 and s2
func Union[U comparable, T any](s1, s2 Set[U, T]) Set[U, T] {
	union := s1.Clone()
	for _, v := range s2.Elems {
		union.Add(v)
	}
	return union
}

// Update the set by taking the intersection with the other set
func (s Set[U, T]) Intersect(other Set[U, T]) {
	for k, v := range s.Elems {
		_, ok := other.Elems[k]
		if !ok {
			s.Remove(v)
		}
	}
}

// Compute the intersection of s1 and s2
func Intersect[U comparable, T any](s1, s2 Set[U, T]) Set[U, T] {
	intersection := NewEmptySet(s1.hasher)
	for k, v := range s1.Elems {
		_, ok := s2.Elems[k]
		if ok {
			intersection.Add(v)
		}
	}
	return intersection
}

// Update the set by taking the difference with the other set
func (s Set[U, T]) Difference(other Set[U, T]) {
	for k, v := range s.Elems {
		_, ok := other.Elems[k]
		if ok {
			s.Remove(v)
		}
	}
}

// Compute the difference of s1 from s2
func Difference[U comparable, T any](s1, s2 Set[U, T]) Set[U, T] {
	diff := NewEmptySet(s1.hasher)
	for k, v := range s1.Elems {
		_, ok := s2.Elems[k]
		if !ok {
			diff.Add(v)
		}
	}
	return diff
}

// Checks whether the set is a subset of the other set
func (s Set[U, T]) IsSubset(other Set[U, T]) bool {
	return Intersect(s, other).Len() == s.Len()
}

// Checks whether s1 is a subset of s2
func IsSubset[U comparable, T any](s1, s2 Set[U, T]) bool {
	return s1.IsSubset(s2)
}

// Checks whether the set is a superset of the other set
func (s Set[U, T]) IsSuperset(other Set[U, T]) bool {
	return other.IsSubset(s)
}

// Checks whether s1 is a superset of s2
func IsSuperset[U comparable, T any](s1, s2 Set[U, T]) bool {
	return s1.IsSuperset(s2)
}

// Checks whether the set has no intersection with the other set
func (s Set[U, T]) IsDisjoint(other Set[U, T]) bool {
	return Intersect(s, other).Len() == 0
}

// Checks whether s1 and s2 have no intersection
func IsDisjoint[U comparable, T any](s1, s2 Set[U, T]) bool {
	return s1.IsDisjoint(s2)
}

// Checks whether s contains the same elements as the other
func (s Set[U, T]) Equal(other Set[U, T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for k := range s.Elems {
		_, ok := other.Elems[k]
		if !ok {
			return false
		}
	}

	return true
}

// Checks whether s1 and s2 contain the same elements
func Equal[U comparable, T any](s1, s2 Set[U, T]) bool {
	return s1.Equal(s2)
}
