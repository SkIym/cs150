package main

import "strings"

func g(strs []string) [][]string {

	stringMatrix := [][]string{}

	for _, str := range strs {
		words := strings.SplitAfter(str, " ")
		splitString := []string{}
		for _, word := range words {
			splitString = append(splitString, strings.Title(word))
		}

		stringMatrix = append(stringMatrix, splitString)
	}

	return stringMatrix
}