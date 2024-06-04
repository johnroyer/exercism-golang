// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

import "math"

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
    if 0 == math.Mod(float64(year), 400) {
        return true
    }
    if 0 == math.Mod(float64(year), 100) {
        return false
    }
    if 0 == math.Mod(float64(year), 4) {
        return true
    }
    return false;
}
