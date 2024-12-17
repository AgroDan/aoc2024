package disk

import (
	"fmt"
	"strconv"
)

// This is a disk module used for the day 9 challenge on AOC

type File struct {
	Id         int
	F          []rune
	compressed bool
}

func NewFile(in string, id int) File {
	// accepts only two characters in the string, throw
	// an error otherwise
	if len(in) != 2 {
		panic("need 2 numbers to define file")
	}

	fileUsed, err := strconv.Atoi(string(in[0]))
	if err != nil {
		panic(err)
	}
	contigFree, err := strconv.Atoi(string(in[1]))
	if err != nil {
		panic(err)
	}

	retFile := File{
		Id:         id,
		F:          make([]rune, fileUsed+contigFree),
		compressed: false,
	}
	for i := 0; i < len(retFile.F); i++ {
		if i <= fileUsed-1 {
			retFile.F[i] = rune('0' + retFile.Id) // THIS IS SO DUCKING FUMB GO, WTF
			continue
		}
		retFile.F[i] = '.'
	}

	return retFile
}

func (f File) IsCompressed() bool {
	// returns true if any file has been
	// stuffed in its contiguous free space
	return f.compressed
}

func (f File) FreeSpace() int {
	// returns how much free space is left in the
	// contiguous block
	var counter int = 0
	for i := 0; i < len(f.F); i++ {
		if f.F[i] == '.' {
			counter++
		}
	}
	return counter
}

func (f File) IsEmpty() bool {
	return f.F[0] == '.'
}

func (f File) OriginFileSize() int {
	// returns the size of the origin file if it's there
	var counter int = 0
	for i := 0; i < len(f.F); i++ {
		if f.F[i] == rune(f.Id+'0') {
			counter++
		}
	}
	return counter
}

func (f *File) Push(d rune) error {
	// pushes a part of a file into the contiguous free
	// space of another file. Returns an error if there
	// is no more space left.
	if f.FreeSpace() < 1 {
		return fmt.Errorf("cannot push, no free space left in contiguous block")
	}
	// otherwise, add the rune
	// fmt.Printf("Looping\n")
	for i := 0; i < len(f.F); i++ {
		// fmt.Printf("loop iter %d\n", i)
		if f.F[i] == '.' {
			// fmt.Printf("assigning\n")
			f.F[i] = d
			break
		}
	}
	f.compressed = true
	return nil
}

func (f *File) Pop() (rune, error) {
	// removes a file piece from the right-most part of
	// the data, returning the file piece. Returns an error
	// if empty.

	// f.Print()
	// in the interest of speed...
	if f.F[0] == '.' {
		return '?', fmt.Errorf("cannot pop, empty block")
	}
	for i := len(f.F) - 1; i >= 0; i-- {
		// count backwards
		if f.F[i] != '.' {
			retval := f.F[i]
			f.F[i] = '.'
			return retval, nil
		}
	}
	return '?', fmt.Errorf("not sure how i got here???")
}

func (f *File) PopOriginId() (rune, error) {
	// like the above, this will pop ONLY the origin ID
	// of the file, and not any additional contiguous
	// space left in the file.
	if f.F[0] != rune(f.Id+'0') {
		return '?', fmt.Errorf("cannot pop, origin file non-existant")
	}
	for i := len(f.F) - 1; i >= 0; i-- {
		if f.F[i] == rune(f.Id+'0') {
			// pop it like it's hawt
			retval := f.F[i]
			f.F[i] = '.'
			return retval, nil
		}
	}
	return '?', fmt.Errorf("Not sure how i got here???")
}

func (f File) HowLargeIsFile() int {
	// returns how many space of free space there are
	// in contiguous free space.
	for i := 0; i < len(f.F); i++ {
		if f.F[i] == '.' {
			return i
		}
	}
	return len(f.F)
}

func (f *File) recvOrigin(fromFile *File) error {
	// Will take one data piece of the origin file from a file to another one.
	// will throw an error if this operation cannot be done.
	data, pErr := fromFile.PopOriginId()
	if pErr != nil {
		return pErr
	}
	rErr := f.Push(data)
	if rErr != nil {
		return rErr
	}
	return nil
}

func (f *File) Recv(fromFile *File) error {
	// Will take one data piece from a file to another one.
	// will throw an error if this operation cannot be done.
	data, pErr := fromFile.Pop()
	if pErr != nil {
		return pErr
	}
	rErr := f.Push(data)
	if rErr != nil {
		return rErr
	}
	return nil
}

func (f File) Print() {
	// prints out a repr of the file obj
	fmt.Printf("File ID: %d, ", f.Id)
	for _, v := range f.F {
		fmt.Printf("%c", v)
	}
	fmt.Printf("\n")
}

// For part two, this will receive a file object ONLY if
// the file is big enough. If it is not big enough (meaning
// there isn't enough contiguous free space), then throw
// an error.

func (f *File) RecvOriginFile(fromFile *File) error {
	if f.FreeSpace() >= fromFile.OriginFileSize() {
		for {
			err := f.recvOrigin(fromFile)
			if err != nil {
				// done moving
				break
			}
		}
		return nil
	}
	return fmt.Errorf("Could not transfer file, not enough free space in contiguous free")
}
