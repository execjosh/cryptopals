package aes128

// mul02 multiplies the polynomial by x (x is 2 in AES)
func mul02(a byte) byte {
	if a&0b1000_0000 == 0 {
		return a << 1
	}

	return (a << 1) ^ 0x1b
}

// mul03 multiplies the polynomial by x + 1 (x is 2 in AES)
func mul03(a byte) byte {
	return mul02(a) ^ a
}

// mul09 is 8 + 1
func mul09(a byte) byte {
	two := mul02(a)
	four := mul02(two)
	eight := mul02(four)
	return eight ^ a
}

// mul0B is 8 + 2 + 1
func mul0B(a byte) byte {
	two := mul02(a)
	four := mul02(two)
	eight := mul02(four)
	return eight ^ two ^ a
}

// mul0D is 8 + 4 + 1
func mul0D(a byte) byte {
	two := mul02(a)
	four := mul02(two)
	eight := mul02(four)
	return eight ^ four ^ a
}

// mul0E is 8 + 4 + 2
func mul0E(a byte) byte {
	two := mul02(a)
	four := mul02(two)
	eight := mul02(four)
	return eight ^ four ^ two
}
