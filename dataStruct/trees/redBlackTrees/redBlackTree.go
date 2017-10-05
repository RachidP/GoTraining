package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	key    int64
	color  rune
	left   *Node //left Child
	right  *Node //right Child
	father *Node // parent
}

type Root struct {
	root *Node
}

func main() {
	tree := &Root{}
	sentinelT := &Node{key: -1, color: 'B', left: nil, right: nil, father: nil}
	//tree.root= &Node{key: -1, color: 'B', left: sentinelT, right: sentinelT, father: sentinelT}

	tree.insertMy(sentinelT, 10)
	tree.insertMy(sentinelT, 20)
	tree.insertMy(sentinelT, 30)
	tree.insertMy(sentinelT, 15)
	tree.insertMy(sentinelT, 50)
	tree.insertMy(sentinelT, 34)
	tree.insertMy(sentinelT, 55)
	tree.insertMy(sentinelT, 1)
	tree.insertMy(sentinelT, 33)
	tree.insertMy(sentinelT, 32)
	tree.insertMy(sentinelT, 31)
	tree.insertMy(sentinelT, 84)

	/*tree.insertMy(sentinelT, 20)
	tree.insertMy(sentinelT, 21)
	tree.insertMy(sentinelT, 22)
	tree.insertMy(sentinelT, 23)*/
	print(os.Stdout, tree.root, 0, 'R')

	/*find minimum in the tree*/
	//fmt.Printf("\nMinimum in the tree is:  %+v \n\n", tree.root.findMinimum(sentinelT))

	/*find Maximum in the tree*/
	//fmt.Printf("Maximum in the tree is:  %+v \n\n", tree.root.findMaximum(sentinelT))
	/*search key*/

	var key int64
	key = 32
	tmp := tree.root.searchKey(sentinelT, key)

	if tmp == sentinelT {
		fmt.Printf("the key = %v doesn't find in the tree \n", key)

	} else {
		fmt.Printf("key=%v finded in the tree:  %+v \n", key, tmp)
		tree.rbDelete(sentinelT, tmp)
		fmt.Printf("key=%v deleted:  %+v \n", key, tmp)
		if tree.root == sentinelT {
			fmt.Println("Empty Tree")

		} else {
			print(os.Stdout, tree.root, 0, 'R')
		}
		/*find Successor of the node x in the tree*/
		//	fmt.Printf("Successor of the key(%v) is:  %+v \n", key, z.findSuccessor())
	}

}

func (T *Root) insertMy(sentinelT *Node, data int64) {

	z := &Node{key: data, color: 'B', left: sentinelT, right: sentinelT, father: sentinelT}
	if T.root == nil {
		T.root = sentinelT
	}

	x := T.root
	y := sentinelT

	for x != sentinelT {

		y = x

		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.father = y

	if y == sentinelT {

		T.root = z

	} else if z.key < y.key {
		y.left = z

	} else {
		y.right = z

	}
	z.left = sentinelT
	z.right = sentinelT
	z.color = 'R'

	//x.rbInsertFixup(sentinelT, root)
	T.rbInsertFixup(sentinelT, z)

}

func (T *Root) rbInsertFixup(sentinelT *Node, z *Node) {

	for z.father.color == 'R' {

		if z.father == z.father.father.left {
			y := z.father.father.right
			if y.color == 'R' {

				z.father.color = 'B'
				y.color = 'B'
				z.father.father.color = 'R'
				z = z.father.father

			} else {
				if z == z.father.right {
					z = z.father
					T.leftRotate(sentinelT, z) //?????
				}
				z.father.color = 'B'
				z.father.father.color = 'R'
				T.rightRotate(sentinelT, z.father.father)
			}
		} else {

			y := z.father.father.left
			//y := z.father.father.right
			if y.color == 'R' {

				z.father.color = 'B'
				y.color = 'B'
				z.father.father.color = 'R'
				z = z.father.father

			} else {
				if z == z.father.left {
					z = z.father
					T.rightRotate(sentinelT, z) //
				}
				z.father.color = 'B'
				z.father.father.color = 'R'
				T.leftRotate(sentinelT, z.father.father) //
			}

		}

	}
	T.root.color = 'B'

}

/*transforms the conﬁguration of the two nodes on the right into the conﬁguration on the left by changing
a constant number of pointers. */
func (T *Root) leftRotate(sentinelT *Node, x *Node) {
	y := x.right     // set y
	x.right = y.left //turn y's left subtree into x's right subtree
	if y.left != sentinelT {
		y.left.father = x
	}
	y.father = x.father // link x’s parent to y

	if x.father == sentinelT {
		T.root = y
	} else {
		if x == x.father.left {
			x.father.left = y
		} else {
			x.father.right = y
		}
	}
	y.left = x
	x.father = y

}

/*transforms the conﬁguration of the two nodes on the right into the conﬁguration on the left by changing
a constant number of pointers. */
func (T *Root) rightRotate(sentinelT *Node, x *Node) {
	y := x.left      // set y
	x.left = y.right //turn y's left subtree into x's right subtree
	if y.right != sentinelT {
		y.right.father = x
	}
	y.father = x.father // link x’s parent to y

	if x.father == sentinelT {
		T.root = x //??? y
	} else if x == x.father.left {
		x.father.left = y
	} else {
		x.father.right = y
	}

	y.right = x
	x.father = y

}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v[%v]\n", ch, node.key, string(node.color))
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func (T *Root) rbTransplant(sentinelT, u, v *Node) {
	if u.father == sentinelT {
		T.root = v
	} else if u == u.father.left {
		u.father.left = v
	} else {
		u.father.right = v
	}
	v.father = u.father

}

func (T *Root) rbDelete(sentinelT, z *Node) {

	var x *Node
	y := z
	//parent := z.father
	yOriginalColor := y.color
	if z.left == sentinelT {
		x = z.right
		T.rbTransplant(sentinelT, z, z.right)
	} else if z.right == sentinelT {

		x = z.left
		T.rbTransplant(sentinelT, z, z.left)
	} else {
		y = z.right.findMinimum(sentinelT) //??????

		yOriginalColor = y.color
		x = y.right
		//?? differente dal libro
		//sito https://github.com/sakeven/RbTree/blob/master/rbtree.go
		if y.father == z {
			/*	if x == sentinelT {
					parent=y
				}else{
					x.father = y
				}
			*/

			x.father = y

		} else {
			T.rbTransplant(sentinelT, y, y.right)
			y.right = z.right
			y.right.father = y
		}
		T.rbTransplant(sentinelT, z, y)
		y.left = z.left
		y.left.father = y
		y.color = z.color
	}
	if yOriginalColor == 'B' {
		T.rbDeleteFixup(sentinelT, x)
	}

}

func (T *Root) rbDeleteFixup(sentinelT, x *Node) {
	var w *Node
	for x != T.root && x.color == 'B' {
		if x == x.father.left {
			w = x.father.right
			if w.color == 'R' {
				w.color = 'B'
				x.father.color = 'R'
				T.leftRotate(sentinelT, x.father)
				w = x.father.right
			}
			if w.left.color == 'B' && w.right.color == 'B' {
				w.color = 'R'
				x = x.father
			} else if w.right.color == 'B' {
				w.left.color = 'B'
				w.color = 'R'
				T.rightRotate(sentinelT, w)
				w = x.father.right
			}
			w.color = x.father.color
			x.father.color = 'B'
			w.right.color = 'B'
			T.leftRotate(sentinelT, x.father)
			x = T.root
		} else {

			w = x.father.left
			if w.color == 'R' {
				w.color = 'B'
				x.father.color = 'R'
				T.leftRotate(sentinelT, x.father)
				w = x.father.left
			}
			if w.right.color == 'B' && w.left.color == 'B' {
				w.color = 'R'
				x = x.father
			} else if w.left.color == 'B' {
				w.right.color = 'B'
				w.color = 'R'
				T.rightRotate(sentinelT, w)
				w = x.father.right
			}
			w.color = x.father.color
			x.father.color = 'B'
			w.left.color = 'B'
			T.leftRotate(sentinelT, x.father)
			x = T.root

		}

	}
	x.color = 'B'
}

/*Search a key in the tree*/
func (x *Node) searchKey(sentinelT *Node, key int64) *Node {

	if x == sentinelT || key == x.key {

		return x
	}
	if key < x.key {
		return x.left.searchKey(sentinelT, key)
	}
	return x.right.searchKey(sentinelT, key)

}

/*find the minimun in the tree*/
func (x *Node) findMinimum(sentinelT *Node) *Node {

	if x.left != sentinelT {

		x = x.left.findMinimum(sentinelT)
	}
	return x
}

/*find the Maximum in the tree*/
func (x *Node) findMaximum(sentinelT *Node) *Node {

	if x.right != sentinelT {

		x = x.right.findMaximum(sentinelT)
	}
	return x

}
