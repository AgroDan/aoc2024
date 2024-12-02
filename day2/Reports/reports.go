package Reports

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	UNSET = iota
	INCREASING
	DECREASING
)

type ReportObject struct {
	numbers   []int
	issafe    bool
	direction int
}

func NewReportObject(line string) ReportObject {
	r := ReportObject{}
	vals := strings.Fields(line)
	for _, v := range vals {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("Not a number")
		}
		r.numbers = append(r.numbers, num)
	}
	r.issafe, r.direction = isSafe(r.numbers)
	return r
}

func (r ReportObject) PrintReport() {
	for _, v := range r.numbers {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("Direction: ")
	switch r.direction {
	case INCREASING:
		fmt.Printf("Increasing ")
	case DECREASING:
		fmt.Printf("Decreasing ")
	default:
		fmt.Printf("Unset ")
	}

	if r.issafe {
		fmt.Printf("Safe: Yes\n")
	} else {
		fmt.Printf("Safe: No, ")
		fmt.Printf("Safe after contingent: ")
		if r.ProblemDampenerContingent() {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}
}

func (r ReportObject) Safe() bool {
	// just an interface to the issafe var
	return r.issafe
}

func remove(s []int, idx int) []int {
	// a lot of troubleshooting to figure this out...Golang doesn't
	// deep copy a slice. Will modify the whole thing in place.
	dst := make([]int, 0, len(s)-1)
	dst = append(dst, s[:idx]...)
	dst = append(dst, s[idx+1:]...)
	return dst
}

func (r ReportObject) ProblemDampenerContingent() bool {
	// returns if this is safe even with the problem dampener
	// contingent. This will loop through every number and remove one
	// to check to see if it can be safe through any removed number.

	// first, is this safe regardless?
	if r.issafe {
		return true
	}

	for i := range r.numbers {
		newNumset := remove(r.numbers, i)
		// fmt.Printf("Idx: %d New numset: %v+\n", i, newNumset)
		newSafe, _ := isSafe(newNumset)
		if newSafe {
			return true
		}
	}
	return false
	// Now lets remove one from everything
	// for i := range r.numbers {
	// 	var newNumset []int

	// 	if i >= len(r.numbers) {
	// 		newNumset = r.numbers[:i]
	// 	} else {
	// 		newNumset = r.numbers[:i]
	// 		newNumset = append(newNumset, r.numbers[i+1:]...)
	// 		fmt.Printf("New numset: %v+\n", newNumset)
	// 	}
	// 	newSafe, _ := isSafe(newNumset)
	// 	if newSafe {
	// 		return true
	// 	}
	// }
	// return false
}

func isSafe(vals []int) (bool, int) {
	// This will determine if a number set is _SAFE_ based on
	// the following criteria:
	//
	// Levels are either all increasing or all decreasing
	// Two adjacent levels differ by at least one and at most 3
	status := UNSET
	for i := 0; i < len(vals)-1; i++ {
		res := vals[i+1] - vals[i]

		if status == UNSET {
			if res < 0 && res >= -3 {
				status = DECREASING
				continue
			}

			if res > 0 && res <= 3 {
				status = INCREASING
				continue
			}

			// Otherwise it's over the threshold
			return false, UNSET
		} else if status == INCREASING {
			if res <= 0 {
				// bouncing around rather than increasing
				return false, UNSET
			}

			if res > 3 {
				// too much off base
				return false, UNSET
			}

			// otherwise you're safe
		} else if status == DECREASING {
			if res >= 0 {
				return false, UNSET
			}

			if res < -3 {
				return false, UNSET
			}
		} else {
			// Don't know how you'd get here
			panic("We shouldn't be here")
		}
	}
	return true, status
}
