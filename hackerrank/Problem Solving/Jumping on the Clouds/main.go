package main

import "fmt"

func jumpingOnClouds(c []int32) {
	var Jumps int32
	for i, v := range c {
		if v == 1 || v == 0 {
			if v == 1 {
				Jumps++
				fmt.Println("Index:", i, "Cloud:", v, "Next cloud:", c[i+1], "Next next cloud:", c[i+2], "TOTAL:", Jumps)
				continue
			}

			if i+2 < len(c) {
				if c[i+1] == 0 && c[i+2] == 0 {
					fmt.Println("Index:", i, "Cloud:", v, "Next cloud:", c[i+1], "Next next cloud:", c[i+2], "TOTAL:", Jumps)
					c[i+1] = 2
					Jumps++
				}
			}

			if i == len(c)-1 {
				Jumps++
			}
		}
	}
	fmt.Println(Jumps)
}

func main() {
	x := []int32{0, 0, 1, 0, 0, 1, 0}
	// x := []int32{0, 0, 0, 1, 0, 0}
	// fmt.Println(jumpingOnClouds(x))
	jumpingOnClouds(x)
}
