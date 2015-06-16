//
// =========================================================================
//
//       Filename:  box_test.go
//
//    Description:  Unit test for the box.go file. Test should have the form
//                  Test_unitOfWork_scenario_expectedBehavior()
//
//           TODO:  Transform all tests to the form above.
//
//        Version:  1.0
//        Created:  06/10/2015 10:54:15 AM
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
	"sort"
	"testing"
)

func Test_Size_normalInput_returnSizeOfBox(t *testing.T) {
	tests := []struct {
		in   *box
		want uint8
	}{
		{
			in:   &box{0, 0, 1, 1, 101},
			want: 1,
		},
		{
			in:   &box{3, 0, 4, 1, 104},
			want: 4,
		},
	}

	for _, test := range tests {
		got := test.in.Size()
		if got != test.want {
			t.Errorf("(%v).Size() == %d, want %d", test.in, got, test.want)
		}
	}
} // -----  end of function Test_Size_normalInput_returnSizeOfBox  -----
func Test_Size_emptyBoxInput_returnZero(t *testing.T) {
	tests := []struct {
		in   *box
		want uint8
	}{
		{
			in:   &box{},
			want: 0,
		},
		{
			in:   &emptybox,
			want: 0,
		},
	}

	for _, test := range tests {
		got := test.in.Size()
		if got != test.want {
			t.Errorf("(%v).Size() == %d, want %d", test.in, got, test.want)
		}
	}
} // -----  end of function Test_Size_emptyBoxInput_returnZero  -----

func Test_HasValidSize_validBoxInput_returnTrue(t *testing.T) {
	tests := []struct {
		in   *box
		want bool
	}{
		{
			in:   &box{0, 0, 1, 1, 101},
			want: true,
		},
		{
			in:   &box{0, 0, 3, 1, 103},
			want: true,
		},
		{
			in:   &box{1, 2, 4, 4, 103},
			want: true,
		},
	}

	for _, test := range tests {
		got := test.in.HasValidSize()
		if got != test.want {
			t.Errorf("(%v).HasValidSize() == %t, want %t", test.in, got, test.want)
		}
	}
} // -----  end of function Test_HasValidSize  -----
func Test_HasValidSize_boxIstoBig_returnFalse(t *testing.T) {
	tests := []struct {
		in   *box
		want bool
	}{
		// to big in first input
		{
			in:   &box{3, 0, 4, 6, 104},
			want: false,
		},
		// to big in second input
		{
			in:   &box{3, 3, 8, 1, 104},
			want: false,
		},
		// to big in both inputs
		{
			in:   &box{3, 3, 5, 5, 104},
			want: false,
		},
	}

	for _, test := range tests {
		got := test.in.HasValidSize()
		if got != test.want {
			t.Errorf("(%v).HasValidSize() == %t, want %t", test.in, got, test.want)
		}
	}
} // -----  end of function Test_HasValidSize  -----
func Test_HasValidSize_boxSizeIsValidButCoordinateAreOutOfBound_returnTrue(t *testing.T) {
	tests := []struct {
		in   *box
		want bool
	}{
		{
			in:   &box{2, 7, 2, 2, 104},
			want: true,
		},
		{
			in:   &box{7, 1, 2, 2, 104},
			want: true,
		},
		{
			in:   &box{7, 7, 2, 2, 104},
			want: true,
		},
	}

	for _, test := range tests {
		got := test.in.HasValidSize()
		if got != test.want {
			t.Errorf("(%v).HasValidSize() == %t, want %t", test.in, got, test.want)
		}
	}
} // -----  end of function Test_HasValidSize  -----
func Test_HasValidSize_emptyBoxInput_returnTrue(t *testing.T) {
	tests := []struct {
		in   *box
		want bool
	}{
		{
			in:   &emptybox,
			want: true,
		},
		{
			in:   &box{},
			want: true,
		},
	}

	for _, test := range tests {
		got := test.in.HasValidSize()
		if got != test.want {
			t.Errorf("(%v).HasValidSize() == %t, want %t", test.in, got, test.want)
		}
	}
} // -----  end of function Test_HasValidSize  -----

func Test_HasValidCoordinates(t *testing.T) {
	tests := []struct {
		in   *box
		want bool
	}{
		// valid box
		{
			in:   &box{0, 0, 1, 1, 101},
			want: true,
		},
		{
			in:   &box{3, 3, 1, 1, 101},
			want: true,
		},
		// too far in x
		{
			in:   &box{4, 0, 4, 6, 104},
			want: false,
		},
		// too far in y
		{
			in:   &box{2, 4, 4, 6, 104},
			want: false,
		},
		// too far in x and y
		{
			in:   &box{4, 4, 4, 6, 104},
			want: false,
		},
		// empty box
		{
			in:   &emptybox,
			want: true,
		},
	}

	for _, test := range tests {
		got := test.in.HasValidCoordinates()
		if got != test.want {
			t.Errorf("(%v).HasValidCoordinates() == %t, want %t", test.in, got, test.want)
		}
	}
} // -----  end of function Test_HasValidCoordinates  -----

type BoxesEqualTest struct {
	a box
	b box
}

func Test_BoxesAreEqual(t *testing.T) {
	tests := []struct {
		in   BoxesEqualTest
		want bool
	}{
		// equal boxes
		{
			in: BoxesEqualTest{
				box{0, 0, 1, 1, 101},
				box{0, 0, 1, 1, 101},
			},
			want: true,
		},
		// different id
		{
			in: BoxesEqualTest{
				box{0, 0, 1, 1, 102},
				box{0, 0, 1, 1, 101},
			},
			want: false,
		},
		// different origin
		{
			in: BoxesEqualTest{
				box{1, 0, 1, 1, 101},
				box{0, 0, 1, 1, 101},
			},
			want: false,
		},
		// emptybox
		{
			in: BoxesEqualTest{
				box{1, 0, 1, 1, 101},
				emptybox,
			},
			want: false,
		},
	}

	for _, test := range tests {
		got := BoxesAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("Comparing boxes: \n %v \n      == \n %v \n want %t, got %t", test.in.a, test.in.b, test.want, got)
		}
	}
} // -----  end of function Test_BoxesAreEqual  -----

type pBoxesEqualTest struct {
	a *box
	b *box
}

func Test_pBoxesAreEqual(t *testing.T) {
	tests := []struct {
		in   pBoxesEqualTest
		want bool
	}{
		// equal boxes
		{
			in: pBoxesEqualTest{
				&box{0, 0, 1, 1, 101},
				&box{0, 0, 1, 1, 101},
			},
			want: true,
		},
		// different id
		{
			in: pBoxesEqualTest{
				&box{0, 0, 1, 1, 102},
				&box{0, 0, 1, 1, 101},
			},
			want: false,
		},
		// different origin
		{
			in: pBoxesEqualTest{
				&box{1, 0, 1, 1, 101},
				&box{0, 0, 1, 1, 101},
			},
			want: false,
		},
		// emptybox
		{
			in: pBoxesEqualTest{
				&box{1, 0, 1, 1, 101},
				&emptybox,
			},
			want: false,
		},
		// invalid pointers
		//	{
		//		in: pBoxesEqualTest{
		//			nil,
		//			nil,
		//		},
		//		want: true,
		//	},
	}

	for _, test := range tests {
		got := BoxesAreEqual(*test.in.a, *test.in.b)
		if got != test.want {
			t.Errorf("Comparing boxes:")
			t.Errorf("%v", test.in.a)
			t.Errorf("    == ")
			t.Errorf("%v", test.in.b)
			t.Errorf("want %t", test.want)
			t.Errorf("got  %t", got)
		}
	}
} // -----  end of function Test_BoxesAreEqual  -----

type PalletsEqualTest struct {
	a pallet
	b pallet
}

func Test_PalletsAreEqual(t *testing.T) {
	tests := []struct {
		in   PalletsEqualTest
		want bool
	}{
		// two equal pallets
		{
			in: PalletsEqualTest{
				pallet{
					[]box{
						box{0, 0, 1, 1, 101},
						box{0, 0, 1, 1, 101},
					},
				},
				pallet{
					[]box{
						box{0, 0, 1, 1, 101},
						box{0, 0, 1, 1, 101},
					},
				},
			},
			want: true,
		},
		// two different pallets
		{
			in: PalletsEqualTest{
				pallet{
					[]box{
						box{0, 0, 1, 1, 101},
						box{0, 0, 1, 1, 101},
					},
				},
				pallet{
					[]box{
						box{0, 0, 1, 1, 101},
						box{1, 0, 1, 1, 101},
					},
				},
			},
			want: false,
		},
		// different number of pallets
		{
			in: PalletsEqualTest{
				pallet{
					[]box{
						box{0, 0, 1, 1, 101},
					},
				},
				pallet{
					[]box{
						box{0, 0, 1, 1, 101},
						box{1, 0, 1, 1, 101},
					},
				},
			},
			want: false,
		},
		// case: two empty pallets
		{
			in: PalletsEqualTest{
				pallet{
					[]box{},
				},
				pallet{
					[]box{},
				},
			},
			want: true,
		},
	}

	for _, test := range tests {
		got := PalletsAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("Comparing pallets \n %v \n            ==\n %v\n want %t, got %t", test.in.a.boxes, test.in.b.boxes, test.want, got)
		}
	}
} // -----  end of function Test_PalletsAreEqual  -----

type BoxArrayEqualTest struct {
	a []box
	b []box
}

func Test_BoxArraysAreEqual(t *testing.T) {
	tests := []struct {
		in   BoxArrayEqualTest
		want bool
	}{
		// two equal pallets
		{
			in: BoxArrayEqualTest{
				[]box{
					box{0, 0, 1, 1, 101},
					box{0, 0, 1, 1, 101},
				},
				[]box{
					box{0, 0, 1, 1, 101},
					box{0, 0, 1, 1, 101},
				},
			},
			want: true,
		},
		// two different pallets
		{
			in: BoxArrayEqualTest{
				[]box{
					box{0, 0, 1, 1, 101},
					box{0, 0, 1, 1, 101},
				},
				[]box{
					box{0, 0, 1, 1, 101},
					box{1, 0, 1, 1, 101},
				},
			},
			want: false,
		},
		// different number of pallets
		{
			in: BoxArrayEqualTest{
				[]box{
					box{0, 0, 1, 1, 101},
				},
				[]box{
					box{0, 0, 1, 1, 101},
					box{1, 0, 1, 1, 101},
				},
			},
			want: false,
		},
		// case: two empty pallets
		{
			in: BoxArrayEqualTest{
				[]box{},
				[]box{},
			},
			want: true,
		},
	}

	for _, test := range tests {
		got := BoxArraysAreEqual(test.in.a, test.in.b)
		if got != test.want {
			t.Errorf("Comparing boxlist \n %v \n            ==\n %v\n want %t, got %t", test.in.a, test.in.b, test.want, got)
		}
	}
} // -----  end of function Test_BoxArraysAreEqual  -----

type PutOnPalletTest struct {
	b box
	p *pallet
}

func Test_PutOnPallet(t *testing.T) {
	tests := []struct {
		in   PutOnPalletTest
		want pallet
	}{
		// empty box on empty pallet
		{
			in: PutOnPalletTest{
				emptybox,
				&pallet{},
			},
			want: pallet{},
		},
		// box on empty pallet
		{
			in: PutOnPalletTest{
				box{0, 0, 1, 1, 100},
				&pallet{},
			},
			want: pallet{
				[]box{
					box{0, 0, 1, 1, 100},
				},
			},
		},
		// box on filled pallet
		{
			in: PutOnPalletTest{
				box{1, 1, 1, 1, 100},
				&pallet{
					[]box{
						box{0, 0, 1, 1, 100},
					},
				},
			},
			want: pallet{
				[]box{
					box{0, 0, 1, 1, 100},
					box{1, 1, 1, 1, 100},
				},
			},
		},
		// box with invalid coordinates on filled pallet
		{
			in: PutOnPalletTest{
				box{4, 5, 1, 1, 100},
				&pallet{
					[]box{
						box{0, 0, 1, 1, 100},
					},
				},
			},
			want: pallet{
				[]box{
					box{0, 0, 1, 1, 100},
				},
			},
		},
	}

	for _, test := range tests {
		test.in.b.PutOnPallet(test.in.p)
		if !PalletsAreEqual(*test.in.p, test.want) {
			t.Errorf("Comparing pallets \n   %v \n!=\n   %v", test.in.p.boxes, test.want.boxes)
		}
	}
} // -----  end of function Test_PutOnPallet  -----

func Test_Sort(t *testing.T) {
	tests := []struct {
		in   []box
		want []box
	}{
		{
			// all boxes at 0,0
			in: []box{
				box{0, 0, 4, 4, 101},
				box{0, 0, 2, 2, 102},
				box{0, 0, 2, 1, 103},
				box{0, 0, 3, 2, 104},
			},
			want: []box{
				box{0, 0, 2, 1, 103},
				box{0, 0, 2, 2, 102},
				box{0, 0, 3, 2, 104},
				box{0, 0, 4, 4, 101},
			},
		},
		{
			// boxes at different coordinates
			in: []box{
				box{0, 0, 4, 4, 101},
				box{1, 1, 2, 2, 102},
				box{3, 1, 2, 1, 103},
				box{0, 2, 3, 2, 104},
			},
			want: []box{
				box{3, 1, 2, 1, 103},
				box{1, 1, 2, 2, 102},
				box{0, 2, 3, 2, 104},
				box{0, 0, 4, 4, 101},
			},
		},
		{
			// two equivalent boxes
			in: []box{
				box{0, 0, 4, 4, 101},
				box{1, 1, 2, 2, 102},
				box{3, 1, 2, 1, 103},
				box{0, 2, 2, 2, 104},
			},
			want: []box{
				box{3, 1, 2, 1, 103},
				box{1, 1, 2, 2, 102},
				box{0, 2, 2, 2, 104},
				box{0, 0, 4, 4, 101},
			},
		},
	}

	for _, test := range tests {
		original := test.in
		sort.Sort(BySize(test.in))

		if !BoxArraysAreEqual(test.in, test.want) {
			t.Errorf("Sorting     %v", original)
			t.Errorf("Resulted in %v", test.in)
			t.Errorf("Should be   %v", test.want)
		}
	}
} // -----  end of function Test_Sort  -----

type SetOriginTest struct {
	b    *box
	x, y uint8
}

func Test_SetOrigin(t *testing.T) {
	tests := []struct {
		in   SetOriginTest
		want *box
	}{
		{
			in: SetOriginTest{
				&box{0, 0, 1, 1, 100},
				1,
				1,
			},
			want: &box{1, 1, 1, 1, 100},
		},
		{
			in: SetOriginTest{
				&box{0, 0, 1, 1, 100},
				0,
				0,
			},
			want: &box{0, 0, 1, 1, 100},
		},
		// coordinates are out of bound
		{
			in: SetOriginTest{
				&box{0, 0, 1, 1, 100},
				4,
				2,
			},
			want: &box{0, 0, 1, 1, 100},
		},
	} // -----  end tests  -----

	for _, test := range tests {
		original := *test.in.b
		test.in.b.SetOrigin(test.in.x, test.in.y)
		if !BoxesAreEqual(*test.want, *test.in.b) {
			t.Errorf("SetOrigin in %v\n Got  %v\n want %v", original, test.in.b, test.want)
		} // -----  end if  -----
	} // -----  end for  -----
} // -----  end of function Test_SetOrigin  -----

func Test_Rotate(t *testing.T) {
	tests := []struct {
		in   *box
		want *box
	}{
		{
			in:   &box{0, 0, 1, 2, 100},
			want: &box{0, 0, 2, 1, 100},
		},
		{
			in:   &box{0, 0, 1, 1, 100},
			want: &box{0, 0, 1, 1, 100},
		},
		// invalid length
		{
			in:   &box{0, 0, 1, 5, 100},
			want: &box{0, 0, 5, 1, 100},
		},
		// invalid coordinates
		{
			in:   &box{0, 5, 1, 3, 100},
			want: &box{0, 5, 3, 1, 100},
		},
	}
	for _, test := range tests {
		original := *test.in
		test.in.Rotate()
		if !BoxesAreEqual(*test.want, *test.in) {
			space := "       "
			t.Errorf("Rotate %v\n %s Got    %v\n %s want   %v", original, space, test.in, space, test.want)
		} // -----  end if  -----
	} // -----  end for  -----
} // -----  end of function Test_Rotate  -----

func Test_IsSquare(t *testing.T) {
	tests := []struct {
		in   *box
		want bool
	}{
		// square box
		{
			in:   &box{0, 0, 1, 1, 101},
			want: true,
		},
		// rectangular box
		{
			in:   &box{3, 0, 4, 1, 104},
			want: false,
		},
		// square box at undefined coordinates
		{
			in:   &box{2, 7, 2, 2, 104},
			want: true,
		},
	}

	for _, test := range tests {
		got := test.in.IsSquare()
		if got != test.want {
			t.Errorf("(%v).IsSquare() == %t, want %t", test.in, got, test.want)
		}
	}
} // -----  end of function Test_IsSquare  -----

// Template
//type FunctionTest struct {
//	a *BoxList
//	b *BoxList
//}
//
//func Test_(t *testing.T) {
//	tests := []struct {
//		in   int
//		want   int
//	}{
//		{
//			in:  3
//			want: 3
//		},
//		{
//			in:  0
//			want: 0
//		},
//	}
//	for _, test := range tests {
//		got := YourFunctionToTest(test.in)
//		if got != test.want {
//			t.Errorf("(%v).Function() == %t, want %t", test.in, got, test.want)
//		}
//	}
//}
