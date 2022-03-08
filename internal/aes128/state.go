package aes128

type state [4][4]byte

func (s *state) addRoundKey(key [blockSize]byte) {
	k := blockToMatrix(key)
	for c := 0; c < 4; c++ {
		s0 := s[0][c] ^ k[0][c]
		s1 := s[1][c] ^ k[1][c]
		s2 := s[2][c] ^ k[2][c]
		s3 := s[3][c] ^ k[3][c]
		s[0][c] = s0
		s[1][c] = s1
		s[2][c] = s2
		s[3][c] = s3
	}
}

func (s *state) invSubBytes() {
	for i := 0; i < 4; i++ {
		invSBox.subWord(&s[i])
	}
}

func (s *state) invShiftRows() {
	s[1][0], s[1][1], s[1][2], s[1][3] = s[1][3], s[1][0], s[1][1], s[1][2]
	s[2][0], s[2][1], s[2][2], s[2][3] = s[2][2], s[2][3], s[2][0], s[2][1]
	s[3][0], s[3][1], s[3][2], s[3][3] = s[3][1], s[3][2], s[3][3], s[3][0]
}

func (s *state) invMixColumns() {
	for c := 0; c < 4; c++ {
		s0 := mul0E(s[0][c]) ^ mul0B(s[1][c]) ^ mul0D(s[2][c]) ^ mul09(s[3][c])
		s1 := mul09(s[0][c]) ^ mul0E(s[1][c]) ^ mul0B(s[2][c]) ^ mul0D(s[3][c])
		s2 := mul0D(s[0][c]) ^ mul09(s[1][c]) ^ mul0E(s[2][c]) ^ mul0B(s[3][c])
		s3 := mul0B(s[0][c]) ^ mul0D(s[1][c]) ^ mul09(s[2][c]) ^ mul0E(s[3][c])
		s[0][c] = s0
		s[1][c] = s1
		s[2][c] = s2
		s[3][c] = s3
	}
}
