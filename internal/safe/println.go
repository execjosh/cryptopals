package safe

import (
	"fmt"
	"unicode"
)

func Println(s []byte) {
	for _, b := range s {
		if unicode.IsPrint(rune(b)) {
			fmt.Printf("%c", b)
		} else {
			fmt.Printf("%#U", b)
		}
	}
	fmt.Println()
}
