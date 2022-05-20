package piscine

import "ft"

func PrintReverseAlphabet() {

	for i := 'z'; i >= 'a'; i-- {
		ft.PrintRune(rune(i))
	}
	ft.PrintRune('\n')
}
