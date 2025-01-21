package pseudorandom

import "fmt"

// more monkeys. I'm just going to create an object showing a monkey with the
// precomputed values for the first 2000 rounds. This will allow me to work
// on the computation to determine proper market flucuations.

type Monkey struct {
	initialValue int
	Numbers      [][3]int
}

func NewMonkey(seed int) *Monkey {
	// This will generate everything needed
	// to assign to this particular monkey.
	allNumbers := GenerateAll(seed, 2000) // I can pass this via a parameter but this is part 2 already so
	return &Monkey{
		initialValue: seed,
		Numbers:      GetCharacteristics(allNumbers),
	}
}

func (m *Monkey) SearchForSequence(seq [4]int) (int, error) {
	// This will search for this particular sequence given and
	// return the cost of the bananas at that place in the
	// list of secret numbers. Returns err if not found

	for i := 4; i < len(m.Numbers); i++ {
		if seq[0] == m.Numbers[i-3][2] && seq[1] == m.Numbers[i-2][2] && seq[2] == m.Numbers[i-1][2] && seq[3] == m.Numbers[i][2] {
			return m.Numbers[i][1], nil
		}
	}
	return -1, fmt.Errorf("sequence not found")
}

// func (m *Monkey) ReturnAllSequences(number int) [][4]int {
// 	// Given a number, will return all of the previous 4 sequences
// 	// before that number. This number will most likely be coming up
// 	// many times in one, so it will return all of the sequences
// 	// it can find. Maybe we can memoize this?

// 	retval := make([][4]int, 0)
// 	for i, val := range m.Numbers {
// 		// we only care about the 4th number on
// 		if i < 3 {
// 			continue
// 		}
// 		if val[1] == number {
// 			retval = append(retval, [4]int{m.Numbers[i-3][2], m.Numbers[i-2][2], m.Numbers[i-1][2], m.Numbers[i][2]})
// 		}
// 	}
// 	return retval
// }

func (m *Monkey) MemoizeAllSequences(memo map[[4]int]int) {
	// This will look for every possible sequence, returning a map
	// that we can just use to memoize everything.
	// remember we need to skip the first number because it doesn't have
	// a valid "difference" to check against.

	for i := 4; i < len(m.Numbers); i++ {
		memo[[4]int{m.Numbers[i-3][2], m.Numbers[i-2][2], m.Numbers[i-1][2], m.Numbers[i][2]}]++
	}
}

func GetAllCosts(monkeys []*Monkey) map[[4]int]int {
	// I'm going to cross my fingers and hope this isn't ridiculously long to
	// compute, but the basic idea I'm going for here is I'm going to look for
	// all sequences behind every single number in the particular number, then
	// memoize each sequence. The sequence with the highest amount of times
	// found _once_ per monkey will be the winner.

	memo := make(map[[4]int]int) // This is pointless now but will help me figure out
	// how long this will take to compute at least.
	for i := 0; i < 10; i++ {
		// starting with one because starting with zero is pointless.
		for _, m := range monkeys {
			// this will keep track of the sequences found for
			// this iteration, so as not to count too many per
			// monkey.
			m.MemoizeAllSequences(memo)
			// foundMemo := make(map[[4]int]bool)
			// sequences := m.ReturnAllSequences(i)
			// for _, seq := range sequences {
			// 	if foundMemo[seq] {
			// 		continue
			// 	}
			// 	foundMemo[seq] = true
			// 	memo[seq]++
			// }
		}
	}
	return memo
}

func GetSequenceValue(monkeys []*Monkey, seq [4]int) int {
	// This, given a specific sequence, will return the value
	// of the bananas sold given this particular market flucutaion
	// sequence. Note that as soon as this sequence is found with
	// a given monkey's secret number set, it will cease looking through
	// the number set and move onto the next monkey.
	total := 0
	for _, m := range monkeys {
		val, err := m.SearchForSequence(seq)
		if err == nil {
			total += val
		}
	}
	return total
}
