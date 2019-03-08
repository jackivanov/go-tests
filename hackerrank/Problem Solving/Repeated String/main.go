package main

import (
	"fmt"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	return int64(int64(strings.Count(s, "a"))*(n/int64(len(s))) + int64(strings.Count(s[:n%int64(len(s))], "a")))
}

func main() {
	x := "epsxyyflvrrrxzvnoenvpegvuonodjoxfwdmcvwctmekpsnamchznsoxaklzjgrqruyzavshfbmuhdwwmpbkwcuomqhiyvuztwvq"
	fmt.Println(repeatedString(x, 549382313570))
	// repeatedString(x, 549382313570)
}
