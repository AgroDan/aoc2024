package utils

import (
	"bufio"
	"os"
)

// This will parse a file based on how you expect the data.

func GetTextBlob(filename string) (string, error) {
	// This will return a string of the entire file that you can use to
	// cut up as you want. This will not alter the file in any way.
	readFile, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	// Similar to day 13, this takes a whole challenge blob and passes
	// it unadulterated to the parser
	return string(readFile), nil
}

func GetFileLines(filename string) ([]string, error) {
	// This will return a slice of strings, each string being a line from the file.
	// useful for iterating over things where the entire input is the same type, unlike
	// a challenge where half of the file is one type and the other half is completely
	// different.

	readFile, err := os.Open(filename)

	var retval []string

	if err != nil {
		return retval, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}
