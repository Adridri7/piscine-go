package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	str_info :=
		`--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.`

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Println(str_info)
	} else if len(args) == 1 && args[0][0:9] != "--insert" && args[0][0:3] != "-i" && args[0][0:8] != "--order" && args[0][0:3] != "-o" {
		fmt.Println("caca")
	}
}
