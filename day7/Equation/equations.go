package equation

import (
	"strconv"
	"strings"
	"utils"
)

const (
	Plus = iota
	Mult
	Cat
)

type Equation struct {
	Answer int
	values []int
}

func NewEquation(line string) Equation {
	// Creates the Equation{} object by parsing the string
	outer := strings.Split(line, ":")
	thisAnswer, err := strconv.Atoi(outer[0])
	if err != nil {
		panic(err)
	}

	var theseValues []int
	nums := strings.Split(strings.TrimSpace(outer[1]), " ")
	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}
		theseValues = append(theseValues, n)
	}
	return Equation{
		Answer: thisAnswer,
		values: theseValues,
	}
}

func (e Equation) IsValid() bool {
	// Try all permutations to determine equality
	permutations := len(e.values) - 1
	attempt := make([]int, permutations)
	var attemptPerms [][]int

	// generate permutations
	GenerateBinaryPermutations(attempt, permutations, &attemptPerms)

	for _, perms := range attemptPerms {

		// set up a queue for these permutations
		permQueue := utils.NewQueue()
		for i := range perms {
			permQueue.Enqueue(perms[i])
		}
		var total int = 0

		for i := 0; i < len(e.values); i++ {
			// since these are paired, it shouldn't ever hit an
			// index error...right?
			if i == 0 {
				// first item on the list
				total = e.values[i]
				continue
			}
			// pop the queue
			operation := permQueue.Dequeue()

			switch operation {
			case Plus:
				total += e.values[i]
			case Mult:
				total *= e.values[i]
			default:
				// I guess we hit the end of the queue?
				panic("hit the end of the queue")
			}
		}
		if e.Answer == total {
			return true
		}
	}
	return false
}

func (e Equation) IsValidPartTwo() bool {
	// this will now get even more permutations! yaaaaay
	permutations := len(e.values) - 1
	attempt := make([]int, permutations)
	var attemptPerms [][]int

	permutationItems := []int{Plus, Mult, Cat}
	// Generate Permutations

	GenerateTrinaryPermutations(attempt, permutations, permutationItems, &attemptPerms)

	for _, perms := range attemptPerms {
		// set up a queue for these permutations
		permQueue := utils.NewQueue()
		for i := range perms {
			permQueue.Enqueue(perms[i])
		}
		var total int = 0

		for i := 0; i < len(e.values); i++ {
			if i == 0 {
				total = e.values[i]
				continue
			}

			operation := permQueue.Dequeue()

			switch operation {
			case Plus:
				total += e.values[i]
			case Mult:
				total *= e.values[i]
			case Cat:
				total = ConcatNumbers(total, e.values[i])
			default:
				panic("hit the end of the queue")
			}
		}
		if e.Answer == total {
			return true
		}
	}
	return false
}
