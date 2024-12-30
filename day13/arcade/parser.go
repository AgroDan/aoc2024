package arcade

import (
	"strings"
)

// this will parse the input data into machine object types

func Parser(lines string) []Machine {
	arcades := strings.Split(lines, "\n\n")

	var retval []Machine
	for _, arc := range arcades {
		retval = append(retval, NewMachine(arc))
	}
	return retval
}
