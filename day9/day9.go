package main

import (
	"bufio"
	"day9/disk"
	"flag"
	"fmt"
	"os"
	"strings"
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
	fullLine := strings.Join(lines, "")
	fList := disk.ParseDiskReport(fullLine)
	// for i := range fList {
	// 	fList[i].Print()
	// 	// fmt.Printf("%v+\n", fList[i])
	// }

	fmt.Printf("Defraggin'...\n")
	newList := disk.Defragment(fList)
	var partOneCounter, partOneCheckSum int = 0, 0
	for i := range newList {
		for j := range newList[i].F {
			if newList[i].F[j] == '.' {
				partOneCounter++
				continue
			}
			digit := int(newList[i].F[j] - '0')
			partOneCheckSum += digit * partOneCounter
			partOneCounter++
		}
		// i'll just do the checksum here
		// newList[i].Print()
	}

	partTwoList := disk.ParseDiskReport(fullLine)
	disk.DefragmentPartTwo(&partTwoList)
	var partTwoCounter, partTwoCheckSum int = 0, 0
	for i := range partTwoList {
		for j := range partTwoList[i].F {
			if partTwoList[i].F[j] == '.' {
				partTwoCounter++
				continue
			}
			digit := int(partTwoList[i].F[j] - '0')
			partTwoCheckSum += digit * partTwoCounter
			partTwoCounter++
		}
		// i'll just do the checksum here
		// newList[i].Print()
	}

	fmt.Printf("Answer to part 1: %d\n", partOneCheckSum)
	fmt.Printf("Answer to part 2: %d\n", partTwoCheckSum)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
