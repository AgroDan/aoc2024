package disk

func ParseDiskReport(line string) []File {
	// takes a giant string and returns a file.
	// Will panic if the length of the string
	// is not divisible by two, which would mean
	// the challenge would not work
	// if len(line)%2 != 0 {
	// 	panic("Not an even number of characters in input string")
	// }

	var retval []File
	var counter int = 0
	for i := 0; i < len(line); i += 2 {
		// two at a time, if odd then no free space
		// bottom, top := i, i+1
		var checkStr string
		if i+1 >= len(line) {
			checkStr = string(line[i]) + "0"
		} else {
			checkStr = line[i : i+2]
		}
		thisFile := NewFile(checkStr, counter)
		counter++
		retval = append(retval, thisFile)
	}
	return retval
}
