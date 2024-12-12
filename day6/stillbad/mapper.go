package main

import (
	"fmt"
	"utils"
)

// I'm all over the place with this one, but the
// day6.go file was getting crowded.

// So here's ANOTHER attempt at logic. Now I'm just
// going to put an obstacle on every single path on
// the original map and re-run the algorithm to count
// how many times we loop. Screw it, I'm running out
// of patience now.

// I'll extend the runemap object

type Walker struct {
	r   utils.Runemap
	bc  utils.Breadcrumb
	loc utils.Coord
	dir int
}

func (w Walker) CheckForward() utils.Coord {
	// returns a coordinate of what is
	// in front of the guard
	return w.loc.Peek(w.dir)
}

func (w Walker) MoveForward() error {
	// Moves guard one space, or turns. no bounds
	// checking, that's for us to find out outside
	// of this function
	checkSpace := w.CheckForward()
	space, err := w.r.Get(checkSpace)
	if err != nil {
		// out of bounds
		return err
	}

	if space == '#' {
		// obstacle
		w.dir = utils.TurnRight(w.dir)
		return nil
	}

	w.loc.Move(w.dir)
	return nil
}

func NewWalker(parsedMap utils.Runemap) Walker {

	guardLoc, err := parsedMap.Find('^', '>', 'v', '<')
	if err != nil {
		panic("Couldn't find guard")
	}

	guardChar, _ := parsedMap.Get(guardLoc)
	var guardDir int
	switch guardChar {
	case '^':
		guardDir = utils.N
	case '>':
		guardDir = utils.E
	case 'v':
		guardDir = utils.S
	case '<':
		guardDir = utils.W
	}

	return Walker{
		r:   parsedMap,
		bc:  *utils.NewBreadcrumb(),
		loc: guardLoc,
		dir: guardDir,
	}
}

func (w Walker) ReleaseTheCrumbs() *utils.Breadcrumb {
	// returns a copy of the breadcrumbs
	return w.bc.DeepCopy()
}

func (w Walker) Walk() (int, int, error) {
	// this function will walk and return how many steps the guard
	// takes. It will also check to see if we are in an infinite
	// loop. If we are in an infinite loop, return an error. If
	// we hit the boundary, return how many steps we took and how
	// many unique spots we found. Note that
	// this won't change the original data in memory.
	var steps int = 0

	for {
		checkAhead, err := w.r.Get(w.CheckForward())
		if err != nil {
			// out of bounds
			steps++
			// create a breadcrumb to make sure the
			// amount is there
			w.bc.Add(w.loc, w.dir)
			return steps, w.bc.Amount(), nil
		}

		if checkAhead == '#' {
			// obstacle
			w.dir = utils.TurnRight(w.dir)
		}

		if w.bc.Contains(w.loc) && w.bc.GetDir(w.loc) == w.dir {
			// infinite loop
			return steps, w.bc.Amount(), fmt.Errorf("Infinite loop detected after %d steps", steps)
		}

		// plant a breadcrumb
		w.bc.Add(w.loc, w.dir)

		// move
		steps++
		w.loc.Move(w.dir)
	}
}

func (w Walker) WalkWithObstacle(obs utils.Coord) bool {
	// given an obstacle location, rebuild the map
	// with the obstacle in that space, then check
	// if we get an infinite loop.

	// remember not to count the starting space
	if obs == w.loc {
		// god i hope this works
		return false
	}

	mapCopy := w.r.DeepCopy()
	if err := mapCopy.Set(obs, '#'); err != nil {
		panic("Incorrect value set")
	}

	newGuard := NewWalker(*mapCopy)
	_, _, err := newGuard.Walk()
	if err != nil {
		return true
	}
	return false
}
