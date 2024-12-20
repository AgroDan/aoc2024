package topmap

import (
	"fmt"
	"utils"
)

type Topo struct {
	area       utils.Runemap
	trailheads []utils.Coord
}

func NewTopoMap(lines []string) Topo {
	// Reads in a topographical map and fills out the trail heads
	r := utils.NewRunemap(lines)
	th := r.FindAll('0')
	return Topo{
		area:       r,
		trailheads: th,
	}
}

func (t Topo) Trailheads() []utils.Coord {
	// returns coordinates of all the trailheads
	return t.trailheads
}

func (t Topo) Print() {
	// prints the topo map with some additional coords
	t.area.Print()

	fmt.Printf("\n")
	for c := range t.trailheads {
		fmt.Printf("TrailHead - X: %d, Y: %d\n", t.trailheads[c].X, t.trailheads[c].Y)
	}
}
