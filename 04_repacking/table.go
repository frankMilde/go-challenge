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
	"errors"
	"fmt"
)

const (
	TABLESIZE = 10
	SQUAREBOX = 5
)

type Table []Stack

type HashError int

// NewTable returns a new Table of size TABLESIZE = 10
func NewTable() Table {
	store := make([]Stack, TABLESIZE)
	// In case we change the stack to work with *box we need to initialize the
	// individual stacks
	//	for i := 0; i != TABLESIZE; i++ {
	//		store[i].Init()
	//	}
	return store
}

// Hash returns the hash [0-9] of box b from its size s=b.Size(). If the box
// has invalid dimensions or the size s is wrong, an error is returned.
func Hash(b *box) (uint8, error) {

	ErrInvalidSize := errors.New("hash: Box has invalid size.")
	var errHash uint8 = 10

	if !b.HasValidDimensions() {
		return errHash, ErrInvalidSize
	}

	var hash uint8
	s := b.Size()

	switch s {
	case 1, 2, 3, 6:
		hash = s - 1
	case 4:
		if b.IsSquare() {
			hash = s
		} else {
			hash = s - 1
		}
	case 8:
		hash = 6
	case 9:
		hash = 7
	case 12:
		hash = 8
	case 16:
		hash = 9
	default:
		return errHash, ErrInvalidSize
	}

	return hash, nil
}

// Add pushes a box b to the appropriate box stack in Table t according to
// its size. An error is returned when input is invalid box.
// TODO: Add Test for error returns
func (t Table) Add(b box) error {

	// this also covers the case of an emptybox
	if !b.HasValidDimensions() {
		return errors.New("Add box to table: Box has invalid size.")
	}

	hash, errHash := Hash(&b)

	if errHash == nil {
		t[hash].Push(b)
		return nil
	}
	return errHash
}

// TablesAreEqual returns true if Table t1 and t2 have the same length and
// their stacks are equal.
func TablesAreEqual(t1, t2 Table) bool {
	if len(t1) != len(t2) {
		return false
	}

	for i, s := range t1 {
		if !StacksAreEqual(s, t2[i]) {
			return false
		}
	}
	return true
}

// String interface to pretty print a Table
func (t Table) String() string {
	var total string
	for i, stack := range t {
		var label string
		switch i {
		case 0, 1, 2, 3, 5:
			label = fmt.Sprintf(" %d", i+1)
		case 4:
			label = fmt.Sprintf("4s")
		case 6:
			label = fmt.Sprintf(" %d", 8)
		case 7:
			label = fmt.Sprintf(" %d", 9)
		case 8:
			label = fmt.Sprintf("%d", 12)
		case 9:
			label = fmt.Sprintf("%d", 16)
		default:
			fmt.Println("default")
		}
		total += fmt.Sprintf("[%s]  -->  %v\n", label, stack)
	}
	return total
}
