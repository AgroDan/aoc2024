package towels

import (
	"fmt"
	"strings"
)

// This will define a towels object which will store the type of towels in the challenge.

func CanWeCombine(towelInventory map[string]bool, request string) bool {
	// the way I'm going to work this is to make it recursive.
	if len(request) == 0 {
		panic("cant request empty string")
	}

	charBuf := ""
	for i := 0; i < len(request); i++ {
		charBuf += string(request[i])
		if _, ok := towelInventory[charBuf]; ok {
			fmt.Printf("Found %s in inventory\n", charBuf)
			charBuf = ""
		}
	}
	return len(charBuf) == 0

}

func CanWeCombineReverse(towelInventory map[string]bool, request string) bool {
	// this works like CanWeCombine, but instead loops through every single
	// towelInventory object and adds up its length. If it's greater than
	// the length of the request string, then it's possible to create this
	// combination
	score := 0
	for towel := range towelInventory {
		towelLength := len(towel)
		amt := strings.Count(request, towel)
		fmt.Printf("Towel: %s, amount found: %d\n", towel, amt)
		score += (towelLength * amt)
	}

	return score >= len(request)
}

func CanWeCombineFlag(towelInventory map[string]bool, request string) bool {
	// This time I'll make sure all letters are accounted for.
	target := make(map[string]int)
	for _, char := range request {
		target[string(char)]++
	}

	found := make(map[string]int)

	for towel := range towelInventory {
		amt := strings.Count(request, towel)
		if amt > 0 {
			fmt.Printf("Found %s in %s %d times\n", towel, request, amt)
			for _, char := range towel {
				found[string(char)] += amt
			}
		}
	}

	totTarget := 0
	totFound := 0

	for _, v := range target {
		totTarget += v
	}

	for _, v := range found {
		totFound += v
	}
	fmt.Printf("total Target: %d, total Found: %d\n", totTarget, totFound)
	return totTarget <= totFound
}

// Now I'll just try to consider the items as "building blocks" and attempt to build
// the item bit by bit until we reach the end.

func CanWeBuild(towelInventory []string, challenge string, memo map[string]bool) bool {
	// if the challenge is empty, then it can always be built
	if challenge == "" {
		return true
	}

	// check memoized result
	if result, found := memo[challenge]; found {
		return result
	}

	// otherwise try each item in the set
	for _, towel := range towelInventory {
		if strings.HasPrefix(challenge, towel) {
			// if the item matches the prefix, check the remaining string recursively
			remaining := strings.TrimPrefix(challenge, towel)
			if CanWeBuild(towelInventory, remaining, memo) {
				memo[challenge] = true
				return true
			}
		}
	}

	// otherwise if no combo works then it ain't possible
	memo[challenge] = false
	return false
}

func CanFormTargetDP(target string, items []string) bool {
	// Create a DP table to track if substrings can be formed
	dp := make([]bool, len(target)+1)
	dp[0] = true // Base case: empty target can always be formed

	// Iterate over the target string
	for i := 1; i <= len(target); i++ {
		for _, item := range items {
			if i >= len(item) && target[i-len(item):i] == item && dp[i-len(item)] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(target)]
}

func PossibleBuildCandidates(target string, towels []string) int {
	// This will count how many possible permutations there are of a possible combination.
	memo := make(map[string]int)
	return possibleBCHandler(target, towels, memo)
}

func possibleBCHandler(target string, towels []string, memo map[string]int) int {
	// This will recurse adding up how many possible permutations
	// there are of a possible combination.
	if val, found := memo[target]; found {
		return val
	}

	if target == "" {
		// if the target is blank, then count it as one possible
		// combination
		return 1
	}

	total := 0

	for i := 0; i < len(towels); i++ {
		if strings.HasPrefix(target, towels[i]) {
			remaining := strings.TrimPrefix(target, towels[i])
			total += possibleBCHandler(remaining, towels, memo)
		}
	}

	// store results in memo
	memo[target] = total

	// return the total count of possible build candidates
	return total
}
