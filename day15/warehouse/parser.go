package warehouse

import (
	"strings"
	"utils"
)

func ParseChallenge(challenge string) Warehouse {
	// this will take the whole challenge as a blob of text and parse accordingly.
	parts := strings.Split(challenge, "\n\n")

	// part[0] is the map
	warehouseMap := strings.Split(parts[0], "\n")

	retval := Warehouse{
		obj:  make(map[utils.Coord]rune),
		rows: len(warehouseMap),
		cols: len(warehouseMap[0]),
	}

	for y := 0; y < len(warehouseMap); y++ {
		for x := 0; x < len(warehouseMap[y]); x++ {
			thisCoord := utils.Coord{
				X: x,
				Y: y,
			}
			obj := rune(warehouseMap[y][x])
			switch obj {
			case '#':
				// wall
				retval.obj[thisCoord] = obj
			case 'O':
				// box
				retval.obj[thisCoord] = obj
			case '@':
				// robot
				retval.r = NewRobot(thisCoord)
			default:
				continue
				// unnecessary but i'm pedantic so
			}
		}
	}

	// Now the instructions
	instructions := strings.Replace(parts[1], "\n", "", -1)
	retval.r.addInstructions(instructions)

	return retval
}

func ParseChallengePartTwo(challenge string) WideWarehouse {
	parts := strings.Split(challenge, "\n\n")

	warehouseMap := strings.Split(parts[0], "\n")

	retval := WideWarehouse{
		obj:  make(map[utils.Coord]*item),
		rows: len(warehouseMap),
		cols: len(warehouseMap) * 2, // remember, this is wider
	}

	for y := 0; y < len(warehouseMap); y++ {
		for x := 0; x < len(warehouseMap[y]); x++ {
			leftCoord := utils.Coord{
				X: x * 2,
				Y: y,
			}
			rightCoord := utils.Coord{
				X: (x * 2) + 1,
				Y: y,
			}

			obj := rune(warehouseMap[y][x])
			switch obj {
			case '#':
				// wall
				i := NewItem(false)
				retval.obj[leftCoord] = i
				retval.obj[rightCoord] = i
			case 'O':
				// box
				i := NewItem(true)
				retval.obj[leftCoord] = i
				retval.obj[rightCoord] = i
			case '@':
				retval.r = NewRobot(leftCoord)
			default:
				continue
			}
		}
	}

	instructions := strings.Replace(parts[1], "\n", "", -1)
	retval.r.addInstructions(instructions)
	return retval
}
