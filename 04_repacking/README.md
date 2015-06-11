Introduction
============
In the logistics industry poorly packed boxes are wasting space, and worse,
tying up pallets that could be resold for a profit. We are in charge of
programming robots to repack the boxes.

You are given a truck full of a pallets with boxes on them that may or may
not be correctly packed. Your task is to implement an algorithm that packs
boxes onto the pallets correctly. A correctly packed pallet is one where
none of the boxes on it overlap, and none of the boxes hang over the edge of
the pallet. Pallets are packed in only two dimensions, with a single layer
of boxes which have an arbitrary height. All of the trucks are going the
same place anyway, so it doesn't matter which truck a box goes in, as long
as it is packed correctly on a pallet.

Empty pallets left over after repacking are pure profit. More empties = more
better! And if a truck leaves the warehouse with more pallets on it than it
came with, it comes out of your profit. So pack carefully!


Boxes
=====

A box is a `box struct`, including its position on the pallet `x`,`y` and its width
and length `w`,`l`. Its `id` is unique across all the boxes in one input file.
```
type box struct {
	x, y uint8
	w, l uint8
	id   uint32
}
```

Possible box types with respected area (using hex values `c=12` and `f=16`) are:
```
1 22 333 4444

44 666 8888
44 666 8888

999 cccc
999 cccc
999 cccc

ffff
ffff
ffff
ffff
```
These are all boxes that can be placed on a 4x4 pallet. However, the input will give
us boxes that are even bigger that `f`. These have to be filtered out.

Note, that an area uniquely identifies the box type, except for an area of 4.

Palettes
========
A pallet holds a collections of `boxes`, each in a certain place on a grid.
```
type pallet struct {
	boxes []box
}
```
All pallets have
```
const palletWidth = 4
const palletLength = 4
```
A palette string is a comma separated list of boxes and look like this:
```
0 0 1 1 101, 1 1 1 1 102, 2 2 1 1 103, 3 0 4 1 104
```
This particular pallet could be visualized as follows:
```
| @       |
|   &     |
|     #   |
| $ $ $ $ |
```

Trucks
======

A truck has an unique `id` and contains a slice of `pallets`.
```
type truck struct {
	id      int
	pallets []pallet
}
```
A truck `string` starts with `truck <id>`, and ends with `endtruck`. Inside
of a truck, there's one pallet per line.
```
truck 1
0 0 1 1 101,1 1 1 1 102,2 2 1 1 103,3 0 4 1 104
0 0 1 1 101,0 0 1 1 102
0 0 5 5 101
endtruck
```

Functions
=========

Function `paint` will take a pallet (as a list of boxes) and tries to fill a
pallet grid with them.

Repackaging trucks
==================

Pack pallets as tight as possible. If a pallet is not full, hold it back
until it can be filled nicely and put it on the next truck.

Algorithm Idea
==============

Each time we place a box onto a free space, we divide the remaining free
space into two regions. We then try to fill the bigger of the two spaces.

![Free space tree structure](tree.png)

When we joint the open ends of the tree we get the total remaining free space.

![Combined free space](add-free-space.png)

Resources
=========
- http://golang-challenge.com/go-challenge4/
- http://0xax.blogspot.de/2014/08/binary-tree-and-some-generic-tricks.html
- http://nathanleclaire.com/blog/2014/07/19/demystifying-golangs-io-dot-reader-and-io-dot-writer-interfaces/
- https://golang.org/pkg/fmt/

