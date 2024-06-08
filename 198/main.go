/* In the context of arrays in computer science, "global inversions" and "local inversions" refer to specific patterns within an array.

1. **Global Inversions**: A global inversion occurs when there is an inversion (i.e., when the order of
	 two elements in the array is reversed) between any pair of elements (i, j) where i < j in the array A.
	 In simpler terms,
 a global inversion is an inversion that spans the entire array.

2. **Local Inversions**: A local inversion occurs when there is an inversion between adjacent elements in
the array A. In other words, for a local inversion to exist,
it means that there are elements A[i] and A[i+1] such that A[i] > A[i+1] for some index i.

Here's an example to illustrate the concepts:

Consider the array A = [1, 3, 2, 4, 5].

- **Global Inversions**:
  - (1, 3), (1, 2), (1, 4), (1, 5), (3, 2), (3, 4), (3, 5), (2, 4), (2, 5) are all global inversions.
  - (2, 4) and (2, 5) are global inversions because 2 comes before 4 and 5 in the array.

- **Local Inversions**:
  - (3, 2) is a local inversion because 3 comes before 2 in the array.

In this example, there are both global and local inversions. However,
not all global inversions are local inversions. All local inversions are global inversions,
but not all global inversions are local, as global inversions can span across non-adjacent elements. */

package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 4, 5}

	// Check for global inversions
	if globalInversion(arr) {
		fmt.Println("No global inversions")
	} else {
		fmt.Println("Global inversions exist")
	}

	// Check for local inversions
	if localInversion(arr) {
		fmt.Println("No local inversions")
	} else {
		fmt.Println("Local inversions exist")
	}

}

// GlobalInversions returns true if there are no global inversions in the array,
// otherwise returns false.
func globalInversion(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			// If there exists a pair (i, j) such that arr[i] > arr[j] and i < j,
			// it indicates a global inversion.
			if arr[i] > arr[j] {
				return false
			}
		}
	}
	return true

}

// LocalInversions returns true if there are no local inversions in the array,
// otherwise returns false.
func localInversion(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		// If there exists an adjacent pair (i, i+1) such that arr[i] > arr[i+1],
		// it indicates a local inversion.

		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
