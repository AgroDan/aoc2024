package main

import (
	"bufio"
	topmap "day10/Topography"
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

	thisMap := topmap.NewTopoMap(lines)

	// thisMap.Print()

	myTrailheads := thisMap.Trailheads()

	var partOneScore int = 0
	for i := range myTrailheads {
		lilMan := topmap.NewHiker(myTrailheads[i], thisMap)
		lilMan.Plot()
		// fmt.Printf("Plotted from Trailhead X: %d, Y: %d -- Score: %d\n", myTrailheads[i].X, myTrailheads[i].Y, lilMan.Score())
		partOneScore += lilMan.Score()
	}

	// now do part two
	var partTwoScore int = 0
	for i := range myTrailheads {
		lilMan := topmap.NewHiker(myTrailheads[i], thisMap)
		topmap.PlotRating(&lilMan)
		// fmt.Printf("Plotted from Trailhead X: %d, Y: %d -- Rating: %d\n", myTrailheads[i].X, myTrailheads[i].Y, lilMan.Rating())
		partTwoScore += lilMan.Rating()
	}

	fmt.Printf("Total score for part one: %d\n", partOneScore)
	fmt.Printf("Total score for part two: %d\n", partTwoScore)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
