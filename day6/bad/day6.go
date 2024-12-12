package main

import (
	"bufio"
	guardmap "day6/GuardMap"
	"flag"
	"fmt"
	"os"
	"time"
)

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
	GM := guardmap.NewGuardMap(lines)
	// GM.PrintOriginalMap()

	fmt.Printf("\n\nNew map:\n")

	// Now initialize the lil' man
	startPos, startDir := GM.ReturnStart()
	G := guardmap.NewGuard(startPos, startDir, &GM)

	var loopCount int = 0
	proposedObstacleCoords := []guardmap.Coord{}

	obsBreadCrumb := guardmap.NewBreadcrumb()

	// For buffered STDOUT
	writer := bufio.NewWriter(os.Stdout)
	var steps int = 0
	for { // just keep swimming

		// if steps == 3213 {
		// 	fmt.Printf("\n")
		// 	G.PrintCurrentState()
		// 	break
		// }
		if !G.March() {
			break
		}
		// Now check for loops at each step, but if we put an obstacle there
		// already then just don't bother

		check, c := G.CheckForContinuousLoop()
		if check && !obsBreadCrumb.Contains(c) {
			loopCount++
			proposedObstacleCoords = append(proposedObstacleCoords, c)
			obsBreadCrumb.Add(c, 0) // 0 is meaningless but it works
		}
		steps++
		fmt.Fprintf(writer, "\rSteps Taken: %d, Confirmed loops: %d", steps, loopCount)
		writer.Flush()
	}
	fmt.Printf("\n")

	G.PrintObstacleMap(proposedObstacleCoords)

	// // now let's do it one more time!
	// G.Reset()
	// for {
	// 	if !G.March() {
	// 		break
	// 	}
	// 	G.CheckForLoop()
	// }
	// G.PrintCurrentState()
	fmt.Printf("Amount of steps taken for part 1: %d\n", G.GetUnique())
	fmt.Printf("Amount of potential infinite loops found for part 2: %d\n", loopCount)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
