package compumaze

import "utils"

// now for each point, a flood-fill algorithm should be employed to find
// the border points of every single piece in the wall.

func (c Compumaze) IsWallBorder(loc utils.Coord) bool {
	// Given a wall, will return a true if any of its neighbors
	// contains the '.' rune
	neighbors := loc.AllAvailable()
	for _, n := range neighbors {
		piece, err := c.Get(n)
		if err != nil {
			// out of bounds
			continue
		}
		if piece == '.' {
			return true
		}
	}
	return false
}

func (c Compumaze) GetValidPathsFromWall(loc utils.Coord) []utils.Coord {
	// given a wall, this will look all around for the '.' rune,
	// and if it finds it it will return it
	thisChar, err := c.Get(loc)
	if err != nil || thisChar != '#' {
		panic("starting location is not a wall")
	}

	neighbors := loc.AllAvailable()
	retval := make([]utils.Coord, 0)
	for _, n := range neighbors {
		piece, err := c.Get(n)
		if err != nil {
			// out of bounds
			continue
		}
		if piece == '.' {
			retval = append(retval, n)
		}
	}
	return retval
}

func FloodFill(starting utils.Coord, m *Compumaze, maxDistance int) []utils.Coord {
	// Given a starting point, will flood-fill the map and return
	// all the possible valid paths we can phase through to, given
	// the max distance we can phase.
	retval := make([]utils.Coord, 0)

	// some sanity checking...
	if startChar, err := m.Get(starting); err != nil || startChar != '#' {
		panic("starting location is not a wall")
	}

	queue := utils.NewGQueue[utils.Coord]()
	queue.Enqueue(starting)

	visited := make(map[utils.Coord]struct{})
	var beenthere struct{}

	totalPossiblePaths := make(map[utils.Coord]struct{})
	var exists struct{}

	for {
		if queue.IsEmpty() {
			break
		}

		next, _ := queue.Dequeue()

		// is this past the max distance?
		if utils.ManhattanDistance(starting, next) > maxDistance {
			// skip it, too far
			continue
		}

		if _, ok := visited[next]; ok {
			// been here before
			continue
		}

		// is this a wall?
		if thisChar, _ := m.Get(next); thisChar != '#' {
			// not a wall
			continue
		}

		// otherwise, get all points
		neighbors := next.AllAvailable()
		for _, n := range neighbors {
			getChar, err := m.Get(n)
			if err != nil {
				// hit a border, we don't care
				continue
			}
			if getChar == '.' {
				// valid path
				totalPossiblePaths[n] = exists
				continue
			}

			// otherwise we're a wall so queue it up...
			// as long as we haven't been there
			if _, ok := visited[n]; !ok {
				queue.Enqueue(n)
			}
		}

		visited[next] = beenthere
	}
	for k := range totalPossiblePaths {
		retval = append(retval, k)
	}
	return retval
}

// This, knowing now what I do about how the challenge is stated because
// I'm a complete moron and didn't read it properly, is a function that,
// given a point, will return the "Manhattan Radius" of all the points
// that are reachable from that point within a specific distance. This
// will ONLY return points in which it ends on a path, NOT a wall.

func (cm Compumaze) getAllPossible(c utils.Coord, radius int) []utils.Coord {
	retval := make([]utils.Coord, 0)
	AllPoints := utils.ManhattanRadius(c, radius)
	for i := 0; i < len(AllPoints); i++ {
		char, err := cm.Get(AllPoints[i])
		if err != nil {
			continue
		}
		if char == '#' {
			continue
		}

		// otherwise this is valid so add it
		retval = append(retval, AllPoints[i])
	}
	return retval
}

func ScoresInRadius(cm *Compumaze, radius int, start utils.Coord, steps map[utils.Coord]int) map[utils.Coord]int {
	// This will review every single coordinate in the points value
	// and compare it to the values in the steps map. If there are
	// any values which can actually save time, it will return those
	// coordinates with the amount of picoseconds saved
	points := cm.getAllPossible(start, radius)
	retval := make(map[utils.Coord]int)

	startVal := steps[start]
	for i := 0; i < len(points); i++ {
		checkVal := steps[points[i]]
		if utils.ManhattanDistance(start, points[i])+startVal < checkVal {
			// this is valid, let's record
			retval[points[i]] = checkVal - (startVal + utils.ManhattanDistance(start, points[i]))
		}
	}
	return retval
}
