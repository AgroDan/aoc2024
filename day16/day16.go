package main

import (
	"bufio"
	"day16/maze"
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

	reindeerMap := maze.NewMaze(lines)
	reindeerMap.Print()

	lowestAStar := maze.AStarSolverPartOne(reindeerMap)
	fmt.Printf("Lowest score discovered with A* for part one: %d\n", lowestAStar)
	fmt.Printf("Now attempting BFS for narrowing down score...\n")
	lowestBFS := maze.MazeSolverPartOne(reindeerMap, lowestAStar)
	fmt.Printf("Lowest score discovered with BFS for part one: %d\n", lowestBFS)

	fmt.Printf("Running through one more time to find all points in path...\n")
	uniquePoints := maze.CountUniquePoints(reindeerMap, lowestBFS)
	fmt.Printf("Lowest score part 2: %d, total points on ALL valid paths: %d\n", lowestBFS, uniquePoints)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
