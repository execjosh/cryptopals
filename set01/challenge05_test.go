package set01_test

import (
	"testing"

	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/execjosh/cryptopals/internal/xor"

	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/5
func TestChallenge05(t *testing.T) {
	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte("ICE")
	want := hex.StringToBytes("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	got := xor.Repeating(input, key)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
