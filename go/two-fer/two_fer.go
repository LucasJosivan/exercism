// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith returns a string whith given name when is not empty,
// and without name when is empty.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
