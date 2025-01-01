package main

import (
	"bufio"
	"day14/robotmap"
	"flag"
	"fmt"
	"os"
	"time"
	"utils"
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

	rows, cols := 103, 101
	// rows, cols := 7, 11 // testinput
	iter := 100
	var theseRobots []*robotmap.Robot
	for _, line := range lines {
		rob := robotmap.NewRobot(line, rows, cols)
		// rob.Print()
		theseRobots = append(theseRobots, rob)
	}

	quadLocations := robotmap.GetQuadrants(rows, cols)
	var quadrants = []int{0, 0, 0, 0}
	for i := range theseRobots {
		theseRobots[i].March(iter)

		for j := 0; j < len(quadrants); j++ {
			if utils.IsInSquare(theseRobots[i].Loc(), quadLocations[j][0], quadLocations[j][1]) {
				// fmt.Printf("Found in q %d\n", j)
				quadrants[j]++
				// fmt.Printf("Quadrants: %v+\n", quadrants)
				break
			}
		}
	}
	partOneAnswer := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	fmt.Printf("Safety score of part one: %d\n", partOneAnswer)
	// robotmap.PrintMap(rows, cols, theseRobots)

	// I'm going to just start over and search for a potential tree.

	var treeRobots []*robotmap.Robot
	for _, line := range lines {
		rob := robotmap.NewRobot(line, rows, cols)
		treeRobots = append(treeRobots, rob)
	}

	var partTwoAnswer int = 0
	for {
		// check for trees
		if robotmap.TreeDetection(treeRobots, cols, rows, 65.0) {
			break
		}

		for i := range treeRobots {
			treeRobots[i].March(1)
		}
		partTwoAnswer++
	}
	robotmap.PrintTree(rows, cols, treeRobots)
	fmt.Printf("Is this a tree? If so then it took %d seconds.\n", partTwoAnswer)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
