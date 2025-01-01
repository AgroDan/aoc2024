package robotmap

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

// This will contain the robot objects.

type Robot struct {
	startPos, vel, loc utils.Coord
	rows, cols         int
}

func NewRobot(in string, rows, cols int) *Robot {
	// This will create a pointer to a robot
	// based on the input string given in the
	// challenge.
	attr := strings.Split(strings.TrimSpace(in), " ")
	p := strings.Split(attr[0], "=")
	p2 := strings.Split(p[1], ",")
	pX, _ := strconv.Atoi(p2[0])
	pY, _ := strconv.Atoi(p2[1])
	pLoc := utils.Coord{
		X: pX,
		Y: pY,
	}

	v := strings.Split(attr[1], "=")
	v2 := strings.Split(v[1], ",")
	vX, _ := strconv.Atoi(v2[0])
	vY, _ := strconv.Atoi(v2[1])
	vel := utils.Coord{
		X: vX,
		Y: vY,
	}

	return &Robot{
		startPos: pLoc,
		vel:      vel,
		loc:      pLoc,
		rows:     rows,
		cols:     cols,
	}
}

func (r Robot) Print() {
	fmt.Printf("Robot starting location: X -> %d, Y -> %d\n", r.startPos.X, r.startPos.Y)
	fmt.Printf("Robot current location: X -> %d, Y -> %d\n", r.loc.X, r.loc.Y)
	fmt.Printf("Robot velocity: X -> %d, Y -> %d\n", r.vel.X, r.vel.Y)
	fmt.Printf("Known rows: %d, cols: %d\n", r.rows, r.cols)
}

func (r Robot) Loc() utils.Coord {
	return r.loc
}

func (r *Robot) March(iter int) {
	// This moves the robot as many times as the iteration states
	for i := 0; i < iter; i++ {
		r.loc = utils.Coord{
			X: utils.EuclideanModulus(r.loc.X+r.vel.X, r.cols),
			Y: utils.EuclideanModulus(r.loc.Y+r.vel.Y, r.rows),
		}
	}
}

func PrintMap(rows, cols int, robots []*Robot) {
	// Prints the map in the robots' current state. This is kinda inefficient
	// but generally speaking this is just for debugging purposes anyway. also
	// this will look like trash with the challenge data so this should really
	// only be used for the sample input
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			thisLoc := utils.Coord{
				X: c,
				Y: r,
			}

			counter := 0
			for _, v := range robots {
				if v.loc == thisLoc {
					counter++
				}

				// also for sanity
				if v.rows != rows || v.cols != cols {
					panic(fmt.Sprintf("robots dont have same cols and rows: r.c = %d, r.r = %d, cols = %d, rows = %d", v.cols, v.rows, cols, rows))
				}
			}

			if counter == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", counter)
			}
		}
		fmt.Printf("\n")
	}
}

func PrintTree(rows, cols int, robots []*Robot) {
	// this version doesn't have spaces between, but otherwise is the same thing.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			thisLoc := utils.Coord{
				X: c,
				Y: r,
			}

			counter := 0
			for _, v := range robots {
				if v.loc == thisLoc {
					counter++
				}

				// also for sanity
				if v.rows != rows || v.cols != cols {
					panic(fmt.Sprintf("robots dont have same cols and rows: r.c = %d, r.r = %d, cols = %d, rows = %d", v.cols, v.rows, cols, rows))
				}
			}

			if counter == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}
