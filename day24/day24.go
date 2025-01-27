package main

import (
	"day24/gates"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here
	workingWire := flag.String("w", "z25", "Working wire to analyze")

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	challengeText, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// Insert code here
	fmt.Printf("Parsing challenge...\n")
	wires, gateList := gates.ParseChallenge(challengeText)
	fmt.Printf("Executing gates...\n")
	afterExecution := gates.Execute(gateList, wires)
	fmt.Printf("Getting wire values...\n")
	valXStr, _ := afterExecution.GetWires("x")
	valYStr, _ := afterExecution.GetWires("y")
	valStr, valInt := afterExecution.GetWires("z")
	fmt.Printf("Value of wire x: %s\n", valXStr)
	fmt.Printf("Value of wire y: %s\n", valYStr)
	fmt.Printf("Value of wire z: %s, or in decimal: %d\n", valStr, valInt)

	// Now I'm going to try and re-execute by filling a 1 in for each digit and validating
	// testExec := gates.Execute(gateList, wires)
	// xLen, _, _ := wires.GetXYZBitlength()
	// for i := 0; i < xLen; i++ {
	// 	fmt.Printf("Testing bit %d\n", i)
	// 	testVal := strings.Repeat("0", xLen-i-1) + "1" + strings.Repeat("0", i)

	// 	// fmt.Printf("Test value: %s\n", testVal)
	// 	testWires := gates.SetWires("x", testVal, wires)
	// 	testWires = gates.SetWires("y", testVal, testWires)
	// 	testExec := gates.Execute(gateList, testWires)
	// 	testValXStr, _ := testExec.GetWires("x")
	// 	testValYStr, _ := testExec.GetWires("y")
	// 	testValStr, _ := testExec.GetWires("z")
	// 	// fmt.Printf("Value of wire x: %s\n", testValXStr)
	// 	// fmt.Printf("Value of wire y: %s\n", testValYStr)
	// 	// fmt.Printf("Value of wire z: %s, or in decimal: %d\n", testValStr, testValInt)
	// 	if !gates.ValidateAND(testValXStr, testValYStr, testValStr) {
	// 		fmt.Printf("Validation mistmatch at bit %d\n", i)
	// 	}
	// }

	// bitsInvalid := gates.ValidateAllBitsAND(wires, gateList)
	// fmt.Printf("Invalid bits: %v\n", bitsInvalid)

	// fmt.Printf("Trying to fix...\n")
	// fixed := gates.TryBitCombinationsAND(bitsInvalid, wires, gateList)
	// fmt.Printf("Fixed: %v\n", fixed)

	badBits := gates.ValidateAllBitsADD(wires, gateList)
	fmt.Printf("Invalid bits: %v\n", badBits)

	// fmt.Printf("Obtaining all affected gates...\n")
	// badGates := make([]string, 0)
	// uniqueBadGates := make(map[string]int)
	// for _, b := range badBits {
	// 	selectedGate := gates.GetGateFromOutput(gateList, b)
	// 	affectedGates := gates.FollowGates(gateList, selectedGate, b)

	// 	for _, v := range affectedGates {
	// 		if !slices.Contains(badGates, v) {
	// 			// badGates = append(badGates, v)
	// 			uniqueBadGates[v]++
	// 		}
	// 	}
	// }

	// for k, v := range uniqueBadGates {
	// 	if v == 1 {
	// 		badGates = append(badGates, k)
	// 	}
	// }

	// fmt.Printf("Unique bad gates: %v\n", badGates)
	// fmt.Printf("Fixing...\n")
	// allWires := gates.GetAllWires(gateList)
	// for _, bg := range badBits {
	// 	// Get the wires leading up to this bad bit
	// 	leadWires := gates.FollowGates(gateList, gates.GetGateFromOutput(gateList, bg), wires, bg)
	// 	fmt.Printf("Fixing gate %s\n", bg)
	// 	for _, lw := range leadWires {
	// 		if !slices.Contains(badGates, lw) {
	// 			// only care if this leadwire is in the known unique bad gates
	// 			continue
	// 		}

	// 		for _, w := range allWires {
	// 			testGate := gates.GatelistDeepCopy(gateList)
	// 			leftGate := gates.GetGateFromOutput(testGate, lw)
	// 			rightGate := gates.GetGateFromOutput(testGate, w)
	// 			if leftGate == nil || rightGate == nil {
	// 				continue
	// 			}
	// 			gates.SwapOutputs(leftGate, rightGate)
	// 			if gates.ValidateBitsFromIndex(wires, testGate, bg) {
	// 				fmt.Printf("I fixed a gate! by swapping %s and %s\n", lw, w)
	// 			}
	// 		}
	// 	}
	// }

	// fixed := gates.TryBitCombinationsADD(badGates, wires, gateList)

	// fmt.Printf("Fixed: %v\n", fixed)

	// Let's get a list of all the gates leading up to
	impactedGates := gates.GetImpactedGates(gateList, *workingWire)
	fmt.Printf("Impacted gates for %s:\n", *workingWire)
	gates.PrintImpactedGates(gateList, impactedGates)

	/*
		This is my template for what is considered a valid Adder circuit.
		Generally speaking, for each bit, it cascades down until it reaches
		z00, which is the first. So for each bit, it should compare itself
		with the previous bit's logic circuit, and then ONLY print the
		difference between them. Based on that, this is a standard pattern
		for a WORKING circuit, in this case z25:

		x25 XOR y25 -> vnt
		hfm OR wwm -> kcd
		tcj XOR swd -> z24
		y24 AND x24 -> hfm
		tcj AND swd -> wwm
		kcd XOR vnt -> z25

		So in this case for z25, the following needs to exist:
		- XOR gate for the current number
		- AND gate for the previous number
		- an OR gate
		- another AND gate
		- another XOR gate

		after doing it by hand, I believe this is the order:
		z10, vcf z17 fhg fsq dvb tnc z39
	*/
	vals := []string{"z10", "vcf", "z17", "fhg", "fsq", "dvb", "tnc", "z39"}
	slices.Sort(vals)
	fmt.Printf("Order: %s\n", strings.Join(vals, ","))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
