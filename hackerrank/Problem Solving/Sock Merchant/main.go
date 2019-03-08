package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	var Pairs int32
	ColorsOfSocks := make(map[int32]int32)

	for _, v := range ar {
		ColorsOfSocks[v]++
	}

	for _, v := range ColorsOfSocks {
		if v > 1 {
			Pairs = Pairs + v/2
		}
	}
	return Pairs
}

func main() {
	x := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	fmt.Println(sockMerchant(10, x))
}
