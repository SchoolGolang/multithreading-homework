package mock

import (
	"context"
	"errors"
	"math/rand"
	"multithreading/homework/message/http"
	httpRec "multithreading/homework/receiver/http"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//EmulateUser Can emulate input to given receiver, context needed to end execution after cancellation function was called
func EmulateUser(receiver httpRec.Receiver[string, http.Message[string]], ctx context.Context) {
	go func() {
		rand.Seed(time.Now().UnixNano())

		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg := http.Message[string]{}

				msg.SetType(http.MessageType(rand.Intn(3)))
				msg.SetData(randStringRunes(rand.Intn(20)))
				if rand.Intn(1) > 0 {
					msg.SetError(errors.New("some random stuff happens"))
				}

				receiver.PutMessage(msg)
				time.Sleep(time.Duration(100*rand.Intn(10)) * time.Millisecond)
			}
		}
	}()
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
