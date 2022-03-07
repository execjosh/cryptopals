package set01_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/execjosh/cryptopals/internal/base64"
	"github.com/execjosh/cryptopals/internal/xor"
	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/6
func TestChallenge06(t *testing.T) {
	dat, err := os.ReadFile("../data/6.txt")
	if err != nil {
		t.Fatal(err)
	}

	ciphertext := bytes.ReplaceAll(dat, []byte("\n"), nil)
	ciphertext = base64.Decode(ciphertext)

	got, _ := xor.FindKey(ciphertext)
	want := base64.DecodeString(`SSdtIGJhY2sgYW5kIEknbSByaW5naW4nIHRoZSBiZWxsIApBIHJvY2tpbicgb24gdGhlIG1pa2Ugd2hpbGUgdGhlIGZseSBnaXJscyB5ZWxsIApJbiBlY3N0YXN5IGluIHRoZSBiYWNrIG9mIG1lIApXZWxsIHRoYXQncyBteSBESiBEZXNoYXkgY3V0dGluJyBhbGwgdGhlbSBaJ3MgCkhpdHRpbicgaGFyZCBhbmQgdGhlIGdpcmxpZXMgZ29pbicgY3JhenkgClZhbmlsbGEncyBvbiB0aGUgbWlrZSwgbWFuIEknbSBub3QgbGF6eS4gCgpJJ20gbGV0dGluJyBteSBkcnVnIGtpY2sgaW4gCkl0IGNvbnRyb2xzIG15IG1vdXRoIGFuZCBJIGJlZ2luIApUbyBqdXN0IGxldCBpdCBmbG93LCBsZXQgbXkgY29uY2VwdHMgZ28gCk15IHBvc3NlJ3MgdG8gdGhlIHNpZGUgeWVsbGluJywgR28gVmFuaWxsYSBHbyEgCgpTbW9vdGggJ2NhdXNlIHRoYXQncyB0aGUgd2F5IEkgd2lsbCBiZSAKQW5kIGlmIHlvdSBkb24ndCBnaXZlIGEgZGFtbiwgdGhlbiAKV2h5IHlvdSBzdGFyaW4nIGF0IG1lIApTbyBnZXQgb2ZmICdjYXVzZSBJIGNvbnRyb2wgdGhlIHN0YWdlIApUaGVyZSdzIG5vIGRpc3NpbicgYWxsb3dlZCAKSSdtIGluIG15IG93biBwaGFzZSAKVGhlIGdpcmxpZXMgc2EgeSB0aGV5IGxvdmUgbWUgYW5kIHRoYXQgaXMgb2sgCkFuZCBJIGNhbiBkYW5jZSBiZXR0ZXIgdGhhbiBhbnkga2lkIG4nIHBsYXkgCgpTdGFnZSAyIC0tIFllYSB0aGUgb25lIHlhJyB3YW5uYSBsaXN0ZW4gdG8gCkl0J3Mgb2ZmIG15IGhlYWQgc28gbGV0IHRoZSBiZWF0IHBsYXkgdGhyb3VnaCAKU28gSSBjYW4gZnVuayBpdCB1cCBhbmQgbWFrZSBpdCBzb3VuZCBnb29kIAoxLTItMyBZbyAtLSBLbm9jayBvbiBzb21lIHdvb2QgCkZvciBnb29kIGx1Y2ssIEkgbGlrZSBteSByaHltZXMgYXRyb2Npb3VzIApTdXBlcmNhbGFmcmFnaWxpc3RpY2V4cGlhbGlkb2Npb3VzIApJJ20gYW4gZWZmZWN0IGFuZCB0aGF0IHlvdSBjYW4gYmV0IApJIGNhbiB0YWtlIGEgZmx5IGdpcmwgYW5kIG1ha2UgaGVyIHdldC4gCgpJJ20gbGlrZSBTYW1zb24gLS0gU2Ftc29uIHRvIERlbGlsYWggClRoZXJlJ3Mgbm8gZGVueWluJywgWW91IGNhbiB0cnkgdG8gaGFuZyAKQnV0IHlvdSdsbCBrZWVwIHRyeWluJyB0byBnZXQgbXkgc3R5bGUgCk92ZXIgYW5kIG92ZXIsIHByYWN0aWNlIG1ha2VzIHBlcmZlY3QgCkJ1dCBub3QgaWYgeW91J3JlIGEgbG9hZmVyLiAKCllvdSdsbCBnZXQgbm93aGVyZSwgbm8gcGxhY2UsIG5vIHRpbWUsIG5vIGdpcmxzIApTb29uIC0tIE9oIG15IEdvZCwgaG9tZWJvZHksIHlvdSBwcm9iYWJseSBlYXQgClNwYWdoZXR0aSB3aXRoIGEgc3Bvb24hIENvbWUgb24gYW5kIHNheSBpdCEgCgpWSVAuIFZhbmlsbGEgSWNlIHllcCwgeWVwLCBJJ20gY29taW4nIGhhcmQgbGlrZSBhIHJoaW5vIApJbnRveGljYXRpbmcgc28geW91IHN0YWdnZXIgbGlrZSBhIHdpbm8gClNvIHB1bmtzIHN0b3AgdHJ5aW5nIGFuZCBnaXJsIHN0b3AgY3J5aW4nIApWYW5pbGxhIEljZSBpcyBzZWxsaW4nIGFuZCB5b3UgcGVvcGxlIGFyZSBidXlpbicgCidDYXVzZSB3aHkgdGhlIGZyZWFrcyBhcmUgam9ja2luJyBsaWtlIENyYXp5IEdsdWUgCk1vdmluJyBhbmQgZ3Jvb3ZpbicgdHJ5aW5nIHRvIHNpbmcgYWxvbmcgCkFsbCB0aHJvdWdoIHRoZSBnaGV0dG8gZ3Jvb3ZpbicgdGhpcyBoZXJlIHNvbmcgCk5vdyB5b3UncmUgYW1hemVkIGJ5IHRoZSBWSVAgcG9zc2UuIAoKU3RlcHBpbicgc28gaGFyZCBsaWtlIGEgR2VybWFuIE5hemkgClN0YXJ0bGVkIGJ5IHRoZSBiYXNlcyBoaXR0aW4nIGdyb3VuZCAKVGhlcmUncyBubyB0cmlwcGluJyBvbiBtaW5lLCBJJ20ganVzdCBnZXR0aW4nIGRvd24gClNwYXJrYW1hdGljLCBJJ20gaGFuZ2luJyB0aWdodCBsaWtlIGEgZmFuYXRpYyAKWW91IHRyYXBwZWQgbWUgb25jZSBhbmQgSSB0aG91Z2h0IHRoYXQgCllvdSBtaWdodCBoYXZlIGl0IApTbyBzdGVwIGRvd24gYW5kIGxlbmQgbWUgeW91ciBlYXIgCic4OSBpbiBteSB0aW1lISBZb3UsICc5MCBpcyBteSB5ZWFyLiAKCllvdSdyZSB3ZWFrZW5pbicgZmFzdCwgWU8hIGFuZCBJIGNhbiB0ZWxsIGl0IApZb3VyIGJvZHkncyBnZXR0aW4nIGhvdCwgc28sIHNvIEkgY2FuIHNtZWxsIGl0IApTbyBkb24ndCBiZSBtYWQgYW5kIGRvbid0IGJlIHNhZCAKJ0NhdXNlIHRoZSBseXJpY3MgYmVsb25nIHRvIElDRSwgWW91IGNhbiBjYWxsIG1lIERhZCAKWW91J3JlIHBpdGNoaW4nIGEgZml0LCBzbyBzdGVwIGJhY2sgYW5kIGVuZHVyZSAKTGV0IHRoZSB3aXRjaCBkb2N0b3IsIEljZSwgZG8gdGhlIGRhbmNlIHRvIGN1cmUgClNvIGNvbWUgdXAgY2xvc2UgYW5kIGRvbid0IGJlIHNxdWFyZSAKWW91IHdhbm5hIGJhdHRsZSBtZSAtLSBBbnl0aW1lLCBhbnl3aGVyZSAKCllvdSB0aG91Z2h0IHRoYXQgSSB3YXMgd2VhaywgQm95LCB5b3UncmUgZGVhZCB3cm9uZyAKU28gY29tZSBvbiwgZXZlcnlib2R5IGFuZCBzaW5nIHRoaXMgc29uZyAKClNheSAtLSBQbGF5IHRoYXQgZnVua3kgbXVzaWMgU2F5LCBnbyB3aGl0ZSBib3ksIGdvIHdoaXRlIGJveSBnbyAKcGxheSB0aGF0IGZ1bmt5IG11c2ljIEdvIHdoaXRlIGJveSwgZ28gd2hpdGUgYm95LCBnbyAKTGF5IGRvd24gYW5kIGJvb2dpZSBhbmQgcGxheSB0aGF0IGZ1bmt5IG11c2ljIHRpbGwgeW91IGRpZS4gCgpQbGF5IHRoYXQgZnVua3kgbXVzaWMgQ29tZSBvbiwgQ29tZSBvbiwgbGV0IG1lIGhlYXIgClBsYXkgdGhhdCBmdW5reSBtdXNpYyB3aGl0ZSBib3kgeW91IHNheSBpdCwgc2F5IGl0IApQbGF5IHRoYXQgZnVua3kgbXVzaWMgQSBsaXR0bGUgbG91ZGVyIG5vdyAKUGxheSB0aGF0IGZ1bmt5IG11c2ljLCB3aGl0ZSBib3kgQ29tZSBvbiwgQ29tZSBvbiwgQ29tZSBvbiAKUGxheSB0aGF0IGZ1bmt5IG11c2ljIAo=`)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
