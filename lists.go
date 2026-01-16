package main

// Работа со связанными списками

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) AppendInt(value int) {
	newNode := &ListNode{Val: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	}

	ll.Tail.Next = newNode
	ll.Tail = newNode

	ll.Size++
}

func (ll *LinkedList) AppendList(value *LinkedList) {

	if ll.Head == nil {
		// ll.Head = value.Head
		// ll.Tail = value.Tail
		ll = value
	}

	ll.Tail.Next = value.Head
	ll.Tail = value.Tail

	ll.Size += value.Size
}

func (ll *LinkedList) Length() int {
	return ll.Size
}

// Решение задачи 21 с литкода. Слияние сортирвоанных списков, решение, используя struct

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

// Решение задачи 160. Intersection of Two Linked Lists

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptrA, ptrB := headA, headB

	for ptrA != ptrB {
		if ptrA.Next != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = headB
		}
		if ptrB.Next != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = headA
		}
	}

	return ptrA
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head

	for i := 0; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

func (node *ListNode) String() string {
	var result strings.Builder

	for node != nil {
		result.WriteString(strconv.Itoa(node.Val))
		if node.Next != nil {
			result.WriteString(" -> ")
		}
		node = node.Next
	}

	return result.String()
}

// func (node *ListNode) Merge(toMerge *ListNode) *ListNode {
// 	if node == nil {
// 		return toMerge
// 	}

// 	for
// }
