// dog is a package whith a function Years to calculates the age that the dog would have if it was a human being
package dog

// Years is a function which calculates the final age by multiplying the input by 7
//
// it takes a int in entry for the dog'age and returns a int for the human's age
func Years(a int) int {
	transformedAge := a * 7
	return transformedAge
}
