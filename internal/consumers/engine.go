package consumers

import (
	"context"
	"fmt"
	"os"

	"github.com/IBM/sarama"
)

type EngineConsumer struct {
	Ready chan bool
}

// NewEngineConsumer creates a new EngineConsumer
func NewEngineConsumer(ctx context.Context) (*EngineConsumer, error) {
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

func (c *EngineConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Fprint(os.Stdout, "message channel was closed...\n")
				return nil
			}

			fmt.Fprintf(os.Stdout, "message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
