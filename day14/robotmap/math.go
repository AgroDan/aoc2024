package robotmap

import "utils"

func GetQuadrants(rows, cols int) [4][2]utils.Coord {
	// i hate this function so much
	var retval [4][2]utils.Coord

	midRow := int(rows / 2)
	midCol := int(cols / 2)

	// X .
	// . .
	topLeftq1 := utils.Coord{
		X: 0,
		Y: 0,
	}
	botRightq1 := utils.Coord{
		X: midCol - 1,
		Y: midRow - 1,
	}

	// . X
	// . .
	topLeftq2 := utils.Coord{
		X: midCol + 1,
		Y: 0,
	}
	botRightq2 := utils.Coord{
		X: cols - 1,
		Y: midRow - 1,
	}

	// . .
	// X .
	topLeftq3 := utils.Coord{
		X: 0,
		Y: midRow + 1,
	}
	botRightq3 := utils.Coord{
		X: midCol - 1,
		Y: rows - 1,
	}

	// . .
	// . X
	topLeftq4 := utils.Coord{
		X: midCol + 1,
		Y: midRow + 1,
	}
	botRightq4 := utils.Coord{
		X: cols - 1,
		Y: rows - 1,
	}

	retval[0][0] = topLeftq1
	retval[0][1] = botRightq1

	retval[1][0] = topLeftq2
	retval[1][1] = botRightq2

	retval[2][0] = topLeftq3
	retval[2][1] = botRightq3

	retval[3][0] = topLeftq4
	retval[3][1] = botRightq4

	return retval
}
