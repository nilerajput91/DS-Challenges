package main

import "fmt"

func main() {
	// Example 1
	arr1 := []int{0, 2, 1, 2, 0}
	sort012(arr1)
	fmt.Println(arr1) // Output: [0 0 1 2 2]

	// Example 2
	arr2 := []int{0, 1, 0}
	sort012(arr2)
	fmt.Println(arr2) // Output: [0 0 1]

}

func sort012(arr []int) {
	low,mid,high:=0,0,len(arr)-1

	for mid<=high{

		switch(arr[mid]){

		case 0:
			arr[low],arr[mid]=arr[mid],arr[low]
			low++
			mid++

		case 1:
				mid++

		case 2:

		arr[low],arr[high]	=arr[high],arr[low]
		high --

		}
	}

}
