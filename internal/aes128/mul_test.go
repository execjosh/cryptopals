package aes128

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMul02(t *testing.T) {
	// example from https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf
	for have, want := range map[byte]byte{
		0x57: 0xAE,
		0xAE: 0x47,
		0x47: 0x8E,
		0x8E: 0x07,
	} {
		have, want := have, want
		t.Run(fmt.Sprintf("0x%02x --> 0x%02x", have, want), func(t *testing.T) {
			got := mul02(have)
			if !cmp.Equal(want, got) {
				t.Fatal(cmp.Diff(want, got))
			}
		})
	}
}
