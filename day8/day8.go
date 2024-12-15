package main

import (
	"bufio"
	"day8/antenna"
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
	ourMap := antenna.NewAntennaMap(lines)
	ourMap.PrintMap()

	antinodes := ourMap.FindAllAntinodes()
	fmt.Printf("Part one answer: %d\n", len(antinodes))
	resonantFreq := ourMap.FindAllResonantAntinodes()
	fmt.Printf("Part two answer: %d\n", len(resonantFreq))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
