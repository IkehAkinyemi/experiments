package main

import (
	"fmt"
	"strings"
)

type LL struct {
	Item int
	Next *LL
}

func ReorderStudents(ll *LL) {
	if ll == nil || ll.Next == nil {
		return
	}

	var h *LL
	s, f := ll, ll
	for {
		if f == nil {
			break
		}
		h = s
		if f.Next == nil {
			s, f = s.Next, f.Next
			break
		}
		s, f = s.Next, f.Next.Next
	}

	var prev *LL
	curr := s

	for {
		if curr == nil {
			h.Next = prev
			return
		}

		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

func BuildLinkedList(items ...int) *LL {
	if len(items) < 1 {
		return &LL{}
	}

	ll := &LL{items[0], nil}
	lastNode := ll
	for _, item := range items[1:] {
		newNode := &LL{item, nil}
		lastNode.Next = newNode
		lastNode = newNode
	}

	return ll
}

func (ll *LL) String() string {
	if ll == nil {
		return fmt.Sprintln("nil")
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d --> ", ll.Item))

	i := ll.Next
	for {
		if i == nil {
			break
		}
		sb.WriteString(fmt.Sprintf("%d --> ", i.Item))
		i = i.Next
	}

	sb.WriteString("nil")
	return sb.String()
}

func main() {
	ll := BuildLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("%s\n", ll)

	ReorderStudents(ll)

	fmt.Printf("%s\n", ll)
}
