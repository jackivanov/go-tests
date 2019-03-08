package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	var gmailN []string
	for NItr := 0; NItr < int(N); NItr++ {
		firstNameEmailID := strings.Split(readLine(reader), " ")

		firstName := firstNameEmailID[0]

		emailID := firstNameEmailID[1]

		if ok, _ := regexp.MatchString(`[a-z]+@gmail.com$`, emailID); ok {
			gmailN = append(gmailN, firstName)
		}
	}
	sort.Strings(gmailN)

	for _, v := range gmailN {
		fmt.Println(v)
	}
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
