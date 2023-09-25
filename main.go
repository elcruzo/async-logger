package main

import (
	"github.com/elcruzo/async-logger"
	"time"
)

func main() {
	logger := asynclogger.NewLogger()

	for i := 0; i < 10; i++ {
		go func(index int) {
			for j := 0; j < 5; j++ {
				message := fmt.Sprintf("Log message %d-%d", index, j)
				logger.Log(message)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	logger.Stop()
}
