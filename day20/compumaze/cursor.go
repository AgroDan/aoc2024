package compumaze

import "utils"

// This will be a cursor that will traverse the map and return
// a map of every coordinate pointing to the amount of steps
// taken

type Cursor struct {
	utils.Coord
	steps int
}

func NewCursor(c utils.Coord) Cursor {
	return Cursor{c, 0}
}

func (m Compumaze) Race() map[utils.Coord]int {
	// This is a recursive function that will return a map of coordinates
	// to the amount of steps taken to get there
	// It will take in a cursor and a map, and return a map
	// It will also take in a map of visited coordinates

	racer := NewCursor(m.start)

	// This will also serve as the return value
	visited := make(map[utils.Coord]int)
	visited[racer.Coord] = racer.steps

	// might as well make this DFS
	queue := utils.NewGQueue[utils.Coord]()

	neighbors := m.GetNeighbors(racer.Coord)

	// this should only be one but just in case...
	for _, n := range neighbors {
		queue.Enqueue(n)
	}
	for {
		if queue.IsEmpty() {
			break
		}

		next, _ := queue.Dequeue()
		if _, ok := visited[next]; ok {
			// been here before
			continue
		}

		racer.Coord = next
		racer.steps++

		// lay down the breadcrumb
		visited[racer.Coord] = racer.steps

		neighbors = m.GetNeighbors(racer.Coord)

		for _, n := range neighbors {
			queue.Enqueue(n)
		}

		if racer.Coord == m.end {
			// we're done
			break
		}
	}

	return visited
}
