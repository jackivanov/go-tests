package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func consecutiveOne(n int32) {
	binaryN := fmt.Sprintf("%b", n)
	count := make(map[int]int)
	index := 0
	for i := 0; i < len(binaryN); i++ {
		if binaryN[i] == 49 {
			count[index]++
		} else {
			index++
		}
	}

	biggest := count[0]
	for _, v := range count {
		if v > biggest {
			biggest = v
		}
	}

	fmt.Println(biggest)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	consecutiveOne(n)
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
