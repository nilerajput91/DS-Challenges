// User
// You have records of employee name as string (Ename) and salary as positive integer (S). You have to sort the records on the basis of employee salary, if salary is same then use employee name for comparison.

// Example 1:

// Input: N = 2
// arr[] = {{xbnnskd, 100}, {geek, 50}}
// Output: {geek 50}, {xbnnskd 100}
// Explanation: geek has lowest salary
// as 50 and xbnnskd has more salary.
// Example 2:

// Input: N = 2
// arr[] = {{shyam, 50}, {ram, 50}}
// Output: ram 50 shyam 50

package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name   string
	salary int64
}

type BySalaryAndName []Employee

// Len implements sort.Interface.
func (BySalaryAndName) Len() int {
	panic("unimplemented")
}

// Less implements sort.Interface.
func (BySalaryAndName) Less(i int, j int) bool {
	panic("unimplemented")
}

// Swap implements sort.Interface.
func (BySalaryAndName) Swap(i int, j int) {
	panic("unimplemented")
}

func (a BySalaryAndName) len() int {

	return len(a)

}

func (a BySalaryAndName) swap(i, j int) {

	a[j], a[i] = a[i], a[j]
}

func (a BySalaryAndName) less(i, j int) bool {
	if a[i].salary != a[j].salary {

		return a[i].salary < a[j].salary
	}

	return a[i].Name < a[j].Name

}
func main() {
	arr1 := BySalaryAndName{{"xbnnskd", 100}, {"geek", 50}}
	sort.Sort(BySalaryAndName(arr1))
	fmt.Println(arr1)
}
