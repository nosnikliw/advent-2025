package sets

type IntSet struct {
	members map[int64]struct{}
}

func NewIntSet() *IntSet {
	return &IntSet{
		members: make(map[int64]struct{}),
	}
}

func (s *IntSet) Add(val int64) {
	s.members[val] = struct{}{}
}

func (s *IntSet) AddRange(min int64, max int64) {
	for i := min; i <= max; i++ {
		s.Add(i)
	}
}

func (s *IntSet) Remove(val int64) {
	delete(s.members, val)
}

func (s *IntSet) Count() int {
	return len(s.members)
}

func (s *IntSet) Contains(val int64) bool {
	_, found := s.members[val]
	return found
}

func (s *IntSet) Members() (list []int64) {
	for i := range s.members {
		list = append(list, i)
	}
	return
}
