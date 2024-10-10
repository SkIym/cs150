package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(A, B int) int {
	return A + B
}

func subtract(A, B int) int {
	return A - B
}

func multiply(A, B int) int {
	return A * B
}

func main() {
	argv := os.Args

	// fmt.Println(argv)
	if len(argv) != 3 {
		fmt.Println("ERROR")
		return
	}

	A, err1 := strconv.Atoi(argv[1])
	B, err2 := strconv.Atoi(argv[2])

	if err1 != nil || err2 != nil {
		fmt.Println("ERROR")
		return
	}

	fmt.Println(add(A, B))
	fmt.Println(subtract(A, B))
	fmt.Println(multiply(A, B))
}
