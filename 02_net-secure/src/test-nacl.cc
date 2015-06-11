//
// =========================================================================
//
//       Filename:  test-nacl.c
//
//    Description:  
//
//        Version:  1.0
//        Created:  04/08/2015 02:48:40 PM
//       Revision:  none
//       Compiler:  g++
//
//          Usage:  
//
//         Output:  
//
//         Author:  Frank Milde (FM), frank.milde (at) posteo.de
//        Company:  
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde
//
// =========================================================================

#include <iostream>

#include <sodium.h>
     
int main() {
	std::string pk;
	std::string sk;

	if (sodium_init() == -1) {
		return 1;
	}

	pk = crypto_box_keypair(&sk);

	std::cout << pk << std::endl;
	std::cout << sk << std::endl;

	return 0;
}


