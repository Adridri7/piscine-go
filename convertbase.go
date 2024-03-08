package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	base10 := AtoiBase(nbr, baseFrom)
	return PrintNbrBase(base10, baseTo)
}
