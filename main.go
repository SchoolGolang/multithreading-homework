package main

import "multithreading/multithreading-homework/application"

// Agenda: We have application that processes some user requests.
// Receiver structs are needed for messages routing after they come.
// Processor structs are needed to process messages. After processing Processor must give log of success or error in console.
// mock.EmulateUser(receiver, ctx) emulates user that sends some requests to your receiver
//
// Task: Implement processing loop using waitgroups and context
//
// Disclaimer: You are allowed to edit application.go
func main() {
	application.Run()
}
