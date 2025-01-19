package utils

import (
	"errors"
	"fmt"
)

// This package creates a rune map. If a map is a series of characters
// then this will apply coordinates to the map.

type Runemap struct {
	m [][]rune
}

// Note that this follows the way I read in input data presently,
// line by line and as a string. This builds from the top left
// to the bottom right.
func NewRunemap(in []string) Runemap {
	newRunemap := Runemap{}
	for _, r := range in {
		var rowSlice []rune
		for _, c := range r {
			rowSlice = append(rowSlice, rune(c))
		}
		newRunemap.m = append(newRunemap.m, rowSlice)
	}
	return newRunemap
}

func (r Runemap) Print() {
	// Just prints the runemap
	for y := range r.m {
		for x := range r.m[y] {
			fmt.Printf("%c", r.m[y][x])
		}
		fmt.Printf("\n")
	}
}

func (r Runemap) IsInBounds(c Coord) bool {
	// Given a coordinate, will return a true if it is inside
	// the map's boundaries, false if outside. This assumes
	// the map is a rectangle or square ONLY.
	if c.Y >= len(r.m) || c.Y < 0 {
		return false
	}
	if c.X >= len(r.m[c.Y]) || c.X < 0 {
		return false
	}

	return true
}

func (r Runemap) Get(c Coord) (rune, error) {
	// Given a coordinate, will return an item
	// inside that runemap. Will error out if
	// out of bounds
	if r.IsInBounds(c) {
		return r.m[c.Y][c.X], nil
	}
	return '?', fmt.Errorf("out of bounds: X: %d, Y: %d", c.X, c.Y)
}

func (r *Runemap) Set(c Coord, setRune rune) error {
	// Will set a specific coord to a supplied rune. Returns an error
	// if the supplied coordinate is invalid
	if r.IsInBounds(c) {
		r.m[c.Y][c.X] = setRune
		return nil
	}
	return fmt.Errorf("supplied coord is invalid, X: %d, Y: %d", c.X, c.Y)
}

func (r Runemap) Find(items ...rune) (Coord, error) {
	// Given a bunch of runes, will find the first instance of the first
	// item that it is given to find. Returns an error if it can't
	// find anything. You can use the "Get()" function to determine
	// which item it found I guess.
	retval := Coord{}
	for _, item := range items {
		for y := range r.m {
			for x := range r.m[y] {
				if r.m[y][x] == item {
					retval.X = x
					retval.Y = y
					return retval, nil
				}
			}
		}
	}
	return retval, errors.New("unable to find supplied items")
}

func (r Runemap) FindAll(item rune) []Coord {
	// like Find(), this only accepts ONE item and returns a list of
	// coordinates that this item was found in.
	var retval []Coord
	for y := range r.m {
		for x := range r.m[y] {
			if r.m[y][x] == item {
				retval = append(retval, Coord{
					X: x,
					Y: y,
				})
			}
		}
	}
	return retval
}

func (r Runemap) DeepCopy() *Runemap {
	// This function makes a deep copy of the runemap
	// elsewhere in memory and returns the address to
	// that runemap.
	newRunemap := Runemap{}
	for _, r := range r.m {
		var rowSlice []rune
		for _, c := range r {
			rowSlice = append(rowSlice, rune(c))
		}
		newRunemap.m = append(newRunemap.m, rowSlice)
	}
	return &newRunemap
}

func (r Runemap) Width() int {
	// returns the width of the rectangular map
	return len(r.m[0])
}

func (r Runemap) Height() int {
	// returns the height of the rectangular map
	return len(r.m)
}

func (r Runemap) GetRaw() [][]rune {
	// This will simply return the raw runemap
	retval := make([][]rune, 0)
	for _, row := range r.m {
		rows := make([]rune, len(row))
		copy(rows, row)
		retval = append(retval, rows)
	}
	return retval
}
