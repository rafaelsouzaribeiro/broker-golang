package consumer

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func ListenPartition(broker *[]string, data *payload.Message, message chan<- payload.Message) {

	consumer, err := sarama.NewConsumer(*broker, GetConfig())

	if err != nil {
		panic(err)
	}

	pc, err := consumer.ConsumePartition(data.Topic, data.Partition, data.Offset)

	if err != nil {
		panic(err)
	}

	for msgs := range pc.Messages() {
		var listHeaders []payload.Header

		for _, h := range msgs.Headers {
			header := payload.Header{Key: string(h.Key), Value: string(h.Value)}
			listHeaders = append(listHeaders, header)
		}

		message <- payload.Message{
			Topic:     msgs.Topic,
			GroupID:   data.GroupID,
			Value:     string(msgs.Value),
			Key:       string(msgs.Key),
			Partition: msgs.Partition,
			Headers:   &listHeaders,
			Time:      msgs.Timestamp,
			Offset:    msgs.Offset,
		}

	}

	pc.Close()

}
