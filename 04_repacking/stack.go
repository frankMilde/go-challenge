//
// =========================================================================
//
//       Filename:  stack.go
//
//    Description:  Implements a stack by using slices not lists.
//
//        Version:  1.0
//        Created:  06/16/2015 07:07:49 PM
//       Revision:  none
//       Compiler:  go
//
//          Usage:  <+USAGE+>
//
//         Output:  <+OUTPUT+>
//
//           TODO:  Try to get it thread safe. Resources:
//      						https://github.com/hishboy/gocommons/blob/master/lang/stack.go
//                  https://gist.github.com/moräs/2141121
//
//         Author:  Frank Milde (FM), frank.milde (at) posteo.de
//        Company:
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde
//
// =========================================================================
//

package main

type Stack []box

func NewStack() Stack {
	var s []box
	return s
}

func (s Stack) IsEmpty() bool { return len(s) == 0 }

// Front returns last element of slice, which is the front of the stack
func (s Stack) Front() box {
	if s.IsEmpty() {
		return emptybox
	}
	return s[len(s)-1]
}

// Push adds box b to stack pointer sp
func (sp *Stack) Push(b box) {
	*sp = append(*sp, b)
}

func (sp *Stack) Pop() box {
	if (*sp).IsEmpty() {
		return emptybox
	}

	last := len(*sp) - 1
	b := (*sp)[last]
	//	s[last] = nil // or the zero value of T
	(*sp) = (*sp)[:last]
	return b
}

// StacksAreEqual compares two stacks s1,s2 and returns true if s1 and
// s2 have the same length and the same boxes at the same positions.
func StacksAreEqual(s1, s2 Stack) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, b := range s1 {
		if !BoxesAreEqual(b, s2[i]) {
			return false
		}
	}
	return true
}

// === old ===

/*
// Element is an element of a linked list.
type Element struct {
	next  *Element
	stack *Stack
	b     *box
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.stack != nil && p != &e.stack.root {
		return p
	}
	return nil
}

// Box returns the box of element e.
func (e *Element) Box() *box {
	if e.b != nil {
		return e.b
	}
	return nil
}

// Stack represents a singly-linked list.
// The zero value for Stack is an empty stack ready to use.
type Stack struct {
	root   Element
	length uint
}

// Init initializes or clears stack s.
func (s *Stack) Init() *Stack {
	s.root.next = &s.root
	s.length = 0
	return s
}

// New returns an initialized list.
func NewStack() *Stack {
	return new(Stack).Init()
}

// Len returns the number of elements of stack s.
// The complexity is O(1).
func (s *Stack) Len() uint {
	return s.length
}

// Front returns the first element of stack s or nil.
func (s *Stack) Front() *Element {
	if s.length == 0 {
		return nil
	}
	return s.root.next
}

// Push inserts Element e created from box b at front of stack s, increments
// s.length.
func (s *Stack) Push(b *box) {
	r := &s.root
	n := r.next

	e := &Element{n, s, b}

	r.next = e
	s.length++
}

// Pop removes first element e of stack s and decrements s.length.
// Pop returns nil if stack is empty or the *box from said first element.
func (s *Stack) Pop() *box {

	if s.Len() == 0 {
		return nil
	}

	f := s.Front()

	s.root = *f.next

	// avoid memory leaks
	f.next = nil
	f.stack = nil

	s.length--

	return f.b
}

// StacksAreEqual compares to two stacks a,b and returns a true if a and b
// have the same length and the same boxes at the same positions in the
// stack.
func StacksAreEqual(a, b *Stack) bool {
	if a.Len() != b.Len() {
		return false
	}

	if a.Len() == 0 && b.Len() == 0 {
		return true
	}

	for p, q := &a.root, &b.root; p != nil && q != nil; p, q = p.Next(), q.Next() {
		if !BoxesAreEqual(*p.b, *q.b) {
			return false
		}
	}

	return true
}
*/
