package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputString := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, World.")
	for inputString.Scan() {
		fmt.Println(inputString.Text())
	}
}
