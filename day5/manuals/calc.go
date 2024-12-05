package manuals

import "fmt"

/*
 * This section will perform all of the calculations
 * on the challenge to determine state
 */

// First, let's create something that, given two numbers, will
// determine if the second number is valid in coming _after_ the
// first. Otherwise will blow up.

func (m Manual) ValidOrder(first, second int) bool {
	// if for whatever reason the number doesn't exist, there
	// are no rules for it so pass. But check both directions
	// to be sure! First check has more weight.

	fVal, ok := m.Instructions.por[first]
	if ok {
		for _, v := range fVal {
			if v == second {
				return true
			}
		}
	}

	// If you got here, then the number doesn't have a rule appended
	// to it, so check the reverse -- there may be a rule stating that
	// the second number should come _before_ the first.
	sVal, ok := m.Instructions.por[second]
	if ok {
		for _, v := range sVal {
			if v == first {
				return false // invalid!
			}
		}
	}

	fmt.Printf("No known instructions for this: %d %d\n", first, second)
	return true
}

func (m Manual) ReturnValidInstructions() [][]int {
	// This loops through each instruction set and returns the instructions
	// that are valid as the returned 2-d array
	var retval [][]int
	for i := range m.Pages.p {
		for j := range m.Pages.p[i] {
			if j >= len(m.Pages.p[i])-1 {
				// made it here, it's valid
				retval = append(retval, m.Pages.p[i])
				break
			}

			// otherwise let's gooooo
			if !m.ValidOrder(m.Pages.p[i][j], m.Pages.p[i][j+1]) {
				break
			}
		}
	}
	return retval
}

func (m Manual) ReturnInvalidInstructions() [][]int {
	// Like the above, just negates it. I can just negate the above but
	// i don't care leave me alone
	var retval [][]int
	for i := range m.Pages.p {
		for j := range m.Pages.p[i] {
			if j >= len(m.Pages.p[i])-1 {
				break
			}
			if !m.ValidOrder(m.Pages.p[i][j], m.Pages.p[i][j+1]) {
				retval = append(retval, m.Pages.p[i])
				break
			}
		}
	}
	return retval
}
