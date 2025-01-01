package utils

func EuclideanModulus(x, y int) int {
	// This is the euclidean modulus, which is actually different
	// from Go's default modulus operation, which is actually
	// "technically correct." In this case, -1 % 10 will equal
	// 9, where in Go -1 % 10 will equal -1. This is useful for
	// things like rotating around a map where if you go over the
	// map border, it will teleport you to the opposite side.
	rem := x % y
	if rem < 0 {
		rem += y
	}

	return rem
}
