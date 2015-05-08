package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/williamsodell/goBallClockgo/helpers"
)

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

		helpers.Mode2(balls, mins)
		os.Exit(0)
	}

	helpers.Mode1(balls)
}
