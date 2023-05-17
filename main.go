
package main

/*
Problem #107
Print the nodes in a binary tree level-wise. For example, the following should print 1, 2, 3, 4, 5.

  1
 / \
2   3
   / \
  4   5

Main:
	node2  := NewNode(2, nil, nil)
	node4  := NewNode(4, nil, nil)
	node5  := NewNode(5, nil, nil)

	node3  := NewNode(3, node4, node5)
	root  := NewNode(1, node2, node3)

	ReadBinaryByLowest(root)
*/
type Node struct{
	Data int
	Left *Node
	Right *Node
}

func NewNode(data int, left, right *Node) *Node{
	return &Node{
		Data: data,
		Left: left,
		Right: right,
	}
}

func ReadBinaryByLowest(node *Node) string {

	fmt.Printf("%d, ", node.Data)

	if (node.Left != nil){
		ReadBinaryByLowest(node.Left)
	}else{
		return ""
	}

	return ReadBinaryByLowest(node.Right)
}
