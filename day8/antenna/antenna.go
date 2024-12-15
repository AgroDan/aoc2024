package antenna

// this module will specify an antenna object, which will outline all
// known coordinates of a specific antenna

// this code is unused but I'm keeping it here anyway
import (
	"slices"
	"utils"
)

type Antenna struct {
	freq rune
	loc  []utils.Coord
}

func NewAntenna(a rune, loc utils.Coord) Antenna {
	var locArr []utils.Coord
	locArr = append(locArr, loc)
	return Antenna{
		freq: a,
		loc:  locArr,
	}
}

func (a *Antenna) NewLoc(loc utils.Coord) {
	// adds a new location if it doesn't already exist
	if !slices.Contains(a.loc, loc) {
		a.loc = append(a.loc, loc)
	}

	// no doubles, silently fail if doesn't exist
}
