package xor

// Fixed XORs a with b, both of which must be the same length
func Fixed(a []byte, b []byte) []byte {
	n := len(a)

	if n != len(b) {
		panic("input lengths differ")
	}

	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = a[i] ^ b[i]
	}

	return v
}
