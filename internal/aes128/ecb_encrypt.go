package aes128

func encryptBlockECB(block [blockSize]byte, key [blockSize]byte) [blockSize]byte {
	var aes state

	roundKeys := keyExpansion(key)

	aes = blockToMatrix(block)
	aes.addRoundKey(roundKeys[0])
	for i := 1; i <= 9; i++ {
		aes.subBytes()
		aes.shiftRows()
		aes.mixColumns()
		aes.addRoundKey(roundKeys[i])
	}
	aes.subBytes()
	aes.shiftRows()
	aes.addRoundKey(roundKeys[10])

	return matrixToBlock(aes)
}

func (s *state) subBytes() {
	for i := 0; i < 4; i++ {
		mainSBox.subWord(&s[i])
	}
}

func (s *state) shiftRows() {
	s[1][0], s[1][1], s[1][2], s[1][3] = s[1][1], s[1][2], s[1][3], s[1][0]
	s[2][0], s[2][1], s[2][2], s[2][3] = s[2][2], s[2][3], s[2][0], s[2][1]
	s[3][0], s[3][1], s[3][2], s[3][3] = s[3][3], s[3][0], s[3][1], s[3][2]
}

func (s *state) mixColumns() {
	for c := 0; c < 4; c++ {
		s0 := mul02(s[0][c]) ^ mul03(s[1][c]) ^ s[2][c] ^ s[3][c]
		s1 := s[0][c] ^ mul02(s[1][c]) ^ mul03(s[2][c]) ^ s[3][c]
		s2 := s[0][c] ^ s[1][c] ^ mul02(s[2][c]) ^ mul03(s[3][c])
		s3 := mul03(s[0][c]) ^ s[1][c] ^ s[2][c] ^ mul02(s[3][c])
		s[0][c] = s0
		s[1][c] = s1
		s[2][c] = s2
		s[3][c] = s3
	}
}
