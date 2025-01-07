package byteshower

import (
	"strconv"
	"strings"
	"utils"
)

func ParseChallenge(line []string, width int) Shower {
	// This will take the challenge line one at a time
	// and return a Shower object ready for falling.
	shower := Shower{
		width:  width,
		idx:    0,
		fallen: make(map[utils.Coord]struct{}),
	}

	for _, l := range line {
		vals := strings.Split(l, ",")
		x, _ := strconv.Atoi(vals[0])
		y, _ := strconv.Atoi(vals[1])

		shower.coords = append(shower.coords, utils.Coord{
			X: x,
			Y: y,
		})
	}

	return shower
}
