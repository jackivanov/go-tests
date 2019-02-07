package main

import (
	"fmt"
)

type guys struct {
	name string
	age int
}

type fuckers struct {
	guys
	stupid bool
}

func main() {
	fuckers := fuckers{
      guys: guys{
  			name: "jack",
  			age: 28,
      },
			stupid: false,
	}

	fmt.Println(fuckers)
}
