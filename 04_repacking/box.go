//       Filename:  box.go
//    Description:  Handles all things related to boxes.
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

import "errors"

// HasValidDimensions returns true, if box b is small enough to fit on an
// empty pallet and has a non-zero length and width.
func (b box) HasValidDimensions() bool {
	return (b.l <= palletWidth) && (b.w <= palletLength) && (b.l > 0) && (b.w > 0)
}

// ValidCoordinates returns true, if coord x,y are within pallet bounds.
func ValidCoordinates(x, y uint8) bool {
	return (y < palletWidth) && (x < palletLength)
}

// HasValidCoordinates returns true, if origin of box is within pallet
// bounds.
func (b box) HasValidCoordinates() bool {
	return ValidCoordinates(b.x, b.y)
}

// IsWithinBounds returns true if a box fits within the pallet bounds.
func (b box) IsWithinBounds(x, y uint8) bool {
	boxIsTooWide := (b.l + x) > palletWidth
	boxIsTooLong := (b.w + y) > palletLength
	return (!boxIsTooWide && !boxIsTooLong)
}

func (b box) Size() uint8    { return b.l * b.w }
func (b box) IsSquare() bool { return b.l == b.w }
func (b *box) Rotate() {
	tmp := b.l
	b.l = b.w
	b.w = tmp
}

// Display prints a graphic representation of box b on the terminal.
func (b box) Display() string {
	c := b.canon()

	var out string
	var i, j uint8

	for i = 0; i < c.w; i++ {
		for j = 0; j < c.l; j++ {
			out += "x "
		}
		out += "\n"
	}
	return out
}

// BoxesAreEqual returns true, if two boxes a,b are equal. It serves as a
// basis for BoxArraysAreEqual() and PalletsAreEqual().
// The input is as value, not pointer to use this method in
// PalletsAreEqual() as we cannot range over pointers.
func BoxesAreEqual(a, b box) bool {
	if a.x != b.x {
		return false
	}
	if a.y != b.y {
		return false
	}
	if a.l != b.l {
		return false
	}
	if a.w != b.w {
		return false
	}
	if a.id != b.id {
		return false
	}

	return true
}

// BoxArraysAreEqual returns true, if two []box a,b are equal.
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
}

// PalletsAreEqual returns true, if two pallets a,b are equal.
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
}

// Add takes a box b and appends it to the boxlist of pallet pointer
// pp. If the box is empty or in other form invalid nothing happens.
// TODO: Add Error hanging.
func (pp *pallet) Add(b box) {
	if BoxesAreEqual(b, emptybox) {
		return
	}
	if !b.HasValidCoordinates() {
		return
	}

	pp.boxes = append(pp.boxes, b)
}

// SetOrigin sets origin coordinates x,y of box pointer bp and checks if box is still
// valid. Returns Error when failed.
func (bp *box) SetOrigin(x, y uint8) error {
	if !ValidCoordinates(x, y) {
		return errors.New("box: Origin coordinates out of bounds.")
	}
	if !bp.HasValidDimensions() {
		return errors.New("box: Has invalid size.")
	}
	if bp.IsWithinBounds(x, y) {
		bp.x = x
		bp.y = y
		return nil
	} else {
		bp.Rotate()
		if bp.IsWithinBounds(x, y) {
			bp.x = x
			bp.y = y
			return nil
		} else {
			return errors.New("box.SetOrigin: Hangs over pallet edge. Unable to place box on grid")
		}
		return errors.New("box.SetOrigin: Hangs over pallet edge. Unable to place box on grid")
	}
}

// Implementing Sort interface.
// Will order boxes from lowest to highest size.
// Use as:
//          boxes = []box
//          sort.Sort(BySize(boxes))
//
//	  			box{0, 0, 4, 4, 101},       box{0, 0, 2, 1, 103},
//	  			box{0, 0, 2, 2, 102},  -->  box{0, 0, 2, 2, 102},
//	  			box{0, 0, 2, 1, 103},       box{0, 0, 3, 2, 104},
//	  			box{0, 0, 3, 2, 104},       box{0, 0, 4, 4, 101},
type BySize []box

func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return a[i].Size() < a[j].Size() }

// -----  end of Sort Interface  -----
