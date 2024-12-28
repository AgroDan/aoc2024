package main

import (
	"bufio"
	"day12/gardens"
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

	gardenMap := utils.NewRunemap(lines)

	partOneTotal := gardens.CrawlMap(gardenMap, false)
	fmt.Printf("Answer to part 1 is: %d\n", partOneTotal)

	partTwoTotal := gardens.CrawlMapPartTwo(gardenMap, false)
	fmt.Printf("Answer to part 2 is: %d\n", partTwoTotal)

	// // to test, this should show 12 coords
	// startLoc := utils.Coord{
	// 	X: 0,
	// 	Y: 0,
	// }
	// testLoc := gardens.GetRegion(gardenMap, startLoc)
	// fmt.Printf("Amount of regional spots for coord X: %d, Y: %d is %d\n", startLoc.X, startLoc.Y, len(testLoc))
	// fmt.Printf("Full perimeter of this region is: %d\n", gardens.GetRegionPerimeter(gardenMap, testLoc))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
