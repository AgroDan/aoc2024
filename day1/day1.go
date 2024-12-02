package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func appearanceFreq(c int, arr []int) int {
	// Given a number 'c', will return how many
	// times that number appears in the given list.
	// THIS LIST MUST BE SORTED! This is for efficiency.
	var counter int = 0
	for i := 0; i < len(arr); i++ {
		if c == arr[i] {
			counter++
		}

		if c != arr[i] && counter > 0 {
			return counter
		}
	}
	return counter
}

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

	var leftArr, rightArr []int

	for _, line := range lines {
		vals := strings.Fields(line)
		lVal, err := strconv.Atoi(vals[0])
		if err != nil {
			panic("Not a number")
		}
		rVal, err := strconv.Atoi(vals[1])
		if err != nil {
			panic("Not a number")
		}
		leftArr = append(leftArr, lVal)
		rightArr = append(rightArr, rVal)
	}

	// sort the slices
	slices.Sort(leftArr)
	slices.Sort(rightArr)

	// Now add them all up

	var totalPartOne int = 0
	for i := 0; i < len(leftArr); i++ {
		// I could just get Abs of this
		// but that means I have to move my
		// mouse and look up how to do it again
		var diff int = 0
		if leftArr[i] <= rightArr[i] {
			diff = rightArr[i] - leftArr[i]
		} else {
			diff = leftArr[i] - rightArr[i]
		}
		fmt.Printf("%d - %d = %d\n", leftArr[i], rightArr[i], diff)
		totalPartOne += diff
	}
	fmt.Printf("Part 1 total: %d\n", totalPartOne)

	var totalPartTwo int = 0
	for _, val := range leftArr {
		freq := appearanceFreq(val, rightArr)
		totalPartTwo += freq * val
		fmt.Printf("%d * freq of %d = %d\n", val, freq, freq*val)
	}
	fmt.Printf("Part 2 total: %d\n", totalPartTwo)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
