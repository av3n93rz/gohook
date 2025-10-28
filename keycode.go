// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package hook

import "github.com/vcaesar/keycode"

// MouseMap defines the robotgo hook mouse's code map
var MouseMap = keycode.MouseMap

// Keycode defines the robotgo hook key's code map
var Keycode = keycode.Keycode

// Special defines the special key map
var Special = keycode.Special

// init function to add extended function keys
func init() {
	// Add extended function keys F13-F24
	// These are common keycodes for extended function keys on Windows
	extendedFKeys := map[string]uint16{
		"f13": 124, // F13
		"f14": 125, // F14
		"f15": 126, // F15
		"f16": 127, // F16
		"f17": 128, // F17
		"f18": 129, // F18
		"f19": 130, // F19
		"f20": 131, // F20
		"f21": 132, // F21
		"f22": 133, // F22
		"f23": 134, // F23
		"f24": 135, // F24
	}
	
	// Add the extended function keys to the Keycode map
	for key, code := range extendedFKeys {
		Keycode[key] = code
	}
}
