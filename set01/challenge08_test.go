package set01_test

import (
	"fmt"
	"os"
	"sort"
	"sync"
	"testing"

	"github.com/execjosh/cryptopals/internal/bytes"
	"github.com/execjosh/cryptopals/internal/hex"
	"github.com/google/go-cmp/cmp"
)

// https://cryptopals.com/sets/1/challenges/8
func TestChallenge08(t *testing.T) {
	dat, err := os.ReadFile("../data/8.txt")
	if err != nil {
		t.Fatal(err)
	}

	type item struct {
		score int
		data  []byte
	}

	var wg sync.WaitGroup
	ch := make(chan item)

	lines := bytes.SplitLines(dat)
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		wg.Add(1)
		go func(line []byte) {
			defer wg.Done()

			line = hex.Decode(line)

			chunks := bytes.Chunk(line, 16)
			if len(chunks) < 2 {
				panic("need at least 2 chunks")
			}

			cache := map[string]struct{}{}
			var repetitions int
			for _, c := range chunks {
				v := fmt.Sprintf("%x", c)
				if _, seen := cache[v]; seen {
					repetitions++
				} else {
					cache[v] = struct{}{}
				}
			}

			ch <- item{
				score: repetitions,
				data:  line,
			}
		}(line)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var items []item
	for i := range ch {
		items = append(items, i)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].score > items[j].score
	})

	got := items[0].data
	want := hex.StringToBytes(`d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a`)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
