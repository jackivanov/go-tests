package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func swap(x, y int32) (int32, int32) {
	return y, x
}

func bubbleSort(a []int32) {
	var numberOfSwaps int32
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = swap(a[j], a[j+1])
				numberOfSwaps++
			}
		}

		if numberOfSwaps == 0 {
			break
		}
	}

	fmt.Println("Array is sorted in", numberOfSwaps, "swaps.")
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[len(a)-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	bubbleSort(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
