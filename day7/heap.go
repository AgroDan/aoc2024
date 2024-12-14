package main

func Heap(arr []int, n int, result *[][]int) {
	if n == 1 {
		// Make a copy of the current permutation and append it to the result
		perm := make([]int, len(arr))
		copy(perm, arr)
		*result = append(*result, perm)
		return
	}

	for i := 0; i < n; i++ {
		Heap(arr, n-1, result)
		if n%2 == 0 {
			// If n is even, swap the ith element with the last element
			arr[i], arr[n-1] = arr[n-1], arr[i]
		} else {
			// If n is odd, swap the first element with the last element
			arr[0], arr[n-1] = arr[n-1], arr[0]
		}
	}
}
