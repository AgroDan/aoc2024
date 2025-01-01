package robotmap

import (
	"fmt"
	"utils"
)

// ok this is probably the weirdest one I've done yet but I'll take it.

// Ok now let's look for robots with _at least_ 2 neighbors. If it hits
// a threshold then it may be a tree.

func TreeDetection(robots []*Robot, cols, rows int, threshold float64) bool {
	// In this detection, we'll look for a number of robots with at least
	// two neighbors and throw a true if the _amount_ of robots is over
	// the threshold. The threshold is a percentage, so panic if >100
	if threshold > 100 || threshold < 0 {
		panic("threshold must be a percentage, <100 >0")
	}

	// create a finite set
	robotMap := make(map[utils.Coord]int)
	for _, r := range robots {
		robotMap[r.loc]++
	}

	// get the threshold number
	limit := int((threshold / 100.0) * float64(len(robots)))

	// loop over all the robots again
	var possibleTreePoint int = 0
	for i, r := range robots {
		// neighbor counter
		var countAdjacent int = 0

		neighbors := r.loc.Neighbors()
		for _, n := range neighbors {
			if _, exists := robotMap[n]; exists {
				countAdjacent++

				if countAdjacent >= 2 {
					// optimization purposes
					possibleTreePoint++
					break
				}
			}
		}
		if possibleTreePoint >= limit {
			fmt.Printf("Possible tree points: %d, limit: %d\n", possibleTreePoint, limit)
			return true
		}

		// otherwise if we can't ever hit the threshold then just bail out
		if possibleTreePoint+(len(robots)-i) < limit {
			break
		}
	}
	return false
}

// Let's assume there's a tree if there are, let's say, ten robots vertically
// down the center. We can visually inspect.

func TreeDetectionFail2(robots []*Robot, cols, rows int) bool {
	robotMap := make(map[utils.Coord]int)

	for _, r := range robots {
		robotMap[r.loc]++
	}

	midPoint := utils.Coord{
		X: int(cols / 2),
		Y: int(rows / 2),
	}

	var checkSlice []utils.Coord
	for i := -5; i < 5; i++ {
		s := utils.Coord{
			X: midPoint.X,
			Y: midPoint.Y + i,
		}
		checkSlice = append(checkSlice, s)
	}

	for _, v := range checkSlice {
		if robotMap[v] <= 0 {
			return false
		}
	}
	return true
}

// Not gonna uncomment this but this didn't work well
func TreeDetectionFail(robots []*Robot, cols int) bool {
	// This is kinda arbitrary, but I'll assume that the christmas tree
	// starts at Y=0 and X=len(cols)/2, then Y=1 and X=len(cols)/2 as well
	// as next to it, etc. Detect that pattern I guess. More layers == more precision
	robotMap := make(map[utils.Coord]int)

	for _, r := range robots {
		robotMap[r.loc]++
	}

	midCol := int(cols / 2)
	// layer 1
	layer1 := utils.Coord{
		X: midCol,
		Y: 0,
	}
	if robotMap[layer1] <= 0 {
		return false
	}

	// layer 2
	layer2 := make([]utils.Coord, 3)
	layer2[0] = utils.Coord{
		X: midCol - 1,
		Y: 1,
	}
	layer2[1] = utils.Coord{
		X: midCol,
		Y: 1,
	}
	layer2[2] = utils.Coord{
		X: midCol + 1,
		Y: 1,
	}

	for _, v := range layer2 {
		if robotMap[v] <= 0 {
			return false
		}
	}

	// layer 3
	layer3 := make([]utils.Coord, 5)
	layer3[0] = utils.Coord{
		X: midCol - 2,
		Y: 2,
	}
	layer3[1] = utils.Coord{
		X: midCol - 1,
		Y: 2,
	}
	layer3[2] = utils.Coord{
		X: midCol,
		Y: 2,
	}
	layer3[3] = utils.Coord{
		X: midCol + 1,
		Y: 2,
	}
	layer3[4] = utils.Coord{
		X: midCol + 2,
		Y: 2,
	}

	for _, v := range layer3 {
		if robotMap[v] <= 0 {
			return false
		}
	}

	// probably a tree?
	return true

}
