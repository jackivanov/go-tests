package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkNumber(n int32) {
	if n%2 != 0 {
		fmt.Println("Weird")
	} else if n%2 == 0 && n >= 2 && n <= 5 {
		fmt.Println("Not Wierd")
	} else if n%2 == 0 && n >= 6 && n <= 20 {
		fmt.Println("Wierd")
	} else if n%2 == 0 && n > 20 {
		fmt.Println("Not Wierd")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	checkNumber(N)
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
