package equation

import "strconv"

// generate permutations
func GenerateBinaryPermutations(arr []int, size int, result *[][]int) {
	if size == 0 {
		// append copy of the current arr to the result
		perm := make([]int, len(arr))
		copy(perm, arr)
		*result = append(*result, perm)
		return
	}

	// set to Plus and recurse
	arr[len(arr)-size] = Plus
	GenerateBinaryPermutations(arr, size-1, result)

	// set the current pos to mult and recurse
	arr[len(arr)-size] = Mult
	GenerateBinaryPermutations(arr, size-1, result)
}

func GenerateTrinaryPermutations(arr []int, size int, items []int, result *[][]int) {
	if size == 0 {
		// apend a copy of the current array to result
		perm := make([]int, len(arr))
		copy(perm, arr)
		*result = append(*result, perm)
		return
	}

	// iterate over the possible items
	for _, item := range items {
		arr[len(arr)-size] = item
		GenerateTrinaryPermutations(arr, size-1, items, result)
	}
}

func ConcatNumbers(left, right int) int {
	// concatenates two digits into one giant digit. I am going to cheat a little
	// with this one by converting it to a string, concatenating the string, and
	// converting back to an integer

	leftStr := strconv.Itoa(left)
	rightStr := strconv.Itoa(right)
	retval, err := strconv.Atoi(leftStr + rightStr)
	if err != nil {
		panic("could not convert")
	}
	return retval
}
