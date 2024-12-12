package guardmap

import "fmt"

/*
 * This will store the functions for the guardmap datatype, as well as the
 * parser function which will ingest the challenge data into a useable map.
 */

const (
	N = iota
	E
	S
	W
)

type Coord struct {
	X, Y int
}

type GuardMap struct {
	m           [][]rune // literal map of items
	startingPos Coord    // coordinate of starting position
	startingDir int      // direction facing of starting coord
}

func NewGuardMap(line []string) GuardMap {
	// This function takes in the challenge input
	// and builds the GuardMap object from it.
	m := GuardMap{}
	for row := 0; row < len(line); row++ {
		var rowSlice []rune
		for col, j := range line[row] {
			// go character by character. Note that this
			// paints the map from the top left to the bottom
			// right
			// WITNESS MY OVERENGINEERING!
			// WITNESS ME
			// WITNESS ME DAMN YOU
			switch j {
			case '^':
				m.startingPos.X = col
				m.startingPos.Y = row
				m.startingDir = N
				rowSlice = append(rowSlice, '.')
				continue
			case 'v':
				m.startingPos.X = col
				m.startingPos.Y = row
				m.startingDir = S
				rowSlice = append(rowSlice, '.')
				continue
			case '>':
				m.startingPos.X = col
				m.startingPos.Y = row
				m.startingDir = E
				rowSlice = append(rowSlice, '.')
				continue
			case '<':
				m.startingPos.X = col
				m.startingPos.Y = row
				m.startingDir = W
				rowSlice = append(rowSlice, '.')
				continue
			}
			rowSlice = append(rowSlice, j)
		}
		m.m = append(m.m, rowSlice)
	}
	return m
}

func (m GuardMap) PrintOriginalMap() {
	// Prints the original map with the starting position
	// of the guard, ensures it was parsed properly
	for row, r := range m.m {
		for col, c := range r {
			if row == m.startingPos.Y && col == m.startingPos.X {
				switch m.startingDir {
				case N:
					fmt.Printf("^")
				case E:
					fmt.Printf(">")
				case S:
					fmt.Printf("v")
				case W:
					fmt.Printf("<")
				}
				continue
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func (m GuardMap) PrintObstacleMap(o []Coord) {
	// Prints the map with proposed obstacles as O's
	for row, r := range m.m {
	colLoop:
		for col, c := range r {
			for _, obs := range o {
				if obs.X == col && obs.Y == row {
					fmt.Printf("O")
					if col == len(r)-1 {
						fmt.Printf("\n")
					}
					continue colLoop
				}
			}
			if row == m.startingPos.Y && col == m.startingPos.X {
				switch m.startingDir {
				case N:
					fmt.Printf("^")
				case E:
					fmt.Printf(">")
				case S:
					fmt.Printf("v")
				case W:
					fmt.Printf("<")
				}
				continue
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}
