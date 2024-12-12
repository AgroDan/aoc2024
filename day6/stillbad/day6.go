package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
	"utils"
)

// type Guard struct {
// 	loc utils.Coord
// 	dir int
// 	bc  utils.Breadcrumb
// }

// func (g Guard) CheckForward() utils.Coord {
// 	// returns a coordinate of what is
// 	// in front of the guard
// 	return g.loc.Peek(g.dir)
// }

// func (g *Guard) MoveGuard(rm *utils.Runemap) error {
// 	// Moves the guard one space, or turns.
// 	// No bounds checking, that's for us to find out
// 	// outside of this function
// 	checkSpace := g.CheckForward()

// 	space, err := rm.Get(checkSpace)
// 	if err != nil {
// 		return err
// 	}

// 	if space == '#' {
// 		// obstacle
// 		g.dir = utils.TurnRight(g.dir)
// 		return nil
// 	}

// 	g.loc.Move(g.dir)
// 	return nil
// }

// func (g *Guard) ScanForLoop(rm *utils.Runemap) bool {
// 	// This will create a new guard and loop until we
// 	// hit an infinite loop. This is meant to run for
// 	// each step.

// 	// First some sanity checks. Is the piece in front
// 	// either out of bounds or already has an obstacle?
// 	checkAhead, err := rm.Get(g.CheckForward())
// 	if err != nil {
// 		// out of bounds
// 		return false
// 	}

// 	if checkAhead == '#' {
// 		// technically if an obstacle is already there
// 		// then it wouldn't be an infinite loop.
// 		return false
// 	}

// 	potentialObstacle := g.CheckForward()

// 	scanner := Guard{
// 		loc: g.loc,
// 		dir: utils.TurnRight(g.dir),
// 		bc:  *utils.NewBreadcrumb(),
// 	}

// 	// Now let's see if we get what we need.
// 	for {
// 		aheadCoord := scanner.CheckForward()
// 		scanAhead, err := rm.Get(aheadCoord)
// 		if err != nil {
// 			// out of bounds
// 			return false
// 		}

// 		// Now see if we've been there in the same
// 		// direction, if so infinite loop
// 		if scanner.bc.Contains(scanner.loc) && scanner.bc.GetDir(scanner.loc) == scanner.dir {
// 			return true
// 		}

// 		if scanAhead == '#' || aheadCoord == potentialObstacle {
// 			scanner.dir = utils.TurnRight(scanner.dir)
// 			continue
// 		}

// 		// otherwise move
// 		scanner.bc.Add(scanner.loc, scanner.dir)
// 		scanner.loc.Move(scanner.dir)
// 	}
// }

// func (g *Guard) AmountVisited() int {
// 	return g.bc.Amount()
// }

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here

	thisMap := utils.NewRunemap(lines)

	// // Find the guard
	// guardLoc, err := thisMap.Find('^', 'v', '<', '>')
	// if err != nil {
	// 	panic("Couldn't find guard")
	// }

	// var guardDir int
	// guardChar, _ := thisMap.Get(guardLoc)
	// switch guardChar {
	// case '^':
	// 	guardDir = utils.N
	// case '>':
	// 	guardDir = utils.E
	// case 'v':
	// 	guardDir = utils.S
	// case '<':
	// 	guardDir = utils.W
	// }

	// // Now that we have the guard location, let's modify the
	// // map to be a blank space where the guard is.
	// thisMap.Set(guardLoc, '.')

	// // Now create the guard
	// g := Guard{
	// 	loc: guardLoc,
	// 	dir: guardDir,
	// 	bc:  *utils.NewBreadcrumb(),
	// }

	// fmt.Printf("Guard start loc: %v+\n", g.loc)
	// writer := bufio.NewWriter(os.Stdout)
	// var loopCount int = 0
	// for {
	// 	// Start moving
	// 	if !thisMap.IsInBounds(g.loc) {
	// 		break
	// 	}

	// 	// fmt.Printf("Guard loc: %v+\n", g.loc)

	// 	g.bc.Add(g.loc, g.dir)

	// 	// fmt.Printf("Loops?\n")
	// 	if g.ScanForLoop(&thisMap) {
	// 		loopCount++
	// 		fmt.Fprintf(writer, "\rConfirmed loops: %d", loopCount)
	// 		writer.Flush()
	// 	}

	// 	err := g.MoveGuard(&thisMap)
	// 	if err != nil {
	// 		fmt.Printf("Exited map!\n")
	// 		break
	// 	}
	// }

	// fmt.Printf("Amount visited: %d\n", g.AmountVisited())
	// fmt.Printf("Infinite loops detected: %d\n", loopCount)

	myGuard := NewWalker(thisMap)
	steps, visited, err := myGuard.Walk()
	if err != nil {
		panic(err)
	}

	listOfCrumbs := myGuard.ReleaseTheCrumbs()
	// now let's try this whole thing again
	var amtOfInfLoops int = 0
	for k := range listOfCrumbs.List() {
		if myGuard.WalkWithObstacle(k) {
			amtOfInfLoops++
		}
	}
	fmt.Printf("Steps taken: %d, Unique places: %d\n", steps, visited)
	fmt.Printf("Amount of infinite loops detected: %d\n", amtOfInfLoops)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
