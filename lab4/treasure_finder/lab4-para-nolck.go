package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"strconv"
)

func main() {
	allArgs := os.Args
	if len(allArgs) != 3 {
		log.Fatal("MORE ARGS")
	}
	n := allArgs[1]
	number_n, err_n := strconv.Atoi(n)
	seed := allArgs[2]
	k := "2"
	num_k, err_k := strconv.Atoi(k)
	found := ""

	if err_n != nil || err_k != nil {
		log.Fatal(err_n)
	}

	wg := sync.WaitGroup{}

	// Create channel
	ch := make(chan string)
	doneWriting := sync.WaitGroup{}
	doneWriting.Add(1)

	go func() {
		for cell := range ch {
			if found != "" {
				found += ":"
			}
			found += cell
		}
		doneWriting.Done()
	}()

	// One goroutine per row
	for row := 0; row < number_n; row++ {
		wg.Add(1)
		cell := strconv.Itoa(row) + ",0"

		go func() {
			defer wg.Done()

			for {
				var out []byte
				var err error
				var exec_result *exec.Cmd

				// run the command
				exec_result = exec.Command("./treasure", n, seed, k, cell)
				out, err = exec_result.Output()

				if err != nil {
					log.Fatal(err)
				}

				// get the results of the command execution
				run_result := strings.Split(string(out), "\n")[:2]
				found_result := run_result[0]

				if found_result == "found" {
					ch <- cell
				}

				cellSplit := strings.Split(cell, ",")
				row, err_row := strconv.Atoi(cellSplit[0])
				col, err_col := strconv.Atoi(cellSplit[1])

				if err_row != nil || err_col != nil {
					log.Fatal(err_row)
				}

				// append depending if there's nearby or not
				if run_result[1] == "none" {
					col += num_k
				} else {
					col += 1
				}

				// if all row elements has been searched
				max_col, _ := strconv.Atoi(n)
				if col >= max_col {
					return
				}

				// Changing cell to search
				cell = strconv.Itoa(row) + "," + strconv.Itoa(col)
			}
		}()
	}

	wg.Wait()
	close(ch)
	doneWriting.Wait()

	fmt.Println(found)
}
