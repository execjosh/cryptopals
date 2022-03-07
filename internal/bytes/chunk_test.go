package bytes_test

import (
	"fmt"
	"testing"

	"github.com/execjosh/cryptopals/internal/bytes"
	"github.com/execjosh/cryptopals/internal/hex"

	"github.com/google/go-cmp/cmp"
)

func TestChunk(t *testing.T) {
	for tn, tc := range []struct {
		have      []byte
		chunkSize int
		want      [][]byte
	}{
		{
			have:      hex.StringToBytes(`abcdefabcdefabcd`),
			chunkSize: 3,
			want: [][]byte{
				hex.StringToBytes(`abcdef`),
				hex.StringToBytes(`abcdef`),
				hex.StringToBytes(`abcd`),
			},
		},
		{
			have:      hex.StringToBytes(`6bc1bee22e409f96e93d7e117393172aae2d8a571e03ac9c9eb76fac45af8e5130c81c46a35ce411e5fbc1191a0a52eff69f2445df4f9b17ad2b417be66c3710`),
			chunkSize: 16,
			want: [][]byte{
				hex.StringToBytes(`6bc1bee22e409f96e93d7e117393172a`),
				hex.StringToBytes(`ae2d8a571e03ac9c9eb76fac45af8e51`),
				hex.StringToBytes(`30c81c46a35ce411e5fbc1191a0a52ef`),
				hex.StringToBytes(`f69f2445df4f9b17ad2b417be66c3710`),
			},
		},
	} {
		tn, tc := tn, tc
		t.Run(fmt.Sprint(tn), func(t *testing.T) {
			got := bytes.Chunk(tc.have, tc.chunkSize)
			if !cmp.Equal(tc.want, got) {
				t.Fatal(cmp.Diff(tc.want, got))
			}
		})
	}

}
