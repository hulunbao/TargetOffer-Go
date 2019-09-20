package problem16

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	fmt.Println("")
	print(ReverseList(nil))
	fmt.Printf("\n")

	l1 := &NodeList{1, nil}
	fmt.Println("1 -> ")
	print(ReverseList(l1))
	fmt.Printf("\n")

	l3 := &NodeList{3, nil}
	l2 := &NodeList{2, l3}
	l1 = &NodeList{1, l2}
	fmt.Println("1 -> 2 -> 3")
	print(ReverseList(l1))
	fmt.Printf("\n")
}

func TestReverseList1(t *testing.T) {
	fmt.Println("")
	print(ReverseList1(nil))
	fmt.Printf("\n")

	l1 := &NodeList{1, nil}
	fmt.Println("1 -> ")
	print(ReverseList1(l1))
	fmt.Printf("\n")

	l3 := &NodeList{3, nil}
	l2 := &NodeList{2, l3}
	l1 = &NodeList{1, l2}
	fmt.Println("1 -> 2 -> 3")
	print(ReverseList1(l1))
	fmt.Printf("\n")
}

func TestReverseList2(t *testing.T) {
	fmt.Println("")
	print(ReverseList2(nil))
	fmt.Printf("\n")

	l1 := &NodeList{1, nil}
	fmt.Println("1 -> ")
	print(ReverseList2(l1))
	fmt.Printf("\n")

	l3 := &NodeList{3, nil}
	l2 := &NodeList{2, l3}
	l1 = &NodeList{1, l2}
	fmt.Println("1 -> 2 -> 3")
	print(ReverseList2(l1))
	fmt.Printf("\n")
}
