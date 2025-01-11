package compumaze

import (
	"fmt"
	"utils"
)

// i suck at naming things

type Compumaze struct {
	utils.Runemap
	start, end utils.Coord
}

func NewCompumaze(in []string) Compumaze {
	newCompumaze := Compumaze{}
	newCompumaze.Runemap = utils.NewRunemap(in)
	newCompumaze.start, _ = newCompumaze.Find('S')
	newCompumaze.end, _ = newCompumaze.Find('E')
	newCompumaze.Set(newCompumaze.start, '.')
	newCompumaze.Set(newCompumaze.end, '.')
	return newCompumaze
}

func (c Compumaze) Print() {
	c.Runemap.Print()
	fmt.Printf("Start: X=%d, Y=%d\n", c.start.X, c.start.Y)
	fmt.Printf("End: X=%d, Y=%d\n", c.end.X, c.end.Y)
}

// Get all neighbors of all traversable coordinates
func (c Compumaze) GetNeighbors(loc utils.Coord) []utils.Coord {
	neighbors := loc.AllAvailable()
	retval := make([]utils.Coord, 0)

	for i := 0; i < len(neighbors); i++ {
		piece, err := c.Get(neighbors[i])
		if err != nil {
			// out of bounds
			continue
		}
		if piece == '.' {
			retval = append(retval, neighbors[i])
		}
	}
	return retval
}

// Get directions of walls (potentially to cheat)
func (c Compumaze) GetWalls(loc utils.Coord) []int {
	directions := []int{
		utils.N,
		utils.E,
		utils.S,
		utils.W,
	}
	retval := make([]int, 0)

	for _, n := range directions {
		thisDir := loc.Peek(n)
		piece, err := c.Get(thisDir)
		if err != nil {
			// border, ignore
			continue
		}
		if piece == '#' {
			retval = append(retval, n)
		}
	}
	return retval
}

func (c Compumaze) GetWallCoords(loc utils.Coord) []utils.Coord {
	// same thing as above, just makes sure to get the coordinates
	// of the walls instead of the directions.
	directions := []int{
		utils.N,
		utils.E,
		utils.S,
		utils.W,
	}
	retval := make([]utils.Coord, 0)

	for _, n := range directions {
		thisDir := loc.Peek(n)
		piece, err := c.Get(thisDir)
		if err != nil {
			// border, ignore
			continue
		}
		if piece == '#' {
			retval = append(retval, thisDir)
		}
	}
	return retval
}

func (c Compumaze) PeekCheat(loc utils.Coord, dir, howMany int) utils.Coord {
	// Given a location and a direction, will return the coordinate
	// howMany steps in that direction. We can use this to see if we've
	// been on that map piece before and if so we can find out how many
	// picoseconds it will save if we cheat in that direction

	newLoc := loc
	for i := 0; i < howMany; i++ {
		newLoc = newLoc.Peek(dir)
	}
	return newLoc
}

func GetCheatOptions(m *Compumaze, breadcrumbs map[utils.Coord]int, howMany int) map[int]int {
	// will parse through the breadcrumbs object and get the potential
	// places that a racer can phase through a wall and cheat, showing
	// the times that are saved. This will ONLY report a list of times
	// that are SAVED, and not if it goes through a wall to a place
	// that the racer has been before.
	retval := make(map[int]int)

	for k, v := range breadcrumbs {
		dirs := m.GetWalls(k)
		for _, d := range dirs {
			peek := m.PeekCheat(k, d, howMany)
			if peekScore, ok := breadcrumbs[peek]; ok {
				// we have this listing...
				if peekScore > v {
					// this is a cheat move, so let's score it
					// remember that moving through the wall takes 2 picoseconds
					// so make sure to account for that
					retval[(peekScore-v)-2]++
				}
			}
		}
	}
	return retval
}
