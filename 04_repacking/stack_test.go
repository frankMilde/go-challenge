//       Filename:  stack_test.go
//    Description:  Testing box stack.
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

import "testing"

func Test_NewStack_LengthIsZero(t *testing.T) {
	s := NewStack()

	got := len(s)
	want := 0

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
func Test_NewStack_FrontIsNil(t *testing.T) {
	got := NewStack()
	if got != nil {
		t.Errorf("Got %v, want nil", got)
	}
}

func Test_IsEmpty_EmptyStack_ReturnTrue(t *testing.T) {
	s := NewStack()

	got := s.IsEmpty()
	want := true

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
func Test_IsEmpty_NonEmptyStack_ReturnFalse(t *testing.T) {
	s := NewStack()
	s.Push(box{0, 0, 1, 1, 100})

	got := s.IsEmpty()
	want := false

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func Test_Front_NonEmptyStack_ReturnLastElement(t *testing.T) {
	s := NewStack()

	b1 := box{0, 0, 1, 1, 100}
	b2 := box{0, 0, 1, 1, 101}

	s.Push(b1)
	s.Push(b2)

	got := s.Front()
	want := b2

	if !BoxesAreEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
func Test_Front_EmptyStack_ReturnEmptyBox(t *testing.T) {
	s := NewStack()

	got := s.Front()
	want := emptybox

	if !BoxesAreEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func Test_Push_AddBoxToEmptyStack(t *testing.T) {

	s := NewStack()
	b := box{0, 0, 1, 1, 100}

	s.Push(b)

	got := s[0]
	want := b

	if !BoxesAreEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
func Test_Push_AddBoxToNonEmptyStack(t *testing.T) {

	s := NewStack()
	b1 := box{0, 0, 1, 1, 100}
	b2 := box{0, 0, 1, 1, 101}

	s.Push(b1)
	s.Push(b2)

	got := s[1]
	want := b2

	if !BoxesAreEqual(got, want) {
		t.Errorf("Boxes: got %v, want %v", got, want)
	}
	if len(s) != 2 {
		t.Errorf("Length: got %d, want %d", len(s), 2)
	}
}

func Test_Pop_NonEmptyStackUntilEmpty_ReturnLastElement(t *testing.T) {
	s := NewStack()
	b1 := box{0, 0, 1, 1, 100}
	b2 := box{0, 0, 1, 1, 101}

	s.Push(b1)
	s.Push(b2)

	type wants struct {
		l int
		b box
	}

	tests := []struct {
		want wants
	}{
		{
			want: wants{1, b2},
		},
		{
			want: wants{0, b1},
		},
		{
			want: wants{0, emptybox},
		},
	} // end tests

	for i, test := range tests {
		gotb := s.Pop()
		gotl := len(s)
		if !BoxesAreEqual(gotb, test.want.b) {
			t.Errorf("Run %d", i)
			t.Errorf("Boxes: Got %v, want %v", gotb, test.want.b)
		}
		if gotl != test.want.l {
			t.Errorf("Run %d", i)
			t.Errorf("Length: Got %v, want %v", gotl, test.want.l)
		}
	}
}

func Test_StacksAreEqual_GetEmptyStacks_ReturnTrue(t *testing.T) {
	got := StacksAreEqual(NewStack(), NewStack())
	want := true
	if got != want {
		t.Errorf("Got %b, want %b", got, want)
	}
} // -----  end of function Test_StacksAreEqual  -----
func Test_StacksAreEqual_GetEqualStacks_ReturnTrue(t *testing.T) {
	s1 := NewStack()
	s2 := NewStack()

	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 2, 102}
	b3 := box{0, 0, 1, 3, 103}
	b4 := box{0, 0, 4, 4, 110}

	boxes := []box{b1, b2, b3, b4}

	for _, box := range boxes {
		s1.Push(box)
		s2.Push(box)
	}
	got := StacksAreEqual(s1, s2)
	want := true
	if got != want {
		t.Errorf("Got %b, want %b", got, want)
	}
} // -----  end of function Test_StacksAreEqual  -----
func Test_StacksAreEqual_GetNonEqualStacks_ReturnFalse(t *testing.T) {
	s1 := NewStack()
	s2 := NewStack()

	b1 := box{0, 0, 1, 1, 101}
	b2 := box{0, 0, 1, 2, 102}
	b3 := box{0, 0, 1, 3, 103}
	b4 := box{0, 0, 4, 4, 110}

	boxes := []box{b1, b2, b3, b4}

	for _, box := range boxes {
		s1.Push(box)
		s2.Push(box)
	}
	s1.Push(b1)

	got := StacksAreEqual(s1, s2)
	want := false
	if got != want {
		t.Errorf("Got %b, want %b", got, want)
	}
} // -----  end of function Test_StacksAreEqual  -----
