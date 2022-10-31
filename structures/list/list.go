package list

import (
	"errors"
	"fmt"
	"strings"
)

type node[T any] struct {
	Value T
	Next  *node[T]
	Prev  *node[T]
}

type List[T any] struct {
	First  *node[T]
	Last   *node[T]
	length int
}

func (l *List[T]) Iter() <-chan T {
	iter := make(chan T)
	go func() {
		curr := l.First
		for curr != nil {
			iter <- curr.Value
			curr = curr.Next
		}
		close(iter)
	}()
	return iter
}

func (l *List[T]) String() string {
	var builder strings.Builder
	curr := l.First
	builder.WriteString("|")
	for curr.Next != nil {
		builder.WriteString(fmt.Sprintf(" %v <->", curr.Value))
		curr = curr.Next
	}
	builder.WriteString(fmt.Sprintf(" %v ", curr.Value))
	builder.WriteString("|")
	return builder.String()
}

func (l *List[T]) Size() int {
	return l.length
}

func (l *List[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *List[T]) AddFront(elem T) {
	if l.Size() == 0 {
		Node := new(node[T])
		Node.Value = elem
		Node.Prev = nil
		Node.Next = nil
		l.First = Node
		l.Last = Node
	} else {
		Node := new(node[T])
		Node.Value = elem
		Node.Prev = nil
		Node.Next = l.First
		l.First.Prev = Node
		l.First = Node
	}
	l.length++
}

func (l *List[T]) AddBack(elem T) {
	if l.Size() == 0 {
		Node := new(node[T])
		Node.Value = elem
		Node.Prev = nil
		Node.Next = nil
		l.First = Node
		l.Last = Node
	} else {
		Node := new(node[T])
		Node.Value = elem
		Node.Next = nil
		Node.Prev = l.Last
		l.Last.Next = Node
		l.Last = Node
	}
	l.length++
}

func (l *List[T]) At(ind int) (T, error) {
	var zero T
	if l.IsEmpty() {
		return zero, errors.New("list is empty")
	}
	if ind > l.Size() || ind < 0 {
		return zero, errors.New("index is out of bound")
	}
	curr := l.First
	for i := 0; i < ind; i++ {
		curr = curr.Next
	}
	return curr.Value, nil
}

func (l *List[T]) DeleteElem(ind int) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if ind > l.Size() || ind < 0 {
		return errors.New("index is out of bound")
	}
	curr := l.First
	for i := 0; i < ind; i++ {
		curr = curr.Next
	}
	prev := curr.Prev
	next := curr.Next
	prev.Next = next
	next.Prev = prev
	return nil
}
