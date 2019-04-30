package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	saramaConfig *sarama.Config
	brokers      []string
	topics       string
	producer     sarama.SyncProducer
	msg          *sarama.ProducerMessage
	MsgPool      []string
)


func init(){
	saramaConfig = sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 5
	saramaConfig.Producer.Return.Successes = true
	brokers, topics = []string{"192.168.11.145:9092"},"test"
	var err error
	producer, err = sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
	panic(err)
	}

	msg = &sarama.ProducerMessage{
	Topic: topics,
	//Partition:  //todo: 指定partition
	//Value: sarama.StringEncoder("Something Cool"),
	}
}

func Producer(m string) error {
	//defer producer.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic,%v", r)
		}
	}()
	msg.Value = sarama.StringEncoder(m)
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("Send Message `%s`  Fail", m)
	}
	//fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topics, partition, offset)
	return err
}
