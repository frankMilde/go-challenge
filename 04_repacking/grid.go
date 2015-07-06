//       Filename:  grid.go
//
//    Description:  Implements a grid by using slices not lists.
//
//           TODO:  Concurrent implement of Put and Update
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

import (
	"fmt"
	"log"
)

type Orientation uint8

const (
	HORIZONTAL Orientation = iota
	VERTICAL
	SQUAREGRID
)

type GridElement struct {
	x, y   uint8 // coordinates of origin
	w, l   uint8 // width and length
	size   int
	orient Orientation // horizontal, vertical, square
}

type FreeGrid []GridElement

var emptygrid = GridElement{}

func (g FreeGrid) IsEmpty() bool { return len(g) == 0 }

// NewGrid returns an empty Freegrid
func NewGrid() FreeGrid {
	var g []GridElement
	return g
}

// New4x4Grid returns a FreeGrid with a single 4x4 GridElement.
func New4x4Grid() FreeGrid {
	init := GridElement{0, 0, 4, 4, 16, SQUAREGRID}
	f := []GridElement{init}
	return f
}

// NewSubGrid initializes a new FreeGrid with Element e.
func NewSubGrid(e GridElement) FreeGrid {
	f := []GridElement{e}
	return f
}

// SetProperties sets the size and orientation of a GridElement pointer ep
func (ep *GridElement) SetProperties() {

	ep.size = int(ep.l * ep.w)

	switch {
	case ep.l == ep.w:
		ep.orient = SQUAREGRID
	case ep.w > ep.l:
		ep.orient = HORIZONTAL
	case ep.w < ep.l:
		ep.orient = VERTICAL
	}
}

// Put takes a box b and puts it in the top left corner of Gridelement e.
// If b does not cover e completely, the remaining free space of grid e is
// returned. This return value is of type FreeGrid := []GridElement and
// contains up to three elements into which the original e has been
// split by the box: (1) bottom, (2) right, (3) bottom right
//  | b b b 2 |
//  | b b b 2 |
//  | 1 1 1 3 |
//  | 1 1 1 3 |
func Put(b *box, e GridElement) FreeGrid {

	errCoor := b.SetOrigin(e.x, e.y)
	if errCoor != nil {
		log.Println("Error when setting origin ", e.x, e.y, " of box ", b.l, b.w, b.id)
		log.Println(e)
		log.Println(b)
	}

	bottom := GridElement{
		x: b.x + b.l,
		y: b.y,
		w: b.w,
		l: e.l - b.l,
	}
	right := GridElement{
		x: b.x,
		y: b.y + b.w,
		w: e.w - b.w,
		l: b.l,
	}
	bottomRight := GridElement{
		x: b.x + b.l,
		y: b.y + b.w,
		w: e.w - b.w,
		l: e.l - b.l,
	}
	bottom.SetProperties()
	right.SetProperties()
	bottomRight.SetProperties()

	elements := []GridElement{right, bottomRight, bottom}

	var split FreeGrid

	for _, e := range elements {
		if e.size != 0 {
			split = append(split, e)
		}
	}

	return split
}

// Update cuts out last element of FreeGrid pointer fp and replaces it with
// a new FreeGrid newG.
func (fp *FreeGrid) Update(newG FreeGrid) {

	// Cut out last element
	last := len(*fp) - 1
	(*fp) = (*fp)[:last]

	// Append new FreeGrid and sort
	if !newG.IsEmpty() {
		*fp = append(*fp, newG...)
	}
}

// GridElementsAreEqual returns true, if two GridElements a,b are equal in
// each field.
func GridElementsAreEqual(a, b GridElement) bool {
	if a.size != b.size {
		return false
	}
	if a.orient != b.orient {
		return false
	}
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

	return true
}

// FreeGridsAreEqual returns true, if all GridElements of two FreeGrids a,b
// are equal.
func FreeGridsAreEqual(a, b FreeGrid) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !GridElementsAreEqual(v, b[i]) {
			return false
		}
	}
	return true
}

// String interface to pretty print Orientation.
func (orient Orientation) String() string {

	var s string

	switch orient {
	case HORIZONTAL:
		s = "horizontal"
	case VERTICAL:
		s = "vertical"
	case SQUAREGRID:
		s = "square"
	}

	return s
}

// String interface to pretty print GridElement.
func (e GridElement) String() string {

	var s string
	s += fmt.Sprintf("[%d %d %d %d] ", e.x, e.y, e.w, e.l)
	s += fmt.Sprintf("%d %v ", e.size, e.orient)
	return s
}

// String interface to pretty print FreeGrid.
func (g FreeGrid) String() string {

	var s string
	for i, g := range g {
		boxtmp := box{g.x, g.y, g.w, g.l, 1}
		grid := pallet{[]box{boxtmp}}
		if i < 10 {
			s += fmt.Sprintf("[ %d]   -->   %v,%v\n", i, g, grid)
		} else {
			s += fmt.Sprintf("[%d]   -->   %v,%v\n", i, g, grid)
		}
	}
	return s
}

//  Implementing Sort interface
//  Will order grids from lowest to highest size.
//  Use as:
//          grid = []GridElements
//          sort.Sort(ByArea(grid))
type ByArea []GridElement

func (a ByArea) Len() int           { return len(a) }
func (a ByArea) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByArea) Less(i, j int) bool { return a[i].size < a[j].size }

// -----  end of Sort Interface  -----
