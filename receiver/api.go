package receiver

import "multithreading/multithreading-homework/message"

type Receiver[U any, T message.Message[U]] interface {
	YieldMessage() T
	PutMessage(T)
}
