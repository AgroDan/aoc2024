package manuals

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * This is just the parser, will ingest data from the provided
 * input in accordance with the challenge specifications.
 */

type InstructionSet struct {
	// por means Page Ordering Rules
	por map[int][]int
}

func NewInstructionSet() *InstructionSet {
	// sets up the instructionset
	return &InstructionSet{
		por: make(map[int][]int),
	}
}

func (i *InstructionSet) ParsePageOrder(p string) {
	// This will parse a page that looks like this:
	// 57|20
	// and add it to the por K:V set. Ultimately that
	// means that page 57 in the above case has a rule
	// that means it will come _before_ page 20.
	trimmedInst := strings.TrimSpace(p)
	instSlice := strings.Split(trimmedInst, "|")

	// Now to ingest as numbers
	var numSlice []int
	for _, v := range instSlice {
		extractedNum, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numSlice = append(numSlice, extractedNum)
	}

	// Now add the rule
	i.por[numSlice[0]] = append(i.por[numSlice[0]], numSlice[1])
}

// Now let's create a page set object because why not
type PageSet struct {
	p [][]int
}

func NewPageSet() *PageSet {
	return &PageSet{
		p: make([][]int, 0),
	}
}

func (p *PageSet) ParsePageSet(pages string) {
	// This, like ParsePageOrder, will ingest a string that looks like this:
	// 75,47,61,53,29
	// The order for this is important! It will ensure that everything is in
	// the order listed. This is for a single instruction.
	trimmedPages := strings.TrimSpace(pages)
	pageSlice := strings.Split(trimmedPages, ",")
	var numSlice []int
	for _, v := range pageSlice {
		extractedNum, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numSlice = append(numSlice, extractedNum)
	}

	p.p = append(p.p, numSlice)
}

// Now let's tie the room together
type Manual struct {
	Instructions *InstructionSet
	Pages        *PageSet
}

// Let's parse the whole file as one giant thing

func ParseChallenge(t string) Manual {
	// First, trim whitespace because why not
	trimmedFile := strings.TrimSpace(t)
	newT := strings.Split(trimmedFile, "\n\n") // separate by 2 newlines

	m := Manual{
		Instructions: NewInstructionSet(),
		Pages:        NewPageSet(),
	}
	// Work on the first part
	instLines := strings.Split(newT[0], "\n")
	for _, line := range instLines {
		m.Instructions.ParsePageOrder(line)
	}

	// On the second part
	pageLines := strings.Split(newT[1], "\n")
	for _, line := range pageLines {
		m.Pages.ParsePageSet(line)
	}
	return m
}

func (m Manual) PrintAll() {
	// Just prints everything out to ensure it was ingested properly
	fmt.Printf("Page instructions:\n")
	for k, v := range m.Instructions.por {
		fmt.Printf("%d:", k)
		for _, vv := range v {
			fmt.Printf(" %d", vv)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nPage Sets:\n")
	for _, v := range m.Pages.p {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Printf("\n")
	}
}
