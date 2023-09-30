package main

type Word struct {
	word  string
	point int
}

type WordHeap []Word

func (h WordHeap) Len() int           { return len(h) }
func (h WordHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h WordHeap) Less(i, j int) bool { return h[i].point > h[j].point }
func (h *WordHeap) Push(e interface{}) {
	*h = append(*h, e.(Word))
}
func (h *WordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
