package main

import (
	"errors"
	"fmt"

	"github.com/AaronZheng815/loggather/log_agent/confData"
	"github.com/astaxie/beego/config"
)

func loadCollectConf(conf config.Configer) (err error) {

	var cc confData.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect - path")
		return
	}

	cc.Topic = conf.String("collect::log_topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect - topic")
		return
	}

	//保持到切片中
	confData.AppConfig.CConf = append(confData.AppConfig.CConf, cc)

	return
}

func loadConfig(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	//保持配置信息
	confData.AppConfig = &confData.Config{}

	confData.AppConfig.LogLevel = conf.String("log::log_level")
	if len(confData.AppConfig.LogLevel) == 0 {
		confData.AppConfig.LogLevel = "Debug"
	}

	confData.AppConfig.LogPath = conf.String("log::log_path")
	if len(confData.AppConfig.LogLevel) == 0 {
		confData.AppConfig.LogLevel = "./logs"
	}

	confData.AppConfig.ChanSize, err = conf.Int("log::channel_size")
	if err != nil {
		err = errors.New("invalid collect - channel size")
		confData.AppConfig.ChanSize = 100 //默认值100
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Println("load collect config failed")
		return
	}

	//保持kafka配置信息
	confData.AppConfig.KfkSerIp = conf.String("kafka::server_ip")
	if len(confData.AppConfig.KfkSerIp) == 0 {
		confData.AppConfig.KfkSerIp = "192.168.126.5"
	}
	confData.AppConfig.KfkSerPort, err = conf.Int("kafka::server_port")
	if err != nil {
		fmt.Println("failed to get kafka port from config file,set to default.")
		confData.AppConfig.KfkSerPort = 9092
	}

	return
}
