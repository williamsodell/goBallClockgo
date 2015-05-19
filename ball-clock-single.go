package main

import (
	"os"
	"strconv"
	"fmt"
	"encoding/json"
)

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

const HOUR_BALL_CAP = 11
const FIVE_MIN_BALL_CAP = 11
const ONE_MIN_BALL_CAP = 4

func IsFinished(queue *BallQueue, balls int) bool {
	if queue.Len() != balls {
		return false
	}

	q := *queue

	if q[0] != 1 || q[balls - 1] != balls {
		return false
	}

	for i := 2; i < balls; i++ {
		if q[i - 1] != i {
		 return false
		}
	}

	return true
}

func CycleQueue(queue, mainQueue *BallQueue) {
	for queue.Len() > 0 {
		mainQueue.Push(queue.Pop())
	}
}

func Process(mainQueue, minQueue, fiveMinQueue, hourQueue *BallQueue) int {
	ball := mainQueue.Shift()

	if minQueue.Len() < ONE_MIN_BALL_CAP {
		minQueue.Push(ball)
		return 0
	}

  CycleQueue(minQueue, mainQueue)

	if fiveMinQueue.Len() < FIVE_MIN_BALL_CAP {
		fiveMinQueue.Push(ball)
		return 0
	}

  CycleQueue(fiveMinQueue, mainQueue)

	if hourQueue.Len() < HOUR_BALL_CAP {
		hourQueue.Push(ball)
		return 0
	}

  CycleQueue(hourQueue, mainQueue)

	mainQueue.Push(ball)
	return 1
}

func Mode1(balls int) {
	fmt.Println("Mode 1")

	//Initialize all the variables
	twelveHours := 0
	finished := false

	mainQueue := &BallQueue{}
	minQueue := &BallQueue{}
	fiveMinQueue := &BallQueue{}
	hourQueue := &BallQueue{}

	for i := 1; i <= balls; i++ {
		mainQueue.Push(i)
	}

	for !finished {
		twelveHours += Process(mainQueue, minQueue, fiveMinQueue, hourQueue)

		finished = IsFinished(mainQueue, balls)
	}

	fmt.Println(twelveHours / 2)//Float.. and round
}

type Mode2Result struct {
	Mins []int
	FiveMins []int
	Hours []int
	Main []int
}

func Mode2(balls, mins int) {
	fmt.Println("Mode 2")

	mainQueue := &BallQueue{}
	minQueue := &BallQueue{}
	fiveMinQueue := &BallQueue{}
	hourQueue := &BallQueue{}

	for i := 1; i <= balls; i++ {
		mainQueue.Push(i)
	}

	for i := 0; i < mins; i++ {
		Process(mainQueue, minQueue, fiveMinQueue, hourQueue)
	}

	result := &Mode2Result{
		Mins: *minQueue,
		FiveMins: *fiveMinQueue,
		Hours: *hourQueue,
		Main: *mainQueue,
	}

	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("The balls argument is required")
		os.Exit(1)
	}

	balls, err := strconv.Atoi(os.Args[1]);

	if err != nil || balls < 27 || balls > 127 {
		fmt.Println("The balls argument must be an int between 27 and 127(inclusive)")
		os.Exit(2)
	}

	if len(os.Args) > 2 {
		mins, err := strconv.Atoi(os.Args[2]);

		if err != nil {
			fmt.Println("The mins argument must be an int")
			os.Exit(3)
		}

		Mode2(balls, mins)
		os.Exit(0)
	}

	Mode1(balls)
}
