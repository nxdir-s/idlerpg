package consumers

import (
	"fmt"
	"os"

	"github.com/IBM/sarama"
)

type EngineConsumer struct {
	Ready chan bool
}

func NewEngineConsumer() (*EngineConsumer, error) {
	return &EngineConsumer{
		Ready: make(chan bool),
	}, nil
}

func (c *EngineConsumer) Reset() {
	c.Ready = make(chan bool)
}

func (c *EngineConsumer) AwaitSetup() {
	<-c.Ready
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *EngineConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *EngineConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *EngineConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Fprint(os.Stdout, "message channel was closed\n")
				return nil
			}

			fmt.Fprintf(os.Stdout, "Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
