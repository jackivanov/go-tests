package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	var way, valleys int32
	for i := 0; i < len(s); i++ {

		if string(s[i]) == "U" {
			way++
		} else if string(s[i]) == "D" {
			way--
			if way == -1 {
				valleys++
			}
		}
	}

	return valleys
}

func main() {
	x := "UDDDUUUUDDDUU"
	fmt.Println(countingValleys(10, x))
	// countingValleys(10, x)
}
