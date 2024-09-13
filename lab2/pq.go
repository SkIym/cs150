package main

import (
	"errors"
	// "fmt"
	"strings"
)

type strPriorityQueue struct {
	head *strNode
}

type strNode struct {
	data  string
	left  *strNode
	right *strNode
}

func (pq *strPriorityQueue) enqueue(str string) {

	if pq.head == nil {
		pq.head = &strNode{str, nil, nil}
		return
	}

	newNode := strNode{str, nil, nil}

	if strings.Compare(str, pq.head.data) < 0 {
		pq.head.left = &newNode
		newNode.right = pq.head
		pq.head = &newNode
		// fmt.Println("hello at main", pq.head.data, pq.head.right.data)
		return
	}

	temp := pq.head
	for strings.Compare(temp.data, str) < 0 {

		if temp.right == nil {
			break
		}

		temp = temp.right
	}

	if temp.right == nil && strings.Compare(temp.data, str) < 0 {
		temp.right = &newNode
		newNode.left = temp
		// fmt.Println("Adding at the tail")
	} else {
		temp.left.right = &newNode
		newNode.left = temp.left
		temp.left = &newNode
		newNode.right = temp
		// fmt.Println("Adding in between!")
	}

}

func (pq *strPriorityQueue) dequeue() (string, error) {

	if pq.head == nil {

		return "", errors.New("cannot dequeue: string is alrady empty")
	}

	temp := pq.head

	if pq.head.right == nil {
		pq.head = nil
	} else {
		pq.head = pq.head.right
		pq.head.left = nil
	}

	return temp.data, nil

}

func (pq *strPriorityQueue) length() int {

	if pq.head == nil {
		return 0
	}

	ctr := 0

	temp := pq.head
	for temp != nil {
		ctr += 1
		temp = temp.right
	}

	return ctr
}
