package main

import (
	"fmt"
	"strings"
)

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type Node struct {
	key    int
	color  Color
	parent *Node
	left   *Node
	right  *Node
}

type RedBlackTree struct {
	root *Node
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (t *RedBlackTree) rotateLeft(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RedBlackTree) rotateRight(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *RedBlackTree) insertFixup(z *Node) {
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.rotateLeft(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rotateRight(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rotateLeft(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RedBlackTree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (t *RedBlackTree) minimum(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func (t *RedBlackTree) deleteFixup(x *Node) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					t.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					t.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}
	x.color = BLACK
}

func (t *RedBlackTree) Insert(key int) {
	z := &Node{key: key, color: RED}
	var y *Node
	x := t.root
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	t.insertFixup(z)
}

func (t *RedBlackTree) Remove(key int) {
	z := t.root
	for z != nil {
		if key == z.key {
			break
		} else if key < z.key {
			z = z.left
		} else {
			z = z.right
		}
	}
	if z == nil {
		return
	}
	y := z
	yOriginalColor := y.color
	var x *Node
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimum(z.right)
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			if x != nil {
				x.parent = y
			}
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yOriginalColor == BLACK && x != nil {
		t.deleteFixup(x)
	}
}

func (t *RedBlackTree) printInorder(x *Node) {
	if x != nil {
		t.printInorder(x.left)
		fmt.Printf("%d ", x.key)
		t.printInorder(x.right)
	}
}

func (t *RedBlackTree) PrintInorder() {
	t.printInorder(t.root)
	fmt.Println()
}

func (t *RedBlackTree) getRoot() *Node {
	return t.root
}

func printTree(root *Node, space int, count int) {
	if root == nil {
		return
	}
	space += count
	printTree(root.right, space, count)
	fmt.Println()
	fmt.Print(strings.Repeat(" ", space))
	color := "Ч"
	if root.color == RED {
		color = "К"
	}
	fmt.Printf("%d(%s)\n", root.key, color)
	printTree(root.left, space, count)
}

func main() {
	rbTree := NewRedBlackTree()

	for {
		fmt.Println("1. Вставить\n2. Удалить\n3. Вывести дерево\n4. Выйти\nВведите ваш выбор: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Введите значение для вставки: ")
			var value int
			fmt.Scan(&value)
			rbTree.Insert(value)
		case 2:
			fmt.Println("Введите значение для удаления: ")
			var value int
			fmt.Scan(&value)
			rbTree.Remove(value)
		case 3:
			fmt.Println("Красно-черное дерево:")
			printTree(rbTree.getRoot(), 0, 10)
			fmt.Println()
		case 4:
			return
		default:
			fmt.Println("Неверный выбор! Попробуйте снова.")
		}
	}
}
