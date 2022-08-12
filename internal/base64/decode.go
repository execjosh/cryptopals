package base64

import (
	"bytes"
	"fmt"
	"strings"
)

func DecodeString(s string) []byte {
	return Decode([]byte(s))
}

func Decode(input []byte) []byte {
	input = bytes.TrimRight(input, "=")

	if len(input) == 0 {
		return []byte{}
	}

	padLen := len(input) % 4
	if padLen == 1 {
		panic("illegal input: length % 4 must not be 1")
	}

	var output bytes.Buffer

	var block uint32
	var idx int
	for _, x := range input {
		b := strings.IndexByte(alphabet, x)
		if b < 0 {
			panic(fmt.Errorf("invalid base64 byte: %#U", x))
		}
		switch idx {
		case 0:
			block = uint32(b) << 18
			idx++
		case 1:
			block |= uint32(b) << 12
			idx++
		case 2:
			block |= uint32(b) << 6
			idx++
		case 3:
			block |= uint32(b)
			output.Write([]byte{
				uint8((block & 0b11111111_00000000_00000000) >> 16),
				uint8((block & 0b00000000_11111111_00000000) >> 8),
				uint8((block & 0b00000000_00000000_11111111) >> 0),
			})
			idx = 0
		default:
			panic("should never happen")
		}
	}

	switch padLen {
	case 0:
		// nothing to do; we've processed the final quantum
	case 2:
		// From rfc4648:
		// The final quantum of encoding input is exactly 8 bits; here, the
		// final unit of encoded output will be two characters followed by
		// two "=" padding characters.
		output.Write([]byte{
			uint8((block & 0b11111111_00000000_00000000) >> 16),
		})
	case 3:
		// From rfc4648:
		// The final quantum of encoding input is exactly 16 bits; here, the
		// final unit of encoded output will be three characters followed by
		// one "=" padding character.
		output.Write([]byte{
			uint8((block & 0b11111111_00000000_00000000) >> 16),
			uint8((block & 0b00000000_11111111_00000000) >> 8),
		})
	default:
		panic("this should never happen")
	}

	return output.Bytes()
}
