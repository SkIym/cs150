package main

// Marcelo, Abram Josh C.
// 2021-12540

import (
	// "fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	head := strNode{"hello", nil, nil}
	mypq := strPriorityQueue{&head}

	if mypq.head == nil {
		t.Error("Head should contain string data")
	}

}

func TestEnqueue(t *testing.T) {
	t.Run("Enqueuing in an empty pq", func(t *testing.T) {
		mypq := strPriorityQueue{nil}

		mypq.enqueue("a")

		if mypq.head == nil {
			t.Error("Head should be filled if priority queue is empty")
		}

		if mypq.head.data != "a" {
			t.Error("Head contains the wrong data when priority queue is previously empty")
		}
	})
	t.Run("Enqueuing in a pq with only one element (head) and the string to be enqueued is lexicographically larger than head", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("a")
		mypq.enqueue("b")

		if mypq.head.right == nil {
			t.Error("Head should point right to a node with lexicographically larger string")
		}

		if mypq.head.right.data != "b" {
			t.Error("Fails when the node to be enqueued is a lexicographically larger string")
		}

		if mypq.head.right.left == nil && mypq.head.right.left.data != "a" {
			t.Error("Newly added node should point left to head")
		}

	})
	t.Run("Enqueuing in a pq with only one element (head) and the string to be enqueued is lexicographically smaller than head (hence smallest)", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("h")
		mypq.enqueue("b")

		// fmt.Println("hello at test", mypq.head.data, mypq.head.right.data)

		if mypq.head == nil {
			t.Error("Head is missing")
		}

		if mypq.head.data != "b" {
			t.Error("Head should be the newly enqueued node")
		}

		if mypq.head.right == nil {
			t.Error("New head should be pointing right to the previous node")
		}

		if mypq.head.right.data != "h" {
			t.Error("New head should be pointing right to the previous head")
		}

		if mypq.head.right.left == nil || mypq.head.right.left.data != "b" {
			t.Error("Previous head should be pointing left to new head (newly added node)")
		}

	})

	t.Run("Enqueuing in a pq with several elements and the string to be enqueued is lexicographically the largest", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("f")
		mypq.enqueue("o")
		mypq.enqueue("c")
		mypq.enqueue("y")

		if mypq.head == nil || mypq.head.data != "c" {
			t.Error("Head should not change")
		}

		temp := mypq.head
		for temp.right != nil {
			temp = temp.right
		}

		if temp.data != "y" {
			t.Error("Tail should be the lexicographically largest string")
		}

		if temp.left == nil {
			t.Error("Newly enqueued node should point left to another node")
		}

		if temp.left.data != "o" {
			t.Error("Newly enqueued node should point left to previous lexicographically largest string")
		}

	})

	t.Run("Enqueuing in a pq with several elements and the string to be enqueued is neither the smallest nor the largest lexicographically", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("j")
		mypq.enqueue("b")
		mypq.enqueue("z")
		mypq.enqueue("w")
		mypq.enqueue("l")

		if mypq.head == nil || mypq.head.data != "b" {
			t.Error("Head should not change")
		}

		temp := mypq.head

		for temp != nil {
			if temp.data == "l" {
				break
			}
			temp = temp.right
		}

		if temp == nil {
			t.Error("Failed to enqueue string somewhere at the middle of the queue")
			return
		}

		if temp.right == nil {
			t.Error("Newly enqueued node should point right to another node")
			return
		}

		if temp.right.data != "w" {
			t.Error("String was incorrectly queued")
		}

		if temp.left == nil {
			t.Error("Newly enqueued node should point left to another node")
			return
		}

		if temp.left.data != "j" {
			t.Error("String was incorrectly queued")
		}

	})

}

func TestDequeue(t *testing.T) {
	t.Run("Dequeuing an empty queue", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		ret, err := mypq.dequeue()
		if err == nil {
			t.Error("It should return an error")
		}

		if ret != "" {
			t.Error("It should return an empty string")
		}
	})

	t.Run("Dequeuing a queue with a single element (head)", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("NEVER GONNA GIVE YOU UP")
		ret, err := mypq.dequeue()

		if mypq.head != nil {
			t.Error("Priority queue must be empty")
		}

		if ret != "NEVER GONNA GIVE YOU UP" {
			t.Error("Should return the data in head")
		}

		if err != nil {
			t.Error("Should not return an error")
		}

	})

	t.Run("Dequeuing a filled queue", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("q")
		mypq.enqueue("f")
		mypq.enqueue("b")
		mypq.enqueue("t")
		mypq.enqueue("m")

		ret, err := mypq.dequeue()

		if ret == "" {
			t.Error("Should return a string")
		}

		if err != nil {
			t.Error("Should not return an error")
		}

		if ret != "b" {
			t.Error("Should return the data in head")
		}

		if mypq.head == nil {
			t.Error("Head should still be present")
			return
		}

		if mypq.head.data == "b" {
			t.Error("Dequeued string should be in head")
		}

		if mypq.head.data != "f" {
			t.Error("New head should be updated to the next in line node")
		}

		if mypq.head.left != nil {
			t.Error("Head should still point left to nothing")
		}

	})
}

func TestLength(t *testing.T) {
	t.Run("Length of an empty queue", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		length := mypq.length()

		if length != 0 {
			t.Error("Length should be zero")
		}
	})

	t.Run("Length of a queue with a single element", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("AND HURT YOUUUU")
		length := mypq.length()

		if length != 1 {
			t.Error("Length should only be one")
		}
	})

	t.Run("Length of a queue with multiple elements", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("DESERT YOUUUU")
		mypq.enqueue("and at last i see the light")
		mypq.enqueue("and it's like the the fog has lifted")
		mypq.enqueue("REESES PUFFS")
		length := mypq.length()

		if length != 4 {
			t.Error("Length should accurately reflect the number of elements in the queue")
		}
	})

	t.Run("Length after removing the sole element", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("DESERT YOUUUU")
		mypq.dequeue()
		length := mypq.length()

		if length == 1 {
			t.Error("Length should now be back to zero")
		}
	})

	t.Run("Length after removing one element of many", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("DESERT YOUUUU")
		mypq.enqueue("and at last i see the light")
		mypq.enqueue("and it's like the the fog has lifted")
		mypq.enqueue("REESES PUFFS")
		old_length := mypq.length()
		mypq.dequeue()
		new_length := mypq.length()

		if old_length == new_length {
			t.Error("Length should be decreased by one")
		}
	})

	t.Run("Length after removing all elements", func(t *testing.T) {
		mypq := strPriorityQueue{nil}
		mypq.enqueue("DESERT YOUUUU")
		mypq.enqueue("and at last i see the light")
		mypq.enqueue("and it's like the the fog has lifted")
		mypq.enqueue("REESES PUFFS")
		old_length := mypq.length()
		mypq.dequeue()
		mypq.dequeue()
		mypq.dequeue()
		mypq.dequeue()
		new_length := mypq.length()

		if old_length == new_length {
			t.Error("Length should be decreased")
		}

		if new_length != 0 {
			t.Error("Length should be zero")
		}
	})
}
