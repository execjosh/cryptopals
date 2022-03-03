package set01_test

import (
	"fmt"
	"testing"

	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/execjosh/cryptopals/internal/safe"
	"github.com/execjosh/cryptopals/internal/xor"

	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/3
func TestChallenge03(t *testing.T) {
	ciphertext := hex.StringToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	want := string(hex.StringToBytes("436f6f6b696e67204d432773206c696b65206120706f756e64206f66206261636f6e"))

	score, key, plaintext := xor.FindSingleByteKey(ciphertext)

	fmt.Printf("Score is: %f\n", score)
	fmt.Printf("Key is: 0x%02X\n", key)
	safe.Println(plaintext)

	if got := string(plaintext); !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
