package base64_test

import (
	"testing"

	"github.com/execjosh/cryptopals/internal/base64"
	"github.com/google/go-cmp/cmp"
)

func TestDecode(t *testing.T) {
	for have, want := range map[string]string{
		"dGVzdA":     "test",
		"dGVzdDE":    "test1",
		"dGVzdDEy":   "test12",
		"dGVzdDEyMw": "test123",

		// Test vectors from rfc4648
		"Zm9vYmFy": "foobar",
		"Zm9vYmE=": "fooba",
		"Zm9vYg==": "foob",
		"Zm9v":     "foo",
		"Zm8=":     "fo",
		"Zg==":     "f",
		"":         "",
	} {
		t.Run(have, func(t *testing.T) {
			got := base64.DecodeString(have)
			if !cmp.Equal([]byte(want), got) {
				t.Fatal(cmp.Diff(want, got))
			}
		})
	}
}
