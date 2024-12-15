package antenna

import (
	"fmt"
	"slices"
	"utils"
)

type AntennaMap struct {
	area utils.Runemap
	ant  map[rune][]utils.Coord
}

func NewAntennaMap(lines []string) AntennaMap {
	// parses the map and returns an antennamap
	// with all the antennas added
	thisMap := AntennaMap{
		area: utils.NewRunemap(lines),
		ant:  make(map[rune][]utils.Coord),
	}

	// now loop through the map
	for i := 0; i < thisMap.area.Height(); i++ {
		for j := 0; j < thisMap.area.Width(); j++ {
			checkCoord := utils.Coord{
				X: j,
				Y: i,
			}
			areaItem, err := thisMap.area.Get(checkCoord)
			if err != nil {
				panic(err)
			}
			if areaItem == '.' {
				continue
			}

			// otherwise, record an antenna
			thisMap.ant[areaItem] = append(thisMap.ant[areaItem], checkCoord)
		}
	}
	return thisMap
}

func (a AntennaMap) PrintMap() {
	// prints the map to ensure it has been parsed properly
	for i := 0; i < a.area.Height(); i++ {
		for j := 0; j < a.area.Width(); j++ {
			printCoord := utils.Coord{
				X: j,
				Y: i,
			}
			areaItem, err := a.area.Get(printCoord)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%c", areaItem)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\nFound antennae:\n")
	for k, v := range a.ant {
		fmt.Printf("Symbol: %c\n", k)
		for i := range v {
			fmt.Printf("\tX: %d, Y: %d\n", v[i].X, v[i].Y)
		}
	}
}

func (a AntennaMap) DoesAntennaExist(c utils.Coord) bool {
	// Given a coordinate, will determine if an antenna
	// exists in that coordinate.
	areaCheck, err := a.area.Get(c)
	if err != nil {
		return false
	}
	if areaCheck == '.' {
		return false
	}
	return true
}

func (a AntennaMap) FindAllAntinodes() []utils.Coord {
	// given an antenna map, will return a list of _valid_ antinodes. A valid
	// antinode is considered within the boundaries of the map. Additionally,
	// an antinode is considered valid if it is NOT in a space where an
	// antenna currently exists

	var retval []utils.Coord
	for _, v := range a.ant {
		// For each signal sent, check for the antinodes between each
		// antenna location
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				firstAntinode, secondAntinode := GetAntinodes(v[i], v[j])

				// if a.area.IsInBounds(firstAntinode) && !a.DoesAntennaExist(firstAntinode) {
				// 	retval = append(retval, firstAntinode)
				// }
				// if a.area.IsInBounds(secondAntinode) && !a.DoesAntennaExist(secondAntinode) {
				// 	retval = append(retval, secondAntinode)
				// }

				if a.area.IsInBounds(firstAntinode) {
					// make sure it's unique
					if !slices.Contains(retval, firstAntinode) {
						retval = append(retval, firstAntinode)
					}
				}
				if a.area.IsInBounds(secondAntinode) {
					if !slices.Contains(retval, secondAntinode) {
						retval = append(retval, secondAntinode)
					}
				}
			}
		}
	}
	return retval
}

func (a AntennaMap) FindAllResonantAntinodes() []utils.Coord {
	// just like FindAllAntinodes(), but this accounts for all the resonant frequencies too.
	// note that this will find only unique frequencies.
	var retval []utils.Coord
	for _, v := range a.ant {
		// for each signal sent, check for antinodes and resonant antinodes
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				resonantFreq := a.GetResonanceFrequencyDirection(v[i], v[j])
				for o := range resonantFreq {
					if !slices.Contains(retval, resonantFreq[o]) {
						retval = append(retval, resonantFreq[o])
					}
				}
			}
		}
	}
	return retval
}

func (a AntennaMap) GetResonanceFrequencyDirection(antA, antB utils.Coord) []utils.Coord {
	// This will return all the locations of antinodes accounting for resonance
	// frequency, which means it will reflect out in two different directions
	// until it hits the boundary of the map.

	var retval []utils.Coord
	// in this instance, I will denote a direction as a coordinate, which I will
	// then use to add to a specific coordinate to determine the resonant frequency
	// of a signal.

	directionA := utils.Coord{
		X: antA.X - antB.X,
		Y: antA.Y - antB.Y,
	}
	directionB := utils.Coord{
		X: antB.X - antA.X,
		Y: antB.Y - antA.Y,
	}

	// since it's only two directions i'm just gonna do it twice
	Cursor := utils.Coord{
		X: antA.X,
		Y: antA.Y,
	}
	for {
		if a.area.IsInBounds(Cursor) {
			retval = append(retval, Cursor)
		} else {
			break
		}
		Cursor.X += directionA.X
		Cursor.Y += directionA.Y
	}

	Cursor = utils.Coord{
		X: antB.X,
		Y: antB.Y,
	}
	for {
		if a.area.IsInBounds(Cursor) {
			retval = append(retval, Cursor)
		} else {
			break
		}
		Cursor.X += directionB.X
		Cursor.Y += directionB.Y
	}
	return retval

}
