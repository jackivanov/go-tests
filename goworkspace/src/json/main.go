package main

import (
	"encoding/json"
	"fmt"
)

type who struct {
	Name  string
	Sname string
	Age   int
}

func main() {
	WhatToUnmarshal := `[{"Name":"jack","Sname":"ivanov","Age":28},{"Name":"julia","Sname":"lisitcyna","Age":26}]`
	bs := []byte(WhatToUnmarshal)

	var UnmIvlis []who

	json.Unmarshal(bs, &UnmIvlis)

	fmt.Println(UnmIvlis)
}
