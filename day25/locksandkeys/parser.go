package locksandkeys

import (
	"strings"
	"utils"
)

// First we'll convert split up the input file by splitting on double newlines,
// then split on newlines. Convert each line into a runemap and work out exactly
// what we're looking at.

// we're expecting a gigantic blob of strings
func ParseInput(challengetext string) ([]Key, []Lock) {
	keys := []Key{}
	locks := []Lock{}

	// split the text blob into blocks
	blocks := strings.Split(challengetext, "\n\n")

	// Now for each block, split into lines, and cover for any extra whitespace
	for _, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		runemap := utils.NewRunemap(lines)

		// Now we need to determine if this is a key or a lock
		if IsKey(runemap) {
			keys = append(keys, *NewKey(runemap))
		} else if IsLock(runemap) {
			locks = append(locks, *NewLock(runemap))
		}
	}

	return keys, locks
}
