package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	t := ItoT(n)
	sortTab(t)
	for _, val := range t {
		z01.PrintRune(rune(val + '0'))
	}
}

func ItoT(n int) []int {
	tab := []int{}
	for n != 0 {
		tab = append(tab, n%10)
		n /= 10
	}
	return tab
}

func sortTab(tab []int) {
	for i := 0; i <= len(tab)-1; i++ {
		for j := 0; j <= len(tab)-1; j++ {
			if tab[j] > tab[i] {
				Swap(&tab[i], &tab[j])
			}
		}
	}
}
