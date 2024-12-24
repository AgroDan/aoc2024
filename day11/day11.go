package main

import (
	"bufio"
	"day11/stones"
	"flag"
	"fmt"
	"os"
	"strings"
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
	ourStones := stones.Parse(strings.Join(lines, ""))
	ourStones.Print()

	ourStones.Iter(25)
	fmt.Printf("Answer to part 1: %d\n", ourStones.Count())

	// now lets try memoization
	partTwoStones := stones.Calculate(strings.Join(lines, ""), 75)
	fmt.Printf("Answer to part 2: %d\n", partTwoStones)
	// test := stones.NewStoneSet(stones.NewStone("30"))
	// test.Iter(50)
	// test.Count()

	// // 50 more times
	// ourStones.Iter(50)
	// fmt.Printf("Answer to part 2: %d\n", ourStones.Count())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
