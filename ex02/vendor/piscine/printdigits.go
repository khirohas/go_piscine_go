package piscine

import "ft"

func PrintDigits() {
	
	for i := 0; i < 10; i++ {
		ft.PrintRune(rune(int('0') + i))
	}
	ft.PrintRune('\n')
}