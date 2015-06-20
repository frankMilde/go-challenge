//
// =========================================================================
//
//       Filename:  stack_test.go
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

func Test_NewGrid(t *testing.T) {
	g := NewGrid()

	if g != nil {
		t.Errorf("Wrong new grid")
	}
}

func Test_NewInitialGrid(t *testing.T) {
	g := NewInitialGrid()

	if !(g[0].x == 0 &&
		g[0].y == 0 &&
		g[0].w == 4 &&
		g[0].l == 4 &&
		g[0].s == 16 &&
		g[0].o == SQUAREGRID) {
		t.Errorf("Wrong initial grid")
	}
}

func Test_NewSubGrid(t *testing.T) {
	init := GridElement{0, 0, 0, 0, 0, 0}
	g := NewSubGrid(init)

	var want uint8 = 0

	if !(g[0].x == want &&
		g[0].y == want &&
		g[0].w == want &&
		g[0].l == want &&
		g[0].s == want &&
		g[0].o == HORIZONTAL) {
		t.Errorf("Non zero")
	}

	init2 := GridElement{1, 2, 2, 2, 4, SQUAREGRID}
	g2 := NewSubGrid(init2)

	if !(g2[0].x == 1 &&
		g2[0].y == 2 &&
		g2[0].w == 2 &&
		g2[0].l == 2 &&
		g2[0].s == 4 &&
		g2[0].o == SQUAREGRID) {
		t.Errorf("Non zero")
	}

}

func Test_SetProperties(t *testing.T) {

	e := GridElement{1, 2, 2, 2, 0, 0}

	e.SetProperties()

	if !(e.s == 4 && e.o == SQUAREGRID) {
		t.Errorf("Settings Wrong")
	}
}

func Test_IsEmpty(t *testing.T) {
	g := NewGrid()

	if !g.IsEmpty() {
		t.Errorf("Grid not empty")
	}
}

func Test_Split(t *testing.T) {
	e := GridElement{0, 0, 4, 4, 16, SQUAREGRID}

	b := box{0, 0, 3, 2, 100}

	g := e.Split(b)

	want := FreeGrid{
		GridElement{3, 0, 1, 2, 2, VERTICAL},
		GridElement{3, 2, 1, 2, 2, VERTICAL},
		GridElement{0, 2, 3, 2, 6, HORIZONTAL},
	}

	if !FreeGridsAreEqual(g, want) {
		t.Errorf("Spliting wrong")
		t.Errorf("got:  \n%v", g)
		t.Errorf("want: \n%v", want)
	}
}

func Test_GridElementsAreEqual_InputAreGridElements(t *testing.T) {
	type inputs struct {
		a GridElement
		b GridElement
	}

	tests := []struct {
		in   inputs
		want bool
	}{
		// two emptybox
		{
			in: inputs{
				emptygrid,
				emptygrid,
			},
			want: true,
		},
		// equal grids
		{
			in: inputs{
				GridElement{0, 0, 2, 2, 4, SQUAREGRID},
				GridElement{0, 0, 2, 2, 4, SQUAREGRID},
			},
			want: true,
		},
		// different id
		{
			in: inputs{
				GridElement{0, 0, 1, 4, 4, VERTICAL},
				GridElement{0, 0, 2, 2, 4, SQUAREGRID},
			},
			want: false,
		},
		// different origin
		{
			in: inputs{
				GridElement{0, 0, 2, 2, 4, SQUAREGRID},
				GridElement{1, 2, 2, 2, 4, SQUAREGRID},
			},
			want: false,
		},
		// one emptybox
		{
			in: inputs{
				GridElement{1, 2, 2, 2, 4, SQUAREGRID},
				emptygrid,
			},
			want: false,
		},
	}

	for _, test := range tests {
		got := GridElementsAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("Comparing GridElements: \n %v \n      == \n %v \n want %t, got %t", test.in.a, test.in.b, test.want, got)
		}
	}
} // -----  end of function Test_BoxesAreEqual_InputAreBoxes  -----
func Test_FreeGridsAreEqual(t *testing.T) {
	type inputs struct {
		a FreeGrid
		b FreeGrid
	}
	tests := []struct {
		in   inputs
		want bool
	}{
		// two equal FreeGrids
		{
			in: inputs{
				FreeGrid{
					GridElement{0, 0, 2, 2, 4, SQUAREGRID},
					GridElement{1, 2, 2, 2, 4, SQUAREGRID},
				},
				FreeGrid{
					GridElement{0, 0, 2, 2, 4, SQUAREGRID},
					GridElement{1, 2, 2, 2, 4, SQUAREGRID},
				},
			},
			want: true,
		},
		// two different FreeGrids
		{
			in: inputs{
				FreeGrid{
					GridElement{0, 0, 2, 2, 4, SQUAREGRID},
					GridElement{1, 2, 3, 2, 6, HORIZONTAL},
				},
				FreeGrid{
					GridElement{0, 0, 2, 2, 4, SQUAREGRID},
					GridElement{1, 2, 2, 2, 4, SQUAREGRID},
				},
			},
			want: false,
		},
		// different number of FreeGrids
		{
			in: inputs{
				FreeGrid{
					GridElement{0, 0, 2, 2, 4, SQUAREGRID},
					GridElement{1, 2, 3, 2, 6, HORIZONTAL},
				},
				FreeGrid{
					GridElement{0, 0, 2, 2, 4, SQUAREGRID},
				},
			},
			want: false,
		},
		// case: two empty FreeGrids
		{
			in: inputs{
				FreeGrid{
					GridElement{},
				},
				FreeGrid{
					GridElement{},
				},
			},
			want: true,
		},
	}

	for _, test := range tests {
		got := FreeGridsAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("got: %t, want: %t", got, test.want)
			t.Errorf("a: \n%v", test.in.a)
			t.Errorf("b: \n%v", test.in.b)
		}
	}
} // -----  end of function Test_PalletsAreEqual  -----

//func Test_NewStack_LengthIsZero(t *testing.T) {
//	s := NewStack()
//
//	got := len(s)
//	want := 0
//
//	if got != want {
//		t.Errorf("Got %d, want %d", got, want)
//	}
//}
//func Test_NewStack_FrontIsNil(t *testing.T) {
//	got := NewStack()
//	if got != nil {
//		t.Errorf("Got %v, want nil", got)
//	}
//}
//
//func Test_IsEmpty_EmptyStack_ReturnTrue(t *testing.T) {
//	s := NewStack()
//
//	got := s.IsEmpty()
//	want := true
//
//	if got != want {
//		t.Errorf("Got %d, want %d", got, want)
//	}
//}
//func Test_IsEmpty_NonEmptyStack_ReturnFalse(t *testing.T) {
//	s := NewStack()
//	s.Push(box{0, 0, 1, 1, 100})
//
//	got := s.IsEmpty()
//	want := false
//
//	if got != want {
//		t.Errorf("Got %d, want %d", got, want)
//	}
//}
//
//func Test_Front_NonEmptyStack_ReturnLastElement(t *testing.T) {
//	s := NewStack()
//
//	b1 := box{0, 0, 1, 1, 100}
//	b2 := box{0, 0, 1, 1, 101}
//
//	s.Push(b1)
//	s.Push(b2)
//
//	got := s.Front()
//	want := b2
//
//	if !BoxesAreEqual(got, want) {
//		t.Errorf("Got %v, want %v", got, want)
//	}
//}
//func Test_Front_EmptyStack_ReturnEmptyBox(t *testing.T) {
//	s := NewStack()
//
//	got := s.Front()
//	want := emptybox
//
//	if !BoxesAreEqual(got, want) {
//		t.Errorf("Got %v, want %v", got, want)
//	}
//}
//
//func Test_Push_AddBoxToEmptyStack(t *testing.T) {
//
//	s := NewStack()
//	b := box{0, 0, 1, 1, 100}
//
//	s.Push(b)
//
//	got := s[0]
//	want := b
//
//	if !BoxesAreEqual(got, want) {
//		t.Errorf("Got %v, want %v", got, want)
//	}
//}
//func Test_Push_AddBoxToNonEmptyStack(t *testing.T) {
//
//	s := NewStack()
//	b1 := box{0, 0, 1, 1, 100}
//	b2 := box{0, 0, 1, 1, 101}
//
//	s.Push(b1)
//	s.Push(b2)
//
//	got := s[1]
//	want := b2
//
//	if !BoxesAreEqual(got, want) {
//		t.Errorf("Boxes: got %v, want %v", got, want)
//	}
//	if len(s) != 2 {
//		t.Errorf("Length: got %d, want %d", len(s), 2)
//	}
//}
//
//func Test_Pop_NonEmptyStackUntilEmpty_ReturnLastElement(t *testing.T) {
//	s := NewStack()
//	b1 := box{0, 0, 1, 1, 100}
//	b2 := box{0, 0, 1, 1, 101}
//
//	s.Push(b1)
//	s.Push(b2)
//
//	type wants struct {
//		l int
//		b box
//	}
//
//	tests := []struct {
//		want wants
//	}{
//		{
//			want: wants{1, b2},
//		},
//		{
//			want: wants{0, b1},
//		},
//		{
//			want: wants{0, emptybox},
//		},
//	} // end tests
//
//	for i, test := range tests {
//		gotb := s.Pop()
//		gotl := len(s)
//		if !BoxesAreEqual(gotb, test.want.b) {
//			t.Errorf("Run %d", i)
//			t.Errorf("Boxes: Got %v, want %v", gotb, test.want.b)
//		}
//		if gotl != test.want.l {
//			t.Errorf("Run %d", i)
//			t.Errorf("Length: Got %v, want %v", gotl, test.want.l)
//		}
//	}
//}
//
//func Test_StacksAreEqual_GetEmptyStacks_ReturnTrue(t *testing.T) {
//	got := StacksAreEqual(NewStack(), NewStack())
//	want := true
//	if got != want {
//		t.Errorf("Got %b, want %b", got, want)
//	}
//} // -----  end of function Test_StacksAreEqual  -----
//func Test_StacksAreEqual_GetEqualStacks_ReturnTrue(t *testing.T) {
//	s1 := NewStack()
//	s2 := NewStack()
//
//	b1 := box{0, 0, 1, 1, 101}
//	b2 := box{0, 0, 1, 2, 102}
//	b3 := box{0, 0, 1, 3, 103}
//	b4 := box{0, 0, 4, 4, 110}
//
//	boxes := []box{b1, b2, b3, b4}
//
//	for _, box := range boxes {
//		s1.Push(box)
//		s2.Push(box)
//	}
//	got := StacksAreEqual(s1, s2)
//	want := true
//	if got != want {
//		t.Errorf("Got %b, want %b", got, want)
//	}
//} // -----  end of function Test_StacksAreEqual  -----
//func Test_StacksAreEqual_GetNonEqualStacks_ReturnFalse(t *testing.T) {
//	s1 := NewStack()
//	s2 := NewStack()
//
//	b1 := box{0, 0, 1, 1, 101}
//	b2 := box{0, 0, 1, 2, 102}
//	b3 := box{0, 0, 1, 3, 103}
//	b4 := box{0, 0, 4, 4, 110}
//
//	boxes := []box{b1, b2, b3, b4}
//
//	for _, box := range boxes {
//		s1.Push(box)
//		s2.Push(box)
//	}
//	s1.Push(b1)
//
//	got := StacksAreEqual(s1, s2)
//	want := false
//	if got != want {
//		t.Errorf("Got %b, want %b", got, want)
//	}
//} // -----  end of function Test_StacksAreEqual  -----
