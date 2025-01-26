package valobj

import "context"

type Event struct {
	Ctx      context.Context
	Msg      Message
	Consumed chan struct{}
}

type Message struct {
	Value string `json:"value"`
}
