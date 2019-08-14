package problem05

import "fmt"

// NodeList 链表节点
type NodeList struct {
	Value int
	Next  *NodeList
}

// PrintListTailToHead 从尾到头遍历链表
func PrintListTailToHead(head *NodeList) {
	if head != nil {
		PrintListTailToHead(head.Next)
		fmt.Println(head.Value)
	}
}
