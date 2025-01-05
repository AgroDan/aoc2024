package maze

import (
	"container/heap"
	"math"
	"utils"
)

// I'm going to attempt to use A* for traversal, since DFS was
// just incredibly CPU intensive and took far too long. Plus
// I really just want to understand how to implement A*

// This is effectively the reindeer I guess?
type Node struct {
	utils.Coord         // I only JUST discovered embedded structs
	G, H        float64 // G: Cost from start, H: Heuristic cost to goal
	Parent      *Node
	Steps       int
	Turns       int
	dir         int // current direction
}

func (n *Node) F() float64 {
	return n.G + n.H
}

// This priorityQueue object will implement the Heap interface
// by adding the capabilities of Less(), Len(), Swap(), Push() and Pop()
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// In this case, prioritize lower F() cost
	return pq[i].F() < pq[j].F()
}

func (pq PriorityQueue) Swap(i, j int) {
	// swap values
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func heuristic(a, b utils.Coord) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

func isValid(point utils.Coord, m *Maze) bool {
	// is this a valid place to move to? If so it would be a '.' or 'E' not a '#'
	rows, cols := m.m.Height(), m.m.Width()
	gridChar, _ := m.m.Get(point)
	return point.X >= 0 && point.X < rows && point.Y >= 0 && point.Y < cols && (gridChar == '.' || gridChar == 'E')
}

func reconstructPath(node *Node) []utils.Coord {
	var path []utils.Coord
	for node != nil {
		path = append([]utils.Coord{node.Coord}, path...)
		node = node.Parent
	}

	return path
}

func aStar(start, end utils.Coord, startingDir int, m *Maze) ([]utils.Coord, float64) {
	// The actual aStar path traversal algorithm
	openSet := &PriorityQueue{}

	heap.Init(openSet)
	heap.Push(openSet, &Node{
		Coord: start,
		G:     0,
		H:     heuristic(start, end),
		Steps: 0,
		Turns: 0,
		dir:   startingDir,
	})

	bestScore := make(map[utils.Coord]float64)
	bestScore[start] = 0

	var bestPath []utils.Coord
	lowestScore := math.Inf(1)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)
		if current.Coord == end {
			if current.G < lowestScore {
				lowestScore = current.G
				bestPath = reconstructPath(current)
			}
			continue
		}

		// right, left, and fwd
		var directions []int
		directions = append(directions, current.dir)
		directions = append(directions, utils.TurnRight(current.dir))
		directions = append(directions, utils.TurnLeft(current.dir))

		for _, d := range directions {
			neighbor := current.Peek(d)

			if !isValid(neighbor, m) {
				continue
			}

			turns := current.Turns
			if d != current.dir {
				// we're turning
				turns++
			}

			newSteps := current.Steps + 1
			newG := float64(newSteps) + (float64(turns) * 1000)
			newH := heuristic(neighbor, end)
			totalScore := newG + newH

			// check if this path to the neighbor is better...
			if best, exists := bestScore[neighbor]; !exists || totalScore < best {
				bestScore[neighbor] = totalScore
				neighborNode := &Node{
					Coord: neighbor,
					G:     newG,
					H:     newH,
					Steps: newSteps,
					Turns: turns,
					dir:   d,
				}
				heap.Push(openSet, neighborNode)
			}
		}
	}

	return bestPath, lowestScore // no path found
}
