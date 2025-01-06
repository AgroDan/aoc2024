package computer

import "fmt"

// Now I'll write the functions designed to reverse engineer _my_ challenge.
// This is specific to my challenge so this is kinda punk rock. Though on
// further review hey maybe this would work for you too.

// Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0
// b = a % 8     bst
// b = b ^ 1     bxl
// c = a >> b    cdv
// b = b ^ 5     bxl
// b = b ^ c     bxc
// a = a >> 3    adv
// out += b % 8  out
// loop          jnz

func (c Computer) GenCompute(aVal int) []int {
	// This will run the program with a given a value and return
	// the output, for simplicity's sake
	newRound := c.Initialize(aVal)
	newRound.Run()
	return newRound.Output
}

func ValidCandidate(program, candidate []int) bool {
	// This will compare two attempts, and will eval to true if
	// the candidate array is the exact same as the END of the
	// program array
	if len(candidate) > len(program) {
		panic("candidate must be less than program. You probably found the answer!")
	}
	for i := 0; i < len(candidate); i++ {
		cTop := (len(candidate) - 1) - i
		pTop := (len(program) - 1) - i

		if program[pTop] != candidate[cTop] {
			return false
		}
	}
	return true
}

func (c Computer) Solver() {
	// going to work backwards now. I'm so sorry for this function.
	die := false
	candidates := []int{0}
	for {
		if die {
			break
		}
		die = true
		newCandidates := []int{}
		for _, v := range candidates {
			target := v << 3
			for i := target; i < target+100; i++ {
				attempt := c.GenCompute(i)
				if ValidCandidate(c.Program, attempt) {
					newCandidates = append(newCandidates, i)
					die = false
				}
			}
		}
		copy(candidates, newCandidates)
		fmt.Printf("Candidates: %v\n", candidates)
	}
}
