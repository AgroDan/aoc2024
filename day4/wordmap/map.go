package wordmap

import (
	"fmt"
)

type Wordmap struct {
	m [][]rune
}

func NewWordmap(lines []string) Wordmap {
	w := Wordmap{}
	for _, v := range lines {
		// create a row
		var row []rune
		for _, k := range v {
			// now a column
			row = append(row, k)
		}
		w.m = append(w.m, row)
	}
	return w
}

func (w Wordmap) PrintMap() {
	// just so we can rest assured it was parsed properly
	for _, v := range w.m {
		for _, k := range v {
			fmt.Printf("%c", k)
		}
		fmt.Printf("\n")
	}
}

func (w Wordmap) Letter(wp WPointer) (rune, error) {
	// Returns a letter, given X/Y coordinates. Remember that
	// Y comes first, then X. Returns an error if out of bounds.
	if wp.Y < 0 || wp.X < 0 {
		return '?', fmt.Errorf("out of bounds, too low. X: %d Y: %d", wp.X, wp.Y)
	}
	if wp.Y >= len(w.m) || wp.X >= len(w.m[wp.Y]) {
		return '?', fmt.Errorf("out of bounds, too high, X: %d Y: %d", wp.X, wp.Y)
	}
	return w.m[wp.Y][wp.X], nil
}
