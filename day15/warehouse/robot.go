package warehouse

import "utils"

type Robot struct {
	loc  utils.Coord // location
	inst []rune      // list of instructions
	idx  int         // index of current instruction
}

func NewRobot(c utils.Coord) *Robot {
	return &Robot{
		loc:  c,
		inst: make([]rune, 0), // will need this later
		idx:  0,
	}
}

func (r *Robot) addInstructions(directions string) {
	// This will take one giant string of instructions
	// and just add it to the queue of the robot. Make
	// sure you remove linebreaks in the directions string
	for i := range directions {
		r.inst = append(r.inst, rune(directions[i]))
	}
}

func (r *Robot) Next() (int, bool) {
	// will increment the index and return the next command.
	// will return a false if we reach the end of the index.
	// also will reset the index.
	var retval int
	switch r.inst[r.idx] {
	case '^':
		retval = utils.N
	case '>':
		retval = utils.E
	case 'v':
		retval = utils.S
	case '<':
		retval = utils.W
	default:
		panic("invalid direction")
	}
	r.idx++

	if r.idx >= len(r.inst) {
		// hit the end
		r.idx = 0
		return retval, false
	}
	return retval, true
}
