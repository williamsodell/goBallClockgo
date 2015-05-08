package main

import (
	"encoding/json"
	"os"
	"strconv"
	"fmt"
	"github.com/williamsodell/goBallClockgo/helpers"
)

type Mode2 struct {
	Mins []int
	FiveMins []int
	Hours []int
	Main []int
}

func mode1(balls int) {
	fmt.Println("Mode 1")

	//Initialize all the variables
	twelveHours := 0
	finished := false

	mainQueue := &helpers.BallQueue{}
	minQueue := &helpers.BallQueue{}
	fiveMinQueue := &helpers.BallQueue{}
	hourQueue := &helpers.BallQueue{}

	for i := 1; i <= balls; i++ {
		mainQueue.Push(i)
	}

	for !finished {
		twelveHours += helpers.Process(mainQueue, minQueue, fiveMinQueue, hourQueue)

		finished = helpers.IsFinished(mainQueue, balls)
	}

	fmt.Println(twelveHours / 2)//Float.. and round
}

func mode2(balls, mins int) {
	fmt.Println("Mode 2")

	mainQueue := &helpers.BallQueue{}
	minQueue := &helpers.BallQueue{}
	fiveMinQueue := &helpers.BallQueue{}
	hourQueue := &helpers.BallQueue{}

	for i := 1; i <= balls; i++ {
		mainQueue.Push(i)
	}

	for i := 0; i < mins; i++ {
		helpers.Process(mainQueue, minQueue, fiveMinQueue, hourQueue)
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
