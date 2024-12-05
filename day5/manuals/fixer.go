package manuals

import "slices"

/*
 * Figured I'd give the "fixer" application it's own file, so here it is.
 * A ***MAJOR*** props to /u/Fabianofski for his extremely elegant solution
 * that really helped me understand the best approach here.
 */

func (inst InstructionSet) Fix(pages []int) []int {
	// This, given a pages slice, will order it appropriately in accordance
	// with the supplied ruleset.
	for i := 0; i < len(pages); i++ {
		// Now loop through the instruction set associated with the number, if any
		for _, is := range inst.por[pages[i]] {

			// Get a list of all preceeding numbers
			prev := pages[:i]

			// This gives the index where the number _should be behind_
			idx := slices.Index(prev, is)

			if idx != -1 {
				// meaning we found the index of a preceeding number
				pages = slices.Delete(pages, idx, idx+1)
				pages = slices.Insert(pages, i, is)
				i = 0 // start the loop over
				break
			}
		}
	}
	return pages
}
