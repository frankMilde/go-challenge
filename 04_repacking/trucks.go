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

// ===  FUNCTION  ==========================================================
//         Name:  Unload
//  Description:  Takes a truck pointer as input and returns all boxes that
//                where on this truck.
// =========================================================================
func (t *truck) Unload() []box {
	allBoxes := make([]box, 0)
	for _, p := range t.pallets {
		allBoxes = append(allBoxes, p.boxes...)
	}

	return allBoxes
} // -----  end of function extractAllBoxesFrom -----

// ===  FUNCTION  ==========================================================
//         Name:  AttachInfoTo
//  Description:  Takes a list of lain boxes and attached useful info to
//                each box. Returns list of boxes with info.
// =========================================================================
func AttachInfoTo(boxlist []box) []boxWithInfo {
	allBoxes := make([]boxWithInfo, len(boxlist))
	for _, box := range boxlist {
		square := box.IsSquare()
		newBoxWithInfo := boxWithInfo{box, box.Area(), square, false}
		allBoxes = append(allBoxes, newBoxWithInfo)
	}
	return allBoxes
} // -----  end of function AttachInfoTo -----
