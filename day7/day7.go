package main

import (
	"bufio"
	equation "day7/Equation"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here
	// var myEquations []Equation
	var partOneTotal int = 0
	var partTwoTotal int = 0
	for i := range lines {
		myEquation := equation.NewEquation(lines[i])
		if myEquation.IsValid() {
			partOneTotal += myEquation.Answer
		}
		if myEquation.IsValidPartTwo() {
			partTwoTotal += myEquation.Answer
		}
	}
	fmt.Printf("Part one total: %d\n", partOneTotal)
	fmt.Printf("Part two total: %d\n", partTwoTotal)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
