package xor

import (
	"bytes"
	"math"
	"sync"

	"github.com/execjosh/cryptopals/internal/english"
)

func FindSingleByteKey(str []byte) (float64, byte, []byte) {
	n := len(str)

	type item struct {
		score float64
		key   byte
		val   []byte
	}
	var best item
	best.score = math.Inf(-1)

	var wg sync.WaitGroup
	items := make(chan item)
	for key := byte(0); key < 0xFF; key++ {
		wg.Add(1)
		go func(str []byte, key byte) {
			defer wg.Done()

			res := Fixed(str, bytes.Repeat([]byte{key}, n))
			score := english.Score(res)
			items <- item{
				score: score,
				key:   key,
				val:   res,
			}
		}(str, key)
	}
	go func() {
		wg.Wait()
		close(items)
	}()

	for m := range items {
		if m.score > best.score {
			best.score = m.score
			best.key = m.key
			best.val = m.val
		}
	}

	return best.score, best.key, best.val
}
