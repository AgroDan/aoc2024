package main

import (
	"day24/logic"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

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
	wires, instQueue := logic.ParseGates(challengeText)
	// for k, v := range wires {
	// 	fmt.Printf("%s: value: %d, ready? %t\n", k, logic.ConvertBoolToInt(v.Value), v.Ready)
	// }
	// fmt.Printf("Instruction queue: %v\n", instQueue)
	logic.ParseInstructions(wires, instQueue)
	partOneAnswer := logic.GetZWires(wires)
	// logic.PrintAllWires(wires)
	fmt.Printf("Answer for part one: %d\n", partOneAnswer)

	// Now I'll do a sanity check on my code.
	// testText, err := utils.GetTextBlob("./day24/testinput2")
	// if err != nil {
	// 	fmt.Println("Fatal:", err)
	// }
	// testWires, testInstQueue := logic.ParseGates(testText)
	// testVal1, testVal2 := 11, 13
	// newQueue := logic.SwapInstructionsFromQueue("x00 AND y00 -> z05", "x05 AND y05 -> z00", testInstQueue)
	// AnotherNewQueue := logic.SwapInstructionsFromQueue("x01 AND y01 -> z02", "x02 AND y02 -> z01", newQueue)
	// fmt.Printf("AnotherNewQueue: %v\n", AnotherNewQueue)
	// val := logic.TestValues(testVal1, testVal2, 0, testWires, AnotherNewQueue)
	// if val {
	// 	fmt.Println("Test passed!")
	// 	fmt.Println("Value found:", logic.GetZWires(testWires))
	// } else {
	// 	fmt.Println("Test failed!")
	// 	fmt.Println("Value found:", logic.GetZWires(testWires))
	// 	fmt.Printf("Tested: %d + %d, expected %d\n", testVal1, testVal2, testVal1+testVal2)
	// }
	// logic.PrintAllWires(testWires)

	testX, testY := "1000", "0"
	fmt.Printf("Test input: %s + %s, output of compute: %s\n", testX, testY, logic.InputAndExecute(testX, testY, challengeText))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
