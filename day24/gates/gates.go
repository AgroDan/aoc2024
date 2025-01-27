package gates

import (
	"fmt"
	"strconv"
	"strings"
)

// This will read an instruction and return the Gate type.

type Gate struct {
	operand1, operand2 string
	operator           string
	output             string
}

func NewGate(instLine string) *Gate {
	// The line will look like aaa XOR bbb -> ccc
	operations := strings.Fields(instLine)
	return &Gate{
		operand1: operations[0],
		operand2: operations[2],
		operator: operations[1],
		output:   operations[4],
	}
}

func (g Gate) String() string {
	return g.operand1 + " " + g.operator + " " + g.operand2 + " -> " + g.output
}

func (g Gate) Operate(w WireList) bool {
	// This will output a true if the operation were successful.
	// what denotes unsuccessful is if any of the operands are not ready.
	if !w[g.operand1].Ready || !w[g.operand2].Ready {
		return false
	}

	switch g.operator {
	case "AND":
		if w[g.operand1].Value == "1" && w[g.operand2].Value == "1" {
			w[g.output].Value = "1"
		} else {
			w[g.output].Value = "0"
		}
		w[g.output].Ready = true
	case "OR":
		if w[g.operand1].Value == "1" || w[g.operand2].Value == "1" {
			w[g.output].Value = "1"
		} else {
			w[g.output].Value = "0"
		}
		w[g.output].Ready = true
	case "XOR":
		if w[g.operand1].Value != w[g.operand2].Value {
			w[g.output].Value = "1"
		} else {
			w[g.output].Value = "0"
		}
		w[g.output].Ready = true
	default:
		panic("Invalid operator")
	}
	return true
}

func GatelistDeepCopy(g []*Gate) []*Gate {
	// This will return a deep copy of the gate list
	ret := make([]*Gate, len(g))
	for i, gate := range g {
		copyGate := &Gate{
			operand1: gate.operand1,
			operand2: gate.operand2,
			operator: gate.operator,
			output:   gate.output,
		}
		ret[i] = copyGate
	}
	return ret
}

// func FollowGates(gates []*Gate, thisGate *Gate, wires WireList, start string) []string {
// 	// This will start from the supplied output and return all of the wires
// 	// that have executions performed on them to get to that wire, up until
// 	// it hits an X or Y wire.
// 	ret := make([]string, 0)

// 	// first, find the gate that has the output
// 	ret = append(ret, thisGate.output)

// 	if !strings.HasPrefix(thisGate.operand1, "x") || !strings.HasPrefix(thisGate.operand1, "y") {
// 		op1Gate := GetGateFromOutput(gates, thisGate.operand1)
// 		if op1Gate != nil {
// 			ret = append(ret, FollowGates(gates, op1Gate, wires, op1Gate.output)...)
// 		}
// 	}

// 	if !strings.HasPrefix(thisGate.operand2, "x") || !strings.HasPrefix(thisGate.operand2, "y") {
// 		op2Gate := GetGateFromOutput(gates, thisGate.operand2)
// 		if op2Gate != nil {
// 			ret = append(ret, FollowGates(gates, op2Gate, wires, op2Gate.output)...)
// 		}
// 	}

// 	return ret
// }

func FollowGates(gates []*Gate, thisGate *Gate, start string) []string {
	// Similar to the above, but is less restrictive
	ret := make([]string, 0)

	ret = append(ret, thisGate.output)

	op1Gate := GetGateFromOutput(gates, thisGate.operand1)
	if op1Gate != nil {
		ret = append(ret, FollowGates(gates, op1Gate, op1Gate.output)...)
	}

	op2Gate := GetGateFromOutput(gates, thisGate.operand2)
	if op2Gate != nil {
		ret = append(ret, FollowGates(gates, op2Gate, op2Gate.output)...)
	}

	// fmt.Printf("Finished: %v\n", ret)
	return ret
}

func GetImpactedGates(gates []*Gate, initialGate string) []string {
	// This takes the initialGate and the one before it and returns THE DIFFERENCE in
	// unique gates between them, showing only what makes THAT particular digit.

	// I don't have time for error validation. NO TIME!
	gateNum, _ := strconv.Atoi(initialGate[1:])
	previousGate := fmt.Sprintf("z%02d", gateNum-1)

	thisGate := FollowGates(gates, GetGateFromOutput(gates, initialGate), initialGate)
	// fmt.Printf("Got this gate...\n")
	prevGate := FollowGates(gates, GetGateFromOutput(gates, previousGate), previousGate)

	// get the unique ones
	indivGates := make(map[string]int)
	for _, v := range thisGate {
		indivGates[v]++
	}

	for _, v := range prevGate {
		indivGates[v]++
	}

	// this is ghetto but i'm just so tired
	retVal := make([]string, 0)
	for k, v := range indivGates {
		if v == 1 {
			retVal = append(retVal, k)
		}
	}
	return retVal
}
