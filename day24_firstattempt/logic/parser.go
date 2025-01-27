package logic

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type Wire struct {
	Name  string
	Ready bool
	Value bool
}

func NewWire(name string, isReady bool, value bool) *Wire {
	return &Wire{
		Name:  name,
		Ready: isReady,
		Value: value,
	}
}

func ParseGates(s string) (map[string]*Wire, utils.GQueue[string]) {
	// Going to store the initial values, then store the instructions as items
	// in a queue.

	wireList := make(map[string]*Wire)
	parts := strings.Split(strings.TrimSpace(s), "\n\n")

	partOne := strings.Split(parts[0], "\n")
	for _, line := range partOne {
		vals := strings.Split(line, ":")

		formattedResult := strings.TrimSpace(vals[1])
		if formattedResult == "1" {
			wireList[vals[0]] = NewWire(vals[0], true, true)
		} else {
			wireList[vals[0]] = NewWire(vals[0], true, false)
		}
	}

	// Now for the second half...
	partTwo := strings.Split(parts[1], "\n")
	outQueue := utils.NewGQueue[string]()
	for _, line := range partTwo {
		vals := strings.Fields(line)

		// Create the gate if not ready yet
		if _, exists := wireList[vals[0]]; !exists {
			wireList[vals[0]] = NewWire(vals[0], false, false)
		}

		if _, exists := wireList[vals[2]]; !exists {
			wireList[vals[2]] = NewWire(vals[2], false, false)
		}

		if _, exists := wireList[vals[4]]; !exists {
			wireList[vals[4]] = NewWire(vals[4], false, false)
		}

		// now push the instruction to the queue
		outQueue.Enqueue(line)
	}
	return wireList, outQueue
}

func ParseInstructions(wireList map[string]*Wire, instQueue utils.GQueue[string]) {
	// instructions will look like:
	// x00 AND y00 -> z00
	for !instQueue.IsEmpty() {
		inst, _ := instQueue.Dequeue()
		vals := strings.Fields(inst)

		// If the wire isn't ready, skip this instruction
		if !wireList[vals[0]].Ready {
			instQueue.Enqueue(inst)
			continue
		}

		// If the wire isn't ready, skip this instruction
		if !wireList[vals[2]].Ready {
			instQueue.Enqueue(inst)
			continue
		}

		// If we're here, let's parse the instruction.
		switch vals[1] {
		case "AND":
			wireList[vals[4]].Value = And(wireList[vals[0]].Value, wireList[vals[2]].Value)
			wireList[vals[4]].Ready = true
		case "OR":
			wireList[vals[4]].Value = Or(wireList[vals[0]].Value, wireList[vals[2]].Value)
			wireList[vals[4]].Ready = true
		case "XOR":
			wireList[vals[4]].Value = Xor(wireList[vals[0]].Value, wireList[vals[2]].Value)
			wireList[vals[4]].Ready = true
		}
	}
}

func GetZWires(wireList map[string]*Wire) int {
	// This will return the decimal number of all the z wires, in the order they are received.
	// remember, least significant bit is z00, most being z99 if it exists.
	zWires := make([]string, 0)
	for k := range wireList {
		if strings.HasPrefix(k, "z") {
			zWires = append(zWires, k)
		}
	}

	sort.Strings(zWires)
	outBin := ""
	for _, v := range zWires {
		if wireList[v].Value {
			outBin = "1" + outBin
		} else {
			outBin = "0" + outBin
		}
	}
	decimal, err := strconv.ParseInt(outBin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(decimal)

}

func PrintAllWires(wireList map[string]*Wire) {
	sortableWireList := make([]string, 0)
	for k := range wireList {
		sortableWireList = append(sortableWireList, k)
	}
	sort.Strings(sortableWireList)

	for _, v := range sortableWireList {
		fmt.Printf("%s: value: %d\n", v, ConvertBoolToInt(wireList[v].Value))
	}
}

func ParseOnlyInstructions(s string) []string {
	// This gets rid of all the cruft and just returns the instructions
	// in the order it was received. "s" should be the entire challenge text.
	parts := strings.Split(strings.TrimSpace(s), "\n\n")
	return strings.Split(parts[1], "\n")
}
