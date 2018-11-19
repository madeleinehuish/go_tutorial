package main

import (
	"fmt"
)

func main () {
	intSlice := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("intSlice: ", intSlice)
	for i := 0; i < len(intSlice); i++ {
		if i%2 == 0 {
			fmt.Println(i, " even")
		} else {
			fmt.Println(i, " odd")
		}
	}
}