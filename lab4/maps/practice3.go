package main

import (
	"os"
	"fmt"
)


func main() {
	argv := os.Args

	m := make(map[string]struct{})

	for _, arg := range argv[1:] {
		m[arg] = struct{}{}
	}

	for arg := range m {
		fmt.Println(arg)
	}
}