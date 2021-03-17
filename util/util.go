package util

type Notifier interface {
	Send(string) error
}

func Thing(n *Notifier) error {

	return (*n).Send("string")
}
