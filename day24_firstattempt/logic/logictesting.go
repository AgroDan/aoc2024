package logic

import (
	"slices"
	"sort"
	"strconv"
	"strings"
	"utils"
)

// This will test the values of X and Y and test whether or not they add up to the correct number.
// For this, I will simply add 10,000 and 20,000 together to get 30,000, or in binary:
// 0010 0111 0001 0000 + 0100 1110 0010 0000 = 0111 0101 0011 0000

// Since the full value of X can only be 22 bits, this value will be 4_194_303 (aka 0011_1111_1111_1111_1111_1111)
// and

func TestValues(x, y, expected int, wireList map[string]*Wire, instructionList utils.GQueue[string]) bool {
	// re run the test, but this time check the Z values for the supplied X and Y values to determine if
	// they are correct.
	xList := make([]string, 0)
	yList := make([]string, 0)

	for k := range wireList {
		if strings.HasPrefix(k, "x") {
			xList = append(xList, k)
		}
		if strings.HasPrefix(k, "y") {
			yList = append(yList, k)
		}
	}
	sort.Strings(xList)
	slices.Reverse(xList)
	sort.Strings(yList)
	slices.Reverse(yList) // little endian!

	xNum := strconv.FormatInt(int64(x), 2)
	xPadding := len(xList) - len(xNum)
	for i := 0; i < xPadding; i++ {
		xNum = "0" + xNum
	}

	yNum := strconv.FormatInt(int64(y), 2)
	yPadding := len(yList) - len(yNum)
	for i := 0; i < yPadding; i++ {
		yNum = "0" + yNum
	}

	// hopefully this shouldn't throw any weird out-of-index errors...
	for i := 0; i < len(xList); i++ {
		wireList[xList[i]].Value = xNum[i] == '1'
		wireList[xList[i]].Ready = true
	}

	for i := 0; i < len(yList); i++ {
		wireList[yList[i]].Value = yNum[i] == '1'
		wireList[yList[i]].Ready = true
	}

	ParseInstructions(wireList, instructionList)
	zWires := make([]string, 0)
	for k := range wireList {
		if strings.HasPrefix(k, "z") {
			zWires = append(zWires, k)
		}
	}
	sort.Strings(zWires)
	outBin := ""
	for _, v := range zWires {
		if wireList[v].Value {
			outBin = "1" + outBin
		} else {
			outBin = "0" + outBin
		}
	}
	decimal, err := strconv.ParseInt(outBin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(decimal) == x+y

}
