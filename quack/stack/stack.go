package stack

import "errors"

type Stack struct {
	stk []int
}

type StackInterface interface {
	Push(x int)
	Pop() (int, error)
	IsEmpty() bool
}

func (s *Stack) Push(x int) {
	_ = append(s.stk, x)
}

func (s *Stack) Pop() (int, error) {
	if !s.IsEmpty() {
		v := s.stk[len(s.stk)-1]
		if len(s.stk) > 1 {
			s.stk = s.stk[:len(s.stk)-2]
		} else {
			s.stk = nil
		}
		return v, nil
	}
	return 0, errors.New("Stack is empty")
}

func (s *Stack) IsEmpty() {
	return len(s.stk) == 0
}
