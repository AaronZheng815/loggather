package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AaronZheng815/loggather/log_agent/confData"

	"github.com/AaronZheng815/loggather/log_agent/kafkaAgent"

	"github.com/AaronZheng815/loggather/log_agent/tailf"

	"github.com/astaxie/beego/logs"
)

func main() {

	filename := "./conf/logagent.conf"
	err := loadConfig("ini", filename)
	if err != nil {
		fmt.Println("load config file failed")
		panic("load config file failed")
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
	}

	logs.Debug("Initialize logger Success")

	err = kafkaAgent.KafkaAgentInit()
	if err != nil {
		fmt.Printf("load kafka Agent failed.")
		panic("load kafka agent failed")
	}

	err = tailf.InitTailf(confData.AppConfig.CConf, confData.AppConfig.ChanSize)
	if err != nil {
		logs.Error("Failed to initialize tailf, err:", err)
		return
	}

	err = kafkaAgent.KafkaAgentInit()
	if err != nil {
		logs.Error("Faile to initailize kafka agent, err", err)
		return
	}

	go func() {
		for {
			logs.Debug("Hello golang logger")
			time.Sleep(time.Second)
		}
	}()

	err = ServiceRun()
	if err != nil {
		log.Fatal("Failed to run log agent service")
		return
	}

	logs.Info("log agent exited.")
}
