package stones

import (
	"strings"
)

// This will parse the input and set up the Stoneset.

func Parse(line string) *StoneSet {
	// will parse the line and generate a linked list
	// based on the input
	// fmt.Printf("Using: %s\n", line)
	ts := strings.TrimSpace(line)
	// fmt.Printf("Now: %s\n", ts)
	entries := strings.Split(ts, " ")
	// fmt.Printf("Length: %d\n", len(entries))

	// Get the first entry for the linked list
	thisSet := NewStoneSet(NewStone(entries[0]))
	// fmt.Printf("Found %s\n", entries[0])
	currPtr := thisSet
	for i := 1; i < len(entries); i++ {
		// Create the stone
		// fmt.Printf("Found %s\n", entries[i])
		myStone := NewStone(entries[i])
		currPtr.AddAfter(myStone)
		currPtr = currPtr.Next
	}
	return thisSet
}
