package main

import (
	"container/heap"
	"fmt"
)

type CharCount struct {
	char  rune
	count int
}

type ChCount []CharCount

func (h ChCount) Len() int           { return len(h) }
func (h ChCount) Less(i, j int) bool { return h[i].count > h[j].count }
func (h ChCount) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ChCount) Push(x interface{}) {
	*h = append(*h, x.(CharCount))
}

func (h *ChCount) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func rearrangeString(s string) string {

	if len(s) > 500 {
		return "Error: String is too long"
	}

	charCount := make(map[rune]int)
	maxHeap := &ChCount{}

	for _, char := range s {
		charCount[char]++
	}

	for char, count := range charCount {
		heap.Push(maxHeap, CharCount{char, count})
	}

	result := []rune{}
	var prev CharCount

	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(CharCount)
		result = append(result, current.char)

		if prev.count > 0 {
			heap.Push(maxHeap, prev)
		}

		current.count--
		prev = current
	}

	if len(result) != len(s) {
		return ""
	}

	return string(result)
}

func main() {
	test_500 := "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like)."
	test_valid := "aab"
	test_invalid := "aaab"
	fmt.Println(rearrangeString(test_500))
	fmt.Println(rearrangeString(test_valid))
	fmt.Println(rearrangeString(test_invalid))
}
