package bytes

import "bytes"

// SplitLines splits input on the newline byte
func SplitLines(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}
