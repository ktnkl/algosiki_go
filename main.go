package main

import "fmt"

func main() {
	list := &LinkedList{}
	list2 := &LinkedList{}
	list3 := &LinkedList{}

	list3.AppendInt(6)
	list3.AppendInt(2)

	list2.AppendInt(1)
	list2.AppendInt(3)
	list2.AppendInt(5)
	list2.AppendInt(7)

	list.AppendInt(3)
	list.AppendInt(3)
	list.AppendInt(5)
	list.AppendInt(7)

	list2.AppendList(list)
	list3.AppendList(list)

	intersection := getIntersectionNode(list2.Head, list3.Head)

	fmt.Println(intersection.String())
}
