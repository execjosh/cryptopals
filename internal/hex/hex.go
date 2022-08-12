package hex

import (
	"bytes"
	"fmt"
)

func StringToBytes(s string) []byte {
	return Decode([]byte(s))
}

func Decode(s []byte) []byte {
	if len(s)%2 != 0 {
		panic("invalid input length not multiple of 2")
	}

	var buf bytes.Buffer

	var digitBuf uint8
	for i, d := range s {
		v, ok := hexDigitToByte[d]
		if !ok {
			panic(fmt.Errorf("unknown hex digit %#U", d))
		}

		if i%2 == 0 {
			digitBuf = (v & 0xF) << 4
			continue
		}
		digitBuf |= v & 0xF

		buf.WriteByte(digitBuf)
	}

	return buf.Bytes()
}

var hexDigitToByte = map[byte]byte{
	'0': 0x0,
	'1': 0x1,
	'2': 0x2,
	'3': 0x3,
	'4': 0x4,
	'5': 0x5,
	'6': 0x6,
	'7': 0x7,
	'8': 0x8,
	'9': 0x9,
	'a': 0xa, 'A': 0xa,
	'b': 0xb, 'B': 0xb,
	'c': 0xc, 'C': 0xc,
	'd': 0xd, 'D': 0xd,
	'e': 0xe, 'E': 0xe,
	'f': 0xf, 'F': 0xf,
}
