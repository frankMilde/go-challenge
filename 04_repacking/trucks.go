//       Filename:  trucks.go
//
//    Description:  Handles all truck related things.
//           TODO:  Add test for TrucksAreEuals.
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

import "fmt"

var emptyTruck = truck{}

// Unload adds all boxes b from truck pointer tp to Table table and returns
// number of pallets nr.
func (tp truck) Unload(table *Table) int {
	var nrPallets int
	for i, p := range tp.pallets {
		for _, b := range p.boxes {
			table.Add(b.canon())
		}
		nrPallets = i
	}
	return nrPallets + 1
}

// TrucksAreEqual returns true if Trucks t1, t2 have the same number of
// pallets and the pallets are equal.
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
}

// String interface to pretty print truck.
func (t truck) String() string {
	s := fmt.Sprintf("Truck %d\n", t.id)
	for i, p := range t.pallets {
		if i < 10 {
			s += fmt.Sprintf("    [ %d]%v\n", i, p)
		} else {
			s += fmt.Sprintf("    [%d]%v\n", i, p)
		}
	}
	return s
}
