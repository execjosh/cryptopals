package xor

import (
	"sort"
	"sync"

	"github.com/execjosh/cryptopals/internal/bytes"
	"github.com/execjosh/cryptopals/internal/english"
)

// FindKey finds the repeating XOR key for a given cipertext
func FindKey(ciphertext []byte) (plaintext []byte, key []byte) {
	type plaintextinfo struct {
		score float64
		value []byte
		key   []byte
	}

	keysizes := findProbableKeySizes(ciphertext)

	var wg sync.WaitGroup
	plaintextinfos := make(chan plaintextinfo)

	for _, keysize := range keysizes {
		wg.Add(1)
		go func(keysize int) {
			defer wg.Done()

			chunks := bytes.Chunk(ciphertext[:], keysize)
			trans := transpose(chunks)
			key := []byte{}
			for _, x := range trans {
				_, k, _ := FindSingleByteKey(x)
				key = append(key, k)
			}

			plaintext := Repeating(ciphertext, key)

			plaintextinfos <- plaintextinfo{
				score: english.Score(plaintext),
				value: plaintext,
				key:   key,
			}
		}(keysize)
	}
	go func() {
		wg.Wait()
		close(plaintextinfos)
	}()

	pts := []plaintextinfo{}
	for pt := range plaintextinfos {
		pts = append(pts, pt)
	}

	sort.Slice(pts, func(i, j int) bool {
		return pts[i].score > pts[j].score
	})

	return pts[0].value, pts[0].key
}

func findProbableKeySizes(ciphertext []byte) []int {
	type keyinfo struct {
		hammingDistance float64
		keysize         int
	}

	var wg sync.WaitGroup
	keyinfos := make(chan keyinfo)

	for keysize := 2; keysize < 40; keysize++ {
		wg.Add(1)
		go func(keysize int) {
			defer wg.Done()

			// this code assumes that there are at least 4*keysize bytes
			// it could be improved to sniff and adapt

			distance := bytes.AverageHammingDistance(
				ciphertext[keysize*0:keysize*1],
				ciphertext[keysize*1:keysize*2],
				ciphertext[keysize*2:keysize*3],
				ciphertext[keysize*3:keysize*4],
			)

			keyinfos <- keyinfo{
				hammingDistance: distance,
				keysize:         keysize,
			}
		}(keysize)
	}
	go func() {
		wg.Wait()
		close(keyinfos)
	}()

	kis := []keyinfo{}
	for ki := range keyinfos {
		kis = append(kis, ki)
	}

	sort.Slice(kis, func(i, j int) bool {
		return kis[i].hammingDistance < kis[j].hammingDistance
	})

	// take only top three contestants
	keysizes := [3]int{}
	for i, v := range kis[:3] {
		keysizes[i] = v.keysize
	}

	return keysizes[:]
}
