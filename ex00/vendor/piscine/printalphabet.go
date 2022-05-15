package piscine

import "ft"

func PrintAlphabet() {
	
	for i := 0; i < 26; i++ {
		ft.PrintRune(rune(int('a') + i))
	}
	ft.PrintRune('\n')
}