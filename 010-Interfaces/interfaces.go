package main

import "fmt"

// Any struct with a void-returning "Log" function that accepts only one
// string as an argument will be considered an implementation of "Logger"
type Logger interface {
	Log(message string)
}

type SqlLogger struct{}

func (l SqlLogger) Log(message string) {
	fmt.Println("Sql log:", message)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println("Console log:", message)
}

type FileLogger struct{}

func (l FileLogger) Log(message string) {
	fmt.Println("File log:", message)
}

func main() {
	logger := &ConsoleLogger{}
	// We pass a concrete implementation to the function
	use_logger(logger, "This is a test log message")
}

// This function's argument has the type of the interface
func use_logger(logger Logger, message string) {
	logger.Log(message)
}
