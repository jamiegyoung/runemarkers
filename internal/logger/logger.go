package logger

import (
	"fmt"
	"time"
)

func New(prefix string) func(string) {
	return func(message string) {
		now := time.Now().Format("15:04:05")
		fmt.Printf("[%+v %+v] %+v\n", now, prefix, message)
	}
}
