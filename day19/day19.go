package main

import (
	"day19/towels"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here
	// chalPtr := flag.Int("t", 0, "Challenge number")

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here
	towelInventory, challenges := towels.ParseChallenge(lines)
	towels.PrintChallenge(towelInventory, challenges)

	// chal := *chalPtr
	// memo := make(map[string]bool)
	// fmt.Printf("Using towel %d: %s, result: %t\n", chal, challenges[chal], towels.CanWeBuild(towelInventory, challenges[chal], memo))
	// fmt.Printf("Using towel %d: %s, result %d\n", chal, challenges[chal], towels.PossibleBuildCandidates(challenges[chal], towelInventory))

	totalPartOne := 0
	totalPartTwo := 0
	for _, v := range challenges {
		// the memo for each one
		// memo := make(map[string]bool)
		// if towels.CanWeBuild(towelInventory, v, memo) {
		// 	totalPartOne++
		// }

		// if towels.CanFormTargetDP(v, towelInventory) {
		// 	totalPartOne++
		// }
		howMany := towels.PossibleBuildCandidates(v, towelInventory)
		fmt.Printf("Found %d possible combinations for %s\n", howMany, v)
		if howMany > 0 {
			totalPartOne++
		}
		totalPartTwo += howMany
	}
	fmt.Printf("Total towel combinations possible for part 1: %d\n", totalPartOne)
	fmt.Printf("Total possible combinations of all towels for part 2: %d\n", totalPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
