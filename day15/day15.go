package main

import (
	"day15/warehouse"
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

	// Similar to day 13, this takes a whole challenge blob and passes
	// it unadulterated to the parser
	challengeText := string(readFile)

	// Insert code here

	thisWarehouse := warehouse.ParseChallenge(challengeText)

	for {
		val := thisWarehouse.Move()

		if !val {
			break
		}
	}
	fmt.Printf("GPS sum of part 1 is: %d\n", thisWarehouse.PartOneCalc())

	// now let's do part two
	thisWideWarehouse := warehouse.ParseChallengePartTwo(challengeText)

	for {
		val := thisWideWarehouse.Move()

		if !val {
			break
		}
	}
	// thisWideWarehouse.Print()
	fmt.Printf("GPS sum of part 2 is %d\n", thisWideWarehouse.PartTwoCalc())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
