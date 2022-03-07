package bytes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHammingDistance(t *testing.T) {
	have1 := []byte("this is a test")
	have2 := []byte("wokka wokka!!!")
	if want, got := 37, HammingDistance(have1, have2); !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
