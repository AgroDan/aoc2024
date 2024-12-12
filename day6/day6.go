package main

import (
	"bufio"
	"day6/guard"
	"flag"
	"fmt"
	"os"
	"time"
	"utils"
)

func Part1(chalMap utils.Runemap) int {
	// just returns a person marching and counting how
	// many unique places they've been to.
	g := guard.NewGuard(chalMap)
	crumbs := utils.NewBreadcrumb()

	for {
		// first, are we in bounds?
		if !chalMap.IsInBounds(g.Pos) {
			return crumbs.Amount()
		}

		// set a breadcrumb
		crumbs.Add(g.Pos, g.Dir)

		// peek ahead
		ahead, _ := chalMap.Get(g.PeekForward())
		if ahead != '#' {
			g.MoveForward()
		} else {
			g.TurnRight()
		}
	}
}

func CheckForLoop(chalMap utils.Runemap) bool {
	// This simply returns a true or false if the map provided
	// contains an infinite loop.
	g := guard.NewGuard(chalMap)
	crumbs := utils.NewBreadcrumb()

	for {
		// first, are we in bounds?
		if !chalMap.IsInBounds(g.Pos) {
			return false
		}

		// have we been on this space and in this direction before?
		if crumbs.Contains(g.Pos) && crumbs.GetDir(g.Pos) == g.Dir {
			// by definition, if we have been on this space before
			// and facing the same direction, this is an infinite loop
			return true
		}

		// peek ahead
		ahead, _ := chalMap.Get(g.PeekForward())
		if ahead != '#' {
			crumbs.Add(g.Pos, g.Dir)
			g.MoveForward()
		} else {
			g.TurnRight()
		}
	}
}

func Part2(chalMap utils.Runemap) int {
	// This will brute force every single piece of the map
	// and change it to an obstacle then march it looking for
	// an infinite loop
	var loopCount int = 0
	for Y := 0; Y < chalMap.Height(); Y++ {
		for X := 0; X < chalMap.Width(); X++ {
			Cursor := utils.Coord{
				X: X,
				Y: Y,
			}
			mapPiece, err := chalMap.Get(Cursor)
			if err != nil {
				panic("invalid cursor placement")
			}

			if mapPiece == '^' {
				// ignore the starting position
				continue
			}

			if mapPiece != '#' {
				// fmt.Printf("Looking at %v+\n", Cursor)
				newMap := chalMap.DeepCopy()
				newMap.Set(Cursor, '#')

				if CheckForLoop(*newMap) {
					loopCount++
				}
			}
		}
	}
	return loopCount
}

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

	chalMap := utils.NewRunemap(lines)
	partOne := Part1(chalMap)

	partTwo := Part2(chalMap)

	fmt.Printf("Part 1 answer: %d\n", partOne)
	fmt.Printf("Part 2 answer: %d\n", partTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
