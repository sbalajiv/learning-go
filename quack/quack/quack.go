package quack

import (
	"errors"
	"stack"
)

type QuackInteface interface {
	Push(x int)
	Pop() (int, error)
	Pull() (int, error)
}

type Quack struct {
	S1   stack.Stack
	S2   stack.Stack
	S3   stack.Stack
	nele int
}

func (q Quack) Push(x int) {
	q.S1.Push(x)
	q.S2.Push(x)
	q.nele++
}

func (q Quack) Pop() (int, error) {
	v := 0

	if q.nele > 0 {
		v = q.S1.Pop()
		q.nele--
		return v, nil
	}
	return v, errors.New("Quack is empty")
}
