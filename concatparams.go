package piscine

func ConcatParams(args []string) (str string) {

	for idx, arg := range args {
		str += arg
		if idx != len(args)-1 {
			str += "\n"
		}
	}
	return
}
