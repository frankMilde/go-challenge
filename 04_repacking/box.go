//
// =========================================================================
//
//       Filename:  box.go
//
//    Description:  Handles all things related to boxes.
//
//        Version:  1.0
//        Created:  06/10/2015 10:47:42 AM
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

type boxWithInfo struct {
	b          box
	size       uint8
	isSquare   bool
	isOnPallet bool
}

type boxError struct {
	//    err   error
	msg  string
	code int
}

func (e *boxError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}

// ===  FUNCTION  ==========================================================
//         Name:  HasValidSize
//  Description:  Checks if a box is small enough to fit on an empty pallet.
// =========================================================================
func (b *box) HasValidSize() bool {
	return (b.w <= palletWidth) && (b.l <= palletLength) && (b.w > 0) && (b.l > 0)
} // -----  end of function HasValidSize  -----

// ===  FUNCTION  ==========================================================
//         Name:  HasValidCoordinates
//  Description:  Checks if a box is small enough to fit on an empty pallet.
// =========================================================================
func (b *box) HasValidCoordinates() bool {
	return (b.x < palletWidth) && (b.y < palletLength)
} // -----  end of function HasValidCoordinates  -----

// ===  FUNCTION  ==========================================================
//         Name:  Size
//  Description:  Calculates Size of a box.
// =========================================================================
func (b *box) Size() uint8 {
	return b.w * b.l
}

func (b *box) IsSquare() bool {
	return b.w == b.l
}

func (b *box) Display() string {
	c := b.canon()

	var out string
	var i, j uint8

	for i = 0; i < c.l; i++ {
		for j = 0; j < c.w; j++ {
			out += "x "
		}
		out += "\n"
	}
	return out
}

// ===  FUNCTION  ==========================================================
//         Name:  BoxesAreEqual
//  Description:  Compares if two Boxes are equal. Since we cannot simply
//                range over structs we have do it manually. The input is as
//                a value, not pointer to use this method in
//                `PalletsAreEqual` as we cannot range over pointers
// =========================================================================
func BoxesAreEqual(a, b box) bool {
	if a.x != b.x {
		return false
	}
	if a.y != b.y {
		return false
	}
	if a.w != b.w {
		return false
	}
	if a.l != b.l {
		return false
	}
	if a.id != b.id {
		return false
	}

	return true

	// It would be more elegant to use reflections to get the values of the
	// respected fields of the box struct. See also:
	// https://stackoverflow.com/qüstions/18926303/iterate-through-a-struct-in-go
	//
	// However, using reflect to iterate over the box structure fails as the
	// data field variables are all lower case in `box` and thus are invisible
	// outside the defining package and reflect is an outside package. See
	// https://groups.google.com/forum/#!topic/golang-nuts/UYgse9hnfoc
	//
	//
	//	A := reflect.ValueOf(a)
	//	B := reflect.ValueOf(b)
	//
	//	A_values := make([]interface{}, A.NumField())
	//	B_values := make([]interface{}, B.NumField())
	//
	//	for i := 0; i < A.NumField(); i++ {
	//		A_values[i] = A.Field(i).Interface()
	//	}
	//	for i := 0; i < B.NumField(); i++ {
	//		B_values[i] = B.Field(i).Interface()
	//	}
	//
	//	for i, v := range A_values {
	//		if v != B_values[i] {
	//			return false
	//		}
	//	}
	//	return true
} // -----  end of function BoxesAreEqual  -----

// ===  FUNCTION  ==========================================================
//         Name:  PalletssAreEqual
//  Description:  Compares if two Palletss are equal.
// =========================================================================
func PalletsAreEqual(a, b pallet) bool {
	if len(a.boxes) != len(b.boxes) {
		return false
	}
	for i, v := range a.boxes {
		if !BoxesAreEqual(v, b.boxes[i]) {
			return false
		}
	}
	return true
} // -----  end of function PalletssAreEqual  -----

// ===  FUNCTION  ==========================================================
//         Name:  BoxArraysAreEqual
//  Description:  Compares if two Palletss are equal.
// =========================================================================
func BoxArraysAreEqual(a, b []box) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !BoxesAreEqual(v, b[i]) {
			return false
		}
	}
	return true
} // -----  end of function BoxArraysAreEqual  -----

// ===  FUNCTION  ==========================================================
//         Name:  PutOnPallet
//  Description:
// =========================================================================
func (b box) PutOnPallet(p *pallet) {
	if BoxesAreEqual(b, emptybox) {
		return
	}
	if !b.HasValidCoordinates() {
		return
	}

	p.boxes = append(p.boxes, b)

} // -----  end of function PutOnPallet  -----

// ===  FUNCTION  ==========================================================
//         Name:  SetOrigin
//  Description:  TODO: Error Handling if inpuit x,y are out of bound.
// =========================================================================
func (b *box) SetOrigin(x, y uint8) {
	if x < 3 && y < 3 {
		b.x = x
		b.y = y
	}
} // -----  end of function SetOrigin  -----

// ===  FUNCTION  ==========================================================
//         Name:  Rotate
//  Description:
// =========================================================================
func (b *box) Rotate() {
	tmp := b.w
	b.w = b.l
	b.l = tmp
} // -----  end of function Rotate  -----

// =========================================================================
//  Implementing Sort interface
//  Will order boxes from lowest to highest size.
//  Use as:
//          boxes = []box
//          sort.Sort(BySize(boxes))
//
//	  			box{0, 0, 4, 4, 101},       box{0, 0, 2, 1, 103},
//	  			box{0, 0, 2, 2, 102},  -->  box{0, 0, 2, 2, 102},
//	  			box{0, 0, 2, 1, 103},       box{0, 0, 3, 2, 104},
//	  			box{0, 0, 3, 2, 104},       box{0, 0, 4, 4, 101},
// =========================================================================
type BySize []box

func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return a[i].Size() < a[j].Size() }

// -----  end of Sort Interface  -----
//func isPalleteFilled(p pallet) bool {
//}
//
//func fillPalleteWithBoxes(allBoxes []boxInfo)
//
