package bytes

func Chunk(input []byte, size int) [][]byte {
	chunks := [][]byte{}

	for {
		n := len(input)
		if n < 1 {
			break
		}

		if n < size {
			size = n
		}

		chunks = append(chunks, input[:size])
		input = input[size:]
	}

	return chunks
}
