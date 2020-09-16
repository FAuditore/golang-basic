package main

import "fmt"

func main() {
	var a = 1

LOOP:
	for a < 10 {
		if a == 5 {
			a++
			goto LOOP
		}
		fmt.Printf("%d ", a)
		a++
	}
}
