//       Filename:  repack.go
//
//    Description:
//           TODO:  Add concurrency. Write Test.
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

// A repacker repacks trucks.
type repacker struct {
}

// This repacker is the worst possible, since it uses a new pallet for
// every box. Your job is to replace it with something better.
func oneBoxPerPallet(t *truck) (out *truck) {
	out = &truck{id: t.id}
	for _, p := range t.pallets {
		for _, b := range p.boxes {
			b.x, b.y = 0, 0
			out.pallets = append(out.pallets, pallet{boxes: []box{b}})
		}
	}
	return
}

// betterPacker takes an input truck pointer tp and unloads all its boxes
// into the Table pointer store, that might be empty or might contains boxes
// which were not packed on the last truck.
// Then the store is emptied box by box and the pallets of tp are filled
// until the store is empty, i.e. all boxes are packed.
func betterPacker(tp *truck, store *Table) (out *truck) {
	out = &truck{id: tp.id}

	// put all boxes of t in store
	nrPallets := tp.Unload(store)

	for i := 0; i < nrPallets && !store.IsEmpty(); i++ {
		var p pallet

		// freeGridSpace will track the free space on pallet
		freeGridSpace := NewInitialGrid()

		// As long as there is space keep on packing
		for !freeGridSpace.IsEmpty() && !store.IsEmpty() {

			// grab the last element of g, which hopefully has biggest size
			last := len(freeGridSpace) - 1
			e := freeGridSpace[last]

			// get corresponding box to that element
			b, _ := store.GetBoxThatFitsOrIsEmpty(e.size, e.orient)

			// If there are no more boxes in store that are as big or smaller than
			// freeGridSpace element e then break.
			if b == emptybox {
				break
			}

			// Put box on freeGridElement and set its coordinates b.x and b.y
			// correspondingly.
			// Return the splitting of the remaining freeSpace in
			// newFreeGridElements.
			newFreeGridElements := Put(&b, e)

			// Finally add box to pallet
			(&p).Add(b)

			freeGridSpace.Update(newFreeGridElements)
		} // end loop

		out.pallets = append(out.pallets, p)
	} //  end for pallets
	return
}

func newRepacker(in <-chan *truck, out chan<- *truck) *repacker {
	store := NewTable()
	go func() {
		for t := range in {
			// The last truck is indicated by its id. You might
			// need to do something special here to make sure you
			// send all the boxes.
			if t.id == idLastTruck {
				// TODO: not sure what is wanted here.
			}

			t = betterPacker(t, store)
			out <- t
		}
		// The repacker must close channel out after it detects that
		// channel in is closed so that the driver program will finish
		// and print the stats.
		close(out)
	}()
	return &repacker{}
}
