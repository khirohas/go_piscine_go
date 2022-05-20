package piscine

import "ft"

func PrintAlphabet() {

	for i := 'a'; i <= 'z'; i++ {
		ft.PrintRune(rune(i))
	}
	ft.PrintRune('\n')
}
