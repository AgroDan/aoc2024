package stones

import (
	"errors"
	"fmt"
)

const (
	DENULL = iota
	SPLIT
	MULT
)

// This will define a stoneset. This is a slice of pointers
// pointing at stones. It allows for stones to be added to
// as well. To make it easier, this will be a linked list
// because...i dunno it just makes more sense to me. Doesn't
// need to be bidirectional...i think.

type StoneSet struct {
	ThisStone Stone
	Next      *StoneSet
}

func NewStoneSet(s Stone) *StoneSet {
	return &StoneSet{
		ThisStone: s,
		Next:      nil,
	}
}

func (s *StoneSet) NextStep() (*StoneSet, error) {
	if s.Next == nil {
		return nil, errors.New("dangling list entry")
	}
	return s.Next, nil
}

func (s *StoneSet) AddAfter(newStone Stone) {
	// adds a stone after the current one.
	n := NewStoneSet(newStone)
	if s.Next == nil {
		s.Next = n
	} else {
		n.Next = s.Next
		s.Next = n
	}
}

func (s *StoneSet) Act() int {
	// this will perform the actions on ONE particular
	// stoneset, which is a linked list.

	// fmt.Printf("Working on %s\n", s.ThisStone.Id)
	// if stone is 0
	if s.ThisStone.Id == "0" {
		s.ThisStone.Id = "1"
		return DENULL
	}

	// if stone has even number of digits
	if len(s.ThisStone.Id)%2 == 0 {
		// fmt.Printf("%s has an even length\n", s.ThisStone.Id)
		left, right := s.ThisStone.Split()
		s.ThisStone.Id = left.Id
		s.AddAfter(right)
		return SPLIT
	}

	// otherwise, multiply
	s.ThisStone.Mult()
	return MULT
}

func (s *StoneSet) Iter(n int) {
	// iterates through each item as many times as the iteration count.
	// MAKE SURE WE'RE AT THE BEGINNING THOUGH!
	for i := 0; i < n; i++ {
		initPtr := s
		for {
			result := initPtr.Act()
			if result == SPLIT {
				if initPtr.Next.Next == nil {
					break
				}
				initPtr = initPtr.Next.Next
				continue
			}

			if initPtr.Next == nil {
				break
			}

			// otherwise...
			initPtr = initPtr.Next
		}
	}
}

func (s *StoneSet) Print() {
	// Just prints each stoneset
	initPtr := s
	fmt.Printf("State: ")
	for {
		fmt.Printf("%s ", initPtr.ThisStone.Id)
		if initPtr.Next == nil {
			break
		}
		initPtr = initPtr.Next
	}
	fmt.Printf("\n")
}

func (s *StoneSet) Count() int {
	// returns how many stones we have at this time
	iPtr := s
	var counter int = 0
	for {
		if iPtr == nil {
			break
		}
		counter++
		iPtr = iPtr.Next
	}
	return counter
}
