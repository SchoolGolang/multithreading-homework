package application

import (
	"context"
	"errors"
	"fmt"
	"multithreading/homework/logger"
	"multithreading/homework/message/http"
	"multithreading/homework/mock"
	httpProc "multithreading/homework/processor/http"
	httpRec "multithreading/homework/receiver/http"
)

func Run() {
	//This is just an example of how you can use given functionality
	receiver := httpRec.NewHTTPReceiver[string](10)
	processor := httpProc.NewHTTPProcessor[string]()
	mock.EmulateUser(receiver, ctx) // emulates user that sends some requests to your receiver

	var msg = http.Message[string]{}

	msg.SetData("adasdasd")
	msg.SetType(http.GET)
	receiver.PutMessage(msg)

	msg.SetType(http.POST)
	receiver.PutMessage(msg)

	msg.SetType(3)
	receiver.PutMessage(msg)

	msg.SetType(http.GET)
	msg.SetError(errors.New("some random stuff happen"))
	receiver.PutMessage(msg)

	for i := 0; i < 3; i++ {
		err := processor.ProcessMassage(receiver.YieldMessage())
		if err != nil {
			logger.GetInstance().Warn(fmt.Sprintf("An error catched from processor: %s\n", err))
		}
	}
}
