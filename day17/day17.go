package main

import (
	"day17/computer"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()

	readFile, err := os.ReadFile(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// Similar to day 13, this takes a whole challenge blob and passes
	// it unadulterated to the parser
	challengeText := string(readFile)

	// Insert code here
	myComputer := computer.ParseProgram(challengeText)
	myComputer.Print()

	myComputer.Run()
	myComputer.PrintOutput()

	// now count until we hit something...

	// var counter int = 627674170 << 3
	// for {
	// 	attempt := myComputer.GenCompute(counter)
	// 	if computer.ValidCandidate(myComputer.Program, attempt) {
	// 		break
	// 	}
	// 	counter++
	// }
	// fmt.Printf("Discovered %d to achieve copy of same program.\n", counter)

	myComputer.Solver()
	// x := []int{5, 3, 9, 0}
	// y := []int{0}
	// fmt.Printf("Are we equal? %t\n", computer.ValidCandidate(x, y))
	// guess := myComputer.Initialize(627674170)
	// guess.Run()
	// guess.PrintOutput()

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
