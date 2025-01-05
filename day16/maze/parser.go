package maze

import (
	"fmt"
	"utils"
)

func NewMaze(lines []string) Maze {
	retval := Maze{
		m: utils.NewRunemap(lines),
	}

	retval.start, _ = retval.m.Find('S')
	retval.end, _ = retval.m.Find('E')
	return retval
}

func (m Maze) Print() {
	m.m.Print()
	fmt.Printf("Start: X-%d, Y-%d\n", m.start.X, m.start.Y)
	fmt.Printf("End: X-%d, Y-%d\n", m.end.X, m.end.Y)
}
