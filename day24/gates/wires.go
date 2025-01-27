package gates

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// This will hold the values of all the gates.
// If they do not have a signal, then the "ready" field
// will be false.

type Wire struct {
	Name  string
	Value string // This is easier to work with if we just stringify "1" or "0"
	Ready bool
}

func NewWire(name, value string) *Wire {
	return &Wire{
		Name:  name,
		Value: value,
		Ready: value != "",
	}
}

type WireList map[string]*Wire

// this will hold the collection of gates.

func (w WireList) GetWires(letter string) (string, int) {
	// This will return the values of the wire in accordance with the challenge
	// stipulation, where the least signficant bit is in the 0th index, so count
	// backward to give the actual value.
	if letter != "x" && letter != "y" && letter != "z" {
		return "", 0
	}

	retstr := ""
	wires := make([]string, 0)
	for k := range w {
		if strings.HasPrefix(k, letter) {
			wires = append(wires, k)
		}
	}

	// now sort the wires
	sort.Slice(wires, func(i, j int) bool {
		return wires[i] > wires[j]
	})

	// slices.Reverse(wires)
	for i := range wires {
		retstr += w[wires[i]].Value
	}

	if retstr == "" {
		return "", 0
	}

	numVal, err := strconv.ParseInt(retstr, 2, 64)
	if err != nil {
		fmt.Println("Fatal:", err)
	}
	return retstr, int(numVal)
}

func (w WireList) GetXYZBitlength() (int, int, int) {
	// This will return the bitlength of the x, y, and z wires
	// in that order.
	x, _ := w.GetWires("x")
	y, _ := w.GetWires("y")
	z, _ := w.GetWires("z")
	return len(x), len(y), len(z)
}

func (w WireList) PrintAllWires() {
	for k, v := range w {
		fmt.Printf("Wire: %s, Value: %s, Ready?: %t\n", k, v.Value, v.Ready)
	}
}

func (w WireList) DeepCopy() WireList {
	// This will create a deep copy of the wirelist
	newList := make(WireList)
	for k, v := range w {
		newList[k] = NewWire(v.Name, v.Value)
	}
	return newList
}
