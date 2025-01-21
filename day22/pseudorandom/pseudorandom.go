package pseudorandom

func Generator(seed, rounds int) int {
	// According to the challenge:
	// 1) multiply the seed by 64.
	// 2) mix this result into the seed (see below)
	// 3) prune the seed
	// 4) divide the secret number by 32, rounding down
	// 5) mix this result into the seed
	// 6) prune the seed
	// 7) multiply the seed by 2048
	// 8) mix this result into the seed
	// 9) prune the seed.
	//
	// To mix: get XOR of the given value and the seed, the seed now becomes this result
	// to prune: get the value of the seed and modulo it by 16777216

	secretNumber := seed // just so I can manipulate it, ignore my pedantic nature

	for i := 0; i < rounds; i++ {
		// Step one
		step1val := secretNumber * 64

		// step two
		secretNumber = step1val ^ secretNumber

		// step three
		secretNumber = secretNumber % 16777216

		// step four
		step4val := int(secretNumber / 32)

		// step five
		secretNumber = step4val ^ secretNumber

		// step 6
		secretNumber = secretNumber % 16777216

		// step 7
		step7val := secretNumber * 2048

		// step 8
		secretNumber = step7val ^ secretNumber

		// step 9
		secretNumber = secretNumber % 16777216
	}
	return secretNumber
}

func GenerateAll(secretNumber, rounds int) []int {

	retval := make([]int, rounds+1)
	retval[0] = secretNumber
	for i := 0; i < rounds; i++ {
		// Step one
		step1val := secretNumber * 64

		// step two
		secretNumber = step1val ^ secretNumber

		// step three
		secretNumber = secretNumber % 16777216

		// step four
		step4val := int(secretNumber / 32)

		// step five
		secretNumber = step4val ^ secretNumber

		// step 6
		secretNumber = secretNumber % 16777216

		// step 7
		step7val := secretNumber * 2048

		// step 8
		secretNumber = step7val ^ secretNumber

		// step 9
		secretNumber = secretNumber % 16777216

		retval[i+1] = secretNumber
	}
	return retval
}

func GetCharacteristics(secretNumbers []int) [][3]int {
	// This will create a two dimensional array, where
	// each row[0] is the provided "secret number", followed
	// by row[1] being the ones digit, and row[2] being the
	// difference betweeen it and the number previous to it.

	retval := make([][3]int, len(secretNumbers))
	for i, secretNumber := range secretNumbers {
		retval[i][0] = secretNumber
		retval[i][1] = secretNumber % 10
		if i == 0 {
			retval[i][2] = 0
		} else {
			retval[i][2] = retval[i][1] - (secretNumbers[i-1] % 10)
		}
	}
	return retval
}
