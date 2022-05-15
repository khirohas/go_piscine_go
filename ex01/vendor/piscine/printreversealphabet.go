package piscine

import "ft"

func PrintReverseAlphabet() {
	
	for i := 0; i < 26; i++ {
		ft.PrintRune(rune(int('z') - i))
	}
	ft.PrintRune('\n')
}