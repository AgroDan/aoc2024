package warehouse

import (
	"fmt"
	"utils"
)

// this is an item inside the warehouse. The warehouse item
// will have two points in the map pointing to the same struct
// all we need to know is if this is moveable or not. if #, then
// the answer is false
type item struct {
	IsMovable bool
}

func NewItem(m bool) *item {
	return &item{
		IsMovable: m,
	}
}

// This will effectively be the same object as a Warehouse, but with
// different methods pertaining to this particular data structure
type WideWarehouse struct {
	obj        map[utils.Coord]*item
	r          *Robot
	rows, cols int
}

func (w WideWarehouse) GetWidth(loc utils.Coord) (utils.Coord, utils.Coord) {
	// given a location, will check west and east to determine both sides of
	// the object

	myListing := w.obj[loc]
	// check west
	if checkLoc := loc.Peek(utils.W); w.obj[checkLoc] == myListing {
		return checkLoc, loc
	}
	// check east
	if checkLoc := loc.Peek(utils.E); w.obj[checkLoc] == myListing {
		return loc, checkLoc
	}

	// otherwise what are we doing????
	panic("what the hell are we checking")
}

func (w WideWarehouse) GetPiece(loc utils.Coord) rune {
	// This will return the piece of a given location using the logic
	// to determine a [ or ], unless of course it's a wall then whatever.

	// first, is anything here
	if _, exists := w.obj[loc]; !exists {
		return '.'
	}

	if w.obj[loc].IsMovable {
		// this is a box, so find out if this loc is the
		// left or right side
		myListing := w.obj[loc]
		if checkLoc := loc.Peek(utils.W); w.obj[checkLoc] == myListing {
			// this is the right side
			return ']'
		}

		if checkLoc := loc.Peek(utils.E); w.obj[checkLoc] == myListing {
			// this is the left side
			return '['
		}

		// ...otherwise wtf???
		return '?'
	}

	// otherwise it's a wall
	return '#'
}

func (w WideWarehouse) Print() {
	// did we do this right? Let's find out!
	fmt.Printf("Map:\n")
	for y := 0; y < w.rows; y++ {
		for x := 0; x < w.cols; x++ {
			thisCoord := utils.Coord{
				X: x,
				Y: y,
			}
			if w.r.loc == thisCoord {
				fmt.Printf("@")
				continue
			}

			fmt.Printf("%c", w.GetPiece(thisCoord))
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\nInstructions")
	for i := 0; i < len(w.r.inst); i++ {
		if i%100 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", w.r.inst[i])
	}
	fmt.Printf("\n")
}

func (w WideWarehouse) CanIMove(dir int, pos utils.Coord) bool {
	// this will operate similarly to the previous function only
	// because things are wider it will account for the X axis by 2.
	// also we have to make sure that both sides of the object can
	// move!

	if _, tile := w.obj[pos]; tile {
		// tile contains something
		if !w.obj[pos].IsMovable {
			return false
		}

		// otherwise, tile is a box...so make sure we can move the
		// entire box in that direction to accommodate!
		if dir == utils.E || dir == utils.W {
			fwd := pos.Peek(dir)
			return w.CanIMove(dir, fwd)
		}

		// otherwise, we're moving north or south.
		left, right := w.GetWidth(pos)

		leftFwd, rightFwd := left.Peek(dir), right.Peek(dir)
		return w.CanIMove(dir, leftFwd) && w.CanIMove(dir, rightFwd)
	}

	// nothing there so move
	return true
}

func (w *WideWarehouse) Push(dir int, pos utils.Coord) {
	// this will push boxes in the direction stated. If we're moving
	// east or west, move in that direction until a blank object is
	// found. Put every known object into the object map and move in
	// that dir. Also MAKE SURE YOU CHECK IF YOU CAN MOVE IN THIS DIR
	// or else it will all fall apart

	checkLoc := pos.Peek(dir)

	if dir == utils.E || dir == utils.W {
		if _, exists := w.obj[checkLoc]; exists {
			if !w.obj[checkLoc].IsMovable {
				panic("cannot push, hit a wall going E/W")
			}
			w.Push(dir, checkLoc)
		}

		// nothing here, so keep moving
		w.obj[checkLoc] = w.obj[pos]
		delete(w.obj, pos)
	} else {
		// otherwise we're going north or south, so the rules
		// change a little bit...we're going to move to the object
		// north or south of us, get both sides, and fork off from
		// there.
		left, right := w.GetWidth(pos)
		leftFwd, rightFwd := left.Peek(dir), right.Peek(dir)

		if _, exists := w.obj[leftFwd]; exists {
			if !w.obj[leftFwd].IsMovable {
				panic("cannot push, left side hit wall going N/S")
			}
			w.Push(dir, leftFwd)
		}

		// nothing here now so keep moving
		w.obj[leftFwd] = w.obj[left]
		delete(w.obj, left)

		if _, exists := w.obj[rightFwd]; exists {
			if !w.obj[rightFwd].IsMovable {
				panic("cannot push, right side hit wall going N/S")
			}
			w.Push(dir, rightFwd)
		}

		// same here
		w.obj[rightFwd] = w.obj[right]
		delete(w.obj, right)
	}
}

func (w *WideWarehouse) Move() bool {
	// just like the warehouse move, this will move one step in the process.
	// will return true if there are more instructions afterward. Otherwise
	// false (but still move on the last instruction) at the end of the
	// instruction set
	dir, retval := w.r.Next()

	// get the square in the direction of the next step
	checkLoc := w.r.loc.Peek(dir)

	if w.CanIMove(dir, checkLoc) {
		// is there a box?
		if _, exists := w.obj[checkLoc]; exists {
			// push the box
			w.Push(dir, checkLoc)
		}

		// move the robot
		w.r.loc = checkLoc
	} // otherwise we can't move but incremented anyway

	return retval
}

func (w WideWarehouse) PartTwoCalc() int {
	// This divies up the scores as per the instructions for part one

	var retval int = 0

	var exists struct{}
	boxLocations := make(map[utils.Coord]struct{})
	for k, v := range w.obj {
		if v.IsMovable {
			left, _ := w.GetWidth(k)
			boxLocations[left] = exists
		}
	}

	for loc := range boxLocations {
		retval += loc.X + (loc.Y * 100)
	}
	return retval
}
