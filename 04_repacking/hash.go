//
// =========================================================================
//
//       Filename:  hash.go
//
//    Description:  Implements the hash table to store the box stacks in.
//
//        Version:  1.0
//        Created:  06/18/2015 05:45:39 PM
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

import (
	"fmt"
)

const (
	TABLESIZE = 17
	SQUAREBOX = 5
)

type Table []Stack

func NewTable() Table {
	store := make([]Stack, TABLESIZE)
	// In case we change the stack to work with *box we need to initialize the
	// individual stacks
	//	for i := 0; i != TABLESIZE; i++ {
	//		store[i].Init()
	//	}
	return store
}

func (t Table) Add(b box) {
	if b != emptybox {
		size := b.Size()
		if size == 4 && b.IsSquare() {
			t[SQUAREBOX].Push(b)
		} else {
			t[size].Push(b)
		}
	}
}

func (t Table) String() string {
	var total string
	for i, stack := range t {
		if i < 10 {
			total += fmt.Sprintf("[ %d]  -->  %v\n", i, stack)
		} else {
			total += fmt.Sprintf("[%d]  -->  %v\n", i, stack)
		}
	}
	return total
}
