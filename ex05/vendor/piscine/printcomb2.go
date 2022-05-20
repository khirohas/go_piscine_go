package piscine

import "ft"

func PrintComb2() {

	f := [2]int{0, 0}
	s := [2]int{0, 0}

	for f[0] = 0; f[0] <= 9; f[0]++ {
		for f[1] = 0; f[1] <= 9; f[1]++ {
			for s[0] = 0; s[0] <= 9; s[0]++ {
				for s[1] = 0; s[1] <= 9; s[1]++ {
					if f != s && isMorethan(s, f) {
						outputComb(f, s)
					}
				}
			}
		}
	}
}

func isMorethan(s, f [2]int) bool {

	if s[0] > f[0] || (s[0] == f[0] && s[1] > f[1]) {
		return (true)
	} else {
		return (false)
	}
}

func outputComb(f, s [2]int) {
	ft.PrintRune((rune)('0' + f[0]))
	ft.PrintRune((rune)('0' + f[1]))
	ft.PrintRune(' ')
	ft.PrintRune((rune)('0' + s[0]))
	ft.PrintRune((rune)('0' + s[1]))
	if !(f[0] == 9 && f[1] == 8) {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	} else {
		ft.PrintRune('\n')
	}

}
