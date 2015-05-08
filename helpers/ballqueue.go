package helpers

type BallQueue []int

func (h BallQueue) Len() int { return len(h) }

func (h *BallQueue) Push(x int) {
	*h = append(*h, x)
}

func (h *BallQueue) Pop() int {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func (h *BallQueue) Shift() int {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1 : n]
	return x
}
