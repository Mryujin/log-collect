package kafka

import (
	"time"
	"github.com/Shopify/sarama"
)
var asyncProducer sarama.AsyncProducer

func init() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true   //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
    p, err := sarama.NewAsyncProducer(strings.Split("localhost:9092", ","), config)
    defer p.Close()
    if err != nil {
        return
    }

    //必须有这个匿名函数内容
    go func(p sarama.AsyncProducer) {
        errors := p.Errors()
        success := p.Successes()
        for {
            select {
            case err := <-errors:
                if err != nil {
                    glog.Errorln(err)
                }
            case <-success:
            }
        }
	}(p)
	
	asyncProducer = p
}

func SendMessage(topic, message string) {
	msg := &sarama.ProducerMessage{
        Topic: topics,
        Value: sarama.ByteEncoder(v),
    }
    asyncProducer.Input() <- msg
}

