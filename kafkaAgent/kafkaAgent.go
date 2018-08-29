package kafkaAgent

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/AaronZheng815/loggather/log_agent/confData"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Log2kafka(msg string, topic string) (err error) {

	kmsg := &sarama.ProducerMessage{}
	kmsg.Topic = topic
	kmsg.Value = sarama.StringEncoder(msg)

	// pid, offset, err := client.SendMessage(kmsg)
	_, _, err = client.SendMessage(kmsg)
	if err != nil {
		logs.Error("Send kafka message failed. err:", err)
		return
	}
	//fmt.Printf("\nSuccess pid:%v, offset:%v\n", pid, offset)
	return
}

// func Log2kafka(msg string) (err error) {

// 	kmsg := &sarama.ProducerMessage{}
// 	kmsg.Topic = confData.AppConfig.CConf[0].Topic
// 	kmsg.Value = sarama.StringEncoder(msg)

// 	pid, offset, err := client.SendMessage(kmsg)
// 	if err != nil {
// 		logs.Error("Send kafka message failed. err:", err)
// 		return
// 	}

// 	fmt.Printf("\nSuccess pid:%v, offset:%v\n", pid, offset)
// 	return
// }
func KafkaAgentInit() (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	kafaAddr := confData.AppConfig.KfkSerIp + ":" + fmt.Sprintf("%d", confData.AppConfig.KfkSerPort)
	fmt.Println("kafka address: ", kafaAddr)

	client, err = sarama.NewSyncProducer([]string{kafaAddr}, config)
	if err != nil {
		logs.Error("New kafka producer failed, err:", err)
		return
	}

	// defer clt.Close()
	logs.Debug("Initial kafka agent success!")
	return
}
