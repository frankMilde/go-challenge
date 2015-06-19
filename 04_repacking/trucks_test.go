//
// =========================================================================
//
//       Filename:  trucks_test.go
//
//    Description:  Unit test for the trucks.go file.
//
//        Version:  1.0
//        Created:  06/10/2015 02:10:44 PM
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

func Test_UnloadTo_TruckWithSomePallets(t *testing.T) {

	truck1 := &truck{
		1, []pallet{
			pallet{
				[]box{
					box{0, 0, 1, 1, 101},
				},
			},
			pallet{
				[]box{
					box{1, 1, 1, 1, 102},
					box{2, 2, 1, 1, 103},
				},
			},
			pallet{
				[]box{
					box{3, 0, 4, 1, 104},
					box{0, 0, 1, 1, 105},
					box{0, 0, 1, 1, 106},
					box{0, 0, 4, 3, 107},
				},
			},
		},
	}

	want := Table{
		Stack{},
		Stack{ // 1
			box{0, 0, 1, 1, 101},
			box{1, 1, 1, 1, 102},
			box{2, 2, 1, 1, 103},
			box{0, 0, 1, 1, 105},
			box{0, 0, 1, 1, 106},
		},
		Stack{}, // 2
		Stack{}, // 3
		Stack{ // 4
			box{3, 0, 4, 1, 104},
		},
		Stack{}, // 5
		Stack{}, // 6
		Stack{}, // 7
		Stack{}, // 8
		Stack{}, // 9
		Stack{}, // 10
		Stack{}, // 11
		Stack{ // 12
			box{0, 0, 4, 3, 107},
		},
		Stack{}, // 13
		Stack{}, // 14
		Stack{}, // 15
		Stack{}, // 16
	} // end Stack

	got := NewTable()

	truck1.UnloadTo(got)
	if !TablesAreEqual(got, want) {
		t.Errorf("Comparing Tables:\n")
		t.Errorf("Got: \n%v ", got)
		t.Errorf("Want:\n%v ", want)
	}
}

func Test_UnloadTo_TruckWithEmptyPallets(t *testing.T) {

	truck1 := &truck{
		1, []pallet{
			pallet{
				[]box{
					box{0, 0, 1, 1, 101},
				},
			},
			pallet{
				[]box{
					box{1, 1, 1, 1, 102},
					box{2, 2, 1, 1, 103},
				},
			},
			pallet{
				[]box{},
			},
			pallet{},
		},
	}
	want := Table{
		Stack{},
		Stack{ // 1
			box{0, 0, 1, 1, 101},
			box{1, 1, 1, 1, 102},
			box{2, 2, 1, 1, 103},
		},
		Stack{}, // 2
		Stack{}, // 3
		Stack{}, // 4
		Stack{}, // 5
		Stack{}, // 6
		Stack{}, // 7
		Stack{}, // 8
		Stack{}, // 9
		Stack{}, // 10
		Stack{}, // 11
		Stack{}, // 12
		Stack{}, // 13
		Stack{}, // 14
		Stack{}, // 15
		Stack{}, // 16
	} // end Stack

	got := NewTable()

	truck1.UnloadTo(got)
	if !TablesAreEqual(got, want) {
		t.Errorf("Comparing Tables:\n")
		t.Errorf("Got: \n%v ", got)
		t.Errorf("Want:\n%v ", want)
	}
}
