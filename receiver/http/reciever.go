package http

import (
	"fmt"
	"multithreading/multithreading-homework/logger"
	"multithreading/multithreading-homework/message/http"
)

type Receiver[U any, T http.Message[U]] struct {
	ch chan http.Message[U]
}

func (r Receiver[U, T]) PutMessage(message http.Message[U]) {
	logger.GetInstance().Info(fmt.Sprintf("Received a message: %s", message))
	r.ch <- message
}

func (r Receiver[U, T]) YieldMessage() http.Message[U] {
	message := <-r.ch
	logger.GetInstance().Info(fmt.Sprintf("Yielding a message: %s", message))
	return message
}

func NewHTTPReceiver[U any, T http.Message[U]](queueVol int) Receiver[U, T] {
	return Receiver[U, T]{
		ch: make(chan http.Message[U], queueVol),
	}
}
