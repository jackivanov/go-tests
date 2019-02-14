package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var even, odd string
		w := scanner.Text()
		if len(w) > 2 {
			for i := 0; i < len(w); i++ {
				if i%2 == 0 {
					even += string(w[i])
				} else {
					odd += string(w[i])
				}
			}
			fmt.Println(even, odd)
		}
	}
}
