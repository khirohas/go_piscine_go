package piscine

import "ft"

func PrintComb() {

	f, s, t := 0, 0, 0
	for f = 0; f <= 9; f++ {
		for s = 0; s <= 9; s++ {
			for t = 0; t <= 9; t++ {
				if f < s && s < t {
					outputComb(f, s, t)
				}
			}
		}
	}
}

func outputComb(f, s, t int) {
	ft.PrintRune((rune)('0' + f))
	ft.PrintRune((rune)('0' + s))
	ft.PrintRune((rune)('0' + t))
	if !(f == 7 && s == 8 && t == 9) {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	} else {
		ft.PrintRune('\n')
	}
}
