package base64

import "bytes"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encode(str []byte) string {
	var output bytes.Buffer
	var block uint32
	var idx int = 0
	for _, b := range str {
		switch idx {
		case 0:
			block = uint32(b) << 16
			idx++
		case 1:
			block |= uint32(b) << 8
			idx++
		case 2:
			block |= uint32(b)
			output.Write([]byte{
				alphabet[uint8((block&0b111111_000000_000000_000000)>>18)],
				alphabet[uint8((block&0b000000_111111_000000_000000)>>12)],
				alphabet[uint8((block&0b000000_000000_111111_000000)>>6)],
				alphabet[uint8((block&0b000000_000000_000000_111111)>>0)],
			})
			idx = 0
		default:
			panic("this should never happen")
		}
	}

	switch idx {
	case 0:
		// nothing to do; we're at the end of both the input and output
	case 1:
		// there is one byte in the buffer
		// so we must output two chunks
		output.Write([]byte{
			alphabet[uint8((block&0b111111_000000_000000_000000)>>18)],
			alphabet[uint8((block&0b000000_111111_000000_000000)>>12)],
		})
	case 2:
		// there are two bytes in the buffer
		// so we must output three chunks
		output.Write([]byte{
			alphabet[uint8((block&0b111111_000000_000000_000000)>>18)],
			alphabet[uint8((block&0b000000_111111_000000_000000)>>12)],
			alphabet[uint8((block&0b000000_000000_111111_000000)>>6)],
		})
	default:
		panic("this should never happen")
	}

	return output.String()
}
