package http

import (
	"errors"
	"fmt"
	"multithreading/homework/logger"
	"multithreading/homework/message/http"
)

type Processor[U any, T http.Message[U]] struct{}

func (p Processor[U, T]) ProcessMassage(message http.Message[U]) error {
	err := message.GetError()
	if err != nil {
		return fmt.Errorf("message was delivered with an error: %w", err)
	}

	messageType := message.GetType()
	if messageType == http.GET {
		logger.GetInstance().Info(fmt.Sprintf("Processing GET message with data %s\n", message.GetData()))
	} else if messageType == http.POST {
		logger.GetInstance().Info(fmt.Sprintf("Processing POST message with data %s\n", message.GetData()))
	} else {
		return errors.New("wrong message type")
	}

	return nil
}

func NewHTTPProcessor[U any, T http.Message[U]]() Processor[U, T] {
	return Processor[U, T]{}
}
