package main

import (
	"fmt"
	"unicode/utf8"
)

/*
A Go string is a read-only slice of bytes.
The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
*/
func main() {
	// “hello” in the Thai language
	const s3 = "ส"
	const s2 = "สวั"
	const s1 = "สวัส"
	const s = "สวัสดี"

	// the length of the raw bytes stored within
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	// Some Thai characters are represented by multiple UTF-8 code points,
	// so the result of this count may be surprising
	fmt.Println("Rune count:", utf8.RuneCountInString(s3)) // 1
	fmt.Println("Rune count:", utf8.RuneCountInString(s2)) // 3
	fmt.Println("Rune count:", utf8.RuneCountInString(s1)) // 4
	fmt.Println("Rune count:", utf8.RuneCountInString(s))  // 6

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
	// Values enclosed in single quotes are rune literals.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
