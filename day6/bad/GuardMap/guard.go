package guardmap

import "fmt"

/*
 * This is the little man or woman that marches around
 * the map, bumping into objects and turning right or
 * something. I think I can use interfaces here but
 * maybe later when I learn more about implementation.
 */

type Guard struct {
	guardMap *GuardMap // the map we will traverse
	currPos  Coord
	currDir  int
	steps    int
	bc       *Breadcrumb
	// lc       *Breadcrumb // the loop crumb! Let it traverse twice!
}

func NewGuard(c Coord, dir int, m *GuardMap) Guard {
	// returns a new little man to play with
	deezBreadcrumbs := NewBreadcrumb() // i really suck at names
	deezBreadcrumbs.Add(c, dir)
	// deezLoopcrumbs := NewBreadcrumb()
	return Guard{
		guardMap: m,
		currPos:  c,
		currDir:  dir,
		steps:    0,
		bc:       deezBreadcrumbs,
		// lc:       deezLoopcrumbs,
	}
}

// func (g Guard) GetLoopCrumbCount() int {
// 	return g.lc.Amount()
// }

func (g Guard) GetUnique() int {
	return g.bc.Amount()
}

func (g Guard) PrintObstacleMap(o []Coord) {
	// shut up
	g.guardMap.PrintObstacleMap(o)
}

func (g Guard) PrintMapCurrentState() {
	// Prints the map in the current state, where the guard is currently.
	for row, r := range g.guardMap.m {
		for col, c := range r {
			if row == g.currPos.Y && col == g.currPos.Y {
				switch g.currDir {
				case N:
					fmt.Printf("^")
				case E:
					fmt.Printf(">")
				case S:
					fmt.Printf("v")
				case W:
					fmt.Printf("<")
				}
				continue
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func (g *Guard) March() bool {
	// This function does all the heavy lifting. It takes *one* step.
	// If it hits an obstacle, it changes direction by moving to the
	// right. It only increments a step if it actually moves forward
	// by one. Returns True if any type of movement is possible.
	// Returns false if moving forward means we are outside of the
	// boundary.

	potentialMove := Move(g.currPos, g.currDir)
	if g.guardMap.IsValid(potentialMove) {
		// we can move here, so do so.
		g.currPos = potentialMove
		g.steps++
		g.bc.Add(potentialMove, g.currDir)
		return true
	}

	// Is this an obstacle? if IsInBounds returns true, then it's
	// an obstacle.
	if g.guardMap.IsInBounds(potentialMove) {
		// Move to the right
		switch g.currDir {
		case N:
			g.currDir = E
		case E:
			g.currDir = S
		case S:
			g.currDir = W
		case W:
			g.currDir = N
		}

		// update the breadcrumb
		g.bc.Add(g.currPos, g.currDir)
		return true
	}

	// otherwise it's out of bounds, so default false
	return false
}

// Helper function
func Move(c Coord, dir int) Coord {
	// This does no error checking, just returns the coordinate if
	// you move one step in the provided direction from the
	// provided coordinate
	switch dir {
	case N:
		return Coord{
			X: c.X,
			Y: c.Y - 1,
		}
	case E:
		return Coord{
			X: c.X + 1,
			Y: c.Y,
		}
	case S:
		return Coord{
			X: c.X,
			Y: c.Y + 1,
		}
	case W:
		return Coord{
			X: c.X - 1,
			Y: c.Y,
		}
	default:
		panic("invalid direction")
	}
}

// Turn function
func TurnRight(dir int) int {
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

func (g Guard) PrintCurrentState() {
	// Just for debugging purposes, this will print the current
	// map and the state of the little man running around.
	for row, r := range g.guardMap.m {
		for col, c := range r {
			if row == g.currPos.Y && col == g.currPos.X {
				switch g.currDir {
				case N:
					fmt.Printf("^")
				case E:
					fmt.Printf(">")
				case S:
					fmt.Printf("v")
				case W:
					fmt.Printf("<")
				}
				continue
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func (g Guard) GetSteps() int {
	// just returns the amount of steps
	return g.steps
}

// Lets create a scanner object complete with breadcrumbs too...
type Scanner struct {
	c   Coord
	dir int
	bc  *Breadcrumb
}

func NewScanner(currentPos Coord, currDir int) Scanner {
	return Scanner{
		c: Coord{
			X: currentPos.X,
			Y: currentPos.Y,
		},
		dir: currDir,
		bc:  NewBreadcrumb(),
	}
}

func (g *Guard) CheckForContinuousLoop() (bool, Coord) {
	// So this is what I'm going to do. I am going to create a scanner object
	// which will spawn on the current coordinate of the guard and turn right
	// 90 degrees. It will then loop through the method of traversing until
	// ONE of the following occurs: the boundary is hit, or a previously-
	// traversed coordinate is landed on AND it was landed on in the SAME
	// direction that we are currently facing.

	// This is what we will return if this coordinate is a suitable
	// place for an obstacle to generate an infinite loop, but ONLY
	// if this coordinate is an empty space!
	potentialObstacle := Move(g.currPos, g.currDir)
	if !g.guardMap.IsValid(potentialObstacle) {
		// This is invalid because either an obstacle is there already
		// or it is out of bounds.
		return false, potentialObstacle
	}

	// What defines an infinite loop is if we land on a space where the
	// scanner has previously been on facing the same direction as when
	// the scanner placed the breadcrumb.

	// turn the scanner right, since we are checking if we should place
	// the obstacle in front of the scanner and it turns right anyway
	scanner := NewScanner(g.currPos, TurnRight(g.currDir))

	for {
		// Did we hit a breadcrumb already? Is it going the same direction?
		if scanner.bc.Contains(scanner.c) && scanner.bc.GetDir(scanner.c) == scanner.dir {
			// technically also an infinite loop
			return true, potentialObstacle
		}

		// Let's set a breadcrumb for good measure.
		scanner.bc.Add(scanner.c, scanner.dir)

		// scan ahead
		checkSpace := Move(scanner.c, scanner.dir)

		// Is the space in front of the scanner the same
		// as the potential obstacle we proposed? If so, infinite loop.
		if potentialObstacle == checkSpace {
			return true, potentialObstacle
		}

		// Is it an empty space? If so move there and start over
		if g.guardMap.IsValid(checkSpace) {
			scanner.c = checkSpace
			continue
		}

		// Is it out of bounds? If so then we will never be in an infinite loop
		if !g.guardMap.IsInBounds(checkSpace) {
			return false, potentialObstacle
		}

		// Otherwise it's an obstacle. Turn right and try again
		scanner.dir = TurnRight(scanner.dir)
	}
}

// func (g *Guard) CheckForContinuousLoop() (bool, Coord) {
// 	// So this is what I'm going to do. I am going to create a scanner object
// 	// which will spawn on the current coordinate of the guard and turn right
// 	// 90 degrees. It will then loop through the method of traversing until
// 	// ONE of the following occurs: the boundary is hit, or a previously-
// 	// traversed coordinate is landed on AND it was landed on in the SAME
// 	// direction that we are currently facing.

// 	// This is what we will return if this coordinate is a suitable
// 	// place for an obstacle to generate an infinite loop, but ONLY
// 	// if this coordinate is an empty space!
// 	potentialObstacle := Move(g.currPos, g.currDir)
// 	if !g.guardMap.IsValid(potentialObstacle) {
// 		// This is invalid because either an obstacle is there already
// 		// or it is out of bounds.
// 		return false, potentialObstacle
// 	}

// 	scanner := NewScanner(g.currPos, TurnRight(g.currDir))
// 	// Don't forget the breadcrumb, but set it to the previous direction
// 	// so as not to close the for loop immediately, and besides it would
// 	// turn right anyway here because of the obstacle
// 	// scanner.bc.Add(scanner.c, g.currDir)

// 	for {
// 		// First, check if the current space we are on satisfies an infinite loop
// 		// based on the previous breadcrumbs from the guard.
// 		if g.bc.Contains(scanner.c) && g.bc.GetDir(scanner.c) == scanner.dir {
// 			// We are in an infinite loop
// 			return true, potentialObstacle
// 		}

// 		// Now check if the current space we are on satisifes an infinite loop
// 		// based on the breadcrumbs of the scanner
// 		if scanner.bc.Contains(scanner.c) && scanner.bc.GetDir(scanner.c) == scanner.dir &&
// 			g.guardMap.startingPos != scanner.c {
// 			return true, potentialObstacle
// 		}

// 		// Get the coordinate in front of the scanner in its new direction
// 		checkSpace := Move(scanner.c, scanner.dir)

// 		// Is it an empty space? If so move there and start over
// 		if g.guardMap.IsValid(checkSpace) {
// 			// Add the breadcrumb BEFORE we move.
// 			scanner.bc.Add(scanner.c, scanner.dir)

// 			scanner.c = checkSpace
// 			continue
// 		}

// 		// Is it out of bounds? If so then we will never be in an infinite loop
// 		if !g.guardMap.IsInBounds(checkSpace) {
// 			return false, potentialObstacle
// 		}

// 		// Otherwise it's an obstacle. Turn right and try again
// 		scanner.dir = TurnRight(scanner.dir)
// 	}
// }

// // For part two, I need to determine infinite loops. To perform an infinite loop,
// // I need to search to the right of every step to check if the following are true:
// // A previously traversed box exists where the direction is OPPOSITE of my current heading
// // An obstacle exists NEXT to it
// func (g *Guard) CheckForLoop() bool {
// 	var checkDir, oppDir int
// 	switch g.currDir {
// 	case N:
// 		checkDir = E
// 		oppDir = S
// 	case E:
// 		checkDir = S
// 		oppDir = W
// 	case S:
// 		checkDir = W
// 		oppDir = N
// 	case W:
// 		checkDir = N
// 		oppDir = E
// 	default:
// 		panic("invalid direction")
// 	}

// 	scanCursor := Coord{
// 		X: g.currPos.X,
// 		Y: g.currPos.Y,
// 	}

// 	// Now we have a direction to scan
// 	for {
// 		scanCursor = Move(scanCursor, checkDir)

// 		if g.guardMap.IsValid(scanCursor) {
// 			// We are in a blank space, we can continue to check
// 			// have we been here before, and if so were we moving
// 			// in the opposite direction?
// 			if g.bc.Contains(scanCursor) && g.bc.GetDir(scanCursor) == oppDir {
// 				// We found an intersection...is there an obstacle in the next step?
// 				checkOne := Move(scanCursor, checkDir)
// 				obj, err := g.guardMap.Get(checkOne)
// 				if err != nil {
// 					// we are out of bounds, this is invalid
// 					return false
// 				}

// 				if obj == '#' {
// 					// otherwise we found a valid loop!
// 					// First, note the loopcrumb to specify where the obstacle should be
// 					// g.lc.Add(Move(g.currPos, g.currDir), N)
// 					return true
// 				}

// 				// otherwise keep scanning until we're out of bounds.
// 				continue
// 			}
// 			// otherwise it's valid but we weren't there before, so start over
// 			continue
// 		}
// 		// Otherwise we either hit a wall without determining that there is a loop
// 		// or the outside of the map, so return false
// 		return false
// 	}

// }

// New method...now let's continue searching until EITHER
// We hit an out-of-bounds, which is false, OR we hit a
// section which we've already visited AND going in the
// same direction
// func (g *Guard) CheckForContinuousLoop() bool {
// 	var checkDir int
// 	switch g.currDir {
// 	case N:
// 		checkDir = E
// 	case E:
// 		checkDir = S
// 	case S:
// 		checkDir = W
// 	case W:
// 		checkDir = N
// 	default:
// 		panic("invalid direction")
// 	}

// 	scanCursor := Coord{
// 		X: g.currPos.X,
// 		Y: g.currPos.Y,
// 	}

// 	for {
// 		scanCursor = Move(scanCursor, checkDir)

// 		if g.guardMap.IsValid(scanCursor) {
// 			// this is in a blank space. Have we been there?
// 			// are we going in the right direction?
// 			if g.bc.Contains(scanCursor) && g.bc.GetDir(scanCursor) == checkDir {
// 				// we've been here, AND we are going in the same direction!
// 				return true
// 			}

// 			// otherwise, ignore everything. Check the next entry.
// 			continue
// 		}

// 		// Is the next space an obstacle?
// 		nextSpace, err := g.guardMap.Get(Move(scanCursor, checkDir))
// 		if err != nil {
// 			// we hit the boundary
// 			return false
// 		}

// 		if nextSpace == '#' {
// 			// obviously it's an obstacle...
// 			switch checkDir {
// 			case N:
// 				checkDir = E
// 			case E:
// 				checkDir = S
// 			case S:
// 				checkDir = W
// 			case W:
// 				checkDir = N
// 			default:
// 				panic("invalid direction")
// 			}
// 			if g.bc.GetDir(scanCursor) == checkDir {
// 				// this should be a loop
// 				return true
// 			}
// 			continue
// 		}
// 		// otherwise we're out of bounds
// 		return false
// 	}
// }

// func (g *Guard) Reset() {
// 	// resets to the beginning
// 	g.currPos = g.guardMap.startingPos
// 	g.currDir = g.guardMap.startingDir
// }
