//
// =========================================================================
//
//       Filename:  boxstack.go
//
//    Description:  Implements a stack by shamelessly using go's lists.
//
//        Version:  1.0
//        Created:  06/16/2015 07:07:49 PM
//       Revision:  none
//       Compiler:  g++
//
//          Usage:  <+USAGE+>
//
//         Output:  <+OUTPUT+>
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

type Element struct {
	next  *Element
	stack *Stack
	b     *box
}

type Stack struct {
	root   Element
	length uint
}

func (e *Element) Next() *Element {
	return e.next
}
func (e *Element) Box() box {
	return *e.b
}

func (s *Stack) Len() uint {
	return s.length
}

func NewStack() *Stack {
	return new(Stack).Init()
}
func (s *Stack) Init() *Stack {
	s.root.next = &s.root
	s.length = 0
	return s
}

func (s *Stack) Push(newBox *box) {
	e := &Element{&s.root, s, newBox}
	s.root = *e
	s.length++
}

func (s *Stack) Pop() *box {
	b := s.root.b
	at := &s.root

	s.root = *at.Next()
	s.length--

	return b
}

// function to help run the tests
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
