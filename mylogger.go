package asynclogger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	logChannel  chan string
	stopChannel chan bool
	wg          sync.WaitGroup
}

func NewLogger() *Logger {
	logger := &Logger{
		logChannel:  make(chan string, 1000),
		stopChannel: make(chan bool),
	}
	logger.startLogger()
	return logger
}

func (l *Logger) Log(message string) {
	l.logChannel <- message
}

func (l *Logger) startLogger() {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			return
		}
		defer file.Close()

		for {
			select {
			case logMessage := <-l.logChannel:
				logLine := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), logMessage)
				_, err := file.WriteString(logLine)
				if err != nil {
					fmt.Printf("Error writing to log file: %v\n", err)
				}
			case <-l.stopChannel:
				return
			}
		}
	}()
}

func (l *Logger) Stop() {
	close(l.stopChannel)
	l.wg.Wait()
}

