package logger

import "fmt"

func New(prefix string) func(string) {
  return func(message string) {
    fmt.Println("[" + prefix + "] " + message)
  }
}
