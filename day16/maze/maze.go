package maze

import "utils"

type Maze struct {
	m          utils.Runemap
	start, end utils.Coord
}

type Reindeer struct {
	pos               utils.Coord
	dir, steps, turns int
	maze              *Maze
	bc                *utils.Breadcrumb
}

func NewReindeer(pos utils.Coord, maze *Maze, dir, steps, turns int, bc *utils.Breadcrumb) *Reindeer {
	thisBc := bc.DeepCopy()
	thisBc.Add(pos, dir)
	return &Reindeer{
		pos:   pos,
		dir:   dir,
		steps: steps,
		turns: turns,
		maze:  maze,
		bc:    thisBc,
	}
}

func (r Reindeer) CheckDirections() []*Reindeer {
	// This will check every possible direction (except backwards) and, if it is
	// possible to go in that direction, it will create another reindeer object
	// going in that direction, increasing the steps, direction, and turns.
	var retval []*Reindeer

	// first, forward.
	fwd := r.pos.Peek(r.dir)
	mapChar, err := r.maze.m.Get(fwd)
	if err != nil {
		panic("somehow looking out of bounds, don't know how that happened")
	}
	if mapChar != '#' && !r.bc.Contains(fwd) {
		// if we didn't hit a wall...
		newBC := r.bc.DeepCopy()
		newBC.Add(fwd, r.dir)
		retval = append(retval, NewReindeer(
			fwd, r.maze, r.dir, r.steps+1, r.turns, newBC,
		))
	}

	// check left
	left := utils.TurnLeft(r.dir)
	leftLoc := r.pos.Peek(left)
	leftChar, err := r.maze.m.Get(leftLoc)
	if err != nil {
		panic("turned left, out of bounds")
	}
	if leftChar != '#' && !r.bc.Contains(leftLoc) {
		newBC := r.bc.DeepCopy()
		newBC.Add(leftLoc, left)
		retval = append(retval, NewReindeer(
			leftLoc, r.maze, left, r.steps+1, r.turns+1, newBC,
		))
	}

	// check right
	right := utils.TurnRight(r.dir)
	rightLoc := r.pos.Peek(right)
	rightChar, err := r.maze.m.Get(rightLoc)
	if err != nil {
		panic("turned right, out of bounds")
	}
	if rightChar != '#' && !r.bc.Contains(rightLoc) {
		newBC := r.bc.DeepCopy()
		newBC.Add(rightLoc, right)
		retval = append(retval, NewReindeer(
			rightLoc, r.maze, right, r.steps+1, r.turns+1, newBC,
		))
	}

	return retval
}

func (r Reindeer) GetScore() int {
	// according to the first part of the challenge, this will
	// return the score of the reindeer.
	return r.steps + (r.turns * 1000)
}
