package computer

import (
	"strconv"
	"strings"
)

func ParseProgram(challenge string) Computer {
	// this will parse the program, because of the weird formatting
	// I'll expect a giant blob of data to be input instead of line-by-line
	thisComputer := Computer{
		instPtr: 0,
		Output:  []int{},
	}
	sections := strings.Split(challenge, "\n\n")

	registers := strings.Split(sections[0], "\n")

	regA := strings.Split(registers[0], ":")
	regANum, _ := strconv.Atoi(strings.TrimSpace(regA[1]))

	regB := strings.Split(registers[1], ":")
	regBNum, _ := strconv.Atoi(strings.TrimSpace(regB[1]))

	regC := strings.Split(registers[2], ":")
	regCNum, _ := strconv.Atoi(strings.TrimSpace(regC[1]))

	thisComputer.A, thisComputer.B, thisComputer.C = regANum, regBNum, regCNum

	// Now the program
	prog := strings.Split(sections[1], ":")
	progNums := strings.Split(strings.TrimSpace(prog[1]), ",")
	for i := range progNums {
		num, _ := strconv.Atoi(progNums[i])
		thisComputer.Program = append(thisComputer.Program, num)
	}
	return thisComputer
}
