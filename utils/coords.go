package utils

// This has to do with coordinates.
// modified slightly to allow for 45 degree directions

const (
	N = iota
	E
	S
	W
	NE
	SE
	SW
	NW
)

type Coord struct {
	X, Y int
}

// the below is based on how data is read in, from the top left to the bottom right.
// So one line down means that Y is added to, one line up is Y is subtracted from.

func (c *Coord) Move(dir int) {
	// Changes the coord position
	// based on the implied direction
	switch dir {
	case N:
		c.Y--
	case E:
		c.X++
	case S:
		c.Y++
	case W:
		c.X--
	case NE:
		c.Y--
		c.X++
	case SE:
		c.Y++
		c.X++
	case SW:
		c.Y++
		c.X--
	case NW:
		c.Y--
		c.X--
	default:
		panic("invalid direction")
	}
}

func (c Coord) Peek(dir int) Coord {
	// Like moving, but returns a separate
	// coord object instead of modifying
	// the current struct object
	check := c
	switch dir {
	case N:
		check.Y--
	case E:
		check.X++
	case S:
		check.Y++
	case W:
		check.X--
	case NE:
		check.Y--
		check.X++
	case SE:
		check.Y++
		check.X++
	case SW:
		check.Y++
		check.X--
	case NW:
		check.Y--
		check.X--
	default:
		panic("invalid direction")
	}
	return check
}

func (c Coord) AllAvailable() []Coord {
	// this will give a list of all coordinates in all 4
	// directions from where we are.
	retval := make([]Coord, 4)
	retval[0] = c.Peek(N) // look north
	retval[1] = c.Peek(E) // look east
	retval[2] = c.Peek(S) // look south
	retval[3] = c.Peek(W) // look west
	return retval
}

func (c Coord) Neighbors() []Coord {
	// Similar to allavailable, but this considers diagonal as well
	retval := make([]Coord, 8)
	retval[0] = c.Peek(N)  // look north
	retval[1] = c.Peek(NE) // look northeast
	retval[2] = c.Peek(E)  // look east
	retval[3] = c.Peek(SE) // look southeast
	retval[4] = c.Peek(S)  // look south
	retval[5] = c.Peek(SW) // look southwest
	retval[6] = c.Peek(W)  // look west
	retval[7] = c.Peek(NW) // look northwest
	return retval
}

func TurnRight(dir int) int {
	// given the above declarations of directions, will return the
	// direction which is 90 degrees right from that particular direction.
	switch dir {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	default:
		panic("invalid direction")
	}
}

func TurnLeft(dir int) int {
	// just like TurnRight(), this will turn left.
	switch dir {
	case N:
		return W
	case E:
		return N
	case S:
		return E
	case W:
		return S
	default:
		panic("invalid direction")
	}
}

func Opposite(dir int) int {
	// Returns the opposite direction of whatever
	// is provided
	switch dir {
	case N:
		return S
	case NE:
		return SW
	case E:
		return W
	case SE:
		return NW
	case S:
		return N
	case SW:
		return NE
	case W:
		return E
	case NW:
		return SE
	default:
		panic("invalid direction")
	}
}

func IsInSquare(loc, topLeft, botRight Coord) bool {
	// Given loc, check to see if it is in the rect where the top left coordinate
	// and the bottom right coordinate exist as the opposite edges of the rectangle

	if loc.X >= topLeft.X && loc.X <= botRight.X &&
		loc.Y >= topLeft.Y && loc.Y <= botRight.Y {
		return true
	}
	return false
}
