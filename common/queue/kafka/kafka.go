package kafka

import "github.com/IBM/sarama"

type Config struct {
	Addr []string
}

func NewKafkaAsyncProducer(conf *Config) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Addr, c)
	if err != nil {
		panic(err)
	}
	return p
}

func NewKafkaSyncProducer(conf *Config) sarama.SyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewSyncProducer(conf.Addr, c)
	if err != nil {
		panic(err)
	}
	return p
}

func NewKafkaConsumer(conf *Config) sarama.Consumer {
	c := sarama.NewConfig()
	p, err := sarama.NewConsumer(conf.Addr, c)
	if err != nil {
		panic(err)
	}
	return p
}
