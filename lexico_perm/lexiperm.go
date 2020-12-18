package main

import (
	"fmt"
)

func getDigits(n int) []int {
	s := []int{}

	for n > 0 {
		d := n % 10
		s = append([]int{d}, s...)
		n /= 10
	}

	return s
}

func getLexicoPerm(n int) int {
	s := getDigits(n)

	fmt.Println("Digits are : ", s)

	return 0
}

func main() {
	fmt.Println("In main")

	num := 1234

	fmt.Printf("The number is: %d\n", num)

	fmt.Println("The next lexico perm is : ", getLexicoPerm(num))
}
