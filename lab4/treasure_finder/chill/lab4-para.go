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

// disregards metal detector

func searchRow(cell string, n string, seed string, k string, found *string, lock chan struct{}) {

	local_found := ""

	for {
		var out []byte
		var err error
		var exec_result *exec.Cmd

		if local_found == "" {
			exec_result = exec.Command("./treasure", n, seed, k, cell)
		} else {
			exec_result = exec.Command("./treasure", n, seed, k, cell, local_found)
		}
		out, err = exec_result.Output()

		if err != nil {
			log.Fatal(err)
		}

		run_result := strings.Split(string(out), "\n")[:2]
		found_result := run_result[0]

		if found_result == "found" {
			if local_found != "" {
				local_found += ":"
			}
			local_found += cell

			// avoid overwrite
			lock <- struct{}{}
			if *found != "" {
				*found += ":"
			}
			*found += cell
			<-lock
		}

		cellSplit := strings.Split(cell, ",")
		row, err_row := strconv.Atoi(cellSplit[0])
		col, err_col := strconv.Atoi(cellSplit[1])

		if err_row != nil || err_col != nil {
			log.Fatal(err_row)
		}

		// if all row elements has been searched
		col += 1
		max_col, _ := strconv.Atoi(n)
		if col == max_col {
			return
		}

		// Changing cell to search
		cell = strconv.Itoa(row) + "," + strconv.Itoa(col)
		fmt.Println("Goroutine assigned to row" + cellSplit[0])
		fmt.Println("new cell:", cell)
		fmt.Println("found:", local_found)

	}
}

func main() {
	allArgs := os.Args
	if len(allArgs) != 3 {
		log.Fatal("MORE ARGS")
	}
	n := allArgs[1]
	number_n, err_n := strconv.Atoi(n)
	seed := allArgs[2]
	k := "2"
	_, err_k := strconv.Atoi(k)
	found := ""

	if err_n != nil || err_k != nil {
		log.Fatal(err_n)
	}

	wg := sync.WaitGroup{}

	// Create channel for locking (appending values of found)
	lock := make(chan struct{}, 1)

	// One goroutine per row
	for row := 0; row < number_n; row++ {
		cell := strconv.Itoa(row) + ",0"
		wg.Add(1)
		go func(c string) {
			defer wg.Done()
			fmt.Println(c)
			searchRow(c, n, seed, k, &found, lock)
		}(cell)
	}

	wg.Wait()
	close(lock)

	fmt.Println(found)
}
