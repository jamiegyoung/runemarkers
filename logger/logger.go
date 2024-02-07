package logger

import "fmt"

func Logger(prefix string) func(string) {
	return func(message string) {
		fmt.Printf("[%s] %s\n", prefix, message)
	}
}
