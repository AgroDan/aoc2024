package main

import (
	"bufio"
	"day2/Reports"
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
	var r []Reports.ReportObject
	for _, v := range lines {
		r = append(r, Reports.NewReportObject(v))
	}

	var totalPartOne int = 0
	var totalPartTwo int = 0
	for _, v := range r {
		v.PrintReport()
		if v.Safe() {
			totalPartOne++
		}
		if v.ProblemDampenerContingent() {
			totalPartTwo++
		}
	}
	fmt.Printf("Total safe entries: %d\n", totalPartOne)
	fmt.Printf("Total safe entries after contingent: %d\n", totalPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
