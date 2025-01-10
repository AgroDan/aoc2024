package towels

import (
	"fmt"
	"strings"
)

func ParseChallenge(lines string) ([]string, []string) {
	// this will parse the challenge

	chal := strings.Split(lines, "\n\n")
	var towelInventory []string
	towels := strings.Split(chal[0], ",")
	for _, towel := range towels {
		towelInventory = append(towelInventory, strings.TrimSpace(towel))
	}

	var challenges []string
	combos := strings.Split(chal[1], "\n")
	for _, combo := range combos {
		challenges = append(challenges, strings.TrimSpace(combo))
	}
	return towelInventory, challenges

}

func PrintChallenge(towels []string, challenges []string) {
	// this will print the challenge
	fmt.Printf("Towels: ")
	var first bool = true
	for i := range towels {
		if first {
			fmt.Printf("%s", towels[i])
			first = false
		} else {
			fmt.Printf(", %s", towels[i])
		}
	}

	fmt.Printf("\n\nChallenges:\n")
	for _, chal := range challenges {
		fmt.Printf("%s\n", chal)
	}
}
