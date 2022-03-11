package main

import "fmt"

const tii string = "Mushfiqur"

func main() {
	fmt.Println(tii)

	const n = 50000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
