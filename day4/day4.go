package main

import (
	"bufio"
	"day4/wordmap"
	"flag"
	"fmt"
	"os"
	"time"
)

/*
 * I over-engineered this. Witness my failure.
 */

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

	wm := wordmap.NewWordmap(lines)
	wm.PrintMap()

	// now find possible starts of the letter X

	searcher := wordmap.NewSearchIdx(&wm)
	foundXs := searcher.FindPossibleStarts()
	fmt.Printf("Found posible X's: %d\n", len(foundXs))

	var partOneCounter int = 0
	for _, entries := range foundXs {
		partOneCounter += searcher.FindPossibleMatches(entries)
	}

	// vals := searcher.FindPossibleMatches(foundXs[5])
	fmt.Printf("First possible entries for part 1: %d\n", partOneCounter)

	// Now find the possible X patterns
	foundAs := searcher.FindPossibleStartsPartTwo()
	var partTwoCounter int = 0
	for _, entry := range foundAs {
		if searcher.FindPossibleXs(entry) && searcher.IsValidX(entry) {
			partTwoCounter++
		}
	}
	fmt.Printf("Valid X's for part 2: %d\n", partTwoCounter)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
