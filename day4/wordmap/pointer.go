package wordmap

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

type WPointer struct {
	X, Y int
}

func NewWPointer(X, Y int) WPointer {
	return WPointer{
		X: X,
		Y: Y,
	}
}

func (w *WPointer) Move(dir int) {
	// This moves the actual pointer
	switch dir {
	case N:
		w.Y--
	case NE:
		w.X++
		w.Y--
	case E:
		w.X++
	case SE:
		w.X++
		w.Y++
	case S:
		w.Y++
	case SW:
		w.X--
		w.Y++
	case W:
		w.X--
	case NW:
		w.X--
		w.Y--
	}
}

func ShiftDir(wp WPointer, dir int) WPointer {
	// Given a direction from the supplied wp,
	// returns a new WPointer from the direction
	// that was supplied. This does NOT do any
	// bounds-checking, this has to be done from
	// the function that returns a letter.
	switch dir {
	case N:
		return NewWPointer(wp.X, wp.Y-1)
	case NE:
		return NewWPointer(wp.X+1, wp.Y-1)
	case E:
		return NewWPointer(wp.X+1, wp.Y)
	case SE:
		return NewWPointer(wp.X+1, wp.Y+1)
	case S:
		return NewWPointer(wp.X, wp.Y+1)
	case SW:
		return NewWPointer(wp.X-1, wp.Y+1)
	case W:
		return NewWPointer(wp.X-1, wp.Y)
	case NW:
		return NewWPointer(wp.X-1, wp.Y-1)
	default:
		// just a defaultcase that really does nothing
		return NewWPointer(wp.X, wp.Y)
	}
}
