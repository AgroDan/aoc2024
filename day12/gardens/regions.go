package gardens

import "utils"

// This will contain a BFS search to determine a map

func GetRegion(r utils.Runemap, start utils.Coord) []utils.Coord {
	// this function will start at the provided coordinate in the
	// runemap and fan out looking for the entire region of that
	// particular coordinate. It will return an array of coordinates
	// showcasing every possible coordinate of that particular
	// region.

	var retval []utils.Coord

	regionLetter, err := r.Get(start)
	if err != nil {
		panic("must start with a valid coordinate")
	}

	// look everywhere
	pDirections := start.AllAvailable()

	// BFS means using a queue not a stack
	dirQueue := utils.NewGQueue[utils.Coord]()

	// Set up a visited set
	var exists struct{}
	visited := make(map[utils.Coord]struct{})

	// Now claim the visited spot
	visited[start] = exists

	// it's part of the region after all
	retval = append(retval, start)

	for i := range pDirections {
		checkLetter, err := r.Get(pDirections[i])
		if err != nil {
			// out of bounds or something
			continue
		}

		if checkLetter == regionLetter {
			dirQueue.Enqueue(pDirections[i])
		}
	}

	// Now start traversing
	for {
		checkSpace, containsAny := dirQueue.Dequeue()
		if !containsAny {
			// empty queue
			break
		}

		if _, ok := visited[checkSpace]; ok {
			// have we been here before? if yes then...
			continue
		}

		// otherwise this space is new. add it to the queue
		visited[checkSpace] = exists
		retval = append(retval, checkSpace)

		checkEverywhere := checkSpace.AllAvailable()
		for i := range checkEverywhere {
			checkLetter, err := r.Get(checkEverywhere[i])
			if err != nil {
				// out of bounds
				continue
			}

			if checkLetter == regionLetter {
				dirQueue.Enqueue(checkEverywhere[i])
			}
		}
	}

	return retval
}

func GetPerimeter(r utils.Runemap, loc utils.Coord) int {
	// based on the formula given, this will find a perimeter
	// of a specific location on the map. If a direction contains
	// EITHER an out-of-bounds error, OR a different letter from
	// the current letter, then that side counts towards 1 of the
	// perimeter value

	var retval int = 0

	regionLetter, err := r.Get(loc)
	if err != nil {
		// this region is invalid
		panic("invalid location to check")
	}

	// otherwise, check everywhere
	surrounding := loc.AllAvailable()
	for i := range surrounding {
		nextDoor, err := r.Get(surrounding[i])
		if err != nil {
			// map border
			retval++
			continue
		}

		// otherwise check
		if nextDoor != regionLetter {
			retval++
		}
	}
	return retval
}

func GetRegionPerimeter(r utils.Runemap, locSlice []utils.Coord) int {
	// This takes GetPerimeter a step further. Assuming you used the
	// GetRegion function to determine a full region, this will take
	// the output of that function and return the full perimeter of
	// the entire region
	var retval int = 0
	for i := range locSlice {
		retval += GetPerimeter(r, locSlice[i])
	}
	return retval
}
