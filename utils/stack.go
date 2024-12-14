package utils

// a FIFO stack

// This will create a stack that holds _anything_.

// this is ghetto
// type Stack []interface{}

// func (s *Stack) Push(item interface{}) {
// 	*s = append(*s, item)
// }

// func (s *Stack) Pop() interface{} {
// 	if len(*s) == 0 {
// 		return nil
// 	}

// 	top := (*s)[len(*s)-1]
// 	*s = (*s)[:len(*s)-1]
// 	return top
// }

// func (s *Stack) Peek() interface{} {
// 	return (*s)[len(*s)-1]
// }

type Stack struct {
	elements []interface{} // slice to hold stack elements
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack) Peek() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
