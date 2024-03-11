package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) == 1 {
		file, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(file))
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	}
}
