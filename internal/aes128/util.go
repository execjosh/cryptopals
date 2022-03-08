package aes128

func blockToMatrix(b [blockSize]byte) [4][4]byte {
	var m [4][4]byte

	for k, i := 0, 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[j][i] = b[k]
			k++
		}
	}

	return m
}

func matrixToBlock(m [4][4]byte) [blockSize]byte {
	var b [blockSize]byte

	for k, i := 0, 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			b[k] = m[j][i]
			k++
		}
	}

	return b
}

func splitIntoBlocks(input []byte) [][blockSize]byte {
	input = pad(input)
	if len(input)%blockSize != 0 {
		panic("input not multiple of block size")
	}

	var blocks [][blockSize]byte
	for {
		n := len(input)
		if n < 1 {
			break
		}

		if n < blockSize {
			panic("input not multiple of block size")
		}

		var block [blockSize]byte
		copy(block[:], input[:blockSize])
		blocks = append(blocks, block)
		input = input[blockSize:]
	}

	return blocks
}

func pad(a []byte) []byte {
	n := len(a) % blockSize
	if n == 0 {
		return a
	}

	for i := 0; i < n; i++ {
		a = append(a, byte(n))
	}

	return a
}
