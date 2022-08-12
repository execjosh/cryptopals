package base64_test

import (
	"testing"

	"github.com/execjosh/cryptopals/internal/base64"
	"github.com/google/go-cmp/cmp"
)

func TestEncode(t *testing.T) {
	for have, want := range map[string]string{
		"test":    "dGVzdA",
		"test1":   "dGVzdDE",
		"test12":  "dGVzdDEy",
		"test123": "dGVzdDEyMw",

		// Test vectors from rfc4648
		"foobar": "Zm9vYmFy",
		"fooba":  "Zm9vYmE",
		"foob":   "Zm9vYg",
		"foo":    "Zm9v",
		"fo":     "Zm8",
		"f":      "Zg",
		"":       "",
	} {
		t.Run(have, func(t *testing.T) {
			got := base64.Encode([]byte(have))
			if !cmp.Equal(want, got) {
				t.Fatal(cmp.Diff(want, got))
			}
		})
	}
}
