package main

import (
	"day5/manuals"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()

	// Going to do this differently to parse it nicer
	// readFile, err := os.Open(*filePtr)

	// if err != nil {
	// 	fmt.Println("Fatal:", err)
	// }
	// defer readFile.Close()

	// fileScanner := bufio.NewScanner(readFile)
	// fileScanner.Split(bufio.ScanLines)

	// var lines []string

	// for fileScanner.Scan() {
	//     lines = append(lines, fileScanner.Text())
	// }

	content, err := os.ReadFile(*filePtr)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// convert to string
	text := string(content)

	// Insert code here
	chal := manuals.ParseChallenge(text)
	chal.PrintAll()
	valid := chal.ReturnValidInstructions()

	for i := range valid {
		for j := range valid[i] {
			fmt.Printf("%d - ", valid[i][j])
		}
		fmt.Printf("\n")
	}

	var partOneTotal int = 0
	for i := range valid {
		middle := len(valid[i]) / 2
		fmt.Printf("%d ", valid[i][middle])
		partOneTotal += valid[i][middle]
	}
	fmt.Printf("Total for Part One: %d\n", partOneTotal)

	// Now find the invalid ones
	invalid := chal.ReturnInvalidInstructions()

	var partTwoTotal int = 0
	for i := range invalid {
		res := chal.Instructions.Fix(invalid[i])
		fmt.Printf("Fixed: ")
		for r := range res {
			fmt.Printf("%d ", res[r])
		}
		fmt.Printf("\n")

		middle := len(res) / 2
		partTwoTotal += res[middle]
	}
	fmt.Printf("Total for Part Two: %d\n", partTwoTotal)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
