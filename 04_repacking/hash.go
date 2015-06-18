//
// =========================================================================
//
//       Filename:  hash.go
//
//    Description:  Implements the hash table to store the box stacks in.
//
//        Version:  1.0
//        Created:  06/18/2015 05:45:39 PM
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

const (
	MAPSIZE = 17
)

func NewMap() map[uint]Stack {
	return make(map[uint]Stack)
}
