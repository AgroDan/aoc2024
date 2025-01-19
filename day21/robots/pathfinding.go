package robots

import (
	"fmt"
	"utils"
)

// +---+---+
// | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

func GetOptimalPaths(code string, k *Keypad, currentPath string, result *[]string) {
	// this function will return the optimal paths for the given code.
	// The code MUST be two characters long or it will panic. This is
	// because I'm going to attempt to return the most optimal paths
	// from the start position to the end position and return it as
	// as set of <^>vA directions, where A will be pushing the button
	// as the last sequence always.
	if len(code) != 2 {
		panic("need string length of 2")
	}
	start, _ := k.keys.Find(rune(code[0]))
	end, _ := k.keys.Find(rune(code[1]))

	if start == end {
		*result = append(*result, currentPath+"A")
		return
		// pushing A because if we are on the same button
		// then the most we will need to do is just push
		// the button.
	}

	neighbors := start.TrueAllAvailable()
	currentCost := utils.ManhattanDistance(start, end)
	for dir, n := range neighbors {
		char, outOfBounds := k.keys.Get(n)
		if outOfBounds != nil || char == ' ' {
			continue
			// don't care about these
		}

		checkCost := utils.ManhattanDistance(n, end)
		if checkCost >= currentCost {
			continue
			// not worth our time
		}

		newCode := fmt.Sprintf("%s%s", string(char), string(code[1]))

		switch dir {
		case utils.N:
			GetOptimalPaths(newCode, k, currentPath+"^", result)
		case utils.E:
			GetOptimalPaths(newCode, k, currentPath+">", result)
		case utils.S:
			GetOptimalPaths(newCode, k, currentPath+"v", result)
		case utils.W:
			GetOptimalPaths(newCode, k, currentPath+"<", result)
		}
	}
}
