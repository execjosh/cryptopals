package english

import (
	"math"
)

// generated from https://www.gutenberg.org/cache/epub/67571/pg67571.txt
var expectedHist = [256]float64{
	'\x00': 0.000000,
	'\x01': 0.000000,
	'\x02': 0.000000,
	'\x03': 0.000000,
	'\x04': 0.000000,
	'\x05': 0.000000,
	'\x06': 0.000000,
	'\a':   0.000000,
	'\b':   0.000000,
	'\t':   0.000000,
	'\n':   2.046033,
	'\v':   0.000000,
	'\f':   0.000000,
	'\r':   2.046033,
	'\x0e': 0.000000,
	'\x0f': 0.000000,
	'\x10': 0.000000,
	'\x11': 0.000000,
	'\x12': 0.000000,
	'\x13': 0.000000,
	'\x14': 0.000000,
	'\x15': 0.000000,
	'\x16': 0.000000,
	'\x17': 0.000000,
	'\x18': 0.000000,
	'\x19': 0.000000,
	'\x1a': 0.000000,
	'\x1b': 0.000000,
	'\x1c': 0.000000,
	'\x1d': 0.000000,
	'\x1e': 0.000000,
	'\x1f': 0.000000,
	' ':    16.723582,
	'!':    0.005028,
	'"':    0.768886,
	'#':    0.000419,
	'$':    0.001257,
	'%':    0.000419,
	'&':    0.000000,
	'\'':   0.477254,
	'(':    0.008380,
	')':    0.008380,
	'*':    0.006704,
	'+':    0.000000,
	',':    0.938166,
	'-':    0.182270,
	'.':    1.348798,
	'/':    0.003352,
	'0':    0.009637,
	'1':    0.031845,
	'2':    0.006285,
	'3':    0.005866,
	'4':    0.007542,
	'5':    0.006285,
	'6':    0.005447,
	'7':    0.003771,
	'8':    0.006704,
	'9':    0.004609,
	':':    0.007123,
	';':    0.004190,
	'<':    0.000000,
	'=':    0.000000,
	'>':    0.000000,
	'?':    0.082545,
	'@':    0.000000,
	'A':    0.123189,
	'B':    0.192326,
	'C':    0.064947,
	'D':    0.062014,
	'E':    0.072908,
	'F':    0.068718,
	'G':    0.085059,
	'H':    0.224590,
	'I':    0.346522,
	'J':    0.082126,
	'K':    0.035616,
	'L':    0.058243,
	'M':    0.135760,
	'N':    0.059081,
	'O':    0.068299,
	'P':    0.078355,
	'Q':    0.000419,
	'R':    0.167605,
	'S':    0.174728,
	'T':    0.380043,
	'U':    0.029750,
	'V':    0.005447,
	'W':    0.437867,
	'X':    0.002933,
	'Y':    0.066204,
	'Z':    0.000000,
	'[':    0.000419,
	'\\':   0.000000,
	']':    0.000419,
	'^':    0.000000,
	'_':    0.010056,
	'`':    0.000000,
	'a':    5.986416,
	'b':    0.992219,
	'c':    1.911111,
	'd':    3.743448,
	'e':    9.264342,
	'f':    1.505927,
	'g':    1.697415,
	'h':    4.536217,
	'i':    4.532027,
	'j':    0.107267,
	'k':    0.842632,
	'l':    2.611279,
	'm':    1.515564,
	'n':    4.866398,
	'o':    5.922726,
	'p':    1.122112,
	'q':    0.044415,
	'r':    4.411771,
	's':    3.916499,
	't':    6.711305,
	'u':    1.936671,
	'v':    0.697235,
	'w':    1.599785,
	'x':    0.102658,
	'y':    1.614870,
	'z':    0.024303,
	'{':    0.000000,
	'|':    0.000000,
	'}':    0.000000,
	'~':    0.000000,
	'\x7f': 0.000000,
	'\x80': 0.000000,
	'\x81': 0.000000,
	'\x82': 0.000000,
	'\x83': 0.000000,
	'\x84': 0.000000,
	'\x85': 0.000000,
	'\x86': 0.000000,
	'\x87': 0.000000,
	'\x88': 0.000000,
	'\x89': 0.000000,
	'\x8a': 0.000000,
	'\x8b': 0.000000,
	'\x8c': 0.000000,
	'\x8d': 0.000000,
	'\x8e': 0.000000,
	'\x8f': 0.000000,
	'\x90': 0.000000,
	'\x91': 0.000000,
	'\x92': 0.000000,
	'\x93': 0.000000,
	'\x94': 0.000000,
	'\x95': 0.000000,
	'\x96': 0.000000,
	'\x97': 0.000000,
	'\x98': 0.000000,
	'\x99': 0.000000,
	'\x9a': 0.000000,
	'\x9b': 0.000000,
	'\x9c': 0.000000,
	'\x9d': 0.000000,
	'\x9e': 0.000000,
	'\x9f': 0.000000,
	'\xa0': 0.000000,
	'¡':    0.000000,
	'¢':    0.000000,
	'£':    0.000000,
	'¤':    0.000000,
	'¥':    0.000000,
	'¦':    0.000000,
	'§':    0.000000,
	'¨':    0.000000,
	'©':    0.000419,
	'ª':    0.000000,
	'«':    0.000000,
	'¬':    0.000000,
	'\xad': 0.000000,
	'®':    0.000000,
	'¯':    0.000000,
	'°':    0.000000,
	'±':    0.000000,
	'²':    0.000000,
	'³':    0.000000,
	'´':    0.000000,
	'µ':    0.000000,
	'¶':    0.000000,
	'·':    0.000000,
	'¸':    0.000000,
	'¹':    0.000000,
	'º':    0.000000,
	'»':    0.000000,
	'¼':    0.000000,
	'½':    0.000000,
	'¾':    0.000000,
	'¿':    0.000000,
	'À':    0.000000,
	'Á':    0.000000,
	'Â':    0.000000,
	'Ã':    0.000000,
	'Ä':    0.000000,
	'Å':    0.000000,
	'Æ':    0.000000,
	'Ç':    0.000000,
	'È':    0.000000,
	'É':    0.000000,
	'Ê':    0.000000,
	'Ë':    0.000000,
	'Ì':    0.000000,
	'Í':    0.000000,
	'Î':    0.000000,
	'Ï':    0.000000,
	'Ð':    0.000000,
	'Ñ':    0.000000,
	'Ò':    0.000000,
	'Ó':    0.000000,
	'Ô':    0.000000,
	'Õ':    0.000000,
	'Ö':    0.000000,
	'×':    0.000000,
	'Ø':    0.000000,
	'Ù':    0.000000,
	'Ú':    0.000000,
	'Û':    0.000000,
	'Ü':    0.000000,
	'Ý':    0.000000,
	'Þ':    0.000000,
	'ß':    0.000000,
	'à':    0.000000,
	'á':    0.000000,
	'â':    0.000000,
	'ã':    0.000000,
	'ä':    0.000000,
	'å':    0.000000,
	'æ':    0.000000,
	'ç':    0.000000,
	'è':    0.000000,
	'é':    0.000419,
	'ê':    0.000000,
	'ë':    0.000000,
	'ì':    0.000000,
	'í':    0.000000,
	'î':    0.000000,
	'ï':    0.000000,
	'ð':    0.000000,
	'ñ':    0.000000,
	'ò':    0.000000,
	'ó':    0.000000,
	'ô':    0.000000,
	'õ':    0.000000,
	'ö':    0.000000,
	'÷':    0.000000,
	'ø':    0.000000,
	'ù':    0.000000,
	'ú':    0.000000,
	'û':    0.000000,
	'ü':    0.000000,
	'ý':    0.000000,
	'þ':    0.000000,
	'ÿ':    0.000000,
}

func Score(input []byte) float64 {
	if len(input) < 1 {
		return math.Inf(-1)
	}

	return scoreBhattacharyya(expectedHist, calcHist(input))
}

func scoreBhattacharyya(p [256]float64, q [256]float64) float64 {
	n := len(p)
	var sum float64
	for i := 0; i < n; i++ {
		sum += math.Sqrt(p[i] * q[i])
	}
	return sum
}

func calcHist(a []byte) [256]float64 {
	var hist [256]float64
	for _, b := range a {
		hist[b]++
	}

	n := float64(len(a))
	for i := range hist {
		hist[i] /= n
		hist[i] *= 100
	}

	return hist
}
