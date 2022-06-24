package aes128

import (
	"fmt"
	"testing"

	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/google/go-cmp/cmp"
)

func TestEncryptBlockECB(t *testing.T) {
	for i, tc := range []struct {
		plaintext string
		key       string
		output    string
	}{
		// examples from https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf
		{
			plaintext: "3243f6a8885a308d313198a2e0370734",
			key:       "2b7e151628aed2a6abf7158809cf4f3c",
			output:    "3925841d02dc09fbdc118597196a0b32",
		},
		{
			plaintext: "00112233445566778899aabbccddeeff",
			key:       "000102030405060708090a0b0c0d0e0f",
			output:    "69c4e0d86a7b0430d8cdb78070b4c55a",
		},
	} {
		tc := tc
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var have, key, want [16]byte
			copy(have[:], hex.StringToBytes(tc.plaintext))
			copy(key[:], hex.StringToBytes(tc.key))
			copy(want[:], hex.StringToBytes(tc.output))
			got := encryptBlockECB(have, key)
			if !cmp.Equal(want, got) {
				t.Fatal(cmp.Diff(want, got))
			}
		})
	}
}
