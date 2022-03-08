package aes128

import "github.com/execjosh/cryptopals/internal/xor"

// keyExpansion expands key to 11 for 128-bit key
func keyExpansion(key [blockSize]byte) [11][blockSize]byte {
	var keys [11][blockSize]byte

	copy(keys[0][:], key[:])

	prev := wordifyKey(keys[0])
	temp := prev[3]
	for round := 1; round < 11; round++ {
		rotWord(&temp)
		mainSBox.subWord(&temp)

		copy(temp[:], xor.Fixed(temp[:], rcon[round][:]))

		copy(temp[:], xor.Fixed(prev[0][:], temp[:]))
		copy(keys[round][0x00:0x04], temp[:])
		copy(prev[0][:], temp[:])

		copy(temp[:], xor.Fixed(prev[1][:], temp[:]))
		copy(keys[round][0x04:0x08], temp[:])
		copy(prev[1][:], temp[:])

		copy(temp[:], xor.Fixed(prev[2][:], temp[:]))
		copy(keys[round][0x08:0x0C], temp[:])
		copy(prev[2][:], temp[:])

		copy(temp[:], xor.Fixed(prev[3][:], temp[:]))
		copy(keys[round][0x0C:0x10], temp[:])
		copy(prev[3][:], temp[:])
	}

	return keys
}

func rotWord(w *[4]byte) {
	w[0], w[1], w[2], w[3] = w[1], w[2], w[3], w[0]
}

func wordifyKey(block [blockSize]byte) [4][4]byte {
	var m [4][4]byte
	for k, i := 0, 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = block[k]
			k++
		}
	}
	return m
}

var rcon = [][4]byte{
	{0x00, 0x00, 0x00, 0x00},
	{0x01, 0x00, 0x00, 0x00},
	{0x02, 0x00, 0x00, 0x00},
	{0x04, 0x00, 0x00, 0x00},
	{0x08, 0x00, 0x00, 0x00},
	{0x10, 0x00, 0x00, 0x00},
	{0x20, 0x00, 0x00, 0x00},
	{0x40, 0x00, 0x00, 0x00},
	{0x80, 0x00, 0x00, 0x00},
	{0x1b, 0x00, 0x00, 0x00},
	{0x36, 0x00, 0x00, 0x00},
}
