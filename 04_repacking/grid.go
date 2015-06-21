//
// =========================================================================
//
//       Filename:  grid.go
//
//    Description:  Implements a grid by using slices not lists.
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

import (
	"fmt"
	"sort"
)

type Orientation uint8

const (
	HORIZONTAL Orientation = iota
	VERTICAL
	SQUAREGRID
)

type GridElement struct {
	x, y uint8       //origin
	w, l uint8       //width length
	s    uint8       // size
	o    Orientation //horizontal, vertical, square
}

type FreeGrid []GridElement

var emptygrid = GridElement{}

func NewGrid() FreeGrid {
	var g []GridElement
	return g
}
func NewInitialGrid() FreeGrid {
	init := GridElement{0, 0, 4, 4, 16, SQUAREGRID}
	s := []GridElement{init}
	return s
}
func NewSubGrid(g GridElement) FreeGrid {
	s := []GridElement{g}
	return s
}

func (e *GridElement) SetProperties() {

	e.s = e.w * e.l

	if e.w == e.l {
		e.o = SQUAREGRID
	}

	if e.w > e.l {
		e.o = HORIZONTAL
	}
	if e.w < e.l {
		e.o = VERTICAL
	}
}

func (g FreeGrid) IsEmpty() bool { return len(g) == 0 }

func (o Orientation) String() string {

	var s string

	switch o {
	case HORIZONTAL:
		s = "horizontal"
	case VERTICAL:
		s = "vertical"
	case SQUAREGRID:
		s = "square"
	}

	return s
}
func (e GridElement) String() string {

	var s string
	s += fmt.Sprintf("[%d %d %d %d] ", e.x, e.y, e.w, e.l)
	s += fmt.Sprintf("%d %v ", e.s, e.o)
	return s
}
func (g FreeGrid) String() string {

	var s string
	for i, grid := range g {
		if i < 10 {
			s += fmt.Sprintf("[ %d]  -->  %v\n", i, grid)
		} else {
			s += fmt.Sprintf("[%d]  -->  %v\n", i, grid)
		}
	}
	return s
}

// =========================================================================
//  Implementing Sort interface
//  Will order boxes from lowest to highest size.
//  Use as:
//          boxes = []box
//          sort.Sort(BySize(boxes))
//
//	  			box{0, 0, 4, 4, 101},       box{0, 0, 2, 1, 103},
//	  			box{0, 0, 2, 2, 102},  -->  box{0, 0, 2, 2, 102},
//	  			box{0, 0, 2, 1, 103},       box{0, 0, 3, 2, 104},
//	  			box{0, 0, 3, 2, 104},       box{0, 0, 4, 4, 101},
// =========================================================================
type ByArea []GridElement

func (a ByArea) Len() int           { return len(a) }
func (a ByArea) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByArea) Less(i, j int) bool { return a[i].s < a[j].s }

// -----  end of Sort Interface  -----

// Put takes a box b and puts it in the lower left corner of Gridelement e.
// If b does not cover e completely, the remaining free space of grid e is
// returned. This return value is of type FreeGrid := []GridElement and
// contains up to three elements into which the original e has been
// split by the box: (1) top, (2) right, (3) top right
//  | 1 1 1 3 |
//  | 1 1 1 3 |
//  | b b b 2 |
//  | b b b 2 |
func Put(b box, e GridElement) FreeGrid {

	top := GridElement{
		x: b.x,
		y: b.y + b.l,
		w: b.w,
		l: e.l - b.l,
	}
	right := GridElement{
		x: b.x + b.w,
		y: b.y,
		w: e.w - b.w,
		l: b.l,
	}
	topRight := GridElement{
		x: b.x + b.w,
		y: b.y + b.l,
		w: e.w - b.w,
		l: e.l - b.l,
	}

	top.SetProperties()
	right.SetProperties()
	topRight.SetProperties()

	elements := []GridElement{top, right, topRight}

	var split FreeGrid

	for _, e := range elements {
		if e.s != 0 {
			split = append(split, e)
		}
	}

	sort.Sort(ByArea(split))

	return split
}

// GridElementsAreEqual compares each field of two GridElements a,b and
// return true if they are equal.
func GridElementsAreEqual(a, b GridElement) bool {
	if a.s != b.s {
		return false
	}
	if a.o != b.o {
		return false
	}
	if a.x != b.x {
		return false
	}
	if a.y != b.y {
		return false
	}
	if a.w != b.w {
		return false
	}
	if a.l != b.l {
		return false
	}

	return true
} // -----  end of function GridElementsAreEqual  -----

// FreeGridsAreEqual compares all GridElements of two FreeGrids  a,b and
// return true if they are equal.
func FreeGridsAreEqual(a, b FreeGrid) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !GridElementsAreEqual(v, b[i]) {
			return false
		}
	}
	return true
} // -----  end of function FreeGridssAreEqual  -----

//func (s Stack) IsEmpty() bool { return len(s) == 0 }
//
//// Front returns last element of slice, which is the front of the stack
//func (s Stack) Front() box {
//	if s.IsEmpty() {
//		return emptybox
//	}
//	return s[len(s)-1]
//}
//
//// Push adds box b to stack pointer sp
//func (sp *Stack) Push(b box) {
//	*sp = append(*sp, b)
//}
//
//func (sp *Stack) Pop() box {
//	if (*sp).IsEmpty() {
//		return emptybox
//	}
//
//	last := len(*sp) - 1
//	b := (*sp)[last]
//	//	s[last] = nil // or the zero value of T
//	(*sp) = (*sp)[:last]
//	return b
//}
//
//// StacksAreEqual compares two stacks s1,s2 and returns true if s1 and
//// s2 have the same length and the same boxes at the same positions.
//func StacksAreEqual(s1, s2 Stack) bool {
//	if len(s1) != len(s2) {
//		return false
//	}
//
//	for i, b := range s1 {
//		if !BoxesAreEqual(b, s2[i]) {
//			return false
//		}
//	}
//	return true
//}
//
//// === old ===
//
///*
//// Element is an element of a linked list.
//type Element struct {
//	next  *Element
//	stack *Stack
//	b     *box
//}
//
//// Next returns the next list element or nil.
//func (e *Element) Next() *Element {
//	if p := e.next; e.stack != nil && p != &e.stack.root {
//		return p
//	}
//	return nil
//}
//
//// Box returns the box of element e.
//func (e *Element) Box() *box {
//	if e.b != nil {
//		return e.b
//	}
//	return nil
//}
//
//// Stack represents a singly-linked list.
//// The zero value for Stack is an empty stack ready to use.
//type Stack struct {
//	root   Element
//	length uint
//}
//
//// Init initializes or clears stack s.
//func (s *Stack) Init() *Stack {
//	s.root.next = &s.root
//	s.length = 0
//	return s
//}
//
//// New returns an initialized list.
//func NewStack() *Stack {
//	return new(Stack).Init()
//}
//
//// Len returns the number of elements of stack s.
//// The complexity is O(1).
//func (s *Stack) Len() uint {
//	return s.length
//}
//
//// Front returns the first element of stack s or nil.
//func (s *Stack) Front() *Element {
//	if s.length == 0 {
//		return nil
//	}
//	return s.root.next
//}
//
//// Push inserts Element e created from box b at front of stack s, increments
//// s.length.
//func (s *Stack) Push(b *box) {
//	r := &s.root
//	n := r.next
//
//	e := &Element{n, s, b}
//
//	r.next = e
//	s.length++
//}
//
//// Pop removes first element e of stack s and decrements s.length.
//// Pop returns nil if stack is empty or the *box from said first element.
//func (s *Stack) Pop() *box {
//
//	if s.Len() == 0 {
//		return nil
//	}
//
//	f := s.Front()
//
//	s.root = *f.next
//
//	// avoid memory leaks
//	f.next = nil
//	f.stack = nil
//
//	s.length--
//
//	return f.b
//}
//
//// StacksAreEqual compares to two stacks a,b and returns a true if a and b
//// have the same length and the same boxes at the same positions in the
//// stack.
//func StacksAreEqual(a, b *Stack) bool {
//	if a.Len() != b.Len() {
//		return false
//	}
//
//	if a.Len() == 0 && b.Len() == 0 {
//		return true
//	}
//
//	for p, q := &a.root, &b.root; p != nil && q != nil; p, q = p.Next(), q.Next() {
//		if !BoxesAreEqual(*p.b, *q.b) {
//			return false
//		}
//	}
//
//	return true
//}
//*/
