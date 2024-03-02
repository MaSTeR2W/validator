package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, _ = strconv.ParseInt("922337203685477082", 10, 32)
	fmt.Println(int32(num))
}
