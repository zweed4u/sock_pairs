package main

import "fmt"

func sockMerchant(n int, ar []int) (pairs int) {
	colorsSeen := make(map[int]int)
	for _, sockColorID := range ar {
		_, ok := colorsSeen[sockColorID]
		if !ok {
			colorsSeen[sockColorID] = 1
		} else {
			colorsSeen[sockColorID] += 1
		}
	}

	for _, sockColorAmount := range colorsSeen {
		pairs += sockColorAmount / 2
	}
	fmt.Println(pairs)
	return
}

func main() {
	arr1 := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	sockMerchant(len(arr1), arr1)

	arr2 := []int{1, 2, 1, 2, 1, 3, 2}
	sockMerchant(len(arr2), arr2)
}
