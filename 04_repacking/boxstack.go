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
	next *Element
	b    *box
}

type Stack struct {
	root   *Element
	length uint
}

func (s *Stack) Len() uint {
	return s.length
}

func NewStack() *Stack {
	return &Stack{&Element{}, 0}
}

func (e *Element) Next() *Element {
	return e.next
}

func StacksAreEqual(a, b *Stack) bool {
	if a.Len() != b.Len() {
		return false
	}

	for p, q := a.root, b.root; p != nil && q != nil; p, q = p.Next(), q.Next() {
		if !BoxesAreEqual(*p.b, *q.b) {
			return false
		}
	}

	return true
}
