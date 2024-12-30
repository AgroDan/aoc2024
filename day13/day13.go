package main

import (
	"day13/arcade"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.ReadFile(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// changing things around because this is a weird file to parse
	challengeText := string(readFile)

	// Insert code here
	arcadeSet := arcade.Parser(challengeText)

	// remember, pushing button A costs 3 tokens.
	// pushing button B costs 1 token.

	var partOneTotal int = 0
	for _, a := range arcadeSet {
		a.Print()
		fmt.Printf("\n")

		calc := a.PrizeCalc()
		if len(calc) > 0 {
			var runningTally int = 99999999999999

			for _, v := range calc {
				aCost := v[0] * 3
				bCost := v[1]
				if aCost+bCost < runningTally {
					runningTally = aCost + bCost
				}
			}

			partOneTotal += runningTally
		}
		fmt.Printf("Prizes found: %v+\n", calc)
	}
	fmt.Printf("Lowest cost of button pushes for part 1: %d\n", partOneTotal)

	// Now let's try for part two...
	var partTwoTotal int = 0
	for i, thisArcade := range arcadeSet {
		A, B := thisArcade.SolvePart2()
		if A == 0 && B == 0 {
			continue
		}
		aCost := A * 3
		bCost := B
		partTwoTotal += (aCost + bCost)
		fmt.Printf("Arcade at position %d has value %d\n", i+1, aCost+bCost)
	}

	fmt.Printf("Lowest cost for button pushes for part 2: %d\n", partTwoTotal)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
