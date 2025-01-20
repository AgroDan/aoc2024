package main

import (
	"bufio"
	"day18/byteshower"
	"flag"
	"fmt"
	"os"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	mapWidth := flag.Int("w", 70, "Width of the map")
	bytesFallen := flag.Int("b", 1024, "Number of bytes to fall")

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
	bs := byteshower.ParseChallenge(lines, *mapWidth)
	bs.Fall(*bytesFallen)
	bs.DrawMap()
	path, gVal := byteshower.AStar(bs.Start(), bs.Goal(), &bs)
	fmt.Printf("Determined most efficient path steps after %d bytes fall: %.2f\n", *bytesFallen, gVal)

	bs.PrintPathway(path)

	// now we'll repeat until we find when the path gets cut off
	var thisCoord utils.Coord
	var iter int = *bytesFallen
	for {
		thisCoord = bs.FallAndGetCoord()
		iter++
		_, whoops := byteshower.AStar(bs.Start(), bs.Goal(), &bs)
		if whoops == 0 {
			break
		}
	}
	fmt.Printf("Path cut off at: (%d, %d) after %d bytes fell.\n", thisCoord.X, thisCoord.Y, iter)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
