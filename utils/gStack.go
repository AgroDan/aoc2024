package utils

// Just like the gQueue, I'm going to create a stack that takes
// generic datatypes without using the interface{} datatype.
// This allows me to add whatever data I want without creating
// a wrapper function in the stack

type GStack[T any] struct {
	element []T
}

func NewGStack[T any]() GStack[T] {
	return GStack[T]{}
}

func (s *GStack[T]) Push(element T) {
	// Pushes an element onto the generic stack
	s.element = append(s.element, element)
}

func (s *GStack[T]) Pop() (T, bool) {
	if len(s.element) == 0 {
		var zero T
		return zero, false
	}

	// otherwise pop the stack
	top := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return top, true
}

func (s *GStack[T]) Peek() (T, bool) {
	// peeks at the top of the stack,
	// doesn't manipulate otherwise
	if len(s.element) == 0 {
		var zero T
		return zero, false
	}

	return s.element[len(s.element)-1], true
}

func (s *GStack[T]) IsEmpty() bool {
	return len(s.element) == 0
}
