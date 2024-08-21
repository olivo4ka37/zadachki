package BinarySearchTree

import "fmt"

var count = 0

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.insert(150)
	tree.insert(300)
	tree.insert(155)
	tree.insert(3)
	tree.insert(7)
	tree.insert(1337)
	tree.insert(1)
	tree.insert(0)
	tree.insert(141)
	tree.insert(228)
	tree.insert(1488)
	tree.insert(555)
	tree.insert(777)
	tree.insert(32)
	fmt.Println(tree)
	fmt.Println(tree.search(32))
	fmt.Println(count)
}

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

//Search

func (n *Node) search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.search(k)
	} else if n.Key > k {
		return n.Left.search(k)
	}

	return true
}
