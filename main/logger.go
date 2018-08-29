package main

import (
	"encoding/json"
	"fmt"

	"github.com/AaronZheng815/loggather/log_agent/confData"
	"github.com/astaxie/beego/logs"
)

func convLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogger() (err error) {

	config := make(map[string]interface{})
	config["filename"] = confData.AppConfig.LogPath
	config["level"] = convLogLevel(confData.AppConfig.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("init logger failed, marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	return
}
