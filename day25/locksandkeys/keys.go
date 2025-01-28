package locksandkeys

import (
	"fmt"
	"utils"
)

type Key struct {
	peaks   utils.Runemap
	heights [5]int
}

// No I'll do an isKey() function just in case
// some weirdo edge case appears

func IsKey(r utils.Runemap) bool {
	rawMap := r.GetRaw()
	for x := range rawMap[0] {
		if rawMap[len(rawMap)-1][x] != '#' {
			return false
		}
	}
	return true
}

func NewKey(r utils.Runemap) *Key {
	newKey := Key{}
	newKey.peaks = r

	// first check if this is a key
	if !IsKey(r) {
		panic("Not a key")
	}
	rawMap := r.GetRaw()
	for x := range r.Width() {
		for y := r.Height() - 2; y >= 0; y-- {
			if rawMap[y][x] == '.' {
				newKey.heights[x] = r.Height() - (y + 2)
				break
			}
		}
	}
	return &newKey
}

func (k Key) Print() {
	for _, y := range k.peaks.GetRaw() {
		for _, x := range y {
			fmt.Printf("%c", x)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Heights: %d, %d, %d, %d, %d\n", k.heights[0], k.heights[1], k.heights[2], k.heights[3], k.heights[4])
}
