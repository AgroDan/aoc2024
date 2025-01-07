package byteshower

import (
	"fmt"
	"utils"
)

// First let's define the byteshower object, which will consist
// of two things. An array of coordinates in the order that they
// were discovered, and a resulting map of the coordinates _after
// they've fallen_.

type Shower struct {
	coords []utils.Coord
	fallen map[utils.Coord]struct{}
	width  int
	idx    int
}

func (s *Shower) Fall(howMany int) {
	// This will drop as many "bytes" into the fallen map as nessessary.
	var dropped struct{}
	for i := 0; i < howMany; i++ {
		s.fallen[s.coords[s.idx]] = dropped
		s.idx++
	}
}

func (s *Shower) FallAndGetCoord() utils.Coord {
	var dropped struct{}
	retval := s.coords[s.idx]
	s.fallen[s.coords[s.idx]] = dropped
	s.idx++
	return retval
}

func (s Shower) DrawMap() {
	// Draws the map in accordance with what has so far fallen
	for y := 0; y <= s.width; y++ {
		for x := 0; x <= s.width; x++ {
			if _, ok := s.fallen[utils.Coord{X: x, Y: y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (s Shower) Start() utils.Coord {
	return utils.Coord{X: 0, Y: 0}
}

func (s Shower) Goal() utils.Coord {
	return utils.Coord{X: s.width, Y: s.width}
}

func (s Shower) PrintFallen() {
	// prints all the fallen bytes
	for k := range s.fallen {
		fmt.Printf("(X: %d, Y: %d)\n", k.X, k.Y)
	}
}

func (s Shower) PrintPathway(path []utils.Coord) {
	// to make things work a little faster I'll create
	// a map of all the paths
	var exists struct{}
	pathMap := make(map[utils.Coord]struct{})
	for _, v := range path {
		pathMap[v] = exists
	}

	// now print the map as we did before
	for y := 0; y <= s.width; y++ {
		for x := 0; x <= s.width; x++ {
			if _, ok := s.fallen[utils.Coord{X: x, Y: y}]; ok {
				fmt.Printf("#")
			} else if _, ok := pathMap[utils.Coord{X: x, Y: y}]; ok {
				fmt.Printf("O")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
