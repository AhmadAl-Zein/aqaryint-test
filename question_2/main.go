package main

import (
	"container/heap"
	"fmt"
)

const maxChar = 26

type charFreq struct {
	ch   rune
	freq int
}

type charFreqHeap []charFreq

func (h charFreqHeap) Len() int           { return len(h) }
func (h charFreqHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h charFreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *charFreqHeap) Push(x interface{}) {
	*h = append(*h, x.(charFreq))
}

func (h *charFreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeString(str string) string {
	N := len(str)

	count := make([]int, maxChar)
	for i := 0; i < N; i++ {
		count[int(str[i]-'a')]++
	}

	pq := &charFreqHeap{}
	heap.Init(pq)
	for c := 'a'; c <= 'z'; c++ {
		val := int(c - 'a')
		if count[val] > 0 {
			heap.Push(pq, charFreq{ch: c, freq: count[val]})
		}
	}

	result := ""

	var prev charFreq

	for pq.Len() > 0 {
		k := heap.Pop(pq).(charFreq)
		result += string(k.ch)

		if prev.freq > 0 {
			heap.Push(pq, prev)
		}

		k.freq--
		prev = k
	}

	// If length is not the same => not possible
	if N != len(result) {
		return ""
	}

	return result
}

func main() {
	str := "aab"
	result := rearrangeString(str)
	fmt.Println(result)
}
