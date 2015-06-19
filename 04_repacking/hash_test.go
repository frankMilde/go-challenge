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
//	"testing"
)

//func Test_NewStore_InitialSizeIs17(t *testing.T) {
//
//	store := NewStore()
//
//	if len(store) != MAPSIZE {
//		t.Errorf("New Store length is %v, wanted %v", len(store), MAPSIZE)
//	}
//
//	for i, s := range store {
//		if s.Len() != 0 {
//			t.Errorf("Stack %d in New Store has length %d, want 0 ", i, s.Len())
//		}
//		if s.root.next == nil {
//			t.Errorf("Stack %d in New Store has root %v, want non-nil", i, s.root.next)
//		}
//	}
//
//}

//func Test_Add_SingleBoxPerStack(t *testing.T) {
//
//	s := NewStore()
//
//	b1 := &box{0, 0, 1, 1, 101}
//	b2 := &box{0, 0, 1, 2, 102}
//	b3 := &box{0, 0, 1, 3, 103}
//	b4 := &box{0, 0, 1, 4, 104}
//	b5 := &box{0, 0, 2, 2, 105}
//	b6 := &box{0, 0, 2, 3, 106}
//	b8 := &box{0, 0, 2, 4, 107}
//	b9 := &box{0, 0, 3, 3, 108}
//	b12 := &box{0, 0, 3, 4, 109}
//	b16 := &box{0, 0, 4, 4, 110}
//
//	boxes := []*box{nil, b1, b2, b3, b4, b5, b6, nil, b8, b9, nil, nil, b12, nil, nil, nil, b16}
//
//	for _, box := range boxes {
//		s.Add(box)
//	}
//
//	for i := 1; i < MAPSIZE; i++ {
//		if s[i].Len() != 0 {
//			got := s[i].Pop()
//			if !BoxesAreEqual(*got, *boxes[i]) {
//				t.Errorf("%d Want (%v)", i, boxes[i])
//				t.Errorf("%d Got  (%v) with len %d", i, got, s[i].Len())
//			}
//		}
//	}
//}
