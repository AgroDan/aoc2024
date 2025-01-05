package maze

import (
	"fmt"
	"utils"
)

// this will house the pathfinding algorithm

func AStarSolverPartOne(m Maze) int {
	path, score := aStar(m.start, m.end, utils.E, &m)
	if path != nil {
		return int(score)
	} else {
		return 999999999999999999
	}
}

func MazeSolverPartOne(m Maze, lowestAStar int) int {
	// this will use a breadth-first-search to find the shortest route
	// to the end square. Since I already wrote the AStar algorithm,
	// I'll use the lowest score from that to determine the maximum score
	// from this.

	// our starting reindeer, told to start facing east.
	rudolph := NewReindeer(m.start, &m, utils.E, 0, 0, utils.NewBreadcrumb())

	// travel queue
	tQueue := utils.NewGQueue[*Reindeer]()
	tQueue.Enqueue(rudolph)

	// copy the lowest score
	lowest := lowestAStar

	// create a visited metric
	visited := make(map[utils.Coord]int)

	// now let's goooooo
	for {
		if tQueue.IsEmpty() {
			break
		}

		// don't care about the second result because of isEmpty()
		r, _ := tQueue.Dequeue()

		// store the score since we don't want to be too computationally expensive
		thisScore := r.GetScore()

		if r.pos == m.end {
			if lowest > thisScore {
				lowest = thisScore
				fmt.Printf("Current lowest: %d\n", lowest)
			}
			continue
		}

		// Skip this if we've been here before with a better score
		if best, exists := visited[r.pos]; exists && thisScore >= best {
			continue
		}

		// Mark visited too
		visited[r.pos] = thisScore

		// utils.FlushScreen(fmt.Sprintf("Lowest score: %d", lowest))

		// otherwise, look around
		pDirs := r.CheckDirections()
		for i := range pDirs {
			// fmt.Printf("This reindeer: %v+ has a score of %d\n", pDirs[i], pDirs[i].GetScore())
			tQueue.Enqueue(pDirs[i])
		}
	}

	return lowest
}

func FindAllPointsInPath(m Maze, maxScore int) map[utils.Coord]struct{} {
	// This will traverse the map as before using BFS, but rather than look for the best
	// score, will ONLY return the coordinates that a reindeer used to find the end with
	// exactly that score.
	var exists struct{}
	retval := make(map[utils.Coord]struct{})

	// start with good ol' rudolph
	rudolph := NewReindeer(m.start, &m, utils.E, 0, 0, utils.NewBreadcrumb())

	// set the breadcrumb for metrics
	rudolph.bc.Add(rudolph.pos, rudolph.dir)

	// travel queue
	tQueue := utils.NewGQueue[*Reindeer]()
	tQueue.Enqueue(rudolph)

	// visited metric to speed things up
	visited := make(map[utils.Coord]int)

	for {
		if tQueue.IsEmpty() {
			break
		}

		r, _ := tQueue.Dequeue()

		thisScore := r.GetScore()

		if thisScore > maxScore {
			// this reindeer has outlived its potential
			continue
		}

		if r.pos == m.end && thisScore == maxScore {
			// this is a preferred route
			for k := range r.bc.List() {
				retval[k] = exists
			}
			continue
		}

		if best, exists := visited[r.pos]; exists && thisScore > best {
			continue
		}

		// mark visited
		visited[r.pos] = thisScore

		// look around
		pDirs := r.CheckDirections()
		for i := range pDirs {
			tQueue.Enqueue(pDirs[i])
		}
	}

	return retval
}

func CountUniquePoints(m Maze, lowestKnown int) int {
	// let's try this method now...count all the paths with the lowest score.
	rudolph := NewReindeer(
		m.start,
		&m,
		utils.E,
		0,
		0,
		utils.NewBreadcrumb(),
	)

	// set the breadcrumb
	rudolph.bc.Add(rudolph.pos, rudolph.dir)

	mQueue := utils.NewGQueue[*Reindeer]()
	mQueue.Enqueue(rudolph)

	pointScore := make(map[utils.Coord]int) // lowest score for each node

	var exists struct{}
	uniquePoints := make(map[utils.Coord]struct{}) // just lists unique points

	for {
		if mQueue.IsEmpty() {
			break
		}

		current, _ := mQueue.Dequeue()

		// save on calculations
		thisScore := current.GetScore()

		// drop a breadcrumb
		current.bc.Add(current.pos, current.dir)

		// check the score

		// NOTE: why am I adding 1000 to the pointscore of this tile? Well it's because
		// in a situation like this map:
		// #################
		// #...#...#...#..E#
		// #.#.#.#.#.#.#.#O#
		// #.#.#.#...#...#O#
		// #.#.#.#.###.#.#O#
		// #OOO#.#.#.....#O#
		// #O#O#.#.#.#####O#
		// #O#O..#.#.#OOOOX#
		// #O#O#####.#O###O#
		// #O#O#..OOOOO#OOO#
		// #O#O###O#####O###
		// #O#O#OOO#..OOO#.#
		// #O#O#O#####O###.#
		// #O#O#XOOOOOO..#.#
		// #O#O#O#########.#
		// #S#OOO..........#
		// #################

		// Where: S=Start, E=End, O=A valid path, and X=an intersection in question
		// at the "X" position there's a convergance of two different paths. Because
		// the score is recorded as +1000 when there's a turn, it's considered not
		// a viable path. So I added a threshold of 1000 to account for intersections
		// such as this where two paths may be heading towards the same goal but one
		// involves a turn at this junction.

		if ps, exists := pointScore[current.pos]; exists && (ps+1000) < thisScore {
			// this is a useless step since we've been here
			// before with a better score
			continue
		} else {
			pointScore[current.pos] = thisScore
		}

		if current.pos == m.end {
			// we're at the end
			if thisScore < lowestKnown {
				panic(fmt.Sprintf("i need the absolute lowest score, got %d", thisScore))
			} else if thisScore == lowestKnown {
				// we're at the end and have the lowest known score
				for k := range current.bc.List() {
					uniquePoints[k] = exists
				}
			}
			continue
		}

		// I'll try an anonymous function for once in my dumb life
		// this just checks to see if a location is a valid tile that
		// i could potentially move to, regardless if we've been there
		// or not
		isValid := func(loc utils.Coord) bool {
			thisChar, err := m.m.Get(loc)
			if err != nil {
				return false
			}
			if thisChar != '#' {
				return true
			}
			return false
		}

		// get all the directions. remember we increment turns by one
		// if we look left or right, and we won't turn around either
		// because that's inefficient
		directions := [3][2]int{
			{current.dir, 0},                  // fwd
			{utils.TurnLeft(current.dir), 1},  // left
			{utils.TurnRight(current.dir), 1}, // right
		}

		for _, d := range directions {

			lookLoc := current.pos.Peek(d[0])
			if !isValid(lookLoc) {
				// hit a wall, forget it
				continue
			}

			if current.bc.Contains(lookLoc) {
				// we've been there, forget it
				continue
			}

			// manage turns
			turns := current.turns + d[1]

			// new bc tracking
			newBC := current.bc.DeepCopy()

			// generate the neighbor reindeer
			qReindeer := NewReindeer(
				lookLoc,
				&m,
				d[0],
				current.steps+1,
				turns,
				newBC,
			)

			mQueue.Enqueue(qReindeer)
		}
	}

	return len(uniquePoints)
}
