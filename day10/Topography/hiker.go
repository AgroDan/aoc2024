package topmap

import (
	"utils"
)

// This will show the functionality of the hiker in question.
// Effectively this will act as a cursor using breadcrumbs.
// Because I used breadcrumbs before, they involved a direction.
// Now it doesn't _look_ like I'll need to record the direction
// so I'm just going to say all the directions are 0, or North.

type Hiker struct {
	tmap   Topo
	loc    utils.Coord
	crumbs utils.Breadcrumb
	score  int
	rating int
}

func NewHiker(start utils.Coord, thisMap Topo) Hiker {
	return Hiker{
		tmap:   thisMap,
		loc:    start,
		crumbs: *utils.NewBreadcrumb(),
		score:  0,
		rating: 0, // this was unnecessary but whatever
	}
}

func (h *Hiker) DeepCopy() Hiker {
	// makes a deep copy of the hiker.
	return Hiker{
		tmap:   h.tmap,
		loc:    h.loc,
		crumbs: *h.crumbs.DeepCopy(),
		score:  h.score,
		rating: h.rating,
	}
}

func (h Hiker) Score() int {
	return h.score
}

func (h Hiker) Rating() int {
	return h.rating
}

func (h Hiker) ThisTile() rune {
	// returns the tile the hiker is
	// currently on
	thisRune, err := h.tmap.area.Get(h.loc)
	if err != nil {
		panic("invalid space")
	}
	return thisRune
}

func (h Hiker) ValidStep(c utils.Coord) bool {
	// This will determine if the Coordinate presented is
	// a valid step. If so, it will return true. False if no.
	// Generally it will check to see if the current step we
	// are on allows for it to be valid, by being 1 away from
	// the current, as well as if we've been there before.
	if h.crumbs.Contains(c) {
		return false
	}

	// otherwise, do some math
	currentRune, err := h.tmap.area.Get(h.loc)
	if err != nil {
		panic("we're not supposed to be on this loc")
	}
	currentVal := int(currentRune + '0') // again, so dumb

	// checker
	checkRune, err := h.tmap.area.Get(c)
	if err != nil {
		return false // invalid map space
	}
	checkVal := int(checkRune + '0')

	// if currentVal-checkVal < -1 || currentVal-checkVal > 1 {
	// 	// too big a drop or too high a wall
	// 	return false
	// }

	// above is an instruction i misread
	if checkVal-currentVal != 1 {
		return false
	}

	return true
}

func (h *Hiker) Plot() {
	// This will plot the course of each hiker
	// first set the breadcrumb
	h.crumbs.Add(h.loc, 0)

	// Set up the stack
	dirStack := NewCoordStack()

	// Get all the directions
	allDirs := h.loc.AllAvailable()

	for i := range allDirs {
		if h.ValidStep(allDirs[i]) {
			dirStack.Push(allDirs[i])
		}
	}

	// now we loop
	for {
		currDir, inUse := dirStack.Pop()
		if !inUse {
			// stack is empty
			break
		}

		// set the breadcrumb
		h.crumbs.Add(currDir, 0)

		// set the cursor location
		h.loc = currDir

		// Are we at a height of 9?
		thisTile, err := h.tmap.area.Get(h.loc)
		if err != nil {
			// this should be a valid square, what happened?
			panic("invalid square")
		}
		if thisTile == '9' {
			h.score++
			continue
		}

		lookingAround := h.loc.AllAvailable()
		for i := range lookingAround {
			if h.ValidStep(lookingAround[i]) {
				dirStack.Push(lookingAround[i])
			}
		}
	}
}

// so as not to get TOO confusing, I'm going to make this a helper
// function and not a method hanging off of the Hiker struct object

// Because I want this to be recursive, you have to pass this
// a new hiker object with the same breadcrumbs as before!

func PlotRating(thisHiker *Hiker) {

	// update the breadcrumb
	thisHiker.crumbs.Add(thisHiker.loc, 0)

	// are we on an end piece?
	if thisHiker.ThisTile() == '9' {
		thisHiker.rating++
	}

	// if not, look around
	allDirs := thisHiker.loc.AllAvailable()

	var allValidDirs []utils.Coord
	for i := range allDirs {
		if thisHiker.ValidStep(allDirs[i]) {
			allValidDirs = append(allValidDirs, allDirs[i])
		}
	}

	// step and recurse
	for i := range allValidDirs {
		// Let's generate another hiker
		newHiker := thisHiker.DeepCopy()
		newHiker.loc = allValidDirs[i]
		PlotRating(&newHiker)
		thisHiker.rating = newHiker.rating
	}
}
