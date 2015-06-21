//
// =========================================================================
//
//       Filename:  hash_test.go
//
//    Description:
//
//        Version:  1.0
//        Created:  06/18/2015 05:46:10 PM
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
	//	"fmt"
	"errors"
	"testing"
)

func Test_NewTable_InitialSizeIs17(t *testing.T) {

	store := NewTable()

	if len(store) != TABLESIZE {
		t.Errorf("New Table length is %v, wanted %v", len(store), TABLESIZE)
	}

	for i, s := range store {
		if len(s) != 0 {
			t.Errorf("Stack %d in New Table has length %d, want 0 ", i, len(s))
		}
		if s.Front() != emptybox {
			t.Errorf("Stack %d in New Table has Front %v, want emptybox %v", i, s.Front(), emptybox)
		}
	}

}

func Test_Add_SingleBoxPerStack(t *testing.T) {

	s := NewTable()

	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 2, 102}
	b3 := box{0, 0, 1, 3, 103}
	b4 := box{0, 0, 1, 4, 104}
	b5 := box{0, 0, 2, 2, 105}
	b6 := box{0, 0, 2, 3, 106}
	b8 := box{0, 0, 2, 4, 107}
	b9 := box{0, 0, 3, 3, 108}
	b12 := box{0, 0, 3, 4, 109}
	b16 := box{0, 0, 4, 4, 110}

	boxes := []box{b1, b2, b3, b4, b5, b6, b8, b9, b12, b16}

	for _, box := range boxes {
		s.Add(box)
	}

	for i, stack := range s {
		got := stack.Pop()
		if !BoxesAreEqual(got, boxes[i]) {
			t.Errorf("%d Want (%v)", i, boxes[i])
			t.Errorf("%d Got  (%v) with len %d", i, got, len(stack))
		}
	}
}
func Test_Add_MultipleBoxesPerStack(t *testing.T) {

	s := NewTable()

	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 1, 102}
	b3 := box{0, 0, 1, 1, 103}
	b4 := box{0, 0, 1, 1, 104}
	b5 := box{0, 0, 2, 4, 105}
	b6 := box{0, 0, 2, 4, 106}
	b8 := box{0, 0, 2, 4, 107}
	b9 := box{0, 0, 3, 3, 108}
	b12 := box{0, 0, 3, 3, 109}

	boxes := []box{b1, b2, b3, b4, b5, b6, b8, b9, b12}

	for _, box := range boxes {
		s.Add(box)
	}

	// all 1x1 boxes
	var s1 Stack = Stack{b1, b2, b3, b4}
	// all 2x4 boxes
	var s8 Stack = Stack{b5, b6, b8}
	// all 3x4 boxes
	var s9 Stack = Stack{b9, b12}

	if !StacksAreEqual(s[0], s1) {
		t.Errorf("%d Want (%v)", s1)
		t.Errorf("%d Got  (%v)", s[0])
	}
	if !StacksAreEqual(s[6], s8) {
		t.Errorf("%d Want (%v)", s8)
		t.Errorf("%d Got  (%v)", s[6])
	}
	if !StacksAreEqual(s[7], s9) {
		t.Errorf("%d Want (%v)", s9)
		t.Errorf("%d Got  (%v)", s[7])
	}
}
func Test_Add_InvalidInput_CheckErrMsg(t *testing.T) {

	s := NewTable()

	zeroWidth := box{0, 0, 0, 1, 101}
	tooLong := box{0, 0, 2, 5, 102}

	invSizeErr := errors.New("Add box to table: Box has invalid size.")

	boxes := []box{emptybox, zeroWidth, tooLong}
	want := invSizeErr

	errMsgs := make([]error, 3)

	// collect error messages
	for i, box := range boxes {
		errMsgs[i] = s.Add(box)
	}

	for _, got := range errMsgs {
		if got.Error() != want.Error() {
			t.Errorf("Got (%v), want (%v)", got, want)
		}
	}
}

func Test_Hash_InputRegularBoxes(t *testing.T) {
	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 2, 102}
	b3 := box{0, 0, 1, 3, 103}
	b4 := box{0, 0, 1, 4, 104}
	b5 := box{0, 0, 2, 2, 105}
	b6 := box{0, 0, 2, 3, 106}
	b8 := box{0, 0, 2, 4, 107}
	b9 := box{0, 0, 3, 3, 108}
	b12 := box{0, 0, 3, 4, 109}
	b16 := box{0, 0, 4, 4, 110}

	boxes := []box{b1, b2, b3, b4, b5, b6, b8, b9, b12, b16}
	wants := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, box := range boxes {
		got, _ := Hash(&box)
		want := wants[i]
		if want != got {
			t.Errorf("Got %d, want %d", got, want)
		}
	}
}
func Test_Hash_InputInvalidBoxes_CheckErrMsg(t *testing.T) {
	zeroWidth := box{0, 0, 0, 1, 101}
	tooLong := box{0, 0, 2, 5, 102}

	boxes := []box{emptybox, zeroWidth, tooLong}
	want := errors.New("hash: Box has invalid size.")

	for _, box := range boxes {
		_, got := Hash(&box)
		if want.Error() != got.Error() {
			t.Errorf("Got (%v), want (%v)", got, want)
		}
	}
}

func Test_TablesAreEqual(t *testing.T) {
	t1 := NewTable()
	t2 := NewTable()

	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 1, 102}
	b3 := box{0, 0, 1, 1, 103}
	b4 := box{0, 0, 1, 1, 104}
	b5 := box{0, 0, 2, 4, 105}
	b6 := box{0, 0, 2, 4, 106}
	b8 := box{0, 0, 2, 4, 107}
	b9 := box{0, 0, 3, 3, 108}
	b12 := box{0, 0, 3, 3, 109}

	boxes := []box{b1, b2, b3, b4, b5, b6, b8, b9, b12}

	for _, box := range boxes {
		t1.Add(box)
		t2.Add(box)
	}
	t1.Add(b1)

	got := TablesAreEqual(t1, t2)
	want := false

	if got != want {
		t.Errorf("%v ", t1)
		t.Errorf("%v ", t2)
	}

}

func Test_String(t *testing.T) {
	s := NewTable()

	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 2, 102}
	b3 := box{0, 0, 1, 3, 103}
	b4 := box{0, 0, 1, 4, 104}
	b5 := box{0, 0, 2, 2, 105}
	b6 := box{0, 0, 2, 3, 106}
	b8 := box{0, 0, 2, 4, 107}
	b9 := box{0, 0, 3, 3, 108}
	b12 := box{0, 0, 3, 4, 109}
	b16 := box{0, 0, 4, 4, 110}

	boxes := []box{b1, b2, b3, b4, b5, b6, b8, b9, b12, b16}

	for _, box := range boxes {
		s.Add(box)
	}

	got := s.String()

	want := `[ 1]  -->  [0 0 1 1 101]
[ 2]  -->  [0 0 1 2 102]
[ 3]  -->  [0 0 1 3 103]
[ 4]  -->  [0 0 1 4 104]
[4s]  -->  [0 0 2 2 105]
[ 6]  -->  [0 0 2 3 106]
[ 8]  -->  [0 0 2 4 107]
[ 9]  -->  [0 0 3 3 108]
[12]  -->  [0 0 3 4 109]
[16]  -->  [0 0 4 4 110]
`
	if got != want {
		t.Errorf("String is wrong")
		t.Errorf("got:\n%s", got)
		t.Errorf("want\n%s", want)
	}

}
