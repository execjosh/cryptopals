package bytes

func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		panic("a and b must have the same length")
	}

	var acc int
	for i := 0; i < len(a); i++ {
		if v := a[i] ^ b[i]; v != 0 {
			for x := 0; x < 8; x++ {
				if v&1 != 0 {
					acc++
				}
				v >>= 1
			}
		}
	}
	return acc
}

func AverageHammingDistance(chunks ...[]byte) float64 {
	numChunks := len(chunks)
	if numChunks < 2 {
		panic("need at least two chunks")
	}

	var dist float64
	for i := 0; i < numChunks-1; i++ {
		a := chunks[i]
		for j := i + 1; j < numChunks; j++ {
			b := chunks[j]
			if len(a) != len(b) {
				panic("all chunks must have the same length")
			}
			dist += float64(HammingDistance(a, b))
		}
	}

	return (dist / float64(len(chunks))) / float64(len(chunks[0]))
}
