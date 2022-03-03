package set01_test

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/execjosh/cryptopals/internal/safe"
	"github.com/execjosh/cryptopals/internal/xor"

	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/4
func TestChallenge04(t *testing.T) {
	want := string(hex.StringToBytes("4e6f77207468617420746865207061727479206973206a756d70696e670a"))

	dat, err := os.ReadFile("../data/4.txt")
	if err != nil {
		panic(err)
	}

	type item struct {
		score     float64
		key       byte
		plaintext []byte
	}
	var best item
	best.score = math.Inf(-1)

	ss := bytes.Split(dat, []byte("\n"))
	for _, s := range ss {
		ciphertext := hex.StringToBytes(string(s))
		score, key, plaintext := xor.FindSingleByteKey(ciphertext)
		if score > best.score {
			best.score = score
			best.key = key
			best.plaintext = plaintext
		}
	}

	fmt.Printf("Score is: %f\n", best.score)
	fmt.Printf("Key is: 0x%02X\n", best.key)
	safe.Println(best.plaintext)

	if got := string(best.plaintext); !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
