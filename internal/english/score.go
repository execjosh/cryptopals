package english

import "math"

// https://en.wikipedia.org/wiki/Letter_frequency
var expectedFreqWikipedia = map[int]float64{
	'A': 8.087,
	'B': 1.493,
	'C': 2.781,
	'D': 4.253,
	'E': 12.702,
	'F': 2.228,
	'G': 2.015,
	'H': 6.094,
	'I': 6.966,
	'J': 0.153,
	'K': 0.772,
	'L': 4.094,
	'M': 2.587,
	'N': 6.749,
	'O': 7.507,
	'P': 1.929,
	'Q': 0.096,
	'R': 5.987,
	'S': 6.234,
	'T': 9.056,
	'U': 2.758,
	'V': 0.978,
	'W': 2.360,
	'X': 0.150,
	'Y': 1.974,
	'Z': 0.074,
}

var lexicon = []int{
	'E',
	'T',
	'A',
	'O',
	'I',
	'N',
	'S',
	'R',
	'H',
	'D',
	'L',
	'U',
	'C',
	'M',
	'F',
	'Y',
	'W',
	'G',
	'P',
	'B',
	'V',
	'K',
	'X',
	'Q',
	'J',
	'Z',
}

func Score(input []byte) float64 {
	n := len(input)
	if n < 1 {
		return math.Inf(-1)
	}

	hist := calcHist(input)
	for k, v := range hist {
		if v == 0 {
			continue
		}

		if r := rune(k); (r < 0x20 || r > 0x7E) && (r != '\n') {
			return math.Inf(-1)
		}
	}

	return math.Abs(1 / scoreFreq(n, hist))
}

func scoreFreq(n int, hist [256]float64) float64 {
	var score float64
	for _, i := range lexicon {
		freq := (hist[i] / float64(n)) * 100
		score += freq - expectedFreqWikipedia[i]
	}

	return score
}

func calcHist(a []byte) [0x100]float64 {
	hist := [0x100]float64{}
	for _, b := range a {
		switch {
		case 'a' <= b && b <= 'z':
			b -= 32 // tally as if upper-case
		}
		hist[b]++
	}
	return hist
}
