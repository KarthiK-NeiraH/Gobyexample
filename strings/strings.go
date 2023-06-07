package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "கார்த்திக் "

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)

	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("fount tee")
	} else if r == 'a' {
		fmt.Println("found so sua")
	}
}
