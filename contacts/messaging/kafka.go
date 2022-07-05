package messaging

import (
	"context"

	"github.com/golang/glog"
	"github.com/segmentio/kafka-go"
)

type ProduceMessage struct {
	Brokers string // = "localhost:29092"
	Topic   string
	Key     []byte
	Data    []byte
}

type SubscribeMessage struct {
	Brokers string // = "localhost:29092"
	Topic   string
}

func (m *ProduceMessage) Produce(ctx context.Context) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{m.Brokers},
		Topic:   m.Topic,
		//Dialer:  dialer,
	})
	return w.WriteMessages(ctx, kafka.Message{Key: m.Key, Value: m.Data})
}

func (s *SubscribeMessage) Subscribe(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{s.Brokers},
		Topic:   s.Topic,
		//Dialer:  dialer,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			glog.Error(err)
		}
		glog.Info("---> Kafka Mesage Subscriber-> Received:", string(msg.Value))
	}
}
