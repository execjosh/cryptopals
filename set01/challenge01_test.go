package set01_test

import (
	"testing"

	"github.com/execjosh/cryptopals/internal/base64"
	"github.com/execjosh/cryptopals/internal/hex"

	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/1
func TestChallenge01(t *testing.T) {
	have := hex.StringToBytes("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	got := base64.Encode(have)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
