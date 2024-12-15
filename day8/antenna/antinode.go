package antenna

import "utils"

// These are helper functions to determine antinodes

func GetAntinodes(antA, antB utils.Coord) (antinodeA, antinodeB utils.Coord) {
	antinodeA.X = antA.X + (antA.X - antB.X)
	antinodeA.Y = antA.Y + (antA.Y - antB.Y)
	antinodeB.X = antB.X + (antB.X - antA.X)
	antinodeB.Y = antB.Y + (antB.Y - antA.Y)
	return
}

/*
	determine these:
	X: 2, Y: 1
	and
	X: 3, Y: 4

	technically, antinode should be
	X: 4 Y: 7

-2	. # . . .
-1	. . . . .
0	. . . . .
1	. . A . .
2	. . . . .
3	. . . . .
4	. . . A .
5	. . . . .
6	. . . . .
7	. . . . #

	Diff of coords: X: 1, Y: 3
	to find from bottom right, get difference from the bottom-right
	node and add that to the coord
*/
