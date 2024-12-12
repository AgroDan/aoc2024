package utils

// This has to do with coordinates.

const (
	N = iota
	E
	S
	W
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
	default:
		panic("invalid direction")
	}
	return check
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
	case E:
		return W
	case S:
		return N
	case W:
		return E
	default:
		panic("invalid direction")
	}
}
