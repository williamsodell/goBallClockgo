package helpers

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
