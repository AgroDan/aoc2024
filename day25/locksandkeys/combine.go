package locksandkeys

// This will just check to see if a key and lock go together

func DoTheyFit(key *Key, lock *Lock) bool {
	if len(key.heights) != len(lock.heights) {
		panic("Key and lock don't match")
	}

	// otherwise...
	for i := 0; i < len(key.heights); i++ {
		if key.heights[i]+lock.heights[i] > 5 {
			return false
		}
	}
	return true
}
