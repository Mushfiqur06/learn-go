package main

import "fmt"

func main() {
	var i int = 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for j := 2; j <= 7; j++ {
		fmt.Println(j)
	}
}
