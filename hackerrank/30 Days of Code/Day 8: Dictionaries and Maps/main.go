package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	index, _ := strconv.ParseInt(input[0], 10, 64)

	phoneBook := make(map[string]int64)

	for i := 1; i <= int(index); i++ {
		line := strings.Split(input[i], " ")
		phone, _ := strconv.ParseInt(line[1], 10, 64)
		phoneBook[line[0]] = phone
	}

	for i := index + 1; i < int64(len(input)); i++ {

		if v, ok := phoneBook[input[i]]; ok {
			fmt.Print(input[i], "=", v, "\n")
		} else {
			fmt.Println("Not found")
		}
	}
}
