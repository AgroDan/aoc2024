package byteshower

import (
	"container/heap"
	"math"
	"utils"
)

// Going to use A* to find the best path to reach the goal.
// It's defined that 0,0 is the starting point and width,width
// is the end.

type Node struct {
	utils.Coord
	G, H   float64
	Parent *Node
}

func (n *Node) F() float64 {
	return n.G + n.H
}

// so we can use a heap, need to implement it with
// basic functions to use it
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
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

func isValid(point utils.Coord, bs *Shower) bool {
	// is this a valid place to move to? If so then there wouldn't be
	// a corrupted space in that coord
	// first, are we in bounds?

	if point.X < 0 || point.Y < 0 || point.X > bs.width || point.Y > bs.width {
		return false
	}
	if _, exists := bs.fallen[point]; exists {
		return false
	}
	return true
}

func reconstructPath(node *Node) []utils.Coord {
	var path []utils.Coord
	for node != nil {
		path = append([]utils.Coord{node.Coord}, path...)
		node = node.Parent
	}
	return path
}

func AStar(start, end utils.Coord, bs *Shower) ([]utils.Coord, float64) {
	openSet := &PriorityQueue{}
	closedSet := make(map[utils.Coord]struct{})
	var beenthere struct{}

	heap.Init(openSet)
	heap.Push(openSet, &Node{
		Coord:  start,
		G:      0,
		H:      heuristic(start, end),
		Parent: nil,
	})

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)

		if current.Coord == end {
			return reconstructPath(current), current.G
		}

		closedSet[current.Coord] = beenthere

		// check all neighbors
		directions := current.Coord.AllAvailable()

		for _, neighbor := range directions {

			if !isValid(neighbor, bs) {
				continue
			}

			if _, exists := closedSet[neighbor]; exists {
				continue
			}

			gScore := current.G + 1
			hScore := heuristic(neighbor, end)

			// check if this is a new node
			found := false
			for _, node := range *openSet {
				if node.Coord == neighbor {
					found = true
					if gScore < node.G {
						node.G = gScore
						node.Parent = current
					}
				}
			}

			if !found {
				heap.Push(openSet, &Node{
					Coord:  neighbor,
					G:      gScore,
					H:      hScore,
					Parent: current,
				})
			}
		}
	}

	return nil, 0
}
