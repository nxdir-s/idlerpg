package secondary

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type ClientError struct {
	err error
}

func (e ClientError) Error() string {
	return "failed to start sarama producer: " + e.err.Error()
}

type SaramaAdapter struct {
	producer sarama.SyncProducer
}

func NewSaramaAdapter(brokers []string) (*SaramaAdapter, error) {
	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	// tlsConfig := createTlsConfiguration()
	// if tlsConfig != nil {
	// 	config.Net.TLS.Config = tlsConfig
	// 	config.Net.TLS.Enable = true
	// }

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return &SaramaAdapter{}, ClientError{err}
	}

	return &SaramaAdapter{
		producer: producer,
	}, nil
}

// SendPlayerEvent marshals and sends player event to kafka
func (a *SaramaAdapter) SendPlayerEvent(ctx context.Context, event *valobj.PlayerEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	_, _, err = a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "player-events",
		Value: sarama.StringEncoder(data),
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *SaramaAdapter) Close() error {
	return a.producer.Close()
}
