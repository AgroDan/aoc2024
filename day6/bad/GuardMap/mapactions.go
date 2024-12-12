package guardmap

import "fmt"

/*
 * This file will contain things that you can do with the map,
 * such as returning an item from the map, checking if a provided
 * coordinate is out of bounds, etc
 */

func (g GuardMap) Get(c Coord) (rune, error) {
	// will return the rune at the provided coordinate, returns
	// an error if it's out of bounds.

	if !g.IsInBounds(c) {
		return '?', fmt.Errorf("out of bounds: X: %d Y: %d", c.X, c.Y)
	}
	return g.m[c.Y][c.X], nil
}

func (g GuardMap) ReturnStart() (Coord, int) {
	// returns the coordinate of the starting position as well as
	// the direction of the character
	return g.startingPos, g.startingDir
}

func (g GuardMap) IsValid(c Coord) bool {
	// This will check to see if the provided coordinate is an open
	// space and not a wall or something.

	char, err := g.Get(c)
	if err != nil {
		// we're out of bounds
		return false
	}
	if char == '.' {
		return true
	}
	return false
}

func (g GuardMap) IsInBounds(c Coord) bool {
	// Given a coordinate, returns a true if the coordinate is within
	// the constraints of the map. False if not.

	if c.Y >= len(g.m) || c.Y < 0 {
		return false
	}

	if c.X >= len(g.m[c.Y]) || c.X < 0 {
		return false
	}
	return true
}
