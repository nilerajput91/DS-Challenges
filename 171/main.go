package main

import "container/heap"

func main() {

}

type WordFreq struct {
	word  string
	count int
}

type wordFreqHeap []WordFreq

func (h wordFreqHeap) Len() int { return len(h) }
func (h wordFreqHeap) Less(i, j int) bool {
	return h[i].count < h[j].count || (h[i].count == h[j].count && h[i].word > h[j].word)
}
func (h wordFreqHeap) Swap(i, j int) {
	h[j], h[i] = h[i], h[j]
}
func (h *wordFreqHeap) Push(x interface{}) { *h = append(*h, x.(WordFreq)) }

func (h *wordFreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordFreqMap := make(map[string]int)

	for _, word := range words {
		wordFreqMap[word]++
	}

	h := wordFreqHeap{}
	heap.Init(h)

	for word, count := range wordFreqMap {
		heap.Push(h, WordFreq{word, count})
		if h.Len() > k {
			heap.Pop(h)
		}

	}
	result := make([]string, k)

	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(WordFreq).word
	}
	return result

}
