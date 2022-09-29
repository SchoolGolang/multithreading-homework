package http

import "fmt"

const (
	GET  = 0
	POST = 1
)

type MessageType int

type Message[T any] struct {
	method MessageType
	data   T
	err    error
}

func (m Message[T]) GetData() T {
	return m.data
}

func (m Message[T]) GetType() MessageType {
	return m.method
}

func (m Message[T]) GetError() error {
	return m.err
}

func (m *Message[T]) SetData(data T) {
	m.data = data
}

func (m *Message[T]) SetType(messageType MessageType) {
	m.method = messageType
}

func (m *Message[T]) SetError(err error) {
	m.err = err
}

func (m Message[T]) String() string {
	return fmt.Sprintf("{data: %s, type: %d, error: %s}", m.data, m.method, m.err)
}
