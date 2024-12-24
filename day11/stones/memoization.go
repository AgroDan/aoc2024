package stones

import (
	"strconv"
	"strings"
)

// Now I will attempt memoization for running part 2 optimally. This
// means that now that I see that order _doesn't even matter in the
// slightest_, I can concentrate on just sheer amount of rocks. I will
// use a finite set/maps to calculate this.

// This will be a parser because the original parser I wrote created
// a stoneset, which was a linked list. This was handy for keeping
// an order, but since that doesn't matter I'm just going to concentrate
// on this

func Calculate(stoneLine string, iter int) int {
	// Returns only the amount of stones after the appropriate amount
	// of times have been calculated.
	workingStoneSet := make(map[Stone]int)
	theseStones := strings.Split(strings.TrimSpace(stoneLine), " ")
	for i := range theseStones {
		makeStone := NewStone(theseStones[i])
		workingStoneSet[makeStone]++
	}

	var count int = 0
	// now lets loop
	for i := 0; i < iter; i++ {
		iterStoneSet := make(map[Stone]int)
		for st := range workingStoneSet {
			if st.Id == "0" {
				iterStoneSet[NewStone("1")] += workingStoneSet[st]
				continue
			}

			if len(st.Id)%2 == 0 {
				// even digits
				l, r := st.Split()
				iterStoneSet[l] += workingStoneSet[st]
				iterStoneSet[r] += workingStoneSet[st]
				continue
			}

			// gonna do this manually because I think it
			// was messing with the iterator...

			multStone := NewStone(strconv.Itoa(st.NumberVal() * 2024))
			iterStoneSet[multStone] += workingStoneSet[st]
		}

		workingStoneSet = iterStoneSet
	}

	for st := range workingStoneSet {
		count += workingStoneSet[st]
	}

	return count
}
