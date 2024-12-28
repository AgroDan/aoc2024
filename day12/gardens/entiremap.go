package gardens

import (
	"fmt"
	"utils"
)

// Because I try to minimize the amount of logic on the day##.go files,
// I'll put the puzzle logic here. I will go through each box in the map
// and determine the size of the region. Additionally I will cache places
// that I've been to already and skip over them if that's the case.

type region struct {
	area, perimeter int
}

func CrawlMapGrouped(r utils.Runemap, debug bool) int {
	// list of regions and their stats
	total := make(map[rune]region)

	// crumbs
	var beenthere struct{}
	visited := make(map[utils.Coord]struct{})

	for Y := 0; Y < r.Height(); Y++ {
		for X := 0; X < r.Width(); X++ {
			thisSpot := utils.Coord{
				X: X,
				Y: Y,
			}
			// Have we done this before
			if _, exists := visited[thisSpot]; exists {
				continue
				// we've been here before, no need to check again
			}

			// otherwise set the crumb
			visited[thisSpot] = beenthere

			thisChar, err := r.Get(thisSpot)
			if err != nil {
				panic("invalid space")
			}

			thisRegion := GetRegion(r, thisSpot)

			// note the region as visited
			for i := range thisRegion {
				visited[thisRegion[i]] = beenthere
			}

			if entry, exists := total[thisChar]; exists {
				entry.area += len(thisRegion)
				entry.perimeter += GetRegionPerimeter(r, thisRegion)
				total[thisChar] = entry
			} else {
				total[thisChar] = region{
					area:      len(thisRegion),
					perimeter: GetRegionPerimeter(r, thisRegion),
				}
			}
		}
	}

	if debug {
		for i := range total {
			a := total[i].area
			p := total[i].perimeter
			fmt.Printf("A region of %c plants with price %d * %d = %d.\n", i, a, p, a*p)
		}
	}

	var sum int = 0
	for i := range total {
		sum += total[i].area * total[i].perimeter
	}

	return sum
}

func CrawlMap(r utils.Runemap, debug bool) int {
	var sum int = 0
	// crumbs
	var beenthere struct{}
	visited := make(map[utils.Coord]struct{})

	for Y := 0; Y < r.Height(); Y++ {
		for X := 0; X < r.Width(); X++ {
			thisSpot := utils.Coord{
				X: X,
				Y: Y,
			}
			// Have we done this before
			if _, exists := visited[thisSpot]; exists {
				continue
				// we've been here before, no need to check again
			}

			// otherwise set the crumb
			visited[thisSpot] = beenthere

			thisChar, err := r.Get(thisSpot)
			if err != nil {
				panic("invalid space")
			}

			thisRegion := GetRegion(r, thisSpot)

			// note the region as visited
			for i := range thisRegion {
				visited[thisRegion[i]] = beenthere
			}

			a := len(thisRegion)
			p := GetRegionPerimeter(r, thisRegion)

			if debug {
				fmt.Printf("A region of %c plants with price %d * %d = %d.\n", thisChar, a, p, a*p)
			}
			sum += a * p
		}
	}

	return sum
}

func CrawlMapPartTwo(r utils.Runemap, debug bool) int {
	var sum int = 0
	// crumbs
	var beenthere struct{}
	visited := make(map[utils.Coord]struct{})

	for Y := 0; Y < r.Height(); Y++ {
		for X := 0; X < r.Width(); X++ {
			thisSpot := utils.Coord{
				X: X,
				Y: Y,
			}
			// Have we done this before
			if _, exists := visited[thisSpot]; exists {
				continue
				// we've been here before, no need to check again
			}

			// otherwise set the crumb
			visited[thisSpot] = beenthere

			thisChar, err := r.Get(thisSpot)
			if err != nil {
				panic("invalid space")
			}

			thisRegion := GetRegion(r, thisSpot)

			// note the region as visited
			for i := range thisRegion {
				visited[thisRegion[i]] = beenthere
			}

			a := len(thisRegion)
			s := CountCorners(&r, thisRegion)

			if debug {
				fmt.Printf("A region of %c plants with price %d * %d = %d.\n", thisChar, a, s, a*s)
			}
			sum += a * s
		}
	}

	return sum
}
