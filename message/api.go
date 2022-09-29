package message

type Message[T any] interface {
	GetData() T
}
