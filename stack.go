package stack

type Stack struct {
	size int
	val  []interface{}
}

func (s *Stack) Init() {
	s.size = 0
	s.val = make([]interface{}, 0)
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(val interface{}) {
	s.val = append(s.val, val)
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	val := s.val[s.size-1]
	s.val = s.val[:s.size-1]
	s.size--
	return val
}

func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.val[s.size-1]
}

func (s *Stack) Empty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s *Stack) Traverse() []interface{} {
	valList := make([]interface{}, s.size)
	for i := 0; i < s.size; i++ {
		valList[i] = s.val[i]
	}
	return valList
}
