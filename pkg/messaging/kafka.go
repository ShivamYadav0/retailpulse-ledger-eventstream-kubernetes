package messaging

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

type KafkaConsumer struct {
	reader *kafka.Reader
}

// EnsureTopicExists blocks until the topic is created and ready on the cluster
func EnsureTopicExists(brokers []string, topic string, partitions int, replicationFactor int) error {
	if len(brokers) == 0 {
		return fmt.Errorf("no kafka brokers provided")
	}

	// Retry loop for cluster availability
	for i := 0; i < 10; i++ {
		conn, err := kafka.DialContext(context.Background(), "tcp", brokers[0])
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		
		controller, err := conn.Controller()
		if err != nil {
			conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		controllerAddr := net.JoinHostPort(controller.Host, fmt.Sprint(controller.Port))
		controllerConn, err := kafka.DialContext(context.Background(), "tcp", controllerAddr)
		if err != nil {
			conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		err = controllerConn.CreateTopics(kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		})
		
		controllerConn.Close()
		conn.Close()

		if err == nil || strings.Contains(err.Error(), "TopicExists") {
			return nil // Topic is ready
		}

		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("timed out waiting for topic %s to initialize", topic)
}

func NewProducer(brokers []string, topic string) (*KafkaProducer, error) {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			BatchTimeout: 10 * time.Millisecond,
		},
	}, nil
}

func (p *KafkaProducer) Produce(ctx context.Context, key string, payload []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: payload,
		Time:  time.Now(),
	})
}
func NewConsumer(brokers []string, groupID, topic string) *KafkaConsumer {
    return &KafkaConsumer{
        reader: kafka.NewReader(kafka.ReaderConfig{
            Brokers:  brokers,
            GroupID:  groupID,
            Topic:    topic,
            MinBytes: 10e3,
            MaxBytes: 10e6,
        }),
    }
}


func (p *KafkaProducer) Close() error {
	if p.writer == nil {
		return nil
	}
	return p.writer.Close()
}

func NewConsumerWithOffset(brokers []string, groupID, topic string, startOffset int64) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:          brokers,
			GroupID:          groupID,
			Topic:            topic,
			MinBytes:         1,
			MaxBytes:         10e6,
			StartOffset:      startOffset,
			
		}),
	}
}

func (c *KafkaConsumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return c.reader.ReadMessage(ctx)
}

func (c *KafkaConsumer) Close() error {
	if c.reader == nil {
		return nil
	}
	return c.reader.Close()
}
func (c *KafkaConsumer) CommitMessages(ctx context.Context, msgs ...kafka.Message) error {
    return c.reader.CommitMessages(ctx, msgs...)
}
