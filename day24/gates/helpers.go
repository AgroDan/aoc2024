package gates

import "fmt"

// I'm writing a new file for this because it's similar to tools but that was becoming too big

// GeneratePermutations generates all permutations of two pairs of items from a given slice.
func GeneratePermutations(items []string) [][][2]string {
	var result [][][2]string
	n := len(items)
	if n < 4 {
		return result // Not enough items to form two pairs
	}

	// Helper function to generate combinations and recursively build pairs
	var backtrack func(start int, pairs [][2]string)
	backtrack = func(start int, pairs [][2]string) {
		// If we have two pairs, add them to the result
		if len(pairs) == 2 {
			result = append(result, append([][2]string{}, pairs...))
			return
		}

		// Iterate through the items to form pairs
		for i := start; i < n; i++ {
			for j := i + 1; j < n; j++ {
				// Form a pair and continue recursively
				pair := [2]string{items[i], items[j]}
				rest := append(append([]string{}, items[:i]...), items[i+1:j]...)
				rest = append(rest, items[j+1:]...)
				backtrack(i+1, append(pairs, pair))
			}
		}
	}

	backtrack(0, nil)
	return result
}

func GetGateFromOutput(gates []*Gate, output string) *Gate {
	for _, g := range gates {
		if g.output == output {
			return g
		}
	}
	return nil
}

// This is going to be the same thing as the above, but will now generate 4 pairs instead of
// two, so I can use this on the main program to generate the permutations of the x, y, and z
func GeneratePermutationsFourPairs(items []string) [][][2]string {
	var result [][][2]string
	n := len(items)
	if n < 8 {
		return result // Not enough items to form four pairs
	}

	// Helper function to generate combinations and recursively build pairs
	var backtrack func(start int, pairs [][2]string)
	backtrack = func(start int, pairs [][2]string) {
		// If we have four pairs, add them to the result
		if len(pairs) == 4 {
			result = append(result, append([][2]string{}, pairs...))
			return
		}

		// Iterate through the items to form pairs
		for i := start; i < n; i++ {
			for j := i + 1; j < n; j++ {
				// Form a pair and continue recursively
				pair := [2]string{items[i], items[j]}
				rest := append(append([]string{}, items[:i]...), items[i+1:j]...)
				rest = append(rest, items[j+1:]...)
				backtrack(i+1, append(pairs, pair))
			}
		}
	}

	backtrack(0, nil)
	return result
}

func MergeMaps(m1, m2 map[*Gate]int) map[*Gate]int {
	for k, v := range m2 {
		m1[k] += v
	}
	return m1
}

func PrintImpactedGates(gates []*Gate, impactedGateList []string) {
	// This is just so I can pretty print it
	for _, v := range impactedGateList {
		selectedGate := GetGateFromOutput(gates, v)
		fmt.Println(selectedGate.String())
	}
}
