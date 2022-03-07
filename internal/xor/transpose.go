package xor

func transpose(input [][]byte) [][]byte {
	if len(input) < 1 {
		panic("expected there to be some chunks")
	}

	size := len(input[0])
	trans := make([][]byte, size)

	for i := 0; i < size; i++ {
		for _, x := range input {
			if len(x) <= i {
				continue // hopefully this is at the end
			}
			trans[i] = append(trans[i], x[i])
		}
	}

	return trans
}
