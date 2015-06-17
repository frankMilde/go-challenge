//
// =========================================================================
//
//       Filename:  boxstack_test.go
//
//    Description:  Testing box stack.
//
//        Version:  1.0
//        Created:  06/16/2015 07:08:46 PM
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
	"testing"
)

func Test_Len(t *testing.T) {
	tests := []struct {
		in   *Stack
		want uint
	}{
		{
			in:   &Stack{&Element{}, 0},
			want: 0,
		},
		{
			in:   &Stack{&Element{nil, &box{0, 0, 1, 1, 100}}, 1},
			want: 1,
		},
	} // -----  end of tests  -----

	for _, test := range tests {
		got := test.in.Len()
		if got != test.want {
			t.Errorf("Got %d, want %d", got, test.want)
		}
	} // -----  end of for  -----
} // -----  end of function Test_Len  -----

func Test_NewStack_CreateNewStack_GetEmptyStack(t *testing.T) {

	got := NewStack()
	want := &Stack{}

	if (got.root != want.root) && (got.length != want.length) {
		t.Errorf("Got %v, want %v", &got, &want)
	}
} // -----  end of function Test_NewStack  -----

func Test_Next(t *testing.T) {
	tests := []struct {
		in   *Element
		want *Element
	}{
		{
			in: &Element{
				&Element{nil, &box{0, 0, 1, 1, 101}},
				&box{0, 0, 1, 1, 100},
			},
			want: &Element{nil, &box{0, 0, 1, 1, 101}},
		},
	} // -----  end of tests  -----

	for _, test := range tests {
		if test.in.next.next != test.want.next {
			t.Errorf("Next pointers: got (%v), want (%v)", test.in.next, test.want.next)
		}
		got := test.in.Next()
		if !BoxesAreEqual(*got.b, *test.want.b) {
			t.Errorf("Boxes: got (%v), want (%v)", got.b, test.want.b)
		}
	}

} // -----  end of function Test_Next  -----

func Test_StacksAreEqual_GetEmptyStacks_ReturnTrue(t *testing.T) {
	type inputs struct {
		a, b *Stack
	}
	tests := []struct {
		in   inputs
		want bool
	}{
		{
			in:   inputs{&Stack{}, &Stack{}},
			want: true,
		},
	} // -----  end of tests  -----

	for _, test := range tests {
		got := StacksAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("Got %b, want %b", got, test.want)
		}
	}
} // -----  end of function Test_StacksAreEqual  -----

func Test_StacksAreEqual_GetEqualStacks_ReturnTrue(t *testing.T) {
	type inputs struct {
		a, b *Stack
	}
	tests := []struct {
		in   inputs
		want bool
	}{
		{
			in: inputs{
				&Stack{
					&Element{
						&Element{nil, &box{0, 0, 1, 1, 101}},
						&box{0, 0, 1, 1, 100},
					},
					2,
				},
				&Stack{
					&Element{
						&Element{nil, &box{0, 0, 1, 1, 101}},
						&box{0, 0, 1, 1, 100},
					},
					2,
				},
			}, // -----  end of inputs  -----
			want: true,
		},
	} // -----  end of tests  -----

	for _, test := range tests {
		got := StacksAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("Got %t, want %t", got, test.want)
		}
	}
} // -----  end of function Test_StacksAreEqual  -----
