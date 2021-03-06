// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/rfTJFcUn9m

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the Search method for hockey, call into the Search method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import "strings"

// matcher defines the behavior required for performing searches.
type matcher interface {
	search(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// search checks the value for the specified term.
func (s sport) search(searchTerm string) bool {
	if strings.Contains(s.team, searchTerm) ||
		strings.Contains(s.city, searchTerm) {
		return true
	}

	return false
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.

// Implement the matcher interface for hockey.
func ( /* receiver type */ ) search(searchTerm string) bool {
	// Make sure you call into Search method for the embedded
	// sport type.

	// Implement the search for the new fields.
	return false
}

// main is the entry point for the application.
func main() {
	// Define the term to search.

	// Create a slice of matcher values to search.

	// Display what we are searching for.

	// Range of each matcher value and check the search term.
}
