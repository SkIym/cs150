package main

func f(m [][]int, k int) {
	for _, row := range m {
		for i, _ := range row {
			row[i] = row[i] * k
		}
	}
}

