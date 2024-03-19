// merge sorted Array
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1
	j := n - 1
	k := m + n - 1

	//merge num1 and num2 from theend by comparing element

	for i >= 0 && j >= 0 {

		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

		k--
	}

	// Copy any remaining elements from nums2

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--

	}

}
