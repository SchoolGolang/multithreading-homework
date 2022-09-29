package processor

import "multithreading/multithreading-homework/message"

type Processor[U any, T message.Message[U]] interface {
	ProcessMassage(T) error
}
