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

func NewTable() Table {
	store := make([]Stack, TABLESIZE)
	// In case we change the stack to work with *box we need to initialize the
	// individual stacks
	//	for i := 0; i != TABLESIZE; i++ {
	//		store[i].Init()
	//	}
	return store
}

func (t Table) Add(b box) error {
	errAdd := errors.New("Add box to table: Invalid box.")

	if b != emptybox {

		hash, errHash := Hash(&b)

		if errHash == nil {
			t[hash].Push(b)
			return nil
		}
		return errHash
	} // -----  end if b -----
	return errAdd
}

func Hash(b *box) (uint8, error) {

	err := errors.New("hash: Box has invalid size.")
	var errHash uint8 = 10

	if !b.HasValidDimensions() {
		return errHash, err
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
		return errHash, err
	}

	return hash, nil
}

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
