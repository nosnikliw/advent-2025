package sets

type Set[T comparable] struct {
	members map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		members: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(val T) {
	s.members[val] = struct{}{}
}

func (s *Set[T]) Remove(val T) {
	delete(s.members, val)
}

func (s *Set[T]) Count() int {
	return len(s.members)
}

func (s *Set[T]) Contains(val T) bool {
	_, found := s.members[val]
	return found
}

func (s *Set[T]) Members() (list []T) {
	for i := range s.members {
		list = append(list, i)
	}
	return
}

func (s *Set[T]) Intersects(other Set[T]) bool {
	for member := range other.members {
		if s.Contains(member) {
			return true
		}
	}
	return false
}

func (s *Set[T]) Union(other Set[T]) *Set[T] {
	u := NewSet[T]()
	for v := range s.members {
		u.Add(v)
	}
	for v := range other.members {
		u.Add(v)
	}
	return u
}

func (s *Set[T]) Intersection(other Set[T]) *Set[T] {
	n := NewSet[T]()
	for v := range s.members {
		if other.Contains(v) {
			n.Add(v)
		}
	}
	return n
}
