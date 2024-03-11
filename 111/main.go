/* LRU design

For Cache size = 4
Reference Sequence: 10, 20, 10, 30, 40, 50, 30, 40

Expected Cache Behaiviour according to LRU:

Ref Seq Element                                 Hit/Miss                               Cache
     10                                           Miss                              10
     20                                           Miss                              20 10
     10                                           Hit                               10 20
     30                                           Miss                              30 10 20
     40                                           Miss                              40 30 10 20
     50                                           Miss                              50 40 30 10
     30                                           Hit                               30 50 40 10
     40                                           Hit                               40 30 50 10
*/

package main

import "fmt"

func main() {
	cache := Constructor(4)

	// Reference sequence: 10, 20, 10, 30, 40, 50, 30, 40
	// Expected cache behavior according to LRU:
	// Ref Seq Element   Hit/Miss   Cache
	//      10             Miss       10
	//      20             Miss       20 10
	//      10             Hit        10 20
	//      30             Miss       30 10 20
	//      40             Miss       40 30 10 20
	//      50             Miss       50 40 30 10
	//      30             Hit        30 50 40 10
	//      40             Hit        40 30 50 10

	fmt.Printf("Cache after inserting 10: ")
	cache.Put(10, 10)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 20: ")
	cache.Put(20, 20)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 10: ")
	cache.Put(10, 10)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 30: ")
	cache.Put(30, 30)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 40: ")
	cache.Put(40, 40)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 50: ")
	cache.Put(50, 50)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 30: ")
	cache.Put(30, 30)
	fmt.Println(cachetoString(cache))

	fmt.Printf("Cache after inserting 40: ")
	cache.Put(40, 40)
	fmt.Println(cachetoString(cache))

}

// Node represents a node in the doubly linked list
type Node struct {
	Key, Value int
	Prev, Next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

// Constructor initializes the LRUCache with a given capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
}

// Get retrieves the value of the key if the key exists in the cache, otherwise returns -1
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.movetoHead(node)
		return node.Value

	}
	return -1
}

// Put inserts the key-value pair into the cache. If the key already exists, updates the value. If the cache is full, evicts the least recently used item before inserting the new item.
func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.cache[key]; ok {
		node.Value = value
		lru.movetoHead(node)
	} else {
		if len(lru.cache) >= lru.capacity {
			delete(lru.cache, lru.tail.Prev.Key)
			lru.removeNode(lru.tail.Prev)
		}

		newNode := &Node{Key: key, Value: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode)
	}
}

func (lru *LRUCache) movetoHead(node *Node) {

	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	
	node.Next.Prev= node.Prev

}

func (lru *LRUCache) addToHead(node *Node) {
	node.Prev = lru.head
	node.Next = lru.head.Next

	if lru.head.Next != nil {
		lru.head.Next.Prev = node
	}
	lru.head.Next = node

}

// cachetoString converts the cache into a string representation for printing
func cachetoString(cache LRUCache) string {
	result := ""

	node := cache.head.Next

	for node != nil && node != cache.tail {
		result += fmt.Sprintf("%d ", node.Value)
		node = node.Next
	}
	return result

}
