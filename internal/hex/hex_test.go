package hex_test

import (
	"testing"

	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/google/go-cmp/cmp"
)

func TestStringToBytes(t *testing.T) {
	want := []byte("thEre?")
	got := hex.StringToBytes("74684572653f")

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestDecode(t *testing.T) {
	want := []byte("Hello, World!")
	got := hex.Decode([]byte("48656c6c6f2c20576f726c6421"))

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestInvalidLength(t *testing.T) {
	defer func() { _ = recover() }()
	hex.Decode([]byte("123"))
	t.Fatal("expected panic")
}

func TestInvalidByte(t *testing.T) {
	defer func() { _ = recover() }()
	hex.Decode([]byte("ah"))
	t.Fatal("expected panic")
}
