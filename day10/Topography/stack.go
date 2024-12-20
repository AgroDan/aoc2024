package topmap

import "utils"

// This will "infer" from the Queue in the utils package
// and overload it to ensure that it only accepts
// Coordinate structs.

type CoordStack struct {
	stack utils.Stack
}

func NewCoordStack() CoordStack {
	return CoordStack{
		stack: utils.NewStack(),
	}
}

func (s *CoordStack) Push(item utils.Coord) {
	s.stack.Push(item)
}

func (s *CoordStack) Pop() (utils.Coord, bool) {
	element := s.stack.Pop()
	if element == nil {
		return utils.Coord{}, false
	}

	// perform type assertion
	myCoord, ok := element.(utils.Coord)
	return myCoord, ok
}

func (s *CoordStack) Peek() (utils.Coord, bool) {
	element := s.stack.Peek()
	if element == nil {
		return utils.Coord{}, false
	}

	// type assertion again
	myCoord, ok := element.(utils.Coord)
	return myCoord, ok
}

func (s *CoordStack) IsEmpty() bool {
	return s.stack.IsEmpty()
}
