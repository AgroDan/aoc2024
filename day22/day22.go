package main

import (
	"day22/pseudorandom"
	"flag"
	"fmt"
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
	// individual lines:
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

	partOneAnswer := 0
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		randomized := pseudorandom.Generator(val, 2000)
		partOneAnswer += randomized
		// fmt.Printf("%d: %d\n", val, randomized)
	}
	fmt.Printf("Part one answer: %d\n", partOneAnswer)

	// // Going to test with this one...
	// testVal := 123

	// generated := pseudorandom.GenerateAll(testVal, 10)
	// columnized := pseudorandom.GetCharacteristics(generated)
	// for _, row := range columnized {
	// 	// print the first column right-justified
	// 	fmt.Printf("%10d: %d %d\n", row[0], row[1], row[2])
	// }
	fmt.Printf("Generating monkeys...\n")
	monkeySet := make([]*pseudorandom.Monkey, 0)
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		monkeySet = append(monkeySet, pseudorandom.NewMonkey(val))
	}
	fmt.Printf("Monkeys generated.\n")
	fmt.Printf("Sequencing all possibilities...\n")
	allSequences := pseudorandom.GetAllCosts(monkeySet)
	fmt.Printf("Done, %d sequences found.\n", len(allSequences))

	totalVals := make(map[[4]int]int)
	for k := range allSequences {
		// fmt.Printf("Working on sequence: %v+...\n", k)
		totalVals[k] = pseudorandom.GetSequenceValue(monkeySet, k)
		// totalVals = append(totalVals, pseudorandom.GetSequenceValue(monkeySet, k))
	}
	most := 0
	mostSeq := [4]int{0, 0, 0, 0}
	for k, v := range totalVals {
		if v > most {
			most = v
			mostSeq = k
		}
	}
	fmt.Printf("The sequence [%d, %d, %d, %d] has the most bananas: %d\n", mostSeq[0], mostSeq[1], mostSeq[2], mostSeq[3], most)
	// fmt.Printf("Total values: %v+\n", totalVals)
	// Instead of using slices.Max() I'm going to make my own so I can index exactly
	// what I'm seeing...
	// fmt.Printf("Most bananas possible according to part two: %d\n", slices.Max(totalVals))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
