package piscine

func StrRev(s string) string {
	var revstr string
	for _, val := range s {
		revstr = string(val) + revstr
	}
	return revstr
}
