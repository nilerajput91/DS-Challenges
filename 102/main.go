//Subarray with Given Sum

package main

import "fmt"

func main() {
	arr:=[]int{3,2,0,4,7}
	n:=5
	sum:=6

	fmt.Printf("sum achieve :%v\n",subArray(arr,n,sum))

}

func subArray(arr []int, n, sum int) bool {
	for i := 0; i < n; i++ {
		curr:=0

		for j:=i;j<n;j++{
			curr+=arr[i]

			if curr==sum{
				return true
			}
		}
	}

	return false

}
