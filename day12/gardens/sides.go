package gardens

import "utils"

func ContinuedEdge(r *utils.Runemap, loc utils.Coord, lookDir int) bool {
	// This will determine if this particular coordinate is an edge in the direction
	// that we are looking. This means that it will peek in the direction and if the
	// rune in that peeked coordinate is _different_ or _out of bounds_ then it will
	// return a true. Otherwise it is not an edge and will return false.

	// first, get our current char
	currentChar, err := r.Get(loc)
	if err != nil {
		panic("invalid starting place")
	}

	peekChar, err := r.Get(loc.Peek(lookDir))
	if err != nil {
		// this is an edge
		return true
	}

	if peekChar != currentChar {
		// this is an edge
		return true
	}

	// otherwise this is within the region so not an edge based on this direction
	return false
}

// After some research, I discovered that the real way to determine the amount
// of actual sides is to determine the amount of _corners_. If you can count the
// corners of a region, you can count the sides. And the way to determine a corner
// is to check _orthoganlly_. Four right angles in either direction. Check each point
// within the region to determine if it is a corner. The way to accomplish that is
// by referring to the following two examples.
//
// . . .
// # # .
// # # .
//
// For each # plot, check orthoganally in all directions. If both points orthogally
// fall OUTSIDE of the region, then you have discovered an "outer corner." On the
// other hand, if you find an example like this:
//
// . . . .
// # # . .
// # # # .
// # # # .
//
// Then you have to check each point orthogannlly. IF you check each point orthogannly
// and they are both points in the region, BUT you also check diagonally, you have
// found an "inner corner."

func CountCorners(r *utils.Runemap, region []utils.Coord) int {
	// given a region, that is, a region stated as a stack of coordinates, determine
	// how many corners there are, and thus, how many sides.
	var corners int = 0

	// cache region results so we can check with ease
	var exists struct{}
	regionMap := make(map[utils.Coord]struct{})
	for _, v := range region {
		regionMap[v] = exists
	}

	for _, v := range region {
		// Set up directions to check
		var directions = [][3]int{
			{utils.N, utils.NE, utils.E}, {utils.E, utils.SE, utils.S},
			{utils.S, utils.SW, utils.W}, {utils.W, utils.NW, utils.N},
		}
		for _, d := range directions {
			ortho_1 := v.Peek(d[0])
			ortho_2 := v.Peek(d[1])
			ortho_3 := v.Peek(d[2])

			_, ortho_in_region_1 := regionMap[ortho_1]
			_, ortho_in_region_2 := regionMap[ortho_3]

			if ortho_in_region_1 && ortho_in_region_2 {
				// check the diagonal, and i'm so sorry for the number
				// scheme but it makes more sense in my head
				_, ortho_in_region_3 := regionMap[ortho_2]
				if !ortho_in_region_3 {
					// inner corner!
					corners++
					continue
				}
			}

			if !ortho_in_region_1 && !ortho_in_region_2 {
				// outer corner!
				corners++
			}
		}
	}

	return corners
}
