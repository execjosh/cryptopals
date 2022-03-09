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
	for i := 0; i < len(s); i++ {
		v := hexDigit(s[i])
		if i%2 == 0 {
			digitBuf = (v & 0xF) << 4
			continue
		}

		digitBuf |= v & 0xF

		buf.WriteByte(digitBuf)
	}

	return buf.Bytes()
}

func hexDigit(b byte) uint8 {
	switch {
	case '0' <= b && b <= '9':
		return b - '0'
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10
	}
	panic(fmt.Errorf("unknown hex digit %#U", b))
}
