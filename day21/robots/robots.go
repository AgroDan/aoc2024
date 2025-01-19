package robots

import (
	"cmp"
	"fmt"
	"slices"
	"utils"
)

// This package will define the robot charts.

type Keypad struct {
	keys utils.Runemap
	idx  utils.Coord
}

func NewNumberPad() *Keypad {
	// This will return a Numberpad version
	pad := []string{
		"789",
		"456",
		"123",
		" 0A",
	}
	return &Keypad{
		keys: utils.NewRunemap(pad),
		idx: utils.Coord{
			X: 2,
			Y: 3,
		},
	}
}

func NewDirectionPad() *Keypad {
	pad := []string{
		" ^A",
		"<v>",
	}
	return &Keypad{
		keys: utils.NewRunemap(pad),
		idx: utils.Coord{
			X: 2,
			Y: 0,
		},
	}
}

func (k *Keypad) findPaths(current, dst utils.Coord, path string, result *[]string) {
	// This function will recursively find paths to its ultimate destination.
	currDistance := utils.ManhattanDistance(current, dst)

	if current == dst {
		// since we're at the destination, always add a push of 'A'
		path += "A"
		*result = append(*result, path)
	}

	// otherwise, lets get neighbors
	neighbors := current.AllAvailable()
	for i, v := range neighbors {
		if thisButton, err := k.keys.Get(v); thisButton == ' ' || err != nil {
			// either we're outside the boundary or we're on a blank space
			continue
		}
		if checkDistance := utils.ManhattanDistance(v, dst); checkDistance >= currDistance {
			// not worth our time to visit here
			continue
		}
		// create a copy of the path
		newPath := path

		switch i {
		case utils.N:
			newPath += "^"
		case utils.S:
			newPath += "v"
		case utils.E:
			newPath += ">"
		case utils.W:
			newPath += "<"
		}

		k.findPaths(v, dst, newPath, result)
	}
}

func (k *Keypad) GetTotalPathsPer(code string) []string {
	// This, given a sequence of characters, will return all the possible
	// combinations of routes one can take to enter the entire code
	start := k.idx
	combos := make([]string, 0)
	for _, c := range code {
		char := rune(c)
		paths := make([]string, 0)

		destChar, err := k.keys.Find(char)
		if err != nil {
			panic(err)
		}

		k.findPaths(start, destChar, "", &paths)

		start = destChar
		newCombos := make([]string, 0)
		for i := 0; i < len(combos); i++ {
			for _, v := range paths {
				newCombos = append(newCombos, combos[i]+v)
			}
		}
		if len(combos) == 0 {
			combos = paths
		} else {
			combos = newCombos
		}
	}
	return combos
}

func RobotChain(code string, dirRobotsAmt int) []string {
	// this is a helper function which will take a list of directions
	// and iterate down several levels, outputting the total amount of
	// codes it would take to get to the end.

	dirPad := NewDirectionPad()
	result := dirPad.GetTotalPathsPer(code)
	for i := 0; i < dirRobotsAmt; i++ {
		newResult := make([]string, 0)
		for _, v := range result {
			newResult = append(newResult, dirPad.GetTotalPathsPer(v)...)
		}
		result = newResult
	}
	return result
}

func (k *Keypad) GetDistance(buttons string) int {
	if len(buttons) != 2 {
		panic("need string length of 2")
	}
	startChar, _ := k.keys.Find(rune(buttons[0]))
	endChar, _ := k.keys.Find(rune(buttons[1]))
	return utils.ManhattanDistance(startChar, endChar) + 1
}

func GetMinLength(code string, dirRobotsAmt int) int {
	// Similar to the above, this will take a full code and return the minimum
	// length (or complexity) of a code at a given length. This should satisfy
	// part 1

	allPaths := RobotChain(code, dirRobotsAmt)
	return len(slices.MinFunc(allPaths, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}))
}

func CachedGetMinLength(code string, depth int, seq map[string][]string, cache *utils.Cache) int {
	code = "A" + code

	if depth == 1 {
		total := 0
		for i := 0; i < len(code)-1; i++ {
			char1 := string(code[i])
			char2 := string(code[i+1])
			total += len(seq[char1+char2][0])
		}
		return total
	}

	length := 0
	// allPaths := RobotChain(code, depth)
	for i := 0; i < len(code)-1; i++ {
		// Need the code mapping
		movement := fmt.Sprintf("%s%s", string(code[i]), string(code[i+1]))
		move_map := seq[movement]

		res := make([]int, 0)
		for _, v := range move_map {
			cacheKey := fmt.Sprintf("minlength_%s_%d", v, depth-1)
			res = append(res, cache.Get(cacheKey, func() interface{} {
				return CachedGetMinLength(v, depth-1, seq, cache)
			}).(int))
		}
		length += slices.Min(res)
	}
	return length
}

func ComputeSequences(k *Keypad) map[string][]string {
	// This will just compute the total number of sequences for each pair of
	// possibilities.
	pad := k.keys.GetRaw()
	var items []rune
	for Y := range pad {
		items = append(items, pad[Y]...)
	}

	pairs := utils.CartesianProduct(items, items)

	retval := make(map[string][]string)
	for _, v := range pairs {
		if v[0] == ' ' || v[1] == ' ' {
			continue
		}
		convertedCode := fmt.Sprintf("%s%s", string(v[0]), string(v[1]))
		paths := make([]string, 0)
		// fmt.Printf("Working on pair: %s\n", convertedCode)
		GetOptimalPaths(convertedCode, k, "", &paths)
		retval[convertedCode] = append(retval[convertedCode], paths...)
	}

	return retval
}
