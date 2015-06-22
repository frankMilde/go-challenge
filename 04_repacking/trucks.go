//
// =========================================================================
//
//       Filename:  trucks.go
//
//    Description:  Handles all truck related things.
//
//        Version:  1.0
//        Created:  06/10/2015 02:09:41 PM
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
	"fmt"
)

// Unload adds all boxes b from truck pointer tp to Table table and returns
// number of pallets nr.
func (tp *truck) Unload(table Table) int {
	var nr int
	for i, p := range tp.pallets {
		for _, b := range p.boxes {
			table.Add(b.canon())
		}
		nr = i
	}
	return nr + 1
} // -----  end of function Unload -----

// TODO: Add Test.
func TrucksAreEqual(a, b truck) bool {
	if len(a.pallets) != len(b.pallets) {
		return false
	}
	if a.id != b.id {
		return false
	}
	for i, v := range a.pallets {
		if !PalletsAreEqual(v, b.pallets[i]) {
			return false
		}
	}
	return true
} // -----  end of function PalletssAreEqual  -----

func (t truck) String() string {

	s := fmt.Sprintf("Truck %d\n", t.id)
	for i, p := range t.pallets {
		if i < 10 {
			s += fmt.Sprintf("[ %d]  -->  %v\n", i, p)
		} else {
			s += fmt.Sprintf("[%d]  -->  %v\n", i, p)
		}
	}
	return s
}

// ===  FUNCTION  ==========================================================
//         Name:  Unload
//  Description:  Takes a truck pointer as input and returns all boxes that
//                where on this truck.
// =========================================================================
//func (t *truck) Unload() []box {
//	allBoxes := make([]box, 0)
//	for _, p := range t.pallets {
//		allBoxes = append(allBoxes, p.boxes...)
//	}
//	return allBoxes
//} // -----  end of function extractAllBoxesFrom -----

// ===  FUNCTION  ==========================================================
//         Name:  AttachInfoTo
//  Description:  Takes a list of lain boxes and attached useful info to
//                each box. Returns list of boxes with info.
// =========================================================================
//func AttachInfoTo(boxlist []box) []boxWithInfo {
//	allBoxes := make([]boxWithInfo, len(boxlist))
//	for _, box := range boxlist {
//		square := box.IsSquare()
//		newBoxWithInfo := boxWithInfo{box, box.Size(), square, false}
//		allBoxes = append(allBoxes, newBoxWithInfo)
//	}
//	return allBoxes
//} // -----  end of function AttachInfoTo -----
