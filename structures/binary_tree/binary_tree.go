package binary_tree

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	Left   *node[T]
	Right  *node[T]
	Parent *node[T]
	Value  T
}

func NewNode[T any](elem T) *node[T] {
	return &node[T]{
		Value:  elem,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

func nodePrint[T any](prefix string, n *node[T], isLeft bool, sb *strings.Builder) {
	if n != nil {
		sb.WriteString(prefix + func() string {
			if isLeft {
				return fmt.Sprintf("|--%v\n", n.Value)
			}
			return fmt.Sprintf("\\--%v\n", n.Value)
		}())
		nodePrint(prefix+func() string {
			if isLeft {
				return "|   "
			}
			return "    "
		}(), n.Left, true, sb)
		nodePrint(prefix+func() string {
			if isLeft {
				return "|   "
			}
			return "    "
		}(), n.Right, false, sb)
	}
}

type BinaryTree[T any] struct {
	Root *node[T]
}

func (b *BinaryTree[T]) String() string {
	var sb strings.Builder
	nodePrint("", b.Root, false, &sb)
	return sb.String()
}

func (b *BinaryTree[T]) IsEmpty() bool {
	return b.Root == nil
}

func inOrder[T any](node *node[T], iter chan T) {
	if node == nil {
		return
	}
	inOrder(node.Left, iter)
	iter <- node.Value
	inOrder(node.Right, iter)
}

func (b *BinaryTree[T]) Iter() <-chan T {
	iter := make(chan T)
	go func() {
		inOrder(b.Root, iter)
		close(iter)
	}()
	return iter
}

func (b *BinaryTree[T]) Insert(elem T, less func(a, b T) bool) {
	if b.Root == nil {
		b.Root = NewNode(elem)
	} else {
		current := b.Root
		parent := current
		for current != nil {
			parent = current
			if less(elem, current.Value) {
				current = current.Left
			} else {
				current = current.Right
			}
		}
		current = NewNode(elem)
		current.Parent = parent
		if less(current.Value, parent.Value) {
			parent.Left = current
		} else {
			parent.Right = current
		}
	}
}
