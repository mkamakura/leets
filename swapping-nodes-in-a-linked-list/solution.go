package swapping_nodes_in_a_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var toArray []int
	node := head
	for node != nil {
		toArray = append(toArray, node.Val)
		node = node.Next
	}
	fmt.Println(toArray)

	val1 := toArray[k - 1]
	val2 := toArray[len(toArray) - k]
	toArray[k - 1] = val2
	toArray[len(toArray) - k] = val1

	fmt.Println(toArray)

	// make linked list
	newHead := &ListNode{}
	for i := 0; i < len(toArray); i++ {
		newNode := ListNode{}
		newNode.Val = toArray[len(toArray) - i - 1]
		if i != 0 {
			newNode.Next = newHead
			newHead = &newNode
		} else {
			newHead = &newNode
		}
	}

	return newHead
}
