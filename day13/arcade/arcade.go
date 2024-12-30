package arcade

import (
	"fmt"
	"strconv"
	"strings"
)

// let's first organize the data

// this is not a coord, even though it's structurally identical
type ArcadeItem struct {
	X, Y int
}

type Machine struct {
	A, B, Prize ArcadeItem
}

func NewMachine(machineblock string) Machine {
	// this is so specific and complex that I'm ashamed.
	// WITNESS MY SHAME

	// split by newline
	macLines := strings.Split(machineblock, "\n")

	// button a, trust me this splits it all up
	buttonAentries := strings.Split(strings.TrimSpace(strings.Split(macLines[0], ":")[1]), ",")

	// get coords
	axData := strings.Split(buttonAentries[0], "+")[1]
	ayData := strings.Split(buttonAentries[1], "+")[1]

	axInt, err := strconv.Atoi(axData)
	if err != nil {
		panic(fmt.Sprintf("could not interpret button a, axdata: %s", axData))
	}

	ayInt, err := strconv.Atoi(ayData)
	if err != nil {
		panic(fmt.Sprintf("could not interpret button a, aydata: %s", ayData))
	}
	ButtonA := ArcadeItem{
		X: axInt,
		Y: ayInt,
	}

	// button b
	buttonBentries := strings.Split(strings.TrimSpace(strings.Split(macLines[1], ":")[1]), ",")

	// get coords
	bxData := strings.Split(buttonBentries[0], "+")[1]
	byData := strings.Split(buttonBentries[1], "+")[1]

	bxInt, err := strconv.Atoi(bxData)
	if err != nil {
		panic("could not interpret button b, bxdata")
	}

	byInt, err := strconv.Atoi(byData)
	if err != nil {
		panic("could not interpret button b, bydata")
	}

	ButtonB := ArcadeItem{
		X: bxInt,
		Y: byInt,
	}

	// prize
	prizeEntries := strings.Split(strings.TrimSpace(strings.Split(macLines[2], ":")[1]), ",")

	// get coords
	prizeXData := strings.Split(prizeEntries[0], "=")[1]
	prizeYData := strings.Split(prizeEntries[1], "=")[1]

	prizeXInt, err := strconv.Atoi(prizeXData)
	if err != nil {
		panic("could not interpret prize X data")
	}

	prizeYInt, err := strconv.Atoi(prizeYData)
	if err != nil {
		panic("could not interpret prize Y data")
	}

	Prize := ArcadeItem{
		X: prizeXInt,
		Y: prizeYInt,
	}

	return Machine{
		A:     ButtonA,
		B:     ButtonB,
		Prize: Prize,
	}
}

func (m Machine) Print() {
	// prints the data
	fmt.Printf("Button A: X + %d, Y + %d\n", m.A.X, m.A.Y)
	fmt.Printf("Button B: X + %d, Y + %d\n", m.B.X, m.B.Y)
	fmt.Printf("Prize: X = %d, Y = %d\n", m.Prize.X, m.Prize.Y)
}

func (m Machine) PartTwoMult() (int, int) {
	// returns the multiplier of part two
	// ten trillion, damn
	return m.Prize.X + 10_000_000_000_000, m.Prize.Y + 10_000_000_000_000
}
