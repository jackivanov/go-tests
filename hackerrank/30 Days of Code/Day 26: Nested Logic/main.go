package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func whatFine(r, e string) int {
	if e == r {
		return 0
	}

	re := strings.Split(r, " ")
	ex := strings.Split(e, " ")

	var R, E [3]int64

	for i := 0; i < 3; i++ {
		R[i], _ = strconv.ParseInt(re[i], 10, 64)
		E[i], _ = strconv.ParseInt(ex[i], 10, 64)
	}

	if R[2] > E[2] {
		return 10000
	}

	if R[2] == E[2] && R[1] > E[1] {
		return int((R[1] - E[1]) * int64(500))
	}

	if R[2] == E[2] && R[1] == E[1] && R[0] > E[0] {
		return int((R[0] - E[0]) * int64(15))
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}

	r := arr[0]
	e := arr[1]

	fmt.Println(whatFine(r, e))

	// whatFine(r, e)
}
