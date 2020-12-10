package main

/*
   Given an array of numbers representing the stock prices of a company in chronological order,
   write a function that calculates the maximum profit you could have made from buying and
   selling that stock once. You must buy before you can sell it.

   For example, given [9, 11, 8, 5, 7, 10], you should return 5,
   since you could buy the stock at 5 dollars and sell it at 10 dollars.
*/
/*
	Assumption : solution works only with list of +ve integers!!
*/

import (
	"errors"
	"fmt"
)

func getFirstMinElementIndex(v *[]int) (int, error) {
	if len(*v) < 2 {
		return 0, errors.New("Array length is less than 2")
	}

	min := (*v)[0]
	idx := 0
	for i := 1; i < len(*v); i++ {
		if (*v)[i] < min {
			min = (*v)[i]
			idx = i
		}
	}

	return idx, nil
}

func getMaxElementAfterIndex(v *[]int, idx int) (int, error) {
	if len(*v) < idx {
		return 0, errors.New("Index out of range")
	}

	max := (*v)[idx]
	for i := idx + 1; i < len(*v); i++ {
		if (*v)[i] > max {
			max = (*v)[i]
		}
	}

	if max == (*v)[idx] {
		return 0, errors.New("Unable to find element greater than given element")
	}

	return max, nil
}

func main() {
	v := []int{9, 11, 8, 5, 7, 10}

	if idx, err := getFirstMinElementIndex(&v); err == nil {
		fmt.Println("Min value : ", v[idx])
		if e, err := getMaxElementAfterIndex(&v, idx); err == nil {
			fmt.Println("Max Difference : ", e-v[idx])
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
