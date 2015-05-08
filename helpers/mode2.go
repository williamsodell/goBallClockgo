package helpers

import (
	"encoding/json"
	"fmt"
)

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
