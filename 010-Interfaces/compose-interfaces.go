package main

import "fmt"

// ------------------------------------

type Server interface {
	MakeRequest() Response
}

type HelloWorldServer struct {
	Logger   Logger
	Response Response
}

func (s HelloWorldServer) MakeRequest() Response {
	s.Logger.Log(s.Response.Contents())
	return s.Response
}

// ------------------------------------

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println("Console log:", message)
}

// ------------------------------------

type Response interface {
	Contents() string
}

type HelloWorldResponse struct{}

func (l HelloWorldResponse) Contents() string {
	return "Hello World!"
}

// ------------------------------------

func main() {
	server := &HelloWorldServer{
		Logger:   &ConsoleLogger{},
		Response: &HelloWorldResponse{},
	}

	response := server.MakeRequest()
	fmt.Println("Response:", response.Contents())
}
