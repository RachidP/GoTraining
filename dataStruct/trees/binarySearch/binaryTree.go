package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	key    int64
	left   *Node //left Child
	right  *Node //right Child
	father *Node // parent
}

type Root struct {
	root *Node
}

func main() {
	tree := &Root{}
	tree.insert(10)
	tree.insert(20)
	tree.insert(7)
	tree.insert(15)
	tree.insert(30)
	tree.insert(8)
	tree.insert(6)

	fmt.Println(tree.root)

	fmt.Printf("%#v \n", tree.root)

	print(os.Stdout, tree.root, 0, 'R')

	/*search key*/
	var key int64
	key = 20
	z := tree.root.searchKey(key)
	if z == nil {
		fmt.Printf("the key = %v doesn't find in the tree \n", key)

	} else {

		fmt.Printf("key=%v finded in:  %+v \n", key, z)
		/*find Successor of the node x in the tree*/
		fmt.Printf("Successor of the key(%v) is:  %+v \n", key, z.findSuccessor())
	}

	/*find Successor of the node x in the tree*/
	z = tree.root.searchKey(key)
	if z == nil {
		fmt.Printf("the key = %v doesn't find in the tree \n", key)
	} else {
		fmt.Printf("Successor of the key(%v) is:  %+v \n", key, z.findSuccessor())
	}

	/*find minimum in the tree*/
	fmt.Printf("Minimum in the tree is:  %+v \n\n", tree.root.findMinimum())

	/*find Maximum in the tree*/
	fmt.Printf("Maximum in the tree is:  %+v \n\n", tree.root.findMaximum())

	/*delete the node x from the tree*/
	z = tree.root.searchKey(key)
	if z == nil {
		fmt.Printf("the key = %v doesn't find in the tree \n", key)

	} else {
		print(os.Stdout, tree.root, 0, 'S')
		z.delete()
		fmt.Printf("\nthe Node= %v will be deleted \n", z.key)
		print(os.Stdout, tree.root, 0, 'S')

	}

}

func (r *Root) insert(data int64) *Root {
	if r.root == nil {
		r.root = &Node{key: data, left: nil, right: nil, father: nil}

	} else {

		//r.root.insertR(data)
		r.root.insert(data)
	}
	return r

}

func (x *Node) insert(data int64) {
	if x == nil {
		return
	}
	//x := T.root
	var y *Node

	for x != nil {

		y = x

		if data < x.key {
			x = x.left
		} else {
			x = x.right
		}

	}

	if data < y.key {
		y.left = &Node{key: data, left: nil, right: nil, father: y}

	} else {
		y.right = &Node{key: data, left: nil, right: nil, father: y}

	}

}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.key)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func (x *Node) insertR(data int64) bool {
	isWrite := false
	if x == nil {
		return isWrite
	}

	if data < x.key {
		isWrite = x.left.insertR(data)

	} else {
		isWrite = x.right.insertR(data)
	}
	if !isWrite {
		if data < x.key {
			x.left = &Node{key: data, left: nil, right: nil}

		} else {
			x.right = &Node{key: data, left: nil, right: nil}

		}
	}
	return true
}

/*Search a key in the tree*/
func (x *Node) searchKey(key int64) *Node {

	if x == nil || key == x.key {

		return x
	}
	if key < x.key {
		return x.left.searchKey(key)
	}
	return x.right.searchKey(key)

}

/*find the minimun in the tree*/
func (x *Node) findMinimum() *Node {

	if x.left != nil {

		x = x.left.findMinimum()
	}
	return x

}

/*find the Maximum in the tree*/
func (x *Node) findMaximum() *Node {

	if x.right != nil {

		x = x.right.findMaximum()
	}
	return x

}

/*find the successor for the node X in the tree*/
func (x *Node) findSuccessor() *Node {

	if x.right != nil {

		x = x.right.findMinimum()
		return x
	}
	y := x.father

	for y != nil && x == y.right {
		x = y
		y = y.father
	}
	return y

}

/*TRANSPLANT, which replaces one subtree as a child of its parent with
another subtree*/
func (x *Node) transplant(u, v *Node) {
	if u.father == nil {
		x = v
	} else {
		if u == u.father.left {
			u.father.left = v
		} else {
			u.father.right = v
		}

	}

	if v != nil {
		v.father = u.father

	}

}

/*delete a Node in the tree*/
func (x *Node) delete() {

	if x.left == nil {
		x.transplant(x, x.right)

	} else {

		if x.right == nil {
			x.transplant(x, x.left)

		} else {

			y := x.right.findMinimum() //????
			if y.father != x {
				x.transplant(y, y.right)
				y.right = x.right
				y.right.father = y
			}
			x.transplant(x, y)
			y.left = x.left
			y.left.father = y
		}
	}

}
