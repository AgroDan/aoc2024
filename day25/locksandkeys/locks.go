package locksandkeys

import (
	"fmt"
	"utils"
)

// First we'll convert each lock/key into a runemap
type Lock struct {
	tumbler utils.Runemap
	heights [5]int
}

// this is all we need right? I don't need a "isKey()" function
func IsLock(r utils.Runemap) bool {
	rawMap := r.GetRaw()
	for x := range rawMap[0] {
		if rawMap[0][x] != '#' {
			return false
		}
	}
	return true
}

func NewLock(r utils.Runemap) *Lock {
	newLock := Lock{}
	newLock.tumbler = r

	// first check if this is a lock
	if !IsLock(r) {
		panic("Not a lock")
	}
	rawMap := r.GetRaw()
	for x := range r.Width() {
		for y := 1; y < r.Height(); y++ {
			if rawMap[y][x] == '.' {
				newLock.heights[x] = y - 1
				break
			}
		}
	}
	return &newLock
}

func (l Lock) Print() {
	for _, y := range l.tumbler.GetRaw() {
		for _, x := range y {
			fmt.Printf("%c", x)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Tumblers: %d, %d, %d, %d, %d\n", l.heights[0], l.heights[1], l.heights[2], l.heights[3], l.heights[4])
}
