package xor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranspose(t *testing.T) {
	have := [][]byte{
		[]byte("abc"),
		[]byte("abc"),
		[]byte("ab"),
	}
	want := [][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
		[]byte("cc"),
	}

	got := transpose(have)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
