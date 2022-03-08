package aes128

import "bytes"

func DecryptECB(ciphertext []byte, key [blockSize]byte) []byte {
	var buf bytes.Buffer
	blocks := splitIntoBlocks(ciphertext)
	for _, block := range blocks {
		pt := decryptBlockECB(block, key)
		buf.Write(pt[:])
	}
	return depad(buf.Bytes())
}

func decryptBlockECB(block [blockSize]byte, key [blockSize]byte) [blockSize]byte {
	var aes state

	roundKeys := keyExpansion(key)

	aes = blockToMatrix(block)
	aes.addRoundKey(roundKeys[10])
	for i := 9; i >= 1; i-- {
		aes.invShiftRows()
		aes.invSubBytes()
		aes.addRoundKey(roundKeys[i])
		aes.invMixColumns()
	}
	aes.invShiftRows()
	aes.invSubBytes()
	aes.addRoundKey(roundKeys[0])

	return matrixToBlock(aes)
}

func depad(a []byte) []byte {
	val := a[len(a)-1]
	if val > 0x0F {
		return a
	}

	for i := 0; i < int(val); i++ {
		b := a[len(a)-1-i]
		if b != val {
			return a[:len(a)-1-i]
		}
	}

	return a[:len(a)-1-int(val)]
}
