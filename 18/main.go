package main

import (
	"fmt"
	"sort"
)

func main(){
	str:=[]string{"zy","st"}
fmt.Println(check(str))

}

func check(arr []string) bool{
	isSorted:=sort.StringsAreSorted(arr)
	return isSorted
}