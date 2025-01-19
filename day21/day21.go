package main

import (
	"day21/robots"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines
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

	partOneTotal := 0
	for _, line := range lines {
		fmt.Printf("Processing line: %s\n", line)
		numberPad := robots.NewNumberPad()
		dirs := numberPad.GetTotalPathsPer(line)
		least := math.Inf(1)
		for _, d := range dirs {
			// fmt.Printf("Working with %s\n", d)
			smallest := robots.GetMinLength(d, 1)
			if float64(smallest) < least {
				least = float64(smallest)
			}
		}
		fmt.Printf("Complexity: %d\n", int(least))
		num, _ := strconv.Atoi(line[:3])
		partOne := int(least) * num
		fmt.Printf("Value: %d\n", partOne)
		partOneTotal += partOne
	}

	fmt.Printf("Part one answer: %d\n", partOneTotal)

	sequences := robots.ComputeSequences(robots.NewDirectionPad())
	partTwoTotal := 0
	twoDepth := 25
	for _, line := range lines {
		fmt.Printf("Processing line: %s\n", line)
		numberPad := robots.NewNumberPad()
		dirs := numberPad.GetTotalPathsPer(line)
		least := math.Inf(1)
		cache := utils.NewCache()
		for _, d := range dirs {
			smallest := robots.CachedGetMinLength(d, twoDepth, sequences, cache)
			if float64(smallest) < least {
				least = float64(smallest)
			}
		}
		fmt.Printf("Complexity: %d\n", int(least))
		num, _ := strconv.Atoi(line[:3])
		partTwo := int(least) * num
		fmt.Printf("Value: %d\n", partTwo)
		partTwoTotal += partTwo
	}
	fmt.Printf("Answer for part two with depth of %d: %d\n", twoDepth, partTwoTotal)

	// for k, v := range sequences {
	// 	fmt.Printf("Key combo: %s: %v+\n", k, v)
	// }

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
