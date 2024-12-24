package stones

import "strconv"

/*
    If the stone is engraved with the number 0,
	it is replaced by a stone engraved with the number 1.

    If the stone is engraved with a number that has an
	even number of digits, it is replaced by two stones.
	The left half of the digits are engraved on the new
	left stone, and the right half of the digits are
	engraved on the new right stone. (The new numbers
	don't keep extra leading zeroes: 1000 would become
	stones 10 and 0.)

    If none of the other rules apply, the stone is
	replaced by a new stone; the old stone's number
	multiplied by 2024 is engraved on the new stone.
*/

type Stone struct {
	Id string
}

func NewStone(id string) Stone {
	// returns a stone object
	return Stone{
		Id: id,
	}
}

func (s Stone) NumberVal() int {
	// returns the integer value of the
	// stone's ID. Does the conversion
	// from string auto-magically
	val, _ := strconv.Atoi(s.Id)
	return val
}

func (s Stone) Split() (Stone, Stone) {
	// will split a stone into two stones
	// according to the second rule. Note
	// that this DOES NOT do any checking,
	// and will panic if this doesn't abide
	if len(s.Id)%2 != 0 {
		panic("amt of digits in ID must be even")
	}

	half := len(s.Id) / 2

	left := s.Id[:half]
	right := s.Id[half:]

	// convert BACK to a number JUST IN CASE
	// there are any leading zeroes, this should
	// remove them. 001100 should be 1100
	lNum, _ := strconv.Atoi(left)
	left = strconv.Itoa(lNum)

	rNum, _ := strconv.Atoi(right)
	right = strconv.Itoa(rNum)

	// then return as needed
	return NewStone(left), NewStone(right)
}

func (s *Stone) Mult() {
	// Will multiply a stone by 2024
	s.Id = strconv.Itoa(s.NumberVal() * 2024)
}
