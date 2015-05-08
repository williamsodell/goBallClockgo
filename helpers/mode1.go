package helpers

import (
	"fmt"
)

func Mode1(balls int) {
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
		twelveHours += Process(mainQueue, minQueue, fiveMinQueue, hourQueue)

		finished = IsFinished(mainQueue, balls)
	}

	fmt.Println(twelveHours / 2)//Float.. and round
}
