package main

type Stack struct {
	data []int
}
func newstack() *Stack {
	return &Stack{data: make([]int, 0, 1024)}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(num int) {
	s.data = append(s.data, num)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(s.data) - 1
		element := (s.data)[index]
		s.data = (s.data)[:index]
		return element, true
	}
}
