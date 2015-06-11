
Boxes
=====

A box is a box, including its position on the pallet `x`,`y` and its width
`w` and length `l`. Its ID is unique across all the boxes in one input file.
```
type box struct {
	x, y uint8
	w, l uint8
	id   uint32
}
```

With the help of [ASCII flow](http://asciiflow.com/#Draw) possible Box types
with area are:


```
+--+  +------+  +----------+  +--------------+
|1 |  |  2   |  |    3     |  |       4      |
+--+  +------+  +----------+  +--------------+

+------+ +----------+ +--------------+
|      | |          | |              |
|      | |          | |              |
|  4   | |     6    | |      8       |
|      | |          | |              |
+------+ +----------+ +--------------+


+----------+ +--------------+
|          | |              |
|          | |              |
|          | |              |
|    9     | |     12       |
|          | |              |
|          | |              |
|          | |              |
+----------+ +--------------+

+--------------+
|              |
|              |
|              |
|              |
|      16      |
|              |
|              |
|              |
|              |
|              |
+--------------+
```

An area uniquely identifies the box type, except for an area of 4.
So boxes will be displayed using hex values according to their area

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
`
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
