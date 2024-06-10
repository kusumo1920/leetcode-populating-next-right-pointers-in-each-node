package main

import "fmt"

func main() {
	input := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	output := connect(input)
	fmt.Println(output)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	mapper := make(map[int]*Node)
	var recursiveFn func(node *Node, lvl int)
	recursiveFn = func(node *Node, lvl int) {
		if node == nil {
			return
		}
		if node.Left != nil {
			_, ok := mapper[lvl]
			if !ok {
				mapper[lvl] = node.Left
			} else {
				mapper[lvl].Next = node.Left
				mapper[lvl] = mapper[lvl].Next
			}
		}
		if node.Right != nil {
			_, ok := mapper[lvl]
			if !ok {
				mapper[lvl] = node.Right
			} else {
				mapper[lvl].Next = node.Right
				mapper[lvl] = mapper[lvl].Next
			}
		}
		recursiveFn(node.Left, lvl+1)
		recursiveFn(node.Right, lvl+1)
	}
	recursiveFn(root, 1)

	return root
}
