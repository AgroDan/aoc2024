package gates

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"utils"
)

// This will consist of functions used to test values against the wires before
// they are sent to the gate. This will take the value of a pre-generated set
// of wires and manipulate just the wires specific to the challenge, seeing as
// they all have to be the same amount according to the provided input file.

func SetWires(letter, value string, wl WireList) WireList {
	// this will set the wires to the value given in the challenge.
	// remember that it is big-endian, so the least significant bit
	// at the end of the string.
	retWires := wl.DeepCopy()

	// no time for nonsense
	if letter != "x" && letter != "y" && letter != "z" {
		panic("Invalid letter")
	}

	assocWires := make([]string, 0)
	for k := range retWires {
		if strings.HasPrefix(k, letter) {
			assocWires = append(assocWires, k)
		}
	}
	newVal := utils.ZFill(value, len(assocWires))

	if len(newVal) != len(assocWires) {
		panic("wire size mismatch of value")
	}

	newVal = utils.ReverseString(newVal)

	// sort the wires
	sort.Slice(assocWires, func(i, j int) bool {
		return assocWires[i] < assocWires[j]
	})

	for i := 0; i < len(assocWires); i++ {
		setBit := string(newVal[i])
		if setBit != "0" && setBit != "1" {
			panic("invalid bit")
		}
		retWires[assocWires[i]].Value = setBit
		retWires[assocWires[i]].Ready = true
	}
	return retWires
}

func ValidateAND(valX, valY, result string) bool {
	// This function will take the values of valX and valY and AND them, and compare the
	// result to the provided result. If they match, then return true, otherwise false.
	// The values of valX, valY, and result are all strings of 1s and 0s.
	if len(valX) != len(valY) || len(valY) != len(result) {
		panic("mismatched lengths")
	}

	for i := 0; i < len(valX); i++ {
		if valX[i] == '1' && valY[i] == '1' {
			if result[i] != '1' {
				return false
			}
		} else {
			if result[i] != '0' {
				return false
			}
		}
	}
	return true
}

// func ValidateADD(valX, valY, result string) bool {
// 	// this function, similar to validateAND, will take the values of valX and valY and
// 	// add them together and compare the result to the provided result. If they match,
// 	// return true
// 	if len(valX) != len(valY) {
// 		panic("mismatched lengths")
// 	}

// 	carry := 0

// 	// Don't forget to pad the values since the result is usually 1 bit higher
// 	diff := len(result) - len(valX)
// 	for i := 0; i < diff; i++ {
// 		valX = "0" + valX
// 		valY = "0" + valY
// 	}

// 	for i := 0; i < len(valX); i++ {
// 		x, _ := strconv.Atoi(string(valX[i]))
// 		y, _ := strconv.Atoi(string(valY[i]))
// 		r, _ := strconv.Atoi(string(result[i]))

// 		sum := x + y + carry
// 		if sum%2 != r {
// 			return false
// 		}
// 		carry = sum / 2
// 	}
// 	return true
// }

func ValidateADD(valX, valY, result string) bool {
	// This simply takes the values of X, Y, and the result
	// and ensures X and Y add to that result.
	x, _ := strconv.ParseInt(valX, 2, 64)
	y, _ := strconv.ParseInt(valY, 2, 64)
	r, _ := strconv.ParseInt(result, 2, 64)
	return x+y == r
}

func ValidateAllBitsAND(wires WireList, gateList []*Gate) []string {
	// This function will iterate through all bits and return the
	// bit positions that failed the AND validation.
	failedBits := make([]string, 0)
	xLen, _, _ := wires.GetXYZBitlength()
	for i := 0; i < xLen; i++ {
		testVal := strings.Repeat("0", xLen-i-1) + "1" + strings.Repeat("0", i)

		testWires := SetWires("x", testVal, wires)
		testWires = SetWires("y", testVal, testWires)
		testExec := Execute(gateList, testWires)
		testValXStr, _ := testExec.GetWires("x")
		testValYStr, _ := testExec.GetWires("y")
		testValStr, _ := testExec.GetWires("z")
		if !ValidateAND(testValXStr, testValYStr, testValStr) {
			failedBits = append(failedBits, fmt.Sprintf("z%02d", i))
		}
	}
	return failedBits
}

func ValidateAllBitsADD(wires WireList, gateList []*Gate) []string {
	// This function will iterate through all the bits and return the
	// bit positions that failed the ADD validation. This also accounts
	// for the carry bit as well.
	failedBits := make([]string, 0)
	xLen, _, _ := wires.GetXYZBitlength()
	for i := 0; i < xLen; i++ {
		// fmt.Printf("Testing bit %d\n", i)
		testVal := strings.Repeat("0", xLen-i-1) + "1" + strings.Repeat("0", i)
		// yVal := strings.Repeat("0", xLen)
		yVal := testVal
		// fmt.Printf("XValue: %s\n", testVal)
		// fmt.Printf("YValue: %s\n", yVal)
		testWires := SetWires("x", testVal, wires)
		testWires = SetWires("y", yVal, testWires)
		testExec := Execute(gateList, testWires)
		testValXStr, _ := testExec.GetWires("x")
		// fmt.Printf("Value of wire x: %s\n", testValXStr)
		testValYStr, _ := testExec.GetWires("y")
		// fmt.Printf("Value of wire y: %s\n", testValYStr)
		testValStr, _ := testExec.GetWires("z")
		// fmt.Printf("Value of wire z: %s\n", testValStr)
		if !ValidateADD(testValXStr, testValYStr, testValStr) {
			failedBits = append(failedBits, fmt.Sprintf("z%02d", i))
		}
	}
	return failedBits
}

func SwapOutputsIdx(i, j int, gates []*Gate) {
	// This will swap the outputs of the ith and jth gate.
	if i < 0 || i >= len(gates) || j < 0 || j >= len(gates) {
		panic("invalid index")
	}
	gates[i].output, gates[j].output = gates[j].output, gates[i].output
}

func SwapOutputs(i, j *Gate) {
	// This will swap the outputs of the two gates.
	i.output, j.output = j.output, i.output
}

func ValidateWiresAND(wires WireList, gateList []*Gate) bool {
	// Given the wires and the instructions, this will simply test to see
	// if the logic gates are arranged in such a way that the wires will
	// properly "AND" together. To accomplish this, I will send all 1s,
	// and all 0's.
	xLen, _, _ := wires.GetXYZBitlength()
	allOnes := strings.Repeat("1", xLen)
	allZeros := strings.Repeat("0", xLen)

	testOnes := SetWires("x", allOnes, wires)
	testOnes = SetWires("y", allOnes, testOnes)
	testOnesExec := Execute(gateList, testOnes)
	testOnesXStr, _ := testOnesExec.GetWires("x")
	testOnesYStr, _ := testOnesExec.GetWires("y")
	testOnesZstr, _ := testOnesExec.GetWires("z")
	if !ValidateAND(testOnesXStr, testOnesYStr, testOnesZstr) {
		return false
	}

	testZeroes := SetWires("x", allZeros, wires)
	testZeroes = SetWires("y", allZeros, testZeroes)
	testZeroesExec := Execute(gateList, testZeroes)
	testZeroesXStr, _ := testZeroesExec.GetWires("x")
	testZeroesYStr, _ := testZeroesExec.GetWires("y")
	testZeroesZstr, _ := testZeroesExec.GetWires("z")
	if !ValidateAND(testZeroesXStr, testZeroesYStr, testZeroesZstr) {
		return false
	}
	return true
}

func ValidateWiresADD(wires WireList, gateList []*Gate) bool {
	// Similar to ValidateWiresAND, this will test the wires to see if they
	// are arranged in such a way that the wires will properly "ADD" together.
	// To accomplish this, I will send half 0s and 1s added by half 1s and zeros,
	// then all 1s and one 1.
	xLen, _, _ := wires.GetXYZBitlength()
	halfOnes := strings.Repeat("0", xLen/2) + strings.Repeat("1", xLen/2)
	halfZeroes := strings.Repeat("1", xLen/2) + strings.Repeat("0", xLen/2)
	allOnes := strings.Repeat("1", xLen)
	oneOne := strings.Repeat("0", xLen-1) + "1"

	testHalfAndHalf := SetWires("x", halfOnes, wires)
	testHalfAndHalf = SetWires("y", halfZeroes, testHalfAndHalf)
	testHalfAndHalfExec := Execute(gateList, testHalfAndHalf)
	testHalfAndHalfXStr, _ := testHalfAndHalfExec.GetWires("x")
	testHalfAndHalfYStr, _ := testHalfAndHalfExec.GetWires("y")
	testHalfAndHalfZstr, _ := testHalfAndHalfExec.GetWires("z")
	if !ValidateADD(testHalfAndHalfXStr, testHalfAndHalfYStr, testHalfAndHalfZstr) {
		return false
	}

	testAllOnes := SetWires("x", allOnes, wires)
	testAllOnes = SetWires("y", oneOne, testAllOnes)
	testAllOnesExec := Execute(gateList, testAllOnes)
	testAllOnesXStr, _ := testAllOnesExec.GetWires("x")
	testAllOnesYStr, _ := testAllOnesExec.GetWires("y")
	testAllOnesZstr, _ := testAllOnesExec.GetWires("z")
	if !ValidateADD(testAllOnesXStr, testAllOnesYStr, testAllOnesZstr) {
		return false
	}
	return true
}

func TryBitCombinationsAND(badBits []string, wires WireList, gateList []*Gate) []string {
	// This will take the failed bits and try to swap them around to see if any
	// of them properly validates the AND operation. Knowing that, in the test
	// instance, we have 4 gates total that need to be swapped (two pairs), we
	// can get every permutation of the known bad bits and swap them around.
	// This will be a lot simpler than the actual challenge, but I'm going to
	// approach this iteratively.
	permutations := GeneratePermutations(badBits)

	for _, perm := range permutations {
		if perm[0][0] == perm[1][0] || perm[0][1] == perm[1][1] || perm[0][0] == perm[1][1] || perm[0][1] == perm[1][0] {
			continue
		}

		testGate := GatelistDeepCopy(gateList)
		// pair[0] will be the index of the failed bit, so z+pair[0] will be the wire
		leftGateFirst := GetGateFromOutput(testGate, perm[0][0])
		rightGateFirst := GetGateFromOutput(testGate, perm[0][1])
		SwapOutputs(leftGateFirst, rightGateFirst)

		leftGateSecond := GetGateFromOutput(testGate, perm[1][0])
		rightGateSecond := GetGateFromOutput(testGate, perm[1][1])
		SwapOutputs(leftGateSecond, rightGateSecond)
		if ValidateWiresAND(wires, testGate) {
			return []string{perm[0][0], perm[0][1], perm[1][0], perm[1][1]}
		}
	}
	return nil
}

func TryBitCombinationsADD(badBits []string, wires WireList, gateList []*Gate) []string {
	// similar to the above, but this will work against the "ADD" operation.
	fmt.Printf("Amount of bad bits: %d\n", len(badBits))
	permutations := GeneratePermutationsFourPairs(badBits)
	fmt.Printf("Amount of permutations: %d\n", len(permutations))

	for _, perm := range permutations {
		if perm[0][0] == perm[1][0] || perm[0][1] == perm[1][1] || perm[0][0] == perm[1][1] || perm[0][1] == perm[1][0] {
			continue
		}

		testGate := GatelistDeepCopy(gateList)
		// pair[0] will be the index of the failed bit, so z+pair[0] will be the wire
		leftGateFirst := GetGateFromOutput(testGate, perm[0][0])
		rightGateFirst := GetGateFromOutput(testGate, perm[0][1])
		SwapOutputs(leftGateFirst, rightGateFirst)

		leftGateSecond := GetGateFromOutput(testGate, perm[1][0])
		rightGateSecond := GetGateFromOutput(testGate, perm[1][1])
		SwapOutputs(leftGateSecond, rightGateSecond)

		leftGatesThird := GetGateFromOutput(testGate, perm[2][0])
		rightGatesThird := GetGateFromOutput(testGate, perm[2][1])
		SwapOutputs(leftGatesThird, rightGatesThird)

		leftGatesFourth := GetGateFromOutput(testGate, perm[3][0])
		rightGatesFourth := GetGateFromOutput(testGate, perm[3][1])
		SwapOutputs(leftGatesFourth, rightGatesFourth)

		if ValidateWiresADD(wires, testGate) {
			return []string{perm[0][0], perm[0][1], perm[1][0], perm[1][1], perm[2][0], perm[2][1], perm[3][0], perm[3][1]}
		}
	}
	return nil
}

func ValidateBitsFromIndex(wires WireList, gateList []*Gate, zWire string) bool {
	// This will validate the bits from the index given. This will be used to
	// validate the bits that are affected by the bad bits.

	// given the Zwire, this will determine what bit place to check.
	index, err := strconv.Atoi(zWire[1:])
	if err != nil {
		panic(err)
	}

	xLen, _, _ := wires.GetXYZBitlength()
	xVal := strings.Repeat("0", xLen-index-1) + "1" + strings.Repeat("0", index)
	yVal := strings.Repeat("0", xLen)

	testWires := SetWires("x", xVal, wires)
	testWires = SetWires("y", yVal, testWires)
	testExec := Execute(gateList, testWires)
	testValXStr, _ := testExec.GetWires("x")
	testValYStr, _ := testExec.GetWires("y")
	testValStr, _ := testExec.GetWires("z")
	return ValidateADD(testValXStr, testValYStr, testValStr)
}

func GetAllWires(gates []*Gate) []string {
	// This will return all the wires that are used in the gates.
	wires := make([]string, 0)
	for _, g := range gates {
		if !slices.Contains(wires, g.operand1) {
			wires = append(wires, g.operand1)
		}
		if !slices.Contains(wires, g.operand2) {
			wires = append(wires, g.operand2)
		}
		if !slices.Contains(wires, g.output) {
			wires = append(wires, g.output)
		}
	}
	return wires
}
