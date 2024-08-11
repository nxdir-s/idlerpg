package valobj

type Event struct {
	Body     *Message
	Consumed chan bool
}
