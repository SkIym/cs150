package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	A := strconv.Itoa(rand.Intn(101))
	B := strconv.Itoa(rand.Intn(101))
	cmnd := exec.Command("./helper", A, B)
	output, err := cmnd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	outputs := strings.Split(string(output), "\n")
	ints := outputs[:3]
	add, _ := strconv.Atoi(ints[0])
	sub, _ := strconv.Atoi(ints[1])
	mult, _ := strconv.Atoi(ints[2])
	sum := add + sub + mult
	fmt.Println(A, B)
	fmt.Println(ints, sum)
}
