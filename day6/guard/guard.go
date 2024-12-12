package guard

import "utils"

// I am so frustrated.

type Guard struct {
	Pos utils.Coord
	Dir int
}

func NewGuard(areaMap utils.Runemap) Guard {
	// This will create a guard based on the map
	// that it was given. It will look for a '^'
	// to specify which position the guard is in.
	foundGuard, err := areaMap.Find('^')
	if err != nil {
		panic("could not find guard")
	}
	return Guard{
		Pos: foundGuard,
		Dir: utils.N,
	}
}

func (g *Guard) TurnRight() {
	// i'll just do it myself
	switch g.Dir {
	case utils.N:
		g.Dir = utils.E
	case utils.E:
		g.Dir = utils.S
	case utils.S:
		g.Dir = utils.W
	case utils.W:
		g.Dir = utils.N
	default:
		panic("invalid direction")
	}
}

func (g Guard) PeekForward() utils.Coord {
	// returns the coordinate in front of the guard
	retVal := utils.Coord{
		X: g.Pos.X,
		Y: g.Pos.Y,
	}
	switch g.Dir {
	case utils.N:
		retVal.Y--
	case utils.E:
		retVal.X++
	case utils.S:
		retVal.Y++
	case utils.W:
		retVal.X--
	default:
		panic("invalid direction")
	}
	return retVal
}

func (g *Guard) MoveForward() {
	// just sets the guard's position one place in the
	// direction the guard is moving
	g.Pos = g.PeekForward()
}
