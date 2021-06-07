package nstack

import "fmt"

var ErrIsEmpty = fmt.Errorf("stack is empty")

type Stack struct {
	n        int
	elements []Element
}

type Element struct {
	stackId int
	value   interface{}
}

func New(n int) *Stack {
	s := &Stack{
		n: n,
	}

	return s
}

func (s *Stack) Push(stackId int, value interface{}) error {
	s.elements = append(s.elements, Element{stackId: stackId, value: value})

	return nil
}

func (s *Stack) Pop(stackId int) (interface{}, error) {
	value, i, err := s.peek(stackId)
	if err != nil {
		return nil, err
	}

	s.elements = append(s.elements[:i], s.elements[i+1:]...)
	return value, nil
}

func (s *Stack) peek(stackId int) (interface{}, int, error) {
	i, err := s.findLast(stackId)
	if err != nil {
		return nil, 0, err
	}

	value := s.elements[i].value
	return value, i, nil
}

func (s *Stack) Peek(stackId int) (interface{}, error) {
	value, _, err := s.peek(stackId)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (s *Stack) findLast(stackId int) (int, error) {
	for i := len(s.elements) - 1; i >= 0; i-- {
		if s.elements[i].stackId == stackId {
			return i, nil
		}
	}
	return 0, ErrIsEmpty
}
