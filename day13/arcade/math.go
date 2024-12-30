package arcade

import "math"

// This will be where all the heavy math stuff takes place.

const (
	A = iota
	B
)

func (m Machine) Push(button, howMany int) (int, int) {
	// pushes a specific button howMany times, returns
	// the integer value of the X and Y access
	switch button {
	case A:
		return m.A.X * howMany, m.A.Y * howMany
	case B:
		return m.B.X * howMany, m.B.Y * howMany
	default:
		panic("invalid button")
	}
}

func (m Machine) PrizeCalc() [][2]int {
	// returns as many prize calculations as we can find
	var retval [][2]int
buttonA:
	for i := 0; i < 100; i++ {
	buttonB:
		for j := 0; j < 100; j++ {
			aX, aY := m.Push(A, i)
			bX, bY := m.Push(B, j)

			// first, let's see if we pushed it too much
			if aX > m.Prize.X || aY > m.Prize.Y {
				break buttonA
			}

			if bX > m.Prize.X || bY > m.Prize.Y {
				break buttonB
			}

			thisResultX := aX + bX
			thisResultY := aY + bY

			if thisResultX == m.Prize.X && thisResultY == m.Prize.Y {
				retval = append(retval, [2]int{i, j})
			}
		}
	}

	return retval
}

func (m Machine) CalcMaxPushes() (int, int) {
	// calculates the maximum amount of button-presses needed for the
	// A and B buttons

	p2PrizeX, p2PrizeY := m.PartTwoMult()

	// get max a pushes. Lowest number counts because that would mean
	// another push would blow the cap on the other coord
	var aPushes int

	compAX := int(math.Floor(float64(p2PrizeX / m.A.X)))
	compAY := int(math.Floor(float64(p2PrizeY / m.A.Y)))

	if compAX > compAY {
		aPushes = compAY
	} else {
		aPushes = compAX
	}

	var bPushes int

	compBX := int(math.Floor(float64(p2PrizeX / m.B.X)))
	compBY := int(math.Floor(float64(p2PrizeY / m.B.Y)))
	if compBX > compBY {
		bPushes = compBY
	} else {
		bPushes = compBX
	}

	return aPushes, bPushes
}

func (m Machine) SolvePart2() (int, int) {
	// I will now attempt to solve this algebraicly rather than with crazy loops.
	// Since the formula is basically Xa + Ya = ##, Xb + Yb = ##, we can solve this
	// like it's 9th grade algebra again.

	// I'll use the first example to create the forumla here

	// Button A: X+94, Y+34
	// Button B: X+22, Y+67
	// Prize: X=8400, Y=5400

	// So therefore, 94a + 22b = 8400
	// and           34a + 67b = 5400
	// a and b are the same for both equations.

	// So we can eliminiate one by doing this:
	// 34(94a + 22b) = 34*8400
	// 94(34a + 67b) = 94*5400

	// so  3196a + 748b = 285600
	// and 3196a + 6298b = 507600

	// Now we can subtract to eliminate 'a' from the formula

	// (3196a + 6298b) - (3196a + 748b) = 507600 - 285600
	// or 5550b = 222000

	// now solve for b:
	// b = 222000/5550, which = 40

	// now we can substitute 40 in as b and solve

	// 94a + (22 * 40) = 8400
	// 94a + 880 = 8400
	// 94a = 8400 - 880
	// 94a = 7520
	// a = 7520 / 94
	// a = 80

	prizeX, prizeY := m.PartTwoMult()
	// return a 0, 0 if it doesn't compute cleanly!
	// first we'll eliminate button a

	// m.A.X = 94
	// m.A.Y = 34
	// m.B.X = 22
	// m.B.Y = 67
	xCoef := m.A.Y * m.B.X    //748
	newXNum := m.A.Y * prizeX //285600

	yCoef := m.A.X * m.B.Y    //6298
	newYNum := m.A.X * prizeY //507600

	simpLeft := int(math.Abs(float64(xCoef) - float64(yCoef)))    // 5550
	simpAns := int(math.Abs(float64(newXNum) - float64(newYNum))) // 222000

	if simpAns%simpLeft != 0 {
		// doesn't go in evenly
		return 0, 0
	}

	bVal := int(simpAns / simpLeft) // 40
	aLeft := bVal * m.B.X           // 880
	aRight := prizeX - aLeft        // 7520

	if aRight%m.A.X != 0 {
		// doesn't go in evenly
		return 0, 0
	}

	aVal := int(aRight / m.A.X)

	return aVal, bVal

}
