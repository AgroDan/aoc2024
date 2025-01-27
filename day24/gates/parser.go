package gates

import (
	"strings"
	"utils"
)

func ParseChallenge(challengeText string) (WireList, []*Gate) {
	// This will parse the challenge text and return the wire list and the gates.
	wires := make(WireList)
	gates := make([]*Gate, 0)
	for _, line := range strings.Split(challengeText, "\n") {
		if line == "" {
			continue
		}
		if strings.Contains(line, "->") {
			gates = append(gates, NewGate(line))
			// Don't forget to initialize the wires
			// given on the instructions
			parts := strings.Fields(line)

			// some sanity checking first
			w1, w2, w3 := parts[0], parts[2], parts[4]
			if _, ok := wires[w1]; !ok {
				wires[w1] = NewWire(w1, "")
			}
			if _, ok := wires[w2]; !ok {
				wires[w2] = NewWire(w2, "")
			}
			if _, ok := wires[w3]; !ok {
				wires[w3] = NewWire(w3, "")
			}
		} else {
			// This is a wire assignment
			parts := strings.Split(line, ":")
			w := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			wires[w] = NewWire(w, val)
		}
	}
	return wires, gates
}

func Execute(gates []*Gate, wires WireList) WireList {
	// This will execute the operations given to the wires. This does not
	// operate in place and instead returns a new wirelist with the updated
	// values after all operations have performed
	retWires := wires.DeepCopy()

	instQueue := utils.NewGQueue[*Gate]()
	for _, g := range gates {
		instQueue.Enqueue(g)
	}

	for !instQueue.IsEmpty() {
		gate, _ := instQueue.Dequeue()
		if !gate.Operate(retWires) {
			// fmt.Printf("Gate %s not ready\n", gate.output)
			// If the operation was successful, then enqueue the gate
			// again to check if the next operation can be performed.
			instQueue.Enqueue(gate)
		}
	}
	return retWires
}
