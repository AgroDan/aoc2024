package main

import (
	"day20/compumaze"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here
	maze := compumaze.NewCompumaze(lines)
	maze.Print()
	points := maze.Race()
	// for k, v := range points {
	// 	fmt.Printf("X=%d, Y=%d: %d\n", k.X, k.Y, v)
	// }
	cheatOptions := compumaze.GetCheatOptions(&maze, points, 2)
	totalPartOne := 0
	for k, v := range cheatOptions {
		fmt.Printf("There are %d cheats that save %d picoseconds.\n", v, k)
		if k >= 100 {
			totalPartOne += v
		}
	}
	fmt.Printf("Total cheats that save at least 100 picoseconds: %d\n\n", totalPartOne)

	part2CheatOptions := make(map[int]int)
	for k := range points {
		scores := compumaze.ScoresInRadius(&maze, 20, k, points)
		for _, v := range scores {
			part2CheatOptions[v]++
		}
	}

	// part2CheatOptions := make(map[int]int)
	// for p := range points {
	// 	// First, get the walls
	// 	walls := maze.GetWallCoords(p)
	// 	for _, w := range walls {
	// 		// Now get all possible points we can phase through
	// 		// at this wall
	// 		phaseable := compumaze.FloodFill(w, &maze, 21)

	// 		// now for each phaseable path, get the value
	// 		currVal := points[p]
	// 		for _, ph := range phaseable {
	// 			phaseVal := points[ph]
	// 			if phaseVal > currVal && currVal+utils.ManhattanDistance(p, ph) < phaseVal {
	// 				// We can cheat here
	// 				cheat := phaseVal - currVal
	// 				part2CheatOptions[cheat]++
	// 			}
	// 		}
	// 	}
	// }
	totalPartTwo := 0
	for k, v := range part2CheatOptions {
		if k >= 100 {
			fmt.Printf("There are %d cheats that save %d picoseconds.\n", v, k)
			totalPartTwo += v
		}
	}
	fmt.Printf("Total cheats that save at least 100 picoseconds: %d\n", totalPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
