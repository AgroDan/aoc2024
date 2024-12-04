package wordmap

import (
	"errors"
)

/*
 * This will perform the actual searching required to perform this challenge.
 * My intention is to iterate through every item in the wordlist and check
 * each direction for the word. It will return the index of every possible
 * word that it finds.
 */

type SearchIdx struct {
	pos WPointer
	wm  *Wordmap
}

func NewSearchIdx(thisMap *Wordmap) SearchIdx {
	// Creates a new object
	return SearchIdx{
		pos: NewWPointer(0, 0), // starting at 0,0 always
		wm:  thisMap,
	}
}

func (s SearchIdx) Current() (rune, error) {
	// this makes no changes to the current placement, but
	// instead returns the rune at the current place where
	// the pointer is. Returns an error if we are out of
	// bounds
	retval, err := s.wm.Letter(s.pos)
	if err != nil {
		return '?', err
	}
	return retval, nil
}

func (s *SearchIdx) Next() (rune, error) {
	// Updates the search index pointer to the next item in
	// the list, returning the rune that's there. This will
	// return an error if we hit the absolute end of the map!
	if s.pos.X >= len(s.wm.m[s.pos.Y])-1 {
		if s.pos.Y >= len(s.wm.m)-1 {
			return '?', errors.New("end of map")
		}
		s.pos.X = 0
		s.pos.Y++
		return s.wm.Letter(s.pos)
	}

	s.pos.X++
	return s.wm.Letter(s.pos)
}

func (s SearchIdx) FindPossibleStarts() []WPointer {
	// Find all possible starts of a word in the map.
	var retval []WPointer
	// start at the beginning
	current := NewWPointer(0, 0)

	// if there's an error here, something went dreadfully wrong
	check, _ := s.wm.Letter(current)

	if check == 'X' {
		retval = append(retval, NewWPointer(current.X, current.Y))
	}

	// now we can loop
	for {
		check, err := s.Next()
		if check == 'X' {
			retval = append(retval, NewWPointer(s.pos.X, s.pos.Y))
		}

		if err != nil {
			// we hit the end
			break
		}
	}
	return retval
}

func (s SearchIdx) FindPossibleStartsPartTwo() []WPointer {
	// finds all possible starts of a word in the map in accordance with the next
	// part of the challenge. In this case, it's the X-shaped "MAS". So let's look
	// for singular letter "A's"
	var retval []WPointer
	current := NewWPointer(1, 1) // starting here because the first row and col can be ignored

	check, _ := s.wm.Letter(current)
	if check == 'A' {
		retval = append(retval, NewWPointer(current.X, current.Y))
	}

	// now we loop like the dickens
	for {
		check, err := s.Next()
		if check == 'A' {
			retval = append(retval, NewWPointer(s.pos.X, s.pos.Y))
		}

		if err != nil {
			// we hit the end
			break
		}
	}
	return retval
}

func (s SearchIdx) FindPossibleMatches(p WPointer) int {
	// Given a position, finds as many possible instances of the word XMAS
	// in every given direction. Returns the amount it found.
	checkWord := []rune{'M', 'A', 'S'} // not XMAS because we know we have X at least
	directions := []int{N, NE, E, SE, S, SW, W, NW}
	var retval int = 0

	for _, dir := range directions {
		currPos := NewWPointer(p.X, p.Y)
		// for each direction
		var found bool = true
		for _, c := range checkWord {
			// for each direction
			currPos.Move(dir)
			checkLetter, err := s.wm.Letter(currPos)
			if err != nil {
				// hit a border, stop checking this direction
				found = false
				break
			}
			if checkLetter != c {
				// letter does not match, stop checking this direction
				found = false
				break
			}
		}
		// otherwise count it
		if found {
			retval++
		}
	}
	return retval
}

func (s SearchIdx) FindPossibleXs(p WPointer) bool {
	// This works a little bit different. Instead of looping through a bunch of letters
	// in as concise a way as possible, I'm going to just check 4 different directions
	// from the A. Whatever I'm just rushing it kinda. Also this doesn't return the amount
	// of possible things it finds, it returns true or false based on if it's a complete X.
	// Remember that there is more that needs to be done! This just determines if the right
	// letters are included in the X, not necessarily if they're in the right order!

	dirList := []int{NW, NE, SW, SE}

	for _, d := range dirList {
		// for each direction
		currPos := NewWPointer(p.X, p.Y)
		currPos.Move(d)
		thisLetter, err := s.wm.Letter(currPos)
		if err != nil {
			// hit a boundary
			return false
		}
		if thisLetter != 'M' && thisLetter != 'S' {
			return false
		}
	}
	return true
}

func (s SearchIdx) IsValidX(p WPointer) bool {
	// This will do the heavy lifting to determine what direction the "MAS" is facing, either
	// top, bottom, left or right. Will confirm if the X at the given point (REMEMBER THIS
	// MUST BE CONFIRMED WITH THE FindPossibleXs() FUNCTION ABOVE!) and return true if valid
	tl, err := s.peek(p, NW)
	if err != nil {
		return false
	}

	tr, err := s.peek(p, NE)
	if err != nil {
		return false
	}

	bl, err := s.peek(p, SW)
	if err != nil {
		return false
	}

	br, err := s.peek(p, SE)
	if err != nil {
		return false
	}

	// TOP DOWN
	if tl == 'M' && tr == 'M' {
		if bl == 'S' && br == 'S' {
			return true
		}
		return false
	}

	// BOTTOM UP
	if bl == 'M' && br == 'M' {
		if tl == 'S' && tr == 'S' {
			return true
		}
		return false
	}

	// LEFT TO RIGHT
	if tl == 'M' && bl == 'M' {
		if tr == 'S' && br == 'S' {
			return true
		}
		return false
	}

	// RIGHT TO LEFT
	if tr == 'M' && br == 'M' {
		if tl == 'S' && bl == 'S' {
			return true
		}
		return false
	}

	// otherwise quit it
	return false
}

func (s SearchIdx) peek(startPos WPointer, dir int) (rune, error) {
	// returns the rune of the provided direction. error if out of bounds
	thisPos := NewWPointer(startPos.X, startPos.Y)
	thisPos.Move(dir)
	return s.wm.Letter(thisPos)
}
