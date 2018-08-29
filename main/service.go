package main

import (
	"fmt"
	"time"

	"github.com/AaronZheng815/loggather/log_agent/kafkaAgent"

	"github.com/AaronZheng815/loggather/log_agent/tailf"
	"github.com/astaxie/beego/logs"
)

func ServiceRun() (err error) {
	for {
		msgTxt := tailf.GetOneLine()
		err := sendToKafka(msgTxt)
		if err != nil {
			logs.Error("failed to send to kafka!")
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	fmt.Printf("\nread msg: %s, topic:%s\n", msg.Msg, msg.Topic)
	kafkaAgent.Log2kafka(msg.Msg, msg.Topic)
	return
}
