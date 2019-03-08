package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// IsPrime check whether a given integer is a prime number
func IsPrime(value int) string {
	if value <= 1 {
		return "Not prime"
	}

	var m float64 = 1
	if value > 100000 {
		m = 1000
	}

	for i := 2; i <= int(math.Floor(float64(value)/m/2)); i++ {
		if value%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}

func checkPrime(a []int) {
	for i := range a {
		fmt.Println(IsPrime(a[i]))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var arrTmp, arr []int

	for scanner.Scan() {
		w, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		arrTmp = append(arrTmp, int(w))

		arr = arrTmp[1:]

	}
	checkPrime(arr)
}
