package main

import (
	"encoding/json"
	"os"
	"strconv"
	"fmt"
	"github.com/williamsodell/goBallClockgo/helpers"
)

const HOUR_BALL_CAP = 11
const FIVE_MIN_BALL_CAP = 11
const ONE_MIN_BALL_CAP = 4

type Mode2 struct {
	Mins []int
	FiveMins []int
	Hours []int
	Main []int
}

func checkMainQueue(queue *BallQueue, balls int) bool {
	if queue.Len() != balls {
		return false
	}

	q := *queue

	if q[0] != 1 || q[balls - 1] != balls {
		return false
	}

	for i := 1; i <= balls; i++ {
		if q[i - 1] != i {
		 return false
		}
	}

	return true
}

func cycleQueue(queue, mainQueue *BallQueue) {
	for queue.Len() > 0 {
		mainQueue.Push(queue.Pop())
	}
}

func process(mainQueue, minQueue, fiveMinQueue, hourQueue *BallQueue) int {
	ball := mainQueue.Shift()

	if minQueue.Len() < ONE_MIN_BALL_CAP {
		minQueue.Push(ball)
		return 0
	}

	cycleQueue(minQueue, mainQueue)

	if fiveMinQueue.Len() < FIVE_MIN_BALL_CAP {
		fiveMinQueue.Push(ball)
		return 0
	}

	cycleQueue(fiveMinQueue, mainQueue)

	if hourQueue.Len() < HOUR_BALL_CAP {
		hourQueue.Push(ball)
		return 0
	}

	cycleQueue(hourQueue, mainQueue)

	mainQueue.Push(ball)
	return 1
}

func mode1(balls int) {
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
		twelveHours += process(mainQueue, minQueue, fiveMinQueue, hourQueue)

		finished = checkMainQueue(mainQueue, balls)
	}

	fmt.Println(twelveHours / 2)//Float.. and round
}

func mode2(balls, mins int) {
	fmt.Println("Mode 2")

	mainQueue := &BallQueue{}
	minQueue := &BallQueue{}
	fiveMinQueue := &BallQueue{}
	hourQueue := &BallQueue{}

	for i := 1; i <= balls; i++ {
		mainQueue.Push(i)
	}

	for i := 0; i < mins; i++ {
		process(mainQueue, minQueue, fiveMinQueue, hourQueue)
	}

	result := &Mode2{
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

		mode2(balls, mins)
		os.Exit(0)
	}

	mode1(balls)
}
