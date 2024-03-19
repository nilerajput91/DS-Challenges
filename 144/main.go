//	Insert Delete GetRandom O(1)

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomSet := constructor()
	fmt.Println(randomSet.Insert(10))
	fmt.Println(randomSet.Insert(20))
	fmt.Println(randomSet.Insert(30))
	fmt.Println(randomSet.Remove(10))
	fmt.Println(randomSet)

}

type RandmoziedSet struct {
	nums  []int
	index map[int]int
}

func constructor() RandmoziedSet {
	return RandmoziedSet{
		nums:  []int{},
		index: make(map[int]int),
	}
}

func (this *RandmoziedSet) Insert(data int) bool {
	if _, found := this.index[data]; found {
		return false
	}
	this.nums = append(this.nums, data)
	this.index[data] = len(this.nums) - 1
	return true

}

func (this *RandmoziedSet) Remove(data int) bool {
	idx, found := this.index[data]

	if !found {
		return false
	}

	last := len(this.nums) - 1

	lastVal := this.nums[last]
	this.nums[idx] = lastVal
	this.index[lastVal] = idx
	delete(this.index, data)
	this.nums = this.nums[:last]

	return true
}

func (this *RandmoziedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]

}
