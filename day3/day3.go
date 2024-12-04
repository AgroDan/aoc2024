package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseNoTry(startIdx int, instString string) (bool, error) {
	// This will parse do() or don't(). Returns "true" if do(),
	// "false" if don't(). Will return an error if neither.
	var endIdx int
	if startIdx+7 >= len(instString) {
		endIdx = len(instString) - 1
	} else {
		endIdx = startIdx + 7
	}

	workingString := instString[startIdx:endIdx]
	pattern := `^(do\(\)|don\'t\(\))`

	re := regexp.MustCompile(pattern)
	if matches := re.FindStringSubmatch(workingString); matches != nil {
		if matches[1] == "do()" {
			return true, nil
		}
		if matches[1] == "don't()" {
			return false, nil
		}

		// Otherwise WTF does it even say???
		return false, fmt.Errorf("not sure how this got here: %s", matches[1])
	}
	return false, errors.New("neither action found")
}

func ParseMul(startIdx int, instString string) (int, int, error) {
	// Uses regex to do this properly. The most this can check for is
	// 12 characters after the start index according to the rules of the
	// challenge, so let's determine that.
	var endIdx int
	if startIdx+12 >= len(instString) {
		endIdx = len(instString) - 1
	} else {
		endIdx = startIdx + 12
	}
	workingString := instString[startIdx:endIdx]
	// fmt.Printf("Working string: %s\n", workingString)
	pattern := `^mul\((\d{1,3}),(\d{1,3})\)`

	re := regexp.MustCompile(pattern)
	if matches := re.FindStringSubmatch(workingString); matches != nil {
		firstNum, err := strconv.Atoi(matches[1])
		if err != nil {
			return -1, -1, err
		}
		secondNum, err := strconv.Atoi(matches[2])
		if err != nil {
			return -1, -1, err
		}

		// fmt.Printf("Pass\n")
		// fmt.Printf("Working string: %s\n", workingString)
		return firstNum, secondNum, nil
	} else {
		// fmt.Printf("Fail\n")
		return -1, -1, errors.New("invalid number bracket")
	}

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
	var totalCalcPartOne, totalCalcPartTwo int
	var currState bool = true
	workingStr := strings.Join(lines, "") // one giant line pls
	for i := 0; i < len(workingStr); i++ {
		checkStatus, err := ParseNoTry(i, workingStr)
		if err == nil {
			currState = checkStatus
		}

		firstNum, secondNum, err := ParseMul(i, workingStr)
		if err != nil {
			// no mul in this string
			continue
		}

		totalCalcPartOne += (firstNum * secondNum)

		if currState {
			totalCalcPartTwo += (firstNum * secondNum)
		}
	}

	fmt.Printf("Calculation for Stage 1: %d\n", totalCalcPartOne)
	fmt.Printf("Calculation for Stage 2: %d\n", totalCalcPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
