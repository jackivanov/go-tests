package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	type double float64
	type String string

	var i2 int64
	var d2 double
	var s2 String

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// first line
	i2, _ = strconv.ParseInt(input[0], 10, 32)
	fmt.Println(i + uint64(i2))

	// second line
	d2i, _ := strconv.ParseFloat(input[1], 64)
	d2 = double(d2i)
	fmt.Printf("%0.1f\n", d+float64(d2))

	// third line
	s2 = String(input[2])
	fmt.Println(s + string(s2))

}
