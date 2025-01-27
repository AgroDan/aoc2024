package logic

import (
	"sort"
	"strconv"
	"strings"
)

// Here are some helper functions that will help me fiddle with the inputs and outputs.

// zero fill, like what python does
func ZFill(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return strings.Repeat("0", width-len(s)) + s
}

func SetXY(x, y string, regSize int) map[string]*Wire {
	// This will set the X and Y values in a new wirelist
	// equal to the amount of registers specified
	// xList := make([]string, regSize)
	// yList := make([]string, regSize)

	outWire := make(map[string]*Wire)

	x = ZFill(x, regSize)
	y = ZFill(y, regSize)

	for i := 0; i < regSize; i++ {
		alphaDigit := strconv.Itoa(i)
		outWire["x"+ZFill(alphaDigit, 2)] = NewWire("x"+ZFill(alphaDigit, 2), true, x[i] == '1')
		outWire["y"+ZFill(alphaDigit, 2)] = NewWire("y"+ZFill(alphaDigit, 2), true, y[i] == '1')
	}

	return outWire
}

type LogicGate struct {
	reg1, reg2 string
	op         string
	result     string
}

func NewLogicGate(instruction string) *LogicGate {
	parts := strings.Fields(instruction)
	return &LogicGate{
		reg1:   parts[0],
		reg2:   parts[2],
		op:     parts[1],
		result: parts[4],
	}
}

func (lg LogicGate) TrueUp(wireList map[string]*Wire) {
	// This looks at all the registers in the logic gate and
	// confirms they exist in the wireList, and if they do not
	// then add them with "Ready" equal to false.
	if wireList[lg.reg1] == nil {
		wireList[lg.reg1] = NewWire(lg.reg1, false, false)
	}
	if wireList[lg.reg2] == nil {
		wireList[lg.reg2] = NewWire(lg.reg2, false, false)
	}
	if wireList[lg.result] == nil {
		wireList[lg.result] = NewWire(lg.result, false, false)
	}
}

func InputAndExecute(x, y string, challengeText string) string {
	// This will do all the overhead of creating the wirelist, setting the instructions,
	// and returning the result of the Z wires in the wirelist after executing them.

	// First, parse the challenge text, if only to get the size of the registers!
	wireList, instructionQueue := ParseGates(challengeText)
	instructionList := ParseOnlyInstructions(challengeText)

	// Get length of X registers
	count := 0
	for k := range wireList {
		if strings.HasPrefix(k, "x") {
			count++
		}
	}

	// Now throw out the wirelist and replace it with the X and Y values
	newWireList := SetXY(x, y, count)

	logicGates := make([]*LogicGate, 0)
	for _, v := range instructionList {
		newGate := NewLogicGate(v)
		newGate.TrueUp(newWireList)
		logicGates = append(logicGates, newGate)
	}
	// fmt.Printf("NewWireList: %v\n", newWireList)
	// fmt.Printf("OldWireList: %v\n", wireList)

	ParseInstructions(newWireList, instructionQueue)

	zWires := make([]string, 0)
	for k := range newWireList {
		if strings.HasPrefix(k, "z") {
			zWires = append(zWires, k)
		}
	}

	sort.Strings(zWires)
	outBin := ""
	for _, v := range zWires {
		if newWireList[v].Value {
			outBin = "1" + outBin
		} else {
			outBin = "0" + outBin
		}
	}
	return outBin
}
