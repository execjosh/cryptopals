package set01_test

import (
	"testing"

	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/execjosh/cryptopals/internal/xor"

	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/2
func TestChallenge02(t *testing.T) {
	have1 := hex.StringToBytes("1c0111001f010100061a024b53535009181c")
	have2 := hex.StringToBytes("686974207468652062756c6c277320657965")
	want := hex.StringToBytes("746865206b696420646f6e277420706c6179")

	got := xor.Fixed(have1, have2)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
