package xor

// Repeating XORs each byte of a with the corresponding index of key, looping
// around until all bytes of a have been processed.
func Repeating(a []byte, key []byte) []byte {
	n := len(a)
	keylen := len(key)

	v := make([]byte, n)
	for i := 0; i < n; i++ {
		b := key[i%keylen]
		v[i] = a[i] ^ b
	}

	return v
}
