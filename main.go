package main

import "fmt"

func IsNil(a any) bool {
	switch a.(type) {
	case any:
		return false
	case nil:
		return true
	default:
		return false
	}
}

func main() {
	var a = map[string]any{}
	fmt.Println(IsNil(a["hee"]))
}
