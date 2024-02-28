package piscine

import "fmt"

func PrintComb(){
  for i := 0; i < 10; i++ {
        for j := i + 1; j < 10; j++ {
            for k := j + 1; k < 10; k++ {
                fmt.Printf("%d%d%d", i, j, k)
                if !(i == 7 && j == 8 && k == 9) {
                    fmt.Print(", ")
                }
            }
        }
    }
}