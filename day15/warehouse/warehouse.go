package warehouse

import (
	"fmt"
	"utils"
)

// this will contain the map as a series of runes in a map.

type Warehouse struct {
	obj        map[utils.Coord]rune
	r          *Robot
	rows, cols int
}

func (w Warehouse) Print() {
	// prints the warehouse
	fmt.Printf("Map:\n")
	for y := 0; y < w.rows; y++ {
		for x := 0; x < w.cols; x++ {
			thisCoord := utils.Coord{
				X: x,
				Y: y,
			}
			if _, exists := w.obj[thisCoord]; !exists {
				if w.r.loc == thisCoord {
					fmt.Printf("@")
				} else {
					fmt.Printf(".")
				}
				continue
			}
			fmt.Printf("%c", w.obj[thisCoord])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\nInstructions:")
	for i := 0; i < len(w.r.inst); i++ {
		if i%100 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", w.r.inst[i])
	}
	fmt.Printf("\n")
}

func (w Warehouse) CanIMove(dir int, pos utils.Coord) bool {
	// This is a boolean to determine if we can move in this
	// direction. It will be recursive, because it will continue
	// to check in that direction until it hits EITHER a wall
	// OR a blank space. If it hits a wall before it hits a
	// blank space, the answer is false.

	if _, tile := w.obj[pos]; tile {
		// tile contains something.
		if w.obj[pos] == '#' {
			return false
		}

		// otherwise, the tile is a box. recurse!
		fwd := pos.Peek(dir)
		return w.CanIMove(dir, fwd)
	}

	// nothing is there, so yeah we can move
	return true
}

func (w *Warehouse) Push(dir int, pos utils.Coord) {
	// This will push boxes in the direction stated.
	// it will be recursive if there are other boxes.
	// You have to make sure we CAN move first or this
	// will panic!
	checkLoc := pos.Peek(dir)

	if _, exists := w.obj[checkLoc]; exists {
		if w.obj[checkLoc] == '#' {
			panic("cannot push, hit a wall")
		}
		w.Push(dir, checkLoc)
	}

	// otherwise we're clear, so let's move
	w.obj[checkLoc] = w.obj[pos]
	delete(w.obj, pos)
}

func (w *Warehouse) Move() bool {
	// This will act as one step in the process. Will return
	// true if there are more instructions afterward. Will
	// return false (and move on the last instruction) if we
	// hit the end of the instruction set.
	dir, retval := w.r.Next()

	// get the square in the direction of the next step
	checkLoc := w.r.loc.Peek(dir)

	if w.CanIMove(dir, checkLoc) {
		// is there a box?
		if _, exists := w.obj[checkLoc]; exists {
			// push the box
			w.Push(dir, checkLoc)
		}

		// then move the robot
		w.r.loc = checkLoc
	} // otherwise we can't move, but we incremented anyway

	return retval
}

func (w Warehouse) PartOneCalc() int {
	// This divies up the scores as per the instructions for part one

	var retval int = 0
	for k, v := range w.obj {
		if v == 'O' {
			retval += k.X + (k.Y * 100)
		}
	}
	return retval
}
