package main

import "fmt"

func main() {
	var num int
	var pow int
	var result = 1

	num = 2
	pow = 4
	for i := 0; i < pow; i++ {
		result *= num
	}

	fmt.Println("%dの%d乗は%dです。\n", num, pow, result)
}
