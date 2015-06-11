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

func Test_Unload(t *testing.T) {
	var tests = []struct {
		in   *truck
		want []box
	}{
		{
			in: &truck{
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
							box{0, 0, 5, 5, 107},
						},
					},
				},
			},
			want: []box{
				box{0, 0, 1, 1, 101},
				box{1, 1, 1, 1, 102},
				box{2, 2, 1, 1, 103},
				box{3, 0, 4, 1, 104},
				box{0, 0, 1, 1, 105},
				box{0, 0, 1, 1, 106},
				box{0, 0, 5, 5, 107},
			},
		},
		// truck with empty pallets
		{
			in: &truck{
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
			},
			want: []box{
				box{0, 0, 1, 1, 101},
				box{1, 1, 1, 1, 102},
				box{2, 2, 1, 1, 103},
			},
		},
	}

	for _, test := range tests {
		got := test.in.Unload()
		if !BoxlistsAreEqual(got, test.want) {
			t.Errorf("Comparing boxlists \n   %v \n==\n   %v", got, test.want)
		}
	}
}
